# Evaluation Report: service-bus-dp-python-order-processor

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 376.0s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-order-processor` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 16/21 |
| Duration | 376.0s |
| Timestamp | 2026-08-26T08:27:39Z |
| Files Generated | 6 |
| Event Count | 13308 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 242.4s |
| Review | 131.5s |
| **Total** | **376.0s** |

## Configuration

- **name:** python-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-servicebus-py, azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=680832 out=11861 |
| Turn Count | 12 |

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

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-servicebus-py"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-identity-py"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only Python order processing demo using Azure Service Bus queues, sessions for per-customer ordering, scheduled high-priority messages, batching, dead-lettering, and DefaultAzureCredential in sync and async implementations.","learn":true}
```

### 🔧 4. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Python Azure Service Bus guidance for queue sessions, message batching with ServiceBusMessageBatch.add_message, scheduling, correlation properties, dead-letter queues, and async/sync DefaultAzureCredential.","learn":true}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py\\references\\patterns.md"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py\\references\\dead-letter.md"}
```

### 🔧 7. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 8. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

### 🔧 9. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure code-generation best practices to a local-only Python Azure Service Bus order processor.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Find official Azure Service Bus Python documentation for sessions, scheduling, batching, dead-lettering, and message properties.","parameters":{"query":"Azure Service Bus Python SDK queue sessions NEXT_AVAILABLE_SESSION ServiceBusMessageBatch add_message scheduled_enqueue_time_utc correlation_id dead_letter_message dead letter subqueue"}}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python examples for Azure Service Bus sessions, batching, scheduled messages, and DLQ handling.","parameters":{"language":"python","query":"azure-servicebus Python ServiceBusMessageBatch add_message sessions schedule_messages dead_letter_message ServiceBusSubQueue"}}
```

