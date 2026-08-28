# Evaluation Report: service-bus-dp-python-order-processor

**Config:** python-azure-skills/azure-skill-mcp-microsoft-skill | **Result:** ❌ FAILED | **Duration:** 331.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `service-bus-dp-python-order-processor` |
| Config | python-azure-skills/azure-skill-mcp-microsoft-skill |
| Result | ❌ FAILED |
| Score | 16/21 |
| Duration | 331.8s |
| Timestamp | 2026-08-28T21:25:38Z |
| Files Generated | 9 |
| Event Count | 14986 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 216.4s |
| Review | 114.2s |
| **Total** | **331.8s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp-microsoft-skill
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-servicebus-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=399997 out=11968 |
| Turn Count | 8 |

## Tool Calls

### ✅ 1. skill — 3369ms

**Input:**

```json
{"skill":"azure-servicebus-py"}
```

**Output:**

```
Skill loaded successfully ✅

# Azure Service Bus SDK for Python

Enterprise messaging for reliable cloud communication with queues and pub/sub topics.

## Installation

```bash
pip install azure-servicebus azure-identity
```

## Environment Variables

```bash
SERVICEBUS_FULLY_QUALIFIED_NAMESPACE=<namespace>.servicebus.windows.net  # Required for all auth methods
SERVICEBUS_QUEUE_NAME=myqueue  # Required for queue operations
SERVICEBUS_TOPIC_NAME=mytopic  # Required for topic operations
SERVICEBUS_SUBSCRIPTION_NAME=mysubscription  # Required for subscription operations
AZURE_TOKEN_CREDENTIALS=prod # Required only if DefaultAzureCredential is used in production
```

## Authentication & Lifecycle

> **🔑 Two rules apply to every code sample below:**
>
> 1. **Prefer `DefaultAzureCredential`.** It works locally (Azure CLI / VS Code / Developer CLI) and in Azure (managed identity, workload identity) with no code change. Avoid connection strings, account/API keys — they bypass Entra audit and rotation.
>    - Local dev: `DefaultAzureCredential` works as-is.
>    - Production: set `AZURE_TOKEN_CREDENTIALS=prod` (or `AZURE_TOKEN_CREDENTIALS=<specific_credential>`) to constrain the credential chain to production-safe credentials.
> 2. **Wrap every client in a context manager** so HTTP transports, sockets, and token caches are released deterministically:
>    - Sync: `with <Client>(...) as client:`
>    - Async: `async with <Client>(...) as client:` **and** `async with DefaultAzureCredential() as credential:` (from `azure.identity.aio`)
>
> Snippets may abbreviate this setup, but production code should always follow both rules.

```python
from azure.identity import DefaultAzureCredential, ManagedIdentityCredential
from azure.servicebus import ServiceBusClient

# Local dev: DefaultAzureCredential. Production: set AZURE_TOKEN_CREDENTIALS=prod or AZURE_TOKEN_CREDENTIALS=<specific_credential>
credential = DefaultAzureCredential(require_envvar=True)
# Or use a specific creden
... (truncated)
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 9599ms

**Input:**

```json
{"command":"get_azure_bestpractices","intent":"Generate a local-only Python order processing project using Azure Service Bus SDK, DefaultAzureCredential, sync and async clients, sessions for per-customer ordering, size-aware batches, scheduled high-priority messages, and dead-letter reprocessing."}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. view — 10267ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py\\references\\patterns.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/dead-letter.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/dead-letter.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/dead-letter.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/dead-letter.md
@@ -1,419 +1,419 @@
 # Dead-Letter Queue Reference
 
 Handling poison messages and dead-letter queue processing in Azure Service Bus.
 
 ## Dead-Letter Queue Overview
 
 The dead-letter queue (DLQ) is a secondary sub-queue for messages that cannot be processed:
 
 ```
 Main Queue: myqueue
    └── Dead-Letter Queue: myqueue/$deadletterqueue
 ```
 
 ## Why Messages Get Dead-Lettered
 
 | Reason | Description |
 |--------|-------------|
 | `MaxDeliveryCountExceeded` | Message abandoned too many times |
 | `HeaderSizeExceeded` | Message headers too large |
 | `TTLExpiration` | Message expired before delivery |
 | `SessionIdMismatch` | Session ID doesn't match |
 | `MessageSizeExceeded` | Message body too large |
 | Custom reason | Explicitly dead-lettered by application |
 
 ## Receiving from Dead-Letter Queue
 
 ```python
 from azure.servicebus import ServiceBusSubQueue
 from azure.servicebus.aio import ServiceBusClient
 from azure.identity.aio import DefaultAzureCredential
 
 async def receive_dead_letters(namespace: str, queue_name: str):
     """Receive messages from dead-letter queue."""
     credential = DefaultAzureCredential()
     
     async with ServiceBusClient(
         fully_qualified_namespace=namespace,
         credential=credential
     ) as client:
         # Get dead-letter queue receiver
         dlq_receiver = cl
... (truncated)
```

### ✅ 4. view — 9929ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py\\references\\dead-letter.md"}
```

**Output:**

