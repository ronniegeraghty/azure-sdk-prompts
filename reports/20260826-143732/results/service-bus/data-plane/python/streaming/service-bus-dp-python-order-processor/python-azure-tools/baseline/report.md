# Evaluation Report: service-bus-dp-python-order-processor

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 276.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-order-processor` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 16/21 |
| Duration | 276.0s |
| Timestamp | 2026-08-26T08:23:03Z |
| Files Generated | 7 |
| Event Count | 11650 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 188.9s |
| Review | 86.3s |
| **Total** | **276.0s** |

## Configuration

- **name:** python-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=240642 out=9402 |
| Turn Count | 9 |

## Prompt Sent

```
Create a Python project that implements an order processing system using Azure Service Bus.

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- A **model** for an Order with fields for order ID, customer name, product, quantity, total price, and status (pending/processing/completed/failed). It should be serializable to and from JSON. Use a dataclass or dictionary.

- A **sender module** (both sync and async versions) that publishes order messages to a Service Bus queue. It should support sending individual orders and sending a batch of orders efficiently (using `ServiceBusMessageBatch` and checking `add_message()` to respect the maximum batch size so messages aren't rejected). Each message should carry the order ID as a correlation property, and orders above a certain dollar threshold should be sent as high-priority with a scheduled delivery time of 30 seconds in the future (to allow for fraud review before processing).

- A **processor module** (both sync and async versions) that receives and processes orders from the queue. It should handle messages as they arrive, deserialize them, and log the results. If processing fails (e.g., a deserialization error), the message should be sent to the dead-letter queue with a reason string using `dead_letter_message()` rather than being silently abandoned. The processor should also be able to read from the dead-letter queue so failed messages can be inspected and reprocessed. It should guarantee that orders from the same customer are processed in sequence, not interleaved with other customers' orders.

- A **main script** that demos both implementations: connects to the Service Bus namespace (from an environment variable) with `DefaultAzureCredential`, runs the full send/receive/dead-letter cycle using the sync implementation first, then repeats with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace root"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: order_processor/__init__.py\n+\"\"\"Azure Service Bus order processing example.\"\"\"\n+\n+from .model import Order, OrderStatus\n+\n+__all__ = [\"Order\", \"OrderStatus\"]\n*** Add File: order_processor/model.py\n+from __future__ import annotations\n+\n+import json\n+from dataclasses import asdict, dataclass\n+from enum import Enum\n+from typing import Any\n+\n+\n+class OrderStatus(str, Enum):\n+    PENDING = \"pending\"\n+    PROCESSING = \"processing\"\n+    COMPLETED = \"completed\"\n+    FAILED = \"failed\"\n+\n+\n+@dataclass\n+class Order:\n+    order_id: str\n+    customer_name: str\n+    product: str\n+    quantity: int\n+    total_price: float\n+    status: OrderStatus = OrderStatus.PENDING\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.order_id.strip():\n+            raise ValueError(\"order_id must not be empty\")\n+        if not self.customer_name.strip():\n+            raise ValueError(\"customer_name must not be empty\")\n+        if not self.product.strip():\n+            raise ValueError(\"product must not be empty\")\n+        if self.quantity \u003c= 0:\n+            raise ValueError(\"quantity must be greater than zero\")\n+        if self.total_price \u003c 0:\n+            raise ValueError(\"total_price must not be negative\")\n+        if not isinstance(self.status, OrderStatus):\n+            self.status = OrderStatus(self.status)\n+\n+    def to_json(self) -\u003e str:\n+        data = asdict(self)\n+        data[\"status\"] = self.status.value\n+        return json.dumps(data, separators=(\",\", \":\"))\n+\n+    @classmethod\n+    def from_json(cls, payload: str | bytes) -\u003e Order:\n+        data: Any = json.loads(payload)\n+        if not isinstance(data, dict):\n+            raise ValueError(\"order JSON must contain an object\")\n+\n+        required = {\n+            \"order_id\",\n+            \"customer_name\",\n+            \"product\",\n+            \"quantity\",\n+            \"total_price\",\n+            \"status\",\n+        }\n+        missing = required.difference(data)\n+        if missing:\n+            raise ValueError(f\"order JSON is missing fields: {', '.join(sorted(missing))}\")\n+\n+        return cls(\n+            order_id=str(data[\"order_id\"]),\n+            customer_name=str(data[\"customer_name\"]),\n+            product=str(data[\"product\"]),\n+            quantity=int(data[\"quantity\"]),\n+            total_price=float(data[\"total_price\"]),\n+            status=OrderStatus(data[\"status\"]),\n+        )\n*** Add File: order_processor/sender.py\n+from __future__ import annotations\n+\n+from collections import defaultdict\n+from collections.abc import Iterable, Mapping\n+from datetime import datetime, timedelta, timezone\n+from typing import Any\n+\n+from azure.servicebus import ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+from azure.servicebus.exceptions import MessageSizeExceededError\n+from azure.servicebus import ServiceBusClient\n+\n+from .model import Order\n+\n+\n+class _MessageFactory:\n+    def __init__(self, high_priority_threshold: float) -\u003e None:\n+        self.high_priority_threshold = high_priority_threshold\n+\n+    def create(\n+        self,\n+        order: Order,\n+        application_properties: Mapping[str, Any] | None = None,\n+    ) -\u003e ServiceBusMessage:\n+        properties: dict[str, Any] = dict(application_properties or {})\n+        is_high_priority = order.total_price \u003e self.high_priority_threshold\n+        properties[\"priority\"] = \"high\" if is_high_priority else \"normal\"\n+\n+        scheduled_time = (\n+            datetime.now(timezone.utc) + timedelta(seconds=30)\n+            if is_high_priority\n+            else None\n+        )\n+        return ServiceBusMessage(\n+            order.to_json(),\n+            content_type=\"application/json\",\n+            correlation_id=order.order_id,\n+            session_id=order.customer_name,\n+            application_properties=properties,\n+            scheduled_enqueue_time_utc=scheduled_time,\n+        )\n+\n+\n+def _orders_by_customer(orders: Iterable[Order]) -\u003e dict[str, list[Order]]:\n+    grouped: dict[str, list[Order]] = defaultdict(list)\n+    for order in orders:\n+        grouped[order.customer_name].append(order)\n+    return grouped\n+\n+\n+class OrderSender:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self.client = client\n+        self.queue_name = queue_name\n+        self._messages = _MessageFactory(high_priority_threshold)\n+\n+    def send_order(\n+        self,\n+        order: Order,\n+        application_properties: Mapping[str, Any] | None = None,\n+    ) -\u003e None:\n+        with self.client.get_queue_sender(queue_name=self.queue_name) as sender:\n+            sender.send_messages(self._messages.create(order, application_properties))\n+\n+    def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        with self.client.get_queue_sender(queue_name=self.queue_name) as sender:\n+            for customer_orders in _orders_by_customer(orders).values():\n+                batch = sender.create_message_batch()\n+                for order in customer_orders:\n+                    message = self._messages.create(order)\n+                    try:\n+                        batch.add_message(message)\n+                    except MessageSizeExceededError:\n+                        if len(batch) == 0:\n+                            raise\n+                        sender.send_messages(batch)\n+                        batch = sender.create_message_batch()\n+                        batch.add_message(message)\n+                if len(batch):\n+                    sender.send_messages(batch)\n+\n+\n+class AsyncOrderSender:\n+    def __init__(\n+        self,\n+        client: AsyncServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self.client = client\n+        self.queue_name = queue_name\n+        self._messages = _MessageFactory(high_priority_threshold)\n+\n+    async def send_order(\n+        self,\n+        order: Order,\n+        application_properties: Mapping[str, Any] | None = None,\n+    ) -\u003e None:\n+        async with self.client.get_queue_sender(queue_name=self.queue_name) as sender:\n+            await sender.send_messages(\n+                self._messages.create(order, application_properties)\n+            )\n+\n+    async def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        async with self.client.get_queue_sender(queue_name=self.queue_name) as sender:\n+            for customer_orders in _orders_by_customer(orders).values():\n+                batch = await sender.create_message_batch()\n+                for order in customer_orders:\n+                    message = self._messages.create(order)\n+                    try:\n+                        batch.add_message(message)\n+                    except MessageSizeExceededError:\n+                        if len(batch) == 0:\n+                            raise\n+                        await sender.send_messages(batch)\n+                        batch = await sender.create_message_batch()\n+                        batch.add_message(message)\n+                if len(batch):\n+                    await sender.send_messages(batch)\n*** Add File: order_processor/processor.py\n+from __future__ import annotations\n+\n+import logging\n+from dataclasses import dataclass\n+from typing import Any\n+\n+from azure.servicebus import (\n+    NEXT_AVAILABLE_SESSION,\n+    ServiceBusClient,\n+    ServiceBusReceivedMessage,\n+    ServiceBusSubQueue,\n+)\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+from azure.servicebus.exceptions import ServiceBusError\n+\n+from .model import Order, OrderStatus\n+from .sender import AsyncOrderSender, OrderSender\n+\n+LOGGER = logging.getLogger(__name__)\n+\n+\n+@dataclass(frozen=True)\n+class DeadLetterRecord:\n+    order_id: str | None\n+    reason: str | None\n+    description: str | None\n+    body: str\n+\n+\n+def _message_body(message: ServiceBusReceivedMessage) -\u003e str:\n+    return b\"\".join(message.body).decode(\"utf-8\")\n+\n+\n+def _property_is_true(properties: dict[str | bytes, Any] | None, name: str) -\u003e bool:\n+    if not properties:\n+        return False\n+    value = properties.get(name, properties.get(name.encode()))\n+    if isinstance(value, bytes):\n+        value = value.decode(\"utf-8\")\n+    return value is True or value == \"true\"\n+\n+\n+def _process_order(message: ServiceBusReceivedMessage) -\u003e Order:\n+    order = Order.from_json(_message_body(message))\n+    order.status = OrderStatus.PROCESSING\n+    if _property_is_true(message.application_properties, \"simulate_failure\"):\n+        raise RuntimeError(\"simulated processing failure\")\n+    order.status = OrderStatus.COMPLETED\n+    return order\n+\n+\n+def _dead_letter_record(message: ServiceBusReceivedMessage) -\u003e DeadLetterRecord:\n+    correlation_id = message.correlation_id\n+    return DeadLetterRecord(\n+        order_id=str(correlation_id) if correlation_id is not None else None,\n+        reason=message.dead_letter_reason,\n+        description=message.dead_letter_error_description,\n+        body=_message_body(message),\n+    )\n+\n+\n+class OrderProcessor:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        max_wait_time: float = 5,\n+    ) -\u003e None:\n+        self.client = client\n+        self.queue_name = queue_name\n+        self.max_wait_time = max_wait_time\n+\n+    def process(self, max_messages: int | None = None) -\u003e int:\n+        processed = 0\n+        while max_messages is None or processed \u003c max_messages:\n+            receiver = self.client.get_queue_receiver(\n+                queue_name=self.queue_name,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=self.max_wait_time,\n+            )\n+            try:\n+                with receiver:\n+                    while max_messages is None or processed \u003c max_messages:\n+                        messages = receiver.receive_messages(\n+                            max_message_count=(\n+                                min(20, max_messages - processed)\n+                                if max_messages is not None\n+                                else 20\n+                            ),\n+                            max_wait_time=self.max_wait_time,\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            try:\n+                                order = _process_order(message)\n+                            except (TypeError, ValueError, UnicodeError, RuntimeError) as exc:\n+                                LOGGER.error(\n+                                    \"Order %s failed: %s\", message.correlation_id, exc\n+                                )\n+                                receiver.dead_letter_message(\n+                                    message,\n+                                    reason=\"OrderProcessingError\",\n+                                    error_description=str(exc)[:4096],\n+                                )\n+                            else:\n+                                receiver.complete_message(message)\n+                                LOGGER.info(\n+                                    \"Processed order %s for %s: %s\",\n+                                    order.order_id,\n+                                    order.customer_name,\n+                                    order.status.value,\n+                                )\n+                            processed += 1\n+            except ServiceBusError as exc:\n+                LOGGER.debug(\"No available order session: %s\", exc)\n+                break\n+        return processed\n+\n+    def inspect_dead_letters(self, max_messages: int = 20) -\u003e list[DeadLetterRecord]:\n+        receiver = self.client.get_queue_receiver(\n+            queue_name=self.queue_name,\n+            sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+            session_id=NEXT_AVAILABLE_SESSION,\n+            max_wait_time=self.max_wait_time,\n+        )\n+        try:\n+            with receiver:\n+                return [\n+                    _dead_letter_record(message)\n+                    for message in receiver.peek_messages(\n+                        max_message_count=max_messages,\n+                        from_sequence_number=1,\n+                    )\n+                ]\n+        except ServiceBusError as exc:\n+            LOGGER.debug(\"No available dead-letter session: %s\", exc)\n+            return []\n+\n+    def reprocess_dead_letters(\n+        self, sender: OrderSender, max_messages: int = 20\n+    ) -\u003e int:\n+        reprocessed = 0\n+        while reprocessed \u003c max_messages:\n+            receiver = self.client.get_queue_receiver(\n+                queue_name=self.queue_name,\n+                sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=self.max_wait_time,\n+            )\n+            try:\n+                with receiver:\n+                    messages = receiver.receive_messages(\n+                        max_message_count=min(20, max_messages - reprocessed),\n+                        max_wait_time=self.max_wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        try:\n+                            order = Order.from_json(_message_body(message))\n+                            sender.send_order(order)\n+                        except (TypeError, ValueError, UnicodeError) as exc:\n+                            LOGGER.error(\n+                                \"Dead-lettered message %s cannot be reprocessed: %s\",\n+                                message.correlation_id,\n+                                exc,\n+                            )\n+                            receiver.abandon_message(message)\n+                        else:\n+                            receiver.complete_message(message)\n+                            reprocessed += 1\n+                            LOGGER.info(\"Requeued order %s\", order.order_id)\n+            except ServiceBusError as exc:\n+                LOGGER.debug(\"No available dead-letter session: %s\", exc)\n+                break\n+        return reprocessed\n+\n+\n+class AsyncOrderProcessor:\n+    def __init__(\n+        self,\n+        client: AsyncServiceBusClient,\n+        queue_name: str,\n+        max_wait_time: float = 5,\n+    ) -\u003e None:\n+        self.client = client\n+        self.queue_name = queue_name\n+        self.max_wait_time = max_wait_time\n+\n+    async def process(self, max_messages: int | None = None) -\u003e int:\n+        processed = 0\n+        while max_messages is None or processed \u003c max_messages:\n+            receiver = self.client.get_queue_receiver(\n+                queue_name=self.queue_name,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=self.max_wait_time,\n+            )\n+            try:\n+                async with receiver:\n+                    while max_messages is None or processed \u003c max_messages:\n+                        messages = await receiver.receive_messages(\n+                            max_message_count=(\n+                                min(20, max_messages - processed)\n+                                if max_messages is not None\n+                                else 20\n+                            ),\n+                            max_wait_time=self.max_wait_time,\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            try:\n+                                order = _process_order(message)\n+                            except (TypeError, ValueError, UnicodeError, RuntimeError) as exc:\n+                                LOGGER.error(\n+                                    \"Order %s failed: %s\", message.correlation_id, exc\n+                                )\n+                                await receiver.dead_letter_message(\n+                                    message,\n+                                    reason=\"OrderProcessingError\",\n+                                    error_description=str(exc)[:4096],\n+                                )\n+                            else:\n+                                await receiver.complete_message(message)\n+                                LOGGER.info(\n+                                    \"Processed order %s for %s: %s\",\n+                                    order.order_id,\n+                                    order.customer_name,\n+                                    order.status.value,\n+                                )\n+                            processed += 1\n+            except ServiceBusError as exc:\n+                LOGGER.debug(\"No available order session: %s\", exc)\n+                break\n+        return processed\n+\n+    async def inspect_dead_letters(\n+        self, max_messages: int = 20\n+    ) -\u003e list[DeadLetterRecord]:\n+        receiver = self.client.get_queue_receiver(\n+            queue_name=self.queue_name,\n+            sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+            session_id=NEXT_AVAILABLE_SESSION,\n+            max_wait_time=self.max_wait_time,\n+        )\n+        try:\n+            async with receiver:\n+                messages = await receiver.peek_messages(\n+                    max_message_count=max_messages,\n+                    from_sequence_number=1,\n+                )\n+                return [_dead_letter_record(message) for message in messages]\n+        except ServiceBusError as exc:\n+            LOGGER.debug(\"No available dead-letter session: %s\", exc)\n+            return []\n+\n+    async def reprocess_dead_letters(\n+        self, sender: AsyncOrderSender, max_messages: int = 20\n+    ) -\u003e int:\n+        reprocessed = 0\n+        while reprocessed \u003c max_messages:\n+            receiver = self.client.get_queue_receiver(\n+                queue_name=self.queue_name,\n+                sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=self.max_wait_time,\n+            )\n+            try:\n+                async with receiver:\n+                    messages = await receiver.receive_messages(\n+                        max_message_count=min(20, max_messages - reprocessed),\n+                        max_wait_time=self.max_wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        try:\n+                            order = Order.from_json(_message_body(message))\n+                            await sender.send_order(order)\n+                        except (TypeError, ValueError, UnicodeError) as exc:\n+                            LOGGER.error(\n+                                \"Dead-lettered message %s cannot be reprocessed: %s\",\n+                                message.correlation_id,\n+                                exc,\n+                            )\n+                            await receiver.abandon_message(message)\n+                        else:\n+                            await receiver.complete_message(message)\n+                            reprocessed += 1\n+                            LOGGER.info(\"Requeued order %s\", order.order_id)\n+            except ServiceBusError as exc:\n+                LOGGER.debug(\"No available dead-letter session: %s\", exc)\n+                break\n+        return reprocessed\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from order_processor.model import Order\n+from order_processor.processor import AsyncOrderProcessor, OrderProcessor\n+from order_processor.sender import AsyncOrderSender, OrderSender\n+\n+LOGGER = logging.getLogger(__name__)\n+\n+\n+def _settings() -\u003e tuple[str, str]:\n+    namespace = os.environ[\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.environ[\"SERVICE_BUS_QUEUE_NAME\"]\n+    return namespace, queue_name\n+\n+\n+def _orders(prefix: str) -\u003e list[Order]:\n+    return [\n+        Order(f\"{prefix}-001\", \"Ada\", \"Keyboard\", 1, 89.99),\n+        Order(f\"{prefix}-002\", \"Grace\", \"Monitor\", 2, 699.98),\n+        Order(f\"{prefix}-003\", \"Ada\", \"Workstation\", 1, 2_499.00),\n+    ]\n+\n+\n+def run_sync_demo(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting synchronous demo\")\n+    credential = DefaultAzureCredential()\n+    try:\n+        with ServiceBusClient(namespace, credential) as client:\n+            sender = OrderSender(client, queue_name)\n+            processor = OrderProcessor(client, queue_name)\n+\n+            sender.send_orders(_orders(\"sync\"))\n+            sender.send_order(\n+                Order(\"sync-retry\", \"Linus\", \"USB hub\", 1, 39.00),\n+                application_properties={\"simulate_failure\": True},\n+            )\n+            processor.process()\n+\n+            for record in processor.inspect_dead_letters():\n+                LOGGER.info(\n+                    \"Dead letter %s: %s (%s)\",\n+                    record.order_id,\n+                    record.reason,\n+                    record.description,\n+                )\n+\n+            processor.reprocess_dead_letters(sender)\n+            processor.process()\n+    finally:\n+        credential.close()\n+\n+\n+async def run_async_demo(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting asynchronous demo\")\n+    credential = AsyncDefaultAzureCredential()\n+    try:\n+        async with AsyncServiceBusClient(namespace, credential) as client:\n+            sender = AsyncOrderSender(client, queue_name)\n+            processor = AsyncOrderProcessor(client, queue_name)\n+\n+            await sender.send_orders(_orders(\"async\"))\n+            await sender.send_order(\n+                Order(\"async-retry\", \"Margaret\", \"Mouse\", 1, 49.00),\n+                application_properties={\"simulate_failure\": True},\n+            )\n+            await processor.process()\n+\n+            for record in await processor.inspect_dead_letters():\n+                LOGGER.info(\n+                    \"Dead letter %s: %s (%s)\",\n+                    record.order_id,\n+                    record.reason,\n+                    record.description,\n+                )\n+\n+            await processor.reprocess_dead_letters(sender)\n+            await processor.process()\n+    finally:\n+        await credential.close()\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    namespace, queue_name = _settings()\n+    run_sync_demo(namespace, queue_name)\n+    asyncio.run(run_async_demo(namespace, queue_name))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-servicebus\u003e=7.12.2,\u003c8.0.0\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+This sample sends and processes orders with the synchronous and asynchronous\n+Azure Service Bus Python clients.\n+\n+## Prerequisites\n+\n+Use an existing Azure Service Bus queue with **sessions enabled**. The sender\n+sets each message's session ID to the customer name, and each processor drains\n+one session before accepting another. This preserves order for each customer\n+without interleaving customers.\n+\n+Authenticate with any identity supported by `DefaultAzureCredential` and set:\n+\n+```text\n+SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=\u003cnamespace\u003e.servicebus.windows.net\n+SERVICE_BUS_QUEUE_NAME=\u003csession-enabled-queue\u003e\n+```\n+\n+Install dependencies with `python -m pip install -r requirements.txt`, then run\n+`python main.py`.\n+\n+Orders above the sender's high-priority threshold (USD 1,000 by default) carry\n+a `priority=high` application property and are scheduled 30 seconds into the\n+future for a fraud-review window. Service Bus queues do not provide native\n+priority ordering, so consumers can use this property for application-specific\n+handling.\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"python -m compileall -q .","description":"Compile Python project","initial_wait":120}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"python -m pip install --disable-pip-version-check --quiet -r requirements.txt","description":"Install project dependencies","initial_wait":180}
```

### 🔧 6. powershell

**Input:**

```json
{"command":"python -c \"import inspect; from azure.servicebus import ServiceBusMessage; from azure.servicebus.aio import ServiceBusSender; print(inspect.signature(ServiceBusMessage)); print(inspect.signature(ServiceBusSender.create_message_batch))\"","description":"Check Azure SDK method signatures"}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"python -c \"from order_processor.model import Order; from order_processor.sender import _MessageFactory; o=Order('o1','Ada','Keyboard',1,1200); assert Order.from_json(o.to_json()) == o; m=_MessageFactory(1000).create(o); assert m.correlation_id == 'o1' and m.session_id == 'Ada' and m.scheduled_enqueue_time_utc is not None; print('local checks passed')\"","description":"Test model and message creation"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"python -c \"import main; from order_processor.processor import OrderProcessor, AsyncOrderProcessor; print('all modules imported')\"","description":"Import all project modules"}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"python -c \"import inspect; from azure.servicebus.aio import ServiceBusSender, ServiceBusReceiver; print('batch async', inspect.iscoroutinefunction(ServiceBusSender.create_message_batch)); print('peek async', inspect.iscoroutinefunction(ServiceBusReceiver.peek_messages)); print('receive async', inspect.iscoroutinefunction(ServiceBusReceiver.receive_messages))\"","description":"Verify async Azure methods"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"python -c \"from azure.servicebus import ServiceBusMessage; from order_processor.processor import _message_body; m=ServiceBusMessage('{\\\"x\\\":1}'); print(_message_body(m))\"","description":"Test Service Bus body decoding"}
```

### 🔧 11. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: order_processor/processor.py\n@@\n     def reprocess_dead_letters(\n         self, sender: OrderSender, max_messages: int = 20\n     ) -\u003e int:\n         reprocessed = 0\n-        while reprocessed \u003c max_messages:\n+        examined = 0\n+        while examined \u003c max_messages:\n@@\n                     messages = receiver.receive_messages(\n-                        max_message_count=min(20, max_messages - reprocessed),\n+                        max_message_count=min(20, max_messages - examined),\n                         max_wait_time=self.max_wait_time,\n                     )\n@@\n                     for message in messages:\n+                        examined += 1\n                         try:\n@@\n     async def reprocess_dead_letters(\n         self, sender: AsyncOrderSender, max_messages: int = 20\n     ) -\u003e int:\n         reprocessed = 0\n-        while reprocessed \u003c max_messages:\n+        examined = 0\n+        while examined \u003c max_messages:\n@@\n                     messages = await receiver.receive_messages(\n-                        max_message_count=min(20, max_messages - reprocessed),\n+                        max_message_count=min(20, max_messages - examined),\n                         max_wait_time=self.max_wait_time,\n                     )\n@@\n                     for message in messages:\n+                        examined += 1\n                         try:\n*** End Patch\n"
```

