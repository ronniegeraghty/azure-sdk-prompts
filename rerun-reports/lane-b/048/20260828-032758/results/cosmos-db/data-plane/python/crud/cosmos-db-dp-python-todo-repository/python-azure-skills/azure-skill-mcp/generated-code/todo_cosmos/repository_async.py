from __future__ import annotations

import logging
from collections.abc import AsyncIterator, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos.aio import ContainerProxy
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from .model import TodoItem, TodoPage
from .repository import TodoConflictError, _request_charge

logger = logging.getLogger(__name__)


class AsyncTodoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container
        self.last_request_charge = 0.0

    def _capture_charge(self, operation: str):
        def response_hook(headers: Mapping[str, str], _response: Any) -> None:
            self.last_request_charge = _request_charge(headers)
            logger.info(
                "async %s consumed %.2f RU",
                operation,
                self.last_request_charge,
            )

        return response_hook

    async def create(self, item: TodoItem) -> TodoItem:
        document = await self._container.create_item(
            body=item.to_document(),
            response_hook=self._capture_charge("create"),
        )
        return TodoItem.from_document(document)

    async def read(self, item_id: str, category: str) -> TodoItem:
        document = await self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=self._capture_charge("read"),
        )
        return TodoItem.from_document(document)

    async def update(self, item: TodoItem) -> TodoItem:
        if not item.etag:
            raise ValueError(
                "An ETag is required for updates; read the item before updating it."
            )

        try:
            document = await self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=self._capture_charge("update"),
            )
        except CosmosAccessConditionFailedError as exc:
            self.last_request_charge = _request_charge(exc.headers)
            logger.info(
                "conflicted async update consumed %.2f RU",
                self.last_request_charge,
            )
            raise TodoConflictError(
                f"ToDo item {item.id!r} was modified by another process; "
                "read the latest version before retrying."
            ) from exc

        return TodoItem.from_document(document)

    async def delete(self, item_id: str, category: str) -> None:
        await self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=self._capture_charge("delete"),
        )

    async def query_by_category(
        self,
        category: str,
        page_size: int = 100,
    ) -> AsyncIterator[TodoPage]:
        if page_size <= 0:
            raise ValueError("page_size must be greater than zero")

        page_charges: list[float] = []

        def response_hook(headers: Mapping[str, str], _response: Any) -> None:
            page_charges.append(_request_charge(headers))

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
            documents = [document async for document in page]
            charge = page_charges[-1] if page_charges else 0.0
            self.last_request_charge = charge
            items = tuple(TodoItem.from_document(document) for document in documents)
            logger.info(
                "async query_by_category page %d retrieved %d item(s), "
                "consuming %.2f RU",
                page_number,
                len(items),
                charge,
            )
            yield TodoPage(page_number, items, charge)
