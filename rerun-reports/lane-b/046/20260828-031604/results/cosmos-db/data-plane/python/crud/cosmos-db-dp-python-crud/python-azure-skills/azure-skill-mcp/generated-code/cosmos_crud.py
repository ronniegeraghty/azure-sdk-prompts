import os
import sys

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.exceptions import CosmosHttpResponseError


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
ITEM_ID = "item-001"
CATEGORY = "electronics"


def main() -> int:
    endpoint = os.environ.get("COSMOS_ENDPOINT", "https://localhost:8081/")
    key = os.environ.get("COSMOS_KEY")
    if not key:
        print(
            "COSMOS_KEY is required. Set it to the local Cosmos DB emulator key.",
            file=sys.stderr,
        )
        return 2

    client = CosmosClient(endpoint, credential=key)

    try:
        database = client.create_database_if_not_exists(id=DATABASE_NAME)
        container = database.create_container_if_not_exists(
            id=CONTAINER_NAME,
            partition_key=PartitionKey(path="/category"),
        )

        item = {
            "id": ITEM_ID,
            "category": CATEGORY,
            "name": "Wireless Headphones",
            "quantity": 10,
        }
        upserted_item = container.upsert_item(body=item)
        print(f"Upserted: {upserted_item}")

        read_item = container.read_item(
            item=ITEM_ID,
            partition_key=CATEGORY,
        )
        print(f"Read: {read_item}")

        query = "SELECT * FROM c WHERE c.category = @category"
        parameters = [{"name": "@category", "value": CATEGORY}]
        queried_items = list(
            container.query_items(
                query=query,
                parameters=parameters,
                partition_key=CATEGORY,
            )
        )
        print(f"Query results: {queried_items}")

        read_item["quantity"] = 25
        replaced_item = container.replace_item(
            item=ITEM_ID,
            body=read_item,
        )
        print(f"Replaced: {replaced_item}")

        container.delete_item(
            item=ITEM_ID,
            partition_key=CATEGORY,
        )
        print(f"Deleted item: {ITEM_ID}")
    except CosmosHttpResponseError as exc:
        print(
            f"Cosmos DB request failed ({exc.status_code}): {exc.message}",
            file=sys.stderr,
        )
        return 1
    finally:
        client.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
