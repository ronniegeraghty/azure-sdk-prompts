from __future__ import annotations

import logging
from collections.abc import AsyncIterator, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos.aio import ContainerProxy
from azure.cosmos.exceptions import CosmosHttpResponseError

from sync_repository import ConcurrentUpdateError
from todo_model import QueryPage, ToDoItem


class AsyncToDoRepository:
    def __init__(
        self,
        container: ContainerProxy,
        logger: logging.Logger | None = None,
    ) -> None:
        self._container = container
        self._logger = logger or logging.getLogger(__name__)
        self.last_request_charge = 0.0

    @staticmethod
    def _request_charge(headers: Mapping[str, Any]) -> float:
        for key, value in headers.items():
            if key.lower() == "x-ms-request-charge":
                return float(value)
        return 0.0

    def _charge_hook(self, operation: str):
        def capture(headers: Mapping[str, Any], _result: Any = None) -> None:
            self.last_request_charge = self._request_charge(headers)
            self._logger.info(
                "%s consumed %.2f RU", operation, self.last_request_charge
            )

        return capture

    async def create(self, item: ToDoItem) -> ToDoItem:
        document = await self._container.create_item(
            body=item.to_document(),
            response_hook=self._charge_hook("async create"),
        )
        return ToDoItem.from_document(document)

    async def read(self, item_id: str, category: str) -> ToDoItem:
        document = await self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=self._charge_hook("async read"),
        )
        return ToDoItem.from_document(document)

    async def update(self, item: ToDoItem) -> ToDoItem:
        if not item.etag:
            raise ValueError(
                "An ETag is required for update; read the item before modifying it."
            )

        try:
            document = await self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=self._charge_hook("async update"),
            )
        except CosmosHttpResponseError as exc:
            if exc.status_code == 412:
                raise ConcurrentUpdateError(
                    f"ToDo item {item.id!r} was modified by another process; "
                    "read the latest version before retrying."
                ) from exc
            raise
        return ToDoItem.from_document(document)

    async def delete(self, item_id: str, category: str) -> None:
        await self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=self._charge_hook("async delete"),
        )

    async def query_by_category(
        self,
        category: str,
        *,
        page_size: int = 100,
    ) -> AsyncIterator[QueryPage]:
        if page_size <= 0:
            raise ValueError("page_size must be greater than zero")

        query = "SELECT * FROM c WHERE c.category = @category"
        parameters = [{"name": "@category", "value": category}]
        results = self._container.query_items(
            query=query,
            parameters=parameters,
            partition_key=category,
            max_item_count=page_size,
        )
        pages = results.by_page()

        page_number = 0
        async for page in pages:
            page_number += 1
            items = [
                ToDoItem.from_document(document) async for document in page
            ]
            request_charge = self._request_charge(results.get_response_headers())
            self.last_request_charge = request_charge
            self._logger.info(
                "async query page %d retrieved %d item(s) and consumed %.2f RU",
                page_number,
                len(items),
                request_charge,
            )
            yield QueryPage(page_number, items, request_charge)