### 🔧 12. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 13. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the complete official Azure Service Bus Python queue guidance identified by documentation search.","parameters":{"url":"https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues"}}
```

### 🔧 14. web_search (via github-mcp-server)

**Input:**

```json
{"query":"What are the latest stable PyPI versions of azure-servicebus and azure-identity as of August 2026?"}
```

### 🔧 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: order_model.py\n+\"\"\"Order domain model and JSON serialization.\"\"\"\n+\n+from __future__ import annotations\n+\n+import json\n+from dataclasses import dataclass\n+from enum import Enum\n+from typing import Any\n+\n+\n+class OrderStatus(str, Enum):\n+    PENDING = \"pending\"\n+    PROCESSING = \"processing\"\n+    COMPLETED = \"completed\"\n+    FAILED = \"failed\"\n+\n+\n+class OrderDeserializationError(ValueError):\n+    \"\"\"Raised when a message body cannot be converted to an Order.\"\"\"\n+\n+\n+@dataclass\n+class Order:\n+    order_id: str\n+    customer_name: str\n+    product: str\n+    quantity: int\n+    total_price: float\n+    status: OrderStatus = OrderStatus.PENDING\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.order_id.strip():\n+            raise ValueError(\"order_id must not be empty\")\n+        if not self.customer_name.strip():\n+            raise ValueError(\"customer_name must not be empty\")\n+        if not self.product.strip():\n+            raise ValueError(\"product must not be empty\")\n+        if isinstance(self.quantity, bool) or self.quantity \u003c= 0:\n+            raise ValueError(\"quantity must be a positive integer\")\n+        if isinstance(self.total_price, bool) or self.total_price \u003c 0:\n+            raise ValueError(\"total_price must be non-negative\")\n+        if not isinstance(self.status, OrderStatus):\n+            self.status = OrderStatus(self.status)\n+\n+    def to_dict(self) -\u003e dict[str, Any]:\n+        return {\n+            \"order_id\": self.order_id,\n+            \"customer_name\": self.customer_name,\n+            \"product\": self.product,\n+            \"quantity\": self.quantity,\n+            \"total_price\": self.total_price,\n+            \"status\": self.status.value,\n+        }\n+\n+    def to_json(self) -\u003e str:\n+        return json.dumps(self.to_dict(), separators=(\",\", \":\"))\n+\n+    @classmethod\n+    def from_dict(cls, data: dict[str, Any]) -\u003e \"Order\":\n+        try:\n+            return cls(\n+                order_id=str(data[\"order_id\"]),\n+                customer_name=str(data[\"customer_name\"]),\n+                product=str(data[\"product\"]),\n+                quantity=int(data[\"quantity\"]),\n+                total_price=float(data[\"total_price\"]),\n+                status=OrderStatus(data[\"status\"]),\n+            )\n+        except (KeyError, TypeError, ValueError) as exc:\n+            raise OrderDeserializationError(f\"Invalid order data: {exc}\") from exc\n+\n+    @classmethod\n+    def from_json(cls, payload: str) -\u003e \"Order\":\n+        try:\n+            data = json.loads(payload)\n+        except json.JSONDecodeError as exc:\n+            raise OrderDeserializationError(f\"Invalid order JSON: {exc}\") from exc\n+\n+        if not isinstance(data, dict):\n+            raise OrderDeserializationError(\"Order JSON must contain an object\")\n+        return cls.from_dict(data)\n*** Add File: sender.py\n+\"\"\"Synchronous and asynchronous Azure Service Bus order senders.\"\"\"\n+\n+from __future__ import annotations\n+\n+from collections import defaultdict\n+from datetime import datetime, timedelta, timezone\n+from typing import Iterable\n+\n+from azure.servicebus import ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+from azure.servicebus import ServiceBusClient\n+\n+from order_model import Order\n+\n+DEFAULT_HIGH_PRIORITY_THRESHOLD = 1_000.0\n+FRAUD_REVIEW_DELAY = timedelta(seconds=30)\n+\n+\n+def _order_message(order: Order, high_priority_threshold: float) -\u003e ServiceBusMessage:\n+    high_priority = order.total_price \u003e= high_priority_threshold\n+    scheduled_time = (\n+        datetime.now(timezone.utc) + FRAUD_REVIEW_DELAY if high_priority else None\n+    )\n+    return ServiceBusMessage(\n+        order.to_json(),\n+        content_type=\"application/json\",\n+        correlation_id=order.order_id,\n+        message_id=order.order_id,\n+        session_id=order.customer_name,\n+        subject=\"high-priority-order\" if high_priority else \"order\",\n+        application_properties={\"priority\": \"high\" if high_priority else \"normal\"},\n+        scheduled_enqueue_time_utc=scheduled_time,\n+    )\n+\n+\n+def _orders_by_customer(orders: Iterable[Order]) -\u003e dict[str, list[Order]]:\n+    grouped: dict[str, list[Order]] = defaultdict(list)\n+    for order in orders:\n+        grouped[order.customer_name].append(order)\n+    return grouped\n+\n+\n+class OrderSender:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = DEFAULT_HIGH_PRIORITY_THRESHOLD,\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    def send_order(self, order: Order) -\u003e None:\n+        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            sender.send_messages(\n+                _order_message(order, self._high_priority_threshold)\n+            )\n+\n+    def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        grouped_orders = _orders_by_customer(orders)\n+        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            for customer_orders in grouped_orders.values():\n+                batch = sender.create_message_batch()\n+                batch_count = 0\n+                for order in customer_orders:\n+                    message = _order_message(order, self._high_priority_threshold)\n+                    try:\n+                        batch.add_message(message)\n+                        batch_count += 1\n+                    except ValueError:\n+                        if batch_count == 0:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum message size\"\n+                            ) from None\n+                        sender.send_messages(batch)\n+                        batch = sender.create_message_batch()\n+                        try:\n+                            batch.add_message(message)\n+                        except ValueError:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum message size\"\n+                            ) from None\n+                        batch_count = 1\n+                if batch_count:\n+                    sender.send_messages(batch)\n+\n+    def send_invalid_message_for_demo(self, customer_name: str) -\u003e None:\n+        message = ServiceBusMessage(\n+            '{\"order_id\": \"invalid\"',\n+            content_type=\"application/json\",\n+            correlation_id=\"invalid-order\",\n+            message_id=f\"invalid-{customer_name}\",\n+            session_id=customer_name,\n+        )\n+        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            sender.send_messages(message)\n+\n+\n+class AsyncOrderSender:\n+    def __init__(\n+        self,\n+        client: AsyncServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = DEFAULT_HIGH_PRIORITY_THRESHOLD,\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    async def send_order(self, order: Order) -\u003e None:\n+        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            await sender.send_messages(\n+                _order_message(order, self._high_priority_threshold)\n+            )\n+\n+    async def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        grouped_orders = _orders_by_customer(orders)\n+        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            for customer_orders in grouped_orders.values():\n+                batch = await sender.create_message_batch()\n+                batch_count = 0\n+                for order in customer_orders:\n+                    message = _order_message(order, self._high_priority_threshold)\n+                    try:\n+                        batch.add_message(message)\n+                        batch_count += 1\n+                    except ValueError:\n+                        if batch_count == 0:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum message size\"\n+                            ) from None\n+                        await sender.send_messages(batch)\n+                        batch = await sender.create_message_batch()\n+                        try:\n+                            batch.add_message(message)\n+                        except ValueError:\n+                            raise ValueError(\n+                                f\"Order {order.order_id!r} exceeds the maximum message size\"\n+                            ) from None\n+                        batch_count = 1\n+                if batch_count:\n+                    await sender.send_messages(batch)\n+\n+    async def send_invalid_message_for_demo(self, customer_name: str) -\u003e None:\n+        message = ServiceBusMessage(\n+            '{\"order_id\": \"invalid\"',\n+            content_type=\"application/json\",\n+            correlation_id=\"invalid-order\",\n+            message_id=f\"invalid-{customer_name}\",\n+            session_id=customer_name,\n+        )\n+        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            await sender.send_messages(message)\n*** Add File: processor.py\n+\"\"\"Session-aware synchronous and asynchronous order processors.\"\"\"\n+\n+from __future__ import annotations\n+\n+import logging\n+from dataclasses import replace\n+from typing import Any\n+from uuid import uuid4\n+\n+from azure.servicebus import (\n+    NEXT_AVAILABLE_SESSION,\n+    ServiceBusClient,\n+    ServiceBusMessage,\n+    ServiceBusSubQueue,\n+)\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+from azure.servicebus.exceptions import OperationTimeoutError\n+\n+from order_model import Order, OrderDeserializationError, OrderStatus\n+\n+LOGGER = logging.getLogger(__name__)\n+DEAD_LETTER_REASON = \"OrderDeserializationFailed\"\n+\n+\n+def _message_body_as_text(message: Any) -\u003e str:\n+    body = message.body\n+    if isinstance(body, str):\n+        return body\n+    if isinstance(body, (bytes, bytearray)):\n+        return bytes(body).decode(\"utf-8\")\n+    return b\"\".join(body).decode(\"utf-8\")\n+\n+\n+def _process_message_body(message: Any) -\u003e Order:\n+    order = Order.from_json(_message_body_as_text(message))\n+    processing_order = replace(order, status=OrderStatus.PROCESSING)\n+    LOGGER.info(\n+        \"Processing order %s for customer %s\",\n+        processing_order.order_id,\n+        processing_order.customer_name,\n+    )\n+    completed_order = replace(processing_order, status=OrderStatus.COMPLETED)\n+    LOGGER.info(\n+        \"Completed order %s: %s x%d, total $%.2f\",\n+        completed_order.order_id,\n+        completed_order.product,\n+        completed_order.quantity,\n+        completed_order.total_price,\n+    )\n+    return completed_order\n+\n+\n+def _retry_message(message: Any) -\u003e ServiceBusMessage:\n+    properties = dict(message.application_properties or {})\n+    properties[\"reprocessed_from_dlq\"] = True\n+    return ServiceBusMessage(\n+        _message_body_as_text(message),\n+        content_type=message.content_type or \"application/json\",\n+        correlation_id=message.correlation_id,\n+        message_id=f\"{message.message_id}-retry-{uuid4()}\",\n+        session_id=message.session_id,\n+        subject=message.subject,\n+        application_properties=properties,\n+    )\n+\n+\n+class OrderProcessor:\n+    def __init__(self, client: ServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    def process_available(\n+        self, max_wait_time: float = 5, max_messages_per_receive: int = 50\n+    ) -\u003e int:\n+        processed = 0\n+        while True:\n+            receiver = self._client.get_queue_receiver(\n+                queue_name=self._queue_name,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=max_wait_time,\n+            )\n+            try:\n+                with receiver:\n+                    while True:\n+                        messages = receiver.receive_messages(\n+                            max_message_count=max_messages_per_receive,\n+                            max_wait_time=1,\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            try:\n+                                _process_message_body(message)\n+                            except OrderDeserializationError as exc:\n+                                receiver.dead_letter_message(\n+                                    message,\n+                                    reason=DEAD_LETTER_REASON,\n+                                    error_description=str(exc)[:4096],\n+                                )\n+                                LOGGER.error(\n+                                    \"Dead-lettered message %s: %s\",\n+                                    message.message_id,\n+                                    exc,\n+                                )\n+                            else:\n+                                receiver.complete_message(message)\n+                                processed += 1\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    def process_dead_letters(\n+        self, reprocess: bool = False, max_wait_time: float = 5\n+    ) -\u003e int:\n+        handled = 0\n+        sender_context = (\n+            self._client.get_queue_sender(queue_name=self._queue_name)\n+            if reprocess\n+            else None\n+        )\n+        if sender_context is None:\n+            return self._drain_dead_letters(None, max_wait_time)\n+        with sender_context as sender:\n+            return self._drain_dead_letters(sender, max_wait_time)\n+\n+    def _drain_dead_letters(self, sender: Any, max_wait_time: float) -\u003e int:\n+        handled = 0\n+        while True:\n+            receiver = self._client.get_queue_receiver(\n+                queue_name=self._queue_name,\n+                sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=max_wait_time,\n+            )\n+            try:\n+                with receiver:\n+                    while True:\n+                        messages = receiver.receive_messages(\n+                            max_message_count=50, max_wait_time=1\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            LOGGER.warning(\n+                                \"DLQ message %s: reason=%s description=%s body=%s\",\n+                                message.message_id,\n+                                message.dead_letter_reason,\n+                                message.dead_letter_error_description,\n+                                _message_body_as_text(message),\n+                            )\n+                            if sender is not None:\n+                                sender.send_messages(_retry_message(message))\n+                            receiver.complete_message(message)\n+                            handled += 1\n+            except OperationTimeoutError:\n+                break\n+        return handled\n+\n+\n+class AsyncOrderProcessor:\n+    def __init__(self, client: AsyncServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    async def process_available(\n+        self, max_wait_time: float = 5, max_messages_per_receive: int = 50\n+    ) -\u003e int:\n+        processed = 0\n+        while True:\n+            receiver = self._client.get_queue_receiver(\n+                queue_name=self._queue_name,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=max_wait_time,\n+            )\n+            try:\n+                async with receiver:\n+                    while True:\n+                        messages = await receiver.receive_messages(\n+                            max_message_count=max_messages_per_receive,\n+                            max_wait_time=1,\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            try:\n+                                _process_message_body(message)\n+                            except OrderDeserializationError as exc:\n+                                await receiver.dead_letter_message(\n+                                    message,\n+                                    reason=DEAD_LETTER_REASON,\n+                                    error_description=str(exc)[:4096],\n+                                )\n+                                LOGGER.error(\n+                                    \"Dead-lettered message %s: %s\",\n+                                    message.message_id,\n+                                    exc,\n+                                )\n+                            else:\n+                                await receiver.complete_message(message)\n+                                processed += 1\n+            except OperationTimeoutError:\n+                break\n+        return processed\n+\n+    async def process_dead_letters(\n+        self, reprocess: bool = False, max_wait_time: float = 5\n+    ) -\u003e int:\n+        if not reprocess:\n+            return await self._drain_dead_letters(None, max_wait_time)\n+        async with self._client.get_queue_sender(\n+            queue_name=self._queue_name\n+        ) as sender:\n+            return await self._drain_dead_letters(sender, max_wait_time)\n+\n+    async def _drain_dead_letters(self, sender: Any, max_wait_time: float) -\u003e int:\n+        handled = 0\n+        while True:\n+            receiver = self._client.get_queue_receiver(\n+                queue_name=self._queue_name,\n+                sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+                session_id=NEXT_AVAILABLE_SESSION,\n+                max_wait_time=max_wait_time,\n+            )\n+            try:\n+                async with receiver:\n+                    while True:\n+                        messages = await receiver.receive_messages(\n+                            max_message_count=50, max_wait_time=1\n+                        )\n+                        if not messages:\n+                            break\n+                        for message in messages:\n+                            LOGGER.warning(\n+                                \"DLQ message %s: reason=%s description=%s body=%s\",\n+                                message.message_id,\n+                                message.dead_letter_reason,\n+                                message.dead_letter_error_description,\n+                                _message_body_as_text(message),\n+                            )\n+                            if sender is not None:\n+                                await sender.send_messages(_retry_message(message))\n+                            await receiver.complete_message(message)\n+                            handled += 1\n+            except OperationTimeoutError:\n+                break\n+        return handled\n*** Add File: main.py\n+\"\"\"Run synchronous and asynchronous Azure Service Bus order processing demos.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+import time\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from order_model import Order\n+from processor import AsyncOrderProcessor, OrderProcessor\n+from sender import AsyncOrderSender, FRAUD_REVIEW_DELAY, OrderSender\n+\n+LOGGER = logging.getLogger(__name__)\n+\n+\n+def _sample_orders(prefix: str) -\u003e list[Order]:\n+    return [\n+        Order(f\"{prefix}-001\", \"Ada Lovelace\", \"Mechanical keyboard\", 1, 149.99),\n+        Order(f\"{prefix}-002\", \"Ada Lovelace\", \"USB-C cable\", 2, 39.98),\n+        Order(f\"{prefix}-003\", \"Grace Hopper\", \"Developer workstation\", 1, 2499.00),\n+    ]\n+\n+\n+def run_sync_demo(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting synchronous demo\")\n+    orders = _sample_orders(\"sync\")\n+    with DefaultAzureCredential() as credential:\n+        with ServiceBusClient(namespace, credential) as client:\n+            sender = OrderSender(client, queue_name)\n+            processor = OrderProcessor(client, queue_name)\n+\n+            sender.send_order(orders[0])\n+            sender.send_orders(orders[1:])\n+            sender.send_invalid_message_for_demo(\"Sync Invalid Customer\")\n+\n+            processor.process_available()\n+            processor.process_dead_letters(reprocess=True)\n+            processor.process_available()\n+            processor.process_dead_letters()\n+\n+            LOGGER.info(\n+                \"Waiting %d seconds for the fraud-review scheduled order\",\n+                int(FRAUD_REVIEW_DELAY.total_seconds()),\n+            )\n+            time.sleep(FRAUD_REVIEW_DELAY.total_seconds() + 1)\n+            processor.process_available()\n+    LOGGER.info(\"Synchronous demo complete\")\n+\n+\n+async def run_async_demo(namespace: str, queue_name: str) -\u003e None:\n+    LOGGER.info(\"Starting asynchronous demo\")\n+    orders = _sample_orders(\"async\")\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncServiceBusClient(namespace, credential) as client:\n+            sender = AsyncOrderSender(client, queue_name)\n+            processor = AsyncOrderProcessor(client, queue_name)\n+\n+            await sender.send_order(orders[0])\n+            await sender.send_orders(orders[1:])\n+            await sender.send_invalid_message_for_demo(\"Async Invalid Customer\")\n+\n+            await processor.process_available()\n+            await processor.process_dead_letters(reprocess=True)\n+            await processor.process_available()\n+            await processor.process_dead_letters()\n+\n+            LOGGER.info(\n+                \"Waiting %d seconds for the fraud-review scheduled order\",\n+                int(FRAUD_REVIEW_DELAY.total_seconds()),\n+            )\n+            await asyncio.sleep(FRAUD_REVIEW_DELAY.total_seconds() + 1)\n+            await processor.process_available()\n+    LOGGER.info(\"Asynchronous demo complete\")\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    namespace = os.environ[\"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\"]\n+    queue_name = os.environ[\"SERVICEBUS_QUEUE_NAME\"]\n+    run_sync_demo(namespace, queue_name)\n+    asyncio.run(run_async_demo(namespace, queue_name))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.25.3,\u003c2.0.0\n+azure-servicebus\u003e=7.14.3,\u003c8.0.0\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+This sample sends and processes orders with both synchronous and asynchronous\n+Azure SDK clients. It uses Microsoft Entra authentication through\n+`DefaultAzureCredential`; no connection strings or secrets are stored in the\n+project.\n+\n+## Prerequisites\n+\n+- Python 3.9 or newer\n+- An existing Azure Service Bus namespace and **session-enabled queue**\n+- Azure Service Bus Data Sender and Azure Service Bus Data Receiver access\n+- A local identity supported by `DefaultAzureCredential`\n+\n+The queue must have sessions enabled because each message uses the customer name\n+as its session ID. A processor locks and drains one customer session at a time,\n+which preserves FIFO order for that customer and prevents customer sessions from\n+being interleaved by one processor.\n+\n+## Run\n+\n+Install dependencies and set:\n+\n+```text\n+SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=\u003cnamespace\u003e.servicebus.windows.net\n+SERVICEBUS_QUEUE_NAME=\u003csession-enabled-queue\u003e\n+```\n+\n+Then run `python main.py`. The demo sends normal and high-priority orders,\n+delays orders worth at least $1,000 for 30 seconds, explicitly dead-letters an\n+invalid JSON message, inspects and resubmits it, and finally inspects and removes\n+the message after it fails again.\n+\n+References:\n+\n+- https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues\n+- https://learn.microsoft.com/python/api/overview/azure/servicebus-readme\n*** End Patch\n"
```

