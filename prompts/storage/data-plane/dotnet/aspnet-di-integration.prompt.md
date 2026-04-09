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

I'm building an ASP.NET Core web API and need to inject a BlobServiceClient
into my controllers. How do I register Azure SDK clients with dependency
injection the right way?
1. Register BlobServiceClient in Program.cs using AddAzureClients
2. Configure it with DefaultAzureCredential
3. Inject BlobServiceClient into a controller or service class
4. Configure client options (retry policy, diagnostics) via DI

Show the required NuGet packages (Microsoft.Extensions.Azure) and explain
why I should use AddAzureClients instead of manually newing up clients.

## Evaluation Criteria

- `Microsoft.Extensions.Azure` NuGet package
- `builder.Services.AddAzureClients()` registration in Program.cs
- `AddBlobServiceClient()` with URI
- `UseCredential(new DefaultAzureCredential())` global credential
- Constructor injection of `BlobServiceClient` into services/controllers
- `ConfigureDefaults()` for shared client options
- Explanation of singleton lifecycle and connection pooling benefits

## Context

Real .NET developers wire up Azure SDK clients via dependency injection,
not by creating them inline. The Microsoft.Extensions.Azure package is the
recommended approach but many developers don't know it exists. This tests
whether generated code follows the production-ready pattern.
