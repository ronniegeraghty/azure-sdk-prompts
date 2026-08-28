"""Perform a complete CRUD lifecycle in an Azure Cosmos DB for NoSQL container."""

import os
import sys
from typing import Any

from azure.cosmos import CosmosClient, PartitionKey, exceptions


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"


def run_crud_operations() -> None:
    endpoint = os.environ["COSMOS_ENDPOINT"]
    key = os.environ["COSMOS_KEY"]

    client = CosmosClient(endpoint, credential=key)
    database = client.create_database_if_not_exists(id=DATABASE_NAME)
    container = database.create_container_if_not_exists(
        id=CONTAINER_NAME,
        partition_key=PartitionKey(path=PARTITION_KEY_PATH),
        offer_throughput=400,
    )

    item: dict[str, Any] = {
        "id": "item-001",
        "category": "electronics",
        "name": "Wireless Mouse",
        "quantity": 10,
    }

    upserted_item = container.upsert_item(body=item)
    print(f"Upserted item: {upserted_item}")

    read_item = container.read_item(
        item=item["id"],
        partition_key=item["category"],
    )
    print(f"Read item: {read_item}")

    query = "SELECT * FROM items i WHERE i.category = @category"
    parameters = [{"name": "@category", "value": "electronics"}]
    queried_items = list(
        container.query_items(
            query=query,
            parameters=parameters,
            partition_key="electronics",
        )
    )
    print(f"Queried items: {queried_items}")

    item["quantity"] = 15
    replaced_item = container.replace_item(item=item["id"], body=item)
    print(f"Replaced item: {replaced_item}")

    container.delete_item(
        item=item["id"],
        partition_key=item["category"],
    )
    print(f"Deleted item: {item['id']}")


def main() -> int:
    try:
        run_crud_operations()
    except KeyError as error:
        print(
            f"Missing required environment variable: {error.args[0]}",
            file=sys.stderr,
        )
        return 2
    except exceptions.CosmosHttpResponseError as error:
        print(
            f"Cosmos DB request failed ({error.status_code}): {error.message}",
            file=sys.stderr,
        )
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
