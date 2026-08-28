from __future__ import annotations

import logging
from collections.abc import Iterator, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos import ContainerProxy
from azure.cosmos.exceptions import CosmosHttpResponseError

from todo_model import TodoItem

logger = logging.getLogger(__name__)


class TodoConflictError(RuntimeError):
    """Raised when an update is based on a stale or missing ETag."""


class _RequestChargeLogger:
    def __init__(self, operation: str) -> None:
        self.operation = operation

    def __call__(self, headers: Mapping[str, str], _: Any) -> None:
        charge = float(headers.get("x-ms-request-charge", 0.0))
        logger.info("%s request charge: %.2f RU", self.operation, charge)


class SyncTodoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container

    def create(self, item: TodoItem) -> TodoItem:
        document = self._container.create_item(
            body=item.to_document(),
            response_hook=_RequestChargeLogger("create"),
        )
        return TodoItem.from_document(document)

    def read(self, item_id: str, category: str) -> TodoItem:
        document = self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=_RequestChargeLogger("read"),
        )
        return TodoItem.from_document(document)

    def update(self, item: TodoItem) -> TodoItem:
        if not item.etag:
            raise TodoConflictError(
                "Cannot update a ToDo item without an ETag; read or create it first."
            )

        try:
            document = self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                partition_key=item.category,
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=_RequestChargeLogger("update"),
            )
        except CosmosHttpResponseError as error:
            if error.status_code == 412:
                raise TodoConflictError(
                    f"ToDo item {item.id!r} was changed by another process; "
                    "read the latest version before updating."
                ) from error
            raise

        return TodoItem.from_document(document)

    def delete(self, item_id: str, category: str) -> None:
        self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=_RequestChargeLogger("delete"),
        )

    def query_by_category(
        self, category: str, page_size: int = 100
    ) -> Iterator[list[TodoItem]]:
        if page_size < 1:
            raise ValueError("page_size must be at least 1")

        pages = self._container.query_items(
            query="SELECT * FROM c WHERE c.category = @category",
            parameters=[{"name": "@category", "value": category}],
            partition_key=category,
            max_item_count=page_size,
            response_hook=_RequestChargeLogger("query"),
        ).by_page()

        for page_number, page in enumerate(pages, start=1):
            items = [TodoItem.from_document(document) for document in page]
            logger.info("query page %d retrieved %d item(s)", page_number, len(items))
            yield items
