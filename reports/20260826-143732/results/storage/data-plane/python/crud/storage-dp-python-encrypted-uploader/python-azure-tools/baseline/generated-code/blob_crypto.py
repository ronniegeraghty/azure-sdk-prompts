"""Client-side AES-GCM encryption for Azure Blob Storage."""

from __future__ import annotations

import base64
import binascii
import os
from dataclasses import dataclass
from typing import Mapping

from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.storage.blob import ContainerClient
from azure.storage.blob.aio import ContainerClient as AsyncContainerClient
from cryptography.exceptions import InvalidTag
from cryptography.hazmat.primitives.ciphers.aead import AESGCM

from key_management import (
    AsyncKeyVaultKeyManager,
    KeyVaultKeyManager,
    ProtectedDataKey,
    erase_data_key,
    generate_data_key,
)

NONCE_BYTES = 12
FORMAT_VERSION = "1"
AUTHENTICATED_CONTEXT = b"azure-blob-envelope-encryption-v1"


class EncryptedBlobError(RuntimeError):
    """Base error for encrypted blob operations."""


class BlobNotFoundError(EncryptedBlobError):
    """Raised when the requested encrypted blob does not exist."""


class InvalidBlobMetadataError(EncryptedBlobError):
    """Raised when encryption metadata is absent or malformed."""


class BlobDecryptionError(EncryptedBlobError):
    """Raised when ciphertext authentication or decryption fails."""


@dataclass(frozen=True)
class UploadResult:
    blob_name: str
    protected_key: ProtectedDataKey


def _b64encode(value: bytes) -> str:
    return base64.b64encode(value).decode("ascii")


def _metadata(protected_key: ProtectedDataKey, nonce: bytes) -> dict[str, str]:
    return {
        "encryption_version": FORMAT_VERSION,
        "key_id": protected_key.key_id,
        "key_wrap_algorithm": protected_key.algorithm,
        "wrapped_dek": _b64encode(protected_key.wrapped_key),
        "nonce": _b64encode(nonce),
    }


def _parse_metadata(
    metadata: Mapping[str, str] | None,
) -> tuple[ProtectedDataKey, bytes]:
    try:
        if not metadata or metadata["encryption_version"] != FORMAT_VERSION:
            raise InvalidBlobMetadataError("Unsupported encryption metadata version")
        wrapped_key = base64.b64decode(metadata["wrapped_dek"], validate=True)
        nonce = base64.b64decode(metadata["nonce"], validate=True)
        if not wrapped_key:
            raise InvalidBlobMetadataError("Wrapped data key is empty")
        if len(nonce) != NONCE_BYTES:
            raise InvalidBlobMetadataError("AES-GCM nonce must be 12 bytes")
        return (
            ProtectedDataKey(
                key_id=metadata["key_id"],
                algorithm=metadata["key_wrap_algorithm"],
                wrapped_key=wrapped_key,
            ),
            nonce,
        )
    except KeyError as exc:
        raise InvalidBlobMetadataError(
            f"Required encryption metadata {exc.args[0]!r} is missing"
        ) from exc
    except (binascii.Error, ValueError) as exc:
        raise InvalidBlobMetadataError(
            "Encryption metadata contains invalid base64"
        ) from exc


def _decrypt(
    ciphertext: bytes,
    protected_key: ProtectedDataKey,
    nonce: bytes,
    key_manager: KeyVaultKeyManager,
) -> bytes:
    data_key = key_manager.recover_data_key(protected_key)
    try:
        return AESGCM(bytes(data_key)).decrypt(
            nonce, ciphertext, AUTHENTICATED_CONTEXT
        )
    except InvalidTag as exc:
        raise BlobDecryptionError(
            "Ciphertext or its authentication data is invalid"
        ) from exc
    finally:
        erase_data_key(data_key)


class EncryptedBlobClient:
    def __init__(
        self,
        container: ContainerClient,
        key_manager: KeyVaultKeyManager,
    ) -> None:
        self._container = container
        self._key_manager = key_manager

    def upload(self, blob_name: str, plaintext: bytes) -> UploadResult:
        data_key = generate_data_key()
        try:
            nonce = os.urandom(NONCE_BYTES)
            ciphertext = AESGCM(bytes(data_key)).encrypt(
                nonce, plaintext, AUTHENTICATED_CONTEXT
            )
            protected_key = self._key_manager.protect_data_key(data_key)
        finally:
            erase_data_key(data_key)

        try:
            self._container.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=_metadata(protected_key, nonce),
                overwrite=True,
            )
            return UploadResult(blob_name, protected_key)
        except AzureError as exc:
            raise EncryptedBlobError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

    def download(self, blob_name: str) -> bytes:
        try:
            blob = self._container.get_blob_client(blob_name)
            properties = blob.get_blob_properties()
            ciphertext = blob.download_blob().readall()
        except ResourceNotFoundError as exc:
            raise BlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist"
            ) from exc
        except AzureError as exc:
            raise EncryptedBlobError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        protected_key, nonce = _parse_metadata(properties.metadata)
        return _decrypt(ciphertext, protected_key, nonce, self._key_manager)


class AsyncEncryptedBlobClient:
    def __init__(
        self,
        container: AsyncContainerClient,
        key_manager: AsyncKeyVaultKeyManager,
    ) -> None:
        self._container = container
        self._key_manager = key_manager

    async def upload(self, blob_name: str, plaintext: bytes) -> UploadResult:
        data_key = generate_data_key()
        try:
            nonce = os.urandom(NONCE_BYTES)
            ciphertext = AESGCM(bytes(data_key)).encrypt(
                nonce, plaintext, AUTHENTICATED_CONTEXT
            )
            protected_key = await self._key_manager.protect_data_key(data_key)
        finally:
            erase_data_key(data_key)

        try:
            await self._container.upload_blob(
                name=blob_name,
                data=ciphertext,
                metadata=_metadata(protected_key, nonce),
                overwrite=True,
            )
            return UploadResult(blob_name, protected_key)
        except AzureError as exc:
            raise EncryptedBlobError(
                f"Blob Storage could not upload {blob_name!r}: {exc}"
            ) from exc

    async def download(self, blob_name: str) -> bytes:
        try:
            blob = self._container.get_blob_client(blob_name)
            properties = await blob.get_blob_properties()
            stream = await blob.download_blob()
            ciphertext = await stream.readall()
        except ResourceNotFoundError as exc:
            raise BlobNotFoundError(
                f"Encrypted blob {blob_name!r} does not exist"
            ) from exc
        except AzureError as exc:
            raise EncryptedBlobError(
                f"Blob Storage could not download {blob_name!r}: {exc}"
            ) from exc

        protected_key, nonce = _parse_metadata(properties.metadata)
        data_key = await self._key_manager.recover_data_key(protected_key)
        try:
            return AESGCM(bytes(data_key)).decrypt(
                nonce, ciphertext, AUTHENTICATED_CONTEXT
            )
        except InvalidTag as exc:
            raise BlobDecryptionError(
                "Ciphertext or its authentication data is invalid"
            ) from exc
        finally:
            erase_data_key(data_key)
