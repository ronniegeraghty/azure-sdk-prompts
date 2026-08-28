from __future__ import annotations

import os
from contextlib import asynccontextmanager, contextmanager
from collections.abc import AsyncIterator, Iterator

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos import ContainerProxy as SyncContainerProxy
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.cosmos.aio import ContainerProxy as AsyncContainerProxy
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

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
            "Set COSMOS_ENDPOINT to the Azure Cosmos DB account endpoint."
        )

    return (
        endpoint,
        os.environ.get("COSMOS_DATABASE", "todo-db"),
        os.environ.get("COSMOS_CONTAINER", "todos"),
    )


@contextmanager
def sync_container() -> Iterator[SyncContainerProxy]:
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
        yield container
    finally:
        client.close()
        credential.close()


@asynccontextmanager
async def async_container() -> AsyncIterator[AsyncContainerProxy]:
    endpoint, database_name, container_name = _settings()
    credential = AsyncDefaultAzureCredential()
    client = AsyncCosmosClient(endpoint, credential=credential)

    try:
        async with client:
            database = await client.create_database_if_not_exists(id=database_name)
            container = await database.create_container_if_not_exists(
                id=container_name,
                partition_key=PartitionKey(path="/category"),
                default_ttl=DEFAULT_TTL_SECONDS,
                indexing_policy=INDEXING_POLICY,
            )
            yield container
    finally:
        await credential.close()
