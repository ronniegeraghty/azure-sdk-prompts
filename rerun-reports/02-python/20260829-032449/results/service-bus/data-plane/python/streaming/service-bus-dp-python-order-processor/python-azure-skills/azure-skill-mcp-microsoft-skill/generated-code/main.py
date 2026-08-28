from __future__ import annotations

import asyncio
import logging
import os
from uuid import uuid4

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from async_processor import AsyncOrderProcessor
from async_sender import AsyncOrderSender
from order import Order
from sync_processor import SyncOrderProcessor
from sync_sender import SyncOrderSender

HIGH_PRIORITY_THRESHOLD = 1_000.0
SCHEDULE_WAIT_SECONDS = 35
SYNC_POISON_SESSION = "Sync Poison Demo"
ASYNC_POISON_SESSION = "Async Poison Demo"

logger = logging.getLogger(__name__)


def demo_orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada Lovelace", "Mechanical Keyboard", 1, 149.99),
        Order(f"{prefix}-002", "Grace Hopper", "GPU Workstation", 1, 4_999.00),
        Order(f"{prefix}-003", "Ada Lovelace", "USB-C Dock", 2, 259.98),
    ]


def replacement_order(prefix: str, customer_name: str) -> Order:
    return Order(
        order_id=f"{prefix}-recovered",
        customer_name=customer_name,
        product="Recovered Demo Item",
        quantity=1,
        total_price=25.00,
    )


def send_sync_poison_message(client: ServiceBusClient, queue_name: str) -> None:
    message = ServiceBusMessage(
        '{"order_id": "broken"',
        content_type="application/json",
        correlation_id="sync-invalid",
        message_id=str(uuid4()),
        session_id=SYNC_POISON_SESSION,
    )
    with client.get_queue_sender(queue_name=queue_name) as sender:
        sender.send_messages(message)


async def send_async_poison_message(
    client: AsyncServiceBusClient,
    queue_name: str,
) -> None:
    message = ServiceBusMessage(
        '{"order_id": "broken"',
        content_type="application/json",
        correlation_id="async-invalid",
        message_id=str(uuid4()),
        session_id=ASYNC_POISON_SESSION,
    )
    async with client.get_queue_sender(queue_name=queue_name) as sender:
        await sender.send_messages(message)


def run_sync_demo(namespace: str, queue_name: str) -> None:
    logger.info("Starting synchronous Service Bus demo")
    with DefaultAzureCredential() as credential, ServiceBusClient(
        fully_qualified_namespace=namespace,
        credential=credential,
    ) as client:
        sender = SyncOrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)
        processor = SyncOrderProcessor(client, queue_name)
        orders = demo_orders("sync")

        sender.send_order(orders[0])
        sender.send_orders(orders[1:])
        send_sync_poison_message(client, queue_name)

        processor.process_customers(
            [order.customer_name for order in orders],
            max_wait_time=SCHEDULE_WAIT_SECONDS,
        )
        processor.process_customer(SYNC_POISON_SESSION)
        processor.inspect_dead_letters(
            SYNC_POISON_SESSION,
            replacement_factory=lambda _: replacement_order(
                "sync",
                SYNC_POISON_SESSION,
            ),
        )
        processor.process_customer(SYNC_POISON_SESSION)
    logger.info("Synchronous demo complete")


async def run_async_demo(namespace: str, queue_name: str) -> None:
    logger.info("Starting asynchronous Service Bus demo")
    async with AsyncDefaultAzureCredential() as credential, AsyncServiceBusClient(
        fully_qualified_namespace=namespace,
        credential=credential,
    ) as client:
        sender = AsyncOrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)
        processor = AsyncOrderProcessor(client, queue_name)
        orders = demo_orders("async")

        await sender.send_order(orders[0])
        await sender.send_orders(orders[1:])
        await send_async_poison_message(client, queue_name)

        await processor.process_customers(
            [order.customer_name for order in orders],
            max_wait_time=SCHEDULE_WAIT_SECONDS,
        )
        await processor.process_customer(ASYNC_POISON_SESSION)
        await processor.inspect_dead_letters(
            ASYNC_POISON_SESSION,
            replacement_factory=lambda _: replacement_order(
                "async",
                ASYNC_POISON_SESSION,
            ),
        )
        await processor.process_customer(ASYNC_POISON_SESSION)
    logger.info("Asynchronous demo complete")


def get_required_environment_variable(name: str) -> str:
    value = os.getenv(name)
    if not value:
        raise RuntimeError(f"Set the required environment variable {name}")
    return value


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = get_required_environment_variable(
        "SERVICEBUS_FULLY_QUALIFIED_NAMESPACE"
    )
    queue_name = get_required_environment_variable("SERVICEBUS_QUEUE_NAME")

    run_sync_demo(namespace, queue_name)
    asyncio.run(run_async_demo(namespace, queue_name))


if __name__ == "__main__":
    main()
