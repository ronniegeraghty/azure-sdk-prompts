---
id: app-configuration-dp-dotnet-crud
properties:
  service: app-configuration
  plane: data-plane
  language: dotnet
  category: crud
  difficulty: basic
  description: 'Can a developer read and write configuration values and feature flags in Azure App Configuration using the
    .NET SDK?

    '
  sdk_package: Azure.Data.AppConfiguration
  doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/data.appconfiguration-readme
  created: '2025-07-28'
  author: ronniegeraghty
tags:
- app-configuration
- configuration
- feature-flags
- crud
---

# Configuration Values: Azure App Configuration (.NET)

## Prompt

How do I read and write configuration values in Azure App Configuration
using the .NET SDK? I need to:
1. Connect to my App Configuration store securely
2. Set a configuration value with a specific key and value
3. Set a configuration value with a label for environment separation
4. Retrieve a setting by key and print its value
5. List all settings matching a key prefix
6. Delete a setting

Include proper error handling for failed operations.

## Evaluation Criteria

The generated code should include:
- Creates a configuration client with identity-based authentication or connection string
- Sets configuration values with key, value, and optional label
- Retrieves settings by key and label
- Filters and lists settings using a selector or prefix
- Deletes settings and handles errors appropriately

## Context

App Configuration centralizes application settings. This tests whether the .NET
docs cover the key-value model including labels and the feature flag pattern,
which are the two primary usage scenarios.
