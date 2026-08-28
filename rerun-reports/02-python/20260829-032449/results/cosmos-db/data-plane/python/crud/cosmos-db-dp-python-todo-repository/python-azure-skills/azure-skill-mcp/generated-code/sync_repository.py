from __future__ import annotations

import logging
from collections.abc import Iterator

from azure.core import MatchConditions
from azure.cosmos import exceptions
from azure.cosmos.container import ContainerProxy

from repository_common import ConcurrencyConflictError, RequestChargeTracker
from todo_model import OperationResult, QueryPage, ToDoItem


LOGGER = logging.getLogger(__name__)


class SyncToDoRepository:
    def __init__(self, container: ContainerProxy) -> None:
        self._container = container

    def create(self, item: ToDoItem) -> OperationResult[ToDoItem]:
        tracker = RequestChargeTracker()
        try:
            document = self._container.create_item(
                body=item.to_document(), response_hook=tracker
            )
        finally:
            self._log_charge("create", tracker.total)
        return OperationResult(ToDoItem.from_document(document), tracker.total)

    def read(self, item_id: str, category: str) -> OperationResult[ToDoItem]:
        tracker = RequestChargeTracker()
        try:
            document = self._container.read_item(
                item=item_id,
                partition_key=category,
                response_hook=tracker,
            )
        finally:
            self._log_charge("read", tracker.total)
        return OperationResult(ToDoItem.from_document(document), tracker.total)

    def update(self, item: ToDoItem) -> OperationResult[ToDoItem]:
        if item.etag is None:
            raise ValueError(
                "Cannot update an item without an ETag; read it from the repository first."
            )

        tracker = RequestChargeTracker()
        try:
            document = self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                partition_key=item.category,
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=tracker,
            )
        except exceptions.CosmosHttpResponseError as exc:
            if exc.status_code == 412:
                raise ConcurrencyConflictError(
                    f"ToDo item {item.id!r} in category {item.category!r} "
                    "was modified after it was read."
                ) from exc
            raise
        finally:
            self._log_charge("update", tracker.total)
        return OperationResult(ToDoItem.from_document(document), tracker.total)

    def delete(self, item_id: str, category: str) -> OperationResult[None]:
        tracker = RequestChargeTracker()
        try:
            self._container.delete_item(
                item=item_id,
                partition_key=category,
                response_hook=tracker,
            )
        finally:
            self._log_charge("delete", tracker.total)
        return OperationResult(None, tracker.total)

    def query_by_category(
        self, category: str, page_size: int = 100
    ) -> Iterator[QueryPage]:
        if page_size < 1:
            raise ValueError("page_size must be at least 1")

        tracker = RequestChargeTracker()
        result = self._container.query_items(
            query="SELECT * FROM todo WHERE todo.category = @category",
            parameters=[{"name": "@category", "value": category}],
            partition_key=category,
            max_item_count=page_size,
            response_hook=tracker,
        )

        for page_number, raw_page in enumerate(result.by_page(), start=1):
            items = tuple(
                ToDoItem.from_document(document) for document in raw_page
            )
            charge = tracker.take()
            LOGGER.info(
                "query_by_category page=%d items=%d request_charge=%.2f RU",
                page_number,
                len(items),
                charge,
            )
            yield QueryPage(page_number, items, charge)

    @staticmethod
    def _log_charge(operation: str, request_charge: float) -> None:
        LOGGER.info("%s request_charge=%.2f RU", operation, request_charge)