### 🔧 16. powershell

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt \u0026\u0026 python -m compileall -q . \u0026\u0026 python -c \"import inspect; from azure.servicebus import ServiceBusMessage; from azure.servicebus.exceptions import OperationTimeoutError; from order_model import Order; from sender import _order_message; o=Order('o-1','Customer','Product',2,1200.0); r=Order.from_json(o.to_json()); m=_order_message(o,1000.0); assert r == o; assert m.correlation_id == 'o-1'; assert m.session_id == 'Customer'; assert m.scheduled_enqueue_time_utc is not None; print(inspect.signature(ServiceBusMessage)); print('local validation passed')\"","description":"Install dependencies and validate project locally","initial_wait":120,"mode":"sync"}
```

### 🔧 17. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: sender.py\n@@\n def _order_message(order: Order, high_priority_threshold: float) -\u003e ServiceBusMessage:\n-    high_priority = order.total_price \u003e= high_priority_threshold\n+    high_priority = order.total_price \u003e high_priority_threshold\n*** Update File: processor.py\n@@\n     def process_dead_letters(\n         self, reprocess: bool = False, max_wait_time: float = 5\n     ) -\u003e int:\n-        handled = 0\n         sender_context = (\n             self._client.get_queue_sender(queue_name=self._queue_name)\n             if reprocess\n*** Update File: README.md\n@@\n-Then run `python main.py`. The demo sends normal and high-priority orders,\n-delays orders worth at least $1,000 for 30 seconds, explicitly dead-letters an\n+Then run `python main.py`. The demo sends normal and high-priority orders,\n+delays orders worth more than $1,000 for 30 seconds, explicitly dead-letters an\n*** End Patch\n"
```

