# Evaluation Report: identity-dp-python-credential-chain

**Config:** python-azure-skills/azure-skill-mcp | **Result:** ❌ FAILED | **Duration:** 321.4s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-python-credential-chain` |
| Config | python-azure-skills/azure-skill-mcp |
| Result | ❌ FAILED |
| Score | 19/21 |
| Duration | 321.4s |
| Timestamp | 2026-08-28T19:42:16Z |
| Files Generated | 8 |
| Event Count | 14330 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 218.0s |
| Review | 102.2s |
| **Total** | **321.4s** |

## Configuration

- **name:** python-azure-skills/azure-skill-mcp
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | airunway-aks-setup, appinsights-instrumentation, azure-ai, azure-aigateway, azure-app-onboard, azure-app-onboard-prereq, azure-cloud-migrate, azure-compliance, azure-compute, azure-cost, azure-deploy, azure-diagnostics, azure-enterprise-infra-planner, azure-kubernetes, azure-kusto, azure-messaging, azure-prepare, azure-quotas, azure-reliability, azure-resource-lookup, azure-resource-visualizer, azure-storage, azure-upgrade, azure-validate, entra-agent-id, entra-app-registration, microsoft-foundry, python-appservice-deploy, customize-cloud-agent, github-pr-media |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=481920 out=9725 |
| Turn Count | 10 |

## Tool Calls

### ❌ 1. skill — 2062ms

**Input:**

```json
{"skill":"azure-identity"}
```

**Error:**

```
Skill not found: azure-identity
```

### ✅ 2. azure-get_azure_bestpractices (via azure) — 7609ms

**Input:**

```json
{"intent":"Create a local Python sample demonstrating Azure Identity credential chains for developer tools, CI/CD, managed identity, workload identity, CAE token requests, and sync/async credential usage. Do not deploy or modify Azure resources.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'get_azure_bestpractices' tool.
If you do not find a suitable "command", run again with the "learn=true" to get a list of available commands and their parameters.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"get_azure_bestpractices_get","description":"This tool returns a list of best practices for code generation, operations and deployment\r\nwhen working with Azure services. It should be called for any code generation, deployment or\r\noperations involving Azure, Azure Functions, Azure Kubernetes Service (AKS), Azure Container\r\nApps (ACA), Bicep, Terraform, Azure Cache, Redis, CosmosDB, Entra, Azure Active Directory,\r\nAzure App Services, or any other Azure technology or programming language. Only call this function\r\nwhen you are confident the user is discussing Azure. If this tool needs to be categorized,\r\nit belongs to the Azure Best Practices category.","inputSchema":{"type":"object","properties":{"resource":{"type":"string","description":"The Azure resource type for which to get best practices. Options: \u0027general\u0027 (general Azure), \u0027azurefunctions\u0027 (Azure Functions), \u0027static-web-app\u0027 (Azure Static Web Apps), \u0027coding-agent\u0027 (Coding Agent)."},"action":{"type":"string","description":"The action type for the best practices. Options: \u0027all\u0027, \u0027code-generation\u0027, \u0027deployment\u0027. Note: \u0027static-web-app\u0027 and \u0027coding-agent\u0027 resources only supports \u0027all\u0027."}},"required":["resource","action"],"additionalProperties":false}},{"command":"get_azure_bestpractices_ai_app","description":"Returns best practices and code generation guidance for building AI applications in Azure.\r\nUse this command when you need recommendations on how to write code for AI agents, chatbots, workflows, or any AI / LLM fea
... (truncated)
```

### ✅ 3. azure-documentation (via azure) — 10162ms

**Input:**

```json
{"intent":"Find current official Python Azure Identity documentation for ChainedTokenCredential, AzureCliCredential, AzureDeveloperCliCredential, VisualStudioCodeCredential, EnvironmentCredential, AzurePipelinesCredential, ManagedIdentityCredential, WorkloadIdentityCredential, async equivalents, and enable_cae on get_token.","learn":true}
```

**Output:**

```
Here are the available commands and their input schema for 'documentation' tool.
If you do not find a suitable command, run again with the "learn=true" to get a list of available commands and their input schema.
Next, identify the command you want to execute and run again with the "command" and "parameters" arguments, respecting "required" parameters if present.

[{"command":"microsoft_docs_search","description":"Search official Microsoft/Azure documentation to find the most relevant and trustworthy content for a user\u0027s query. This tool returns up to 10 high-quality content chunks (each max 500 tokens), extracted from Microsoft Learn and other official sources. Each result includes the article title, URL, and a self-contained content excerpt optimized for fast retrieval and reasoning. Always use this tool to quickly ground your answers in accurate, first-party Microsoft/Azure knowledge.\n\n## Follow-up Pattern\nTo ensure completeness, use microsoft_docs_fetch when high-value pages are identified by search. The fetch tool complements search by providing the full detail. This is a required step for comprehensive results.","inputSchema":{"type":"object","properties":{"query":{"description":"a query or topic about Microsoft/Azure products, services, platforms, developer tools, frameworks, or APIs","type":"string","default":null}}}},{"command":"microsoft_code_sample_search","description":"Search for code snippets and examples in official Microsoft Learn documentation. This tool retrieves relevant code samples from Microsoft documentation pages providing developers with practical implementation examples and best practices for Microsoft/Azure products and services related coding tasks. This tool will help you use the **LATEST OFFICIAL** code snippets to empower coding capabilities.\n\n## When to Use This Tool\n- When you are going to provide sample Microsoft/Azure related code snippets in your answers.\n- When you are **generating any Microsoft/Azure related code*
... (truncated)
```

### ✅ 4. glob — 8887ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 5. azure-get_azure_bestpractices (via azure) — 6777ms

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get code-generation best practices for a local Python Azure Identity credential-chain sample.","parameters":{"action":"code-generation","resource":"general"}}
```

**Output:**

```
{"status":200,"message":"","results":["Prefer topics about Azure and Azure technologies. Focus on newer and more popular Azure technologies and services such as Azure Containers, Azure Kubernetes Service (AKS), Azure Container Applications (ACA), Azure Container Instances (ACI), Azure Functions, and Azure App Services.\r\nFor azure sdk code samples always prefer managed identity, never use key based authentication, always tell me what you used for references, always list the http link to the sample reference when returning a code sample.\r\nGet the latest version of the azure sdk for any sample code you return.\r\n\r\nYou are an expert in Azure services implementation. When asked to generate code for any Azure service:\r\n- Parse the user's requirements to identify which Azure service(s) to implement\r\n- Follow these comprehensive guidelines to ensure high-quality, secure code\r\n- Validate your output against all requirements before returning\r\n\r\nGenerate secure, efficient, and maintainable Azure service code following these requirements:\r\n- Use appropriate auth: Managed Identity (Azure-hosted), Service Principal (CI/CD), Interactive Browser (user apps), Client Secret (daemons)\r\n- Never hardcode credentials; use Key Vault\r\n- Implement credential rotation and least privilege\r\n- Enable encryption and secure connections\r\n\r\nError Handling & Reliability\r\n- Implement retry logic with exponential backoff for transient failures\r\n- Add proper logging and monitoring\r\n- Include circuit breakers where needed\r\n- Handle service-specific errors\r\n- Ensure proper resource cleanup\r\n\r\nPerformance & Scaling\r\n- Use connection pooling for databases\r\n- Configure concurrent operations and timeouts\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nAuthentication & Security\r\n- Implement caching strategically\r\n- Monitor resource usage\r\n- Optimize batch operations\r\n\r\nDatabase Operations\r\n- Use p
... (truncated)
```

### ✅ 6. azure-documentation (via azure) — 10538ms

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Ground Azure Identity credential behavior in official docs.","parameters":{"query":"Python Azure Identity ChainedTokenCredential AzurePipelinesCredential ManagedIdentityCredential WorkloadIdentityCredential enable_cae get_token async"}}
```

**Output:**

```
{"results":[{"title":"ChainedTokenCredential Class-Methods","content":"# ChainedTokenCredential Class\r\n\r\nA sequence of credentials that is itself a credential.\r\n\r\nIts [get_token](https://learn.microsoft.com/python/api/azure-identity/azure.identity.aio.chainedtokencredential?view=azure-python#azure-identity-aio-chainedtokencredential-get-token) method calls \u0060get_token\u0060 on each credential in the sequence, in order, returning the first valid token received. For more information, see https://aka.ms/azsdk/python/identity/credential-chains#chainedtokencredential-overview.\r\n\r\n#### Examples\r\n\r\nCreate a ChainedTokenCredential.\r\n\r\n\u0060\u0060\u0060python\r\n\r\n   from azure.identity.aio import ChainedTokenCredential, EnvironmentCredential, AzureCliCredential\r\n\r\n   credential_chain = (\r\n       # Try EnvironmentCredential first\r\n       EnvironmentCredential(),\r\n       # Fallback to Azure CLI if EnvironmentCredential fails\r\n       AzureCliCredential(),\r\n   )\r\n   credential = ChainedTokenCredential(*credential_chain)\r\n\r\n\u0060\u0060\u0060\r\n\r\n## Methods\r\n\r\n### get\\_token\r\n\r\nAsynchronously request a token from each credential, in order, returning the first token received.\r\n\r\nIf no credential provides a token, raises [ClientAuthenticationError](https://learn.microsoft.com/python/api/azure-core/azure.core.exceptions.clientauthenticationerror?view=azure-python) with an error message from each credential.\r\n\r\nThis method is called automatically by Azure SDK clients.\r\n\r\n\u0060\u0060\u0060python\r\nasync get_token(*scopes: str, claims: str | None = None, tenant_id: str | None = None, enable_cae: bool = False, **kwargs: Any) -\u003E AccessToken\r\n\u0060\u0060\u0060\r\n\r\n#### Parameters\r\n\r\n| Name | Description |\r\n| --- | --- |\r\n| scopes\u003Cbr\u003E\u003Cbr\u003ERequired | [str](https://docs.python.org/3/library/stdtypes.html#str)\u003Cbr\u003E\u003Cbr\u003Edesired scopes for the access token. This meth
... (truncated)
```

### ✅ 7. azure-documentation (via azure) — 8130ms

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python examples for Azure Identity credential chains and CAE.","parameters":{"language":"python","query":"Python azure.identity ChainedTokenCredential AzurePipelinesCredential ManagedIdentityCredential WorkloadIdentityCredential get_token enable_cae"}}
```

**Output:**

```
{"results":[{"description":"description: Creates a ChainedTokenCredential combining ManagedIdentityCredential and AzureDeveloperCliCredential to provide an ordered sequence of authentication mechanisms.\npackage: azure-identity\nlanguage: python\n","codeSnippet":"azure_credential = ChainedTokenCredential(\r\n    user_assigned_managed_identity_credential,\r\n    azure_dev_cli_credential\r\n)","link":"https://learn.microsoft.com/azure/developer/ai/get-started-securing-your-ai-app?pivots=python#exploring-the-sample-code","language":"python"},{"description":"description: Creates a ChainedTokenCredential instance combining ManagedIdentityCredential with a user-assigned client ID and AzureCliCredential, providing a customized authentication chain.\npackage: azure.identity\nlanguage: python\n","codeSnippet":"credential = ChainedTokenCredential(\r\n    ManagedIdentityCredential(client_id=user_assigned_client_id),\r\n    AzureCliCredential()\r\n)","link":"https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains#defaultazurecredential-overview","language":"python"},{"description":"description: Creates a ChainedTokenCredential instance with two development-time credentials, AzureCliCredential and AzureDeveloperCliCredential, to attempt authentication sequentially using these credentials.\npackage: azure-identity\nlanguage: python\n","codeSnippet":"credential = ChainedTokenCredential(\r\n    AzureCliCredential(),\r\n    AzureDeveloperCliCredential()\r\n)","link":"https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains#chainedtokencredential-overview","language":"python"},{"description":"description: Retrieves an access token for Azure Storage using the default managed identity in a Kubernetes cluster compute environment, without specifying a client_id.\npackage: azure.identity\nlanguage: python\n","codeSnippet":"from azure.identity import ManagedIdentityCredential\r\n    \r\n    credential = ManagedIdentityCredentia
... (truncated)
```

### ✅ 8. azure-documentation (via azure) — 6387ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the authoritative Python credential-chain guidance before implementation.","parameters":{"url":"https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains"}}
```

**Output:**

```
# Credential chains in the Azure Identity library for Python

The Azure Identity library provides *credentials*—public classes that implement the Azure Core library's [TokenCredential](/en-us/python/api/azure-core/azure.core.credentials.tokencredential) protocol. A credential represents a distinct authentication flow for acquiring an access token from Microsoft Entra ID. These credentials can be chained together to form an ordered sequence of authentication mechanisms to be attempted.

## How a chained credential works

At runtime, a credential chain attempts to authenticate using the sequence's first credential. If that credential fails to acquire an access token, the next credential in the sequence is attempted, and so on, until an access token is successfully obtained. The following sequence diagram illustrates this behavior:

![Diagram that shows credential chain sequence.](../media/mermaidjs/chain-sequence.svg)

## Why use credential chains

A chained credential can offer the following benefits:

- **Environment awareness**: Automatically selects the most appropriate credential based on the environment in which the app is running. Without it, you'd have to write code like this:

    ```python
    # Set up credential based on environment (Azure or local development)
    if os.getenv("WEBSITE_HOSTNAME"):
        credential = ManagedIdentityCredential(client_id=user_assigned_client_id)
    else:
        credential = AzureCliCredential()
    ```
- **Seamless transitions**: Your app can move from local development to your staging or production environment without changing authentication code.
- **Improved resiliency**: Includes a fallback mechanism that moves to the next credential when the prior fails to acquire an access token.

## How to choose a chained credential

There are two disparate philosophies to credential chaining:

- **"Tear down" a chain**: Start with a preconfigured chain and exclude what you don't need. For this appr
... (truncated)
```

### ✅ 9. azure-documentation (via azure) — 6058ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the authoritative AzurePipelinesCredential Python API details before implementation.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-identity/azure.identity.azurepipelinescredential?view=azure-python"}}
```

**Output:**

```
# WorkloadIdentityCredential Class

Authenticates using Microsoft Entra Workload ID.

Workload identity authentication is a feature in Azure that allows applications running on virtual machines (VMs) to access other Azure resources without the need for a service principal or managed identity. With workload identity authentication, applications authenticate themselves using their own identity, rather than using a shared service principal or managed identity. Under the hood, workload identity authentication uses the concept of Service Account Credentials (SACs), which are automatically created by Azure and stored securely in the VM. By using workload identity authentication, you can avoid the need to manage and rotate service principals or managed identities for each application on each VM. Additionally, because SACs are created automatically and managed by Azure, you don't need to worry about storing and securing sensitive credentials themselves.

The WorkloadIdentityCredential supports Azure workload identity authentication on Azure Kubernetes and acquires a token using the service account credentials available in the Azure Kubernetes environment. Refer to [this workload identity overview](/en-us/azure/aks/workload-identity-overview) for more information.

## Constructor

```python
WorkloadIdentityCredential(*, tenant_id: str | None = None, client_id: str | None = None, token_file_path: str | None = None, **kwargs: Any)
```

### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| tenant\_id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>ID of the application's Microsoft Entra tenant. Also called its "directory" ID.<br><br>Default value: None |
| client\_id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The client ID of a Microsoft Entra app registration.<br><br>Default value: None |
| token\_file\_path | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The path to a file conta
... (truncated)
```

### ✅ 10. azure-documentation (via azure) — 5738ms

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the authoritative WorkloadIdentityCredential Python API details before implementation.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-identity/azure.identity.workloadidentitycredential?view=azure-python"}}
```

**Output:**

```
# AzurePipelinesCredential Class

Authenticates using Microsoft Entra Workload ID in Azure Pipelines.

This credential enables authentication in Azure Pipelines using workload identity federation for Azure service connections.

## Constructor

```python
AzurePipelinesCredential(*, tenant_id: str, client_id: str, service_connection_id: str, system_access_token: str, **kwargs: Any)
```

### Keyword-Only Parameters

| Name | Description |
| --- | --- |
| tenant\_id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The tenant ID for the service connection. Required. |
| client\_id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The client ID for the service connection. Required. |
| service\_connection\_id | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The service connection ID for the service connection associated with the pipeline. From the service connection's configuration page URL in the Azure DevOps web portal, the ID is the value of the "resourceId" query parameter. Required. |
| system\_access\_token | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>The pipeline's System.AccessToken value. It is recommended to assign the value of System.AccessToken to a secure variable in the Azure Pipelines environment. See [https://learn.microsoft.com/azure/devops/pipelines/build/variables#systemaccesstoken](/en-us/azure/devops/pipelines/build/variables#systemaccesstoken) for more info. Required. |
| authority | [str](https://docs.python.org/3/library/stdtypes.html#str)<br><br>Authority of a Microsoft Entra endpoint, for example "login.microsoftonline.com", the authority for Azure Public Cloud (which is the default). [AzureAuthorityHosts](azure.identity.azureauthorityhosts?view=azure-python) defines authorities for other clouds. |
| additionally\_allowed\_tenants | [List](https://docs.python.org/3/library/typing.html#typing.List)[[str](https://docs.python.org/3/library/stdtypes.ht
... (truncated)
```

### ✅ 11. apply_patch — 820ms

**Input:**

```json
"*** Begin Patch\n*** Add File: credential_factory.py\n+\"\"\"Environment-specific Azure credential chains.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from dataclasses import dataclass\n+\n+from azure.core.credentials import TokenCredential\n+from azure.core.credentials_async import AsyncTokenCredential\n+from azure.identity import (\n+    AzureCliCredential,\n+    AzureDeveloperCliCredential,\n+    AzurePipelinesCredential,\n+    AzurePowerShellCredential,\n+    ChainedTokenCredential,\n+    EnvironmentCredential,\n+    ManagedIdentityCredential,\n+    VisualStudioCodeCredential,\n+    WorkloadIdentityCredential,\n+)\n+from azure.identity.aio import (\n+    AzureCliCredential as AsyncAzureCliCredential,\n+)\n+from azure.identity.aio import (\n+    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,\n+)\n+from azure.identity.aio import (\n+    AzurePipelinesCredential as AsyncAzurePipelinesCredential,\n+)\n+from azure.identity.aio import (\n+    AzurePowerShellCredential as AsyncAzurePowerShellCredential,\n+)\n+from azure.identity.aio import (\n+    ChainedTokenCredential as AsyncChainedTokenCredential,\n+)\n+from azure.identity.aio import (\n+    EnvironmentCredential as AsyncEnvironmentCredential,\n+)\n+from azure.identity.aio import (\n+    ManagedIdentityCredential as AsyncManagedIdentityCredential,\n+)\n+from azure.identity.aio import (\n+    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,\n+)\n+from azure.identity.aio import (\n+    WorkloadIdentityCredential as AsyncWorkloadIdentityCredential,\n+)\n+\n+from environment_detector import RuntimeEnvironment\n+\n+\n+@dataclass(frozen=True)\n+class SyncCredentialSelection:\n+    credential: TokenCredential\n+    strategy: str\n+    enable_cae: bool\n+\n+\n+@dataclass(frozen=True)\n+class AsyncCredentialSelection:\n+    credential: AsyncTokenCredential\n+    strategy: str\n+    enable_cae: bool\n+\n+\n+def _pipeline_settings() -\u003e dict[str, str] | None:\n+    names = {\n+        \"tenant_id\": \"AZURE_TENANT_ID\",\n+        \"client_id\": \"AZURE_CLIENT_ID\",\n+        \"service_connection_id\": \"AZURE_SERVICE_CONNECTION_ID\",\n+        \"system_access_token\": \"SYSTEM_ACCESSTOKEN\",\n+    }\n+    values = {argument: os.getenv(variable, \"\") for argument, variable in names.items()}\n+    return values if all(values.values()) else None\n+\n+\n+def _workload_identity_is_configured() -\u003e bool:\n+    return all(\n+        os.getenv(name)\n+        for name in (\n+            \"AZURE_TENANT_ID\",\n+            \"AZURE_CLIENT_ID\",\n+            \"AZURE_FEDERATED_TOKEN_FILE\",\n+        )\n+    )\n+\n+\n+def _managed_identity_client_id() -\u003e str | None:\n+    # Keep this separate from AZURE_CLIENT_ID, which workload identity also uses.\n+    return os.getenv(\"AZURE_MANAGED_IDENTITY_CLIENT_ID\") or None\n+\n+\n+def build_sync_credential(\n+    environment: RuntimeEnvironment,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e SyncCredentialSelection:\n+    \"\"\"Build a synchronous credential chain for the selected environment.\"\"\"\n+    if environment is RuntimeEnvironment.DEV:\n+        credential = ChainedTokenCredential(\n+            VisualStudioCodeCredential(),\n+            AzureCliCredential(),\n+            AzurePowerShellCredential(),\n+            AzureDeveloperCliCredential(),\n+        )\n+        strategy = \"developer tools: VS Code -\u003e Azure CLI -\u003e Azure PowerShell -\u003e Azure Developer CLI\"\n+    elif environment is RuntimeEnvironment.CI:\n+        credentials: list[TokenCredential] = []\n+        pipeline_settings = _pipeline_settings()\n+        if pipeline_settings:\n+            credentials.append(AzurePipelinesCredential(**pipeline_settings))\n+        credentials.append(EnvironmentCredential())\n+        credential = ChainedTokenCredential(*credentials)\n+        strategy = (\n+            \"Azure Pipelines service connection -\u003e environment credential\"\n+            if pipeline_settings\n+            else \"environment credential (service principal or federated token variables)\"\n+        )\n+    else:\n+        managed_identity_client_id = _managed_identity_client_id()\n+        credentials = [\n+            ManagedIdentityCredential(client_id=managed_identity_client_id)\n+        ]\n+        workload_configured = _workload_identity_is_configured()\n+        if workload_configured:\n+            credentials.append(WorkloadIdentityCredential())\n+        credential = ChainedTokenCredential(*credentials)\n+        identity_kind = (\n+            f\"user-assigned managed identity ({managed_identity_client_id})\"\n+            if managed_identity_client_id\n+            else \"system-assigned managed identity\"\n+        )\n+        strategy = (\n+            f\"{identity_kind} -\u003e Kubernetes workload identity\"\n+            if workload_configured\n+            else identity_kind\n+        )\n+\n+    return SyncCredentialSelection(credential, strategy, enable_cae)\n+\n+\n+def build_async_credential(\n+    environment: RuntimeEnvironment,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e AsyncCredentialSelection:\n+    \"\"\"Build an asynchronous credential chain for the selected environment.\"\"\"\n+    if environment is RuntimeEnvironment.DEV:\n+        credential = AsyncChainedTokenCredential(\n+            AsyncVisualStudioCodeCredential(),\n+            AsyncAzureCliCredential(),\n+            AsyncAzurePowerShellCredential(),\n+            AsyncAzureDeveloperCliCredential(),\n+        )\n+        strategy = \"developer tools: VS Code -\u003e Azure CLI -\u003e Azure PowerShell -\u003e Azure Developer CLI\"\n+    elif environment is RuntimeEnvironment.CI:\n+        credentials: list[AsyncTokenCredential] = []\n+        pipeline_settings = _pipeline_settings()\n+        if pipeline_settings:\n+            credentials.append(AsyncAzurePipelinesCredential(**pipeline_settings))\n+        credentials.append(AsyncEnvironmentCredential())\n+        credential = AsyncChainedTokenCredential(*credentials)\n+        strategy = (\n+            \"Azure Pipelines service connection -\u003e environment credential\"\n+            if pipeline_settings\n+            else \"environment credential (service principal or federated token variables)\"\n+        )\n+    else:\n+        managed_identity_client_id = _managed_identity_client_id()\n+        credentials = [\n+            AsyncManagedIdentityCredential(client_id=managed_identity_client_id)\n+        ]\n+        workload_configured = _workload_identity_is_configured()\n+        if workload_configured:\n+            credentials.append(AsyncWorkloadIdentityCredential())\n+        credential = AsyncChainedTokenCredential(*credentials)\n+        identity_kind = (\n+            f\"user-assigned managed identity ({managed_identity_client_id})\"\n+            if managed_identity_client_id\n+            else \"system-assigned managed identity\"\n+        )\n+        strategy = (\n+            f\"{identity_kind} -\u003e Kubernetes workload identity\"\n+            if workload_configured\n+            else identity_kind\n+        )\n+\n+    return AsyncCredentialSelection(credential, strategy, enable_cae)\n*** Add File: environment_detector.py\n+\"\"\"Detect the deployment environment from well-known environment variables.\"\"\"\n+\n+from __future__ import annotations\n+\n+import os\n+from enum import Enum\n+from typing import Mapping\n+\n+\n+class RuntimeEnvironment(str, Enum):\n+    DEV = \"dev\"\n+    CI = \"ci\"\n+    PRODUCTION = \"production\"\n+\n+\n+_CI_MARKERS = (\n+    \"CI\",\n+    \"TF_BUILD\",\n+    \"BUILD_BUILDID\",\n+    \"BUILD_SOURCESDIRECTORY\",\n+    \"SYSTEM_TEAMPROJECT\",\n+    \"GITHUB_ACTIONS\",\n+    \"GITLAB_CI\",\n+    \"JENKINS_URL\",\n+)\n+\n+_AZURE_HOST_MARKERS = (\n+    \"IDENTITY_ENDPOINT\",\n+    \"MSI_ENDPOINT\",\n+    \"IMDS_ENDPOINT\",\n+    \"WEBSITE_INSTANCE_ID\",\n+    \"CONTAINER_APP_NAME\",\n+)\n+\n+\n+def _is_truthy(value: str | None) -\u003e bool:\n+    return bool(value and value.strip().lower() not in {\"0\", \"false\", \"no\", \"off\"})\n+\n+\n+def detect_environment(\n+    environ: Mapping[str, str] | None = None,\n+) -\u003e RuntimeEnvironment:\n+    \"\"\"Classify the runtime as dev, CI, or production.\n+\n+    APP_ENV can explicitly override detection. CI markers take precedence over\n+    Azure-host markers because hosted pipeline agents can expose both.\n+    \"\"\"\n+    values = os.environ if environ is None else environ\n+    override = values.get(\"APP_ENV\", \"\").strip().lower()\n+    aliases = {\n+        \"dev\": RuntimeEnvironment.DEV,\n+        \"development\": RuntimeEnvironment.DEV,\n+        \"local\": RuntimeEnvironment.DEV,\n+        \"ci\": RuntimeEnvironment.CI,\n+        \"pipeline\": RuntimeEnvironment.CI,\n+        \"prod\": RuntimeEnvironment.PRODUCTION,\n+        \"production\": RuntimeEnvironment.PRODUCTION,\n+    }\n+    if override:\n+        try:\n+            return aliases[override]\n+        except KeyError as error:\n+            valid = \", \".join(sorted(aliases))\n+            raise ValueError(f\"APP_ENV must be one of: {valid}\") from error\n+\n+    if any(_is_truthy(values.get(name)) for name in _CI_MARKERS):\n+        return RuntimeEnvironment.CI\n+\n+    has_workload_identity = all(\n+        values.get(name)\n+        for name in (\n+            \"AZURE_TENANT_ID\",\n+            \"AZURE_CLIENT_ID\",\n+            \"AZURE_FEDERATED_TOKEN_FILE\",\n+        )\n+    )\n+    has_azure_host = any(values.get(name) for name in _AZURE_HOST_MARKERS)\n+    has_kubernetes_host = bool(values.get(\"KUBERNETES_SERVICE_HOST\"))\n+    if has_azure_host or has_workload_identity or has_kubernetes_host:\n+        return RuntimeEnvironment.PRODUCTION\n+\n+    return RuntimeEnvironment.DEV\n*** Add File: connectivity_tester.py\n+\"\"\"Synchronous and asynchronous Azure credential connectivity tests.\"\"\"\n+\n+from __future__ import annotations\n+\n+from dataclasses import dataclass\n+from datetime import datetime, timezone\n+import re\n+\n+from azure.core.credentials import TokenCredential\n+from azure.core.credentials_async import AsyncTokenCredential\n+from azure.core.exceptions import ClientAuthenticationError\n+from azure.identity import CredentialUnavailableError\n+\n+\n+@dataclass(frozen=True)\n+class ConnectivityResult:\n+    succeeded: bool\n+    scope: str\n+    cae_requested: bool\n+    expires_on: datetime | None = None\n+    failure_category: str | None = None\n+    failure_detail: str | None = None\n+\n+\n+_FAILURE_PATTERNS = (\n+    (\n+        r\"AADSTS7000222|certificate.+expired|expired.+certificate|client secret.+expired\",\n+        \"expired client certificate or secret\",\n+    ),\n+    (\n+        r\"AADSTS90002|tenant.+(?:not found|invalid)|invalid tenant\",\n+        \"wrong or unavailable tenant\",\n+    ),\n+    (\n+        r\"AADSTS700016|application.+not found|unauthorized_client\",\n+        \"client ID is unknown in the selected tenant\",\n+    ),\n+    (\n+        r\"AADSTS7000215|invalid client secret|invalid_client\",\n+        \"invalid client secret or client credential\",\n+    ),\n+    (\n+        r\"AADSTS700024|client assertion.+(?:expired|not within)\",\n+        \"expired or not-yet-valid federated assertion\",\n+    ),\n+    (\n+        r\"AADSTS700027|certificate.+(?:invalid|not registered)|thumbprint\",\n+        \"client certificate is invalid or not registered\",\n+    ),\n+    (\n+        r\"AADSTS70021|federated identity credential|subject.+issuer.+audience\",\n+        \"workload identity federation configuration mismatch\",\n+    ),\n+    (\n+        r\"AADSTS70011|invalid_scope|scope.+invalid\",\n+        \"invalid or unauthorized token scope\",\n+    ),\n+    (\n+        r\"managed identity|IMDS|identity endpoint|no identity|unavailable\",\n+        \"no usable managed or workload identity is available\",\n+    ),\n+    (\n+        r\"Azure CLI|Azure Developer CLI|Azure PowerShell|Visual Studio Code\",\n+        \"no authenticated developer-tool account is available\",\n+    ),\n+)\n+\n+\n+def _diagnose(error: BaseException) -\u003e tuple[str, str]:\n+    message = getattr(error, \"message\", None) or str(error)\n+    compact_message = re.sub(r\"\\s+\", \" \", message).strip()\n+    if isinstance(error, CredentialUnavailableError):\n+        category = \"credential unavailable\"\n+    else:\n+        category = \"authentication rejected\"\n+\n+    for pattern, specific_category in _FAILURE_PATTERNS:\n+        if re.search(pattern, compact_message, flags=re.IGNORECASE):\n+            category = specific_category\n+            break\n+\n+    detail = compact_message[:800] or error.__class__.__name__\n+    return category, detail\n+\n+\n+def _print_result(label: str, result: ConnectivityResult) -\u003e None:\n+    cae = \"yes\" if result.cae_requested else \"no\"\n+    if result.succeeded:\n+        expiry = result.expires_on.isoformat() if result.expires_on else \"unknown\"\n+        print(f\"[{label}] SUCCESS\")\n+        print(f\"[{label}] Token expiry (UTC): {expiry}\")\n+        print(f\"[{label}] CAE requested: {cae}\")\n+        return\n+\n+    print(f\"[{label}] FAILURE\")\n+    print(f\"[{label}] Reason: {result.failure_category}\")\n+    print(f\"[{label}] Detail: {result.failure_detail}\")\n+    print(f\"[{label}] CAE requested: {cae}\")\n+\n+\n+def test_credential_sync(\n+    credential: TokenCredential,\n+    scope: str,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e ConnectivityResult:\n+    \"\"\"Request a token synchronously and report a diagnostic result.\"\"\"\n+    try:\n+        token = credential.get_token(scope, enable_cae=enable_cae)\n+        result = ConnectivityResult(\n+            succeeded=True,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            expires_on=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),\n+        )\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        category, detail = _diagnose(error)\n+        result = ConnectivityResult(\n+            succeeded=False,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            failure_category=category,\n+            failure_detail=detail,\n+        )\n+    except Exception as error:\n+        category, detail = _diagnose(error)\n+        result = ConnectivityResult(\n+            succeeded=False,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            failure_category=f\"unexpected {category}\",\n+            failure_detail=detail,\n+        )\n+\n+    _print_result(\"sync\", result)\n+    return result\n+\n+\n+async def test_credential_async(\n+    credential: AsyncTokenCredential,\n+    scope: str,\n+    *,\n+    enable_cae: bool = False,\n+) -\u003e ConnectivityResult:\n+    \"\"\"Request a token asynchronously and report a diagnostic result.\"\"\"\n+    try:\n+        token = await credential.get_token(scope, enable_cae=enable_cae)\n+        result = ConnectivityResult(\n+            succeeded=True,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            expires_on=datetime.fromtimestamp(token.expires_on, tz=timezone.utc),\n+        )\n+    except (CredentialUnavailableError, ClientAuthenticationError) as error:\n+        category, detail = _diagnose(error)\n+        result = ConnectivityResult(\n+            succeeded=False,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            failure_category=category,\n+            failure_detail=detail,\n+        )\n+    except Exception as error:\n+        category, detail = _diagnose(error)\n+        result = ConnectivityResult(\n+            succeeded=False,\n+            scope=scope,\n+            cae_requested=enable_cae,\n+            failure_category=f\"unexpected {category}\",\n+            failure_detail=detail,\n+        )\n+\n+    _print_result(\"async\", result)\n+    return result\n*** Add File: main.py\n+\"\"\"Run synchronous and asynchronous Azure credential connectivity tests.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import os\n+\n+from connectivity_tester import test_credential_async, test_credential_sync\n+from credential_factory import build_async_credential, build_sync_credential\n+from environment_detector import detect_environment\n+\n+\n+ARM_SCOPE = \"https://management.azure.com/.default\"\n+\n+\n+def _cae_enabled() -\u003e bool:\n+    value = os.getenv(\"AZURE_ENABLE_CAE\", \"false\").strip().lower()\n+    if value in {\"1\", \"true\", \"yes\", \"on\"}:\n+        return True\n+    if value in {\"0\", \"false\", \"no\", \"off\"}:\n+        return False\n+    raise ValueError(\n+        \"AZURE_ENABLE_CAE must be one of: true, false, 1, 0, yes, no, on, off\"\n+    )\n+\n+\n+async def _run_async(environment, enable_cae: bool) -\u003e None:\n+    selection = build_async_credential(environment, enable_cae=enable_cae)\n+    print(f\"\\nAsync credential strategy: {selection.strategy}\")\n+    try:\n+        await test_credential_async(\n+            selection.credential,\n+            ARM_SCOPE,\n+            enable_cae=selection.enable_cae,\n+        )\n+    finally:\n+        await selection.credential.close()\n+\n+\n+def main() -\u003e None:\n+    environment = detect_environment()\n+    enable_cae = _cae_enabled()\n+    print(f\"Detected environment: {environment.value}\")\n+    print(f\"CAE token requests enabled: {'yes' if enable_cae else 'no'}\")\n+\n+    selection = build_sync_credential(environment, enable_cae=enable_cae)\n+    print(f\"\\nSync credential strategy: {selection.strategy}\")\n+    try:\n+        test_credential_sync(\n+            selection.credential,\n+            ARM_SCOPE,\n+            enable_cae=selection.enable_cae,\n+        )\n+    finally:\n+        selection.credential.close()\n+\n+    asyncio.run(_run_async(environment, enable_cae))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: tests/test_environment_detector.py\n+import unittest\n+\n+from environment_detector import RuntimeEnvironment, detect_environment\n+\n+\n+class EnvironmentDetectorTests(unittest.TestCase):\n+    def test_defaults_to_dev(self):\n+        self.assertEqual(detect_environment({}), RuntimeEnvironment.DEV)\n+\n+    def test_detects_ci_before_azure_host(self):\n+        environment = detect_environment(\n+            {\"TF_BUILD\": \"True\", \"IDENTITY_ENDPOINT\": \"http://localhost\"}\n+        )\n+        self.assertEqual(environment, RuntimeEnvironment.CI)\n+\n+    def test_detects_managed_identity_host_as_production(self):\n+        self.assertEqual(\n+            detect_environment({\"IDENTITY_ENDPOINT\": \"http://localhost\"}),\n+            RuntimeEnvironment.PRODUCTION,\n+        )\n+\n+    def test_detects_workload_identity_as_production(self):\n+        environment = detect_environment(\n+            {\n+                \"AZURE_TENANT_ID\": \"tenant\",\n+                \"AZURE_CLIENT_ID\": \"client\",\n+                \"AZURE_FEDERATED_TOKEN_FILE\": \"token-file\",\n+            }\n+        )\n+        self.assertEqual(environment, RuntimeEnvironment.PRODUCTION)\n+\n+    def test_explicit_override_wins(self):\n+        self.assertEqual(\n+            detect_environment({\"APP_ENV\": \"local\", \"CI\": \"true\"}),\n+            RuntimeEnvironment.DEV,\n+        )\n+\n+    def test_invalid_override_is_rejected(self):\n+        with self.assertRaisesRegex(ValueError, \"APP_ENV must be one of\"):\n+            detect_environment({\"APP_ENV\": \"staging\"})\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** Add File: tests/test_connectivity_tester.py\n+import asyncio\n+import time\n+import unittest\n+\n+from azure.core.credentials import AccessToken\n+from azure.core.exceptions import ClientAuthenticationError\n+\n+from connectivity_tester import test_credential_async, test_credential_sync\n+\n+\n+class SuccessfulSyncCredential:\n+    def get_token(self, *scopes, **kwargs):\n+        return AccessToken(\"not-logged\", int(time.time()) + 3600)\n+\n+\n+class FailedSyncCredential:\n+    def get_token(self, *scopes, **kwargs):\n+        raise ClientAuthenticationError(\n+            message=\"AADSTS7000222: The provided client secret keys are expired.\"\n+        )\n+\n+\n+class SuccessfulAsyncCredential:\n+    async def get_token(self, *scopes, **kwargs):\n+        return AccessToken(\"not-logged\", int(time.time()) + 3600)\n+\n+\n+class ConnectivityTesterTests(unittest.TestCase):\n+    def test_sync_success_records_cae_and_expiry(self):\n+        result = test_credential_sync(\n+            SuccessfulSyncCredential(), \"scope\", enable_cae=True\n+        )\n+        self.assertTrue(result.succeeded)\n+        self.assertTrue(result.cae_requested)\n+        self.assertIsNotNone(result.expires_on)\n+\n+    def test_sync_failure_reports_specific_reason(self):\n+        result = test_credential_sync(FailedSyncCredential(), \"scope\")\n+        self.assertFalse(result.succeeded)\n+        self.assertEqual(\n+            result.failure_category, \"expired client certificate or secret\"\n+        )\n+\n+    def test_async_success_records_cae_and_expiry(self):\n+        result = asyncio.run(\n+            test_credential_async(\n+                SuccessfulAsyncCredential(), \"scope\", enable_cae=True\n+            )\n+        )\n+        self.assertTrue(result.succeeded)\n+        self.assertTrue(result.cae_requested)\n+        self.assertIsNotNone(result.expires_on)\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** Add File: requirements.txt\n+azure-identity\u003e=1.24.0,\u003c2.0.0\n+azure-identity-broker\u003e=1.3.0,\u003c2.0.0\n*** Add File: README.md\n+# Environment-specific Azure credential chains\n+\n+This sample builds explicit Azure Identity chains instead of using one broad\n+credential everywhere. It only requests an Azure Resource Manager token; it\n+does not create or modify Azure resources.\n+\n+## Setup and run\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+python main.py\n+```\n+\n+Set `AZURE_ENABLE_CAE=true` to request a CAE-enabled token. Azure Resource\n+Manager and the selected identity flow must support CAE for the resulting token\n+to have CAE semantics.\n+\n+Use `APP_ENV=dev`, `APP_ENV=ci`, or `APP_ENV=production` to override automatic\n+detection.\n+\n+## Credential configuration\n+\n+| Environment | Credential order | Configuration |\n+| --- | --- | --- |\n+| Development | VS Code, Azure CLI, Azure PowerShell, Azure Developer CLI | Sign in with the corresponding developer tool. VS Code authentication requires the Azure Resources extension and `azure-identity-broker`. |\n+| CI | Azure Pipelines service connection, then environment credential | For Azure Pipelines, expose `SYSTEM_ACCESSTOKEN`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_SERVICE_CONNECTION_ID` as secret pipeline variables. Other CI systems can configure the environment variables supported by `EnvironmentCredential`. |\n+| Production | Managed identity, then configured Kubernetes workload identity | Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` for a user-assigned managed identity; omit it for system-assigned. Workload identity uses `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE`. |\n+\n+Run the offline tests with:\n+\n+```powershell\n+python -m unittest discover -s tests -v\n+```\n+\n+References:\n+\n+- https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains\n+- https://learn.microsoft.com/python/api/azure-identity/azure.identity.azurepipelinescredential\n+- https://learn.microsoft.com/python/api/azure-identity/azure.identity.workloadidentitycredential\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py
@@ -1,0 +1,189 @@
+"""Environment-specific Azure credential chains."""
+
+from __future__ import annotations
+
+import os
+from dataclasses import dataclass
+
+from azure.core.credentials import TokenCredential
+from azure.core.credentials_async import AsyncTokenCredential
+from azure.identity import (
+    AzureCliCredential,
+    AzureDeveloperCliCredential,
+    AzurePipelinesCredential,
+    AzurePowerShellCredential,
+    ChainedTokenCredential,
+    EnvironmentCredential,
+    ManagedIdentityCredential,
+    VisualStudioCodeCredential,
+    WorkloadIdentityCredential,
+)
+from azure.identity.aio import (
+    AzureCliCredential as AsyncAzureCliCredential,
+)
+from azure.identity.aio import (
+    AzureDeveloperCliCredential as AsyncAzureDeveloperCliCredential,
+)
+from azure.identity.aio import (
+    AzurePipelinesCredential as AsyncAzurePipelinesCredential,
+)
+from azure.identity.aio import (
+    AzurePowerShellCredential as AsyncAzurePowerShellCredential,
+)
+from azure.identity.aio import (
+    ChainedTokenCredential as AsyncChainedTokenCredential,
+)
+from azure.identity.aio import (
+    EnvironmentCredential as AsyncEnvironmentCredential,
+)
+from azure.identity.aio import (
+    ManagedIdentityCredential as AsyncManagedIdentityCredential,
+)
+from azure.identity.aio import (
+    VisualStudioCodeCredential as AsyncVisualStudioCodeCredential,
+)
+from azure.identity.aio import (
+    WorkloadIdentityCredential as AsyncWorkloadIdentityCred
... (truncated)
```

### ✅ 12. powershell — 9847ms

**Input:**

```json
{"command":"python -m pip install --disable-pip-version-check --quiet -r requirements.txt \u0026\u0026 python -m unittest discover -s tests -v \u0026\u0026 python -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py","description":"Install dependencies and run offline tests","initial_wait":180,"mode":"sync"}
```

**Output:**

```
test_async_success_records_cae_and_expiry (test_connectivity_tester.ConnectivityTesterTests.test_async_success_records_cae_and_expiry) ... ok
test_sync_failure_reports_specific_reason (test_connectivity_tester.ConnectivityTesterTests.test_sync_failure_reports_specific_reason) ... ok
test_sync_success_records_cae_and_expiry (test_connectivity_tester.ConnectivityTesterTests.test_sync_success_records_cae_and_expiry) ... ok
test_defaults_to_dev (test_environment_detector.EnvironmentDetectorTests.test_defaults_to_dev) ... ok
test_detects_ci_before_azure_host (test_environment_detector.EnvironmentDetectorTests.test_detects_ci_before_azure_host) ... ok
test_detects_managed_identity_host_as_production (test_environment_detector.EnvironmentDetectorTests.test_detects_managed_identity_host_as_production) ... ok
test_detects_workload_identity_as_production (test_environment_detector.EnvironmentDetectorTests.test_detects_workload_identity_as_production) ... ok
test_explicit_override_wins (test_environment_detector.EnvironmentDetectorTests.test_explicit_override_wins) ... ok
test_invalid_override_is_rejected (test_environment_detector.EnvironmentDetectorTests.test_invalid_override_is_rejected) ... ok

----------------------------------------------------------------------
Ran 9 tests in 0.002s

OK
[async] SUCCESS
[async] Token expiry (UTC): 2026-08-28T20:45:10+00:00
[async] CAE requested: yes
[sync] FAILURE
[sync] Reason: expired client certificate or secret
[sync] Detail: AADSTS7000222: The provided client secret keys are expired.
[sync] CAE requested: no
[sync] SUCCESS
[sync] Token expiry (UTC): 2026-08-28T20:45:10+00:00
[sync] CAE requested: yes
<shellId: 0 completed with exit code 0>
```

### ✅ 13. powershell — 2682ms

**Input:**

```json
{"command":"@'\nimport asyncio\nimport os\nfrom unittest.mock import patch\n\nfrom credential_factory import build_async_credential, build_sync_credential\nfrom environment_detector import RuntimeEnvironment\n\nscenarios = [\n    (RuntimeEnvironment.DEV, {}),\n    (RuntimeEnvironment.CI, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"client\", \"AZURE_CLIENT_SECRET\": \"secret\"}),\n    (RuntimeEnvironment.CI, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"client\", \"AZURE_SERVICE_CONNECTION_ID\": \"connection\", \"SYSTEM_ACCESSTOKEN\": \"token\"}),\n    (RuntimeEnvironment.PRODUCTION, {}),\n    (RuntimeEnvironment.PRODUCTION, {\"AZURE_MANAGED_IDENTITY_CLIENT_ID\": \"managed-client\"}),\n    (RuntimeEnvironment.PRODUCTION, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"workload-client\", \"AZURE_FEDERATED_TOKEN_FILE\": \"token-file\"}),\n]\n\nasync def check_async(environment):\n    selection = build_async_credential(environment, enable_cae=True)\n    assert selection.enable_cae\n    await selection.credential.close()\n\nfor environment, values in scenarios:\n    with patch.dict(os.environ, values, clear=True):\n        selection = build_sync_credential(environment, enable_cae=True)\n        assert selection.strategy and selection.enable_cae\n        selection.credential.close()\n        asyncio.run(check_async(environment))\n\nprint(f\"Constructed and closed sync/async credentials for {len(scenarios)} scenarios.\")\n'@ | python -","description":"Validate credential factory scenarios","initial_wait":120,"mode":"sync"}
```

**Output:**

```
Incomplete environment configuration for EnvironmentCredential. These variables are set: AZURE_TENANT_ID, AZURE_CLIENT_ID
Incomplete environment configuration for EnvironmentCredential. These variables are set: AZURE_TENANT_ID, AZURE_CLIENT_ID
Constructed and closed sync/async credentials for 6 scenarios.
<shellId: 1 completed with exit code 0>
```

### ✅ 14. apply_patch — 727ms

**Input:**

```json
"*** Begin Patch\n*** Update File: credential_factory.py\n@@\n     elif environment is RuntimeEnvironment.CI:\n-        credentials: list[TokenCredential] = []\n         pipeline_settings = _pipeline_settings()\n         if pipeline_settings:\n-            credentials.append(AzurePipelinesCredential(**pipeline_settings))\n-        credentials.append(EnvironmentCredential())\n-        credential = ChainedTokenCredential(*credentials)\n-        strategy = (\n-            \"Azure Pipelines service connection -\u003e environment credential\"\n-            if pipeline_settings\n-            else \"environment credential (service principal or federated token variables)\"\n-        )\n+            credential = ChainedTokenCredential(\n+                AzurePipelinesCredential(**pipeline_settings)\n+            )\n+            strategy = \"Azure Pipelines workload identity service connection\"\n+        else:\n+            credential = ChainedTokenCredential(EnvironmentCredential())\n+            strategy = \"environment credential (service principal variables)\"\n@@\n     elif environment is RuntimeEnvironment.CI:\n-        credentials: list[AsyncTokenCredential] = []\n         pipeline_settings = _pipeline_settings()\n         if pipeline_settings:\n-            credentials.append(AsyncAzurePipelinesCredential(**pipeline_settings))\n-        credentials.append(AsyncEnvironmentCredential())\n-        credential = AsyncChainedTokenCredential(*credentials)\n-        strategy = (\n-            \"Azure Pipelines service connection -\u003e environment credential\"\n-            if pipeline_settings\n-            else \"environment credential (service principal or federated token variables)\"\n-        )\n+            credential = AsyncChainedTokenCredential(\n+                AsyncAzurePipelinesCredential(**pipeline_settings)\n+            )\n+            strategy = \"Azure Pipelines workload identity service connection\"\n+        else:\n+            credential = AsyncChainedTokenCredential(AsyncEnvironmentCredential())\n+            strategy = \"environment credential (service principal variables)\"\n*** Update File: README.md\n@@\n-| CI | Azure Pipelines service connection, then environment credential | For Azure Pipelines, expose `SYSTEM_ACCESSTOKEN`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_SERVICE_CONNECTION_ID` as secret pipeline variables. Other CI systems can configure the environment variables supported by `EnvironmentCredential`. |\n+| CI | Azure Pipelines service connection when fully configured; otherwise environment credential | For Azure Pipelines, expose `SYSTEM_ACCESSTOKEN`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_SERVICE_CONNECTION_ID` as secret pipeline variables. Other CI systems can configure the service principal variables supported by `EnvironmentCredential`. |\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-identity-dp-python-credential-chain-python-azure-skills-azure-skill-mcp-3621057050/credential_factory.py
@@ -105,17 +105,15 @@
         )
         strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
     elif environment is RuntimeEnvironment.CI:
-        credentials: list[TokenCredential] = []
         pipeline_settings = _pipeline_settings()
         if pipeline_settings:
-            credentials.append(AzurePipelinesCredential(**pipeline_settings))
-        credentials.append(EnvironmentCredential())
-        credential = ChainedTokenCredential(*credentials)
-        strategy = (
-            "Azure Pipelines service connection -> environment credential"
-            if pipeline_settings
-            else "environment credential (service principal or federated token variables)"
-        )
+            credential = ChainedTokenCredential(
+                AzurePipelinesCredential(**pipeline_settings)
+            )
+            strategy = "Azure Pipelines workload identity service connection"
+        else:
+            credential = ChainedTokenCredential(EnvironmentCredential())
+            strategy = "environment credential (service principal variables)"
     else:
         managed_identity_client_id = _managed_identity_client_id()
         credentials = [
@@ -154,17 +152,15 @@
         )
         strategy = "developer tools: VS Code -> Azure CLI -> Azure PowerShell -> Azure Developer CLI"
   
... (truncated)
```

### ✅ 15. powershell — 3336ms

**Input:**

```json
{"command":"python -m unittest discover -s tests -v \u0026\u0026 @'\nimport asyncio\nimport os\nfrom unittest.mock import patch\nfrom credential_factory import build_async_credential, build_sync_credential\nfrom environment_detector import RuntimeEnvironment\n\nscenarios = [\n    (RuntimeEnvironment.DEV, {}),\n    (RuntimeEnvironment.CI, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"client\", \"AZURE_CLIENT_SECRET\": \"secret\"}),\n    (RuntimeEnvironment.CI, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"client\", \"AZURE_SERVICE_CONNECTION_ID\": \"connection\", \"SYSTEM_ACCESSTOKEN\": \"token\"}),\n    (RuntimeEnvironment.PRODUCTION, {}),\n    (RuntimeEnvironment.PRODUCTION, {\"AZURE_MANAGED_IDENTITY_CLIENT_ID\": \"managed-client\"}),\n    (RuntimeEnvironment.PRODUCTION, {\"AZURE_TENANT_ID\": \"tenant\", \"AZURE_CLIENT_ID\": \"workload-client\", \"AZURE_FEDERATED_TOKEN_FILE\": \"token-file\"}),\n]\n\nasync def check_async(environment):\n    selection = build_async_credential(environment, enable_cae=True)\n    assert selection.strategy and selection.enable_cae\n    await selection.credential.close()\n\nfor environment, values in scenarios:\n    with patch.dict(os.environ, values, clear=True):\n        selection = build_sync_credential(environment, enable_cae=True)\n        assert selection.strategy and selection.enable_cae\n        selection.credential.close()\n        asyncio.run(check_async(environment))\nprint(f\"Validated {len(scenarios)} sync/async factory scenarios.\")\n'@ | python - \u0026\u0026 python -m py_compile credential_factory.py environment_detector.py connectivity_tester.py main.py","description":"Re-run offline validation","initial_wait":120,"mode":"sync"}
```

**Output:**

```
test_async_success_records_cae_and_expiry (test_connectivity_tester.ConnectivityTesterTests.test_async_success_records_cae_and_expiry) ... ok
test_sync_failure_reports_specific_reason (test_connectivity_tester.ConnectivityTesterTests.test_sync_failure_reports_specific_reason) ... ok
test_sync_success_records_cae_and_expiry (test_connectivity_tester.ConnectivityTesterTests.test_sync_success_records_cae_and_expiry) ... ok
test_defaults_to_dev (test_environment_detector.EnvironmentDetectorTests.test_defaults_to_dev) ... ok
test_detects_ci_before_azure_host (test_environment_detector.EnvironmentDetectorTests.test_detects_ci_before_azure_host) ... ok
test_detects_managed_identity_host_as_production (test_environment_detector.EnvironmentDetectorTests.test_detects_managed_identity_host_as_production) ... ok
test_detects_workload_identity_as_production (test_environment_detector.EnvironmentDetectorTests.test_detects_workload_identity_as_production) ... ok
test_explicit_override_wins (test_environment_detector.EnvironmentDetectorTests.test_explicit_override_wins) ... ok
test_invalid_override_is_rejected (test_environment_detector.EnvironmentDetectorTests.test_invalid_override_is_rejected) ... ok

----------------------------------------------------------------------
Ran 9 tests in 0.002s

OK
[async] SUCCESS
[async] Token expiry (UTC): 2026-08-28T20:45:46+00:00
[async] CAE requested: yes
[sync] FAILURE
[sync] Reason: expired client certificate or secret
[sync] Detail: AADSTS7000222: The provided client secret keys are expired.
[sync] CAE requested: no
[sync] SUCCESS
[sync] Token expiry (UTC): 2026-08-28T20:45:46+00:00
[sync] CAE requested: yes
Validated 6 sync/async factory scenarios.
<shellId: 2 completed with exit code 0>
```

## Generated Files

- `README.md`
- `connectivity_tester.py`
- `credential_factory.py`
- `environment_detector.py`
- `main.py`
- `requirements.txt`
- `tests\test_connectivity_tester.py`
- `tests\test_environment_detector.py`

## Copilot Response

I’ll establish the Azure Identity patterns, then scaffold the local-only sample and verify its behavior without contacting Azure.

The design is now fixed around explicit ordered chains: developer-tool credentials for local use, service-connection/environment credentials for CI, and managed identity before workload identity in production. I’m checking the current API details before writing the modules.

I’m implementing the sample as importable modules plus offline unit tests, with no token values logged and no live Azure resource operations. The production chain only activates workload identity when its three Kubernetes federation settings are complete.

The project files are complete. I’m validating imports and the offline detector/connectivity behavior now; the validation uses fake credentials and will not request a real Azure token.

Construction exposed an avoidable Azure Identity warning when service-connection variables were present: `EnvironmentCredential` interpreted the same tenant/client variables as an incomplete secret-based configuration. I’m making the CI modes mutually exclusive so service connections and environment credentials remain deterministic and warning-free.

Created the Python project with:

- Environment-specific sync and async Azure credential factories
- Automatic dev/CI/production detection
- Developer-tool, pipeline, managed identity, and workload identity strategies
- Optional CAE requests via `AZURE_ENABLE_CAE`
- Detailed authentication failure diagnostics
- Azure Resource Manager connectivity runner
- Offline unit tests covering detection, CAE, success, and failure reporting
- `requirements.txt` and setup documentation in `README.md`

Run with:

```powershell
python main.py
```

Enable CAE with `$env:AZURE_ENABLE_CAE = "true"` before running.

## Reviewer Notes (LLM-as-Judge)

### Summary

0 of 1 criteria passed.

### Strengths

- ClientAuthenticationError and CredentialUnavailableError receive specific diagnostic handling in both synchronous and asynchronous testers.

### Issues

- HttpResponseError and its non-ClientAuthenticationError subclasses are only handled by a generic Exception catch.

## Grader Results

- credential-chain.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (14/14)
      - Uses `ChainedTokenCredential` to compose multiple credentials: Pass
      - Credentials passed in order to the constructor — order matters: Pass
      - **Dev chain**: includes `AzureCliCredential`; may include `VisualStudioCodeCredential`, `AzurePowerShellCredential`: Pass
      - **CI chain**: uses `EnvironmentCredential` or `AzurePipelinesCredential` (not just `DefaultAzureCredential`): Pass
      - **Production chain**: `ManagedIdentityCredential` first (supports user-assigned via `client_id` kwarg), `WorkloadIdentityCredential` as fallback: Pass
      - Enables CAE via `enable_cae=True` in `get_token()` call or token request context: Pass
      - Detects CI (checks `CI`, `TF_BUILD`, `BUILD_SOURCESDIRECTORY`, or similar env vars): Pass
      - Detects production/managed identity (checks `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, or similar): Pass
      - Falls back to dev if neither detected: Pass
      - Calls `get_token()` with correct scope (`https://management.azure.com/.default`): Pass
      - Prints token expiry from `AccessToken.expires_on`: Pass
      - Handles failure with specific exception info from `ClientAuthenticationError`: Pass
      - Async tester uses `azure.identity.aio` async credentials and `await get_token()`: Pass
      - NOT using `DefaultAzureCredential` as the CI credential (too broad): Pass
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
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 100% | 1.00 | 1.0000 | 16.7% | ✅ |
| **Final** | | | **Σ 8.00** | **Σ 6.0000** | **75.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-python-credential-chain --config python-azure-skills/azure-skill-mcp --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
