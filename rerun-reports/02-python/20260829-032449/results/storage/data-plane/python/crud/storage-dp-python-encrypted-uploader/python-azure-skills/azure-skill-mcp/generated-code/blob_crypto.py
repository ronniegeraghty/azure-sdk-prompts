from __future__ import annotations

import base64
import binascii
from dataclasses import dataclass
from pathlib import Path
from secrets import token_bytes
from typing import Mapping

from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.storage.blob import BlobServiceClient
from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from key_management import (
    AsyncKeyManager,
    KeyManager,
    KeyManagementError,
    ProtectedDataKey,
)

NONCE_SIZE_BYTES = 12
ENCRYPTION_VERSION = "1"
CONTENT_ALGORITHM = "A256GCM"
KEY_WRAP_ALGORITHM = "RSA-OAEP-256"


class BlobEncryptionError(RuntimeError):
    """Base error for encrypted blob operations."""


class EncryptedBlobNotFoundError(BlobEncryptionError):
    """Raised when an encrypted blob does not exist."""


class EncryptionMetadataError(BlobEncryptionError):
    """Raised when encryption metadata is absent or malformed."""


class CiphertextIntegrityError(BlobEncryptionError):
    """Raised when AES-GCM authentication fails."""


@dataclass(frozen=True)
class UploadResult:
    key_id: str
    wrapped_key_base64: str


class EncryptedBlobClient:
    def __init__(
        self,
        blob_service_client: BlobServiceClient,
        container_name: str,
        key_manager: KeyManager,
    ) -> None:
        self._container_client = blob_service_client.get_container_client(
            container_name
        )
        self._container_name = container_name
        self._key_manager = key_manager

    def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = False
    ) -> UploadResult:
        data_key = self._key_manager.generate_data_key()
        nonce = token_bytes(NONCE_SIZE_BYTES)
        ciphertext = AESGCM(data_key).encrypt(
            nonce, plaintext, self._associated_data(blob_name)
        )
        protected_key = self._key_manager.protect_data_key(data_key)
        metadata = _build_metadata(protected_key, nonce)

        try:
            self._container_client.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=metadata,
                overwrite=overwrite,
            )
        except AzureError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not upload {blob_name!r}"
            ) from exc

        return UploadResult(
            key_id=protected_key.key_id,
            wrapped_key_base64=metadata["wrapped_dek"],
        )

    def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = False
    ) -> UploadResult:
        return self.upload_bytes(
            blob_name, Path(source).read_bytes(), overwrite=overwrite
        )

    def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            properties = blob_client.get_blob_properties()
            ciphertext = blob_client.download_blob().readall()
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} was not found"
            ) from exc
        except AzureError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not download {blob_name!r}"
            ) from exc

        protected_key, nonce = _parse_metadata(properties.metadata)
        data_key = self._key_manager.recover_data_key(protected_key)
        try:
            return AESGCM(data_key).decrypt(
                nonce, ciphertext, self._associated_data(blob_name)
            )
        except InvalidTag as exc:
            raise CiphertextIntegrityError(
                f"Authentication failed for encrypted blob {blob_name!r}"
            ) from exc

    def download_file(self, blob_name: str, destination: str | Path) -> None:
        Path(destination).write_bytes(self.download_bytes(blob_name))

    def _associated_data(self, blob_name: str) -> bytes:
        return (
            f"azure-blob-envelope-v{ENCRYPTION_VERSION}:"
            f"{self._container_name}/{blob_name}"
        ).encode("utf-8")


