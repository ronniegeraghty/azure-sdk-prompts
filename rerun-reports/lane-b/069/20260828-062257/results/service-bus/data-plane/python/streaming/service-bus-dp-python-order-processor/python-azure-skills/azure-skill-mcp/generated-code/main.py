from __future__ import annotations

import asyncio
import logging
import os
from decimal import Decimal

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order_processing.message_factory import customer_session_id
from order_processing.models import Order
from order_processing.processor import AsyncOrderProcessor, OrderProcessor
from order_processing.sender import AsyncOrderSender, OrderSender

LOGGER = logging.getLogger(__name__)


def sample_orders(prefix: str) -> list[Order]:
    return [
        Order(
            order_id=f"{prefix}-001",
            customer_name="Ada Lovelace",
            product="Mechanical Keyboard",
            quantity=1,
            total_price=Decimal("149.99"),
        ),
        Order(
            order_id=f"{prefix}-002",
            customer_name="Ada Lovelace",
            product="4K Monitor",
            quantity=2,
            total_price=Decimal("1199.98"),
        ),
        Order(
            order_id=f"{prefix}-003",
            customer_name="Grace Hopper",
            product="USB-C Dock",
            quantity=1,
            total_price=Decimal("229.00"),
        ),
    ]


def malformed_message(prefix: str) -> ServiceBusMessage:
    message_id = f"{prefix}-invalid"
    return ServiceBusMessage(
        '{"order_id":',
        content_type="application/json",
        correlation_id=message_id,
        message_id=message_id,
        session_id=customer_session_id("Invalid Demo Customer"),
    )


def repaired_demo_order(prefix: str) -> Order:
    return Order(
        order_id=f"{prefix}-invalid-repaired",
        customer_name="Invalid Demo Customer",
        product="Replacement Cable",
        quantity=1,
        total_price=Decimal("19.99"),
    )


def run_sync(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting synchronous demo")
    with DefaultAzureCredential() as credential:
        with ServiceBusClient(namespace, credential) as client:
            sender = OrderSender(client, queue_name)
            processor = OrderProcessor(client, queue_name)
            sender.send_orders(sample_orders("sync"))
            sender.send_raw_message(malformed_message("sync"))
            processor.process_orders(max_messages=4, wait_time=35)
            processor.inspect_dead_letters(
                max_messages=10,
                reprocess=True,
                repair=lambda _: repaired_demo_order("sync"),
            )
            processor.process_orders(max_messages=1)


async def run_async(namespace: str, queue_name: str) -> None:
    LOGGER.info("Starting asynchronous demo")
    async with AsyncDefaultAzureCredential() as credential:
        async with AsyncServiceBusClient(namespace, credential) as client:
            sender = AsyncOrderSender(client, queue_name)
            processor = AsyncOrderProcessor(client, queue_name)
            await sender.send_orders(sample_orders("async"))
            await sender.send_raw_message(malformed_message("async"))
            await processor.process_orders(max_messages=4, wait_time=35)
            await processor.inspect_dead_letters(
                max_messages=10,
                reprocess=True,
                repair=lambda _: repaired_demo_order("async"),
            )
            await processor.process_orders(max_messages=1)


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = os.environ["SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE"]
    queue_name = os.environ["SERVICE_BUS_QUEUE_NAME"]
    run_sync(namespace, queue_name)
    asyncio.run(run_async(namespace, queue_name))


if __name__ == "__main__":
    main()
