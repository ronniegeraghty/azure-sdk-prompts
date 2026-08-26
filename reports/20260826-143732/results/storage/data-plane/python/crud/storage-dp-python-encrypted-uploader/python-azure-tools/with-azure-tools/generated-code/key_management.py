from __future__ import annotations

import os
from dataclasses import dataclass
from typing import Protocol

from azure.core.credentials import TokenCredential
from azure.core.credentials_async import AsyncTokenCredential
from azure.core.exceptions import AzureError, HttpResponseError, ResourceNotFoundError
from azure.keyvault.keys import KeyClient
from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
from azure.keyvault.keys.crypto import (
    CryptographyClient,
    KeyWrapAlgorithm,
)
from azure.keyvault.keys.crypto.aio import (
    CryptographyClient as AsyncCryptographyClient,
)

DATA_KEY_SIZE_BYTES = 32
KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256
KEY_WRAP_ALGORITHM_NAME = "RSA-OAEP-256"


class KeyManagementError(RuntimeError):
    """Raised when a Key Vault key operation cannot be completed."""


class KeyVaultKey(Protocol):
    id: str | None


@dataclass(frozen=True)
class EnvelopeKey:
    plaintext_key: bytearray
    wrapped_key: bytes
    key_id: str


def erase_key(key: bytearray) -> None:
    for index in range(len(key)):
        key[index] = 0


def _require_key_id(key: KeyVaultKey, key_name: str) -> str:
    if not key.id:
        raise KeyManagementError(
            f"Key Vault returned no versioned key ID for key {key_name!r}"
        )
    return key.id


def _operation_error(
    operation: str, key_reference: str, error: AzureError
) -> KeyManagementError:
    status_code = getattr(error, "status_code", None)
    status = f" (HTTP {status_code})" if status_code else ""
    return KeyManagementError(
        f"Key Vault could not {operation} using {key_reference!r}{status}. "
        "Verify that the key exists, is enabled, permits wrapKey/unwrapKey, "
        "and that the caller has the required RBAC role."
    )


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

    def create_envelope_key(self) -> EnvelopeKey:
        plaintext_key = bytearray(os.urandom(DATA_KEY_SIZE_BYTES))
        try:
            try:
                key = self._key_client.get_key(self._key_name)
                key_id = _require_key_id(key, self._key_name)
                with CryptographyClient(
                    key_id, credential=self._credential
                ) as crypto_client:
                    result = crypto_client.wrap_key(
                        KEY_WRAP_ALGORITHM, bytes(plaintext_key)
                    )
            except ResourceNotFoundError as error:
                raise KeyManagementError(
                    f"Key Vault key {self._key_name!r} was not found"
                ) from error
            except HttpResponseError as error:
                raise _operation_error("wrap a data key", self._key_name, error) from error
            except AzureError as error:
                raise _operation_error("wrap a data key", self._key_name, error) from error

            return EnvelopeKey(plaintext_key, result.encrypted_key, key_id)
        except BaseException:
            erase_key(plaintext_key)
            raise

    def unwrap_key(self, wrapped_key: bytes, key_id: str) -> bytearray:
        try:
            with CryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = crypto_client.unwrap_key(KEY_WRAP_ALGORITHM, wrapped_key)
        except ResourceNotFoundError as error:
            raise KeyManagementError(
                f"The Key Vault key version {key_id!r} was not found"
            ) from error
        except HttpResponseError as error:
            raise _operation_error("unwrap a data key", key_id, error) from error
        except AzureError as error:
            raise _operation_error("unwrap a data key", key_id, error) from error

        if len(result.key) != DATA_KEY_SIZE_BYTES:
            raise KeyManagementError("Key Vault returned an invalid data-key length")
        return bytearray(result.key)


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

    async def create_envelope_key(self) -> EnvelopeKey:
        plaintext_key = bytearray(os.urandom(DATA_KEY_SIZE_BYTES))
        try:
            try:
                key = await self._key_client.get_key(self._key_name)
                key_id = _require_key_id(key, self._key_name)
                async with AsyncCryptographyClient(
                    key_id, credential=self._credential
                ) as crypto_client:
                    result = await crypto_client.wrap_key(
                        KEY_WRAP_ALGORITHM, bytes(plaintext_key)
                    )
            except ResourceNotFoundError as error:
                raise KeyManagementError(
                    f"Key Vault key {self._key_name!r} was not found"
                ) from error
            except HttpResponseError as error:
                raise _operation_error("wrap a data key", self._key_name, error) from error
            except AzureError as error:
                raise _operation_error("wrap a data key", self._key_name, error) from error

            return EnvelopeKey(plaintext_key, result.encrypted_key, key_id)
        except BaseException:
            erase_key(plaintext_key)
            raise

    async def unwrap_key(self, wrapped_key: bytes, key_id: str) -> bytearray:
        try:
            async with AsyncCryptographyClient(
                key_id, credential=self._credential
            ) as crypto_client:
                result = await crypto_client.unwrap_key(
                    KEY_WRAP_ALGORITHM, wrapped_key
                )
        except ResourceNotFoundError as error:
            raise KeyManagementError(
                f"The Key Vault key version {key_id!r} was not found"
            ) from error
        except HttpResponseError as error:
            raise _operation_error("unwrap a data key", key_id, error) from error
        except AzureError as error:
            raise _operation_error("unwrap a data key", key_id, error) from error

        if len(result.key) != DATA_KEY_SIZE_BYTES:
            raise KeyManagementError("Key Vault returned an invalid data-key length")
        return bytearray(result.key)