class AsyncEncryptedBlobClient:
    def __init__(
        self,
        blob_service_client: AsyncBlobServiceClient,
        container_name: str,
        key_manager: AsyncKeyManager,
    ) -> None:
        self._container_client = blob_service_client.get_container_client(
            container_name
        )
        self._container_name = container_name
        self._key_manager = key_manager

    async def upload_bytes(
        self, blob_name: str, plaintext: bytes, *, overwrite: bool = False
    ) -> UploadResult:
        data_key = self._key_manager.generate_data_key()
        nonce = token_bytes(NONCE_SIZE_BYTES)
        ciphertext = AESGCM(data_key).encrypt(
            nonce, plaintext, self._associated_data(blob_name)
        )
        protected_key = await self._key_manager.protect_data_key(data_key)
        metadata = _build_metadata(protected_key, nonce)

        try:
            await self._container_client.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=metadata,
                overwrite=overwrite,
            )
        except AzureError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not upload {blob_name!r}"
            ) from exc

        return UploadResult(
            key_id=protected_key.key_id,
            wrapped_key_base64=metadata["wrapped_dek"],
        )

    async def upload_file(
        self, blob_name: str, source: str | Path, *, overwrite: bool = False
    ) -> UploadResult:
        return await self.upload_bytes(
            blob_name, Path(source).read_bytes(), overwrite=overwrite
        )

    async def download_bytes(self, blob_name: str) -> bytes:
        blob_client = self._container_client.get_blob_client(blob_name)
        try:
            properties = await blob_client.get_blob_properties()
            stream = await blob_client.download_blob()
            ciphertext = await stream.readall()
        except ResourceNotFoundError as exc:
            raise EncryptedBlobNotFoundError(
                f"Encrypted blob {blob_name!r} was not found"
            ) from exc
        except AzureError as exc:
            raise BlobEncryptionError(
                f"Blob Storage could not download {blob_name!r}"
            ) from exc

        protected_key, nonce = _parse_metadata(properties.metadata)
        data_key = await self._key_manager.recover_data_key(protected_key)
        try:
            return AESGCM(data_key).decrypt(
                nonce, ciphertext, self._associated_data(blob_name)
            )
        except InvalidTag as exc:
            raise CiphertextIntegrityError(
                f"Authentication failed for encrypted blob {blob_name!r}"
            ) from exc

    async def download_file(
        self, blob_name: str, destination: str | Path
    ) -> None:
        Path(destination).write_bytes(await self.download_bytes(blob_name))

    def _associated_data(self, blob_name: str) -> bytes:
        return (
            f"azure-blob-envelope-v{ENCRYPTION_VERSION}:"
            f"{self._container_name}/{blob_name}"
        ).encode("utf-8")


def _build_metadata(
    protected_key: ProtectedDataKey, nonce: bytes
) -> dict[str, str]:
    return {
        "encryption_version": ENCRYPTION_VERSION,
        "content_algorithm": CONTENT_ALGORITHM,
        "key_wrap_algorithm": KEY_WRAP_ALGORITHM,
        "key_id": protected_key.key_id,
        "wrapped_dek": _base64_encode(protected_key.wrapped_key),
        "nonce": _base64_encode(nonce),
    }


def _parse_metadata(
    metadata: Mapping[str, str] | None,
) -> tuple[ProtectedDataKey, bytes]:
    if not metadata:
        raise EncryptionMetadataError("Blob has no encryption metadata")

    required = {
        "encryption_version",
        "content_algorithm",
        "key_wrap_algorithm",
        "key_id",
        "wrapped_dek",
        "nonce",
    }
    missing = sorted(required.difference(metadata))
    if missing:
        raise EncryptionMetadataError(
            f"Blob encryption metadata is missing: {', '.join(missing)}"
        )
    if metadata["encryption_version"] != ENCRYPTION_VERSION:
        raise EncryptionMetadataError(
            f"Unsupported encryption version "
            f"{metadata['encryption_version']!r}"
        )
    if metadata["content_algorithm"] != CONTENT_ALGORITHM:
        raise EncryptionMetadataError("Unsupported content encryption algorithm")
    if metadata["key_wrap_algorithm"] != KEY_WRAP_ALGORITHM:
        raise EncryptionMetadataError("Unsupported key wrapping algorithm")

    wrapped_key = _base64_decode(metadata["wrapped_dek"], "wrapped DEK")
    nonce = _base64_decode(metadata["nonce"], "nonce")
    if len(nonce) != NONCE_SIZE_BYTES:
        raise EncryptionMetadataError(
            f"AES-GCM nonce must be {NONCE_SIZE_BYTES} bytes"
        )
    if not metadata["key_id"]:
        raise EncryptionMetadataError("Key Vault key ID is empty")

    return (
        ProtectedDataKey(
            key_id=metadata["key_id"],
            wrapped_key=wrapped_key,
        ),
        nonce,
    )


def _base64_encode(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _base64_decode(value: str, field_name: str) -> bytes:
    try:
        decoded = base64.b64decode(value, validate=True)
    except (binascii.Error, ValueError) as exc:
        raise EncryptionMetadataError(
            f"Blob encryption metadata has invalid base64 for {field_name}"
        ) from exc
    if not decoded:
        raise EncryptionMetadataError(
            f"Blob encryption metadata has an empty {field_name}"
        )
    return decoded
