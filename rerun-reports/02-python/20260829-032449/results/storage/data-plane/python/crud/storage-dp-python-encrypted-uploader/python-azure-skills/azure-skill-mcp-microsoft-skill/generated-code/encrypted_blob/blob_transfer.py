"""AES-GCM encrypted Azure Blob upload and download operations."""

from __future__ import annotations

import asyncio
import base64
import binascii
import os
import tempfile
from dataclasses import dataclass
from pathlib import Path
from typing import Mapping

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from .key_management import AsyncKeyManager, SyncKeyManager

AES_GCM_NONCE_SIZE_BYTES = 12
ENCRYPTION_ALGORITHM = "AES-256-GCM"
FORMAT_VERSION = "1"
KEY_WRAP_ALGORITHM_NAME = "RSA-OAEP-256"

_METADATA_VERSION = "enc_version"
_METADATA_ALGORITHM = "enc_algorithm"
_METADATA_KEY_WRAP_ALGORITHM = "enc_key_wrap_algorithm"
_METADATA_KEY_ID = "enc_key_id"
_METADATA_WRAPPED_KEY = "enc_wrapped_key"
_METADATA_NONCE = "enc_nonce"


class BlobEncryptionError(RuntimeError):
    """Raised when encrypted blob transfer or decryption fails."""


class InvalidBlobMetadataError(BlobEncryptionError):
    """Raised when required encryption metadata is missing or malformed."""


@dataclass(frozen=True)
class UploadResult:
    key_id: str
    wrapped_key_base64: str


@dataclass(frozen=True)
class _EncryptionMetadata:
    key_id: str
    wrapped_key: bytes
    nonce: bytes


