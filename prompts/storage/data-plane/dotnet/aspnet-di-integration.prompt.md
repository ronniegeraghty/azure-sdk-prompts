---
id: storage-dp-dotnet-aspnet-di
service: storage
plane: data-plane
language: dotnet
category: authentication
difficulty: intermediate
description: >
  Can a developer register Azure SDK clients in ASP.NET Core
  dependency injection using the recommended patterns?
sdk_package: Azure.Storage.Blobs
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/storage.blobs-readme
tags:
  - dependency-injection
  - aspnet-core
  - configuration
  - best-practices
created: 2026-04-09
author: JonathanCrd
---

# ASP.NET Core DI Integration: Azure Blob Storage (.NET)

## Prompt

I'm building an ASP.NET Core web API and need to use Azure Blob Storage
from my controllers and services. How do I register Azure SDK clients
with dependency injection properly?
1. Register an Azure Storage client in the service container at startup
2. Authenticate using identity-based credentials
3. Inject the client into controllers or service classes
4. Configure shared client options like retry policies

Explain why I should use the DI registration approach instead of
creating clients manually.

## Evaluation Criteria

- Registers Azure SDK clients using the ASP.NET Core DI extensions
- Configures identity-based authentication globally
- Injects clients into controllers or services via constructor injection
- Configures shared client options (retry, diagnostics) through DI
- Explains lifecycle benefits (singleton, connection pooling)

## Context

Real .NET developers wire up Azure SDK clients via dependency injection,
not by creating them inline. The Microsoft.Extensions.Azure package is the
recommended approach but many developers don't know it exists. This tests
whether generated code follows the production-ready pattern.
