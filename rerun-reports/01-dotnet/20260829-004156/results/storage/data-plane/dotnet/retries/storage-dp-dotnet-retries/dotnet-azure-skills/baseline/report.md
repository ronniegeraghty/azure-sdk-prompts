# Evaluation Report: storage-dp-dotnet-retries

**Config:** dotnet-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 206.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-dotnet-retries` |
| Config | dotnet-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 7/8 |
| Duration | 206.9s |
| Timestamp | 2026-08-28T18:53:41Z |
| Files Generated | 0 |
| Event Count | 4412 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 188.7s |
| Review | 17.5s |
| **Total** | **206.9s** |

## Configuration

- **name:** dotnet-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=194623 out=5352 |
| Turn Count | 6 |

## Tool Calls

### ✅ 1. web_search (via github-mcp-server) — 99317ms

**Input:**

```json
{"query":"Official Azure SDK for .NET Azure.Storage.Blobs BlobClientOptions RetryOptions NetworkTimeout MaxRetries Delay MaxDelay retryable HTTP status codes Azure.Core 2026"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's official guidance, tailored for your scenario using Polly with Azure.Storage.Blobs (for example, `BlobClient`), focusing on Circuit Breaker, operation timeout, and proper `CancellationToken` usage:\n\n---\n\n### 1. **Circuit Breaker with Polly**\n\nPolly's Circuit Breaker will prevent repeated poor-performing operations on cloud services (e.g., Blob Storage). You can configure it for failure ratio, minimum attempts, sample window, and break duration. Polly throws `BrokenCircuitException` when the circuit is open. You usually *combine* Circuit Breaker with Retry and Timeout for robust cloud code【3:0†source】【3:2†source】.\n\n**Example:**\n```csharp\nvar circuitBreakerPolicy = Policy\n    .Handle<Exception>() // Or specify more granular Azure exceptions\n    .CircuitBreakerAsync(\n        handledEventsAllowedBeforeBreaking: 5,\n        durationOfBreak: TimeSpan.FromSeconds(30)\n    );\n```\n\n---\n\n### 2. **Timeout Policy with CancellationToken**\n\nThe Polly Timeout policy ensures that a given operation (such as a blob upload/download) will be abandoned if not completed in time. When used, Polly generates its own timeout `CancellationTokenSource`, combining it with any user-provided `CancellationToken`, so *either* can trigger cancellation. You should always pass the combined token to your blob operation.\n\n**Example:**\n```csharp\nvar timeoutPolicy = Policy\n    .TimeoutAsync<BlobContentInfo>(TimeSpan.FromSeconds(10)); // For uploads, for example\n\n// Usage pattern\nawait timeoutPolicy.ExecuteAsync(async ct =>\n    await blobClient.UploadAsync(stream, cancellationToken: ct), \n    externalCancellationToken // This is your (user-provided) CancellationToken\n);\n```\n**Important:** Your blob operation must honor the `CancellationToken`. Azure.Storage.Blobs operations *do* honor cancellation, but any custom async work *inside* the delegate must as well【3:3†source】.\n\n---\n\n### 3. **Combining Polly Polic
... (truncated)
```

### ✅ 2. web_search (via github-mcp-server) — 99003ms

**Input:**

```json
{"query":"Azure.Core .NET ResponseClassifier retryable status codes source 408 429 500 502 503 504 official GitHub"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The Azure SDK for .NET allows you to configure custom retry and pipeline policies using the `BlobClientOptions` and the extensibility provided by `HttpPipelinePolicy`. Here’s how you can approach each requirement:\n\n---\n\n### 1. **Custom Retry Policy via BlobClientOptions**\n\n- **Default Retry Options:** You can set `RetryOptions` to control retries for transient failures (such as timeouts) directly via `BlobClientOptions`:\n\n```csharp\nvar options = new BlobClientOptions();\noptions.Retry.MaxRetries = 4;\noptions.Retry.Delay = TimeSpan.FromSeconds(2);\noptions.Retry.MaxDelay = TimeSpan.FromSeconds(10);\noptions.Retry.Mode = RetryMode.Exponential; // Or Fixed\n\nvar blobServiceClient = new BlobServiceClient(connectionString, options);\n```\nThese options are sufficient for most retry scenarios【3:5†source】.\n\n---\n\n### 2. **Custom Retry Policy Class**\n\nFor advanced scenarios, implement your own retry policy by deriving from the `RetryPolicy` class, then assign it to `options.RetryPolicy`. This replaces the standard retry behavior:\n\n```csharp\npublic class CustomRetryPolicy : RetryPolicy\n{\n    public override void Process(HttpMessage message, ReadOnlyMemory<HttpPipelinePolicy> pipeline)\n    {\n        // Custom logic here\n        ProcessNext(message, pipeline);\n        // Inspect or mutate the response as needed\n    }\n\n    public override ValueTask ProcessAsync(HttpMessage message, ReadOnlyMemory<HttpPipelinePolicy> pipeline)\n    {\n        // Custom async logic here\n        return ProcessNextAsync(message, pipeline);\n    }\n}\n```\n\nAssign it:\n```csharp\nvar options = new BlobClientOptions\n{\n    RetryPolicy = new CustomRetryPolicy()\n};\n```\nBe aware, when you override the RetryPolicy, you’re expected to manage the retry logic yourself (including backoff, limits, etc.)【3:6†source】【3:3†source】.\n\n---\n\n### 3. **Adding a Circuit Breaker (Custom HttpPipelinePolicy)**\n\nThe pipeline
... (truncated)
```

### ✅ 3. web_search (via github-mcp-server) — 98679ms

**Input:**

```json
{"query":"Microsoft official Azure SDK .NET custom HttpPipelinePolicy retry policy circuit breaker BlobClientOptions"}
```

**Output:**

```
{"type":"output_text","text":{"value":"The official retriable status codes for the default `ResponseClassifier` in Azure.Core for .NET are 408, 429, 500, 502, 503, and 504. This logic is implemented in the `StatusCodeClassifier`, which is the default classifier used. You can verify this from the official Azure SDK for .NET source code on GitHub, specifically in the [ResponseClassifier.cs](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs) file and related implementation.\n\n**Direct code reference (as found in source):**\n```csharp\nprivate static readonly int[] RetriableStatusCodes = {408, 429, 500, 502, 503, 504};\n```\nThese status codes correspond to:\n- 408: Request Timeout\n- 429: Too Many Requests\n- 500: Internal Server Error\n- 502: Bad Gateway\n- 503: Service Unavailable\n- 504: Gateway Timeout\n\nYou can browse or view the source for up-to-date details directly in the Azure SDK for .NET GitHub repository here:  \n[ResponseClassifier.cs on GitHub](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs)【3:1†source】.","annotations":[{"text":"","start_index":1066,"end_index":1078,"url_citation":{"title":"azure-sdk-for-net/sdk/core/Azure.Core/src/ResponseClassifier ... - GitHub","url":"https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs"}}]},"bing_searches":[{"text":"Azure.Core .NET ResponseClassifier retryable status codes official GitHub source","url":"https://www.bing.com/search?q=Azure.Core+.NET+ResponseClassifier+retryable+status+codes+official+GitHub+source"}],"annotations":null}
```

### ✅ 4. web_search (via github-mcp-server) — 98352ms

**Input:**

```json
{"query":"Polly circuit breaker Azure.Storage.Blobs .NET official guidance operation timeout CancellationToken"}
```

**Output:**

```
{"type":"output_text","text":{"value":"Here's a summary for configuring retries for Azure Blob Storage using the official Azure SDK for .NET (`Azure.Storage.Blobs`, `BlobClientOptions`, and `Azure.Core`), as of 2026:\n\n## Key Classes & Namespaces\n\n- **BlobClientOptions**: Used to configure the client. Located in `Azure.Storage.Blobs`.\n- **RetryOptions**: Encapsulates retry settings. Inherited from `Azure.Core.ClientOptions`.\n\n## RetryOptions Properties\n\n| Property     | Type      | Description                                                                 | Default Value      |\n|--------------|-----------|-----------------------------------------------------------------------------|--------------------|\n| `Delay`      | TimeSpan  | The delay between retry attempts (fixed) or initial delay (exponential).    | 0.8 seconds        |\n| `MaxDelay`   | TimeSpan  | Maximum delay between retries (if no `Retry-After`).                        | 1 minute           |\n| `MaxRetries` | int       | Max retry attempts before failing.                                          | 5                  |\n| `Mode`       | enum      | Retry strategy: `Exponential` or `Fixed`.                                   | `Exponential`      |\n| `NetworkTimeout` | TimeSpan | Timeout for an individual network operation.                              | 100 seconds        |\n\n## Setting Retry Options Example\n\n```csharp\nBlobClientOptions options = new BlobClientOptions();\noptions.Retry.Delay = TimeSpan.FromSeconds(1);\noptions.Retry.MaxDelay = TimeSpan.FromSeconds(30);\noptions.Retry.MaxRetries = 4;\noptions.Retry.NetworkTimeout = TimeSpan.FromSeconds(120); // adjust as needed\n\nvar serviceClient = new BlobServiceClient(connectionString, options);\n```\nYou must set these properties before instantiating your BlobServiceClient/BlobClient【3:3†source】.\n\n## Retryable HTTP Status Codes\n\nThe SDK will retry on transient errors, such as:\n\n- HTTP 408 (Request Timeout)\n- HTTP 429 (Too 
... (truncated)
```

### ✅ 5. web_fetch — 3937ms

**Input:**

```json
{"max_length":20000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs"}
```

**Output:**

```
Contents of https://www.pollydocs.org/strategies/circuit-breaker.html:
# Circuit breaker resilience strategy

## About

* * *

The circuit breaker **reactive** resilience strategy shortcuts the execution if the underlying resource is detected as unhealthy. The detection process is done via sampling. If the sampled executions' failure-success ratio exceeds a predefined threshold then a circuit breaker will prevent any new executions by throwing a `BrokenCircuitException`. After a preset duration the circuit breaker performs a probe, because the assumption is that this period was enough for the resource to self-heal. Depending on the outcome of the probe, the circuit will either allow new executions or continue to block them. If an execution is blocked by the circuit breaker, the thrown exception may indicate the amount of time executions will continue to be blocked through its `RetryAfter` property.

##### Note

Be aware that the Circuit Breaker strategy [rethrows all exceptions](https://github.com/App-vNext/Polly/wiki/Circuit-Breaker#exception-handling), including those that are handled. A Circuit Breaker's role is to monitor faults and break the circuit when a certain threshold is reached; it does not manage retries. Combine the Circuit Breaker with a Retry strategy if needed.

## Usage

```
// Circuit breaker with default options.
// See https://www.pollydocs.org/strategies/circuit-breaker#defaults for defaults.
var optionsDefaults = new CircuitBreakerStrategyOptions();

// Circuit breaker with customized options:
// The circuit will break if more than 50% of actions result in handled exceptions,
// within any 10-second sampling duration, and at least 8 actions are processed.
var optionsComplex = new CircuitBreakerStrategyOptions
{
    FailureRatio = 0.5,
    SamplingDuration = TimeSpan.FromSeconds(10),
    MinimumThroughput = 8,
    BreakDuration = TimeSpan.FromSeconds(30),
    ShouldHandle = new PredicateBuilder().Handle<SomeExceptionType>()
};

// Circuit break
... (truncated)
```

### ✅ 6. web_fetch — 3606ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/ResponseClassifier.cs"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/ResponseClassifier.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.IO;

namespace Azure.Core
{
    /// <summary>
    /// A type that analyzes HTTP responses and exceptions and determines if they should be retried,
    /// and/or analyzes responses and determines if they should be treated as error responses.
    /// </summary>
    public class ResponseClassifier
    {
        internal static ResponseClassifier Shared { get; } = new();

        /// <summary>
        /// Specifies if the request contained in the <paramref name="message"/> should be retried.
        /// </summary>
        public virtual bool IsRetriableResponse(HttpMessage message)
        {
            switch (message.Response.Status)
            {
                case 408: // Request Timeout
                case 429: // Too Many Requests
                case 500: // Internal Server Error
                case 502: // Bad Gateway
                case 503: // Service Unavailable
                case 504: // Gateway Timeout
                    return true;
                default:
                    return false;
            }
        }

        /// <summary>
        /// Specifies if the operation that caused the exception should be retried.
        /// </summary>
        public virtual bool IsRetriableException(Exception exception)
        {
            return (exception is IOException) ||
                   (exception is RequestFailedException requestFailed && requestFailed.Status == 0);
        }

        /// <summary>
        /// Specifies if the operation that caused the exception should be retried taking the <see cref="HttpMessage"/> into consideration.
        /// </summary>
        public virtual bool IsRetriable(HttpMessage
... (truncated)
```

### ✅ 7. web_fetch — 3278ms

**Input:**

```json
{"max_length":12000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet:
---
layout: Reference
monikers:
- azure-dotnet
defaultMoniker: azure-dotnet
versioningType: Ranged
title: ClientOptions.RetryPolicy Property (Azure.Core) - Azure for .NET Developers | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet
config_moniker_range: azure-dotnet-legacy || azure-dotnet || azure-dotnet-preview
uid: Azure.Core.ClientOptions.RetryPolicy*
namespace: Azure.Core
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-net/issues
author: azure-sdk
ms.author: azsdkdocs
ms.topic: generated-reference
ms.devlang: csharp
ms.date: 2021-04-13T00:00:00.0000000Z
uhfHeaderId: Azure
breadcrumb_path: ../azure_breadcrumb/toc.json
apiPlatform: dotnet
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
ms.service: azure
ms.custom:
- azure-dotnet-devcenter
- azure-dotnet-devcenter-generated
api_name:
- Azure.Core.ClientOptions.RetryPolicy
- Azure.Core.ClientOptions.get_RetryPolicy
- Azure.Core.ClientOptions.set_RetryPolicy
api_location:
- Azure.Core.dll
topic_type:
- apiref
api_type:
- Assembly
locale: en-us
document_id: 4600e65f-3027-abe5-f2ea-aa6125fe7d79
document_version_independent_id: b0e545bb-9918-8c2f-f6a2-85ac3dd21f2f
updated_at: 2026-08-28T11:57:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/live/xml/Azure.Core/ClientOptions.xml
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-dotnet/blob/c5d2f1e44889c00f914e0d66fa7d0148914ebc8a/xml/Azure.Core/ClientOptions.xml
git_commit_id: c5d2f1e44889c00f914e0d66fa7d0148914ebc8a
default_moniker: azure-dotnet
site_name: Docs
depot_name: Azure.azuredotnet
page_type: dotnet
page_kind: property
ms.assetid: Azure.Core.ClientOptions.RetryPolicy*
description: "Gets or sets the policy to use for re
... (truncated)
```

### ✅ 8. web_fetch — 2953ms

**Input:**

```json
{"max_length":7000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.clientoptions.retrypolicy?view=azure-dotnet"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy:
---
layout: Conceptual
title: Implement a retry policy using the Azure Storage client library for .NET - Azure Storage | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/storage-retry-policy
breadcrumb_path: /azure/bread/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/125/azure-blob-storage/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/a8bb4a47-3525-ec11-b6e6-000d3a4f0f84
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
adobe-target: true
learn_banner_products:
- azure
manager: akashdubey
ms.reviewer: akashdubey-ms
description: Learn about retry policies and how to implement them for Blob Storage. This article helps you set up a retry policy for Blob Storage requests using the Azure Storage client library for .NET.
author: stevenmatthew
ms.author: shaas
ms.service: azure-blob-storage
ms.topic: how-to
ms.date: 2024-11-19T00:00:00.0000000Z
ms.custom: devx-track-dotnet, devguide-csharp
locale: en-us
document_id: 3b78918f-3432-e8b0-aa74-df6998979c4d
document_version_independent_id: 4434fc07-9da4-1db4-6bf5-61401ecf97bc
updated_at: 2025-08-15T05:11:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/storage-retry-policy.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/9656d6c3a103d19d44c6fcc49b1391adc8590214/articles/storage/blobs/storage-retry-policy.md
git_commit_id: 9656d6c3a103d19d44c6fcc49b1391adc8590214
site_name: Docs
depot_name: Azure.azure-documents
page_type: conceptual
toc_rel: toc.json
pdf_url_template: https://learn.microsoft.com/pdfstore/en-us/Azure.azure-documents/{branchName}{pdfName}
word_count: 854
asset_id: storage/blobs/storage-retry-policy
moniker_rang
... (truncated)
```

### ✅ 9. web_fetch — 2627ms

**Input:**

```json
{"max_length":10000,"url":"https://www.pollydocs.org/strategies/circuit-breaker.html"}
```

**Output:**

```
Content type text/plain; charset=utf-8 cannot be simplified to markdown. Here is the raw content:
Contents of https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/RetryPolicy.cs:
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Runtime.ExceptionServices;
using System.Threading;
using System.Threading.Tasks;
using Azure.Core.Diagnostics;

namespace Azure.Core.Pipeline
{
    /// <summary>
    /// Represents a policy that can be overridden to customize whether or not a request will be retried and how long to wait before retrying.
    /// </summary>
    public class RetryPolicy : HttpPipelinePolicy
    {
        private readonly int _maxRetries;

        /// <summary>
        /// Gets the delay to use for computing the interval between retry attempts.
        /// </summary>
        private readonly DelayStrategy _delayStrategy;

        /// <summary>
        /// Initializes a new instance of the <see cref="RetryPolicy"/> class.
        /// </summary>
        /// <param name="maxRetries">The maximum number of retries to attempt.</param>
        /// <param name="delayStrategy">The delay to use for computing the interval between retry attempts.</param>
        public RetryPolicy(int maxRetries = RetryOptions.DefaultMaxRetries, DelayStrategy? delayStrategy = default)
        {
            _maxRetries = maxRetries;
            _delayStrategy = delayStrategy ?? DelayStrategy.CreateExponentialDelayStrategy();
        }

        /// <summary>
        /// This method can be overridden to take full control over the retry policy. If this is overridden and the base method isn't called,
        /// it is the implementer's responsibility to populate the <see cref="HttpMessage.ProcessingContext"/> property.
        /// This method will only be called for async methods.
        /// </summary>
        /// <param
... (truncated)
```

### ❌ 10. web_fetch — 2116ms

**Input:**

```json
{"max_length":12000,"url":"https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/DelayStrategy.cs"}
```

**Error:**

```
Error: Failed to fetch https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.delaystrategy.createexponentialdelaystrategy?view=azure-dotnet - status code 404
```

### ❌ 11. web_fetch — 2115ms

**Input:**

```json
{"max_length":8000,"url":"https://learn.microsoft.com/en-us/dotnet/api/azure.core.pipeline.delaystrategy.createexponentialdelaystrategy?view=azure-dotnet"}
```

**Error:**

```
Error: Failed to fetch https://raw.githubusercontent.com/Azure/azure-sdk-for-net/main/sdk/core/Azure.Core/src/Pipeline/DelayStrategy.cs - status code 404
```

### ✅ 12. web_fetch — 2112ms

**Input:**

```json
{"max_length":6000,"url":"https://www.pollydocs.org/strategies/timeout.html"}
```

**Output:**

```
Contents of https://www.pollydocs.org/strategies/timeout.html:
# Timeout resilience strategy

## About

* * *

The timeout **proactive** resilience strategy cancels the execution if it does not complete within the specified timeout period. If the execution is canceled by the timeout strategy, it throws a `TimeoutRejectedException`. The timeout strategy operates by wrapping the incoming cancellation token with a new one. Should the original token be canceled, the timeout strategy will transparently honor the original cancellation token without throwing a `TimeoutRejectedException`.

##### Important

It is crucial that the user's callback respects the cancellation token. If it does not, the callback will continue executing even after a cancellation request, thereby ignoring the cancellation.

## Usage

```
// To add a timeout with a custom TimeSpan duration
new ResiliencePipelineBuilder().AddTimeout(TimeSpan.FromSeconds(3));

// Timeout using the default options.
// See https://www.pollydocs.org/strategies/timeout#defaults for defaults.
var optionsDefaults = new TimeoutStrategyOptions();

// To add a timeout using a custom timeout generator function
var optionsTimeoutGenerator = new TimeoutStrategyOptions
{
    TimeoutGenerator = static args =>
    {
        // Note: the timeout generator supports asynchronous operations
        return new ValueTask<TimeSpan>(TimeSpan.FromSeconds(123));
    }
};

// To add a timeout and listen for timeout events
var optionsOnTimeout = new TimeoutStrategyOptions
{
    TimeoutGenerator = static args =>
    {
        // Note: the timeout generator supports asynchronous operations
        return new ValueTask<TimeSpan>(TimeSpan.FromSeconds(123));
    },
    OnTimeout = static args =>
    {
        Console.WriteLine($"{args.Context.OperationKey}: Execution timed out after {args.Timeout.TotalSeconds} seconds.");
        return default;
    }
};

// Add a timeout strategy with a TimeoutStrategyOptions instance to the pipeline
new ResilienceP
... (truncated)
```

### ✅ 13. github-mcp-server-search_code (via github-mcp-server) — 2032ms

**Input:**

```json
{"fields":["name","path","sha","text_matches"],"perPage":10,"query":"class StorageResponseClassifier repo:Azure/azure-sdk-for-net path:sdk/storage"}
```

**Output:**

```
{"incomplete_results":false,"items":[{"name":"StorageResponseClassifier.cs","path":"sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs","sha":"b1ed5e397f663ed445f621133d37609923009562","text_matches":[{"fragment":"\nnamespace Azure.Storage\n{\n    internal class StorageResponseClassifier : ResponseClassifier\n    {\n        /// <summary>\n        /// The secondary URI to be used for retries on failed read requests","matches":[{"indices":[40,45],"text":"class"},{"indices":[46,71],"text":"StorageResponseClassifier"},{"indices":[61,66],"text":"Class"},{"indices":[82,87],"text":"Class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"StorageClassifierPolicy.cs","path":"sdk/storage/Azure.Storage.Common/src/Shared/StorageClassifierPolicy.cs","sha":"23f088c6523c6329617386eb49c36041a9fa5812","text_matches":[{"fragment":"    /// </summary>\n    internal class StorageClassifierPolicy : HttpPipelineSynchronousPolicy\n    {\n        private readonly StorageResponseClassifier _storageClassifier;\n\n        public StorageClassifierPolicy(StorageResponseClassifier storageClassifier)\n        {","matches":[{"indices":[32,37],"text":"class"},{"indices":[45,50],"text":"Class"},{"indices":[125,150],"text":"StorageResponseClassifier"},{"indices":[140,145],"text":"Class"},{"indices":[159,164],"text":"Class"},{"indices":[194,199],"text":"Class"},{"indices":[211,236],"text":"StorageResponseClassifier"},{"indices":[226,231],"text":"Class"},{"indices":[244,249],"text":"Class"}],"object_type":"FileContent","object_url":"https://api.github.com/repositories/2928944/contents/sdk/storage/Azure.Storage.Common/src/Shared/StorageClassifierPolicy.cs?ref=f7a9a696b3f71350b43e3b3a675fd9a2144e1327","property":"content"}]},{"name":"StorageRetriableClassifier.cs","path":"sdk/storage/Azur
... (truncated)
```

### ✅ 14. github-mcp-server-get_file_contents (via github-mcp-server) — 2191ms

**Input:**

```json
{"owner":"Azure","path":"sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs","repo":"azure-sdk-for-net"}
```

**Output:**

```
successfully downloaded text file (SHA: b1ed5e397f663ed445f621133d37609923009562)// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

using System;
using Azure.Core;

namespace Azure.Storage
{
    internal class StorageResponseClassifier : ResponseClassifier
    {
        /// <summary>
        /// The secondary URI to be used for retries on failed read requests
        /// </summary>
        public Uri SecondaryStorageUri { get; set; }

        /// <summary>
        /// Overridden version of IsRetriableResponse that allows for Storage specific retry logic.
        /// </summary>
        /// <param name="message">The message containing both Response and Request</param>
        /// <returns></returns>
        public override bool IsRetriableResponse(HttpMessage message)
        {
            // If secondary storage Uri was specified, we want to retry if the current attempt was against the secondary Uri, and we
            // get a response of NotFound. This is because the resource may not have been propagated to secondary Uri yet.
            if (SecondaryStorageUri != null &&
                message.Request.Uri.Host == SecondaryStorageUri.Host &&
                message.Response.Status == Constants.HttpStatusCode.NotFound)
            {
                return true;
            }

            // Retry select Storage service error codes
            if (message.Response.Status >= 400 &&
                message.Response.Headers.TryGetValue(Constants.HeaderNames.ErrorCode, out var error))
            {
                switch (error)
                {
                    case Constants.ErrorCodes.InternalError:
                    case Constants.ErrorCodes.OperationTimedOut:
                    case Constants.ErrorCodes.ServerBusy:
                        return true;
                }
            }

            // Retry select Copy Source Error Codes.
            if (message.Response.Status >= 400 &&
                message.Resp
... (truncated)
```

## Copilot Response

I’ll verify the current `Azure.Core` retry classification and Blob SDK APIs, then provide a complete .NET pattern covering timeouts, operation-specific policies, and circuit breaking.

Configure retries when constructing the client, use cancellation tokens for a whole-operation deadline, and place a circuit breaker around the SDK call. Keep the SDK’s built-in retry policy as the only retry layer to avoid multiplying retries.

```bash
dotnet add package Azure.Storage.Blobs
dotnet add package Azure.Identity
dotnet add package Polly
```

## 1. Configure `BlobClientOptions.Retry`

```csharp
using Azure.Core;
using Azure.Identity;
using Azure.Storage.Blobs;

Uri accountUri = new("https://<account-name>.blob.core.windows.net");

var options = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 5,
        Delay = TimeSpan.FromSeconds(1),
        MaxDelay = TimeSpan.FromSeconds(30),
        NetworkTimeout = TimeSpan.FromSeconds(60)
    }
};

var serviceClient = new BlobServiceClient(
    accountUri,
    new DefaultAzureCredential(),
    options);
```

`MaxRetries = 5` means **one initial attempt plus up to five retries**, for at most six attempts. `Delay` is the basis for exponential backoff, while `MaxDelay` caps SDK-calculated delays. A service-provided `Retry-After` delay takes precedence.

`NetworkTimeout` applies to an **individual network operation**, not the complete high-level upload. A multipart upload can involve several network operations and therefore run much longer.

## 2. Enforce a total per-operation timeout

Use a linked `CancellationTokenSource` to limit the entire upload:

```csharp
using Azure;
using Azure.Storage.Blobs;

static async Task UploadWithDeadlineAsync(
    BlobClient blob,
    Stream content,
    TimeSpan timeout,
    CancellationToken callerToken)
{
    using var deadline = CancellationTokenSource.CreateLinkedTokenSource(callerToken);
    deadline.CancelAfter(timeout);

    try
    {
        await blob.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: deadline.Token);
    }
    catch (OperationCanceledException)
        when (deadline.IsCancellationRequested &&
              !callerToken.IsCancellationRequested)
    {
        throw new TimeoutException(
            $"Blob upload did not complete within {timeout}.");
    }
}
```

This gives you two separate limits:

- `NetworkTimeout`: bounds an individual network read/write.
- `CancelAfter`: bounds the complete upload, including retries, backoff, and multipart requests.

## 3. Distinguish transient failures from authentication failures

```csharp
using Azure;

static bool IsRetryableStatus(int status) =>
    status is 408 or 429 or 500 or 502 or 503 or 504;

static bool IsTransient(RequestFailedException exception) =>
    exception.Status == 0 || // No HTTP response, commonly a network failure
    IsRetryableStatus(exception.Status) ||
    exception.ErrorCode is "InternalError"
        or "OperationTimedOut"
        or "ServerBusy";

static async Task UploadAndHandleErrorsAsync(
    BlobClient blob,
    Stream content,
    CancellationToken cancellationToken)
{
    try
    {
        await blob.UploadAsync(
            content,
            overwrite: true,
            cancellationToken: cancellationToken);
    }
    catch (RequestFailedException ex) when (ex.Status is 401 or 403)
    {
        // Non-transient: refresh/fix credentials, RBAC, SAS, or permissions.
        // Repeating the same request is unlikely to help.
        throw new InvalidOperationException(
            $"Blob authentication or authorization failed: {ex.ErrorCode}", ex);
    }
    catch (RequestFailedException ex) when (IsTransient(ex))
    {
        // The SDK has already exhausted its configured retries.
        // Queue for later processing, return a temporary failure, or alert.
        throw new IOException(
            $"Blob Storage remained unavailable after retries: {ex.ErrorCode}", ex);
    }
}
```

Normally you should let `RequestFailedException` propagate after logging sufficient context. Do not immediately start another application-level retry loop, because the SDK may already have made six attempts.

## 4. Use a different retry policy for high-value operations

Retry settings are client-level rather than call-level. Create a dedicated client for operations requiring a different policy:

```csharp
var highValueOptions = new BlobClientOptions
{
    Retry =
    {
        Mode = RetryMode.Exponential,
        MaxRetries = 8,
        Delay = TimeSpan.FromSeconds(2),
        MaxDelay = TimeSpan.FromMinutes(1),
        NetworkTimeout = TimeSpan.FromMinutes(2)
    }
};

var highValueServiceClient = new BlobServiceClient(
    accountUri,
    new DefaultAzureCredential(),
    highValueOptions);

BlobClient normalBlob = serviceClient
    .GetBlobContainerClient("uploads")
    .GetBlobClient("normal.dat");

BlobClient highValueBlob = highValueServiceClient
    .GetBlobContainerClient("uploads")
    .GetBlobClient("high-value.dat");
```

Clients are thread-safe and intended to be reused, so register both as long-lived singletons rather than creating them for every upload.

For fully custom retry behavior, `BlobClientOptions.RetryPolicy` can replace `RetryOptions` with a custom `HttpPipelinePolicy` or derived `RetryPolicy`. Once set, `RetryPolicy` takes precedence over `Retry`; implementing it correctly requires preserving pipeline processing state, so separate clients with tuned `RetryOptions` are generally safer.

## 5. Add a circuit breaker for sustained transient failures

Polly v8 can wrap the complete SDK operation. This example does **not** add another retry strategy—the Azure SDK already owns retries.

```csharp
using Azure;
using Polly;
using Polly.CircuitBreaker;
using Polly.Timeout;

static bool IsTransient(Exception exception) =>
    exception switch
    {
        TimeoutRejectedException => true,

        RequestFailedException requestFailed =>
            requestFailed.Status == 0 ||
            IsRetryableStatus(requestFailed.Status) ||
            requestFailed.ErrorCode is "InternalError"
                or "OperationTimedOut"
                or "ServerBusy",

        HttpRequestException => true,
        IOException => true,

        AggregateException aggregate when aggregate.InnerExceptions.Count > 0 =>
            aggregate.InnerExceptions.All(IsTransient),

        // Caller cancellation and auth errors are deliberately excluded.
        _ => false
    };

var resiliencePipeline = new ResiliencePipelineBuilder()
    .AddCircuitBreaker(new CircuitBreakerStrategyOptions
    {
        FailureRatio = 0.5,
        MinimumThroughput = 10,
        SamplingDuration = TimeSpan.FromSeconds(30),
        BreakDuration = TimeSpan.FromSeconds(30),
        ShouldHandle = new PredicateBuilder()
            .Handle<Exception>(IsTransient)
    })
    // Total deadline for one complete SDK operation, including SDK retries.
    .AddTimeout(TimeSpan.FromMinutes(10))
    .Build();

static async Task UploadHighValueAsync(
    ResiliencePipeline pipeline,
    BlobClient blob,
    Stream content,
    CancellationToken cancellationToken)
{
    try
    {
        await pipeline.ExecuteAsync(
            async token =>
            {
                await blob.UploadAsync(
                    content,
                    overwrite: true,
                    cancellationToken: token);
            },
            cancellationToken);
    }
    catch (BrokenCircuitException ex)
    {
        throw new InvalidOperationException(
            $"Blob Storage circuit is open. Retry after {ex.RetryAfter}.", ex);
    }
    catch (TimeoutRejectedException ex)
    {
        throw new TimeoutException("The complete blob upload timed out.", ex);
    }
}
```

Because the circuit breaker is outside the SDK, it records a failure only after the SDK has exhausted its internal retries. Authentication failures, authorization failures, and caller-requested cancellation do not contribute to opening the circuit.

## Default retry classification

The underlying `Azure.Core.ResponseClassifier` considers these HTTP status codes retryable:

| Status | Meaning |
|---:|---|
| `408` | Request Timeout |
| `429` | Too Many Requests |
| `500` | Internal Server Error |
| `502` | Bad Gateway |
| `503` | Service Unavailable |
| `504` | Gateway Timeout |

It also retries:

- `IOException`
- `RequestFailedException` with status `0`, meaning no HTTP response was obtained
- Non-user-initiated `OperationCanceledException`, such as an internal network timeout

Azure Storage adds service-specific classification for responses whose `x-ms-error-code` or copy-source error code is:

- `InternalError`
- `OperationTimedOut`
- `ServerBusy`

When geo-redundant secondary reads are configured, a `404` received from the secondary endpoint can also be retried against another endpoint. A normal primary-endpoint `404` is not retryable.

`400`, `401`, `403`, `404`, `409`, and most other client errors are not retryable merely because they are `4xx`. In particular, `401` and `403` normally require correcting credentials, SAS validity, RBAC assignments, or permissions.

Sources: [Microsoft Blob retry-policy guidance](https://learn.microsoft.com/azure/storage/blobs/storage-retry-policy), [`Azure.Core.ResponseClassifier`](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/core/Azure.Core/src/ResponseClassifier.cs), [Storage-specific classifier](https://github.com/Azure/azure-sdk-for-net/blob/main/sdk/storage/Azure.Storage.Common/src/Shared/StorageResponseClassifier.cs), and [Polly circuit-breaker documentation](https://www.pollydocs.org/strategies/circuit-breaker.html).

## Reviewer Notes (LLM-as-Judge)

### Summary

7 of 8 criteria passed. The response thoroughly covers retry configuration, timeout scopes, failure classification, dedicated high-value clients, and Polly-based circuit breaking, but omits actual GeoRedundantSecondaryUri configuration.

### Strengths

- Correctly distinguishes per-request NetworkTimeout from a complete-operation cancellation deadline.
- Clearly explains retry counts, exponential backoff, default retryable statuses, and non-transient authentication errors.
- Uses a dedicated client for operation-specific retry settings and integrates Polly without multiplying retries.

### Issues

- No BlobClientOptions.GeoRedundantSecondaryUri configuration example or explanation is provided.

## Grader Results

- retry-configuration.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (7/8)
      - `BlobClientOptions.Retry` configuration with `RetryOptions`: Pass
      - `MaxRetries`, `Delay`, `MaxDelay`, `Mode` (Exponential vs Fixed): Pass
      - `NetworkTimeout` for per-request timeouts: Pass
      - Default retryable status codes (408, 429, 500, 502, 503, 504): Pass
      - Non-retryable errors (400, 401, 403, 404, 409): Pass
      - Per-operation `CancellationToken` for timeout control: Pass
      - Geo-redundant retry with `GeoRedundantSecondaryUri`: Fail
      - Interaction with Polly or other resilience libraries: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 88% | 1.00 | 0.8750 | 100.0% | ❌ |
| **Final** | | | **Σ 1.00** | **Σ 0.8750** | **87.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-dotnet-retries --config dotnet-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
