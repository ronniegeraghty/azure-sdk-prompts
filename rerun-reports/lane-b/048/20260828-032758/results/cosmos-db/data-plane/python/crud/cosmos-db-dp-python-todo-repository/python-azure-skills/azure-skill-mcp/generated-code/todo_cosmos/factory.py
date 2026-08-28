from __future__ import annotations

import os
from dataclasses import dataclass

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from .repository import TodoRepository
from .repository_async import AsyncTodoRepository

DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60
INDEXING_POLICY = {
    "indexingMode": "consistent",
    "automatic": True,
    "includedPaths": [{"path": "/*"}],
    "excludedPaths": [{"path": '/description/?'}],
}


@dataclass(frozen=True, slots=True)
class CosmosSettings:
    endpoint: str
    database_name: str = "todo-db"
    container_name: str = "todos"

    @classmethod
    def from_environment(cls) -> CosmosSettings:
        endpoint = os.environ.get("AZURE_COSMOS_ENDPOINT")
        if not endpoint:
            raise RuntimeError(
                "Set AZURE_COSMOS_ENDPOINT to the Azure Cosmos DB account endpoint."
            )
        return cls(
            endpoint=endpoint,
            database_name=os.environ.get("AZURE_COSMOS_DATABASE", "todo-db"),
            container_name=os.environ.get("AZURE_COSMOS_CONTAINER", "todos"),
        )


@dataclass(slots=True)
class SyncRepositoryResources:
    repository: TodoRepository
    client: CosmosClient
    credential: DefaultAzureCredential

    def close(self) -> None:
        self.client.close()
        self.credential.close()

    def __enter__(self) -> SyncRepositoryResources:
        return self

    def __exit__(self, *_args: object) -> None:
        self.close()


@dataclass(slots=True)
class AsyncRepositoryResources:
    repository: AsyncTodoRepository
    client: AsyncCosmosClient
    credential: AsyncDefaultAzureCredential

    async def close(self) -> None:
        await self.client.close()
        await self.credential.close()

    async def __aenter__(self) -> AsyncRepositoryResources:
        return self

    async def __aexit__(self, *_args: object) -> None:
        await self.close()


def create_sync_repository(
    settings: CosmosSettings | None = None,
) -> SyncRepositoryResources:
    settings = settings or CosmosSettings.from_environment()
    credential = DefaultAzureCredential()
    client = CosmosClient(settings.endpoint, credential=credential)

    try:
        database = client.create_database_if_not_exists(id=settings.database_name)
        container = database.create_container_if_not_exists(
            id=settings.container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        client.close()
        credential.close()
        raise

    return SyncRepositoryResources(
        repository=TodoRepository(container),
        client=client,
        credential=credential,
    )


async def create_async_repository(
    settings: CosmosSettings | None = None,
) -> AsyncRepositoryResources:
    settings = settings or CosmosSettings.from_environment()
    credential = AsyncDefaultAzureCredential()
    client = AsyncCosmosClient(settings.endpoint, credential=credential)

    try:
        database = await client.create_database_if_not_exists(
            id=settings.database_name
        )
        container = await database.create_container_if_not_exists(
            id=settings.container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        await client.close()
        await credential.close()
        raise

    return AsyncRepositoryResources(
        repository=AsyncTodoRepository(container),
        client=client,
        credential=credential,
    )
