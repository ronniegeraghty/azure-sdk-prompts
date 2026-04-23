---
id: app-configuration-dp-dotnet-aspnet-provider
service: app-configuration
plane: data-plane
language: dotnet
category: crud
difficulty: intermediate
description: >
  Can a developer integrate Azure App Configuration with ASP.NET Core
  using the configuration provider for dynamic settings refresh?
sdk_package: Microsoft.Extensions.Configuration.AzureAppConfiguration
doc_url: https://learn.microsoft.com/en-us/azure/azure-app-configuration/quickstart-aspnet-core-app
tags:
  - aspnet-core
  - configuration-provider
  - dynamic-refresh
  - feature-flags
created: 2026-04-09
author: JonathanCrd
---

# ASP.NET Core Provider: Azure App Configuration (.NET)

## Prompt

I'm building an ASP.NET Core web app and want to use Azure App Configuration
as my centralized configuration source. My settings need to update at runtime
without redeploying the application. How do I:
1. Wire up my app to pull configuration from App Configuration at startup
2. Authenticate securely without hardcoding credentials
3. Load only specific keys and labels relevant to my environment
4. Have configuration values refresh automatically when they change in the store
5. Access the latest configuration values in my controllers and services

## Evaluation Criteria

- Integrates App Configuration into the ASP.NET Core configuration pipeline
- Uses `DefaultAzureCredential` or identity-based authentication
- Configures key filtering with labels for environment separation
- Implements dynamic refresh with a sentinel key pattern
- Exposes refreshed configuration via ASP.NET Core options patterns (e.g., `IOptionsSnapshot<T>` or `IOptionsMonitor<T>`)
- Includes the appropriate NuGet package for the App Configuration provider

## Context

The ASP.NET Core configuration provider is how most .NET developers actually
consume App Configuration in production — not via direct ConfigurationClient
calls. Dynamic refresh with sentinel keys is the recommended pattern for
reacting to configuration changes without application restarts.