### 🔧 18. powershell

**Input:**

```json
{"command":"python -m compileall -q . \u0026\u0026 python -c \"import inspect; from azure.servicebus import ServiceBusClient, ServiceBusSubQueue; from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient; from order_model import Order, OrderDeserializationError; from processor import _process_message_body; from sender import _order_message; normal=Order('n-1','Ada','Cable',1,1000.0); high=Order('h-1','Grace','Workstation',1,1000.01); assert _order_message(normal,1000).scheduled_enqueue_time_utc is None; assert _order_message(high,1000).scheduled_enqueue_time_utc is not None; assert _process_message_body(type('M',(),{'body':iter([normal.to_json().encode()])})()).status.value == 'completed'; exec(\\\"try:\\\\n Order.from_json('{')\\\\n raise AssertionError('expected error')\\\\nexcept OrderDeserializationError:\\\\n pass\\\"); assert 'sub_queue' in str(inspect.signature(ServiceBusClient.get_queue_receiver)); assert 'session_id' in str(inspect.signature(AsyncServiceBusClient.get_queue_receiver)); print('behavior validation passed')\" \u0026\u0026 python -m pip check; if (Test-Path -LiteralPath '.\\__pycache__') { Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force }","description":"Validate behavior and dependency consistency","initial_wait":120,"mode":"sync"}
```

