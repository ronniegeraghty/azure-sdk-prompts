from __future__ import annotations

import asyncio
import logging
import os

from azure.identity import DefaultAzureCredential
from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
from azure.servicebus import ServiceBusClient, ServiceBusMessage
from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient

from order import Order
from processor import AsyncOrderProcessor, OrderProcessor
from sender import AsyncOrderSender, OrderSender


LOGGER = logging.getLogger(__name__)
QUEUE_NAME = os.getenv("SERVICE_BUS_QUEUE_NAME", "orders")


def _sample_orders(prefix: str) -> list[Order]:
    return [
        Order(f"{prefix}-001", "Ada Lovelace", "Keyboard", 1, 129.99),
        Order(f"{prefix}-002", "Grace Hopper", "Laptop", 1, 1_799.00),
        Order(f"{prefix}-003", "Ada Lovelace", "Monitor", 2, 699.98),
    ]


def _recover_demo_order(message: object, _error: Exception) -> Order | None:
    correlation_id = getattr(message, "correlation_id", None)
    if not correlation_id:
        return None
    return Order(
        order_id=str(correlation_id),
        customer_name="Dead Letter Recovery",
        product="Recovered demo item",
        quantity=1,
        total_price=10.00,
    )


def _send_sync_poison_message(client: ServiceBusClient, prefix: str) -> None:
    order_id = f"{prefix}-poison"
    message = ServiceBusMessage(
        b'{"order_id": "invalid"',
        content_type="application/json",
        correlation_id=order_id,
        message_id=order_id,
        session_id="Dead Letter Recovery",
    )
    with client.get_queue_sender(QUEUE_NAME) as sender:
        sender.send_messages(message)


async def _send_async_poison_message(
    client: AsyncServiceBusClient,
    prefix: str,
) -> None:
    order_id = f"{prefix}-poison"
    message = ServiceBusMessage(
        b'{"order_id": "invalid"',
        content_type="application/json",
        correlation_id=order_id,
        message_id=order_id,
        session_id="Dead Letter Recovery",
    )
    async with client.get_queue_sender(QUEUE_NAME) as sender:
        await sender.send_messages(message)


def run_sync_demo(namespace: str) -> None:
    LOGGER.info("Starting synchronous Service Bus demo")
    credential = DefaultAzureCredential()
    try:
        with ServiceBusClient(namespace, credential) as client:
            sender = OrderSender(client, QUEUE_NAME)
            orders = _sample_orders("sync")
            sender.send_order(orders[0])
            sender.send_orders(orders[1:])
            _send_sync_poison_message(client, "sync")

            processor = OrderProcessor(client, QUEUE_NAME)
            result = processor.process_orders(expected_messages=4)
            LOGGER.info("Sync processing result: %s", result)

            processor.inspect_dead_letters()
            resent = processor.reprocess_dead_letters(_recover_demo_order)
            if resent:
                processor.process_orders(expected_messages=resent)
    finally:
        credential.close()


async def run_async_demo(namespace: str) -> None:
    LOGGER.info("Starting asynchronous Service Bus demo")
    credential = AsyncDefaultAzureCredential()
    try:
        async with AsyncServiceBusClient(namespace, credential) as client:
            sender = AsyncOrderSender(client, QUEUE_NAME)
            orders = _sample_orders("async")
            await sender.send_order(orders[0])
            await sender.send_orders(orders[1:])
            await _send_async_poison_message(client, "async")

            processor = AsyncOrderProcessor(client, QUEUE_NAME)
            result = await processor.process_orders(expected_messages=4)
            LOGGER.info("Async processing result: %s", result)

            await processor.inspect_dead_letters()
            resent = await processor.reprocess_dead_letters(_recover_demo_order)
            if resent:
                await processor.process_orders(expected_messages=resent)
    finally:
        await credential.close()


def main() -> None:
    logging.basicConfig(
        level=logging.INFO,
        format="%(asctime)s %(levelname)s %(name)s: %(message)s",
    )
    namespace = os.environ["SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE"]
    run_sync_demo(namespace)
    asyncio.run(run_async_demo(namespace))


if __name__ == "__main__":
    main()
