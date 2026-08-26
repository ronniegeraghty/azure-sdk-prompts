"""Run synchronous and asynchronous Azure Service Bus order processing demos."""

from __future__ import annotations

import asyncio
import logging
import os
import time

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order_model import Order
from processor import AsyncOrderProcessor, OrderProcessor
from sender import AsyncOrderSender, FRAUD_REVIEW_DELAY, OrderSender

LOGGER = logging.getLogger(__name__)


def _sample_orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada Lovelace", "Mechanical keyboard", 1, 149.99),
        Order(f"{prefix}-002", "Ada Lovelace", "USB-C cable", 2, 39.98),
        Order(f"{prefix}-003", "Grace Hopper", "Developer workstation", 1, 2499.00),
    ]


def run_sync_demo(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting synchronous demo")
    orders = _sample_orders("sync")
    with DefaultAzureCredential() as credential:
        with ServiceBusClient(namespace, credential) as client:
            sender = OrderSender(client, queue_name)
            processor = OrderProcessor(client, queue_name)

            sender.send_order(orders[0])
            sender.send_orders(orders[1:])
            sender.send_invalid_message_for_demo("Sync Invalid Customer")

            processor.process_available()
            processor.process_dead_letters(reprocess=True)
            processor.process_available()
            processor.process_dead_letters()

            LOGGER.info(
                "Waiting %d seconds for the fraud-review scheduled order",
                int(FRAUD_REVIEW_DELAY.total_seconds()),
            )
            time.sleep(FRAUD_REVIEW_DELAY.total_seconds() + 1)
            processor.process_available()
    LOGGER.info("Synchronous demo complete")


async def run_async_demo(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting asynchronous demo")
    orders = _sample_orders("async")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncServiceBusClient(namespace, credential) as client:
            sender = AsyncOrderSender(client, queue_name)
            processor = AsyncOrderProcessor(client, queue_name)

            await sender.send_order(orders[0])
            await sender.send_orders(orders[1:])
            await sender.send_invalid_message_for_demo("Async Invalid Customer")

            await processor.process_available()
            await processor.process_dead_letters(reprocess=True)
            await processor.process_available()
            await processor.process_dead_letters()

            LOGGER.info(
                "Waiting %d seconds for the fraud-review scheduled order",
                int(FRAUD_REVIEW_DELAY.total_seconds()),
            )
            await asyncio.sleep(FRAUD_REVIEW_DELAY.total_seconds() + 1)
            await processor.process_available()
    LOGGER.info("Asynchronous demo complete")


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = os.environ["SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.environ["SERVICEBUS_QUEUE_NAME"]
    run_sync_demo(namespace, queue_name)
    asyncio.run(run_async_demo(namespace, queue_name))


if __name__ == "__main__":
    main()
