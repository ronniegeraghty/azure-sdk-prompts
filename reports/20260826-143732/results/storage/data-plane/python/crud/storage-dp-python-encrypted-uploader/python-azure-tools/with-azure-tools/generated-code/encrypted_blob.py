from __future__ import annotations

import base64
import binascii
import json
import os
from dataclasses import dataclass
from pathlib import Path
from typing import Mapping

from azure.core import MatchConditions
from azure.core.exceptions import AzureError, HttpResponseError, ResourceNotFoundError
from azure.storage.blob import BlobServiceClient, ContentSettings
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from key_management import (
    KEY_WRAP_ALGORITHM_NAME,
    AsyncKeyManager,
    KeyManager,
    erase_key,
)

FORMAT_VERSION = "1"
ENCRYPTION_ALGORITHM = "A256GCM"
NONCE_SIZE_BYTES = 12

META_VERSION = "ce_version"
META_ENCRYPTION = "ce_encryption"
META_KEY_WRAP = "ce_key_wrap"
META_NONCE = "ce_nonce"
META_WRAPPED_KEY = "ce_wrapped_key"
META_KEY_ID = "ce_key_id"


class EncryptedBlobError(RuntimeError):
    """Raised when encrypted blob storage or decryption fails."""


@dataclass(frozen=True)
class UploadResult:
    key_id: str
    wrapped_key: bytes


@dataclass(frozen=True)
class EncryptionMetadata:
    nonce: bytes
    wrapped_key: bytes
    key_id: str


def _b64encode(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _b64decode(value: str, field: str) -> bytes:
    try:
        return base64.b64decode(value, validate=True)
    except (binascii.Error, ValueError) as error:
        raise EncryptedBlobError(f"Blob metadata field {field!r} is invalid") from error


def _aad(container_name: str, blob_name: str, key_id: str) -> bytes:
    authenticated_context = {
        "blob": blob_name,
        "container": container_name,
        "encryption": ENCRYPTION_ALGORITHM,
        "key_id": key_id,
        "key_wrap": KEY_WRAP_ALGORITHM_NAME,
        "version": FORMAT_VERSION,
    }
    return json.dumps(
        authenticated_context, sort_keys=True, separators=(",", ":")
    ).encode("utf-8")


def _serialize_metadata(
    nonce: bytes, wrapped_key: bytes, key_id: str
) -> dict[str, str]:
    return {
        META_VERSION: FORMAT_VERSION,
        META_ENCRYPTION: ENCRYPTION_ALGORITHM,
        META_KEY_WRAP: KEY_WRAP_ALGORITHM_NAME,
        META_NONCE: _b64encode(nonce),
        META_WRAPPED_KEY: _b64encode(wrapped_key),
        META_KEY_ID: _b64encode(key_id.encode("utf-8")),
    }


def _parse_metadata(metadata: Mapping[str, str] | None) -> EncryptionMetadata:
    if not metadata:
        raise EncryptedBlobError("Blob has no client-side encryption metadata")

    try:
        version = metadata[META_VERSION]
        encryption = metadata[META_ENCRYPTION]
        key_wrap = metadata[META_KEY_WRAP]
        nonce = _b64decode(metadata[META_NONCE], META_NONCE)
        wrapped_key = _b64decode(metadata[META_WRAPPED_KEY], META_WRAPPED_KEY)
        key_id_bytes = _b64decode(metadata[META_KEY_ID], META_KEY_ID)
    except KeyError as error:
        raise EncryptedBlobError(
            f"Blob encryption metadata is missing {error.args[0]!r}"
        ) from error

    if (
        version != FORMAT_VERSION
        or encryption != ENCRYPTION_ALGORITHM
        or key_wrap != KEY_WRAP_ALGORITHM_NAME
    ):
        raise EncryptedBlobError(
            "Blob uses an unsupported client-side encryption format"
        )
    if len(nonce) != NONCE_SIZE_BYTES:
        raise EncryptedBlobError("Blob metadata contains an invalid AES-GCM nonce")
    if not wrapped_key:
        raise EncryptedBlobError("Blob metadata contains an empty wrapped data key")

    try:
        key_id = key_id_bytes.decode("utf-8")
    except UnicodeDecodeError as error:
        raise EncryptedBlobError("Blob metadata contains an invalid key ID") from error
    if not key_id:
        raise EncryptedBlobError("Blob metadata contains an empty key ID")

    return EncryptionMetadata(nonce, wrapped_key, key_id)


def _storage_error(
    operation: str, container_name: str, blob_name: str, error: AzureError
) -> EncryptedBlobError:
    status_code = getattr(error, "status_code", None)
    status = f" (HTTP {status_code})" if status_code else ""
    return EncryptedBlobError(
        f"Blob Storage could not {operation} "
        f"{container_name}/{blob_name}{status}"
    )


def _decrypt(
    ciphertext: bytes,
    metadata: EncryptionMetadata,
    data_key: bytearray,
    container_name: str,
    blob_name: str,
) -> bytes:
    try:
        return AESGCM(bytes(data_key)).decrypt(
            metadata.nonce,
            ciphertext,
            _aad(container_name, blob_name, metadata.key_id),
        )
    except InvalidTag as error:
        raise EncryptedBlobError(
            "AES-GCM authentication failed; the blob or its metadata was modified"
        ) from error


class EncryptedBlobClient:
    def __init__(
        self,
        blob_service: BlobServiceClient,
        key_manager: KeyManager,
        container_name: str,
    ) -> None:
        self._blob_service = blob_service
        self._key_manager = key_manager
        self._container_name = container_name

    def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = True
    ) -> UploadResult:
        envelope = self._key_manager.create_envelope_key()
        try:
            nonce = os.urandom(NONCE_SIZE_BYTES)
            ciphertext = AESGCM(bytes(envelope.plaintext_key)).encrypt(
                nonce,
                plaintext,
                _aad(self._container_name, blob_name, envelope.key_id),
            )
            metadata = _serialize_metadata(
                nonce, envelope.wrapped_key, envelope.key_id
            )
            blob_client = self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            )
            try:
                blob_client.upload_blob(
                    ciphertext,
                    overwrite=overwrite,
                    metadata=metadata,
                    content_settings=ContentSettings(
                        content_type="application/octet-stream"
                    ),
                )
            except ResourceNotFoundError as error:
                raise EncryptedBlobError(
                    f"Blob container {self._container_name!r} was not found"
                ) from error
            except HttpResponseError as error:
                raise _storage_error(
                    "upload", self._container_name, blob_name, error
                ) from error
            except AzureError as error:
                raise _storage_error(
                    "upload", self._container_name, blob_name, error
                ) from error
            return UploadResult(envelope.key_id, envelope.wrapped_key)
        finally:
            erase_key(envelope.plaintext_key)

    def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = True
    ) -> UploadResult:
        return self.upload_bytes(
            blob_name, Path(source).read_bytes(), overwrite=overwrite
        )

    def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._blob_service.get_blob_client(
            container=self._container_name, blob=blob_name
        )
        try:
            properties = blob_client.get_blob_properties()
            downloader = blob_client.download_blob(
                etag=properties.etag,
                match_condition=MatchConditions.IfNotModified,
            )
            ciphertext = downloader.readall()
        except ResourceNotFoundError as error:
            raise EncryptedBlobError(
                f"Blob {self._container_name}/{blob_name} was not found"
            ) from error
        except HttpResponseError as error:
            raise _storage_error(
                "download", self._container_name, blob_name, error
            ) from error
        except AzureError as error:
            raise _storage_error(
                "download", self._container_name, blob_name, error
            ) from error

        metadata = _parse_metadata(properties.metadata)
        data_key = self._key_manager.unwrap_key(
            metadata.wrapped_key, metadata.key_id
        )
        try:
            return _decrypt(
                ciphertext,
                metadata,
                data_key,
                self._container_name,
                blob_name,
            )
        finally:
            erase_key(data_key)

    def download_file(self, blob_name: str, destination: str | Path) -> None:
        Path(destination).write_bytes(self.download_bytes(blob_name))


