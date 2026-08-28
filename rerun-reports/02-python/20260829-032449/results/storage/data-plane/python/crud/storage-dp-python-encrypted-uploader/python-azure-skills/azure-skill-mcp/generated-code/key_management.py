from __future__ import annotations

from dataclasses import dataclass
from secrets import token_bytes

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import AzureError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import CryptographyClient, KeyWrapAlgorithm
from azure.keyvault.keys.crypto.aio import (
    CryptographyClient as AsyncCryptographyClient,
)

DATA_KEY_SIZE_BYTES = 32
WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256


class KeyManagementError(RuntimeError):
    """Raised when Azure Key Vault cannot protect or recover a data key."""


@dataclass(frozen=True)
class ProtectedDataKey:
    key_id: str
    wrapped_key: bytes


class KeyManager:
    def __init__(
        self,
        key_client: KeyClient,
        credential: TokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name

    @staticmethod
    def generate_data_key() -> bytes:
        return token_bytes(DATA_KEY_SIZE_BYTES)

    def protect_data_key(self, data_key: bytes) -> ProtectedDataKey:
        _validate_data_key(data_key)
        try:
            key = self._key_client.get_key(self._key_name)
            key_id = key.id
            if not key_id:
                raise KeyManagementError(
                    f"Key Vault returned no ID for key {self._key_name!r}"
                )

            # Supplying the ID rather than public key material forces a Key Vault
            # service call, so key enabled-state and policy are always enforced.
            crypto_client = CryptographyClient(key_id, self._credential)
            try:
                result = crypto_client.wrap_key(WRAP_ALGORITHM, data_key)
            finally:
                crypto_client.close()
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key with {self._key_name!r}"
            ) from exc

        return ProtectedDataKey(key_id=key_id, wrapped_key=result.encrypted_key)

    def recover_data_key(self, protected_key: ProtectedDataKey) -> bytes:
        _validate_protected_key(protected_key)
        try:
            crypto_client = CryptographyClient(
                protected_key.key_id, self._credential
            )
            try:
                result = crypto_client.unwrap_key(
                    WRAP_ALGORITHM, protected_key.wrapped_key
                )
            finally:
                crypto_client.close()
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key with "
                f"{protected_key.key_id!r}; the key may be disabled or unavailable"
            ) from exc

        _validate_data_key(result.key)
        return result.key


class AsyncKeyManager:
    def __init__(
        self,
        key_client: AsyncKeyClient,
        credential: AsyncTokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name

    @staticmethod
    def generate_data_key() -> bytes:
        return token_bytes(DATA_KEY_SIZE_BYTES)

    async def protect_data_key(self, data_key: bytes) -> ProtectedDataKey:
        _validate_data_key(data_key)
        try:
            key = await self._key_client.get_key(self._key_name)
            key_id = key.id
            if not key_id:
                raise KeyManagementError(
                    f"Key Vault returned no ID for key {self._key_name!r}"
                )

            crypto_client = AsyncCryptographyClient(key_id, self._credential)
            try:
                result = await crypto_client.wrap_key(WRAP_ALGORITHM, data_key)
            finally:
                await crypto_client.close()
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key with {self._key_name!r}"
            ) from exc

        return ProtectedDataKey(key_id=key_id, wrapped_key=result.encrypted_key)

    async def recover_data_key(
        self, protected_key: ProtectedDataKey
    ) -> bytes:
        _validate_protected_key(protected_key)
        try:
            crypto_client = AsyncCryptographyClient(
                protected_key.key_id, self._credential
            )
            try:
                result = await crypto_client.unwrap_key(
                    WRAP_ALGORITHM, protected_key.wrapped_key
                )
            finally:
                await crypto_client.close()
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key with "
                f"{protected_key.key_id!r}; the key may be disabled or unavailable"
            ) from exc

        _validate_data_key(result.key)
        return result.key


def _validate_data_key(data_key: bytes) -> None:
    if len(data_key) != DATA_KEY_SIZE_BYTES:
        raise ValueError(
            f"The AES-256 data key must be {DATA_KEY_SIZE_BYTES} bytes"
        )


def _validate_protected_key(protected_key: ProtectedDataKey) -> None:
    if not protected_key.key_id:
        raise ValueError("The protected data key has no Key Vault key ID")
    if not protected_key.wrapped_key:
        raise ValueError("The protected data key is empty")
