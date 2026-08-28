from __future__ import annotations

import asyncio
import logging
from dataclasses import replace

from cosmos_factory import create_async_repository, create_sync_repository
from todo_model import OperationResult, TodoItem


def print_result(operation: str, result: OperationResult[object]) -> None:
    print(f"{operation}: {result.value!r} ({result.request_charge:.2f} RU)")


def run_sync_demo() -> None:
    print("\n=== Synchronous repository ===")
    with create_sync_repository() as repository:
        created = repository.create(
            TodoItem.new(
                title="Try the synchronous repository",
                description="Exercise create, read, update, query, and delete.",
                category="demo-sync",
            )
        )
        print_result("create", created)

        read = repository.read(created.value.id, created.value.category)
        print_result("read", read)

        updated = repository.update(
            replace(read.value, completed=True, title="Sync CRUD complete")
        )
        print_result("update", updated)

        for page in repository.query_by_category("demo-sync", page_size=2):
            print(
                f"query page {page.number}: {page.items!r} "
                f"({page.request_charge:.2f} RU)"
            )

        deleted = repository.delete(updated.value.id, updated.value.category)
        print_result("delete", deleted)


async def run_async_demo() -> None:
    print("\n=== Asynchronous repository ===")
    async with create_async_repository() as repository:
        created = await repository.create(
            TodoItem.new(
                title="Try the asynchronous repository",
                description="Exercise create, read, update, query, and delete.",
                category="demo-async",
            )
        )
        print_result("create", created)

        read = await repository.read(created.value.id, created.value.category)
        print_result("read", read)

        updated = await repository.update(
            replace(read.value, completed=True, title="Async CRUD complete")
        )
        print_result("update", updated)

        async for page in repository.query_by_category(
            "demo-async", page_size=2
        ):
            print(
                f"query page {page.number}: {page.items!r} "
                f"({page.request_charge:.2f} RU)"
            )

        deleted = await repository.delete(
            updated.value.id, updated.value.category
        )
        print_result("delete", deleted)


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    run_sync_demo()
    asyncio.run(run_async_demo())


if __name__ == "__main__":
    main()
