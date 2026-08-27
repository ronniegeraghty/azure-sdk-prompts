import os
import sys

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.exceptions import (
    CosmosHttpResponseError,
    CosmosResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY = "electronics"
ITEM_ID = "item-001"


def run_crud_operations() -> None:
    endpoint = os.environ.get("COSMOS_ENDPOINT")
    if not endpoint:
        raise ValueError(
            "COSMOS_ENDPOINT is required, for example "
            "https://<account>.documents.azure.com:443/"
        )

    item = {
        "id": ITEM_ID,
        "category": PARTITION_KEY,
        "name": "Laptop",
        "quantity": 10,
    }

    with DefaultAzureCredential() as credential:
        with CosmosClient(url=endpoint, credential=credential) as client:
            database = client.create_database_if_not_exists(id=DATABASE_NAME)
            container = database.create_container_if_not_exists(
                id=CONTAINER_NAME,
                partition_key=PartitionKey(path="/category"),
            )

            upserted_item = container.upsert_item(body=item)
            print(f"Upserted item: {upserted_item}")

            read_item = container.read_item(
                item=ITEM_ID,
                partition_key=PARTITION_KEY,
            )
            print(f"Read item: {read_item}")

            query = "SELECT * FROM c WHERE c.category = @category"
            queried_items = list(
                container.query_items(
                    query=query,
                    parameters=[
                        {"name": "@category", "value": PARTITION_KEY},
                    ],
                    partition_key=PARTITION_KEY,
                )
            )
            print(f"Queried items: {queried_items}")

            read_item["quantity"] = 15
            replaced_item = container.replace_item(
                item=ITEM_ID,
                body=read_item,
            )
            print(f"Replaced item: {replaced_item}")

            container.delete_item(
                item=ITEM_ID,
                partition_key=PARTITION_KEY,
            )
            print(f"Deleted item: {ITEM_ID}")


def main() -> int:
    try:
        run_crud_operations()
    except ValueError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2
    except CosmosResourceNotFoundError as error:
        print(f"Cosmos DB resource was not found: {error}", file=sys.stderr)
        return 1
    except CosmosHttpResponseError as error:
        print(
            f"Cosmos DB request failed (status {error.status_code}): {error}",
            file=sys.stderr,
        )
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
