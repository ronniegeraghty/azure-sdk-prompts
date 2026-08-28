"""AES-GCM encrypted uploads and downloads for Azure Blob Storage."""

from __future__ import annotations

import asyncio
import base64
import binascii
import os
from dataclasses import dataclass
from pathlib import Path
from typing import Mapping

from azure.core import MatchConditions
from azure.core.exceptions import (
    AzureError,
    ResourceModifiedError,
    ResourceNotFoundError,
)
from azure.storage.blob import ContainerClient
from azure.storage.blob.aio import ContainerClient as AsyncContainerClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from .key_management import AsyncKeyManager, KeyManager, WrappedDataKey

CONTENT_ENCRYPTION_ALGORITHM = "AES-256-GCM"
ENVELOPE_VERSION = "1"
NONCE_BYTES = 12

_METADATA_VERSION = "encryption_version"
_METADATA_CONTENT_ALGORITHM = "content_encryption"
_METADATA_KEY_WRAP_ALGORITHM = "key_wrap_algorithm"
_METADATA_KEY_ID = "key_id"
_METADATA_WRAPPED_KEY = "wrapped_data_key"
_METADATA_NONCE = "nonce"


class EncryptedBlobError(RuntimeError):
    """Base error for encrypted blob operations."""


class EncryptedBlobNotFoundError(EncryptedBlobError):
    """Raised when an encrypted blob does not exist."""


class EncryptedBlobMetadataError(EncryptedBlobError):
    """Raised when encryption metadata is absent, malformed, or unsupported."""


class BlobStorageError(EncryptedBlobError):
    """Raised when Blob Storage rejects an operation."""


class BlobDecryptionError(EncryptedBlobError):
    """Raised when AES-GCM authentication fails."""


@dataclass(frozen=True)
class UploadResult:
    """Details of the envelope used for an uploaded blob."""

    blob_name: str
    key_id: str
    wrapped_data_key_base64: str


@dataclass(frozen=True)
class _EnvelopeMetadata:
    wrapped_data_key: WrappedDataKey
    nonce: bytes

    def to_blob_metadata(self) -> dict[str, str]:
        return {
            _METADATA_VERSION: ENVELOPE_VERSION,
            _METADATA_CONTENT_ALGORITHM: CONTENT_ENCRYPTION_ALGORITHM,
            _METADATA_KEY_WRAP_ALGORITHM: self.wrapped_data_key.algorithm,
            _METADATA_KEY_ID: self.wrapped_data_key.key_id,
            _METADATA_WRAPPED_KEY: _base64_encode(
                self.wrapped_data_key.wrapped_key
            ),
            _METADATA_NONCE: _base64_encode(self.nonce),
        }


