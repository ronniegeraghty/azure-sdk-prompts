# Evaluation Report: ai-projects-dp-dotnet-project-resource-inventory

**Config:** dotnet-azure-tools/baseline | **Result:** ✅ PASSED | **Duration:** 261.1s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `ai-projects-dp-dotnet-project-resource-inventory` |
| Config | dotnet-azure-tools/baseline |
| Result | ✅ PASSED |
| Score | 9/9 |
| Duration | 261.1s |
| Timestamp | 2026-08-26T15:00:16Z |
| Files Generated | 3 |
| Event Count | 3046 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 145.0s |
| Review | 114.7s |
| **Total** | **261.1s** |

## Configuration

- **name:** dotnet-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=342862 out=3057 |
| Turn Count | 11 |

## Tool Calls

### ✅ 1. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"\"GetConnectionsAsync\" \"Azure.AI.Projects\" language:C#"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/src/Custom/AIProjectDeploymentsOperations.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"        /// <returns> The response returned from the service. </returns>\n        [EditorBrowsable(EditorBrowsableState.Never)]\n        [Obsolete(\"This method is obsolete as the clientRequestId parameter is not used. Please use GetDeploymentsAsync(string modelPublisher, string modelName, string deploymentType, RequestOptions options) instead.\")]\n        public virtual AsyncCollectionResult GetDeploymentsAsync(string modelPublisher, string modelName, string deploymentType, string clientRequestId, RequestOptions options)\n        {\n            return GetDeploymentsAsync(modelPublisher, modelName, deploymentType, options);\n        }","matches":[{"indices":[227,246],"text":"GetDeploymentsAsync"},{"indices":[392,411],"text":"GetDeploymentsAsync"},{"indices":[553,572],"text":"GetDeploymentsAsync"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Custom/AIProjectDeploymentsOperations.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/AIProjectDeploymentsOperations.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"        /// <param name=\"options\"> The request options, which can override default behaviors of the client pipeline on a per-call basis. </param>\n        /// <exception cref=\"ClientResultException\"> Service returned a non-success status code. </exception>\n        /// <returns> The response returned from the service. </returns>\n        public virtual AsyncCollectionResult GetDeploymentsAsync(string modelPublisher, string modelName, string deploymentType, RequestOptions options)\n        {\n            using DiagnosticScope scope = ClientDiagnostics.CreateScope(\"AIProjectDeploymentsOperations.GetDeployment
... (truncated)
```

### ✅ 3. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","repository","text_matches"],"perPage":10,"query":"\"GetDeploymentsAsync\" \"Azure.AI.Projects\" language:C#"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/src/Custom/AIProjectConnectionsOperations.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"        /// <returns> The response returned from the service. </returns>\n        [EditorBrowsable(EditorBrowsableState.Never)]\n        [Obsolete(\"This method is obsolete as the clientRequestId parameter is not used. Please use GetConnectionsAsync(string connectionType, bool? defaultConnection, RequestOptions options) instead.\")]\n        public virtual AsyncCollectionResult GetConnectionsAsync(string connectionType, bool? defaultConnection, string clientRequestId, RequestOptions options)\n        {\n            return GetConnectionsAsync(connectionType, defaultConnection, options);\n        }","matches":[{"indices":[227,246],"text":"GetConnectionsAsync"},{"indices":[376,395],"text":"GetConnectionsAsync"},{"indices":[521,540],"text":"GetConnectionsAsync"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Custom/AIProjectConnectionsOperations.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/AIProjectConnectionsOperations.cs","repository":"Azure/azure-sdk-for-net","text_matches":[{"fragment":"        /// <param name=\"options\"> The request options, which can override default behaviors of the client pipeline on a per-call basis. </param>\n        /// <exception cref=\"ClientResultException\"> Service returned a non-success status code. </exception>\n        /// <returns> The response returned from the service. </returns>\n        public virtual AsyncCollectionResult GetConnectionsAsync(string connectionType, bool? defaultConnection, RequestOptions options)\n        {\n            using DiagnosticScope scope = ClientDiagnostics.CreateScope(\"AIProjectConnectionsOperations.GetConnections\");\n            scope.Start();","matches":[{"indices"
... (truncated)
```

### ✅ 4. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/README.md","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 638c747f39c054a61ae7f85670d39f6d512adc3e)# Azure AI Projects client library for .NET
The AI Projects client library is part of the Azure AI Foundry SDK and provides easy access to resources in your Azure AI Foundry Project. Use it to:

* **Create and run Classic Agents** using the `GetPersistentAgentsClient` method on the client.
* **Create Agents** using `AgentAdministrationClient` property.
* **Enumerate AI Models** deployed to your Foundry Project using the `Deployments` operations.
* **Enumerate connected Azure resources** in your Foundry project using the `Connections` operations.
* **Upload documents and create Datasets** to reference them using the `Datasets` operations.
* **Create and enumerate Search Indexes** using the `Indexes` operations.

