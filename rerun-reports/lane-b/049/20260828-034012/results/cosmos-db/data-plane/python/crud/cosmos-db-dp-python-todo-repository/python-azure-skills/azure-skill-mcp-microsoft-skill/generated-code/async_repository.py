from __future__ import annotations

import logging
from collections.abc import AsyncIterator

from azure.core import MatchConditions
from azure.cosmos.aio import ContainerProxy
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from repository_common import RequestChargeTracker
from todo_model import OperationResult, QueryPage, TodoConflictError, TodoItem


logger = logging.getLogger(__name__)


class AsyncTodoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container

    async def create(self, item: TodoItem) -> OperationResult[TodoItem]:
        tracker = RequestChargeTracker(logger, "async_create")
        document = await self._container.create_item(
            body=item.to_document(),
            response_hook=tracker.response_hook,
        )
        return OperationResult(TodoItem.from_document(document), tracker.last_charge)

    async def read(
        self, item_id: str, category: str
    ) -> OperationResult[TodoItem]:
        tracker = RequestChargeTracker(logger, "async_read")
        document = await self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=tracker.response_hook,
        )
        return OperationResult(TodoItem.from_document(document), tracker.last_charge)

    async def update(self, item: TodoItem) -> OperationResult[TodoItem]:
        if not item.etag:
            raise ValueError(
                "An ETag is required for update; read the item before modifying it."
            )

        tracker = RequestChargeTracker(logger, "async_update")
        try:
            document = await self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=tracker.response_hook,
            )
        except CosmosAccessConditionFailedError as error:
            raise TodoConflictError(
                f"Todo item {item.id!r} was modified after it was read; "
                "reload it before retrying the update."
            ) from error

        return OperationResult(TodoItem.from_document(document), tracker.last_charge)

    async def delete(
        self, item_id: str, category: str
    ) -> OperationResult[None]:
        tracker = RequestChargeTracker(logger, "async_delete")
        await self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=tracker.response_hook,
        )
        return OperationResult(None, tracker.last_charge)

    async def query_by_category(
        self, category: str, page_size: int = 100
    ) -> AsyncIterator[QueryPage]:
        if page_size < 1:
            raise ValueError("page_size must be at least 1")

        tracker = RequestChargeTracker(logger, "async_query_by_category")
        results = self._container.query_items(
            query="SELECT * FROM c WHERE c.category = @category",
            parameters=[{"name": "@category", "value": category}],
            partition_key=category,
            max_item_count=page_size,
            response_hook=tracker.response_hook,
        )

        page_number = 0
        async for page in results.by_page():
            page_number += 1
            items = [
                TodoItem.from_document(document) async for document in page
            ]
            logger.info(
                "Retrieved async query page=%d items=%d request_charge=%.2f RU",
                page_number,
                len(items),
                tracker.last_charge,
            )
            yield QueryPage(page_number, items, tracker.last_charge)
