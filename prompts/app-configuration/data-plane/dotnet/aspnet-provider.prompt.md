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

I want to use Azure App Configuration as the configuration source for my
ASP.NET Core web app so settings update without redeploying. How do I:
1. Add the Azure App Configuration provider in Program.cs
2. Connect using a connection string or DefaultAzureCredential
3. Select which keys to load using key filters and labels
4. Enable dynamic refresh so settings update automatically
5. Use the refreshed settings in a controller via IOptionsSnapshot<T>

Show the required NuGet packages (Microsoft.Extensions.Configuration.AzureAppConfiguration)
and explain the sentinel key pattern for triggering refreshes.

## Evaluation Criteria

- `Microsoft.Extensions.Configuration.AzureAppConfiguration` NuGet package
- `builder.Configuration.AddAzureAppConfiguration()` in Program.cs
- `options.Connect()` with connection string or `DefaultAzureCredential`
- `options.Select()` for key filtering with labels
- `options.ConfigureRefresh()` with sentinel key and cache expiration
- `app.UseAzureAppConfiguration()` middleware for dynamic refresh
- `IOptionsSnapshot<T>` or `IOptionsMonitor<T>` for refreshed values

## Context

The ASP.NET Core configuration provider is how most .NET developers actually
consume App Configuration in production — not via direct ConfigurationClient
calls. Dynamic refresh with sentinel keys is the recommended pattern for
reacting to configuration changes without application restarts.
