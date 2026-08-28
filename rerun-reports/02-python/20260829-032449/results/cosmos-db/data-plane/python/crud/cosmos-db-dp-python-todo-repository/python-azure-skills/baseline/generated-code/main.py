from __future__ import annotations

import asyncio
import logging
from dataclasses import replace
from uuid import uuid4

from cosmos_factory import create_async_repository, create_sync_repository
from models import OperationResult, TodoItem


def print_result(operation: str, result: OperationResult[object]) -> None:
    print(f"{operation}: {result.value} ({result.request_charge:.2f} RU)")


def run_sync_demo() -> None:
    print("\n=== Synchronous repository ===")
    resources = create_sync_repository()
    repository = resources.repository
    item_id = f"sync-{uuid4()}"
    category = "demo"

    try:
        created = repository.create(
            TodoItem(
                id=item_id,
                title="Try the synchronous repository",
                description="Run a complete CRUD cycle against Cosmos DB.",
                completed=False,
                category=category,
            )
        )
        print_result("Created", created)

        read = repository.read(item_id, category)
        print_result("Read", read)

        updated_item = replace(
            read.value,
            title="Synchronous repository complete",
            completed=True,
        )
        updated = repository.update(updated_item)
        print_result("Updated", updated)

        print(f"Querying category {category!r} page by page:")
        for page in repository.query_by_category(category, page_size=2):
            print(
                f"  Page {page.page_number} ({page.request_charge:.2f} RU): "
                f"{page.items}"
            )

        deleted = repository.delete(item_id, category)
        print_result("Deleted", deleted)
    finally:
        resources.close()


async def run_async_demo() -> None:
    print("\n=== Asynchronous repository ===")
    resources = await create_async_repository()
    repository = resources.repository
    item_id = f"async-{uuid4()}"
    category = "demo"

    try:
        created = await repository.create(
            TodoItem(
                id=item_id,
                title="Try the asynchronous repository",
                description="Run a complete CRUD cycle with azure.cosmos.aio.",
                completed=False,
                category=category,
            )
        )
        print_result("Created", created)

        read = await repository.read(item_id, category)
        print_result("Read", read)

        updated_item = replace(
            read.value,
            title="Asynchronous repository complete",
            completed=True,
        )
        updated = await repository.update(updated_item)
        print_result("Updated", updated)

        print(f"Querying category {category!r} page by page:")
        async for page in repository.query_by_category(category, page_size=2):
            print(
                f"  Page {page.page_number} ({page.request_charge:.2f} RU): "
                f"{page.items}"
            )

        deleted = await repository.delete(item_id, category)
        print_result("Deleted", deleted)
    finally:
        await resources.close()


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()