def _encode_base64(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _decode_base64(name: str, value: str) -> bytes:
    try:
        return base64.b64decode(value, validate=True)
    except (binascii.Error, ValueError) as error:
        raise InvalidBlobMetadataError(
            f"Blob encryption metadata {name!r} is not valid base64"
        ) from error


def _parse_metadata(metadata: Mapping[str, str] | None) -> _EncryptionMetadata:
    if metadata is None:
        raise InvalidBlobMetadataError("Blob has no encryption metadata")

    required_names = (
        _METADATA_VERSION,
        _METADATA_ALGORITHM,
        _METADATA_KEY_WRAP_ALGORITHM,
        _METADATA_KEY_ID,
        _METADATA_WRAPPED_KEY,
        _METADATA_NONCE,
    )
    missing = [name for name in required_names if not metadata.get(name)]
    if missing:
        raise InvalidBlobMetadataError(
            "Blob is missing encryption metadata: " + ", ".join(missing)
        )
    if metadata[_METADATA_VERSION] != FORMAT_VERSION:
        raise InvalidBlobMetadataError(
            f"Unsupported encrypted blob format {metadata[_METADATA_VERSION]!r}"
        )
    if metadata[_METADATA_ALGORITHM] != ENCRYPTION_ALGORITHM:
        raise InvalidBlobMetadataError("Blob uses an unsupported data algorithm")
    if metadata[_METADATA_KEY_WRAP_ALGORITHM] != KEY_WRAP_ALGORITHM_NAME:
        raise InvalidBlobMetadataError("Blob uses an unsupported key-wrap algorithm")

    wrapped_key = _decode_base64(
        _METADATA_WRAPPED_KEY, metadata[_METADATA_WRAPPED_KEY]
    )
    nonce = _decode_base64(_METADATA_NONCE, metadata[_METADATA_NONCE])
    if not wrapped_key:
        raise InvalidBlobMetadataError("Blob contains an empty wrapped data key")
    if len(nonce) != AES_GCM_NONCE_SIZE_BYTES:
        raise InvalidBlobMetadataError("Blob contains an invalid AES-GCM nonce")
    return _EncryptionMetadata(
        key_id=metadata[_METADATA_KEY_ID],
        wrapped_key=wrapped_key,
        nonce=nonce,
    )


def _metadata_for_upload(
    key_id: str, wrapped_key: bytes, nonce: bytes
) -> dict[str, str]:
    return {
        _METADATA_VERSION: FORMAT_VERSION,
        _METADATA_ALGORITHM: ENCRYPTION_ALGORITHM,
        _METADATA_KEY_WRAP_ALGORITHM: KEY_WRAP_ALGORITHM_NAME,
        _METADATA_KEY_ID: key_id,
        _METADATA_WRAPPED_KEY: _encode_base64(wrapped_key),
        _METADATA_NONCE: _encode_base64(nonce),
    }


def _storage_error(operation: str, blob_name: str, error: HttpResponseError) -> str:
    status = f"HTTP {error.status_code}" if error.status_code else "an HTTP error"
    return f"Blob Storage {operation} failed for {blob_name!r} ({status})"


def _decrypt(
    ciphertext: bytes, metadata: _EncryptionMetadata, plaintext_key: bytearray
) -> bytes:
    try:
        return AESGCM(bytes(plaintext_key)).decrypt(
            metadata.nonce, ciphertext, None
        )
    except InvalidTag as error:
        raise BlobEncryptionError(
            "AES-GCM authentication failed; the blob or metadata was altered"
        ) from error
    finally:
        for index in range(len(plaintext_key)):
            plaintext_key[index] = 0


def _atomic_write(path: Path, data: bytes) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    temporary_name: str | None = None
    try:
        with tempfile.NamedTemporaryFile(
            mode="wb", dir=path.parent, delete=False
        ) as temporary_file:
            temporary_file.write(data)
            temporary_name = temporary_file.name
        os.replace(temporary_name, path)
    finally:
        if temporary_name and os.path.exists(temporary_name):
            os.unlink(temporary_name)


class SyncEncryptedBlobClient:
    def __init__(
        self,
        blob_service: BlobServiceClient,
        key_manager: SyncKeyManager,
        container_name: str,
    ) -> None:
        self._blob_service = blob_service
        self._key_manager = key_manager
        self._container_name = container_name

    def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = True
    ) -> UploadResult:
        material = self._key_manager.generate_and_wrap_data_key()
        nonce = os.urandom(AES_GCM_NONCE_SIZE_BYTES)
        try:
            ciphertext = AESGCM(bytes(material.plaintext_key)).encrypt(
                nonce, plaintext, None
            )
            metadata = _metadata_for_upload(
                material.key_id, material.wrapped_key, nonce
            )
            with self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            ) as blob_client:
                blob_client.upload_blob(
                    ciphertext,
                    overwrite=overwrite,
                    metadata=metadata,
                )
            return UploadResult(
                key_id=material.key_id,
                wrapped_key_base64=metadata[_METADATA_WRAPPED_KEY],
            )
        except HttpResponseError as error:
            raise BlobEncryptionError(
                _storage_error("upload", blob_name, error)
            ) from error
        finally:
            material.destroy()

    def download_bytes(self, blob_name: str) -> bytes:
        try:
            with self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            ) as blob_client:
                downloader = blob_client.download_blob()
                ciphertext = downloader.readall()
                metadata = _parse_metadata(downloader.properties.metadata)
        except ResourceNotFoundError as error:
            raise BlobEncryptionError(
                f"Blob {blob_name!r} does not exist in container "
                f"{self._container_name!r}"
            ) from error
        except HttpResponseError as error:
            raise BlobEncryptionError(
                _storage_error("download", blob_name, error)
            ) from error

        plaintext_key = self._key_manager.unwrap_data_key(
            metadata.wrapped_key, metadata.key_id
        )
        return _decrypt(ciphertext, metadata, plaintext_key)

    def upload_file(
        self, source: str | Path, blob_name: str, *, overwrite: bool = True
    ) -> UploadResult:
        source_path = Path(source)
        return self.upload_bytes(
            blob_name, source_path.read_bytes(), overwrite=overwrite
        )

    def download_file(self, blob_name: str, destination: str | Path) -> bytes:
        plaintext = self.download_bytes(blob_name)
        _atomic_write(Path(destination), plaintext)
        return plaintext


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
        material = await self._key_manager.generate_and_wrap_data_key()
        nonce = os.urandom(AES_GCM_NONCE_SIZE_BYTES)
        try:
            ciphertext = AESGCM(bytes(material.plaintext_key)).encrypt(
                nonce, plaintext, None
            )
            metadata = _metadata_for_upload(
                material.key_id, material.wrapped_key, nonce
            )
            async with self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            ) as blob_client:
                await blob_client.upload_blob(
                    ciphertext,
                    overwrite=overwrite,
                    metadata=metadata,
                )
            return UploadResult(
                key_id=material.key_id,
                wrapped_key_base64=metadata[_METADATA_WRAPPED_KEY],
            )
        except HttpResponseError as error:
            raise BlobEncryptionError(
                _storage_error("upload", blob_name, error)
            ) from error
        finally:
            material.destroy()

    async def download_bytes(self, blob_name: str) -> bytes:
        try:
            async with self._blob_service.get_blob_client(
                container=self._container_name, blob=blob_name
            ) as blob_client:
                downloader = await blob_client.download_blob()
                ciphertext = await downloader.readall()
                metadata = _parse_metadata(downloader.properties.metadata)
        except ResourceNotFoundError as error:
            raise BlobEncryptionError(
                f"Blob {blob_name!r} does not exist in container "
                f"{self._container_name!r}"
            ) from error
        except HttpResponseError as error:
            raise BlobEncryptionError(
                _storage_error("download", blob_name, error)
            ) from error

        plaintext_key = await self._key_manager.unwrap_data_key(
            metadata.wrapped_key, metadata.key_id
        )
        return _decrypt(ciphertext, metadata, plaintext_key)

    async def upload_file(
        self, source: str | Path, blob_name: str, *, overwrite: bool = True
    ) -> UploadResult:
        source_path = Path(source)
        plaintext = await asyncio.to_thread(source_path.read_bytes)
        return await self.upload_bytes(
            blob_name, plaintext, overwrite=overwrite
        )

    async def download_file(
        self, blob_name: str, destination: str | Path
    ) -> bytes:
        plaintext = await self.download_bytes(blob_name)
        await asyncio.to_thread(_atomic_write, Path(destination), plaintext)
        return plaintext
