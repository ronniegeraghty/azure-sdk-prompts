---
id: app-configuration-dp-dotnet-feature-flags
service: app-configuration
plane: data-plane
language: dotnet
category: crud
difficulty: intermediate
description: >
  Can a developer manage feature flags in Azure App Configuration
  and evaluate them in code using the .NET SDK?
sdk_package: Azure.Data.AppConfiguration
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/data.appconfiguration-readme
tags:
  - feature-flags
  - feature-management
  - toggles
created: 2026-04-09
author: JonathanCrd
---

# Feature Flags: Azure App Configuration (.NET)

## Prompt

I want to use Azure App Configuration to manage feature flags for my .NET app.
How do I:
1. Create a feature flag in App Configuration using ConfigurationClient
2. Read the feature flag and check if it's enabled
3. Use Microsoft.FeatureManagement to evaluate feature flags in code
4. Wire up feature flag evaluation from App Configuration in ASP.NET Core

Show the required NuGet packages and explain the difference between storing
feature flags via ConfigurationClient vs. evaluating them with FeatureManager.

## Evaluation Criteria

- `FeatureFlagConfigurationSetting` for creating feature flags via SDK
- `ConfigurationClient.SetConfigurationSetting()` with feature flag settings
- `Microsoft.FeatureManagement` NuGet package for evaluation
- `IFeatureManager.IsEnabledAsync()` for checking flags in code
- `.featureManagement` key prefix convention
- Integration with ASP.NET Core via `AddFeatureManagement()`

## Context

Feature flags are a core App Configuration use case separate from key-value
configuration. Developers need to understand both the storage model
(FeatureFlagConfigurationSetting) and the evaluation model (FeatureManager)
to use them effectively.
