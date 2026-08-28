from __future__ import annotations

import asyncio
import base64
import binascii
import os
from dataclasses import dataclass
from pathlib import Path
from typing import Mapping

from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.storage.blob import ContainerClient
from azure.storage.blob.aio import ContainerClient as AsyncContainerClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from key_management import AsyncKeyManager, KeyManager, WrappedDataKey

AES_NONCE_SIZE_BYTES = 12
ENCRYPTION_ALGORITHM = "AES-256-GCM"
METADATA_VERSION = "1"


class BlobEncryptionError(RuntimeError):
    """Raised when encrypted blob storage or cryptographic processing fails."""


class EncryptedBlobNotFoundError(BlobEncryptionError):
    """Raised when the requested encrypted blob does not exist."""


@dataclass(frozen=True)
class UploadResult:
    blob_name: str
    key_id: str
    wrapped_data_key: bytes


def _b64encode(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _b64decode(metadata: Mapping[str, str], name: str) -> bytes:
    value = metadata.get(name)
    if not value:
        raise BlobEncryptionError(f"Encrypted blob metadata is missing {name!r}")
    try:
        return base64.b64decode(value, validate=True)
    except (ValueError, binascii.Error) as exc:
        raise BlobEncryptionError(
            f"Encrypted blob metadata {name!r} is not valid base64"
        ) from exc


def _encryption_metadata(wrapped_key: WrappedDataKey, nonce: bytes) -> dict[str, str]:
    return {
        "encryption_version": METADATA_VERSION,
        "encryption_algorithm": ENCRYPTION_ALGORITHM,
        "key_wrap_algorithm": wrapped_key.algorithm,
        "key_id": wrapped_key.key_id,
        "wrapped_data_key": _b64encode(wrapped_key.encrypted_key),
        "nonce": _b64encode(nonce),
    }


def _parse_metadata(metadata: Mapping[str, str]) -> tuple[WrappedDataKey, bytes]:
    if metadata.get("encryption_version") != METADATA_VERSION:
        raise BlobEncryptionError("Encrypted blob has an unsupported metadata version")
    if metadata.get("encryption_algorithm") != ENCRYPTION_ALGORITHM:
        raise BlobEncryptionError("Encrypted blob uses an unsupported cipher")

    key_id = metadata.get("key_id")
    key_wrap_algorithm = metadata.get("key_wrap_algorithm")
    if not key_id or not key_wrap_algorithm:
        raise BlobEncryptionError("Encrypted blob metadata is incomplete")

    nonce = _b64decode(metadata, "nonce")
    if len(nonce) != AES_NONCE_SIZE_BYTES:
        raise BlobEncryptionError("Encrypted blob contains an invalid AES-GCM nonce")

    return (
        WrappedDataKey(
            key_id=key_id,
            encrypted_key=_b64decode(metadata, "wrapped_data_key"),
            algorithm=key_wrap_algorithm,
        ),
        nonce,
    )


def _decrypt(ciphertext: bytes, data_key: bytes, nonce: bytes) -> bytes:
    try:
        return AESGCM(data_key).decrypt(nonce, ciphertext, None)
    except (InvalidTag, ValueError) as exc:
        raise BlobEncryptionError(
            "Blob decryption failed; the ciphertext, key, or metadata is invalid"
        ) from exc


class EncryptedBlobClient:
    def __init__(
        self,
        container_client: ContainerClient,
        key_manager: KeyManager,
    ) -> None:
        self._container_client = container_client
        self._key_manager = key_manager

    def upload_bytes(
        self,
        blob_name: str,
        plaintext: bytes,
        *,
        overwrite: bool = False,
    ) -> UploadResult:
        data_key, wrapped_key = self._key_manager.generate_and_wrap_data_key()
        nonce = os.urandom(AES_NONCE_SIZE_BYTES)
        ciphertext = AESGCM(data_key).encrypt(nonce, plaintext, None)
        metadata = _encryption_metadata(wrapped_key, nonce)

        try:
            self._container_client.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=metadata,
                overwrite=overwrite,
            )
        except HttpResponseError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

        return UploadResult(
            blob_name=blob_name,
            key_id=wrapped_key.key_id,
            wrapped_data_key=wrapped_key.encrypted_key,
        )

    def download_bytes(self, blob_name: str) -> bytes:
        try:
            downloader = self._container_client.download_blob(blob_name)
            ciphertext = downloader.readall()
            metadata = downloader.properties.metadata or {}
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist"
            ) from exc
        except HttpResponseError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        wrapped_key, nonce = _parse_metadata(metadata)
        data_key = self._key_manager.unwrap_data_key(wrapped_key)
        return _decrypt(ciphertext, data_key, nonce)

    def upload_file(
        self,
        source: str | Path,
        blob_name: str,
        *,
        overwrite: bool = False,
    ) -> UploadResult:
        try:
            plaintext = Path(source).read_bytes()
        except OSError as exc:
            raise BlobEncryptionError(f"Could not read {str(source)!r}: {exc}") from exc
        return self.upload_bytes(blob_name, plaintext, overwrite=overwrite)

    def download_file(self, blob_name: str, destination: str | Path) -> None:
        plaintext = self.download_bytes(blob_name)
        try:
            Path(destination).write_bytes(plaintext)
        except OSError as exc:
            raise BlobEncryptionError(
                f"Could not write {str(destination)!r}: {exc}"
            ) from exc


class AsyncEncryptedBlobClient:
    def __init__(
        self,
        container_client: AsyncContainerClient,
        key_manager: AsyncKeyManager,
    ) -> None:
        self._container_client = container_client
        self._key_manager = key_manager

    async def upload_bytes(
        self,
        blob_name: str,
        plaintext: bytes,
        *,
        overwrite: bool = False,
    ) -> UploadResult:
        data_key, wrapped_key = await self._key_manager.generate_and_wrap_data_key()
        nonce = os.urandom(AES_NONCE_SIZE_BYTES)
        ciphertext = AESGCM(data_key).encrypt(nonce, plaintext, None)
        metadata = _encryption_metadata(wrapped_key, nonce)

        try:
            await self._container_client.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=metadata,
                overwrite=overwrite,
            )
        except HttpResponseError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

        return UploadResult(
            blob_name=blob_name,
            key_id=wrapped_key.key_id,
            wrapped_data_key=wrapped_key.encrypted_key,
        )

    async def download_bytes(self, blob_name: str) -> bytes:
        try:
            downloader = await self._container_client.download_blob(blob_name)
            ciphertext = await downloader.readall()
            metadata = downloader.properties.metadata or {}
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist"
            ) from exc
        except HttpResponseError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        wrapped_key, nonce = _parse_metadata(metadata)
        data_key = await self._key_manager.unwrap_data_key(wrapped_key)
        return _decrypt(ciphertext, data_key, nonce)

    async def upload_file(
        self,
        source: str | Path,
        blob_name: str,
        *,
        overwrite: bool = False,
    ) -> UploadResult:
        try:
            plaintext = await asyncio.to_thread(Path(source).read_bytes)
        except OSError as exc:
            raise BlobEncryptionError(f"Could not read {str(source)!r}: {exc}") from exc
        return await self.upload_bytes(blob_name, plaintext, overwrite=overwrite)

    async def download_file(
        self,
        blob_name: str,
        destination: str | Path,
    ) -> None:
        plaintext = await self.download_bytes(blob_name)
        try:
            await asyncio.to_thread(Path(destination).write_bytes, plaintext)
        except OSError as exc:
            raise BlobEncryptionError(
                f"Could not write {str(destination)!r}: {exc}"
            ) from exc
