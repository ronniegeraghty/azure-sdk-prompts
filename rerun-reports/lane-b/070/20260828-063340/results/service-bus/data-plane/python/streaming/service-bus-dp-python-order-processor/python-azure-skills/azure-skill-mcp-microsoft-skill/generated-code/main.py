"""Run synchronous and asynchronous Service Bus order processing demos."""

from __future__ import annotations

import asyncio
import logging
import os
from uuid import uuid4

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order_processing.messages import customer_session_id
from order_processing.model import Order
from order_processing.processor import AsyncOrderProcessor, OrderProcessor
from order_processing.sender import AsyncOrderSender, OrderSender

HIGH_PRIORITY_THRESHOLD = 1_000.0


def demo_orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada Lovelace", "Keyboard", 1, 129.99),
        Order(f"{prefix}-002", "Grace Hopper", "Monitor", 2, 699.98),
        Order(f"{prefix}-003", "Ada Lovelace", "Mouse", 1, 79.99),
    ]


def repaired_order(prefix: str) -> Order:
    return Order(f"{prefix}-repaired", "Poison Demo", "Recovered item", 1, 10.0)


def run_sync_demo(namespace: str, queue_name: str) -> None:
    credential = DefaultAzureCredential()
    try:
        with ServiceBusClient(namespace, credential) as client:
            sender = OrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)
            processor = OrderProcessor(client, queue_name)
            sender.send_orders(demo_orders("sync"))

            with client.get_queue_sender(queue_name=queue_name) as raw_sender:
                raw_sender.send_messages(
                    ServiceBusMessage(
                        "{not-valid-json",
                        message_id=f"sync-poison-{uuid4()}",
                        correlation_id="sync-poison",
                        session_id=customer_session_id("Poison Demo"),
                        content_type="application/json",
                    )
                )

            processor.process(max_messages=4)
            processor.inspect_dead_letters(max_messages=1)
            processor.reprocess_dead_letters(
                lambda _body, _message: repaired_order("sync"), max_messages=1
            )
            processor.process(max_messages=1)
    finally:
        credential.close()


async def run_async_demo(namespace: str, queue_name: str) -> None:
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncServiceBusClient(namespace, credential) as client:
            sender = AsyncOrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)
            processor = AsyncOrderProcessor(client, queue_name)
            await sender.send_orders(demo_orders("async"))

            async with client.get_queue_sender(queue_name=queue_name) as raw_sender:
                await raw_sender.send_messages(
                    ServiceBusMessage(
                        "{not-valid-json",
                        message_id=f"async-poison-{uuid4()}",
                        correlation_id="async-poison",
                        session_id=customer_session_id("Poison Demo"),
                        content_type="application/json",
                    )
                )

            await processor.process(max_messages=4)
            await processor.inspect_dead_letters(max_messages=1)
            await processor.reprocess_dead_letters(
                lambda _body, _message: repaired_order("async"), max_messages=1
            )
            await processor.process(max_messages=1)


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = os.environ.get("SERVICEBUS_FULLY_QUALIFIED_NAMESPACE")
    queue_name = os.environ.get("SERVICEBUS_QUEUE_NAME")
    if not namespace or not queue_name:
        raise RuntimeError(
            "Set SERVICEBUS_FULLY_QUALIFIED_NAMESPACE and SERVICEBUS_QUEUE_NAME"
        )

    run_sync_demo(namespace, queue_name)
    asyncio.run(run_async_demo(namespace, queue_name))


if __name__ == "__main__":
    main()
