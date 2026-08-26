from __future__ import annotations

import asyncio
import logging
import os

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order_processor.model import Order
from order_processor.processor import AsyncOrderProcessor, OrderProcessor
from order_processor.sender import AsyncOrderSender, OrderSender

LOGGER = logging.getLogger(__name__)


def _settings() -> tuple[str, str]:
    namespace = os.environ["SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.environ["SERVICE_BUS_QUEUE_NAME"]
    return namespace, queue_name


def _orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada", "Keyboard", 1, 89.99),
        Order(f"{prefix}-002", "Grace", "Monitor", 2, 699.98),
        Order(f"{prefix}-003", "Ada", "Workstation", 1, 2_499.00),
    ]


def run_sync_demo(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting synchronous demo")
    credential = DefaultAzureCredential()
    try:
        with ServiceBusClient(namespace, credential) as client:
            sender = OrderSender(client, queue_name)
            processor = OrderProcessor(client, queue_name)

            sender.send_orders(_orders("sync"))
            sender.send_order(
                Order("sync-retry", "Linus", "USB hub", 1, 39.00),
                application_properties={"simulate_failure": True},
            )
            processor.process()

            for record in processor.inspect_dead_letters():
                LOGGER.info(
                    "Dead letter %s: %s (%s)",
                    record.order_id,
                    record.reason,
                    record.description,
                )

            processor.reprocess_dead_letters(sender)
            processor.process()
    finally:
        credential.close()


async def run_async_demo(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting asynchronous demo")
    credential = AsyncDefaultAzureCredential()
    try:
        async with AsyncServiceBusClient(namespace, credential) as client:
            sender = AsyncOrderSender(client, queue_name)
            processor = AsyncOrderProcessor(client, queue_name)

            await sender.send_orders(_orders("async"))
            await sender.send_order(
                Order("async-retry", "Margaret", "Mouse", 1, 49.00),
                application_properties={"simulate_failure": True},
            )
            await processor.process()

            for record in await processor.inspect_dead_letters():
                LOGGER.info(
                    "Dead letter %s: %s (%s)",
                    record.order_id,
                    record.reason,
                    record.description,
                )

            await processor.reprocess_dead_letters(sender)
            await processor.process()
    finally:
        await credential.close()


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace, queue_name = _settings()
    run_sync_demo(namespace, queue_name)
    asyncio.run(run_async_demo(namespace, queue_name))


if __name__ == "__main__":
    main()
