from __future__ import annotations

import unittest
from dataclasses import replace
from typing import Any

from azure.core import MatchConditions

from async_repository import AsyncTodoRepository
from sync_repository import SyncTodoRepository, TodoConflictError
from todo_model import TodoItem


def _document(item_id: str = "1", etag: str = "etag-1") -> dict[str, Any]:
    return {
        "id": item_id,
        "title": "Task",
        "description": "Description",
        "completed": False,
        "createdAt": "2026-01-01T00:00:00+00:00",
        "category": "work",
        "_etag": etag,
    }


class _SyncPages:
    def __init__(self, pages: list[list[dict[str, Any]]]) -> None:
        self._pages = pages

    def by_page(self) -> Any:
        return iter(self._pages)


class _SyncContainer:
    def __init__(self) -> None:
        self.replace_kwargs: dict[str, Any] = {}
        self.query_kwargs: dict[str, Any] = {}

    def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self.replace_kwargs = kwargs
        kwargs["response_hook"]({"x-ms-request-charge": "3.5"}, None)
        return {**kwargs["body"], "_etag": "etag-2"}

    def query_items(self, **kwargs: Any) -> _SyncPages:
        self.query_kwargs = kwargs
        kwargs["response_hook"]({"x-ms-request-charge": "2.0"}, None)
        return _SyncPages([[_document("1")], [_document("2")]])


class _AsyncPage:
    def __init__(self, items: list[dict[str, Any]]) -> None:
        self._items = items

    def __aiter__(self) -> Any:
        async def iterator() -> Any:
            for item in self._items:
                yield item

        return iterator()


class _AsyncPages:
    def __init__(self, pages: list[list[dict[str, Any]]]) -> None:
        self._pages = pages

    def __aiter__(self) -> Any:
        async def iterator() -> Any:
            for page in self._pages:
                yield _AsyncPage(page)

        return iterator()

    def by_page(self) -> _AsyncPages:
        return self


class _AsyncContainer:
    def __init__(self) -> None:
        self.replace_kwargs: dict[str, Any] = {}
        self.query_kwargs: dict[str, Any] = {}

    async def replace_item(self, **kwargs: Any) -> dict[str, Any]:
        self.replace_kwargs = kwargs
        kwargs["response_hook"]({"x-ms-request-charge": "3.5"}, None)
        return {**kwargs["body"], "_etag": "etag-2"}

    def query_items(self, **kwargs: Any) -> _AsyncPages:
        self.query_kwargs = kwargs
        kwargs["response_hook"]({"x-ms-request-charge": "2.0"}, None)
        return _AsyncPages([[_document("1")], [_document("2")]])


class SyncRepositoryTests(unittest.TestCase):
    def test_update_uses_optimistic_concurrency(self) -> None:
        container = _SyncContainer()
        repository = SyncTodoRepository(container)  # type: ignore[arg-type]
        item = TodoItem.from_document(_document())

        updated = repository.update(replace(item, completed=True))

        self.assertEqual(updated.etag, "etag-2")
        self.assertEqual(container.replace_kwargs["etag"], "etag-1")
        self.assertIs(
            container.replace_kwargs["match_condition"],
            MatchConditions.IfNotModified,
        )

    def test_update_requires_etag(self) -> None:
        repository = SyncTodoRepository(_SyncContainer())  # type: ignore[arg-type]
        with self.assertRaisesRegex(TodoConflictError, "without an ETag"):
            repository.update(
                TodoItem("1", "Task", "Description", False, "work")
            )

    def test_query_is_parameterized_and_paged(self) -> None:
        container = _SyncContainer()
        repository = SyncTodoRepository(container)  # type: ignore[arg-type]

        pages = list(repository.query_by_category("work", page_size=1))

        self.assertEqual([len(page) for page in pages], [1, 1])
        self.assertEqual(
            container.query_kwargs["parameters"],
            [{"name": "@category", "value": "work"}],
        )
        self.assertEqual(container.query_kwargs["max_item_count"], 1)


class AsyncRepositoryTests(unittest.IsolatedAsyncioTestCase):
    async def test_update_uses_optimistic_concurrency(self) -> None:
        container = _AsyncContainer()
        repository = AsyncTodoRepository(container)  # type: ignore[arg-type]
        item = TodoItem.from_document(_document())

        updated = await repository.update(replace(item, completed=True))

        self.assertEqual(updated.etag, "etag-2")
        self.assertEqual(container.replace_kwargs["etag"], "etag-1")
        self.assertIs(
            container.replace_kwargs["match_condition"],
            MatchConditions.IfNotModified,
        )

    async def test_query_is_parameterized_and_paged(self) -> None:
        container = _AsyncContainer()
        repository = AsyncTodoRepository(container)  # type: ignore[arg-type]

        pages = [
            page async for page in repository.query_by_category("work", page_size=1)
        ]

        self.assertEqual([len(page) for page in pages], [1, 1])
        self.assertEqual(
            container.query_kwargs["parameters"],
            [{"name": "@category", "value": "work"}],
        )


if __name__ == "__main__":
    unittest.main()
