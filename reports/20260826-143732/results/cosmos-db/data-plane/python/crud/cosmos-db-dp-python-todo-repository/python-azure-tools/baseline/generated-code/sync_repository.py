from __future__ import annotations

import logging
from collections.abc import Iterator, Mapping
from typing import Any

from azure.core import MatchConditions
from azure.cosmos import ContainerProxy
from azure.cosmos.exceptions import CosmosHttpResponseError

from todo_model import QueryPage, ToDoItem


class ConcurrentUpdateError(RuntimeError):
    """Raised when an item changed after it was read."""


class ToDoRepository:
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

    def create(self, item: ToDoItem) -> ToDoItem:
        document = self._container.create_item(
            body=item.to_document(),
            response_hook=self._charge_hook("create"),
        )
        return ToDoItem.from_document(document)

    def read(self, item_id: str, category: str) -> ToDoItem:
        document = self._container.read_item(
            item=item_id,
            partition_key=category,
            response_hook=self._charge_hook("read"),
        )
        return ToDoItem.from_document(document)

    def update(self, item: ToDoItem) -> ToDoItem:
        if not item.etag:
            raise ValueError(
                "An ETag is required for update; read the item before modifying it."
            )

        try:
            document = self._container.replace_item(
                item=item.id,
                body=item.to_document(),
                etag=item.etag,
                match_condition=MatchConditions.IfNotModified,
                response_hook=self._charge_hook("update"),
            )
        except CosmosHttpResponseError as exc:
            if exc.status_code == 412:
                raise ConcurrentUpdateError(
                    f"ToDo item {item.id!r} was modified by another process; "
                    "read the latest version before retrying."
                ) from exc
            raise
        return ToDoItem.from_document(document)

    def delete(self, item_id: str, category: str) -> None:
        self._container.delete_item(
            item=item_id,
            partition_key=category,
            response_hook=self._charge_hook("delete"),
        )

    def query_by_category(
        self,
        category: str,
        *,
        page_size: int = 100,
    ) -> Iterator[QueryPage]:
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

        for page_number, page in enumerate(pages, start=1):
            items = [ToDoItem.from_document(document) for document in page]
            request_charge = self._request_charge(results.get_response_headers())
            self.last_request_charge = request_charge
            self._logger.info(
                "query page %d retrieved %d item(s) and consumed %.2f RU",
                page_number,
                len(items),
                request_charge,
            )
            yield QueryPage(page_number, items, request_charge)
