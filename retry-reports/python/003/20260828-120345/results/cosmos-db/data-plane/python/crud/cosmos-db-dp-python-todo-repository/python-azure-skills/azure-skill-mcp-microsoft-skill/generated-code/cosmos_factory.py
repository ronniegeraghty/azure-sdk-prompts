from __future__ import annotations

import os
from collections.abc import AsyncIterator, Iterator
from contextlib import asynccontextmanager, contextmanager
from dataclasses import dataclass

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.aio import CosmosClient as AsyncCosmosClient
from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential

from async_repository import AsyncTodoRepository
from sync_repository import SyncTodoRepository

DEFAULT_TTL_SECONDS = 90 * 24 * 60 * 60
INDEXING_POLICY = {
    "indexingMode": "consistent",
    "automatic": True,
    "includedPaths": [{"path": "/*"}],
    "excludedPaths": [{"path": "/description/?"}],
}


@dataclass(frozen=True, slots=True)
class CosmosSettings:
    endpoint: str
    database_name: str
    container_name: str

    @classmethod
    def from_environment(cls) -> CosmosSettings:
        endpoint = os.environ.get("COSMOS_ENDPOINT")
        if not endpoint:
            raise RuntimeError("COSMOS_ENDPOINT environment variable is required")

        return cls(
            endpoint=endpoint,
            database_name=os.environ.get("COSMOS_DATABASE_NAME", "todo-database"),
            container_name=os.environ.get("COSMOS_CONTAINER_NAME", "todo-items"),
        )


@contextmanager
def create_sync_repository(
    settings: CosmosSettings | None = None,
) -> Iterator[SyncTodoRepository]:
    settings = settings or CosmosSettings.from_environment()
    with DefaultAzureCredential() as credential:
        with CosmosClient(url=settings.endpoint, credential=credential) as client:
            database = client.create_database_if_not_exists(id=settings.database_name)
            container = database.create_container_if_not_exists(
                id=settings.container_name,
                partition_key=PartitionKey(path="/category"),
                default_ttl=DEFAULT_TTL_SECONDS,
                indexing_policy=INDEXING_POLICY,
            )
            yield SyncTodoRepository(container)


@asynccontextmanager
async def create_async_repository(
    settings: CosmosSettings | None = None,
) -> AsyncIterator[AsyncTodoRepository]:
    settings = settings or CosmosSettings.from_environment()
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncCosmosClient(
            url=settings.endpoint, credential=credential
        ) as client:
            database = await client.create_database_if_not_exists(
                id=settings.database_name
            )
            container = await database.create_container_if_not_exists(
                id=settings.container_name,
                partition_key=PartitionKey(path="/category"),
                default_ttl=DEFAULT_TTL_SECONDS,
                indexing_policy=INDEXING_POLICY,
            )
            yield AsyncTodoRepository(container)
