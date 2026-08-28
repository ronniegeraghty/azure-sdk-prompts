from __future__ import annotations

import logging
from collections.abc import Callable, Iterator, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos import ContainerProxy
from azure.cosmos.exceptions import CosmosAccessConditionFailedError

from todo_model import TodoItem

logger = logging.getLogger(__name__)

ResponseHook = Callable[[Mapping[str, str], Any], None]


class ConcurrentUpdateError(RuntimeError):
    """Raised when an item changed after it was read."""


class SyncTodoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container

    @staticmethod
    def _response_hook(operation: str) -> ResponseHook:
        def log_request_charge(headers: Mapping[str, str], _response: Any) -> None:
            charge = headers.get("x-ms-request-charge", "unknown")
            logger.info("%s consumed %s RU", operation, charge)

        return log_request_charge

    def create(self, item: TodoItem) -> TodoItem:
        document = self._container.create_item(
            body=item.to_document(),
            response_hook=self._response_hook("create"),
        )
        return TodoItem.from_document(document)

    def read(self, item_id: str, category: str) -> TodoItem:
        document = self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=self._response_hook("read"),
        )
        return TodoItem.from_document(document)

    def update(self, item: TodoItem) -> TodoItem:
        if not item.etag:
            raise ValueError("update requires an item returned by create or read with an etag")

        try:
            document = self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                partition_key=item.category,
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=self._response_hook("update"),
            )
        except CosmosAccessConditionFailedError as error:
            raise ConcurrentUpdateError(
                f"ToDo item {item.id!r} was modified by another process; "
                "read the latest version before retrying"
            ) from error

        return TodoItem.from_document(document)

    def delete(self, item_id: str, category: str) -> None:
        self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=self._response_hook("delete"),
        )

    def query_by_category(
        self, category: str, page_size: int = 100
    ) -> Iterator[list[TodoItem]]:
        if page_size < 1:
            raise ValueError("page_size must be at least 1")

        response_hook = self._response_hook("query")
        results = self._container.query_items(
            query="SELECT * FROM c WHERE c.category = @category",
            parameters=[{"name": "@category", "value": category}],
            partition_key=category,
            max_item_count=page_size,
            response_hook=response_hook,
        )

        for page_number, page in enumerate(results.by_page(), start=1):
            items = [TodoItem.from_document(document) for document in page]
            logger.info("query retrieved page %d with %d item(s)", page_number, len(items))
            yield items
