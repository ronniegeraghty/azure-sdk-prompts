from __future__ import annotations

import logging
from typing import AsyncIterator

from azure.core import MatchConditions
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from models import OperationResult, QueryPage, TodoItem
from repository_common import MissingEtagError, RequestChargeTracker, TodoConflictError


logger = logging.getLogger(__name__)


class AsyncTodoRepository:
    def __init__(self, container: object) -> None:
        self._container = container

    async def create(self, item: TodoItem) -> OperationResult[TodoItem]:
        tracker = RequestChargeTracker()
        document = await self._container.create_item(
            body=item.to_cosmos_document(),
            response_hook=tracker.response_hook,
        )
        return self._item_result("create", document, tracker)

    async def read(self, item_id: str, category: str) -> OperationResult[TodoItem]:
        tracker = RequestChargeTracker()
        document = await self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=tracker.response_hook,
        )
        return self._item_result("read", document, tracker)

    async def update(self, item: TodoItem) -> OperationResult[TodoItem]:
        if not item.etag:
            raise MissingEtagError(
                "An update requires the _etag from a prior create, read, or update operation."
            )

        tracker = RequestChargeTracker()
        try:
            document = await self._container.replace_item(
                item=item.id,
                body=item.to_cosmos_document(),
                partition_key=item.category,
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=tracker.response_hook,
            )
        except CosmosAccessConditionFailedError as error:
            charge = tracker.take_charge()
            logger.info(
                "Cosmos DB rejected stale async update after consuming %.2f RU",
                charge,
            )
            raise TodoConflictError(
                f"ToDo item {item.id!r} was modified by another process; "
                "read the latest version before retrying the update."
            ) from error

        return self._item_result("update", document, tracker)

    async def delete(self, item_id: str, category: str) -> OperationResult[None]:
        tracker = RequestChargeTracker()
        await self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=tracker.response_hook,
        )
        charge = tracker.take_charge()
        logger.info("Cosmos DB async delete consumed %.2f RU", charge)
        return OperationResult(value=None, request_charge=charge)

    async def query_by_category(
        self, category: str, page_size: int = 100
    ) -> AsyncIterator[QueryPage]:
        if page_size <= 0:
            raise ValueError("page_size must be greater than zero")

        tracker = RequestChargeTracker()
        query = "SELECT * FROM c WHERE c.category = @category"
        parameters = [{"name": "@category", "value": category}]
        pages = self._container.query_items(
            query=query,
            parameters=parameters,
            partition_key=category,
            max_item_count=page_size,
            response_hook=tracker.response_hook,
        ).by_page()

        page_number = 0
        async for page in pages:
            page_number += 1
            items = [
                TodoItem.from_cosmos_document(document)
                async for document in page
            ]
            charge = tracker.take_charge()
            logger.info(
                "Cosmos DB async category query page %d returned %d item(s) "
                "and consumed %.2f RU",
                page_number,
                len(items),
                charge,
            )
            yield QueryPage(
                page_number=page_number,
                items=items,
                request_charge=charge,
            )

    @staticmethod
    def _item_result(
        operation: str, document: dict, tracker: RequestChargeTracker
    ) -> OperationResult[TodoItem]:
        charge = tracker.take_charge()
        logger.info("Cosmos DB async %s consumed %.2f RU", operation, charge)
        return OperationResult(
            value=TodoItem.from_cosmos_document(document),
            request_charge=charge,
        )
