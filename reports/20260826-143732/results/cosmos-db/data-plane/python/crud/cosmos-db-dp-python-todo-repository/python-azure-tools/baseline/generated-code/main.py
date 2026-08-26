from __future__ import annotations

import asyncio
import logging
from dataclasses import replace
from uuid import uuid4

from async_repository import AsyncToDoRepository
from cosmos_factory import create_async_resources, create_sync_resources
from sync_repository import ToDoRepository
from todo_model import ToDoItem


def print_operation(operation: str, result: object, request_charge: float) -> None:
    print(f"{operation}: {result} (RU: {request_charge:.2f})")


def run_sync_demo(repository: ToDoRepository) -> None:
    print("\n=== Synchronous repository ===")
    category = f"sync-demo-{uuid4()}"
    created_items: list[ToDoItem] = []

    for number in range(1, 4):
        item = ToDoItem.new(
            id=str(uuid4()),
            title=f"Sync task {number}",
            description=f"Created by the synchronous demo ({number})",
            category=category,
        )
        created = repository.create(item)
        created_items.append(created)
        print_operation("create", created, repository.last_request_charge)

    current = repository.read(created_items[0].id, category)
    print_operation("read", current, repository.last_request_charge)

    updated = repository.update(
        replace(current, completed=True, title="Completed sync task")
    )
    print_operation("update", updated, repository.last_request_charge)

    print("query by category:")
    for page in repository.query_by_category(category, page_size=2):
        print(
            f"  page {page.page_number} (RU: {page.request_charge:.2f}): "
            f"{page.items}"
        )

    for item in created_items:
        repository.delete(item.id, category)
        print_operation("delete", item.id, repository.last_request_charge)


async def run_async_demo(repository: AsyncToDoRepository) -> None:
    print("\n=== Asynchronous repository ===")
    category = f"async-demo-{uuid4()}"
    created_items: list[ToDoItem] = []

    for number in range(1, 4):
        item = ToDoItem.new(
            id=str(uuid4()),
            title=f"Async task {number}",
            description=f"Created by the asynchronous demo ({number})",
            category=category,
        )
        created = await repository.create(item)
        created_items.append(created)
        print_operation("create", created, repository.last_request_charge)

    current = await repository.read(created_items[0].id, category)
    print_operation("read", current, repository.last_request_charge)

    updated = await repository.update(
        replace(current, completed=True, title="Completed async task")
    )
    print_operation("update", updated, repository.last_request_charge)

    print("query by category:")
    async for page in repository.query_by_category(category, page_size=2):
        print(
            f"  page {page.page_number} (RU: {page.request_charge:.2f}): "
            f"{page.items}"
        )

    for item in created_items:
        await repository.delete(item.id, category)
        print_operation("delete", item.id, repository.last_request_charge)


async def main() -> None:
    with create_sync_resources() as sync_resources:
        run_sync_demo(sync_resources.repository)

    async with await create_async_resources() as async_resources:
        await run_async_demo(async_resources.repository)


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s %(message)s")
    asyncio.run(main())
