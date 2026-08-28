from __future__ import annotations

import os
from dataclasses import dataclass

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from repository_async import AsyncTodoRepository
from repository_sync import SyncTodoRepository


DEFAULT_DATABASE_NAME = "TodoDatabase"
DEFAULT_CONTAINER_NAME = "TodoItems"
DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60
INDEXING_POLICY = {
    "indexingMode": "consistent",
    "automatic": True,
    "includedPaths": [{"path": "/*"}],
    "excludedPaths": [{"path": "/description/?"}],
}


def _settings() -> tuple[str, str, str]:
    endpoint = os.environ.get("COSMOS_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "COSMOS_ENDPOINT must contain the Azure Cosmos DB account endpoint."
        )

    return (
        endpoint,
        os.environ.get("COSMOS_DATABASE_NAME", DEFAULT_DATABASE_NAME),
        os.environ.get("COSMOS_CONTAINER_NAME", DEFAULT_CONTAINER_NAME),
    )


@dataclass
class SyncRepositoryResources:
    repository: SyncTodoRepository
    client: CosmosClient
    credential: DefaultAzureCredential

    def close(self) -> None:
        self.client.close()
        self.credential.close()


@dataclass
class AsyncRepositoryResources:
    repository: AsyncTodoRepository
    client: AsyncCosmosClient
    credential: AsyncDefaultAzureCredential

    async def close(self) -> None:
        await self.client.close()
        await self.credential.close()


def create_sync_repository() -> SyncRepositoryResources:
    endpoint, database_name, container_name = _settings()
    credential = DefaultAzureCredential()
    client = CosmosClient(url=endpoint, credential=credential)
    database = client.create_database_if_not_exists(id=database_name)
    container = database.create_container_if_not_exists(
        id=container_name,
        partition_key=PartitionKey(path="/category"),
        default_ttl=DEFAULT_TTL_SECONDS,
        indexing_policy=INDEXING_POLICY,
    )
    return SyncRepositoryResources(
        repository=SyncTodoRepository(container),
        client=client,
        credential=credential,
    )


async def create_async_repository() -> AsyncRepositoryResources:
    endpoint, database_name, container_name = _settings()
    credential = AsyncDefaultAzureCredential()
    client = AsyncCosmosClient(url=endpoint, credential=credential)
    database = await client.create_database_if_not_exists(id=database_name)
    container = await database.create_container_if_not_exists(
        id=container_name,
        partition_key=PartitionKey(path="/category"),
        default_ttl=DEFAULT_TTL_SECONDS,
        indexing_policy=INDEXING_POLICY,
    )
    return AsyncRepositoryResources(
        repository=AsyncTodoRepository(container),
        client=client,
        credential=credential,
    )

