# Evaluation Report: service-bus-dp-python-order-processor

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 642.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-order-processor` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 15/21 |
| Duration | 642.1s |
| Timestamp | 2026-08-27T22:22:57Z |
| Files Generated | 8 |
| Event Count | 13846 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 531.3s |
| Review | 109.9s |
| **Total** | **642.1s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Invoked | azure-messaging |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=479723 out=9877 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. skill — 4184ms

**Input:**

```json
{"skill":"azure-messaging"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Messaging SDK Troubleshooting

## Quick Reference

| Property | Value |
|----------|-------|
| **Services** | Azure Event Hubs, Azure Service Bus |
| **MCP Tools** | `mcp_azure_mcp_eventhubs`, `mcp_azure_mcp_servicebus` |
| **Best For** | Diagnosing SDK connection, auth, and message processing issues |

## When to Use This Skill

- SDK connection failures, auth errors, or AMQP link errors
- Idle timeout, connection inactivity, or slow reconnection after disconnect
- AMQP link detach or detach-forced errors
- Message lock lost, message lock expired, lock renewal failures, or batch lock timeouts
- Session lock lost, session lock expired, or session receiver errors
- Event processor or message handler stops processing
- Duplicate events or checkpoint offset resets
- SDK configuration questions (retry, prefetch, batch size, receive batch behavior)

## MCP Tools

| Tool | Command | Use |
|------|---------|-----|
| `mcp_azure_mcp_eventhubs` | Namespace/hub ops | List namespaces, hubs, consumer groups |
| `mcp_azure_mcp_servicebus` | Queue/topic ops | List namespaces, queues, topics, subscriptions |
| `mcp_azure_mcp_monitor` | `logs_query` | Query diagnostic logs with KQL |
| `mcp_azure_mcp_resourcehealth` | `get` | Check service health status |
| `mcp_azure_mcp_documentation` | Doc search | Search Microsoft Learn for troubleshooting docs |

## Diagnosis Workflow

1. **Identify the SDK and version** — Check the prompt for SDK and version clues; if not stated, proceed with diagnosis and ask later if needed
2. **Check resource health** — Use `mcp_azure_mcp_resourcehealth` to verify the namespace is healthy
3. **Review the error message** — Match against language-specific troubleshooting guide
4. **Look up documentation** — Use `mcp_azure_mcp_documentation` to search Microsoft Learn for the error or topic
5. **Check configuration** — Verify connection string, entity name, consumer group
6. **Recommend fix** — Apply remediatio
... (truncated)
```

### ❌ 2. view — 5528ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\acceptance-criteria.md"}
```

**Error:**

```
Path does not exist
```

### ❌ 3. view — 5524ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging\\examples.md"}
```

**Error:**

```
Path does not exist
```

### ✅ 4. glob — 5521ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. view — 674ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\azure-skills\\default\\.github\\plugins\\azure-skills\\skills\\azure-messaging"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/azure-skills/default/.github/plugins/azure-skills/skills/azure-messaging
@@ -1,0 +1,1 @@
 SKILL.md

```

### ❌ 6. azure-get_azure_bestpractices (via azure) — 180399ms

**Input:**

```json
{"intent":"Generate a local-only Python Azure Service Bus order processing sample using DefaultAzureCredential, sync and async senders/receivers, batching, scheduling, sessions, and dead-letter handling.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 7. azure-get_azure_bestpractices (via azure) — 182455ms

**Input:**

```json
{"command":"get","intent":"Generate a local-only Python Azure Service Bus order processing sample using DefaultAzureCredential, sync and async senders/receivers, batching, scheduling, sessions, and dead-letter handling.","parameters":{}}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ❌ 8. azure-documentation (via azure) — 182457ms

**Input:**

```json
{"intent":"Find current official Python Azure Service Bus SDK guidance for sync and async queue senders/receivers, ServiceBusMessageBatch.add_message overflow behavior, message correlation_id, scheduled enqueue time, queue sessions, dead-lettering, and receiving from the dead-letter subqueue.","learn":true}
```

**Error:**

```
MCP server 'azure': McpError: MCP error -32001: Request timed out
```

### ✅ 9. apply_patch — 833ms

**Input:**

```json
"*** Begin Patch\n*** Add File: order_processing/__init__.py\n+\"\"\"Azure Service Bus order processing sample.\"\"\"\n+\n+from .models import Order, OrderStatus\n+\n+__all__ = [\"Order\", \"OrderStatus\"]\n*** Add File: order_processing/models.py\n+from __future__ import annotations\n+\n+import json\n+from dataclasses import asdict, dataclass\n+from decimal import Decimal\n+from enum import Enum\n+from typing import Any\n+\n+\n+class OrderStatus(str, Enum):\n+    PENDING = \"pending\"\n+    PROCESSING = \"processing\"\n+    COMPLETED = \"completed\"\n+    FAILED = \"failed\"\n+\n+\n+@dataclass(slots=True)\n+class Order:\n+    order_id: str\n+    customer_name: str\n+    product: str\n+    quantity: int\n+    total_price: Decimal\n+    status: OrderStatus = OrderStatus.PENDING\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.order_id.strip():\n+            raise ValueError(\"order_id must not be empty\")\n+        if not self.customer_name.strip():\n+            raise ValueError(\"customer_name must not be empty\")\n+        if self.quantity \u003c= 0:\n+            raise ValueError(\"quantity must be greater than zero\")\n+        if self.total_price \u003c 0:\n+            raise ValueError(\"total_price must not be negative\")\n+\n+    def to_dict(self) -\u003e dict[str, Any]:\n+        data = asdict(self)\n+        data[\"total_price\"] = str(self.total_price)\n+        data[\"status\"] = self.status.value\n+        return data\n+\n+    def to_json(self) -\u003e str:\n+        return json.dumps(self.to_dict(), separators=(\",\", \":\"))\n+\n+    @classmethod\n+    def from_dict(cls, data: dict[str, Any]) -\u003e Order:\n+        return cls(\n+            order_id=str(data[\"order_id\"]),\n+            customer_name=str(data[\"customer_name\"]),\n+            product=str(data[\"product\"]),\n+            quantity=int(data[\"quantity\"]),\n+            total_price=Decimal(str(data[\"total_price\"])),\n+            status=OrderStatus(data[\"status\"]),\n+        )\n+\n+    @classmethod\n+    def from_json(cls, payload: str | bytes) -\u003e Order:\n+        data = json.loads(payload)\n+        if not isinstance(data, dict):\n+            raise ValueError(\"order JSON must contain an object\")\n+        return cls.from_dict(data)\n*** Add File: order_processing/message_factory.py\n+from __future__ import annotations\n+\n+import re\n+from datetime import datetime, timedelta, timezone\n+from decimal import Decimal\n+\n+from azure.servicebus import ServiceBusMessage\n+\n+from .models import Order\n+\n+HIGH_PRIORITY_DELAY = timedelta(seconds=30)\n+\n+\n+def customer_session_id(customer_name: str) -\u003e str:\n+    normalized = re.sub(r\"[^A-Za-z0-9._-]+\", \"-\", customer_name.strip().lower())\n+    return normalized.strip(\"-\") or \"customer\"\n+\n+\n+def create_order_message(\n+    order: Order,\n+    high_priority_threshold: Decimal,\n+) -\u003e ServiceBusMessage:\n+    is_high_priority = order.total_price \u003e high_priority_threshold\n+    message = ServiceBusMessage(\n+        order.to_json(),\n+        content_type=\"application/json\",\n+        correlation_id=order.order_id,\n+        message_id=order.order_id,\n+        session_id=customer_session_id(order.customer_name),\n+        application_properties={\n+            \"order_id\": order.order_id,\n+            \"priority\": \"high\" if is_high_priority else \"normal\",\n+        },\n+    )\n+    if is_high_priority:\n+        message.scheduled_enqueue_time_utc = (\n+            datetime.now(timezone.utc) + HIGH_PRIORITY_DELAY\n+        )\n+    return message\n*** Add File: order_processing/sender.py\n+from __future__ import annotations\n+\n+from collections import defaultdict\n+from collections.abc import Iterable\n+from decimal import Decimal\n+\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from .message_factory import create_order_message, customer_session_id\n+from .models import Order\n+\n+\n+class OrderSender:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: Decimal = Decimal(\"1000\"),\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    def send_order(self, order: Order) -\u003e None:\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            sender.send_messages(\n+                create_order_message(order, self._high_priority_threshold)\n+            )\n+\n+    def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        grouped = _group_orders_by_session(orders)\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            for session_orders in grouped.values():\n+                batch = sender.create_message_batch()\n+                for order in session_orders:\n+                    message = create_order_message(\n+                        order, self._high_priority_threshold\n+                    )\n+                    try:\n+                        batch.add_message(message)\n+                    except ValueError:\n+                        if len(batch) == 0:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum \"\n+                                \"Service Bus message batch size\"\n+                            ) from None\n+                        sender.send_messages(batch)\n+                        batch = sender.create_message_batch()\n+                        try:\n+                            batch.add_message(message)\n+                        except ValueError:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum \"\n+                                \"Service Bus message batch size\"\n+                            ) from None\n+                if len(batch):\n+                    sender.send_messages(batch)\n+\n+    def send_raw_message(self, message: ServiceBusMessage) -\u003e None:\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            sender.send_messages(message)\n+\n+\n+class AsyncOrderSender:\n+    def __init__(\n+        self,\n+        client: AsyncServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: Decimal = Decimal(\"1000\"),\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    async def send_order(self, order: Order) -\u003e None:\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            await sender.send_messages(\n+                create_order_message(order, self._high_priority_threshold)\n+            )\n+\n+    async def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        grouped = _group_orders_by_session(orders)\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            for session_orders in grouped.values():\n+                batch = await sender.create_message_batch()\n+                for order in session_orders:\n+                    message = create_order_message(\n+                        order, self._high_priority_threshold\n+                    )\n+                    try:\n+                        batch.add_message(message)\n+                    except ValueError:\n+                        if len(batch) == 0:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum \"\n+                                \"Service Bus message batch size\"\n+                            ) from None\n+                        await sender.send_messages(batch)\n+                        batch = await sender.create_message_batch()\n+                        try:\n+                            batch.add_message(message)\n+                        except ValueError:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum \"\n+                                \"Service Bus message batch size\"\n+                            ) from None\n+                if len(batch):\n+                    await sender.send_messages(batch)\n+\n+    async def send_raw_message(self, message: ServiceBusMessage) -\u003e None:\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            await sender.send_messages(message)\n+\n+\n+def _group_orders_by_session(orders: Iterable[Order]) -\u003e dict[str, list[Order]]:\n+    grouped: dict[str, list[Order]] = defaultdict(list)\n+    for order in orders:\n+        grouped[customer_session_id(order.customer_name)].append(order)\n+    return grouped\n*** Add File: order_processing/processor.py\n+from __future__ import annotations\n+\n+import logging\n+from collections.abc import Callable\n+\n+from azure.servicebus import (\n+    NEXT_AVAILABLE_SESSION,\n+    ServiceBusClient,\n+    ServiceBusMessage,\n+    ServiceBusReceiveMode,\n+    ServiceBusReceivedMessage,\n+    ServiceBusSubQueue,\n+)\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+from azure.servicebus.exceptions import OperationTimeoutError\n+\n+from .message_factory import create_order_message\n+from .models import Order, OrderStatus\n+\n+LOGGER = logging.getLogger(__name__)\n+OrderHandler = Callable[[Order], None]\n+\n+\n+def _message_body(message: ServiceBusReceivedMessage) -\u003e bytes:\n+    return b\"\".join(bytes(section) for section in message.body)\n+\n+\n+def _mark_completed(order: Order) -\u003e Order:\n+    order.status = OrderStatus.PROCESSING\n+    LOGGER.info(\n+        \"Processing order %s for %s: %s x %s\",\n+        order.order_id,\n+        order.customer_name,\n+        order.quantity,\n+        order.product,\n+    )\n+    order.status = OrderStatus.COMPLETED\n+    return order\n+\n+\n+class OrderProcessor:\n+    def __init__(self, client: ServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    def process_orders(\n+        self,\n+        max_messages: int,\n+        handler: OrderHandler | None = None,\n+        wait_time: float = 5,\n+    ) -\u003e list[Order]:\n+        processed: list[Order] = []\n+        handler = handler or (lambda order: None)\n+\n+        while len(processed) \u003c max_messages:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    self._queue_name,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,\n+                    max_wait_time=wait_time,\n+                )\n+                with receiver:\n+                    messages = receiver.receive_messages(\n+                        max_message_count=max_messages - len(processed),\n+                        max_wait_time=wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        try:\n+                            order = Order.from_json(_message_body(message))\n+                            handler(order)\n+                            _mark_completed(order)\n+                        except Exception as exc:\n+                            reason = f\"{type(exc).__name__}: {exc}\"[:4096]\n+                            LOGGER.exception(\n+                                \"Order message %s failed; dead-lettering it\",\n+                                message.message_id,\n+                            )\n+                            receiver.dead_letter_message(\n+                                message,\n+                                reason=\"OrderProcessingFailed\",\n+                                error_description=reason,\n+                            )\n+                        else:\n+                            receiver.complete_message(message)\n+                            processed.append(order)\n+                            LOGGER.info(\"Completed order %s\", order.order_id)\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    def inspect_dead_letters(\n+        self,\n+        max_messages: int = 20,\n+        reprocess: bool = False,\n+        wait_time: float = 5,\n+    ) -\u003e list[ServiceBusReceivedMessage]:\n+        inspected: list[ServiceBusReceivedMessage] = []\n+        while len(inspected) \u003c max_messages:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    self._queue_name,\n+                    sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,\n+                    max_wait_time=wait_time,\n+                )\n+                with receiver:\n+                    messages = receiver.receive_messages(\n+                        max_message_count=max_messages - len(inspected),\n+                        max_wait_time=wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        inspected.append(message)\n+                        LOGGER.info(\n+                            \"Dead letter %s: reason=%s description=%s\",\n+                            message.message_id,\n+                            message.dead_letter_reason,\n+                            message.dead_letter_error_description,\n+                        )\n+                        if reprocess:\n+                            self._reprocess_message(message)\n+                            receiver.complete_message(message)\n+                        else:\n+                            receiver.abandon_message(message)\n+            except OperationTimeoutError:\n+                break\n+        return inspected\n+\n+    def _reprocess_message(self, message: ServiceBusReceivedMessage) -\u003e None:\n+        order = Order.from_json(_message_body(message))\n+        with self._client.get_queue_sender(self._queue_name) as sender:\n+            sender.send_messages(create_order_message(order, high_priority_threshold=order.total_price))\n+\n+\n+class AsyncOrderProcessor:\n+    def __init__(self, client: AsyncServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    async def process_orders(\n+        self,\n+        max_messages: int,\n+        handler: OrderHandler | None = None,\n+        wait_time: float = 5,\n+    ) -\u003e list[Order]:\n+        processed: list[Order] = []\n+        handler = handler or (lambda order: None)\n+\n+        while len(processed) \u003c max_messages:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    self._queue_name,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,\n+                    max_wait_time=wait_time,\n+                )\n+                async with receiver:\n+                    messages = await receiver.receive_messages(\n+                        max_message_count=max_messages - len(processed),\n+                        max_wait_time=wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        try:\n+                            order = Order.from_json(_message_body(message))\n+                            handler(order)\n+                            _mark_completed(order)\n+                        except Exception as exc:\n+                            reason = f\"{type(exc).__name__}: {exc}\"[:4096]\n+                            LOGGER.exception(\n+                                \"Order message %s failed; dead-lettering it\",\n+                                message.message_id,\n+                            )\n+                            await receiver.dead_letter_message(\n+                                message,\n+                                reason=\"OrderProcessingFailed\",\n+                                error_description=reason,\n+                            )\n+                        else:\n+                            await receiver.complete_message(message)\n+                            processed.append(order)\n+                            LOGGER.info(\"Completed order %s\", order.order_id)\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    async def inspect_dead_letters(\n+        self,\n+        max_messages: int = 20,\n+        reprocess: bool = False,\n+        wait_time: float = 5,\n+    ) -\u003e list[ServiceBusReceivedMessage]:\n+        inspected: list[ServiceBusReceivedMessage] = []\n+        while len(inspected) \u003c max_messages:\n+            try:\n+                receiver = self._client.get_queue_receiver(\n+                    self._queue_name,\n+                    sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                    session_id=NEXT_AVAILABLE_SESSION,\n+                    receive_mode=ServiceBusReceiveMode.PEEK_LOCK,\n+                    max_wait_time=wait_time,\n+                )\n+                async with receiver:\n+                    messages = await receiver.receive_messages(\n+                        max_message_count=max_messages - len(inspected),\n+                        max_wait_time=wait_time,\n+                    )\n+                    if not messages:\n+                        break\n+                    for message in messages:\n+                        inspected.append(message)\n+                        LOGGER.info(\n+                            \"Dead letter %s: reason=%s description=%s\",\n+                            message.message_id,\n+                            message.dead_letter_reason,\n+                            message.dead_letter_error_description,\n+                        )\n+                        if reprocess:\n+                            await self._reprocess_message(message)\n+                            await receiver.complete_message(message)\n+                        else:\n+                            await receiver.abandon_message(message)\n+            except OperationTimeoutError:\n+                break\n+        return inspected\n+\n+    async def _reprocess_message(\n+        self, message: ServiceBusReceivedMessage\n+    ) -\u003e None:\n+        order = Order.from_json(_message_body(message))\n+        async with self._client.get_queue_sender(self._queue_name) as sender:\n+            await sender.send_messages(\n+                create_order_message(order, high_priority_threshold=order.total_price)\n+            )\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+from decimal import Decimal\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from order_processing.message_factory import customer_session_id\n+from order_processing.models import Order\n+from order_processing.processor import AsyncOrderProcessor, OrderProcessor\n+from order_processing.sender import AsyncOrderSender, OrderSender\n+\n+LOGGER = logging.getLogger(__name__)\n+\n+\n+def sample_orders(prefix: str) -\u003e list[Order]:\n+    return [\n+        Order(\n+            order_id=f\"{prefix}-001\",\n+            customer_name=\"Ada Lovelace\",\n+            product=\"Mechanical Keyboard\",\n+            quantity=1,\n+            total_price=Decimal(\"149.99\"),\n+        ),\n+        Order(\n+            order_id=f\"{prefix}-002\",\n+            customer_name=\"Ada Lovelace\",\n+            product=\"4K Monitor\",\n+            quantity=2,\n+            total_price=Decimal(\"1199.98\"),\n+        ),\n+        Order(\n+            order_id=f\"{prefix}-003\",\n+            customer_name=\"Grace Hopper\",\n+            product=\"USB-C Dock\",\n+            quantity=1,\n+            total_price=Decimal(\"229.00\"),\n+        ),\n+    ]\n+\n+\n+def malformed_message(prefix: str) -\u003e ServiceBusMessage:\n+    message_id = f\"{prefix}-invalid\"\n+    return ServiceBusMessage(\n+        '{\"order_id\":',\n+        content_type=\"application/json\",\n+        correlation_id=message_id,\n+        message_id=message_id,\n+        session_id=customer_session_id(\"Invalid Demo Customer\"),\n+    )\n+\n+\n+def run_sync(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting synchronous demo\")\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(namespace, credential) as client:\n+            sender = OrderSender(client, queue_name)\n+            processor = OrderProcessor(client, queue_name)\n+            sender.send_orders(sample_orders(\"sync\"))\n+            sender.send_raw_message(malformed_message(\"sync\"))\n+            processor.process_orders(max_messages=4, wait_time=35)\n+            processor.inspect_dead_letters(max_messages=10)\n+\n+\n+async def run_async(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting asynchronous demo\")\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncServiceBusClient(namespace, credential) as client:\n+            sender = AsyncOrderSender(client, queue_name)\n+            processor = AsyncOrderProcessor(client, queue_name)\n+            await sender.send_orders(sample_orders(\"async\"))\n+            await sender.send_raw_message(malformed_message(\"async\"))\n+            await processor.process_orders(max_messages=4, wait_time=35)\n+            await processor.inspect_dead_letters(max_messages=10)\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    namespace = os.environ[\"SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.environ[\"SERVICE_BUS_QUEUE_NAME\"]\n+    run_sync(namespace, queue_name)\n+    asyncio.run(run_async(namespace, queue_name))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-servicebus\u003e=7.12.0,\u003c8.0.0\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+This sample publishes and processes orders with both synchronous and\n+asynchronous Azure Service Bus clients. It uses Service Bus sessions keyed by\n+customer to preserve per-customer ordering, message batching, scheduled\n+high-value orders, and explicit dead-letter handling.\n+\n+## Prerequisites\n+\n+- Python 3.10 or later\n+- An Azure Service Bus namespace and a **session-enabled queue**\n+- A signed-in identity supported by `DefaultAzureCredential` with Azure Service\n+  Bus Data Sender and Data Receiver permissions\n+\n+Set these environment variables:\n+\n+```text\n+SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE=\u003cnamespace\u003e.servicebus.windows.net\n+SERVICE_BUS_QUEUE_NAME=\u003csession-enabled-queue\u003e\n+```\n+\n+Install dependencies with `python -m pip install -r requirements.txt`, then run\n+`python main.py`.\n+\n+Orders over $1,000 carry a `priority=high` application property and are\n+scheduled 30 seconds into the future. Azure Service Bus queues do not provide a\n+native priority field, so consumers can use this property when applying their\n+own prioritization policy.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/__init__.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/__init__.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/__init__.py
@@ -1,0 +1,6 @@
+"""Azure Service Bus order processing sample."""
+
+from .models import Order, OrderStatus
+
+__all__ = ["Order", "OrderStatus"]
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/models.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/models.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/models.py
@@ -1,0 +1,62 @@
+from __future__ import annotations
+
+import json
+from dataclasses import asdict, dataclass
+from decimal import Decimal
+from enum import Enum
+from typing import Any
+
+
+class OrderStatus(str, Enum):
+    PENDING = "pending"
+    PROCESSING = "processing"
+    COMPLETED = "completed"
+    FAILED = "failed"
+
+
+@dataclass(slots=True)
+class Order:
+    order_id: str
+    customer_name: str
+    product: str
+    quantity: int
+    total_price: Decimal
+    status: OrderStatus = OrderStatus.PENDING
+
+    def __post_init__(self) -> None:
+        if not self.order_id.strip():
+            raise ValueError("order_id must not be empty")
+        if not self.customer_name.strip():
+            raise ValueError("customer_name must not be
... (truncated)
```

### ✅ 10. powershell — 9410ms

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt \u0026\u0026 python -m compileall -q .","description":"Install dependencies and compile project","initial_wait":180,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 11. powershell — 2195ms

**Input:**

```json
{"command":"@'\nimport asyncio\nimport inspect\nfrom decimal import Decimal\nfrom azure.servicebus import ServiceBusMessage, ServiceBusSubQueue\nfrom azure.servicebus.aio import ServiceBusSender as AsyncSender\nfrom order_processing.models import Order\nfrom order_processing.message_factory import create_order_message\n\norder = Order('o1', 'Ada Lovelace', 'Keyboard', 1, Decimal('1001.00'))\nassert Order.from_json(order.to_json()) == order\nmessage = create_order_message(order, Decimal('1000'))\nassert message.correlation_id == 'o1'\nassert message.session_id == 'ada-lovelace'\nassert message.application_properties[b'priority'] == b'high'\nassert message.scheduled_enqueue_time_utc is not None\nassert inspect.iscoroutinefunction(AsyncSender.create_message_batch)\nassert ServiceBusSubQueue.DEAD_LETTER.value == 'deadletter'\nprint('SDK smoke checks passed')\n'@ | python -","description":"Run local SDK smoke checks","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Traceback (most recent call last):
  File "<stdin>", line 14, in <module>
KeyError: b'priority'
<shellId: 1 completed with exit code 1>
```

### ✅ 12. powershell — 2171ms

**Input:**

```json
{"command":"@'\nimport inspect\nfrom decimal import Decimal\nfrom azure.servicebus import ServiceBusReceiver, ServiceBusSender, ServiceBusSubQueue\nfrom azure.servicebus.aio import ServiceBusReceiver as AsyncReceiver\nfrom azure.servicebus.aio import ServiceBusSender as AsyncSender\nfrom order_processing.models import Order\nfrom order_processing.message_factory import create_order_message\n\norder = Order('o1', 'Ada Lovelace', 'Keyboard', 1, Decimal('1001.00'))\nassert Order.from_json(order.to_json()) == order\nmessage = create_order_message(order, Decimal('1000'))\nassert message.correlation_id == 'o1'\nassert message.session_id == 'ada-lovelace'\nassert message.application_properties['priority'] == 'high'\nassert message.scheduled_enqueue_time_utc is not None\nassert inspect.iscoroutinefunction(AsyncSender.create_message_batch)\nassert 'reason' in inspect.signature(ServiceBusReceiver.dead_letter_message).parameters\nassert 'reason' in inspect.signature(AsyncReceiver.dead_letter_message).parameters\nassert ServiceBusSubQueue.DEAD_LETTER.value == 'deadletter'\nprint('SDK smoke checks passed')\n'@ | python -","description":"Repeat local SDK smoke checks","initial_wait":30,"mode":"sync"}
```

**Output:**

```
SDK smoke checks passed
<shellId: 2 completed with exit code 0>
```

### ✅ 13. apply_patch — 955ms

**Input:**

```json
"*** Begin Patch\n*** Update File: order_processing/processor.py\n@@\n import logging\n from collections.abc import Callable\n+from decimal import Decimal\n@@\n OrderHandler = Callable[[Order], None]\n+DeadLetterRepair = Callable[[ServiceBusReceivedMessage], Order]\n@@\n     def inspect_dead_letters(\n         self,\n         max_messages: int = 20,\n         reprocess: bool = False,\n+        repair: DeadLetterRepair | None = None,\n         wait_time: float = 5,\n     ) -\u003e list[ServiceBusReceivedMessage]:\n@@\n                         )\n                         if reprocess:\n-                            self._reprocess_message(message)\n+                            self._reprocess_message(message, repair)\n                             receiver.complete_message(message)\n                         else:\n                             receiver.abandon_message(message)\n+                    if not reprocess:\n+                        # Abandoned messages remain in this session for later inspection.\n+                        return inspected\n             except OperationTimeoutError:\n                 break\n         return inspected\n \n-    def _reprocess_message(self, message: ServiceBusReceivedMessage) -\u003e None:\n-        order = Order.from_json(_message_body(message))\n+    def _reprocess_message(\n+        self,\n+        message: ServiceBusReceivedMessage,\n+        repair: DeadLetterRepair | None,\n+    ) -\u003e None:\n+        order = repair(message) if repair else Order.from_json(_message_body(message))\n         with self._client.get_queue_sender(self._queue_name) as sender:\n-            sender.send_messages(create_order_message(order, high_priority_threshold=order.total_price))\n+            sender.send_messages(\n+                create_order_message(order, high_priority_threshold=Decimal(\"Infinity\"))\n+            )\n@@\n     async def inspect_dead_letters(\n         self,\n         max_messages: int = 20,\n         reprocess: bool = False,\n+        repair: DeadLetterRepair | None = None,\n         wait_time: float = 5,\n     ) -\u003e list[ServiceBusReceivedMessage]:\n@@\n                         )\n                         if reprocess:\n-                            await self._reprocess_message(message)\n+                            await self._reprocess_message(message, repair)\n                             await receiver.complete_message(message)\n                         else:\n                             await receiver.abandon_message(message)\n+                    if not reprocess:\n+                        # Abandoned messages remain in this session for later inspection.\n+                        return inspected\n             except OperationTimeoutError:\n                 break\n         return inspected\n \n     async def _reprocess_message(\n-        self, message: ServiceBusReceivedMessage\n+        self,\n+        message: ServiceBusReceivedMessage,\n+        repair: DeadLetterRepair | None,\n     ) -\u003e None:\n-        order = Order.from_json(_message_body(message))\n+        order = repair(message) if repair else Order.from_json(_message_body(message))\n         async with self._client.get_queue_sender(self._queue_name) as sender:\n             await sender.send_messages(\n-                create_order_message(order, high_priority_threshold=order.total_price)\n+                create_order_message(order, high_priority_threshold=Decimal(\"Infinity\"))\n             )\n*** Update File: main.py\n@@\n def malformed_message(prefix: str) -\u003e ServiceBusMessage:\n@@\n     )\n \n \n+def repaired_demo_order(prefix: str) -\u003e Order:\n+    return Order(\n+        order_id=f\"{prefix}-invalid-repaired\",\n+        customer_name=\"Invalid Demo Customer\",\n+        product=\"Replacement Cable\",\n+        quantity=1,\n+        total_price=Decimal(\"19.99\"),\n+    )\n+\n+\n def run_sync(namespace: str, queue_name: str) -\u003e None:\n@@\n             sender.send_orders(sample_orders(\"sync\"))\n             sender.send_raw_message(malformed_message(\"sync\"))\n             processor.process_orders(max_messages=4, wait_time=35)\n-            processor.inspect_dead_letters(max_messages=10)\n+            processor.inspect_dead_letters(\n+                max_messages=10,\n+                reprocess=True,\n+                repair=lambda _: repaired_demo_order(\"sync\"),\n+            )\n+            processor.process_orders(max_messages=1)\n@@\n             await sender.send_orders(sample_orders(\"async\"))\n             await sender.send_raw_message(malformed_message(\"async\"))\n             await processor.process_orders(max_messages=4, wait_time=35)\n-            await processor.inspect_dead_letters(max_messages=10)\n+            await processor.inspect_dead_letters(\n+                max_messages=10,\n+                reprocess=True,\n+                repair=lambda _: repaired_demo_order(\"async\"),\n+            )\n+            await processor.process_orders(max_messages=1)\n*** Update File: README.md\n@@\n native priority field, so consumers can use this property when applying their\n-own prioritization policy.\n+own prioritization policy. The demo also sends malformed JSON, dead-letters it,\n+repairs and republishes it through the dead-letter inspection API, and then\n+processes the repaired order.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/processor.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/processor.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/processor.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-547699823/order_processing/processor.py
@@ -2,6 +2,7 @@
 
 import logging
 from collections.abc import Callable
+from decimal import Decimal
 
 from azure.servicebus import (
     NEXT_AVAILABLE_SESSION,
@@ -19,6 +20,7 @@
 
 LOGGER = logging.getLogger(__name__)
 OrderHandler = Callable[[Order], None]
+DeadLetterRepair = Callable[[ServiceBusReceivedMessage], Order]
 
 
 def _message_body(message: ServiceBusReceivedMessage) -> bytes:
@@ -95,6 +97,7 @@
         self,
         max_messages: int = 20,
         reprocess: bool = False,
+        repair: DeadLetterRepair | None = None,
         wait_time: float = 5,
     ) -> list[ServiceBusReceivedMessage]:
         inspected: list[ServiceBusReceivedMessage] = []
@@ -123,18 +126,27 @@
                             message.dead_letter_error_description,
                         )
                         if reprocess:
-                            self._reprocess_message(message)
+                            self._reprocess_message(message, repair)
                             receiver.complete_message(message)
                         else:
                             receiver.abandon_message(message)
+                    if not reprocess:
+                        # Abandoned messages remain in this session for later inspection.
+                        return inspected
             except Operatio
... (truncated)
```

### ✅ 14. powershell — 3703ms

**Input:**

```json
{"command":"@'\nimport inspect\nfrom azure.identity import DefaultAzureCredential\nfrom azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\nfrom azure.servicebus import ServiceBusMessageBatch\nfrom order_processing.models import Order, OrderStatus\nfrom main import repaired_demo_order, sample_orders\n\nassert hasattr(DefaultAzureCredential, '__enter__')\nassert hasattr(AsyncDefaultAzureCredential, '__aenter__')\nassert inspect.signature(ServiceBusMessageBatch.add_message)\nassert repaired_demo_order('x').status is OrderStatus.PENDING\nassert len(sample_orders('x')) == 3\nprint('Project validation passed')\n'@ | python - \u0026\u0026 python -m compileall -q main.py order_processing","description":"Validate project imports and SDK APIs","initial_wait":30,"mode":"sync"}
```

**Output:**

```
Project validation passed
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `README.md`
- `main.py`
- `order_processing\__init__.py`
- `order_processing\message_factory.py`
- `order_processing\models.py`
- `order_processing\processor.py`
- `order_processing\sender.py`
- `requirements.txt`

## Copilot Response

The SDK references are unavailable in this environment, so I’m implementing against the current `azure-servicebus` Python API with session-based ordering and bounded demo loops.

The local SDK checks passed. I’m tightening the dead-letter cycle so the demo repairs, republishes, and successfully processes a failed message instead of only viewing it.

Created the Python Azure Service Bus order-processing project with:

- JSON-serializable `Order` dataclass and status enum
- Synchronous and asynchronous senders with size-aware batching
- Correlation IDs, high-priority metadata, and 30-second scheduling
- Session-based per-customer sequential processing
- Explicit dead-letter handling, inspection, repair, and reprocessing
- Full sync-then-async demo in `main.py`
- Azure SDK dependencies in `requirements.txt`
- Setup and session-enabled queue requirements in `README.md`

The demo uses `SERVICE_BUS_FULLY_QUALIFIED_NAMESPACE` and `SERVICE_BUS_QUEUE_NAME` with `DefaultAzureCredential`.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Processing failures are logged and malformed messages are explicitly dead-lettered with a reason and error description.

### Issues

- Missing explicit handling for HttpResponseError and relevant Azure Service Bus exceptions across sync and async operations.

## Grader Results

- order-processor.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (10/14)
      - Sender uses `ServiceBusClient.get_queue_sender()` (or async equivalent): Pass
      - Processor uses `ServiceBusClient.get_queue_receiver()` or session-enabled receiver: Pass
      - Batch sending: creates `ServiceBusMessageBatch`, checks `add_message()` return or catches `MessageSizeExceededError`: Fail
      - Handles the case where a message doesn't fit in the current batch: Fail
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
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 71% | 1.00 | 0.7143 | 12.5% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.5% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.7143** | **71.4%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-order-processor --config python-azure-skills/azure-skill-mcp
```

---

[← Back to Summary](../../../../../../summary.md)
