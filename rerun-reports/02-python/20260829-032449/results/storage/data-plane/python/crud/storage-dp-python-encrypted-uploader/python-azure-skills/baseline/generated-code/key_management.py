from __future__ import annotations

import os
from dataclasses import dataclass

from azure.core.exceptions import HttpResponseError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import CryptographyClient, KeyWrapAlgorithm
from azure.keyvault.keys.crypto.aio import CryptographyClient as AsyncCryptographyClient
from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential

DATA_KEY_SIZE_BYTES = 32
KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256


class KeyManagementError(RuntimeError):
    """Raised when a Key Vault key operation fails."""


@dataclass(frozen=True)
class WrappedDataKey:
    key_id: str
    encrypted_key: bytes
    algorithm: str = KEY_WRAP_ALGORITHM.value


class KeyManager:
    def __init__(
        self,
        key_client: KeyClient,
        credential: TokenCredential,
        key_name: str,
        key_version: str | None = None,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name
        self._key_version = key_version

    def generate_and_wrap_data_key(self) -> tuple[bytes, WrappedDataKey]:
        data_key = os.urandom(DATA_KEY_SIZE_BYTES)
        try:
            key = self._key_client.get_key(self._key_name, self._key_version)
            if not key.id:
                raise KeyManagementError("Key Vault returned a key without an ID")
            crypto_client = CryptographyClient(key, credential=self._credential)
            result = crypto_client.wrap_key(KEY_WRAP_ALGORITHM, data_key)
            return data_key, WrappedDataKey(
                key_id=key.id,
                encrypted_key=result.encrypted_key,
            )
        except HttpResponseError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key with {self._key_name!r}: {exc}"
            ) from exc

    def unwrap_data_key(self, wrapped_key: WrappedDataKey) -> bytes:
        if wrapped_key.algorithm != KEY_WRAP_ALGORITHM.value:
            raise KeyManagementError(
                f"Unsupported key wrap algorithm: {wrapped_key.algorithm!r}"
            )
        try:
            crypto_client = CryptographyClient(
                wrapped_key.key_id,
                credential=self._credential,
            )
            result = crypto_client.unwrap_key(
                KEY_WRAP_ALGORITHM,
                wrapped_key.encrypted_key,
            )
            return result.key
        except HttpResponseError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key with "
                f"{wrapped_key.key_id!r}: {exc}"
            ) from exc


class AsyncKeyManager:
    def __init__(
        self,
        key_client: AsyncKeyClient,
        credential: AsyncTokenCredential,
        key_name: str,
        key_version: str | None = None,
    ) -> None:
        self._key_client = key_client
        self._credential = credential
        self._key_name = key_name
        self._key_version = key_version

    async def generate_and_wrap_data_key(self) -> tuple[bytes, WrappedDataKey]:
        data_key = os.urandom(DATA_KEY_SIZE_BYTES)
        try:
            key = await self._key_client.get_key(self._key_name, self._key_version)
            if not key.id:
                raise KeyManagementError("Key Vault returned a key without an ID")
            crypto_client = AsyncCryptographyClient(
                key,
                credential=self._credential,
            )
            result = await crypto_client.wrap_key(KEY_WRAP_ALGORITHM, data_key)
            return data_key, WrappedDataKey(
                key_id=key.id,
                encrypted_key=result.encrypted_key,
            )
        except HttpResponseError as exc:
            raise KeyManagementError(
                f"Key Vault could not wrap the data key with {self._key_name!r}: {exc}"
            ) from exc

    async def unwrap_data_key(self, wrapped_key: WrappedDataKey) -> bytes:
        if wrapped_key.algorithm != KEY_WRAP_ALGORITHM.value:
            raise KeyManagementError(
                f"Unsupported key wrap algorithm: {wrapped_key.algorithm!r}"
            )
        try:
            crypto_client = AsyncCryptographyClient(
                wrapped_key.key_id,
                credential=self._credential,
            )
            result = await crypto_client.unwrap_key(
                KEY_WRAP_ALGORITHM,
                wrapped_key.encrypted_key,
            )
            return result.key
        except HttpResponseError as exc:
            raise KeyManagementError(
                f"Key Vault could not unwrap the data key with "
                f"{wrapped_key.key_id!r}: {exc}"
            ) from exc
