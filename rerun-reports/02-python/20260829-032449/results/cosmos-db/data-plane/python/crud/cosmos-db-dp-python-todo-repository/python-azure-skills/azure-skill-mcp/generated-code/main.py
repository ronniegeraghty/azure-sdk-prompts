from __future__ import annotations

import asyncio
import logging
import sys
from dataclasses import replace
from typing import TypeVar

from async_repository import AsyncToDoRepository
from cosmos_factory import create_async_resources, create_sync_resources
from sync_repository import SyncToDoRepository
from todo_model import OperationResult, ToDoItem


T = TypeVar("T")


def print_result(operation: str, result: OperationResult[T]) -> None:
    print(f"{operation}: {result.request_charge:.2f} RU")
    if result.value is not None:
        print(f"  {result.value}")


def run_sync_demo() -> None:
    print("\n=== Synchronous repository ===")
    with create_sync_resources() as resources:
        repository = SyncToDoRepository(resources.container)
        created = repository.create(
            ToDoItem.new(
                title="Try the synchronous Cosmos repository",
                description="Create, read, update, query, and delete an item.",
                category="demo-sync",
            )
        )
        print_result("create", created)

        read = repository.read(created.value.id, created.value.category)
        print_result("read", read)

        updated = repository.update(
            replace(read.value, completed=True, title="Sync demo completed")
        )
        print_result("update", updated)

        print("query_by_category:")
        for page in repository.query_by_category("demo-sync", page_size=2):
            print(
                f"  page {page.number}: {page.request_charge:.2f} RU, "
                f"{len(page.items)} item(s)"
            )
            for item in page.items:
                print(f"    {item}")

        deleted = repository.delete(updated.value.id, updated.value.category)
        print_result("delete", deleted)


async def run_async_demo() -> None:
    print("\n=== Asynchronous repository ===")
    async with await create_async_resources() as resources:
        repository = AsyncToDoRepository(resources.container)
        created = await repository.create(
            ToDoItem.new(
                title="Try the asynchronous Cosmos repository",
                description="Create, read, update, query, and delete an item.",
                category="demo-async",
            )
        )
        print_result("create", created)

        read = await repository.read(created.value.id, created.value.category)
        print_result("read", read)

        updated = await repository.update(
            replace(read.value, completed=True, title="Async demo completed")
        )
        print_result("update", updated)

        print("query_by_category:")
        async for page in repository.query_by_category(
            "demo-async", page_size=2
        ):
            print(
                f"  page {page.number}: {page.request_charge:.2f} RU, "
                f"{len(page.items)} item(s)"
            )
            for item in page.items:
                print(f"    {item}")

        deleted = await repository.delete(
            updated.value.id, updated.value.category
        )
        print_result("delete", deleted)


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(levelname)s %(name)s: %(message)s",
        stream=sys.stdout,
    )
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
