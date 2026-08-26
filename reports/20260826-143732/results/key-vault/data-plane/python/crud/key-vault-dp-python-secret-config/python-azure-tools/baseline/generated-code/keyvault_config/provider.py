from datetime import datetime
from typing import Optional

from azure.core.exceptions import ResourceNotFoundError
from azure.keyvault.secrets import SecretClient
from azure.keyvault.secrets.aio import SecretClient as AsyncSecretClient

from .models import SecretInfo


class KeyVaultSecretProvider:
    """Synchronous access to Key Vault secrets."""

    def __init__(self, client: SecretClient) -> None:
        self.client = client

    def get_secret_info(
        self,
        name: str,
        default: Optional[str] = None,
        version: Optional[str] = None,
    ) -> SecretInfo:
        try:
            secret = self.client.get_secret(name, version)
        except ResourceNotFoundError:
            return SecretInfo(name, default, version, None, found=False)

        return SecretInfo(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    def get_secret(
        self,
        name: str,
        default: Optional[str] = None,
        version: Optional[str] = None,
    ) -> Optional[str]:
        return self.get_secret_info(name, default, version).value

    def get_expiry(
        self, name: str, version: Optional[str] = None
    ) -> Optional[datetime]:
        return self.get_secret_info(name, version=version).expires_on


class AsyncKeyVaultSecretProvider:
    """Asynchronous access to Key Vault secrets."""

    def __init__(self, client: AsyncSecretClient) -> None:
        self.client = client

    async def get_secret_info(
        self,
        name: str,
        default: Optional[str] = None,
        version: Optional[str] = None,
    ) -> SecretInfo:
        try:
            secret = await self.client.get_secret(name, version)
        except ResourceNotFoundError:
            return SecretInfo(name, default, version, None, found=False)

        return SecretInfo(
            name=secret.name,
            value=secret.value,
            version=secret.properties.version,
            expires_on=secret.properties.expires_on,
            found=True,
        )

    async def get_secret(
        self,
        name: str,
        default: Optional[str] = None,
        version: Optional[str] = None,
    ) -> Optional[str]:
        info = await self.get_secret_info(name, default, version)
        return info.value

    async def get_expiry(
        self, name: str, version: Optional[str] = None
    ) -> Optional[datetime]:
        info = await self.get_secret_info(name, version=version)
        return info.expires_on
