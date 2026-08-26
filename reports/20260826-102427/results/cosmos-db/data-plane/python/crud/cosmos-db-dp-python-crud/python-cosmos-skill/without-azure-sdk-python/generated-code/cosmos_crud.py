import logging
import os

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.exceptions import (
    CosmosHttpResponseError,
    CosmosResourceNotFoundError,
)
from azure.identity import DefaultAzureCredential


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"


def main() -> None:
    endpoint = os.environ.get("AZURE_COSMOS_ENDPOINT")
    if not endpoint:
        raise RuntimeError("Set AZURE_COSMOS_ENDPOINT to the Cosmos DB account endpoint.")

    credential = DefaultAzureCredential()
    client = CosmosClient(url=endpoint, credential=credential)

    item = {
        "id": "item-001",
        "category": "electronics",
        "name": "Wireless keyboard",
        "quantity": 10,
    }

    try:
        database = client.create_database_if_not_exists(id=DATABASE_NAME)
        container = database.create_container_if_not_exists(
            id=CONTAINER_NAME,
            partition_key=PartitionKey(path=PARTITION_KEY_PATH),
        )

        container.upsert_item(body=item)
        logging.info("Upserted item %s.", item["id"])

        stored_item = container.read_item(
            item=item["id"],
            partition_key=item["category"],
        )
        logging.info("Read item: %s", stored_item)

        query = "SELECT * FROM items i WHERE i.category = @category"
        parameters = [{"name": "@category", "value": "electronics"}]
        matching_items = list(
            container.query_items(
                query=query,
                parameters=parameters,
                partition_key="electronics",
            )
        )
        logging.info("Query returned %d item(s): %s", len(matching_items), matching_items)

        stored_item["quantity"] = 25
        replaced_item = container.replace_item(
            item=stored_item["id"],
            body=stored_item,
        )
        logging.info("Replaced item with quantity %s.", replaced_item["quantity"])

        container.delete_item(
            item=replaced_item["id"],
            partition_key=replaced_item["category"],
        )
        logging.info("Deleted item %s.", replaced_item["id"])
    except CosmosResourceNotFoundError as error:
        logging.error("A Cosmos DB resource was not found: %s", error)
        raise
    except CosmosHttpResponseError as error:
        logging.error("Cosmos DB request failed: %s", error)
        raise
    finally:
        client.close()
        credential.close()


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    main()
