from __future__ import annotations

import asyncio
import logging
from dataclasses import replace
from uuid import uuid4

from async_repository import AsyncTodoRepository
from cosmos_factory import async_container, sync_container
from sync_repository import SyncTodoRepository
from todo_model import TodoItem


def _demo_items(category: str, prefix: str) -> list[TodoItem]:
    return [
        TodoItem(
            id=f"{prefix}-{uuid4()}",
            title=f"Demo task {number}",
            description=f"Description for demo task {number}",
            completed=False,
            category=category,
        )
        for number in range(1, 4)
    ]


def run_sync_demo() -> None:
    category = f"sync-demo-{uuid4()}"
    with sync_container() as container:
        repository = SyncTodoRepository(container)
        created = [repository.create(item) for item in _demo_items(category, "sync")]
        print("Sync created:", created)

        current = repository.read(created[0].id, category)
        print("Sync read:", current)

        updated = repository.update(
            replace(current, title="Updated sync task", completed=True)
        )
        print("Sync updated:", updated)

        print("Sync category query:")
        for page_number, page in enumerate(
            repository.query_by_category(category, page_size=2), start=1
        ):
            print(f"  page {page_number}: {page}")

        for item in created:
            repository.delete(item.id, category)
        print("Sync deleted all demo items")


async def run_async_demo() -> None:
    category = f"async-demo-{uuid4()}"
    async with async_container() as container:
        repository = AsyncTodoRepository(container)
        created = [
            await repository.create(item)
            for item in _demo_items(category, "async")
        ]
        print("Async created:", created)

        current = await repository.read(created[0].id, category)
        print("Async read:", current)

        updated = await repository.update(
            replace(current, title="Updated async task", completed=True)
        )
        print("Async updated:", updated)

        print("Async category query:")
        page_number = 0
        async for page in repository.query_by_category(category, page_size=2):
            page_number += 1
            print(f"  page {page_number}: {page}")

        for item in created:
            await repository.delete(item.id, category)
        print("Async deleted all demo items")


def main() -> None:
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    print("Request charges are logged below for every Cosmos DB operation.")
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