The client library uses version `v1` of the AI Foundry [data plane REST APIs](https://aka.ms/azsdk/azure-ai-projects/ga-rest-api-reference).

[Product documentation][product_doc]
| [Samples][samples]
| [API reference documentation][api_ref_docs]
| [Package (NuGet)][nuget]
| [SDK source code][source_code]

## Table of contents

- [Getting started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Install the package](#install-the-package)
- [Key concepts](#key-concepts)
  - [Create and authenticate the client](#create-and-authenticate-the-client)
- [Examples](#examples)
  - [Performing Agent operations](#performing-agent-operations)
  - [Get an authenticated AzureOpenAI client](#get-an-authenticated-azureopenai-client)
  - [Get an authenticated ChatCompletionsClient](#get-an-authenticated-chatcompletionsclient)
  - [Deployments operations](#deployments-operations)
  - [Connections operations](#connections-operations)
  - [Dataset operations](#dataset-operations)
  - [Indexes operations](#indexes-operations)
  - [Files operations](#files-operations)
  - [Fine-Tuning operations](#fine-tuning-operations)
  - [Memory store operations](#memory-store-operations)
  - [Model weig
... (truncated)
```

### ❌ 5. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/samples/Sample2_Connections.md","repo":"azure-sdk-for-net"}
```

**Error:**

```
MCP server 'github-mcp-server': Failed to get file contents. The path does not point to a file or directory, or the file does not exist in the repository.
```

### ❌ 6. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/samples/Sample3_Deployments.md","repo":"azure-sdk-for-net"}
```

**Error:**

```
MCP server 'github-mcp-server': Failed to get file contents. The path does not point to a file or directory, or the file does not exist in the repository.
```

### ❌ 7. web_fetch

**Input:**

```json
{"max_length":5000,"url":"https://api.nuget.org/v3-flatcontainer/azure.ai.projects/index.json"}
```

**Error:**

```
web_fetch could not retrieve this URL: a transport-level failure (DNS, connection, or TLS handshake).
```

### ✅ 8. powershell

**Input:**

```json
{"command":"dotnet --info | Select-Object -First 25","description":"Check installed .NET SDK"}
```

**Output:**

```
.NET SDK:
 Version:           10.0.111
 Commit:            e2f47b0110
 Workload version:  10.0.100-manifests.330ea142
 MSBuild version:   18.0.11+e2f47b011

Runtime Environment:
 OS Name:     Windows
 OS Version:  10.0.26200
 OS Platform: Windows
 RID:         win-x64
 Base Path:   C:\Program Files\dotnet\sdk\10.0.111\

.NET workloads installed:
There are no installed workloads to display.
Configured to use workload sets when installing new manifests.
No workload sets are installed. Run "dotnet workload restore" to install a workload set.

Host:
  Version:      10.0.11
  Architecture: x64
  Commit:       e2f47b0110

.NET SDKs installed:
  9.0.317 [C:\Program Files\dotnet\sdk]
<shellId: 0 completed with exit code 0>
```

### ✅ 9. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"fields":["name","path","type"],"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/samples","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 4a47a5eff13033d18e6a917963d9864238695cf1)<Project Sdk="Microsoft.NET.Sdk">
  <PropertyGroup>
    <Description>This is the Azure.AI.Projects client library for developing .NET applications with rich experience.</Description>
    <AssemblyTitle>Microsoft Azure.AI.Projects client library</AssemblyTitle>
    <Version>3.0.0-beta.1</Version>
    <!--The ApiCompatVersion is managed automatically and should not generally be modified manually.-->
    <ApiCompatVersion>2.0.0</ApiCompatVersion>
    <PackageTags>Azure.AI.Projects</PackageTags>
    <DisableEnhancedAnalysis>true</DisableEnhancedAnalysis>
    <TargetFrameworks>$(RequiredTargetFrameworks)</TargetFrameworks>
    <GenerateAPIListing>true</GenerateAPIListing>
    <IncludeOperationsSharedSource>true</IncludeOperationsSharedSource>
  </PropertyGroup>

  <PropertyGroup>
    <IncludeSystemClientModelSharedSourceForTracing>true</IncludeSystemClientModelSharedSourceForTracing>
  </PropertyGroup>

  <ItemGroup>
    <PackageReference Include="Azure.Core" />
    <PackageReference Include="Azure.Storage.Blobs" />
    <ProjectReference Include="../../Azure.AI.Extensions.OpenAI/src/Azure.AI.Extensions.OpenAI.csproj" />
    <ProjectReference Include="../../Azure.AI.Projects.Agents/src/Azure.AI.Projects.Agents.csproj" />
  </ItemGroup>

  <!-- Shared source from Azure.Core -->
  <ItemGroup>
    <Compile Include="$(AzureCoreSharedSources)ExperimentalAttribute.cs" LinkBase="Shared" />
    <Compile Include="$(AzureCoreSharedSources)ArrayBufferWriter.cs" LinkBase="Shared" />
    <Compile Include="$(AzureCoreSharedSources)AzureKeyCredentialPolicy.cs" LinkBase="Shared" />
    <Compile Include="$(AzureCoreSharedSources)AzureResourceProviderNamespaceAttribute.cs" LinkBase="Shared" />
    <Compile Include="$(AzureCoreSharedSources)Multipart\MultipartContent.cs" LinkBase="Shared" />
    <Compile Include="$(AzureCoreSharedSources)Multipart\MultipartFormDataContent.cs" LinkBase="Shared" />
  </ItemGroup>
... (truncated)
```

### ✅ 10. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":20,"query":"repo:Azure/azure-sdk-for-net path:sdk/ai/Azure.AI.Projects/samples \"GetConnectionAsync\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectConnection.cs","text_matches":[{"fragment":"        /// <param name=\"credentials\"> The credentials used by the connection. </param>\n        /// <param name=\"metadata\"> Metadata of the connection. </param>\n        /// <param name=\"additionalBinaryDataProperties\"> Keeps track of any properties unknown to the library. </param>\n        internal AIProjectConnection(string name, string id, ConnectionType @type, string target, bool isDefault, AIProjectConnectionBaseCredential credentials, IReadOnlyDictionary<string, string> metadata, IDictionary<string, BinaryData> additionalBinaryDataProperties)\n        {\n            Name = name;\n            Id = id;","matches":[{"indices":[372,378],"text":"target"},{"indices":[385,394],"text":"isDefault"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectConnection.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectConnection.Serialization.cs","text_matches":[{"fragment":"                name,\n                id,\n                @type,\n                target,\n                isDefault,\n                credentials,\n                metadata,","matches":[{"indices":[81,87],"text":"target"},{"indices":[105,114],"text":"isDefault"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectConnection.Serialization.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/AzureAIProjectsModelFactory.cs","text_matches":[{"fragment":"        /// <param name=\"credentials\"> The credentials used by the connection. </param>\n        /// <param name=\"metadata\"> Metadata of the connection. </par
... (truncated)
```

### ✅ 11. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":20,"query":"repo:Azure/azure-sdk-for-net path:sdk/ai/Azure.AI.Projects/samples \"ModelPublisher\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/samples/Sample2_Deployments.md","text_matches":[{"fragment":"var modelDeploymentName = System.Environment.GetEnvironmentVariable(\"FOUNDRY_MODEL_NAME\");\nvar modelPublisher = System.Environment.GetEnvironmentVariable(\"MODEL_PUBLISHER\");\n","matches":[{"indices":[95,109],"text":"modelPublisher"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/samples/Sample2_Deployments.md?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"},{"fragment":"\nConsole.WriteLine($\"List all deployments by the model publisher `{modelPublisher}`:\");\nforeach (AIProjectDeployment deployment in projectClient.Deployments.GetDeployments(modelPublisher: modelPublisher))","matches":[{"indices":[67,81],"text":"modelPublisher"},{"indices":[172,186],"text":"modelPublisher"},{"indices":[188,202],"text":"modelPublisher"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/samples/Sample2_Deployments.md?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]}],"total_count":1}
```

### ✅ 12. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":20,"query":"repo:Azure/azure-sdk-for-net path:sdk/ai/Azure.AI.Projects/src \"IsDefault\" \"Target\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/samples/Sample3_Connections.md","text_matches":[{"fragment":"Console.WriteLine($\"Get the properties of a connection named `{connectionName}`:\");\nAIProjectConnection specificConnection = await projectClient.Connections.GetConnectionAsync(connectionName, includeCredentials: false);\nConsole.WriteLine(specificConnection);","matches":[{"indices":[157,175],"text":"GetConnectionAsync"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/samples/Sample3_Connections.md?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"},{"fragment":"Console.WriteLine(\"Get the properties of a connection with credentials:\");\nAIProjectConnection specificConnectionCredentials = await projectClient.Connections.GetConnectionAsync(connectionName, includeCredentials: true);\nConsole.WriteLine(specificConnectionCredentials);","matches":[{"indices":[159,177],"text":"GetConnectionAsync"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/samples/Sample3_Connections.md?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]}],"total_count":1}
```

### ✅ 13. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/src/Azure.AI.Projects.csproj","repo":"azure-sdk-for-net"}
```

**Output:**

```
[{"name":"Evaluations","path":"sdk/ai/Azure.AI.Projects/samples/Evaluations","type":"dir"},{"name":"FineTuning","path":"sdk/ai/Azure.AI.Projects/samples/FineTuning","type":"dir"},{"name":"README.md","path":"sdk/ai/Azure.AI.Projects/samples/README.md","type":"file"},{"name":"RedTeam","path":"sdk/ai/Azure.AI.Projects/samples/RedTeam","type":"dir"},{"name":"Routines","path":"sdk/ai/Azure.AI.Projects/samples/Routines","type":"dir"},{"name":"Sample11_InferenceImageEmbedding.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample11_InferenceImageEmbedding.md","type":"file"},{"name":"Sample13_DatasetsAsVectorStoreSource.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample13_DatasetsAsVectorStoreSource.md","type":"file"},{"name":"Sample14_IndexCreationForAISearch.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample14_IndexCreationForAISearch.md","type":"file"},{"name":"Sample15_Files.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample15_Files.md","type":"file"},{"name":"Sample16_Agents_CRUD.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample16_Agents_CRUD.md","type":"file"},{"name":"Sample17_Toolboxes_CRUD.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample17_Toolboxes_CRUD.md","type":"file"},{"name":"Sample1_Datasets.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample1_Datasets.md","type":"file"},{"name":"Sample20_MemoryStore.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample20_MemoryStore.md","type":"file"},{"name":"Sample21_MemoryStoreItems.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample21_MemoryStoreItems.md","type":"file"},{"name":"Sample22_Models.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample22_Models.md","type":"file"},{"name":"Sample2_Deployments.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample2_Deployments.md","type":"file"},{"name":"Sample3_Connections.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample3_Connections.md","type":"file"},{"name":"Sample4_Indexes.md","path":"sdk/ai/Azure.AI.Projects/samples/Sample4_Indexes.md","type":"file"},{"name":"Sample6_Cl
... (truncated)
```

### ❌ 14. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/samples/Sample2_Deployments.md","repo":"azure-sdk-for-net"}
```

**Error:**

```
MCP server 'github-mcp-server': Failed to get file contents. The path does not point to a file or directory, or the file does not exist in the repository.
```

### ✅ 15. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/samples/Sample3_Connections.md","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 2d97ceca72f7fe5c5c4e57404c1b20224393dc8c)// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;

namespace Azure.AI.Projects
{
    /// <summary>
    /// Model Deployment Definition
    /// Please note this is the abstract base class. The derived classes available for instantiation are: <see cref="ModelDeployment"/>.
    /// </summary>
    public abstract partial class AIProjectDeployment
    {
        /// <summary> Keeps track of any properties unknown to the library. </summary>
        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;

        /// <summary> Initializes a new instance of <see cref="AIProjectDeployment"/>. </summary>
        /// <param name="type"> The type of the deployment. </param>
        private protected AIProjectDeployment(AIProjectDeploymentType @type)
        {
            Type = @type;
        }

        /// <summary> Initializes a new instance of <see cref="AIProjectDeployment"/>. </summary>
        /// <param name="type"> The type of the deployment. </param>
        /// <param name="name"> Name of the deployment. </param>
        /// <param name="additionalBinaryDataProperties"> Keeps track of any properties unknown to the library. </param>
        internal AIProjectDeployment(AIProjectDeploymentType @type, string name, IDictionary<string, BinaryData> additionalBinaryDataProperties)
        {
            Type = @type;
            Name = name;
            _additionalBinaryDataProperties = additionalBinaryDataProperties;
        }

        /// <summary> The type of the deployment. </summary>
        internal AIProjectDeploymentType Type { get; set; }

        /// <summary> Name of the deployment. </summary>
        public string Name { get; }
    }
}

```

### ✅ 16. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectModelDeployment.cs","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 12f4f337b5fbed2e2caa9167d9c1a8d99cdc1b10)# Sample using Deployments in Azure.AI.Projects

This sample demonstrates how to use the synchronous and asynchronous `.deployments` methods to enumerate AI models deployed to your AI Foundry Project.

## Prerequisites

- Install the Azure.AI.Projects package.
- Set the following environment variables:
  - `PROJECT_ENDPOINT`: The Azure AI Project endpoint, as found in the overview page of your Azure AI Foundry project.
  - `DEPLOYMENT_NAME`: The name of the deployment to retrieve.
  - `MODEL_PUBLISHER`: The publisher of the model to filter by.

## Synchronous Sample

```C# Snippet:AI_Projects_DeploymentExampleSync
var endpoint = System.Environment.GetEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
var modelDeploymentName = System.Environment.GetEnvironmentVariable("FOUNDRY_MODEL_NAME");
var modelPublisher = System.Environment.GetEnvironmentVariable("MODEL_PUBLISHER");

AIProjectClient projectClient = new AIProjectClient(new Uri(endpoint), new DefaultAzureCredential());

Console.WriteLine("List all deployments:");
foreach (AIProjectDeployment deployment in projectClient.Deployments.GetDeployments())
{
    Console.WriteLine(deployment);
}

Console.WriteLine($"List all deployments by the model publisher `{modelPublisher}`:");
foreach (AIProjectDeployment deployment in projectClient.Deployments.GetDeployments(modelPublisher: modelPublisher))
{
    Console.WriteLine(deployment);
}

Console.WriteLine($"Get a single model deployment named `{modelDeploymentName}`:");
ModelDeployment deploymentDetails = (ModelDeployment)projectClient.Deployments.GetDeployment(modelDeploymentName);
Console.WriteLine(deploymentDetails);
```

## Asynchronous Sample

```C# Snippet:AI_Projects_DeploymentExampleAsync
var endpoint = System.Environment.GetEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
var modelDeploymentName = System.Environment.GetEnvironmentVariable("FOUNDRY_MODEL_NAME");
var modelPublisher = Syste
... (truncated)
```

### ✅ 17. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectDeployment.cs","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 9153d9b4252048c9113cab6ef305049fa7bd2607)# Sample using `Connections` in Azure.AI.Projects

In this example, we will demonstrate listing and retrieving connections using the `Connections` client in `Azure.AI.Projects`. This includes listing all connections, filtering by connection type, and retrieving specific connection details.

## Prerequisites

- Install the Azure.AI.Projects package.
- Set the following environment variables:
  - `PROJECT_ENDPOINT`: The Azure AI Project endpoint, as found in the overview page of your Azure AI Foundry project.
  - `CONNECTION_NAME`: The name of the connection to retrieve.

## Synchronous sample:
```C# Snippet:AI_Projects_ConnectionsExampleSync
var endpoint = Environment.GetEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
var connectionName = Environment.GetEnvironmentVariable("CONNECTION_NAME");
AIProjectClient projectClient = new AIProjectClient(new Uri(endpoint), new DefaultAzureCredential());

Console.WriteLine("List the properties of all connections:");
foreach (AIProjectConnection connection in projectClient.Connections.GetConnections())
{
    Console.WriteLine(connection);
    Console.WriteLine(connection.Name);
}

Console.WriteLine("List the properties of all connections of a particular type (e.g., Azure OpenAI connections):");
foreach (AIProjectConnection connection in projectClient.Connections.GetConnections(connectionType: ConnectionType.AzureOpenAI))
{
    Console.WriteLine(connection);
}

Console.WriteLine($"Get the properties of a connection named `{connectionName}`:");
AIProjectConnection specificConnection = projectClient.Connections.GetConnection(connectionName, includeCredentials: false);
Console.WriteLine(specificConnection);

Console.WriteLine("Get the properties of a connection with credentials:");
AIProjectConnection specificConnectionCredentials = projectClient.Connections.GetConnection(connectionName, includeCredentials: true);
Console.WriteLine(specificConnectionCre
... (truncated)
```

### ✅ 18. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net path:sdk/ai/Azure.AI.Projects/src \"GetDeploymentAsync(string\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/src/Custom/AIProjectDeploymentsOperations.cs","text_matches":[{"fragment":"        /// <exception cref=\"ClientResultException\"> Service returned a non-success status code. </exception>\n        /// <returns> The response returned from the service. </returns>\n        [EditorBrowsable(EditorBrowsableState.Never)]\n        [Obsolete(\"This method is obsolete as the clientRequestId parameter is not used. Please use GetDeploymentAsync(string name, RequestOptions options) instead.\")]\n        public virtual async Task<ClientResult> GetDeploymentAsync(string name, string clientRequestId, RequestOptions options)\n        {\n            return await GetDeploymentAsync(name, options).ConfigureAwait(false);","matches":[{"indices":[337,362],"text":"GetDeploymentAsync(string"},{"indices":[453,478],"text":"GetDeploymentAsync(string"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Custom/AIProjectDeploymentsOperations.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/AIProjectDeploymentsOperations.cs","text_matches":[{"fragment":"        /// <returns> The response returned from the service. </returns>\n        public virtual async Task<ClientResult> GetDeploymentAsync(string name, RequestOptions options)\n        {","matches":[{"indices":[121,146],"text":"GetDeploymentAsync(string"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Generated/AIProjectDeploymentsOperations.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"},{"fragment":"        /// <exception cref=\"ClientResultException\"> Service returned a non-success status code. </exception>\n        public virtual async Task<ClientResult<AIProjectDeployment>> GetDeploymentAsync(string name, CancellationToken cance
... (truncated)
```

### ✅ 19. github-mcp-server-search_code (via github-mcp-server)

**Input:**

```json
{"fields":["path","text_matches"],"perPage":10,"query":"repo:Azure/azure-sdk-for-net path:sdk/ai/Azure.AI.Projects/src \"class ModelDeployment\""}
```

**Output:**

```
{"incomplete_results":false,"items":[{"path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeployment.cs","text_matches":[{"fragment":"namespace Azure.AI.Projects\n{\n    /// <summary> Model Deployment Definition. </summary>\n    public partial class ModelDeployment : AIProjectDeployment\n    {\n        /// <summary> Initializes a new instance of <see cref=\"ModelDeployment\"/>. </summary>\n        internal ModelDeployment() : base(AIProjectDeploymentType.ModelDeployment)","matches":[{"indices":[107,128],"text":"class ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeployment.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeployment.Serialization.cs","text_matches":[{"fragment":"namespace Azure.AI.Projects\n{\n    /// <summary> Model Deployment Definition. </summary>\n    public partial class ModelDeployment : AIProjectDeployment, IJsonModel<ModelDeployment>\n    {\n        /// <param name=\"data\"> The data to parse. </param>\n        /// <param name=\"options\"> The client options for reading and writing models. </param>","matches":[{"indices":[107,128],"text":"class ModelDeployment"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeployment.Serialization.cs?ref=9330c33dc33b99dea27cfe2f4a343c5ab3dc4b05","property":"content"}]},{"path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeploymentSku.cs","text_matches":[{"fragment":"namespace Azure.AI.Projects\n{\n    /// <summary> Sku information. </summary>\n    public partial class ModelDeploymentSku\n    {\n        /// <summary> Keeps track of any properties unknown to the library. </summary>\n        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;","matches":[{"in
... (truncated)
```

### ✅ 20. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/AIProjectConnection.cs","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 6afc9671f6d73fae39cbe831a5bee48c0946fc9d)// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;

namespace Azure.AI.Projects
{
    /// <summary> Response from the list and get connections operations. </summary>
    public partial class AIProjectConnection
    {
        /// <summary> Keeps track of any properties unknown to the library. </summary>
        private protected readonly IDictionary<string, BinaryData> _additionalBinaryDataProperties;

        /// <summary> Initializes a new instance of <see cref="AIProjectConnection"/>. </summary>
        internal AIProjectConnection()
        {
            Metadata = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="AIProjectConnection"/>. </summary>
        /// <param name="name"> The friendly name of the connection, provided by the user. </param>
        /// <param name="id"> A unique identifier for the connection, generated by the service. </param>
        /// <param name="type"> Category of the connection. </param>
        /// <param name="target"> The connection URL to be used for this service. </param>
        /// <param name="isDefault"> Whether the connection is tagged as the default connection of its type. </param>
        /// <param name="credentials"> The credentials used by the connection. </param>
        /// <param name="metadata"> Metadata of the connection. </param>
        /// <param name="additionalBinaryDataProperties"> Keeps track of any properties unknown to the library. </param>
        internal AIProjectConnection(string name, string id, ConnectionType @type, string target, bool isDefault, AIProjectConnectionBaseCredential credentials, IReadOnlyDictionary<string, string> metadata, IDictionary<string, BinaryData> additionalBinaryDataProperties)
        {
            Name = name;
            Id = id;
            Type = @type;
            Target = target;
  
... (truncated)
```

### ✅ 21. powershell

**Input:**

```json
{"command":"dotnet nuget list source","description":"Inspect configured NuGet sources"}
```

**Output:**

```
Registered Sources:
  1.  nuget.org [Disabled]
      https://api.nuget.org/v3/index.json
  2.  azure-default [Enabled]
      https://packagefeedproxy.microsoft.io/nuget/v3/index.json
<shellId: 1 completed with exit code 0>
```

### ✅ 22. github-mcp-server-get_file_contents (via github-mcp-server)

**Input:**

```json
{"owner":"Azure","path":"sdk/ai/Azure.AI.Projects/src/Generated/Models/ModelDeployment.cs","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: 78b023cb815e25b3aff6f7c59b62d0a730d394dd)// <auto-generated/>

#nullable disable

using System;
using System.Collections.Generic;

namespace Azure.AI.Projects
{
    /// <summary> Model Deployment Definition. </summary>
    public partial class ModelDeployment : AIProjectDeployment
    {
        /// <summary> Initializes a new instance of <see cref="ModelDeployment"/>. </summary>
        internal ModelDeployment() : base(AIProjectDeploymentType.ModelDeployment)
        {
            Capabilities = new ChangeTrackingDictionary<string, string>();
        }

        /// <summary> Initializes a new instance of <see cref="ModelDeployment"/>. </summary>
        /// <param name="type"> The type of the deployment. </param>
        /// <param name="name"> Name of the deployment. </param>
        /// <param name="additionalBinaryDataProperties"> Keeps track of any properties unknown to the library. </param>
        /// <param name="modelName"> Publisher-specific name of the deployed model. </param>
        /// <param name="modelVersion"> Publisher-specific version of the deployed model. </param>
        /// <param name="modelPublisher"> Name of the deployed model's publisher. </param>
        /// <param name="capabilities"> Capabilities of deployed model. </param>
        /// <param name="sku"> Sku of the model deployment. </param>
        /// <param name="connectionName"> Name of the connection the deployment comes from. </param>
        internal ModelDeployment(AIProjectDeploymentType @type, string name, IDictionary<string, BinaryData> additionalBinaryDataProperties, string modelName, string modelVersion, string modelPublisher, IReadOnlyDictionary<string, string> capabilities, ModelDeploymentSku sku, string connectionName) : base(@type, name, additionalBinaryDataProperties)
        {
            ModelName = modelName;
            ModelVersion = modelVersion;
            ModelPublisher = modelPublisher;
            Capabilities = capabi
... (truncated)
```

### ✅ 23. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: FoundryProjectInventory.csproj\n+\u003cProject Sdk=\"Microsoft.NET.Sdk\"\u003e\n+\n+  \u003cPropertyGroup\u003e\n+    \u003cOutputType\u003eExe\u003c/OutputType\u003e\n+    \u003cTargetFramework\u003enet8.0\u003c/TargetFramework\u003e\n+    \u003cImplicitUsings\u003eenable\u003c/ImplicitUsings\u003e\n+    \u003cNullable\u003eenable\u003c/Nullable\u003e\n+  \u003c/PropertyGroup\u003e\n+\n+  \u003cItemGroup\u003e\n+    \u003cPackageReference Include=\"Azure.AI.Projects\" Version=\"2.0.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Identity\" Version=\"1.17.0\" /\u003e\n+  \u003c/ItemGroup\u003e\n+\n+\u003c/Project\u003e\n*** Add File: Program.cs\n+using Azure.AI.Projects;\n+using Azure.Identity;\n+\n+string endpoint = GetRequiredEnvironmentVariable(\"FOUNDRY_PROJECT_ENDPOINT\");\n+string connectionName = GetRequiredEnvironmentVariable(\"CONNECTION_NAME\");\n+string deploymentName = GetRequiredEnvironmentVariable(\"DEPLOYMENT_NAME\");\n+\n+var projectClient = new AIProjectClient(\n+    new Uri(endpoint, UriKind.Absolute),\n+    new DefaultAzureCredential());\n+\n+Console.WriteLine(\"Project connections\");\n+Console.WriteLine(\"-------------------\");\n+\n+await foreach (AIProjectConnection connection in projectClient.Connections.GetConnectionsAsync())\n+{\n+    PrintConnection(connection);\n+}\n+\n+Console.WriteLine();\n+Console.WriteLine($\"Connection '{connectionName}'\");\n+Console.WriteLine(new string('-', 13 + connectionName.Length));\n+\n+AIProjectConnection selectedConnection =\n+    await projectClient.Connections.GetConnectionAsync(\n+        connectionName,\n+        includeCredentials: false);\n+PrintConnection(selectedConnection);\n+\n+Console.WriteLine();\n+Console.WriteLine(\"Model deployments\");\n+Console.WriteLine(\"-----------------\");\n+\n+await foreach (AIProjectDeployment deployment in projectClient.Deployments.GetDeploymentsAsync())\n+{\n+    if (deployment is ModelDeployment modelDeployment)\n+    {\n+        PrintModelDeployment(modelDeployment);\n+    }\n+}\n+\n+Console.WriteLine();\n+Console.WriteLine($\"Model deployment '{deploymentName}'\");\n+Console.WriteLine(new string('-', 19 + deploymentName.Length));\n+\n+AIProjectDeployment selectedDeployment =\n+    await projectClient.Deployments.GetDeploymentAsync(deploymentName);\n+\n+if (selectedDeployment is not ModelDeployment selectedModelDeployment)\n+{\n+    throw new InvalidOperationException(\n+        $\"Deployment '{deploymentName}' is not a model deployment.\");\n+}\n+\n+PrintModelDeployment(selectedModelDeployment);\n+\n+static string GetRequiredEnvironmentVariable(string name)\n+{\n+    string? value = Environment.GetEnvironmentVariable(name);\n+\n+    if (string.IsNullOrWhiteSpace(value))\n+    {\n+        throw new InvalidOperationException(\n+            $\"Required environment variable '{name}' is not set.\");\n+    }\n+\n+    return value;\n+}\n+\n+static void PrintConnection(AIProjectConnection connection)\n+{\n+    Console.WriteLine(\n+        $\"Name: {connection.Name}; Type: {connection.Type}; \" +\n+        $\"Target: {connection.Target}; Default: {connection.IsDefault}\");\n+}\n+\n+static void PrintModelDeployment(ModelDeployment deployment)\n+{\n+    Console.WriteLine(\n+        $\"Name: {deployment.Name}; Publisher: {deployment.ModelPublisher}; \" +\n+        $\"Model: {deployment.ModelName}; Version: {deployment.ModelVersion}\");\n+}\n*** Add File: README.md\n+# Foundry Project Inventory\n+\n+This console application uses `Azure.AI.Projects` and `DefaultAzureCredential` to\n+list and retrieve connections and model deployments from a Microsoft Foundry\n+project.\n+\n+Set the required environment variables in PowerShell:\n+\n+```powershell\n+$env:FOUNDRY_PROJECT_ENDPOINT = \"https://\u003cresource\u003e.services.ai.azure.com/api/projects/\u003cproject\u003e\"\n+$env:CONNECTION_NAME = \"\u003cconnection-name\u003e\"\n+$env:DEPLOYMENT_NAME = \"\u003cdeployment-name\u003e\"\n+```\n+\n+Restore, build, and run:\n+\n+```powershell\n+dotnet restore\n+dotnet build --no-restore\n+dotnet run --no-build\n+```\n+\n+Authenticate locally with a credential supported by `DefaultAzureCredential`,\n+such as Azure CLI login, Visual Studio, or environment-based service principal\n+credentials. The application requests connection metadata without credentials.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj
@@ -1,0 +1,16 @@
+<Project Sdk="Microsoft.NET.Sdk">
+
+  <PropertyGroup>
+    <OutputType>Exe</OutputType>
+    <TargetFramework>net8.0</TargetFramework>
+    <ImplicitUsings>enable</ImplicitUsings>
+    <Nullable>enable</Nullable>
+  </PropertyGroup>
+
+  <ItemGroup>
+    <PackageReference Include="Azure.AI.Projects" Version="2.0.0" />
+    <PackageReference Include="Azure.Identity" Version="1.17.0" />
+  </ItemGroup>
+
+</Project>
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/Program.cs b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/Program.cs
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/Program.cs
@@ -1,0 +1,83 @@
+using Azure.AI.Projects;
+using Azure.Identity;
+
+string endpoint = GetRequiredEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
+string connectionName = GetRequiredEnvironmentVariable("CONNECTION_NAME");
+string deploymentName = GetRequiredEnvironmentVariable("DEPLOYMENT_NAME");
+
+var projectClient = new AIProjectClient(
+    new Uri(endpoint, UriKind.Absolute),
+    new DefaultAzureCredential());
+
+Console.WriteLine("Project connections");
+Console.WriteLine("-----
... (truncated)
```

### ✅ 24. powershell

**Input:**

```json
{"command":"dotnet restore --nologo --verbosity quiet \u0026\u0026 dotnet build --no-restore --nologo --verbosity minimal","description":"Restore packages and build application","initial_wait":180,"mode":"sync"}
```

**Output:**

```
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678\FoundryProjectInventory.csproj : error NU1605: Warning As Error: Detected package downgrade: Azure.Identity from 1.20.0 to 1.17.0. Reference the package directly from the project to select a different version. 
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678\FoundryProjectInventory.csproj : error NU1605:  FoundryProjectInventory -> Azure.AI.Projects 2.0.0 -> Azure.Identity (>= 1.20.0) 
C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678\FoundryProjectInventory.csproj : error NU1605:  FoundryProjectInventory -> Azure.Identity (>= 1.17.0)
<shellId: 2 completed with exit code 1>
```

### ✅ 25. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: FoundryProjectInventory.csproj\n@@\n-    \u003cPackageReference Include=\"Azure.Identity\" Version=\"1.17.0\" /\u003e\n+    \u003cPackageReference Include=\"Azure.Identity\" Version=\"1.20.0\" /\u003e\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678/FoundryProjectInventory.csproj
@@ -9,7 +9,7 @@
 
   <ItemGroup>
     <PackageReference Include="Azure.AI.Projects" Version="2.0.0" />
-    <PackageReference Include="Azure.Identity" Version="1.17.0" />
+    <PackageReference Include="Azure.Identity" Version="1.20.0" />
   </ItemGroup>
 
 </Project>

```

### ✅ 26. powershell

**Input:**

```json
{"command":"dotnet restore --nologo --verbosity quiet \u0026\u0026 dotnet build --no-restore --nologo --verbosity minimal","description":"Restore updated packages and build","initial_wait":180,"mode":"sync"}
```

**Output:**

```
FoundryProjectInventory -> C:\Users\weidxu\AppData\Local\Temp\hyoka-ai-projects-dp-dotnet-project-resource-inventory-dotnet-azure-tools-baseline-1867024678\bin\Debug\net8.0\FoundryProjectInventory.dll

Build succeeded.
    0 Warning(s)
    0 Error(s)

Time Elapsed 00:00:01.03
<shellId: 3 completed with exit code 0>
```

## Generated Files

- `FoundryProjectInventory.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’ll inspect the workspace, then create the project against the installed/current Azure SDK API and validate it locally.

The current SDK exposes typed `AIProjectConnection` and `ModelDeployment` models, so I’ll implement against the stable 2.0 API with explicit credential-free connection retrieval and model type checks.

Restore identified the SDK’s required Azure Identity baseline; I’m aligning the manifest to that dependency and rebuilding.

Created a complete .NET 8 console application in the workspace:

- `FoundryProjectInventory.csproj`
- `Program.cs`
- `README.md`

It uses asynchronous `Azure.AI.Projects` pageable APIs, prints typed connection and model deployment metadata, retrieves connections without credentials, validates that the named deployment is a model deployment, and includes PowerShell restore/build/run instructions.

## Reviewer Notes (LLM-as-Judge)

### Summary

9/9 criteria passed. The application implements the required asynchronous project inventory and named-resource retrieval workflow.

### Strengths

- Correct asynchronous pageable iteration for connections and deployments.
- Secure metadata-only connection retrieval with credentials explicitly excluded.
- Proper runtime type narrowing and rejection of non-model deployments.
- Includes a runnable project manifest and concise restore, build, and run instructions.

## Grader Results

- project-resource-inventory.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Creates `AIProjectClient` for `FOUNDRY_PROJECT_ENDPOINT`.: Pass
      - Enumerates `Connections.GetConnectionsAsync` with `await foreach` and uses: Pass
      - Calls `Connections.GetConnectionAsync` for `CONNECTION_NAME` with: Pass
      - Enumerates `Deployments.GetDeploymentsAsync` with `await foreach`.: Pass
      - Narrows `AIProjectDeployment` values to `ModelDeployment` before reading: Pass
      - Calls `Deployments.GetDeploymentAsync` for `DEPLOYMENT_NAME` and rejects a result: Pass
      - Does not replace pageable iteration with a single assumed page.: Pass
      - Does not select resources only from the locally enumerated results instead of: Pass
      - Does not request or display connection credentials.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id ai-projects-dp-dotnet-project-resource-inventory --config dotnet-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