```

diff --git a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
index 0000000..0000000 100644
--- a/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
+++ b/C:/Users/weidxu/.hyoka/cache/repos/microsoft/skills/default/.github/plugins/azure-sdk-python/skills/azure-servicebus-py/references/patterns.md
@@ -1,406 +1,406 @@
 # Messaging Patterns Reference
 
 Advanced messaging patterns for Azure Service Bus.
 
 ## Competing Consumers
 
 Multiple receivers processing messages from the same queue in parallel:
 
 ```python
 import asyncio
 from azure.servicebus.aio import ServiceBusClient
 from azure.identity.aio import DefaultAzureCredential
 
 async def worker(worker_id: int, namespace: str, queue_name: str):
     """Worker that processes messages from a shared queue."""
     credential = DefaultAzureCredential()
     
     async with ServiceBusClient(
         fully_qualified_namespace=namespace,
         credential=credential
     ) as client:
         receiver = client.get_queue_receiver(queue_name=queue_name)
         
         async with receiver:
             while True:
                 messages = await receiver.receive_messages(
                     max_message_count=10,
                     max_wait_time=5
                 )
                 
                 if not messages:
                     continue
                 
                 for msg in messages:
                     try:
                         print(f"Worker {worker_id}: Processing {str(msg)}")
                         await process_message(msg)
                         await receiver.complete_message(msg)
                     except Exception as e:
                        
