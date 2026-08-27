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

I want to use Azure App Configuration to manage feature flags for my .NET
application. How do I:
1. Store and manage feature flags in App Configuration
2. Read a feature flag and determine whether it is enabled
3. Evaluate feature flags in application code at runtime
4. Integrate feature flag evaluation into an ASP.NET Core app

Explain the difference between storing feature flags and evaluating them
in code.

## Evaluation Criteria

- Stores feature flags using the App Configuration SDK (e.g., `FeatureFlagConfigurationSetting`)
- Evaluates feature flags at runtime using a feature management library
- Distinguishes between the storage model and the evaluation model
- Integrates feature flag evaluation into ASP.NET Core dependency injection
- Uses the `.featureManagement` key prefix convention

## Context

Feature flags are a core App Configuration use case separate from key-value
configuration. Developers need to understand both the storage model
(FeatureFlagConfigurationSetting) and the evaluation model (FeatureManager)
to use them effectively.
