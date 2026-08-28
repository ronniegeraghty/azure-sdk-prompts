from __future__ import annotations

from datetime import datetime

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets.aio import SecretClient

from .models import SecretInfo


class AsyncSecretProvider:
    def __init__(self, client: SecretClient) -> None:
        self._client = client

    async def get(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> str | None:
        info = await self.get_info(name, default, version=version)
        return info.value

    async def get_info(
        self,
        name: str,
        default: str | None = None,
        *,
        version: str | None = None,
    ) -> SecretInfo:
        try:
            secret = await self._client.get_secret(name, version=version)
        except ResourceNotFoundError:
            return SecretInfo(name, default, None, None)

        return SecretInfo(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
        )

    async def get_expiry(
        self, name: str, *, version: str | None = None
    ) -> datetime | None:
        info = await self.get_info(name, version=version)
        return info.expires_on
