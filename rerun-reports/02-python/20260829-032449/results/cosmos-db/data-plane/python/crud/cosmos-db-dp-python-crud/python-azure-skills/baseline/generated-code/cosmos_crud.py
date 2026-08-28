import os
import sys

from azure.cosmos import CosmosClient, PartitionKey, exceptions


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"


def required_environment_variable(name: str) -> str:
    value = os.environ.get(name)
    if not value:
        raise ValueError(f"Required environment variable {name} is not set.")
    return value


def main() -> None:
    endpoint = required_environment_variable("COSMOS_ENDPOINT")
    key = required_environment_variable("COSMOS_KEY")

    client = CosmosClient(endpoint, credential=key)
    database = client.create_database_if_not_exists(id=DATABASE_NAME)
    container = database.create_container_if_not_exists(
        id=CONTAINER_NAME,
        partition_key=PartitionKey(path="/category"),
    )

    item = {
        "id": "item-001",
        "category": "electronics",
        "name": "Wireless Mouse",
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

    updated_item = dict(read_item)
    updated_item["quantity"] = 20
    replaced_item = container.replace_item(
        item=updated_item["id"],
        body=updated_item,
    )
    print(f"Replaced item: {replaced_item}")

    container.delete_item(
        item=updated_item["id"],
        partition_key=updated_item["category"],
    )
    print(f"Deleted item: {updated_item['id']}")


if __name__ == "__main__":
    try:
        main()
    except exceptions.CosmosResourceNotFoundError as error:
        print(f"Cosmos DB resource was not found: {error}", file=sys.stderr)
        sys.exit(1)
    except exceptions.CosmosHttpResponseError as error:
        print(f"Cosmos DB request failed: {error}", file=sys.stderr)
        sys.exit(1)
    except ValueError as error:
        print(error, file=sys.stderr)
        sys.exit(1)
