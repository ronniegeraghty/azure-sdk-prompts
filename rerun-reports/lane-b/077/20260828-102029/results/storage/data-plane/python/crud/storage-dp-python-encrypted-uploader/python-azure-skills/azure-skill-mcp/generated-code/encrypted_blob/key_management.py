"""Envelope-key management backed by Azure Key Vault Keys."""

from __future__ import annotations

import secrets
from dataclasses import dataclass

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import AzureError, ResourceNotFoundError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import (
    CryptographyClient,
    KeyWrapAlgorithm,
)
from azure.keyvault.keys.crypto.aio import (
    CryptographyClient as AsyncCryptographyClient,
)

DATA_KEY_BYTES = 32
KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256


class KeyManagementError(RuntimeError):
    """Raised when Key Vault cannot resolve, wrap, or unwrap a data key."""


@dataclass(frozen=True)
class WrappedDataKey:
    """A data encryption key protected by a versioned Key Vault key."""

    key_id: str
    algorithm: str
    wrapped_key: bytes


def _missing_key_message(key_name: str, key_version: str | None) -> str:
    version_description = key_version or "current version"
    return (
        f"Key Vault key {key_name!r} ({version_description}) was not found. "
        "Verify the key name, version, and RBAC permissions."
    )


class KeyManager:
    """Generates local DEKs and protects them with a Key Vault key."""

    def __init__(self, key_id: str, credential: TokenCredential) -> None:
        self.key_id = key_id
        self._credential = credential
        self._cryptography_client = CryptographyClient(key_id, credential)

    @classmethod
    def from_key_client(
        cls,
        key_client: KeyClient,
        credential: TokenCredential,
        key_name: str,
        key_version: str | None = None,
    ) -> "KeyManager":
        try:
            key = key_client.get_key(key_name, key_version)
        except ResourceNotFoundError as exc:
            raise KeyManagementError(
                _missing_key_message(key_name, key_version)
            ) from exc
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not resolve key {key_name!r}: {exc}"
            ) from exc

        if not key.id:
            raise KeyManagementError(
                f"Key Vault returned key {key_name!r} without a versioned key ID."
            )
        return cls(key.id, credential)

    def generate_and_wrap_data_key(self) -> tuple[bytes, WrappedDataKey]:
        data_key = secrets.token_bytes(DATA_KEY_BYTES)
        try:
            result = self._cryptography_client.wrap_key(
                KEY_WRAP_ALGORITHM, data_key
            )
        except AzureError as exc:
            raise KeyManagementError(
                "Key Vault could not wrap the data key. The key may be disabled, "
                f"expired, or inaccessible: {exc}"
            ) from exc

        return data_key, WrappedDataKey(
            key_id=result.key_id or self.key_id,
            algorithm=KEY_WRAP_ALGORITHM.value,
            wrapped_key=result.encrypted_key,
        )

    def unwrap_data_key(self, wrapped: WrappedDataKey) -> bytes:
        if wrapped.algorithm != KEY_WRAP_ALGORITHM.value:
            raise KeyManagementError(
                f"Unsupported key-wrap algorithm {wrapped.algorithm!r}."
            )

        client = self._cryptography_client
        temporary_client: CryptographyClient | None = None
        if wrapped.key_id != self.key_id:
            temporary_client = CryptographyClient(wrapped.key_id, self._credential)
            client = temporary_client

        try:
            result = client.unwrap_key(KEY_WRAP_ALGORITHM, wrapped.wrapped_key)
            return result.key
        except AzureError as exc:
            raise KeyManagementError(
                "Key Vault could not unwrap the data key. Its exact key version "
                f"may be disabled, deleted, or inaccessible: {exc}"
            ) from exc
        finally:
            if temporary_client is not None:
                temporary_client.close()

    def close(self) -> None:
        self._cryptography_client.close()

    def __enter__(self) -> "KeyManager":
        return self

    def __exit__(self, exc_type: object, exc_value: object, traceback: object) -> None:
        self.close()


class AsyncKeyManager:
    """Async Key Vault operations for locally generated data keys."""

    def __init__(self, key_id: str, credential: AsyncTokenCredential) -> None:
        self.key_id = key_id
        self._credential = credential
        self._cryptography_client = AsyncCryptographyClient(key_id, credential)

    @classmethod
    async def from_key_client(
        cls,
        key_client: AsyncKeyClient,
        credential: AsyncTokenCredential,
        key_name: str,
        key_version: str | None = None,
    ) -> "AsyncKeyManager":
        try:
            key = await key_client.get_key(key_name, key_version)
        except ResourceNotFoundError as exc:
            raise KeyManagementError(
                _missing_key_message(key_name, key_version)
            ) from exc
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not resolve key {key_name!r}: {exc}"
            ) from exc

        if not key.id:
            raise KeyManagementError(
                f"Key Vault returned key {key_name!r} without a versioned key ID."
            )
        return cls(key.id, credential)

    async def generate_and_wrap_data_key(self) -> tuple[bytes, WrappedDataKey]:
        data_key = secrets.token_bytes(DATA_KEY_BYTES)
        try:
            result = await self._cryptography_client.wrap_key(
                KEY_WRAP_ALGORITHM, data_key
            )
        except AzureError as exc:
            raise KeyManagementError(
                "Key Vault could not wrap the data key. The key may be disabled, "
                f"expired, or inaccessible: {exc}"
            ) from exc

        return data_key, WrappedDataKey(
            key_id=result.key_id or self.key_id,
            algorithm=KEY_WRAP_ALGORITHM.value,
            wrapped_key=result.encrypted_key,
        )

    async def unwrap_data_key(self, wrapped: WrappedDataKey) -> bytes:
        if wrapped.algorithm != KEY_WRAP_ALGORITHM.value:
            raise KeyManagementError(
                f"Unsupported key-wrap algorithm {wrapped.algorithm!r}."
            )

        client = self._cryptography_client
        temporary_client: AsyncCryptographyClient | None = None
        if wrapped.key_id != self.key_id:
            temporary_client = AsyncCryptographyClient(
                wrapped.key_id, self._credential
            )
            client = temporary_client

        try:
            result = await client.unwrap_key(
                KEY_WRAP_ALGORITHM, wrapped.wrapped_key
            )
            return result.key
        except AzureError as exc:
            raise KeyManagementError(
                "Key Vault could not unwrap the data key. Its exact key version "
                f"may be disabled, deleted, or inaccessible: {exc}"
            ) from exc
        finally:
            if temporary_client is not None:
                await temporary_client.close()

    async def close(self) -> None:
        await self._cryptography_client.close()

    async def __aenter__(self) -> "AsyncKeyManager":
        return self

    async def __aexit__(
        self, exc_type: object, exc_value: object, traceback: object
    ) -> None:
        await self.close()
