from __future__ import annotations

import os
from dataclasses import dataclass

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from async_repository import AsyncToDoRepository
from sync_repository import ToDoRepository

DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60
INDEXING_POLICY = {
    "automatic": True,
    "indexingMode": "consistent",
    "includedPaths": [{"path": "/*"}],
    "excludedPaths": [{"path": "/description/?"}],
}


def _settings() -> tuple[str, str, str]:
    endpoint = os.environ.get("COSMOS_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "COSMOS_ENDPOINT must contain the Azure Cosmos DB account endpoint."
        )
    database_name = os.environ.get("COSMOS_DATABASE", "todo-db")
    container_name = os.environ.get("COSMOS_CONTAINER", "todo-items")
    return endpoint, database_name, container_name


@dataclass
class SyncCosmosResources:
    repository: ToDoRepository
    client: CosmosClient
    credential: DefaultAzureCredential

    def close(self) -> None:
        self.client.close()
        self.credential.close()

    def __enter__(self) -> "SyncCosmosResources":
        return self

    def __exit__(self, *_args: object) -> None:
        self.close()


@dataclass
class AsyncCosmosResources:
    repository: AsyncToDoRepository
    client: AsyncCosmosClient
    credential: AsyncDefaultAzureCredential

    async def close(self) -> None:
        await self.client.close()
        await self.credential.close()

    async def __aenter__(self) -> "AsyncCosmosResources":
        return self

    async def __aexit__(self, *_args: object) -> None:
        await self.close()


def create_sync_resources() -> SyncCosmosResources:
    endpoint, database_name, container_name = _settings()
    credential = DefaultAzureCredential()
    client = CosmosClient(endpoint, credential=credential)
    try:
        database = client.create_database_if_not_exists(id=database_name)
        container = database.create_container_if_not_exists(
            id=container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        client.close()
        credential.close()
        raise
    return SyncCosmosResources(ToDoRepository(container), client, credential)


async def create_async_resources() -> AsyncCosmosResources:
    endpoint, database_name, container_name = _settings()
    credential = AsyncDefaultAzureCredential()
    client = AsyncCosmosClient(endpoint, credential=credential)
    try:
        database = await client.create_database_if_not_exists(id=database_name)
        container = await database.create_container_if_not_exists(
            id=container_name,
            partition_key=PartitionKey(path="/category"),
            default_ttl=DEFAULT_TTL_SECONDS,
            indexing_policy=INDEXING_POLICY,
        )
    except Exception:
        await client.close()
        await credential.close()
        raise
    return AsyncCosmosResources(
        AsyncToDoRepository(container), client, credential
    )
