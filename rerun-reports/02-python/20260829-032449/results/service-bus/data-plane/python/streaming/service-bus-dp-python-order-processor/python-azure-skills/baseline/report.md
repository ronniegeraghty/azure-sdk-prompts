# Evaluation Report: service-bus-dp-python-order-processor

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 316.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-order-processor` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 15/21 |
| Duration | 316.9s |
| Timestamp | 2026-08-28T21:13:15Z |
| Files Generated | 11 |
| Event Count | 14780 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 198.1s |
| Review | 118.0s |
| **Total** | **316.9s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** python-azure-skills/baseline

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=242832 out=10320 |
| Turn Count | 9 |

## Tool Calls

### ✅ 1. glob — 8849ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. powershell — 9928ms

**Input:**

```json
{"command":"Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace contents"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 3. apply_patch — 1265ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-servicebus\u003e=7.12.0,\u003c8.0.0\n+\n*** Add File: order_system/__init__.py\n+\"\"\"Azure Service Bus order processing example.\"\"\"\n+\n+from .model import Order, OrderStatus\n+\n+__all__ = [\"Order\", \"OrderStatus\"]\n+\n*** Add File: order_system/model.py\n+from __future__ import annotations\n+\n+import json\n+from dataclasses import asdict, dataclass\n+from enum import Enum\n+from typing import Any\n+\n+\n+class OrderStatus(str, Enum):\n+    PENDING = \"pending\"\n+    PROCESSING = \"processing\"\n+    COMPLETED = \"completed\"\n+    FAILED = \"failed\"\n+\n+\n+@dataclass(slots=True)\n+class Order:\n+    order_id: str\n+    customer_name: str\n+    product: str\n+    quantity: int\n+    total_price: float\n+    status: OrderStatus = OrderStatus.PENDING\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.order_id.strip():\n+            raise ValueError(\"order_id must not be empty\")\n+        if not self.customer_name.strip():\n+            raise ValueError(\"customer_name must not be empty\")\n+        if self.quantity \u003c= 0:\n+            raise ValueError(\"quantity must be greater than zero\")\n+        if self.total_price \u003c 0:\n+            raise ValueError(\"total_price must not be negative\")\n+        if not isinstance(self.status, OrderStatus):\n+            self.status = OrderStatus(self.status)\n+\n+    def to_dict(self) -\u003e dict[str, Any]:\n+        data = asdict(self)\n+        data[\"status\"] = self.status.value\n+        return data\n+\n+    def to_json(self) -\u003e str:\n+        return json.dumps(self.to_dict(), separators=(\",\", \":\"))\n+\n+    @classmethod\n+    def from_json(cls, value: str | bytes) -\u003e Order:\n+        data = json.loads(value)\n+        if not isinstance(data, dict):\n+            raise ValueError(\"order JSON must contain an object\")\n+        return cls(\n+            order_id=str(data[\"order_id\"]),\n+            customer_name=str(data[\"customer_name\"]),\n+            product=str(data[\"product\"]),\n+            quantity=int(data[\"quantity\"]),\n+            total_price=float(data[\"total_price\"]),\n+            status=OrderStatus(data[\"status\"]),\n+        )\n+\n*** Add File: order_system/messages.py\n+from __future__ import annotations\n+\n+from datetime import datetime, timedelta, timezone\n+\n+from azure.servicebus import ServiceBusMessage\n+\n+from .model import Order\n+\n+\n+HIGH_PRIORITY_DELAY = timedelta(seconds=30)\n+\n+\n+def create_order_message(\n+    order: Order, high_priority_threshold: float\n+) -\u003e ServiceBusMessage:\n+    is_high_priority = order.total_price \u003e high_priority_threshold\n+    return ServiceBusMessage(\n+        order.to_json(),\n+        content_type=\"application/json\",\n+        correlation_id=order.order_id,\n+        session_id=order.customer_name,\n+        application_properties={\n+            \"priority\": \"high\" if is_high_priority else \"normal\",\n+        },\n+        scheduled_enqueue_time_utc=(\n+            datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY\n+            if is_high_priority\n+            else None\n+        ),\n+    )\n+\n*** Add File: order_system/sender.py\n+from __future__ import annotations\n+\n+from collections.abc import Iterable\n+from typing import Any\n+\n+from azure.servicebus import ServiceBusClient\n+from azure.servicebus.exceptions import MessageSizeExceededError\n+\n+from .messages import create_order_message\n+from .model import Order\n+\n+\n+class OrderSender:\n+    def __init__(\n+        self,\n+        fully_qualified_namespace: str,\n+        queue_name: str,\n+        credential: Any,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self._client = ServiceBusClient(\n+            fully_qualified_namespace=fully_qualified_namespace,\n+            credential=credential,\n+        )\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    def __enter__(self) -\u003e OrderSender:\n+        return self\n+\n+    def __exit__(self, *args: object) -\u003e None:\n+        self.close()\n+\n+    def close(self) -\u003e None:\n+        self._client.close()\n+\n+    def send_order(self, order: Order) -\u003e None:\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            sender.send_messages(\n+                create_order_message(order, self._high_priority_threshold)\n+            )\n+\n+    def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            batch = sender.create_message_batch()\n+            batch_count = 0\n+\n+            for order in orders:\n+                message = create_order_message(\n+                    order, self._high_priority_threshold\n+                )\n+                try:\n+                    batch.add_message(message)\n+                    batch_count += 1\n+                except MessageSizeExceededError:\n+                    if batch_count == 0:\n+                        raise\n+                    sender.send_messages(batch)\n+                    batch = sender.create_message_batch()\n+                    batch.add_message(message)\n+                    batch_count = 1\n+\n+            if batch_count:\n+                sender.send_messages(batch)\n+\n*** Add File: order_system/async_sender.py\n+from __future__ import annotations\n+\n+from collections.abc import Iterable\n+from typing import Any\n+\n+from azure.servicebus.aio import ServiceBusClient\n+from azure.servicebus.exceptions import MessageSizeExceededError\n+\n+from .messages import create_order_message\n+from .model import Order\n+\n+\n+class AsyncOrderSender:\n+    def __init__(\n+        self,\n+        fully_qualified_namespace: str,\n+        queue_name: str,\n+        credential: Any,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self._client = ServiceBusClient(\n+            fully_qualified_namespace=fully_qualified_namespace,\n+            credential=credential,\n+        )\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    async def __aenter__(self) -\u003e AsyncOrderSender:\n+        return self\n+\n+    async def __aexit__(self, *args: object) -\u003e None:\n+        await self.close()\n+\n+    async def close(self) -\u003e None:\n+        await self._client.close()\n+\n+    async def send_order(self, order: Order) -\u003e None:\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            await sender.send_messages(\n+                create_order_message(order, self._high_priority_threshold)\n+            )\n+\n+    async def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            batch = await sender.create_message_batch()\n+            batch_count = 0\n+\n+            for order in orders:\n+                message = create_order_message(\n+                    order, self._high_priority_threshold\n+                )\n+                try:\n+                    batch.add_message(message)\n+                    batch_count += 1\n+                except MessageSizeExceededError:\n+                    if batch_count == 0:\n+                        raise\n+                    await sender.send_messages(batch)\n+                    batch = await sender.create_message_batch()\n+                    batch.add_message(message)\n+                    batch_count = 1\n+\n+            if batch_count:\n+                await sender.send_messages(batch)\n+\n*** Add File: order_system/processor.py\n+from __future__ import annotations\n+\n+import logging\n+from collections.abc import Callable\n+from dataclasses import dataclass\n+from typing import Any\n+\n+from azure.servicebus import (\n+    NEXT_AVAILABLE_SESSION,\n+    ServiceBusClient,\n+    ServiceBusMessage,\n+    ServiceBusReceivedMessage,\n+    ServiceBusSubQueue,\n+)\n+from azure.servicebus.exceptions import OperationTimeoutError\n+\n+from .model import Order, OrderStatus\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+@dataclass(frozen=True, slots=True)\n+class DeadLetterRecord:\n+    message_id: str | None\n+    reason: str | None\n+    description: str | None\n+    body: str\n+\n+\n+def message_body(message: ServiceBusReceivedMessage) -\u003e bytes:\n+    return b\"\".join(bytes(section) for section in message.body)\n+\n+\n+def default_order_handler(order: Order) -\u003e None:\n+    order.status = OrderStatus.PROCESSING\n+    logger.info(\n+        \"Processing order %s for %s: %d x %s ($%.2f)\",\n+        order.order_id,\n+        order.customer_name,\n+        order.quantity,\n+        order.product,\n+        order.total_price,\n+    )\n+    order.status = OrderStatus.COMPLETED\n+    logger.info(\"Completed order %s\", order.order_id)\n+\n+\n+class OrderProcessor:\n+    def __init__(\n+        self,\n+        fully_qualified_namespace: str,\n+        queue_name: str,\n+        credential: Any,\n+        handler: Callable[[Order], None] = default_order_handler,\n+    ) -\u003e None:\n+        self._client = ServiceBusClient(\n+            fully_qualified_namespace=fully_qualified_namespace,\n+            credential=credential,\n+        )\n+        self._queue_name = queue_name\n+        self._handler = handler\n+\n+    def __enter__(self) -\u003e OrderProcessor:\n+        return self\n+\n+    def __exit__(self, *args: object) -\u003e None:\n+        self.close()\n+\n+    def close(self) -\u003e None:\n+        self._client.close()\n+\n+    def process_available_orders(\n+        self,\n+        max_sessions: int | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        sessions = 0\n+        while max_sessions is None or sessions \u003c max_sessions:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    queue_name=self._queue_name,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    max_wait_time=max_wait_time,\n+                )\n+                with receiver:\n+                    sessions += 1\n+                    while messages := receiver.receive_messages(\n+                        max_message_count=50,\n+                        max_wait_time=max_wait_time,\n+                    ):\n+                        for message in messages:\n+                            if self._process_message(receiver, message):\n+                                processed += 1\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    def _process_message(self, receiver: Any, message: Any) -\u003e bool:\n+        try:\n+            order = Order.from_json(message_body(message))\n+            self._handler(order)\n+        except Exception as exc:\n+            reason = f\"{type(exc).__name__}: order processing failed\"\n+            logger.exception(\"Dead-lettering invalid order message\")\n+            receiver.dead_letter_message(\n+                message,\n+                reason=reason[:4096],\n+                error_description=str(exc)[:4096],\n+            )\n+            return False\n+        receiver.complete_message(message)\n+        return True\n+\n+    def read_dead_letters(\n+        self,\n+        reprocess: bool = False,\n+        max_sessions: int | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e list[DeadLetterRecord]:\n+        records: list[DeadLetterRecord] = []\n+        sender = (\n+            self._client.get_queue_sender(self._queue_name)\n+            if reprocess\n+            else None\n+        )\n+        if sender is not None:\n+            sender.__enter__()\n+        try:\n+            sessions = 0\n+            while max_sessions is None or sessions \u003c max_sessions:\n+                try:\n+                    receiver = self._client.get_queue_receiver(\n+                        queue_name=self._queue_name,\n+                        sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                        session_id=NEXT_AVAILABLE_SESSION,\n+                        max_wait_time=max_wait_time,\n+                    )\n+                    with receiver:\n+                        sessions += 1\n+                        while messages := receiver.receive_messages(\n+                            max_message_count=50,\n+                            max_wait_time=max_wait_time,\n+                        ):\n+                            for message in messages:\n+                                body = message_body(message)\n+                                records.append(\n+                                    DeadLetterRecord(\n+                                        message_id=message.message_id,\n+                                        reason=message.dead_letter_reason,\n+                                        description=message.dead_letter_error_description,\n+                                        body=body.decode(\"utf-8\", errors=\"replace\"),\n+                                    )\n+                                )\n+                                logger.info(\n+                                    \"DLQ message %s: %s\",\n+                                    message.message_id,\n+                                    message.dead_letter_reason,\n+                                )\n+                                if sender is not None:\n+                                    order = Order.from_json(body)\n+                                    sender.send_messages(\n+                                        ServiceBusMessage(\n+                                            order.to_json(),\n+                                            content_type=\"application/json\",\n+                                            correlation_id=order.order_id,\n+                                            session_id=order.customer_name,\n+                                        )\n+                                    )\n+                                    receiver.complete_message(message)\n+                except OperationTimeoutError:\n+                    break\n+        finally:\n+            if sender is not None:\n+                sender.__exit__(None, None, None)\n+        return records\n+\n*** Add File: order_system/async_processor.py\n+from __future__ import annotations\n+\n+import logging\n+from collections.abc import Awaitable, Callable\n+from typing import Any\n+\n+from azure.servicebus import (\n+    NEXT_AVAILABLE_SESSION,\n+    ServiceBusMessage,\n+    ServiceBusSubQueue,\n+)\n+from azure.servicebus.aio import ServiceBusClient\n+from azure.servicebus.exceptions import OperationTimeoutError\n+\n+from .model import Order, OrderStatus\n+from .processor import DeadLetterRecord, message_body\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+async def default_order_handler(order: Order) -\u003e None:\n+    order.status = OrderStatus.PROCESSING\n+    logger.info(\n+        \"Processing order %s for %s: %d x %s ($%.2f)\",\n+        order.order_id,\n+        order.customer_name,\n+        order.quantity,\n+        order.product,\n+        order.total_price,\n+    )\n+    order.status = OrderStatus.COMPLETED\n+    logger.info(\"Completed order %s\", order.order_id)\n+\n+\n+class AsyncOrderProcessor:\n+    def __init__(\n+        self,\n+        fully_qualified_namespace: str,\n+        queue_name: str,\n+        credential: Any,\n+        handler: Callable[[Order], Awaitable[None]] = default_order_handler,\n+    ) -\u003e None:\n+        self._client = ServiceBusClient(\n+            fully_qualified_namespace=fully_qualified_namespace,\n+            credential=credential,\n+        )\n+        self._queue_name = queue_name\n+        self._handler = handler\n+\n+    async def __aenter__(self) -\u003e AsyncOrderProcessor:\n+        return self\n+\n+    async def __aexit__(self, *args: object) -\u003e None:\n+        await self.close()\n+\n+    async def close(self) -\u003e None:\n+        await self._client.close()\n+\n+    async def process_available_orders(\n+        self,\n+        max_sessions: int | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        sessions = 0\n+        while max_sessions is None or sessions \u003c max_sessions:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    queue_name=self._queue_name,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    max_wait_time=max_wait_time,\n+                )\n+                async with receiver:\n+                    sessions += 1\n+                    while messages := await receiver.receive_messages(\n+                        max_message_count=50,\n+                        max_wait_time=max_wait_time,\n+                    ):\n+                        for message in messages:\n+                            if await self._process_message(receiver, message):\n+                                processed += 1\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    async def _process_message(self, receiver: Any, message: Any) -\u003e bool:\n+        try:\n+            order = Order.from_json(message_body(message))\n+            await self._handler(order)\n+        except Exception as exc:\n+            reason = f\"{type(exc).__name__}: order processing failed\"\n+            logger.exception(\"Dead-lettering invalid order message\")\n+            await receiver.dead_letter_message(\n+                message,\n+                reason=reason[:4096],\n+                error_description=str(exc)[:4096],\n+            )\n+            return False\n+        await receiver.complete_message(message)\n+        return True\n+\n+    async def read_dead_letters(\n+        self,\n+        reprocess: bool = False,\n+        max_sessions: int | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e list[DeadLetterRecord]:\n+        records: list[DeadLetterRecord] = []\n+        sender = (\n+            self._client.get_queue_sender(self._queue_name)\n+            if reprocess\n+            else None\n+        )\n+        if sender is not None:\n+            await sender.__aenter__()\n+        try:\n+            sessions = 0\n+            while max_sessions is None or sessions \u003c max_sessions:\n+                try:\n+                    receiver = self._client.get_queue_receiver(\n+                        queue_name=self._queue_name,\n+                        sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                        session_id=NEXT_AVAILABLE_SESSION,\n+                        max_wait_time=max_wait_time,\n+                    )\n+                    async with receiver:\n+                        sessions += 1\n+                        while messages := await receiver.receive_messages(\n+                            max_message_count=50,\n+                            max_wait_time=max_wait_time,\n+                        ):\n+                            for message in messages:\n+                                body = message_body(message)\n+                                records.append(\n+                                    DeadLetterRecord(\n+                                        message_id=message.message_id,\n+                                        reason=message.dead_letter_reason,\n+                                        description=message.dead_letter_error_description,\n+                                        body=body.decode(\"utf-8\", errors=\"replace\"),\n+                                    )\n+                                )\n+                                logger.info(\n+                                    \"DLQ message %s: %s\",\n+                                    message.message_id,\n+                                    message.dead_letter_reason,\n+                                )\n+                                if sender is not None:\n+                                    order = Order.from_json(body)\n+                                    await sender.send_messages(\n+                                        ServiceBusMessage(\n+                                            order.to_json(),\n+                                            content_type=\"application/json\",\n+                                            correlation_id=order.order_id,\n+                                            session_id=order.customer_name,\n+                                        )\n+                                    )\n+                                    await receiver.complete_message(message)\n+                except OperationTimeoutError:\n+                    break\n+        finally:\n+            if sender is not None:\n+                await sender.__aexit__(None, None, None)\n+        return records\n+\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+from collections.abc import Iterable\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from order_system.async_processor import AsyncOrderProcessor\n+from order_system.async_sender import AsyncOrderSender\n+from order_system.model import Order\n+from order_system.processor import OrderProcessor\n+from order_system.sender import OrderSender\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+def sample_orders(prefix: str) -\u003e list[Order]:\n+    return [\n+        Order(f\"{prefix}-001\", \"Ada Lovelace\", \"Keyboard\", 1, 129.99),\n+        Order(f\"{prefix}-002\", \"Ada Lovelace\", \"Monitor\", 2, 699.98),\n+        Order(f\"{prefix}-003\", \"Grace Hopper\", \"Workstation\", 1, 2_499.00),\n+    ]\n+\n+\n+def send_invalid_sync(\n+    namespace: str,\n+    queue_name: str,\n+    credential: DefaultAzureCredential,\n+    prefix: str,\n+) -\u003e None:\n+    with ServiceBusClient(namespace, credential) as client:\n+        with client.get_queue_sender(queue_name) as sender:\n+            sender.send_messages(\n+                ServiceBusMessage(\n+                    \"{not valid JSON\",\n+                    content_type=\"application/json\",\n+                    correlation_id=f\"{prefix}-invalid\",\n+                    session_id=\"Invalid order demo\",\n+                )\n+            )\n+\n+\n+async def send_invalid_async(\n+    namespace: str,\n+    queue_name: str,\n+    credential: AsyncDefaultAzureCredential,\n+    prefix: str,\n+) -\u003e None:\n+    async with AsyncServiceBusClient(namespace, credential) as client:\n+        async with client.get_queue_sender(queue_name) as sender:\n+            await sender.send_messages(\n+                ServiceBusMessage(\n+                    \"{not valid JSON\",\n+                    content_type=\"application/json\",\n+                    correlation_id=f\"{prefix}-invalid\",\n+                    session_id=\"Invalid order demo\",\n+                )\n+            )\n+\n+\n+def run_sync_demo(\n+    namespace: str,\n+    queue_name: str,\n+    orders: Iterable[Order],\n+) -\u003e None:\n+    logger.info(\"Starting synchronous demo\")\n+    credential = DefaultAzureCredential()\n+    try:\n+        with OrderSender(namespace, queue_name, credential) as sender:\n+            sender.send_orders(orders)\n+        send_invalid_sync(namespace, queue_name, credential, \"sync\")\n+        with OrderProcessor(namespace, queue_name, credential) as processor:\n+            processor.process_available_orders(max_wait_time=35)\n+            records = processor.read_dead_letters(max_wait_time=5)\n+            logger.info(\"Synchronous demo found %d DLQ message(s)\", len(records))\n+    finally:\n+        credential.close()\n+\n+\n+async def run_async_demo(\n+    namespace: str,\n+    queue_name: str,\n+    orders: Iterable[Order],\n+) -\u003e None:\n+    logger.info(\"Starting asynchronous demo\")\n+    credential = AsyncDefaultAzureCredential()\n+    try:\n+        async with AsyncOrderSender(\n+            namespace, queue_name, credential\n+        ) as sender:\n+            await sender.send_orders(orders)\n+        await send_invalid_async(namespace, queue_name, credential, \"async\")\n+        async with AsyncOrderProcessor(\n+            namespace, queue_name, credential\n+        ) as processor:\n+            await processor.process_available_orders(max_wait_time=35)\n+            records = await processor.read_dead_letters(max_wait_time=5)\n+            logger.info(\"Asynchronous demo found %d DLQ message(s)\", len(records))\n+    finally:\n+        await credential.close()\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    namespace = os.environ[\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.getenv(\"SERVICE_BUS_QUEUE_NAME\", \"orders\")\n+\n+    run_sync_demo(namespace, queue_name, sample_orders(\"sync\"))\n+    asyncio.run(\n+        run_async_demo(namespace, queue_name, sample_orders(\"async\"))\n+    )\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n+\n*** Add File: README.md\n+# Azure Service Bus order processor\n+\n+This example sends and processes orders with the synchronous and asynchronous\n+Azure Service Bus Python clients. The target queue **must have sessions\n+enabled**. Each message uses the customer name as its session ID, so one\n+processor owns and drains a customer's session before accepting another\n+customer's session.\n+\n+## Run\n+\n+1. Create and activate a virtual environment.\n+2. Install dependencies with `python -m pip install -r requirements.txt`.\n+3. Set `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE` to a namespace such as\n+   `example.servicebus.windows.net`.\n+4. Optionally set `SERVICE_BUS_QUEUE_NAME` (the default is `orders`).\n+5. Authenticate with a credential supported by `DefaultAzureCredential`.\n+6. Run `python main.py`.\n+\n+The demo sends valid and malformed messages, processes available queue\n+sessions, and inspects the dead-letter queue. Orders over $1,000 are marked\n+high priority and scheduled 30 seconds into the future.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/requirements.txt
@@ -1,0 +1,4 @@
+azure-identity>=1.17.0,<2.0.0
+azure-servicebus>=7.12.0,<8.0.0
+
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/__init__.py
@@ -1,0 +1,7 @@
+"""Azure Service Bus order processing example."""
+
+from .model import Order, OrderStatus
+
+__all__ = ["Order", "OrderStatus"]
+
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/model.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/model.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/model.py
@@ -1,0 +1,59 @@
+from __future__ import annotations
+
+import json
+from dataclasses import asdict, dataclass
+from enum import Enum
+from typing import Any
+
+
+class OrderStatus(str, Enum):
+    PENDING = "pending"
+    PROCESSING = "processing
... (truncated)
```

### ✅ 4. powershell — 14669ms

**Input:**

```json
{"command":"python -m compileall -q .","description":"Compile all Python files"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

### ✅ 5. powershell — 13422ms

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt","description":"Install Python dependencies","initial_wait":120,"mode":"sync"}
```

**Output:**

```

<shellId: 2 completed with exit code 0>
```

### ✅ 6. powershell — 13417ms

**Input:**

```json
{"command":"python -c \"import inspect; from azure.servicebus import ServiceBusMessage, ServiceBusClient; from azure.servicebus.aio import ServiceBusSender as AS; from azure.servicebus import ServiceBusSender as S; print(inspect.signature(ServiceBusMessage)); print('sync_batch', inspect.signature(S.create_message_batch)); print('async_batch', inspect.signature(AS.create_message_batch))\"","description":"Inspect Azure SDK signatures"}
```

**Output:**

```
all imports passed
<shellId: 5 completed with exit code 0>
```

### ✅ 7. powershell — 12141ms

**Input:**

```json
{"command":"python -c \"from order_system.model import Order; from order_system.messages import create_order_message; o=Order('1','Ada','Widget',2,1200); assert Order.from_json(o.to_json()) == o; m=create_order_message(o,1000); assert m.correlation_id=='1' and m.session_id=='Ada' and m.scheduled_enqueue_time_utc is not None; print('model and message checks passed')\"","description":"Check model serialization and messages"}
```

**Output:**

```
model and message checks passed
<shellId: 4 completed with exit code 0>
```

### ✅ 8. powershell — 10922ms

**Input:**

```json
{"command":"python -c \"import main; from order_system.sender import OrderSender; from order_system.async_sender import AsyncOrderSender; from order_system.processor import OrderProcessor; from order_system.async_processor import AsyncOrderProcessor; print('all imports passed')\"","description":"Check project imports"}
```

**Output:**

```
(body: 'Optional[Union[str, bytes]]', *, application_properties: "Optional[Dict[Union[str, bytes], 'PrimitiveTypes']]" = None, session_id: 'Optional[str]' = None, message_id: 'Optional[str]' = None, scheduled_enqueue_time_utc: 'Optional[datetime.datetime]' = None, time_to_live: 'Optional[datetime.timedelta]' = None, content_type: 'Optional[str]' = None, correlation_id: 'Optional[str]' = None, subject: 'Optional[str]' = None, partition_key: 'Optional[str]' = None, to: 'Optional[str]' = None, reply_to: 'Optional[str]' = None, reply_to_session_id: 'Optional[str]' = None, **kwargs: 'Any') -> 'None'
sync_batch (self, max_size_in_bytes: Optional[int] = None) -> azure.servicebus._common.message.ServiceBusMessageBatch
async_batch (self, max_size_in_bytes: Optional[int] = None) -> azure.servicebus._common.message.ServiceBusMessageBatch
<shellId: 3 completed with exit code 0>
```

### ✅ 9. powershell — 6956ms

**Input:**

```json
{"command":"python -c \"from azure.servicebus import ServiceBusMessage, ServiceBusMessageBatch; b=ServiceBusMessageBatch(max_size_in_bytes=100000); b.add_message(ServiceBusMessage('a',session_id='A')); b.add_message(ServiceBusMessage('b',session_id='B')); print(len(b))\"","description":"Test mixed-session message batching"}
```

**Output:**

```
2
<shellId: 6 completed with exit code 0>
```

### ✅ 10. powershell — 7110ms

**Input:**

```json
{"command":"python -c \"from azure.servicebus import ServiceBusMessage, ServiceBusMessageBatch; from azure.servicebus.exceptions import MessageSizeExceededError; b=ServiceBusMessageBatch(max_size_in_bytes=100); print('first');\ntry: b.add_message(ServiceBusMessage('x'*1000))\nexcept Exception as e: print(type(e).__name__, isinstance(e, MessageSizeExceededError))\"","description":"Confirm batch overflow exception"}
```

**Output:**

```
first
MessageSizeExceededError True
<shellId: 7 completed with exit code 0>
```

### ✅ 11. apply_patch — 1235ms

**Input:**

```json
"*** Begin Patch\n*** Update File: order_system/processor.py\n@@\n                                 logger.info(\n                                     \"DLQ message %s: %s\",\n                                     message.message_id,\n                                     message.dead_letter_reason,\n                                 )\n                                 if sender is not None:\n-                                    order = Order.from_json(body)\n-                                    sender.send_messages(\n-                                        ServiceBusMessage(\n-                                            order.to_json(),\n-                                            content_type=\"application/json\",\n-                                            correlation_id=order.order_id,\n-                                            session_id=order.customer_name,\n+                                    try:\n+                                        order = Order.from_json(body)\n+                                        sender.send_messages(\n+                                            ServiceBusMessage(\n+                                                order.to_json(),\n+                                                content_type=\"application/json\",\n+                                                correlation_id=order.order_id,\n+                                                session_id=order.customer_name,\n+                                            )\n                                         )\n-                                    )\n-                                    receiver.complete_message(message)\n+                                        receiver.complete_message(message)\n+                                    except (KeyError, TypeError, ValueError) as exc:\n+                                        logger.warning(\n+                                            \"DLQ message %s cannot be reprocessed: %s\",\n+                                            message.message_id,\n+                                            exc,\n+                                        )\n                 except OperationTimeoutError:\n                     break\n*** Update File: order_system/async_processor.py\n@@\n                                 logger.info(\n                                     \"DLQ message %s: %s\",\n                                     message.message_id,\n                                     message.dead_letter_reason,\n                                 )\n                                 if sender is not None:\n-                                    order = Order.from_json(body)\n-                                    await sender.send_messages(\n-                                        ServiceBusMessage(\n-                                            order.to_json(),\n-                                            content_type=\"application/json\",\n-                                            correlation_id=order.order_id,\n-                                            session_id=order.customer_name,\n+                                    try:\n+                                        order = Order.from_json(body)\n+                                        await sender.send_messages(\n+                                            ServiceBusMessage(\n+                                                order.to_json(),\n+                                                content_type=\"application/json\",\n+                                                correlation_id=order.order_id,\n+                                                session_id=order.customer_name,\n+                                            )\n                                         )\n-                                    )\n-                                    await receiver.complete_message(message)\n+                                        await receiver.complete_message(message)\n+                                    except (KeyError, TypeError, ValueError) as exc:\n+                                        logger.warning(\n+                                            \"DLQ message %s cannot be reprocessed: %s\",\n+                                            message.message_id,\n+                                            exc,\n+                                        )\n                 except OperationTimeoutError:\n                     break\n*** Add File: tests/test_order_system.py\n+from __future__ import annotations\n+\n+import asyncio\n+import json\n+import unittest\n+from datetime import datetime, timezone\n+from unittest.mock import AsyncMock, Mock\n+\n+from azure.servicebus import ServiceBusMessage\n+\n+from order_system.async_processor import AsyncOrderProcessor\n+from order_system.messages import HIGH_PRIORITY_DELAY, create_order_message\n+from order_system.model import Order, OrderStatus\n+from order_system.processor import OrderProcessor\n+\n+\n+class OrderModelTests(unittest.TestCase):\n+    def test_json_round_trip(self) -\u003e None:\n+        order = Order(\n+            \"order-1\",\n+            \"Ada Lovelace\",\n+            \"Keyboard\",\n+            2,\n+            259.98,\n+            OrderStatus.PROCESSING,\n+        )\n+\n+        self.assertEqual(Order.from_json(order.to_json()), order)\n+        self.assertEqual(json.loads(order.to_json())[\"status\"], \"processing\")\n+\n+    def test_high_priority_message_is_correlated_and_scheduled(self) -\u003e None:\n+        before = datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY\n+        message = create_order_message(\n+            Order(\"order-2\", \"Grace Hopper\", \"Workstation\", 1, 2_000),\n+            high_priority_threshold=1_000,\n+        )\n+\n+        self.assertEqual(message.correlation_id, \"order-2\")\n+        self.assertEqual(message.session_id, \"Grace Hopper\")\n+        self.assertEqual(message.application_properties[\"priority\"], \"high\")\n+        self.assertGreaterEqual(message.scheduled_enqueue_time_utc, before)\n+\n+    def test_normal_priority_message_is_not_scheduled(self) -\u003e None:\n+        message = create_order_message(\n+            Order(\"order-3\", \"Grace Hopper\", \"Mouse\", 1, 50),\n+            high_priority_threshold=1_000,\n+        )\n+\n+        self.assertEqual(message.application_properties[\"priority\"], \"normal\")\n+        self.assertIsNone(message.scheduled_enqueue_time_utc)\n+\n+\n+class ProcessorTests(unittest.TestCase):\n+    @staticmethod\n+    def _processor(handler: Mock) -\u003e OrderProcessor:\n+        processor = object.__new__(OrderProcessor)\n+        processor._handler = handler\n+        return processor\n+\n+    def test_processing_failure_is_dead_lettered(self) -\u003e None:\n+        receiver = Mock()\n+        message = ServiceBusMessage(\"{invalid\")\n+        processor = self._processor(Mock())\n+\n+        processed = processor._process_message(receiver, message)\n+\n+        self.assertFalse(processed)\n+        receiver.dead_letter_message.assert_called_once()\n+        receiver.complete_message.assert_not_called()\n+\n+    def test_successful_message_is_completed(self) -\u003e None:\n+        receiver = Mock()\n+        order = Order(\"order-4\", \"Ada Lovelace\", \"Mouse\", 1, 50)\n+        message = ServiceBusMessage(order.to_json())\n+        handler = Mock()\n+        processor = self._processor(handler)\n+\n+        processed = processor._process_message(receiver, message)\n+\n+        self.assertTrue(processed)\n+        handler.assert_called_once_with(order)\n+        receiver.complete_message.assert_called_once_with(message)\n+\n+\n+class AsyncProcessorTests(unittest.TestCase):\n+    def test_processing_failure_is_dead_lettered(self) -\u003e None:\n+        async def run() -\u003e None:\n+            processor = object.__new__(AsyncOrderProcessor)\n+            processor._handler = AsyncMock()\n+            receiver = AsyncMock()\n+            message = ServiceBusMessage(\"{invalid\")\n+\n+            processed = await processor._process_message(receiver, message)\n+\n+            self.assertFalse(processed)\n+            receiver.dead_letter_message.assert_awaited_once()\n+            receiver.complete_message.assert_not_awaited()\n+\n+        asyncio.run(run())\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/processor.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/processor.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/processor.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243/order_system/processor.py
@@ -158,16 +158,23 @@
                                     message.dead_letter_reason,
                                 )
                                 if sender is not None:
-                                    order = Order.from_json(body)
-                                    sender.send_messages(
-                                        ServiceBusMessage(
-                                            order.to_json(),
-                                            content_type="application/json",
-                                            correlation_id=order.order_id,
-                                            session_id=order.customer_name,
+                                    try:
+                                        order = Order.from_json(body)
+                                        sender.send_messages(
+                                            ServiceBusMessage(
+                                                order.to_json(),
+                                                content_type="application/json",
+                                                correlation_id=order.order_id,
+                                                session_id=order.customer_name,
+                                            )
+                                        )
+                                        receiver.complete_message
... (truncated)
```

### ✅ 12. powershell — 4793ms

**Input:**

```json
{"command":"python -m unittest discover -s tests -v","description":"Run order system unit tests","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_processing_failure_is_dead_lettered (test_order_system.AsyncProcessorTests.test_processing_failure_is_dead_lettered) ... Dead-lettering invalid order message
Traceback (most recent call last):
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243\order_system\async_processor.py", line 88, in _process_message
    order = Order.from_json(message_body(message))
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243\order_system\model.py", line 47, in from_json
    data = json.loads(value)
  File "C:\ProgramData\anaconda3\Lib\json\__init__.py", line 346, in loads
    return _default_decoder.decode(s)
           ~~~~~~~~~~~~~~~~~~~~~~~^^^
  File "C:\ProgramData\anaconda3\Lib\json\decoder.py", line 345, in decode
    obj, end = self.raw_decode(s, idx=_w(s, 0).end())
               ~~~~~~~~~~~~~~~^^^^^^^^^^^^^^^^^^^^^^^
  File "C:\ProgramData\anaconda3\Lib\json\decoder.py", line 361, in raw_decode
    obj, end = self.scan_once(s, idx)
               ~~~~~~~~~~~~~~^^^^^^^^
json.decoder.JSONDecodeError: Expecting property name enclosed in double quotes: line 1 column 2 (char 1)
ok
test_high_priority_message_is_correlated_and_scheduled (test_order_system.OrderModelTests.test_high_priority_message_is_correlated_and_scheduled) ... ok
test_json_round_trip (test_order_system.OrderModelTests.test_json_round_trip) ... ok
test_normal_priority_message_is_not_scheduled (test_order_system.OrderModelTests.test_normal_priority_message_is_not_scheduled) ... ok
test_processing_failure_is_dead_lettered (test_order_system.ProcessorTests.test_processing_failure_is_dead_lettered) ... Dead-lettering invalid order message
Traceback (most recent call last):
  File "C:\Users\weidxu\AppData\Local\Temp\hyoka-service-bus-dp-python-order-processor-python-azure-skills-baseline-2714481243\order_system\processor.py", line 101, in _process_message
    order 
... (truncated)
```

### ✅ 13. powershell — 4225ms

**Input:**

```json
{"command":"python -m compileall -q main.py order_system tests","description":"Compile project sources"}
```

**Output:**

```

<shellId: 9 completed with exit code 0>
```

## Generated Files

- `README.md`
- `main.py`
- `order_system\__init__.py`
- `order_system\async_processor.py`
- `order_system\async_sender.py`
- `order_system\messages.py`
- `order_system\model.py`
- `order_system\processor.py`
- `order_system\sender.py`
- `requirements.txt`
- `tests\test_order_system.py`

## Copilot Response

I’ll scaffold the project, implement session-aware sync/async senders and processors, then run local syntax/import-safe validation without contacting Azure.

The design will use Service Bus sessions (`session_id=customer_name`) for per-customer FIFO isolation, split oversized batches on `add_message()` failure, and preserve failed messages in the DLQ unless reprocessing succeeds.

Implementation is complete; I’m now validating SDK signatures and local behavior, especially batch overflow handling, session receivers, and scheduled-message construction.

Local SDK checks passed. I’m tightening DLQ reprocessing so an unrecoverable malformed payload remains inspectable without preventing later valid dead-lettered orders from being retried.

Created the Python Azure Service Bus order-processing project with:

- JSON-serializable `Order` model and status enum
- Sync and async senders with size-aware batching, correlation IDs, priority metadata, and 30-second scheduling for orders over $1,000
- Session-based sync and async processors preserving per-customer ordering
- Explicit dead-letter handling, inspection, and safe reprocessing
- `main.py` demonstrating sync then async send/receive/DLQ cycles
- Azure SDK dependencies, unit tests, and setup documentation

The target Service Bus queue must have **sessions enabled**. Configure `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE` and optionally `SERVICE_BUS_QUEUE_NAME`, then run `python main.py`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Batch overflow and session receive timeouts receive targeted exception handling.

### Issues

- Azure SDK HttpResponseError failures can propagate without contextual logging, recovery, or user-facing handling.

## Grader Results

- order-processor.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (11/14)
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
      - Context manager pattern (`with` statements) for all clients, senders, and receivers: Fail
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
      - tool_used: any tool (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 79% | 1.00 | 0.7857 | 16.4% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.9% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.9% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.9% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.9% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 4.7857** | **59.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-order-processor --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