### 🔧 12. powershell

**Input:**

```json
{"command":"python -m compileall -q .; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -c 'from azure.servicebus import ServiceBusMessage; from order_processor.processor import _message_body; m = ServiceBusMessage(\"{`\"x`\":1}\"); assert _message_body(m) == \"{`\"x`\":1}\"; print(\"compile and body checks passed\")'","description":"Compile and test body decoding","initial_wait":120}
```

### 🔧 13. powershell

**Input:**

```json
{"command":"python -c 'from order_processor.model import Order, OrderStatus; cases=[Order(\"1\",\"A\",\"P\",1,0),Order(\"2\",\"B\",\"Q\",2,3.5,OrderStatus.COMPLETED)]; assert [Order.from_json(x.to_json()) for x in cases] == cases; print(\"model round trips passed\")'","description":"Test order JSON round trips"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"python -m compileall -q .; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -c \"from azure.servicebus import ServiceBusMessage; from order_processor.processor import _message_body; m=ServiceBusMessage('hello'); assert _message_body(m) == 'hello'; print('compile and body checks passed')\"","description":"Compile and test body decoding","initial_wait":120}
```

## Generated Files

- `README.md`
- `main.py`
- `order_processor\__init__.py`
- `order_processor\model.py`
- `order_processor\processor.py`
- `order_processor\sender.py`
- `requirements.txt`

## Copilot Response

I’ll scaffold the project around session-enabled queues so per-customer ordering is enforced by Service Bus, then validate the modules locally without contacting Azure.

The implementation will use one Service Bus session per customer and drain each accepted session before moving to another. Batch sends are grouped by customer because session-enabled queues require consistent session routing within a batch.

The project is now written. I’m checking it against the installed SDK signatures and exercising serialization plus message construction locally; no Azure connection or resource operation will be performed.

The SDK checks confirmed the async batch and receiver APIs are awaited correctly. I also bounded dead-letter reprocessing by messages examined, preventing an unrecoverable malformed payload from being repeatedly abandoned and immediately retried in one run.

Created the Azure Service Bus order-processing project with:

- JSON-serializable `Order` dataclass and validated statuses.
- Synchronous and asynchronous senders with individual sends, size-aware batching, correlation IDs, customer sessions, priority metadata, and 30-second scheduling.
- Session-based sync/async processors with dead-lettering, inspection, and reprocessing.
- `DefaultAzureCredential` demo covering sync then async workflows.
- Azure SDK dependencies and session-enabled queue setup documentation.

Set `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE` and `SERVICE_BUS_QUEUE_NAME`, then run `python main.py`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0/1 criteria passed.

### Strengths

- Sync and async receiver paths consistently catch ServiceBusError, while malformed order payloads are handled and dead-lettered with a reason and description.

### Issues

- Azure SDK exceptions are only partially handled, and caught ServiceBusError instances may represent genuine service failures rather than unavailable sessions.

## Grader Results

- order-processor.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (12/14)
      - Sender uses `ServiceBusClient.get_queue_sender()` (or async equivalent): Pass
      - Processor uses `ServiceBusClient.get_queue_receiver()` or session-enabled receiver: Pass
      - Batch sending: creates `ServiceBusMessageBatch`, checks `add_message()` return or catches `MessageSizeExceededError`: Pass
      - Handles the case where a message doesn't fit in the current batch: Pass
      - Scheduled delivery: uses `schedule_messages()` or `ServiceBusMessage(scheduled_enqueue_time_utc=...)` (~30s delay): Pass
      - Correlation: sets order ID via `correlation_id` property on `ServiceBusMessage`: Pass
      - Dead-letter: explicitly dead-letters failed messages with `receiver.dead_letter_message()` and a reason string: Pass
      - Dead-letter queue reading: uses `sub_queue=ServiceBusSubQueue.DEAD_LETTER` on receiver: Pass
      - Session-aware processing: uses `session_id` on messages and session-enabled receiver: Pass
      - Session ID keyed by customer name for ordered processing: Pass
      - Context manager pattern (`with` statements) for all clients, senders, and receivers: Pass
      - Catches `ServiceBusError` and distinguishes transient vs non-transient errors: Fail
      - Error handler logs entity name and error details: Fail
      - Async versions use `azure.servicebus.aio` module: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Fail
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 86% | 1.00 | 0.8571 | 17.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.6% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 4.8571** | **60.7%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-order-processor --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
