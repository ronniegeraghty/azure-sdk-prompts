from __future__ import annotations

import unittest
from unittest.mock import AsyncMock, MagicMock

from azure.core import MatchConditions
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from async_repository import AsyncTodoRepository
from sync_repository import ConcurrentUpdateError, SyncTodoRepository
from todo_model import TodoItem


def document(item_id: str, etag: str = "etag-1") -> dict[str, object]:
    return {
        "id": item_id,
        "title": "Test",
        "description": "Description",
        "completed": False,
        "created_at": "2026-01-01T00:00:00+00:00",
        "category": "work",
        "_etag": etag,
    }


class SyncPageIterator:
    def __init__(self, pages: list[list[dict[str, object]]]) -> None:
        self._pages = pages

    def by_page(self) -> object:
        return iter(self._pages)


class AsyncPage:
    def __init__(self, items: list[dict[str, object]]) -> None:
        self._items = iter(items)

    def __aiter__(self) -> AsyncPage:
        return self

    async def __anext__(self) -> dict[str, object]:
        try:
            return next(self._items)
        except StopIteration as error:
            raise StopAsyncIteration from error


class AsyncPages:
    def __init__(self, pages: list[list[dict[str, object]]]) -> None:
        self._pages = iter(pages)

    def __aiter__(self) -> AsyncPages:
        return self

    async def __anext__(self) -> AsyncPage:
        try:
            return AsyncPage(next(self._pages))
        except StopIteration as error:
            raise StopAsyncIteration from error


class AsyncPageIterator:
    def __init__(self, pages: list[list[dict[str, object]]]) -> None:
        self._pages = pages

    def by_page(self) -> AsyncPages:
        return AsyncPages(self._pages)


class SyncRepositoryTests(unittest.TestCase):
    def test_update_uses_etag_precondition(self) -> None:
        container = MagicMock()
        container.replace_item.return_value = document("1", "etag-2")
        repository = SyncTodoRepository(container)
        item = TodoItem.from_document(document("1"))

        updated = repository.update(item)

        self.assertEqual(updated.etag, "etag-2")
        _, kwargs = container.replace_item.call_args
        self.assertEqual(kwargs["etag"], "etag-1")
        self.assertEqual(kwargs["match_condition"], MatchConditions.IfNotModified)

    def test_update_translates_precondition_failure(self) -> None:
        container = MagicMock()
        container.replace_item.side_effect = CosmosAccessConditionFailedError(
            status_code=412, message="Precondition failed"
        )
        repository = SyncTodoRepository(container)

        with self.assertRaisesRegex(ConcurrentUpdateError, "modified by another process"):
            repository.update(TodoItem.from_document(document("1")))

    def test_query_is_parameterized_and_paged(self) -> None:
        container = MagicMock()
        container.query_items.return_value = SyncPageIterator(
            [[document("1")], [document("2")]]
        )
        repository = SyncTodoRepository(container)

        pages = list(repository.query_by_category("work", page_size=1))

        self.assertEqual([[item.id for item in page] for page in pages], [["1"], ["2"]])
        _, kwargs = container.query_items.call_args
        self.assertEqual(
            kwargs["parameters"], [{"name": "@category", "value": "work"}]
        )
        self.assertEqual(kwargs["partition_key"], "work")
        self.assertEqual(kwargs["max_item_count"], 1)


class AsyncRepositoryTests(unittest.IsolatedAsyncioTestCase):
    async def test_update_uses_etag_precondition(self) -> None:
        container = MagicMock()
        container.replace_item = AsyncMock(return_value=document("1", "etag-2"))
        repository = AsyncTodoRepository(container)

        updated = await repository.update(TodoItem.from_document(document("1")))

        self.assertEqual(updated.etag, "etag-2")
        _, kwargs = container.replace_item.call_args
        self.assertEqual(kwargs["etag"], "etag-1")
        self.assertEqual(kwargs["match_condition"], MatchConditions.IfNotModified)

    async def test_query_iterates_pages_asynchronously(self) -> None:
        container = MagicMock()
        container.query_items.return_value = AsyncPageIterator(
            [[document("1")], [document("2")]]
        )
        repository = AsyncTodoRepository(container)

        pages = [
            page async for page in repository.query_by_category("work", page_size=1)
        ]

        self.assertEqual([[item.id for item in page] for page in pages], [["1"], ["2"]])
        _, kwargs = container.query_items.call_args
        self.assertEqual(
            kwargs["parameters"], [{"name": "@category", "value": "work"}]
        )
        self.assertEqual(kwargs["partition_key"], "work")
