import os
from urllib.parse import urlparse

from azure.cosmos import CosmosClient, PartitionKey, exceptions


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"
LOCAL_HOSTS = {"localhost", "127.0.0.1", "::1"}


def get_local_connection() -> tuple[str, str]:
    endpoint = os.getenv("COSMOS_ENDPOINT", "https://localhost:8081")
    key = os.getenv("COSMOS_KEY")

    if urlparse(endpoint).hostname not in LOCAL_HOSTS:
        raise ValueError("COSMOS_ENDPOINT must point to a local Cosmos DB Emulator.")
    if not key:
        raise ValueError("Set COSMOS_KEY to the Cosmos DB Emulator primary key.")

    return endpoint, key


def main() -> None:
    endpoint, key = get_local_connection()
    client = CosmosClient(endpoint, credential=key)

    item = {
        "id": "item-001",
        "category": "electronics",
        "name": "Wireless Mouse",
        "quantity": 10,
    }

    try:
        database = client.create_database_if_not_exists(id=DATABASE_NAME)
        container = database.create_container_if_not_exists(
            id=CONTAINER_NAME,
            partition_key=PartitionKey(path=PARTITION_KEY_PATH),
        )

        upserted_item = container.upsert_item(body=item)
        print(f"Upserted: {upserted_item}")

        read_item = container.read_item(
            item=item["id"],
            partition_key=item["category"],
        )
        print(f"Read: {read_item}")

        query = "SELECT * FROM items i WHERE i.category = @category"
        parameters = [{"name": "@category", "value": "electronics"}]
        queried_items = list(
            container.query_items(
                query=query,
                parameters=parameters,
                partition_key="electronics",
            )
        )
        print(f"Query results: {queried_items}")

        updated_item = dict(read_item)
        updated_item["quantity"] = 25
        replaced_item = container.replace_item(
            item=item["id"],
            body=updated_item,
        )
        print(f"Replaced: {replaced_item}")

        container.delete_item(
            item=item["id"],
            partition_key=item["category"],
        )
        print(f"Deleted item: {item['id']}")
    except exceptions.CosmosResourceNotFoundError as exc:
        print(f"Cosmos DB resource was not found: {exc}")
        raise
    except exceptions.CosmosHttpResponseError as exc:
        print(f"Cosmos DB request failed (status {exc.status_code}): {exc}")
        raise
    finally:
        client.close()


if __name__ == "__main__":
    try:
        main()
    except (ValueError, exceptions.CosmosHttpResponseError) as exc:
        raise SystemExit(1) from exc