... (truncated)
```

### ✅ 5. glob — 9258ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 6. glob — 8926ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-servicebus-py","pattern":"**/{acceptance-criteria.md,examples.md}"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 7. azure-get_azure_bestpractices (via azure) — 676ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure code-generation guidance for a Python Azure Service Bus project.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 8. apply_patch — 896ms

**Input:**

```json
"*** Begin Patch\n*** Add File: order.py\n+from __future__ import annotations\n+\n+import json\n+from dataclasses import asdict, dataclass\n+from enum import Enum\n+from typing import Any\n+\n+\n+class OrderStatus(str, Enum):\n+    PENDING = \"pending\"\n+    PROCESSING = \"processing\"\n+    COMPLETED = \"completed\"\n+    FAILED = \"failed\"\n+\n+\n+@dataclass(slots=True)\n+class Order:\n+    order_id: str\n+    customer_name: str\n+    product: str\n+    quantity: int\n+    total_price: float\n+    status: OrderStatus = OrderStatus.PENDING\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.order_id.strip():\n+            raise ValueError(\"order_id must not be empty\")\n+        if not self.customer_name.strip():\n+            raise ValueError(\"customer_name must not be empty\")\n+        if not self.product.strip():\n+            raise ValueError(\"product must not be empty\")\n+        if isinstance(self.quantity, bool) or not isinstance(self.quantity, int):\n+            raise TypeError(\"quantity must be an integer\")\n+        if self.quantity \u003c= 0:\n+            raise ValueError(\"quantity must be greater than zero\")\n+        if isinstance(self.total_price, bool) or not isinstance(\n+            self.total_price, (int, float)\n+        ):\n+            raise TypeError(\"total_price must be numeric\")\n+        if self.total_price \u003c 0:\n+            raise ValueError(\"total_price must not be negative\")\n+        if not isinstance(self.status, OrderStatus):\n+            self.status = OrderStatus(self.status)\n+        self.total_price = float(self.total_price)\n+\n+    def to_json(self) -\u003e str:\n+        payload = asdict(self)\n+        payload[\"status\"] = self.status.value\n+        return json.dumps(payload, separators=(\",\", \":\"))\n+\n+    @classmethod\n+    def from_json(cls, value: str | bytes) -\u003e Order:\n+        payload: Any = json.loads(value)\n+        if not isinstance(payload, dict):\n+            raise ValueError(\"order JSON must contain an object\")\n+\n+        required_fields = {\n+            \"order_id\",\n+            \"customer_name\",\n+            \"product\",\n+            \"quantity\",\n+            \"total_price\",\n+            \"status\",\n+        }\n+        missing_fields = required_fields.difference(payload)\n+        if missing_fields:\n+            missing = \", \".join(sorted(missing_fields))\n+            raise ValueError(f\"order JSON is missing fields: {missing}\")\n+\n+        return cls(\n+            order_id=payload[\"order_id\"],\n+            customer_name=payload[\"customer_name\"],\n+            product=payload[\"product\"],\n+            quantity=payload[\"quantity\"],\n+            total_price=payload[\"total_price\"],\n+            status=OrderStatus(payload[\"status\"]),\n+        )\n*** Add File: messages.py\n+from __future__ import annotations\n+\n+from datetime import datetime, timedelta, timezone\n+from uuid import uuid4\n+\n+from azure.servicebus import ServiceBusMessage\n+\n+from order import Order\n+\n+\n+def create_order_message(\n+    order: Order,\n+    high_priority_threshold: float,\n+) -\u003e ServiceBusMessage:\n+    is_high_priority = order.total_price \u003e high_priority_threshold\n+    scheduled_time = (\n+        datetime.now(timezone.utc) + timedelta(seconds=30)\n+        if is_high_priority\n+        else None\n+    )\n+\n+    return ServiceBusMessage(\n+        order.to_json(),\n+        content_type=\"application/json\",\n+        correlation_id=order.order_id,\n+        message_id=str(uuid4()),\n+        session_id=order.customer_name,\n+        scheduled_enqueue_time_utc=scheduled_time,\n+        application_properties={\n+            \"order_id\": order.order_id,\n+            \"priority\": \"high\" if is_high_priority else \"normal\",\n+        },\n+    )\n+\n+\n+def create_reprocessed_message(order: Order, dead_letter_reason: str) -\u003e ServiceBusMessage:\n+    return ServiceBusMessage(\n+        order.to_json(),\n+        content_type=\"application/json\",\n+        correlation_id=order.order_id,\n+        message_id=str(uuid4()),\n+        session_id=order.customer_name,\n+        application_properties={\n+            \"order_id\": order.order_id,\n+            \"priority\": \"normal\",\n+            \"reprocessed_from_dead_letter\": True,\n+            \"original_dead_letter_reason\": dead_letter_reason,\n+        },\n+    )\n*** Add File: sync_sender.py\n+from __future__ import annotations\n+\n+import logging\n+from collections import defaultdict\n+from collections.abc import Iterable\n+\n+from azure.servicebus import ServiceBusClient, ServiceBusMessageBatch\n+\n+from messages import create_order_message\n+from order import Order\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+class SyncOrderSender:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    def send_order(self, order: Order) -\u003e None:\n+        message = create_order_message(order, self._high_priority_threshold)\n+        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            sender.send_messages(message)\n+        logger.info(\"Sent order %s for %s\", order.order_id, order.customer_name)\n+\n+    def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        orders_by_customer: dict[str, list[Order]] = defaultdict(list)\n+        for order in orders:\n+            orders_by_customer[order.customer_name].append(order)\n+\n+        with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            for customer_orders in orders_by_customer.values():\n+                batch = sender.create_message_batch()\n+                for order in customer_orders:\n+                    message = create_order_message(\n+                        order,\n+                        self._high_priority_threshold,\n+                    )\n+                    batch = self._add_or_flush(sender, batch, message)\n+                if len(batch):\n+                    sender.send_messages(batch)\n+\n+        logger.info(\"Sent %d customer order group(s)\", len(orders_by_customer))\n+\n+    @staticmethod\n+    def _add_or_flush(sender, batch: ServiceBusMessageBatch, message):\n+        try:\n+            batch.add_message(message)\n+            return batch\n+        except ValueError:\n+            if not len(batch):\n+                raise ValueError(\"An order message exceeds the queue batch size limit\")\n+\n+        sender.send_messages(batch)\n+        next_batch = sender.create_message_batch()\n+        try:\n+            next_batch.add_message(message)\n+        except ValueError as exc:\n+            raise ValueError(\n+                \"An order message exceeds the queue batch size limit\"\n+            ) from exc\n+        return next_batch\n*** Add File: async_sender.py\n+from __future__ import annotations\n+\n+import logging\n+from collections import defaultdict\n+from collections.abc import Iterable\n+\n+from azure.servicebus import ServiceBusMessageBatch\n+from azure.servicebus.aio import ServiceBusClient\n+\n+from messages import create_order_message\n+from order import Order\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+class AsyncOrderSender:\n+    def __init__(\n+        self,\n+        client: ServiceBusClient,\n+        queue_name: str,\n+        high_priority_threshold: float = 1_000.0,\n+    ) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+        self._high_priority_threshold = high_priority_threshold\n+\n+    async def send_order(self, order: Order) -\u003e None:\n+        message = create_order_message(order, self._high_priority_threshold)\n+        async with self._client.get_queue_sender(queue_name=self._queue_name) as sender:\n+            await sender.send_messages(message)\n+        logger.info(\"Sent order %s for %s\", order.order_id, order.customer_name)\n+\n+    async def send_orders(self, orders: Iterable[Order]) -\u003e None:\n+        orders_by_customer: dict[str, list[Order]] = defaultdict(list)\n+        for order in orders:\n+            orders_by_customer[order.customer_name].append(order)\n+\n+        async with self._client.get_queue_sender(\n+            queue_name=self._queue_name\n+        ) as sender:\n+            for customer_orders in orders_by_customer.values():\n+                batch = await sender.create_message_batch()\n+                for order in customer_orders:\n+                    message = create_order_message(\n+                        order,\n+                        self._high_priority_threshold,\n+                    )\n+                    batch = await self._add_or_flush(sender, batch, message)\n+                if len(batch):\n+                    await sender.send_messages(batch)\n+\n+        logger.info(\"Sent %d customer order group(s)\", len(orders_by_customer))\n+\n+    @staticmethod\n+    async def _add_or_flush(sender, batch: ServiceBusMessageBatch, message):\n+        try:\n+            batch.add_message(message)\n+            return batch\n+        except ValueError:\n+            if not len(batch):\n+                raise ValueError(\"An order message exceeds the queue batch size limit\")\n+\n+        await sender.send_messages(batch)\n+        next_batch = await sender.create_message_batch()\n+        try:\n+            next_batch.add_message(message)\n+        except ValueError as exc:\n+            raise ValueError(\n+                \"An order message exceeds the queue batch size limit\"\n+            ) from exc\n+        return next_batch\n*** Add File: sync_processor.py\n+from __future__ import annotations\n+\n+import json\n+import logging\n+from collections.abc import Callable, Iterable\n+from dataclasses import dataclass\n+\n+from azure.servicebus import (\n+    ServiceBusClient,\n+    ServiceBusReceivedMessage,\n+    ServiceBusSubQueue,\n+)\n+\n+from messages import create_reprocessed_message\n+from order import Order, OrderStatus\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+@dataclass(frozen=True, slots=True)\n+class DeadLetterRecord:\n+    body: str\n+    reason: str\n+    description: str\n+    correlation_id: str | None\n+\n+\n+class SyncOrderProcessor:\n+    def __init__(self, client: ServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    def process_customers(\n+        self,\n+        customer_names: Iterable[str],\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        for customer_name in dict.fromkeys(customer_names):\n+            processed += self.process_customer(customer_name, max_wait_time)\n+        return processed\n+\n+    def process_customer(\n+        self,\n+        customer_name: str,\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        receiver = self._client.get_queue_receiver(\n+            queue_name=self._queue_name,\n+            session_id=customer_name,\n+            max_wait_time=max_wait_time,\n+        )\n+        with receiver:\n+            while True:\n+                messages = receiver.receive_messages(\n+                    max_message_count=20,\n+                    max_wait_time=max_wait_time,\n+                )\n+                if not messages:\n+                    break\n+                for message in messages:\n+                    if self._process_message(receiver, message):\n+                        processed += 1\n+        return processed\n+\n+    def inspect_dead_letters(\n+        self,\n+        customer_name: str,\n+        replacement_factory: Callable[[DeadLetterRecord], Order | None] | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e list[DeadLetterRecord]:\n+        records: list[DeadLetterRecord] = []\n+        receiver = self._client.get_queue_receiver(\n+            queue_name=self._queue_name,\n+            sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+            session_id=customer_name,\n+            max_wait_time=max_wait_time,\n+        )\n+        with receiver, self._client.get_queue_sender(\n+            queue_name=self._queue_name\n+        ) as sender:\n+            messages = receiver.receive_messages(\n+                max_message_count=100,\n+                max_wait_time=max_wait_time,\n+            )\n+            for message in messages:\n+                record = DeadLetterRecord(\n+                    body=str(message),\n+                    reason=message.dead_letter_reason or \"Unknown\",\n+                    description=message.dead_letter_error_description or \"\",\n+                    correlation_id=message.correlation_id,\n+                )\n+                records.append(record)\n+                logger.warning(\n+                    \"DLQ order correlation_id=%s reason=%s description=%s body=%s\",\n+                    record.correlation_id,\n+                    record.reason,\n+                    record.description,\n+                    record.body,\n+                )\n+\n+                replacement = replacement_factory(record) if replacement_factory else None\n+                if replacement is None:\n+                    receiver.abandon_message(message)\n+                    continue\n+\n+                sender.send_messages(\n+                    create_reprocessed_message(replacement, record.reason)\n+                )\n+                receiver.complete_message(message)\n+                logger.info(\"Reprocessed dead-lettered order as %s\", replacement.order_id)\n+        return records\n+\n+    @staticmethod\n+    def _process_message(receiver, message: ServiceBusReceivedMessage) -\u003e bool:\n+        try:\n+            order = Order.from_json(str(message))\n+            order.status = OrderStatus.PROCESSING\n+            logger.info(\n+                \"Processing order %s for %s: %d x %s ($%.2f)\",\n+                order.order_id,\n+                order.customer_name,\n+                order.quantity,\n+                order.product,\n+                order.total_price,\n+            )\n+            order.status = OrderStatus.COMPLETED\n+            receiver.complete_message(message)\n+            logger.info(\"Completed order %s\", order.order_id)\n+            return True\n+        except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:\n+            receiver.dead_letter_message(\n+                message,\n+                reason=\"InvalidOrderPayload\",\n+                error_description=f\"{type(exc).__name__}: {exc}\",\n+            )\n+            logger.error(\"Dead-lettered invalid order: %s\", exc)\n+            return False\n+        except Exception as exc:\n+            receiver.dead_letter_message(\n+                message,\n+                reason=\"OrderProcessingFailed\",\n+                error_description=f\"{type(exc).__name__}: {exc}\",\n+            )\n+            logger.exception(\"Dead-lettered order after processing failure\")\n+            return False\n*** Add File: async_processor.py\n+from __future__ import annotations\n+\n+import json\n+import logging\n+from collections.abc import Callable, Iterable\n+\n+from azure.servicebus import ServiceBusReceivedMessage, ServiceBusSubQueue\n+from azure.servicebus.aio import ServiceBusClient\n+\n+from messages import create_reprocessed_message\n+from order import Order, OrderStatus\n+from sync_processor import DeadLetterRecord\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+class AsyncOrderProcessor:\n+    def __init__(self, client: ServiceBusClient, queue_name: str) -\u003e None:\n+        self._client = client\n+        self._queue_name = queue_name\n+\n+    async def process_customers(\n+        self,\n+        customer_names: Iterable[str],\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        for customer_name in dict.fromkeys(customer_names):\n+            processed += await self.process_customer(customer_name, max_wait_time)\n+        return processed\n+\n+    async def process_customer(\n+        self,\n+        customer_name: str,\n+        max_wait_time: float = 5,\n+    ) -\u003e int:\n+        processed = 0\n+        receiver = self._client.get_queue_receiver(\n+            queue_name=self._queue_name,\n+            session_id=customer_name,\n+            max_wait_time=max_wait_time,\n+        )\n+        async with receiver:\n+            while True:\n+                messages = await receiver.receive_messages(\n+                    max_message_count=20,\n+                    max_wait_time=max_wait_time,\n+                )\n+                if not messages:\n+                    break\n+                for message in messages:\n+                    if await self._process_message(receiver, message):\n+                        processed += 1\n+        return processed\n+\n+    async def inspect_dead_letters(\n+        self,\n+        customer_name: str,\n+        replacement_factory: Callable[[DeadLetterRecord], Order | None] | None = None,\n+        max_wait_time: float = 5,\n+    ) -\u003e list[DeadLetterRecord]:\n+        records: list[DeadLetterRecord] = []\n+        receiver = self._client.get_queue_receiver(\n+            queue_name=self._queue_name,\n+            sub_queue=ServiceBusSubQueue.DEAD_LETTER,\n+            session_id=customer_name,\n+            max_wait_time=max_wait_time,\n+        )\n+        sender = self._client.get_queue_sender(queue_name=self._queue_name)\n+        async with receiver, sender:\n+            messages = await receiver.receive_messages(\n+                max_message_count=100,\n+                max_wait_time=max_wait_time,\n+            )\n+            for message in messages:\n+                record = DeadLetterRecord(\n+                    body=str(message),\n+                    reason=message.dead_letter_reason or \"Unknown\",\n+                    description=message.dead_letter_error_description or \"\",\n+                    correlation_id=message.correlation_id,\n+                )\n+                records.append(record)\n+                logger.warning(\n+                    \"DLQ order correlation_id=%s reason=%s description=%s body=%s\",\n+                    record.correlation_id,\n+                    record.reason,\n+                    record.description,\n+                    record.body,\n+                )\n+\n+                replacement = replacement_factory(record) if replacement_factory else None\n+                if replacement is None:\n+                    await receiver.abandon_message(message)\n+                    continue\n+\n+                await sender.send_messages(\n+                    create_reprocessed_message(replacement, record.reason)\n+                )\n+                await receiver.complete_message(message)\n+                logger.info(\"Reprocessed dead-lettered order as %s\", replacement.order_id)\n+        return records\n+\n+    @staticmethod\n+    async def _process_message(receiver, message: ServiceBusReceivedMessage) -\u003e bool:\n+        try:\n+            order = Order.from_json(str(message))\n+            order.status = OrderStatus.PROCESSING\n+            logger.info(\n+                \"Processing order %s for %s: %d x %s ($%.2f)\",\n+                order.order_id,\n+                order.customer_name,\n+                order.quantity,\n+                order.product,\n+                order.total_price,\n+            )\n+            order.status = OrderStatus.COMPLETED\n+            await receiver.complete_message(message)\n+            logger.info(\"Completed order %s\", order.order_id)\n+            return True\n+        except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:\n+            await receiver.dead_letter_message(\n+                message,\n+                reason=\"InvalidOrderPayload\",\n+                error_description=f\"{type(exc).__name__}: {exc}\",\n+            )\n+            logger.error(\"Dead-lettered invalid order: %s\", exc)\n+            return False\n+        except Exception as exc:\n+            await receiver.dead_letter_message(\n+                message,\n+                reason=\"OrderProcessingFailed\",\n+                error_description=f\"{type(exc).__name__}: {exc}\",\n+            )\n+            logger.exception(\"Dead-lettered order after processing failure\")\n+            return False\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+from uuid import uuid4\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.servicebus import ServiceBusClient, ServiceBusMessage\n+from azure.servicebus.aio import ServiceBusClient as AsyncServiceBusClient\n+\n+from async_processor import AsyncOrderProcessor\n+from async_sender import AsyncOrderSender\n+from order import Order\n+from sync_processor import DeadLetterRecord, SyncOrderProcessor\n+from sync_sender import SyncOrderSender\n+\n+HIGH_PRIORITY_THRESHOLD = 1_000.0\n+SCHEDULE_WAIT_SECONDS = 35\n+SYNC_POISON_SESSION = \"Sync Poison Demo\"\n+ASYNC_POISON_SESSION = \"Async Poison Demo\"\n+\n+logger = logging.getLogger(__name__)\n+\n+\n+def demo_orders(prefix: str) -\u003e list[Order]:\n+    return [\n+        Order(f\"{prefix}-001\", \"Ada Lovelace\", \"Mechanical Keyboard\", 1, 149.99),\n+        Order(f\"{prefix}-002\", \"Grace Hopper\", \"GPU Workstation\", 1, 4_999.00),\n+        Order(f\"{prefix}-003\", \"Ada Lovelace\", \"USB-C Dock\", 2, 259.98),\n+    ]\n+\n+\n+def replacement_order(prefix: str, customer_name: str) -\u003e Order:\n+    return Order(\n+        order_id=f\"{prefix}-recovered\",\n+        customer_name=customer_name,\n+        product=\"Recovered Demo Item\",\n+        quantity=1,\n+        total_price=25.00,\n+    )\n+\n+\n+def send_sync_poison_message(client: ServiceBusClient, queue_name: str) -\u003e None:\n+    message = ServiceBusMessage(\n+        '{\"order_id\": \"broken\"',\n+        content_type=\"application/json\",\n+        correlation_id=\"sync-invalid\",\n+        message_id=str(uuid4()),\n+        session_id=SYNC_POISON_SESSION,\n+    )\n+    with client.get_queue_sender(queue_name=queue_name) as sender:\n+        sender.send_messages(message)\n+\n+\n+async def send_async_poison_message(\n+    client: AsyncServiceBusClient,\n+    queue_name: str,\n+) -\u003e None:\n+    message = ServiceBusMessage(\n+        '{\"order_id\": \"broken\"',\n+        content_type=\"application/json\",\n+        correlation_id=\"async-invalid\",\n+        message_id=str(uuid4()),\n+        session_id=ASYNC_POISON_SESSION,\n+    )\n+    async with client.get_queue_sender(queue_name=queue_name) as sender:\n+        await sender.send_messages(message)\n+\n+\n+def run_sync_demo(namespace: str, queue_name: str) -\u003e None:\n+    logger.info(\"Starting synchronous Service Bus demo\")\n+    with DefaultAzureCredential() as credential, ServiceBusClient(\n+        fully_qualified_namespace=namespace,\n+        credential=credential,\n+    ) as client:\n+        sender = SyncOrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)\n+        processor = SyncOrderProcessor(client, queue_name)\n+        orders = demo_orders(\"sync\")\n+\n+        sender.send_order(orders[0])\n+        sender.send_orders(orders[1:])\n+        send_sync_poison_message(client, queue_name)\n+\n+        processor.process_customers(\n+            [order.customer_name for order in orders],\n+            max_wait_time=SCHEDULE_WAIT_SECONDS,\n+        )\n+        processor.process_customer(SYNC_POISON_SESSION)\n+        processor.inspect_dead_letters(\n+            SYNC_POISON_SESSION,\n+            replacement_factory=lambda _: replacement_order(\n+                \"sync\",\n+                SYNC_POISON_SESSION,\n+            ),\n+        )\n+        processor.process_customer(SYNC_POISON_SESSION)\n+    logger.info(\"Synchronous demo complete\")\n+\n+\n+async def run_async_demo(namespace: str, queue_name: str) -\u003e None:\n+    logger.info(\"Starting asynchronous Service Bus demo\")\n+    async with AsyncDefaultAzureCredential() as credential, AsyncServiceBusClient(\n+        fully_qualified_namespace=namespace,\n+        credential=credential,\n+    ) as client:\n+        sender = AsyncOrderSender(client, queue_name, HIGH_PRIORITY_THRESHOLD)\n+        processor = AsyncOrderProcessor(client, queue_name)\n+        orders = demo_orders(\"async\")\n+\n+        await sender.send_order(orders[0])\n+        await sender.send_orders(orders[1:])\n+        await send_async_poison_message(client, queue_name)\n+\n+        await processor.process_customers(\n+            [order.customer_name for order in orders],\n+            max_wait_time=SCHEDULE_WAIT_SECONDS,\n+        )\n+        await processor.process_customer(ASYNC_POISON_SESSION)\n+        await processor.inspect_dead_letters(\n+            ASYNC_POISON_SESSION,\n+            replacement_factory=lambda _: replacement_order(\n+                \"async\",\n+                ASYNC_POISON_SESSION,\n+            ),\n+        )\n+        await processor.process_customer(ASYNC_POISON_SESSION)\n+    logger.info(\"Asynchronous demo complete\")\n+\n+\n+def get_required_environment_variable(name: str) -\u003e str:\n+    value = os.getenv(name)\n+    if not value:\n+        raise RuntimeError(f\"Set the required environment variable {name}\")\n+    return value\n+\n+\n+def main() -\u003e None:\n+    logging.basicConfig(\n+        level=logging.INFO,\n+        format=\"%(asctime)s %(levelname)s %(name)s: %(message)s\",\n+    )\n+    namespace = get_required_environment_variable(\n+        \"SERVICEBUS_FULLY_QUALIFIED_NAMESPACE\"\n+    )\n+    queue_name = get_required_environment_variable(\"SERVICEBUS_QUEUE_NAME\")\n+\n+    run_sync_demo(namespace, queue_name)\n+    asyncio.run(run_async_demo(namespace, queue_name))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-identity\n+azure-servicebus\n*** Add File: README.md\n+# Azure Service Bus Order Processor\n+\n+This sample sends and processes orders with both synchronous and asynchronous\n+Azure Service Bus clients. Customer names are used as Service Bus session IDs,\n+so the queue must be created with **sessions enabled**. A processor drains one\n+customer session before opening the next, preserving per-customer FIFO order.\n+\n+High-value orders have an application property of `priority=high` and a\n+scheduled enqueue time 30 seconds in the future. Azure Service Bus queues do\n+not provide a native priority queue, so consumers can use this property for\n+telemetry or downstream routing.\n+\n+## Run\n+\n+Use Python 3.10 or newer, install the dependencies, and configure an identity\n+with the Azure Service Bus Data Sender and Azure Service Bus Data Receiver\n+roles on the namespace or queue.\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+pip install -r requirements.txt\n+$env:SERVICEBUS_FULLY_QUALIFIED_NAMESPACE = \"example.servicebus.windows.net\"\n+$env:SERVICEBUS_QUEUE_NAME = \"orders\"\n+python main.py\n+```\n+\n+The demo first runs the synchronous flow and then the asynchronous flow. Each\n+flow sends normal and scheduled orders, intentionally sends malformed JSON,\n+dead-letters that message, inspects it, replaces it with a valid order, and\n+processes the replacement.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/order.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/order.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/order.py
@@ -1,0 +1,78 @@
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
+    total_price: float
+    status: OrderStatus = OrderStatus.PENDING
+
+    def __post_init__(self) -> None:
+        if not self.order_id.strip():
+            raise ValueError("order_id must not be empty")
+        if not self.customer_name.strip():
+            raise ValueError("customer_name must not be empty")
+        if not self.product.strip():
+            raise ValueError("product must not be empty")
+        if isinstance(self.quantity, bool) or not isinstance(self.quantity, int):
+            raise TypeError("quantity must be an integer")
+        if self.quantity <= 0:
+            raise ValueError("quantity must be greater than zero")
+        if isinstance(self.total_price, bool) or not isinstance(
+            self.total_price, (int, float)
+        ):
+            raise TypeError("total_price must be numeric")
+        if self.total_price < 0:
+            raise ValueError("total_price must not be negative")
+        if not isinstance(self.status, OrderStatus):
+            self.status = OrderStatus(self.stat
... (truncated)
```

### ✅ 9. powershell — 6002ms

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt \u0026\u0026 python -m compileall -q . \u0026\u0026 python -c \"import inspect; from azure.servicebus import ServiceBusMessage, ServiceBusClient; from azure.servicebus.aio import ServiceBusClient as A; print(inspect.signature(ServiceBusMessage)); print(inspect.signature(ServiceBusClient.get_queue_receiver)); print(inspect.signature(A.get_queue_receiver))\"","description":"Install dependencies and validate SDK usage","initial_wait":180,"mode":"sync"}
```

**Output:**

```
(body: 'Optional[Union[str, bytes]]', *, application_properties: "Optional[Dict[Union[str, bytes], 'PrimitiveTypes']]" = None, session_id: 'Optional[str]' = None, message_id: 'Optional[str]' = None, scheduled_enqueue_time_utc: 'Optional[datetime.datetime]' = None, time_to_live: 'Optional[datetime.timedelta]' = None, content_type: 'Optional[str]' = None, correlation_id: 'Optional[str]' = None, subject: 'Optional[str]' = None, partition_key: 'Optional[str]' = None, to: 'Optional[str]' = None, reply_to: 'Optional[str]' = None, reply_to_session_id: 'Optional[str]' = None, **kwargs: 'Any') -> 'None'
(self, queue_name: str, *, client_identifier: Optional[str] = None, socket_timeout: Optional[float] = None, session_id: Union[str, Literal[<ServiceBusSessionFilter.NEXT_AVAILABLE: 0>], NoneType] = None, sub_queue: Union[azure.servicebus._common.constants.ServiceBusSubQueue, str, NoneType] = None, receive_mode: Union[azure.servicebus._common.constants.ServiceBusReceiveMode, str] = <ServiceBusReceiveMode.PEEK_LOCK: 'peeklock'>, max_wait_time: Optional[float] = None, auto_lock_renewer: Optional[azure.servicebus._common.auto_lock_renewer.AutoLockRenewer] = None, prefetch_count: int = 0, **kwargs: Any) -> azure.servicebus._servicebus_receiver.ServiceBusReceiver
(self, queue_name: str, *, client_identifier: Optional[str] = None, socket_timeout: Optional[float] = None, session_id: Union[str, Literal[<ServiceBusSessionFilter.NEXT_AVAILABLE: 0>], NoneType] = None, sub_queue: Union[azure.servicebus._common.constants.ServiceBusSubQueue, str, NoneType] = None, receive_mode: Union[azure.servicebus._common.constants.ServiceBusReceiveMode, str] = <ServiceBusReceiveMode.PEEK_LOCK: 'peeklock'>, max_wait_time: Optional[float] = None, auto_lock_renewer: Optional[azure.servicebus._common.auto_lock_renewer.AutoLockRenewer] = None, prefetch_count: int = 0, **kwargs: Any) -> azure.servicebus.aio._servicebus_receiver_async.ServiceBusReceiver
<shellId: 0 completed with exit code 0>
```

### ✅ 10. apply_patch — 944ms

**Input:**

```json
"*** Begin Patch\n*** Update File: sync_processor.py\n@@\n     @staticmethod\n     def _process_message(receiver, message: ServiceBusReceivedMessage) -\u003e bool:\n         try:\n             order = Order.from_json(str(message))\n             order.status = OrderStatus.PROCESSING\n@@\n                 order.product,\n                 order.total_price,\n             )\n             order.status = OrderStatus.COMPLETED\n-            receiver.complete_message(message)\n-            logger.info(\"Completed order %s\", order.order_id)\n-            return True\n         except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:\n             receiver.dead_letter_message(\n                 message,\n                 reason=\"InvalidOrderPayload\",\n                 error_description=f\"{type(exc).__name__}: {exc}\",\n@@\n             logger.exception(\"Dead-lettered order after processing failure\")\n             return False\n+\n+        receiver.complete_message(message)\n+        logger.info(\"Completed order %s\", order.order_id)\n+        return True\n*** Update File: async_processor.py\n@@\n     @staticmethod\n     async def _process_message(receiver, message: ServiceBusReceivedMessage) -\u003e bool:\n         try:\n             order = Order.from_json(str(message))\n             order.status = OrderStatus.PROCESSING\n@@\n                 order.product,\n                 order.total_price,\n             )\n             order.status = OrderStatus.COMPLETED\n-            await receiver.complete_message(message)\n-            logger.info(\"Completed order %s\", order.order_id)\n-            return True\n         except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:\n             await receiver.dead_letter_message(\n                 message,\n                 reason=\"InvalidOrderPayload\",\n                 error_description=f\"{type(exc).__name__}: {exc}\",\n@@\n             logger.exception(\"Dead-lettered order after processing failure\")\n             return False\n+\n+        await receiver.complete_message(message)\n+        logger.info(\"Completed order %s\", order.order_id)\n+        return True\n*** Update File: main.py\n@@\n from async_processor import AsyncOrderProcessor\n from async_sender import AsyncOrderSender\n from order import Order\n-from sync_processor import DeadLetterRecord, SyncOrderProcessor\n+from sync_processor import SyncOrderProcessor\n from sync_sender import SyncOrderSender\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/sync_processor.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/sync_processor.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/sync_processor.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/sync_processor.py
@@ -126,9 +126,6 @@
                 order.total_price,
             )
             order.status = OrderStatus.COMPLETED
-            receiver.complete_message(message)
-            logger.info("Completed order %s", order.order_id)
-            return True
         except (json.JSONDecodeError, UnicodeDecodeError, TypeError, ValueError) as exc:
             receiver.dead_letter_message(
                 message,
@@ -146,3 +143,7 @@
             logger.exception("Dead-lettered order after processing failure")
             return False
 
+        receiver.complete_message(message)
+        logger.info("Completed order %s", order.order_id)
+        return True
+


diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/async_processor.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/async_processor.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-azure-skill-mcp-microsoft-skill-724431187/async_processor.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-service-bus-dp-python-order-processor-python-azure-skills-
... (truncated)
```

### ✅ 11. powershell — 3189ms

**Input:**

```json
{"command":"python -m compileall -q .; @'\nimport asyncio\nfrom datetime import datetime, timezone\n\nfrom async_sender import AsyncOrderSender\nfrom messages import create_order_message\nfrom order import Order\nfrom sync_sender import SyncOrderSender\n\norder = Order(\"o-1\", \"Customer\", \"Product\", 2, 1500, \"pending\")\nassert Order.from_json(order.to_json()) == order\nmessage = create_order_message(order, 1000)\nassert message.correlation_id == \"o-1\"\nassert message.session_id == \"Customer\"\nassert message.application_properties[\"priority\"] == \"high\"\nassert 25 \u003c= (message.scheduled_enqueue_time_utc - datetime.now(timezone.utc)).total_seconds() \u003c= 30\n\nclass Batch:\n    def __init__(self, capacity=1):\n        self.items = []\n        self.capacity = capacity\n    def add_message(self, message):\n        if len(self.items) \u003e= self.capacity:\n            raise ValueError(\"full\")\n        self.items.append(message)\n    def __len__(self):\n        return len(self.items)\n\nclass SyncSender:\n    def __init__(self):\n        self.sent = []\n    def create_message_batch(self):\n        return Batch()\n    def send_messages(self, batch):\n        self.sent.append(list(batch.items))\n\nsync_sender = SyncSender()\nbatch = Batch()\nbatch.add_message(\"first\")\nnext_batch = SyncOrderSender._add_or_flush(sync_sender, batch, \"second\")\nassert sync_sender.sent == [[\"first\"]] and next_batch.items == [\"second\"]\n\nclass AsyncSender(SyncSender):\n    async def create_message_batch(self):\n        return Batch()\n    async def send_messages(self, batch):\n        self.sent.append(list(batch.items))\n\nasync def check_async_batch():\n    sender = AsyncSender()\n    batch = Batch()\n    batch.add_message(\"first\")\n    next_batch = await AsyncOrderSender._add_or_flush(sender, batch, \"second\")\n    assert sender.sent == [[\"first\"]] and next_batch.items == [\"second\"]\n\nasyncio.run(check_async_batch())\nprint(\"Local behavioral checks passed\")\n'@ | python -","description":"Run local behavioral checks","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Local behavioral checks passed
<shellId: 1 completed with exit code 0>
```

## Generated Files

- `README.md`
- `async_processor.py`
- `async_sender.py`
- `main.py`
- `messages.py`
- `order.py`
- `requirements.txt`
- `sync_processor.py`
- `sync_sender.py`

## Copilot Response

I’m grounding the implementation in the Service Bus SDK’s session, batching, scheduling, and dead-letter patterns before creating the project.

The design will use one Service Bus session per customer, drain one session at a time, and preserve that session ID when replaying dead-lettered messages. High-value orders will carry a high-priority application property and a UTC enqueue time 30 seconds ahead.

I’ve resolved the main batching constraint: orders are grouped by customer before filling `ServiceBusMessageBatch` instances, so session affinity and per-customer FIFO ordering are retained even on partitioned/session-enabled queues. I’m now creating the modules and a runnable demo.

The project files are in place. I’m validating them against the installed SDK surface now, especially context-manager support, message scheduling fields, session/DLQ receiver arguments, and batch APIs.

The SDK signatures match the implementation. I’ve also separated local processing from message settlement so a transport failure during `complete_message()` is surfaced instead of being misclassified as an order-processing failure. I’m running local behavioral checks without contacting Azure.

Created the Python order-processing project with:

- JSON-serializable `Order` model and validated statuses.
- Sync/async senders with correlation IDs, customer sessions, scheduled high-priority orders, and size-aware `ServiceBusMessageBatch` handling.
- Sync/async processors with ordered session processing, explicit dead-lettering, DLQ inspection, and reprocessing.
- `main.py` demonstrating complete sync and async cycles using `DefaultAzureCredential`.
- `requirements.txt` and setup instructions in `README.md`.

The target Service Bus queue must have **sessions enabled**. Implementation follows the [Azure Service Bus Python SDK guidance](https://learn.microsoft.com/azure/service-bus-messaging/service-bus-python-how-to-use-queues).

## Reviewer Notes (LLM-as-Judge)

### Summary

0/1 criteria passed. Azure Service Bus operational exceptions are not handled appropriately.

### Strengths

- Payload deserialization errors are handled explicitly and result in dead-lettering with a reason and description.

### Issues

- No explicit handling for HttpResponseError, ServiceBusError, or relevant subclasses around Azure SDK operations.
- The broad processor exception handler conflates application processing failures with SDK failures and attempts dead-letter settlement without handling settlement errors.

## Grader Results

- order-processor.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (11/14)
      - Sender uses `ServiceBusClient.get_queue_sender()` (or async equivalent): Pass
      - Processor uses `ServiceBusClient.get_queue_receiver()` or session-enabled receiver: Pass
      - Batch sending: creates `ServiceBusMessageBatch`, checks `add_message()` return or catches `MessageSizeExceededError`: Fail
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
  - Tool Usage Verification (tool): Pass (1/1)
      - tool_used: any tool (source=mcp, server=azure): Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 79% | 1.00 | 0.7857 | 13.6% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 17.3% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.3% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 17.3% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 17.3% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 17.3% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 5.7857** | **72.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id service-bus-dp-python-order-processor --config python-azure-skills/azure-skill-mcp-microsoft-skill --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
