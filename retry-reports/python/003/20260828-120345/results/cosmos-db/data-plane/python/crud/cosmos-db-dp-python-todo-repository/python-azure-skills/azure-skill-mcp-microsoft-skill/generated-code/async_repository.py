from __future__ import annotations

import logging
from collections.abc import AsyncIterator, Callable, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos.aio import ContainerProxy
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from sync_repository import ConcurrentUpdateError
from todo_model import TodoItem

logger = logging.getLogger(__name__)

ResponseHook = Callable[[Mapping[str, str], Any], None]


class AsyncTodoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container

    @staticmethod
    def _response_hook(operation: str) -> ResponseHook:
        def log_request_charge(headers: Mapping[str, str], _response: Any) -> None:
            charge = headers.get("x-ms-request-charge", "unknown")
            logger.info("%s consumed %s RU", operation, charge)

        return log_request_charge

    async def create(self, item: TodoItem) -> TodoItem:
        document = await self._container.create_item(
            body=item.to_document(),
            response_hook=self._response_hook("async create"),
        )
        return TodoItem.from_document(document)

    async def read(self, item_id: str, category: str) -> TodoItem:
        document = await self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=self._response_hook("async read"),
        )
        return TodoItem.from_document(document)

    async def update(self, item: TodoItem) -> TodoItem:
        if not item.etag:
            raise ValueError("update requires an item returned by create or read with an etag")

        try:
            document = await self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                partition_key=item.category,
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=self._response_hook("async update"),
            )
        except CosmosAccessConditionFailedError as error:
            raise ConcurrentUpdateError(
                f"ToDo item {item.id!r} was modified by another process; "
                "read the latest version before retrying"
            ) from error

        return TodoItem.from_document(document)

    async def delete(self, item_id: str, category: str) -> None:
        await self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=self._response_hook("async delete"),
        )

    async def query_by_category(
        self, category: str, page_size: int = 100
    ) -> AsyncIterator[list[TodoItem]]:
        if page_size < 1:
            raise ValueError("page_size must be at least 1")

        response_hook = self._response_hook("async query")
        results = self._container.query_items(
            query="SELECT * FROM c WHERE c.category = @category",
            parameters=[{"name": "@category", "value": category}],
            partition_key=category,
            max_item_count=page_size,
            response_hook=response_hook,
        )

        page_number = 0
        async for page in results.by_page():
            page_number += 1
            items = [TodoItem.from_document(document) async for document in page]
            logger.info(
                "async query retrieved page %d with %d item(s)",
                page_number,
                len(items),
            )
            yield items
