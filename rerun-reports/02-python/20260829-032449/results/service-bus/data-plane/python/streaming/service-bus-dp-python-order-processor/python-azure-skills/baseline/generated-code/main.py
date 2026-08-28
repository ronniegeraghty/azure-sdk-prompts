from __future__ import annotations

import asyncio
import logging
import os
from collections.abc import Iterable

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order_system.async_processor import AsyncOrderProcessor
from order_system.async_sender import AsyncOrderSender
from order_system.model import Order
from order_system.processor import OrderProcessor
from order_system.sender import OrderSender

logger = logging.getLogger(__name__)


def sample_orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada Lovelace", "Keyboard", 1, 129.99),
        Order(f"{prefix}-002", "Ada Lovelace", "Monitor", 2, 699.98),
        Order(f"{prefix}-003", "Grace Hopper", "Workstation", 1, 2_499.00),
    ]


def send_invalid_sync(
    namespace: str,
    queue_name: str,
    credential: DefaultAzureCredential,
    prefix: str,
) -> None:
    with ServiceBusClient(namespace, credential) as client:
        with client.get_queue_sender(queue_name) as sender:
            sender.send_messages(
                ServiceBusMessage(
                    "{not valid JSON",
                    content_type="application/json",
                    correlation_id=f"{prefix}-invalid",
                    session_id="Invalid order demo",
                )
            )


async def send_invalid_async(
    namespace: str,
    queue_name: str,
    credential: AsyncDefaultAzureCredential,
    prefix: str,
) -> None:
    async with AsyncServiceBusClient(namespace, credential) as client:
        async with client.get_queue_sender(queue_name) as sender:
            await sender.send_messages(
                ServiceBusMessage(
                    "{not valid JSON",
                    content_type="application/json",
                    correlation_id=f"{prefix}-invalid",
                    session_id="Invalid order demo",
                )
            )


def run_sync_demo(
    namespace: str,
    queue_name: str,
    orders: Iterable[Order],
) -> None:
    logger.info("Starting synchronous demo")
    credential = DefaultAzureCredential()
    try:
        with OrderSender(namespace, queue_name, credential) as sender:
            sender.send_orders(orders)
        send_invalid_sync(namespace, queue_name, credential, "sync")
        with OrderProcessor(namespace, queue_name, credential) as processor:
            processor.process_available_orders(max_wait_time=35)
            records = processor.read_dead_letters(max_wait_time=5)
            logger.info("Synchronous demo found %d DLQ message(s)", len(records))
    finally:
        credential.close()


async def run_async_demo(
    namespace: str,
    queue_name: str,
    orders: Iterable[Order],
) -> None:
    logger.info("Starting asynchronous demo")
    credential = AsyncDefaultAzureCredential()
    try:
        async with AsyncOrderSender(
            namespace, queue_name, credential
        ) as sender:
            await sender.send_orders(orders)
        await send_invalid_async(namespace, queue_name, credential, "async")
        async with AsyncOrderProcessor(
            namespace, queue_name, credential
        ) as processor:
            await processor.process_available_orders(max_wait_time=35)
            records = await processor.read_dead_letters(max_wait_time=5)
            logger.info("Asynchronous demo found %d DLQ message(s)", len(records))
    finally:
        await credential.close()


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = os.environ["SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.getenv("SERVICE_BUS_QUEUE_NAME", "orders")

    run_sync_demo(namespace, queue_name, sample_orders("sync"))
    asyncio.run(
        run_async_demo(namespace, queue_name, sample_orders("async"))
    )


if __name__ == "__main__":
    main()

