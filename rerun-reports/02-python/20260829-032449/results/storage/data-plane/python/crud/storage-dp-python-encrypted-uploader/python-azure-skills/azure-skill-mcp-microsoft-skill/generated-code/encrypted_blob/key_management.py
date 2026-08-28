"""Envelope key management backed by Azure Key Vault Keys."""

from __future__ import annotations

import os
from dataclasses import dataclass

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import HttpResponseError, ResourceNotFoundError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import CryptographyClient, KeyWrapAlgorithm
from azure.keyvault.keys.crypto.aio import (
    CryptographyClient as AsyncCryptographyClient,
)

DATA_KEY_SIZE_BYTES = 32
KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256


class KeyManagementError(RuntimeError):
    """Raised when a Key Vault key operation cannot be completed."""


@dataclass
class DataKeyMaterial:
    plaintext_key: bytearray
    wrapped_key: bytes
    key_id: str

    def destroy(self) -> None:
        for index in range(len(self.plaintext_key)):
            self.plaintext_key[index] = 0


def _key_vault_error(operation: str, key_id: str, error: HttpResponseError) -> str:
    status = f"HTTP {error.status_code}" if error.status_code else "an HTTP error"
    return f"Key Vault {operation} failed for key {key_id} ({status})"


class SyncKeyManager:
    """Generates local DEKs and protects them with a versioned Key Vault key."""

    def __init__(
        self,
        key_client: KeyClient,
        credential: TokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name
        self._key_id: str | None = None

    def get_key_id(self) -> str:
        if self._key_id is None:
            try:
                key = self._key_client.get_key(self._key_name)
            except ResourceNotFoundError as error:
                raise KeyManagementError(
                    f"Key Vault key {self._key_name!r} does not exist"
                ) from error
            except HttpResponseError as error:
                raise KeyManagementError(
                    _key_vault_error("key lookup", self._key_name, error)
                ) from error
            if not key.id:
                raise KeyManagementError(
                    f"Key Vault returned no ID for key {self._key_name!r}"
                )
            self._key_id = key.id
        return self._key_id

    def generate_and_wrap_data_key(self) -> DataKeyMaterial:
        key_id = self.get_key_id()
        plaintext_key = bytearray(os.urandom(DATA_KEY_SIZE_BYTES))
        try:
            with CryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = crypto_client.wrap_key(
                    KEY_WRAP_ALGORITHM, bytes(plaintext_key)
                )
            return DataKeyMaterial(plaintext_key, result.encrypted_key, key_id)
        except HttpResponseError as error:
            for index in range(len(plaintext_key)):
                plaintext_key[index] = 0
            raise KeyManagementError(
                _key_vault_error("key wrapping", key_id, error)
            ) from error

    def unwrap_data_key(self, wrapped_key: bytes, key_id: str) -> bytearray:
        try:
            with CryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = crypto_client.unwrap_key(KEY_WRAP_ALGORITHM, wrapped_key)
            plaintext_key = bytearray(result.key)
        except ResourceNotFoundError as error:
            raise KeyManagementError(
                f"The Key Vault key version {key_id} does not exist"
            ) from error
        except HttpResponseError as error:
            raise KeyManagementError(
                _key_vault_error("key unwrapping", key_id, error)
            ) from error
        if len(plaintext_key) != DATA_KEY_SIZE_BYTES:
            for index in range(len(plaintext_key)):
                plaintext_key[index] = 0
            raise KeyManagementError("Key Vault returned an invalid data key length")
        return plaintext_key


class AsyncKeyManager:
    """Async equivalent of SyncKeyManager."""

    def __init__(
        self,
        key_client: AsyncKeyClient,
        credential: AsyncTokenCredential,
        key_name: str,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name
        self._key_id: str | None = None

    async def get_key_id(self) -> str:
        if self._key_id is None:
            try:
                key = await self._key_client.get_key(self._key_name)
            except ResourceNotFoundError as error:
                raise KeyManagementError(
                    f"Key Vault key {self._key_name!r} does not exist"
                ) from error
            except HttpResponseError as error:
                raise KeyManagementError(
                    _key_vault_error("key lookup", self._key_name, error)
                ) from error
            if not key.id:
                raise KeyManagementError(
                    f"Key Vault returned no ID for key {self._key_name!r}"
                )
            self._key_id = key.id
        return self._key_id

    async def generate_and_wrap_data_key(self) -> DataKeyMaterial:
        key_id = await self.get_key_id()
        plaintext_key = bytearray(os.urandom(DATA_KEY_SIZE_BYTES))
        try:
            async with AsyncCryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = await crypto_client.wrap_key(
                    KEY_WRAP_ALGORITHM, bytes(plaintext_key)
                )
            return DataKeyMaterial(plaintext_key, result.encrypted_key, key_id)
        except HttpResponseError as error:
            for index in range(len(plaintext_key)):
                plaintext_key[index] = 0
            raise KeyManagementError(
                _key_vault_error("key wrapping", key_id, error)
            ) from error

    async def unwrap_data_key(
        self, wrapped_key: bytes, key_id: str
    ) -> bytearray:
        try:
            async with AsyncCryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = await crypto_client.unwrap_key(
                    KEY_WRAP_ALGORITHM, wrapped_key
                )
            plaintext_key = bytearray(result.key)
        except ResourceNotFoundError as error:
            raise KeyManagementError(
                f"The Key Vault key version {key_id} does not exist"
            ) from error
        except HttpResponseError as error:
            raise KeyManagementError(
                _key_vault_error("key unwrapping", key_id, error)
            ) from error
        if len(plaintext_key) != DATA_KEY_SIZE_BYTES:
            for index in range(len(plaintext_key)):
                plaintext_key[index] = 0
            raise KeyManagementError("Key Vault returned an invalid data key length")
        return plaintext_key