### 🔧 19. powershell

**Input:**

```json
{"command":"@'\nimport inspect\nfrom azure.servicebus import ServiceBusClient\nfrom azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\nfrom order_model import Order, OrderDeserializationError\nfrom processor import _process_message_body\nfrom sender import _order_message\n\nnormal = Order(\"n-1\", \"Ada\", \"Cable\", 1, 1000.0)\nhigh = Order(\"h-1\", \"Grace\", \"Workstation\", 1, 1000.01)\nassert _order_message(normal, 1000).scheduled_enqueue_time_utc is None\nassert _order_message(high, 1000).scheduled_enqueue_time_utc is not None\nmessage = type(\"Message\", (), {\"body\": iter([normal.to_json().encode()])})()\nassert _process_message_body(message).status.value == \"completed\"\ntry:\n    Order.from_json(\"{\")\nexcept OrderDeserializationError:\n    pass\nelse:\n    raise AssertionError(\"Malformed JSON was accepted\")\nassert \"sub_queue\" in str(inspect.signature(ServiceBusClient.get_queue_receiver))\nassert \"session_id\" in str(inspect.signature(AsyncServiceBusClient.get_queue_receiver))\nprint(\"behavior validation passed\")\n'@ | python -\nif ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }\npython -m pip check\nif ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }\nif (Test-Path -LiteralPath '.\\__pycache__') { Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force }","description":"Rerun local behavior validation","initial_wait":120,"mode":"sync"}
```