def _base64_encode(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _base64_decode(value: str, field_name: str) -> bytes:
    try:
        return base64.b64decode(value, validate=True)
    except (binascii.Error, ValueError) as exc:
        raise EncryptedBlobMetadataError(
            f"Blob metadata field {field_name!r} is not valid base64."
        ) from exc


def _associated_data(key_id: str, key_wrap_algorithm: str) -> bytes:
    return (
        f"version={ENVELOPE_VERSION}\n"
        f"content={CONTENT_ENCRYPTION_ALGORITHM}\n"
        f"key_wrap={key_wrap_algorithm}\n"
        f"key_id={key_id}"
    ).encode("utf-8")


def _parse_metadata(metadata: Mapping[str, str] | None) -> _EnvelopeMetadata:
    if not metadata:
        raise EncryptedBlobMetadataError(
            "Blob has no client-side encryption metadata."
        )

    normalized = {key.lower(): value for key, value in metadata.items()}
    required_fields = (
        _METADATA_VERSION,
        _METADATA_CONTENT_ALGORITHM,
        _METADATA_KEY_WRAP_ALGORITHM,
        _METADATA_KEY_ID,
        _METADATA_WRAPPED_KEY,
        _METADATA_NONCE,
    )
    missing = [field for field in required_fields if not normalized.get(field)]
    if missing:
        raise EncryptedBlobMetadataError(
            f"Blob encryption metadata is missing: {', '.join(missing)}."
        )
    if normalized[_METADATA_VERSION] != ENVELOPE_VERSION:
        raise EncryptedBlobMetadataError(
            f"Unsupported envelope version {normalized[_METADATA_VERSION]!r}."
        )
    if (
        normalized[_METADATA_CONTENT_ALGORITHM]
        != CONTENT_ENCRYPTION_ALGORITHM
    ):
        raise EncryptedBlobMetadataError(
            "Unsupported content-encryption algorithm "
            f"{normalized[_METADATA_CONTENT_ALGORITHM]!r}."
        )

    nonce = _base64_decode(normalized[_METADATA_NONCE], _METADATA_NONCE)
    if len(nonce) != NONCE_BYTES:
        raise EncryptedBlobMetadataError(
            f"AES-GCM nonce must be {NONCE_BYTES} bytes; got {len(nonce)}."
        )
    wrapped_key = _base64_decode(
        normalized[_METADATA_WRAPPED_KEY], _METADATA_WRAPPED_KEY
    )
    if not wrapped_key:
        raise EncryptedBlobMetadataError("Wrapped data key is empty.")

    return _EnvelopeMetadata(
        wrapped_data_key=WrappedDataKey(
            key_id=normalized[_METADATA_KEY_ID],
            algorithm=normalized[_METADATA_KEY_WRAP_ALGORITHM],
            wrapped_key=wrapped_key,
        ),
        nonce=nonce,
    )


def _encrypt(
    plaintext: bytes, key_manager: KeyManager
) -> tuple[bytes, _EnvelopeMetadata]:
    data_key, wrapped_data_key = key_manager.generate_and_wrap_data_key()
    nonce = os.urandom(NONCE_BYTES)
    try:
        ciphertext = AESGCM(data_key).encrypt(
            nonce,
            plaintext,
            _associated_data(
                wrapped_data_key.key_id, wrapped_data_key.algorithm
            ),
        )
    finally:
        del data_key
    return ciphertext, _EnvelopeMetadata(wrapped_data_key, nonce)


def _decrypt(
    ciphertext: bytes,
    envelope: _EnvelopeMetadata,
    key_manager: KeyManager,
) -> bytes:
    data_key = key_manager.unwrap_data_key(envelope.wrapped_data_key)
    try:
        return AESGCM(data_key).decrypt(
            envelope.nonce,
            ciphertext,
            _associated_data(
                envelope.wrapped_data_key.key_id,
                envelope.wrapped_data_key.algorithm,
            ),
        )
    except InvalidTag as exc:
        raise BlobDecryptionError(
            "AES-GCM authentication failed. The ciphertext or its encryption "
            "metadata may have been modified."
        ) from exc
    finally:
        del data_key


async def _encrypt_async(
    plaintext: bytes, key_manager: AsyncKeyManager
) -> tuple[bytes, _EnvelopeMetadata]:
    data_key, wrapped_data_key = await key_manager.generate_and_wrap_data_key()
    nonce = os.urandom(NONCE_BYTES)
    try:
        ciphertext = AESGCM(data_key).encrypt(
            nonce,
            plaintext,
            _associated_data(
                wrapped_data_key.key_id, wrapped_data_key.algorithm
            ),
        )
    finally:
        del data_key
    return ciphertext, _EnvelopeMetadata(wrapped_data_key, nonce)


async def _decrypt_async(
    ciphertext: bytes,
    envelope: _EnvelopeMetadata,
    key_manager: AsyncKeyManager,
) -> bytes:
    data_key = await key_manager.unwrap_data_key(envelope.wrapped_data_key)
    try:
        return AESGCM(data_key).decrypt(
            envelope.nonce,
            ciphertext,
            _associated_data(
                envelope.wrapped_data_key.key_id,
                envelope.wrapped_data_key.algorithm,
            ),
        )
    except InvalidTag as exc:
        raise BlobDecryptionError(
            "AES-GCM authentication failed. The ciphertext or its encryption "
            "metadata may have been modified."
        ) from exc
    finally:
        del data_key


class EncryptedBlobClient:
    """Synchronous encrypted blob upload and download operations."""

    def __init__(
        self, container_client: ContainerClient, key_manager: KeyManager
    ) -> None:
        self._container_client = container_client
        self._key_manager = key_manager

    def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = True
    ) -> UploadResult:
        ciphertext, envelope = _encrypt(plaintext, self._key_manager)
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            blob_client.upload_blob(
                ciphertext,
                metadata=envelope.to_blob_metadata(),
                overwrite=overwrite,
            )
        except AzureError as exc:
            raise BlobStorageError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

        wrapped_key_base64 = _base64_encode(
            envelope.wrapped_data_key.wrapped_key
        )
        return UploadResult(
            blob_name=blob_name,
            key_id=envelope.wrapped_data_key.key_id,
            wrapped_data_key_base64=wrapped_key_base64,
        )

    def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            properties = blob_client.get_blob_properties()
            ciphertext = blob_client.download_blob(
                etag=properties.etag,
                match_condition=MatchConditions.IfNotModified,
            ).readall()
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist."
            ) from exc
        except ResourceModifiedError as exc:
            raise BlobStorageError(
                f"Encrypted blob {blob_name!r} changed while it was downloading."
            ) from exc
        except AzureError as exc:
            raise BlobStorageError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        envelope = _parse_metadata(properties.metadata)
        return _decrypt(ciphertext, envelope, self._key_manager)

    def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = True
    ) -> UploadResult:
        return self.upload_bytes(
            blob_name, Path(source).read_bytes(), overwrite=overwrite
        )

    def download_to_file(self, blob_name: str, destination: str | Path) -> None:
        Path(destination).write_bytes(self.download_bytes(blob_name))


