"""Run a complete CRUD workflow against an Azure Cosmos DB NoSQL container."""

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


def run_crud_workflow() -> None:
    endpoint = os.getenv("COSMOS_ENDPOINT")
    if not endpoint:
        raise RuntimeError("Set COSMOS_ENDPOINT to the Azure Cosmos DB account endpoint.")

    credential = DefaultAzureCredential()
    try:
        with CosmosClient(url=endpoint, credential=credential) as client:
            database = client.create_database_if_not_exists(id=DATABASE_NAME)
            container = database.create_container_if_not_exists(
                id=CONTAINER_NAME,
                partition_key=PartitionKey(path=PARTITION_KEY_PATH),
            )

            item = {
                "id": "item-001",
                "category": "electronics",
                "name": "Wireless Keyboard",
                "quantity": 10,
            }

            upserted_item = container.upsert_item(body=item)
            logging.info("Upserted item %s.", upserted_item["id"])

            read_item = container.read_item(
                item=item["id"],
                partition_key=item["category"],
            )
            logging.info("Read item: %s", read_item)

            query = "SELECT * FROM c WHERE c.category = @category"
            queried_items = list(
                container.query_items(
                    query=query,
                    parameters=[
                        {"name": "@category", "value": "electronics"},
                    ],
                    partition_key="electronics",
                )
            )
            logging.info("Query returned %d item(s): %s", len(queried_items), queried_items)

            read_item["quantity"] = 25
            replaced_item = container.replace_item(
                item=read_item["id"],
                body=read_item,
            )
            logging.info(
                "Replaced item %s with quantity %d.",
                replaced_item["id"],
                replaced_item["quantity"],
            )

            container.delete_item(
                item=replaced_item["id"],
                partition_key=replaced_item["category"],
            )
            logging.info("Deleted item %s.", replaced_item["id"])
    finally:
        credential.close()


def main() -> int:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")

    try:
        run_crud_workflow()
    except CosmosResourceNotFoundError as error:
        logging.error("Cosmos DB resource was not found: %s", error)
        return 1
    except CosmosHttpResponseError as error:
        logging.error(
            "Cosmos DB request failed (status %s): %s",
            error.status_code,
            error.message,
        )
        return 1
    except RuntimeError as error:
        logging.error("%s", error)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