### 🔧 20. glob

**Input:**

```json
{"pattern":"**/__pycache__"}
```

## Generated Files

- `README.md`
- `main.py`
- `order_model.py`
- `processor.py`
- `requirements.txt`
- `sender.py`

## Copilot Response

I’m grounding the project in the SDK’s session, batching, scheduling, and dead-letter patterns before creating the files.

The design will use a session-enabled queue with `customer_name` as `session_id`; each receiver locks and drains one customer session before taking another, while malformed messages are explicitly dead-lettered and optionally resubmitted from the DLQ.

The files are in place; I’m validating the installed SDK surface and local model/message behavior, with no Azure connection or resource operation.

The first local behavior probe was malformed by PowerShell quoting rather than by project code; I’m rerunning it as a proper inline Python script.

Created the Python Azure Service Bus order-processing project with:

- JSON-serializable `Order` dataclass and validated statuses
- Sync/async senders with size-aware batching, correlation IDs, customer sessions, priority metadata, and 30-second fraud-review scheduling
- Sync/async session-based processors with explicit dead-lettering, DLQ inspection, and reprocessing
- Full sync-then-async demo in `main.py`
- Pinned Azure SDK dependencies and setup documentation

The existing queue must be **session-enabled**. Configure `SERVICEBUS_FULLY_QUALIFIED_NAMESPACE` and `SERVICEBUS_QUEUE_NAME`, then run `python main.py`.

References: [Azure Service Bus queues for Python](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues) and [Azure Service Bus Python SDK](https://learn.microsoft.com/python/api/overview/azure/servicebus-readme).

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- Expected deserialization failures are explicitly dead-lettered, and receive timeouts are handled to terminate draining loops.

### Issues

- Azure SDK HTTP/service failures are largely unhandled in both synchronous and asynchronous paths.

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
hyoka run --prompt-id service-bus-dp-python-order-processor --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
