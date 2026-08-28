from __future__ import annotations

import os
import sys
from typing import Any
from urllib.parse import urlparse

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.exceptions import (
    CosmosHttpResponseError,
    CosmosResourceNotFoundError,
)

DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"
ITEM_ID = "item-001"
ITEM_CATEGORY = "electronics"
LOCAL_HOSTS = {"localhost", "127.0.0.1", "::1"}


def get_local_configuration() -> tuple[str, str]:
    endpoint = os.getenv("COSMOS_ENDPOINT", "https://localhost:8081")
    key = os.getenv("COSMOS_KEY")

    if urlparse(endpoint).hostname not in LOCAL_HOSTS:
        raise ValueError(
            "COSMOS_ENDPOINT must target a local Cosmos DB emulator."
        )
    if not key:
        raise ValueError(
            "Set COSMOS_KEY to the key shown by your local Cosmos DB emulator."
        )

    return endpoint, key


def run_crud_operations() -> None:
    endpoint, key = get_local_configuration()

    with CosmosClient(
        url=endpoint,
        credential=key,
        connection_verify=False,
    ) as client:
        database = client.create_database_if_not_exists(id=DATABASE_NAME)
        container = database.create_container_if_not_exists(
            id=CONTAINER_NAME,
            partition_key=PartitionKey(path=PARTITION_KEY_PATH),
            offer_throughput=400,
        )

        item: dict[str, Any] = {
            "id": ITEM_ID,
            "category": ITEM_CATEGORY,
            "name": "Wireless Headphones",
            "quantity": 10,
        }

        upserted_item = container.upsert_item(body=item)
        print(f"Upserted: {upserted_item}")

        read_item = container.read_item(
            item=ITEM_ID,
            partition_key=ITEM_CATEGORY,
        )
        print(f"Read: {read_item}")

        query = "SELECT * FROM items i WHERE i.category = @category"
        parameters = [{"name": "@category", "value": ITEM_CATEGORY}]
        queried_items = list(
            container.query_items(
                query=query,
                parameters=parameters,
                partition_key=ITEM_CATEGORY,
            )
        )
        print(f"Query results: {queried_items}")

        updated_item = dict(read_item)
        updated_item["quantity"] = 25
        replaced_item = container.replace_item(
            item=ITEM_ID,
            body=updated_item,
        )
        print(f"Replaced: {replaced_item}")

        container.delete_item(
            item=ITEM_ID,
            partition_key=ITEM_CATEGORY,
        )
        print(f"Deleted item {ITEM_ID!r}.")


def main() -> int:
    try:
        run_crud_operations()
    except CosmosResourceNotFoundError as error:
        print(f"Cosmos DB resource was not found: {error}", file=sys.stderr)
        return 1
    except CosmosHttpResponseError as error:
        print(
            f"Cosmos DB request failed "
            f"(status {error.status_code}): {error.message}",
            file=sys.stderr,
        )
        return 1
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