class AsyncEncryptedBlobClient:
    """Asynchronous encrypted blob upload and download operations."""

    def __init__(
        self,
        container_client: AsyncContainerClient,
        key_manager: AsyncKeyManager,
    ) -> None:
        self._container_client = container_client
        self._key_manager = key_manager

    async def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = True
    ) -> UploadResult:
        ciphertext, envelope = await _encrypt_async(
            plaintext, self._key_manager
        )
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            await blob_client.upload_blob(
                ciphertext,
                metadata=envelope.to_blob_metadata(),
                overwrite=overwrite,
            )
        except AzureError as exc:
            raise BlobStorageError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

        wrapped_key_base64 = _base64_encode(
            envelope.wrapped_data_key.wrapped_key
        )
        return UploadResult(
            blob_name=blob_name,
            key_id=envelope.wrapped_data_key.key_id,
            wrapped_data_key_base64=wrapped_key_base64,
        )

    async def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            properties = await blob_client.get_blob_properties()
            stream = await blob_client.download_blob(
                etag=properties.etag,
                match_condition=MatchConditions.IfNotModified,
            )
            ciphertext = await stream.readall()
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist."
            ) from exc
        except ResourceModifiedError as exc:
            raise BlobStorageError(
                f"Encrypted blob {blob_name!r} changed while it was downloading."
            ) from exc
        except AzureError as exc:
            raise BlobStorageError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        envelope = _parse_metadata(properties.metadata)
        return await _decrypt_async(ciphertext, envelope, self._key_manager)

    async def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = True
    ) -> UploadResult:
        plaintext = await asyncio.to_thread(Path(source).read_bytes)
        return await self.upload_bytes(blob_name, plaintext, overwrite=overwrite)

    async def download_to_file(
        self, blob_name: str, destination: str | Path
    ) -> None:
        plaintext = await self.download_bytes(blob_name)
        await asyncio.to_thread(Path(destination).write_bytes, plaintext)
