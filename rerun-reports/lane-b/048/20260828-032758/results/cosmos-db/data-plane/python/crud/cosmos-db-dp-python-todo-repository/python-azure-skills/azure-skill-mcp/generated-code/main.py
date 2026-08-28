from __future__ import annotations

import asyncio
import logging
from dataclasses import replace

from todo_cosmos.factory import (
    CosmosSettings,
    create_async_repository,
    create_sync_repository,
)
from todo_cosmos.model import TodoItem


def print_result(operation: str, value: object, request_charge: float) -> None:
    print(f"{operation}: {value}")
    print(f"  Request charge: {request_charge:.2f} RU")


def run_sync_demo(settings: CosmosSettings) -> None:
    print("\n=== Synchronous repository ===")
    with create_sync_repository(settings) as resources:
        repository = resources.repository

        created = repository.create(
            TodoItem.new(
                title="Try the synchronous Cosmos repository",
                description="Run create, read, update, query, and delete.",
                category="demo-sync",
            )
        )
        print_result("Created", created, repository.last_request_charge)

        loaded = repository.read(created.id, created.category)
        print_result("Read", loaded, repository.last_request_charge)

        updated = repository.update(
            replace(loaded, title="Synchronous CRUD complete", completed=True)
        )
        print_result("Updated", updated, repository.last_request_charge)

        print("Query results:")
        for page in repository.query_by_category(updated.category, page_size=1):
            print(
                f"  Page {page.number}: {len(page.items)} item(s), "
                f"{page.request_charge:.2f} RU"
            )
            for item in page.items:
                print(f"    {item}")

        repository.delete(updated.id, updated.category)
        print_result("Deleted", updated.id, repository.last_request_charge)


async def run_async_demo(settings: CosmosSettings) -> None:
    print("\n=== Asynchronous repository ===")
    async with await create_async_repository(settings) as resources:
        repository = resources.repository

        created = await repository.create(
            TodoItem.new(
                title="Try the asynchronous Cosmos repository",
                description="Run create, read, update, query, and delete.",
                category="demo-async",
            )
        )
        print_result("Created", created, repository.last_request_charge)

        loaded = await repository.read(created.id, created.category)
        print_result("Read", loaded, repository.last_request_charge)

        updated = await repository.update(
            replace(loaded, title="Asynchronous CRUD complete", completed=True)
        )
        print_result("Updated", updated, repository.last_request_charge)

        print("Query results:")
        async for page in repository.query_by_category(
            updated.category,
            page_size=1,
        ):
            print(
                f"  Page {page.number}: {len(page.items)} item(s), "
                f"{page.request_charge:.2f} RU"
            )
            for item in page.items:
                print(f"    {item}")

        await repository.delete(updated.id, updated.category)
        print_result("Deleted", updated.id, repository.last_request_charge)


async def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    settings = CosmosSettings.from_environment()
    run_sync_demo(settings)
    await run_async_demo(settings)


if __name__ == "__main__":
    asyncio.run(main())
