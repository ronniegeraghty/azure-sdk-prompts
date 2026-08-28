from __future__ import annotations

import unittest
from collections.abc import AsyncIterator, Iterator
from dataclasses import replace
from typing import Any

from azure.cosmos import exceptions

from async_repository import AsyncToDoRepository
from repository_common import ConcurrencyConflictError
from sync_repository import SyncToDoRepository
from todo_model import ToDoItem


def stored_document(item: ToDoItem, etag: str = '"version-1"') -> dict[str, Any]:
    return {**item.to_document(), "_etag": etag}


class SyncPageResult:
    def __init__(self, pages: list[list[dict[str, Any]]], hook: Any) -> None:
        self._pages = pages
        self._hook = hook

    def by_page(self) -> Iterator[Iterator[dict[str, Any]]]:
        for page in self._pages:
            self._hook({"x-ms-request-charge": "2.5"}, {})
            yield iter(page)


class SyncContainer:
    def __init__(self, document: dict[str, Any]) -> None:
        self.document = document
        self.last_replace: dict[str, Any] = {}

    @staticmethod
    def _charge(kwargs: dict[str, Any], charge: str = "1.25") -> None:
        kwargs["response_hook"]({"x-ms-request-charge": charge}, {})

    def create_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        return self.document

    def read_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        return self.document

    def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self.last_replace = kwargs
        self._charge(kwargs)
        return {**kwargs["body"], "_etag": '"version-2"'}

    def delete_item(self, **kwargs: Any) -> None:
        self._charge(kwargs)

    def query_items(self, **kwargs: Any) -> SyncPageResult:
        return SyncPageResult([[self.document], [self.document]], kwargs["response_hook"])


class ConflictSyncContainer(SyncContainer):
    def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        raise exceptions.CosmosHttpResponseError(status_code=412)


class AsyncPage:
    def __init__(self, documents: list[dict[str, Any]]) -> None:
        self._documents = documents

    def __aiter__(self) -> AsyncIterator[dict[str, Any]]:
        async def iterate() -> AsyncIterator[dict[str, Any]]:
            for document in self._documents:
                yield document

        return iterate()


class AsyncPageResult:
    def __init__(self, pages: list[list[dict[str, Any]]], hook: Any) -> None:
        self._pages = pages
        self._hook = hook

    def by_page(self) -> AsyncIterator[AsyncPage]:
        async def iterate() -> AsyncIterator[AsyncPage]:
            for page in self._pages:
                self._hook({"x-ms-request-charge": "2.5"}, {})
                yield AsyncPage(page)

        return iterate()


class AsyncContainer:
    def __init__(self, document: dict[str, Any]) -> None:
        self.document = document
        self.last_replace: dict[str, Any] = {}

    @staticmethod
    def _charge(kwargs: dict[str, Any], charge: str = "1.25") -> None:
        kwargs["response_hook"]({"x-ms-request-charge": charge}, {})

    async def create_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        return self.document

    async def read_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        return self.document

    async def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self.last_replace = kwargs
        self._charge(kwargs)
        return {**kwargs["body"], "_etag": '"version-2"'}

    async def delete_item(self, **kwargs: Any) -> None:
        self._charge(kwargs)

    def query_items(self, **kwargs: Any) -> AsyncPageResult:
        return AsyncPageResult([[self.document], [self.document]], kwargs["response_hook"])


class ConflictAsyncContainer(AsyncContainer):
    async def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self._charge(kwargs)
        raise exceptions.CosmosHttpResponseError(status_code=412)


class SyncRepositoryTests(unittest.TestCase):
    def setUp(self) -> None:
        self.item = ToDoItem.new("Test", "Description", "tests")
        self.container = SyncContainer(stored_document(self.item))
        self.repository = SyncToDoRepository(self.container)  # type: ignore[arg-type]

    def test_crud_uses_etag_and_reports_request_charge(self) -> None:
        created = self.repository.create(self.item)
        read = self.repository.read(created.value.id, created.value.category)
        updated = self.repository.update(replace(read.value, completed=True))
        deleted = self.repository.delete(updated.value.id, updated.value.category)

        self.assertEqual(created.request_charge, 1.25)
        self.assertEqual(updated.value.etag, '"version-2"')
        self.assertEqual(self.container.last_replace["etag"], '"version-1"')
        self.assertEqual(deleted.request_charge, 1.25)

    def test_query_yields_bounded_pages_and_uses_parameters(self) -> None:
        pages = list(self.repository.query_by_category("tests", page_size=1))
        self.assertEqual([page.number for page in pages], [1, 2])
        self.assertEqual([page.request_charge for page in pages], [2.5, 2.5])
        self.assertTrue(all(len(page.items) == 1 for page in pages))

    def test_stale_update_has_clear_conflict(self) -> None:
        repository = SyncToDoRepository(  # type: ignore[arg-type]
            ConflictSyncContainer(stored_document(self.item))
        )
        with self.assertRaisesRegex(ConcurrencyConflictError, "was modified"):
            repository.update(ToDoItem.from_document(stored_document(self.item)))


class AsyncRepositoryTests(unittest.IsolatedAsyncioTestCase):
    async def asyncSetUp(self) -> None:
        self.item = ToDoItem.new("Test", "Description", "tests")
        self.container = AsyncContainer(stored_document(self.item))
        self.repository = AsyncToDoRepository(self.container)  # type: ignore[arg-type]

    async def test_crud_uses_etag_and_reports_request_charge(self) -> None:
        created = await self.repository.create(self.item)
        read = await self.repository.read(
            created.value.id, created.value.category
        )
        updated = await self.repository.update(
            replace(read.value, completed=True)
        )
        deleted = await self.repository.delete(
            updated.value.id, updated.value.category
        )

        self.assertEqual(created.request_charge, 1.25)
        self.assertEqual(updated.value.etag, '"version-2"')
        self.assertEqual(self.container.last_replace["etag"], '"version-1"')
        self.assertEqual(deleted.request_charge, 1.25)

    async def test_query_yields_pages_asynchronously(self) -> None:
        pages = [
            page
            async for page in self.repository.query_by_category(
                "tests", page_size=1
            )
        ]
        self.assertEqual([page.number for page in pages], [1, 2])
        self.assertEqual([page.request_charge for page in pages], [2.5, 2.5])

    async def test_stale_update_has_clear_conflict(self) -> None:
        repository = AsyncToDoRepository(  # type: ignore[arg-type]
            ConflictAsyncContainer(stored_document(self.item))
        )
        with self.assertRaisesRegex(ConcurrencyConflictError, "was modified"):
            await repository.update(
                ToDoItem.from_document(stored_document(self.item))
            )


if __name__ == "__main__":
    unittest.main()
