"""Perform basic CRUD operations against an Azure Cosmos DB NoSQL container."""

import os
import sys
from typing import Any

from azure.cosmos import CosmosClient, PartitionKey, exceptions


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"


def require_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    endpoint = require_environment_variable("COSMOS_ENDPOINT")
    key = require_environment_variable("COSMOS_KEY")

    client = CosmosClient(endpoint, credential=key)
    database = client.create_database_if_not_exists(id=DATABASE_NAME)
    container = database.create_container_if_not_exists(
        id=CONTAINER_NAME,
        partition_key=PartitionKey(path=PARTITION_KEY_PATH),
    )

    item: dict[str, Any] = {
        "id": "item-001",
        "category": "electronics",
        "name": "Wireless Keyboard",
        "quantity": 10,
    }

    upserted_item = container.upsert_item(item)
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

    read_item["quantity"] = 15
    replaced_item = container.replace_item(
        item=read_item["id"],
        body=read_item,
    )
    print(f"Replaced item: {replaced_item}")

    container.delete_item(
        item=read_item["id"],
        partition_key=read_item["category"],
    )
    print(f"Deleted item: {read_item['id']}")


if __name__ == "__main__":
    try:
        main()
    except RuntimeError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        sys.exit(1)
    except exceptions.CosmosResourceNotFoundError as error:
        print(f"Cosmos DB resource not found: {error}", file=sys.stderr)
        sys.exit(1)
    except exceptions.CosmosHttpResponseError as error:
        print(
            f"Cosmos DB request failed "
            f"(status={error.status_code}, message={error.message})",
            file=sys.stderr,
        )
        sys.exit(1)
