"""Envelope-key management backed by Azure Key Vault Keys."""

from __future__ import annotations

import os
from dataclasses import dataclass

from azure.core.credentials import TokenCredential
from azure.core.exceptions import AzureError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import (
    CryptographyClient,
    KeyWrapAlgorithm,
)
from azure.keyvault.keys.crypto.aio import CryptographyClient as AsyncCryptographyClient
from azure.core.credentials_async import AsyncTokenCredential

AES_KEY_BYTES = 32
WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256


class KeyManagementError(RuntimeError):
    """Raised when Key Vault cannot protect or recover a data key."""


@dataclass(frozen=True)
class ProtectedDataKey:
    key_id: str
    algorithm: str
    wrapped_key: bytes


def generate_data_key() -> bytearray:
    """Generate an ephemeral AES-256 key that callers must erase after use."""
    return bytearray(os.urandom(AES_KEY_BYTES))


def erase_data_key(data_key: bytearray) -> None:
    for index in range(len(data_key)):
        data_key[index] = 0


class KeyVaultKeyManager:
    def __init__(
        self,
        key_client: KeyClient,
        credential: TokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name

    def protect_data_key(self, data_key: bytes | bytearray) -> ProtectedDataKey:
        try:
            vault_key = self._key_client.get_key(self._key_name)
            if not vault_key.id:
                raise KeyManagementError("Key Vault returned a key without an ID")
            crypto_client = CryptographyClient(vault_key.id, self._credential)
            result = crypto_client.wrap_key(WRAP_ALGORITHM, bytes(data_key))
            return ProtectedDataKey(
                key_id=vault_key.id,
                algorithm=WRAP_ALGORITHM.value,
                wrapped_key=result.encrypted_key,
            )
        except KeyManagementError:
            raise
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key: {exc}"
            ) from exc

    def recover_data_key(self, protected_key: ProtectedDataKey) -> bytearray:
        try:
            algorithm = KeyWrapAlgorithm(protected_key.algorithm)
            crypto_client = CryptographyClient(
                protected_key.key_id, self._credential
            )
            result = crypto_client.unwrap_key(algorithm, protected_key.wrapped_key)
            if len(result.key) != AES_KEY_BYTES:
                raise KeyManagementError(
                    f"Key Vault returned an invalid {len(result.key)}-byte data key"
                )
            return bytearray(result.key)
        except (ValueError, KeyManagementError):
            raise
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key: {exc}"
            ) from exc


class AsyncKeyVaultKeyManager:
    def __init__(
        self,
        key_client: AsyncKeyClient,
        credential: AsyncTokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name

    async def protect_data_key(
        self, data_key: bytes | bytearray
    ) -> ProtectedDataKey:
        try:
            vault_key = await self._key_client.get_key(self._key_name)
            if not vault_key.id:
                raise KeyManagementError("Key Vault returned a key without an ID")
            crypto_client = AsyncCryptographyClient(
                vault_key.id, self._credential
            )
            result = await crypto_client.wrap_key(
                WRAP_ALGORITHM, bytes(data_key)
            )
            return ProtectedDataKey(
                key_id=vault_key.id,
                algorithm=WRAP_ALGORITHM.value,
                wrapped_key=result.encrypted_key,
            )
        except KeyManagementError:
            raise
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key: {exc}"
            ) from exc

    async def recover_data_key(
        self, protected_key: ProtectedDataKey
    ) -> bytearray:
        try:
            algorithm = KeyWrapAlgorithm(protected_key.algorithm)
            crypto_client = AsyncCryptographyClient(
                protected_key.key_id, self._credential
            )
            result = await crypto_client.unwrap_key(
                algorithm, protected_key.wrapped_key
            )
            if len(result.key) != AES_KEY_BYTES:
                raise KeyManagementError(
                    f"Key Vault returned an invalid {len(result.key)}-byte data key"
                )
            return bytearray(result.key)
        except (ValueError, KeyManagementError):
            raise
        except AzureError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key: {exc}"
            ) from exc
