from __future__ import annotations

import asyncio
import logging

from cosmos_factory import create_async_repository, create_sync_repository
from todo_model import TodoItem

logger = logging.getLogger(__name__)


def print_page(label: str, page_number: int, items: list[TodoItem]) -> None:
    print(f"{label} page {page_number}:")
    for item in items:
        print(f"  {item}")


def run_sync_demo() -> None:
    print("\n=== Synchronous repository ===")
    with create_sync_repository() as repository:
        created = repository.create(
            TodoItem(
                title="Try the synchronous repository",
                description="Run a complete Cosmos DB CRUD cycle.",
                category="demo-sync",
            )
        )
        print("Created:", created)

        current = repository.read(created.id, created.category)
        print("Read:", current)

        current.completed = True
        updated = repository.update(current)
        print("Updated:", updated)

        for page_number, page in enumerate(
            repository.query_by_category(updated.category, page_size=2), start=1
        ):
            print_page("Sync query", page_number, page)

        repository.delete(updated.id, updated.category)
        print("Deleted:", updated.id)


async def run_async_demo() -> None:
    print("\n=== Asynchronous repository ===")
    async with create_async_repository() as repository:
        created = await repository.create(
            TodoItem(
                title="Try the asynchronous repository",
                description="Run a complete async Cosmos DB CRUD cycle.",
                category="demo-async",
            )
        )
        print("Created:", created)

        current = await repository.read(created.id, created.category)
        print("Read:", current)

        current.completed = True
        updated = await repository.update(current)
        print("Updated:", updated)

        page_number = 0
        async for page in repository.query_by_category(
            updated.category, page_size=2
        ):
            page_number += 1
            print_page("Async query", page_number, page)

        await repository.delete(updated.id, updated.category)
        print("Deleted:", updated.id)


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
