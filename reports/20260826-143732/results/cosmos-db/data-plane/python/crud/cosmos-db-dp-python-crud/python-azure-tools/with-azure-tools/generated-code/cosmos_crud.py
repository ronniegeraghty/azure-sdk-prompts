import logging
import os
import sys
from typing import Any

from azure.cosmos import CosmosClient, PartitionKey
from azure.cosmos.exceptions import CosmosHttpResponseError
from azure.identity import DefaultAzureCredential


DATABASE_NAME = "TestDB"
CONTAINER_NAME = "Items"
PARTITION_KEY_PATH = "/category"
ITEM_ID = "item-001"
ITEM_CATEGORY = "electronics"

logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
logger = logging.getLogger(__name__)


def run_crud_operations() -> None:
    endpoint = os.environ.get("COSMOS_ENDPOINT")
    if not endpoint:
        raise ValueError(
            "COSMOS_ENDPOINT must be set, for example "
            "'https://<account>.documents.azure.com:443/'."
        )

    credential = DefaultAzureCredential()
    item: dict[str, Any] = {
        "id": ITEM_ID,
        "category": ITEM_CATEGORY,
        "name": "Laptop",
        "quantity": 10,
    }

    try:
        with CosmosClient(url=endpoint, credential=credential) as client:
            database = client.create_database_if_not_exists(id=DATABASE_NAME)
            container = database.create_container_if_not_exists(
                id=CONTAINER_NAME,
                partition_key=PartitionKey(path=PARTITION_KEY_PATH),
            )

            upserted_item = container.upsert_item(body=item)
            logger.info("Upserted item %s", upserted_item["id"])

            read_item = container.read_item(
                item=ITEM_ID,
                partition_key=ITEM_CATEGORY,
            )
            logger.info("Read item: %s", read_item)

            query = "SELECT * FROM c WHERE c.category = @category"
            query_results = list(
                container.query_items(
                    query=query,
                    parameters=[
                        {"name": "@category", "value": ITEM_CATEGORY},
                    ],
                    partition_key=ITEM_CATEGORY,
                )
            )
            logger.info(
                "Found %d item(s) in category %s",
                len(query_results),
                ITEM_CATEGORY,
            )

            read_item["quantity"] = 20
            replaced_item = container.replace_item(
                item=read_item["id"],
                body=read_item,
            )
            logger.info(
                "Replaced item %s; quantity is now %s",
                replaced_item["id"],
                replaced_item["quantity"],
            )

            container.delete_item(
                item=ITEM_ID,
                partition_key=ITEM_CATEGORY,
            )
            logger.info("Deleted item %s", ITEM_ID)
    except CosmosHttpResponseError as error:
        logger.error(
            "Cosmos DB request failed (status %s): %s",
            error.status_code,
            error.message,
        )
        raise
    finally:
        credential.close()


if __name__ == "__main__":
    try:
        run_crud_operations()
    except (CosmosHttpResponseError, ValueError):
        sys.exit(1)
