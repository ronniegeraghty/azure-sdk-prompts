import os
import sys

from azure.cosmos import CosmosClient, PartitionKey, exceptions


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"


def main() -> None:
    endpoint = os.environ["COSMOS_ENDPOINT"]
    key = os.environ["COSMOS_KEY"]

    try:
        client = CosmosClient(endpoint, credential=key)
        database = client.create_database_if_not_exists(id=DATABASE_NAME)
        container = database.create_container_if_not_exists(
            id=CONTAINER_NAME,
            partition_key=PartitionKey(path="/category"),
        )

        item = {
            "id": "item-001",
            "category": "electronics",
            "name": "Wireless keyboard",
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
        parameters = [
            {"name": "@category", "value": "electronics"},
        ]
        queried_items = list(
            container.query_items(
                query=query,
                parameters=parameters,
                partition_key="electronics",
            )
        )
        print(f"Queried items: {queried_items}")

        read_item["quantity"] = 25
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
    except exceptions.CosmosResourceNotFoundError as error:
        print(f"Cosmos DB resource was not found: {error}", file=sys.stderr)
        raise
    except exceptions.CosmosAccessConditionFailedError as error:
        print(f"Cosmos DB access condition failed: {error}", file=sys.stderr)
        raise
    except exceptions.CosmosHttpResponseError as error:
        print(
            f"Cosmos DB request failed with status {error.status_code}: {error}",
            file=sys.stderr,
        )
        raise


if __name__ == "__main__":
    try:
        main()
    except KeyError as error:
        variable = error.args[0]
        print(
            f"Missing required environment variable: {variable}",
            file=sys.stderr,
        )
        sys.exit(2)
    except exceptions.CosmosHttpResponseError:
        sys.exit(1)