class AsyncEncryptedBlobClient:
    def __init__(
        self,
        blob_service: AsyncBlobServiceClient,
        key_manager: AsyncKeyManager,
        container_name: str,
    ) -> None:
        self._blob_service = blob_service
        self._key_manager = key_manager
        self._container_name = container_name

    async def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = True
    ) -> UploadResult:
        envelope = await self._key_manager.create_envelope_key()
        try:
            nonce = os.urandom(NONCE_SIZE_BYTES)
            ciphertext = AESGCM(bytes(envelope.plaintext_key)).encrypt(
                nonce,
                plaintext,
                _aad(self._container_name, blob_name, envelope.key_id),
            )
            metadata = _serialize_metadata(
                nonce, envelope.wrapped_key, envelope.key_id
            )
            blob_client = self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            )
            try:
                await blob_client.upload_blob(
                    ciphertext,
                    overwrite=overwrite,
                    metadata=metadata,
                    content_settings=ContentSettings(
                        content_type="application/octet-stream"
                    ),
                )
            except ResourceNotFoundError as error:
                raise EncryptedBlobError(
                    f"Blob container {self._container_name!r} was not found"
                ) from error
            except HttpResponseError as error:
                raise _storage_error(
                    "upload", self._container_name, blob_name, error
                ) from error
            except AzureError as error:
                raise _storage_error(
                    "upload", self._container_name, blob_name, error
                ) from error
            return UploadResult(envelope.key_id, envelope.wrapped_key)
        finally:
            erase_key(envelope.plaintext_key)

    async def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = True
    ) -> UploadResult:
        return await self.upload_bytes(
            blob_name, Path(source).read_bytes(), overwrite=overwrite
        )

    async def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._blob_service.get_blob_client(
            container=self._container_name, blob=blob_name
        )
        try:
            properties = await blob_client.get_blob_properties()
            downloader = await blob_client.download_blob(
                etag=properties.etag,
                match_condition=MatchConditions.IfNotModified,
            )
            ciphertext = await downloader.readall()
        except ResourceNotFoundError as error:
            raise EncryptedBlobError(
                f"Blob {self._container_name}/{blob_name} was not found"
            ) from error
        except HttpResponseError as error:
            raise _storage_error(
                "download", self._container_name, blob_name, error
            ) from error
        except AzureError as error:
            raise _storage_error(
                "download", self._container_name, blob_name, error
            ) from error

        metadata = _parse_metadata(properties.metadata)
        data_key = await self._key_manager.unwrap_key(
            metadata.wrapped_key, metadata.key_id
        )
        try:
            return _decrypt(
                ciphertext,
                metadata,
                data_key,
                self._container_name,
                blob_name,
            )
        finally:
            erase_key(data_key)

    async def download_file(
        self, blob_name: str, destination: str | Path
    ) -> None:
        Path(destination).write_bytes(await self.download_bytes(blob_name))
