---
id: example-expected-packages
properties:
  service: cosmos-db
  plane: data-plane
  language: python
  category: crud
  difficulty: intermediate
  description: 'Example prompt demonstrating expected_packages and expected_tools fields'
  sdk_package: azure-cosmos
  doc_url: https://learn.microsoft.com/python/api/azure-cosmos
  created: '2025-04-04'
  author: hyoka-examples
tags:
- example
- expected-validation
expected_packages:
- azure-cosmos>=4.5.0
- azure-identity>=1.15.0
expected_tools:
- pip
- python3
---

# Example: Expected Packages and Tools

## Prompt

Create a Python script that connects to Azure Cosmos DB and performs basic CRUD operations
on a container. Use the azure-cosmos SDK with DefaultAzureCredential for authentication.

The script should:
- Install required packages using pip
- Create a database client with credential authentication
- Perform create, read, update, delete operations on items

## Evaluation Criteria

The generated code should:
- Install `azure-cosmos` and `azure-identity` packages
- Use Python 3.9+ features
- Include proper error handling
- Use DefaultAzureCredential for authentication

## Context

This example demonstrates how `expected_packages` and `expected_tools` can be used
to validate that the generator:
1. Installs the correct SDK packages with appropriate version constraints
2. Uses the expected development tools (pip, python3)

These fields help measure whether the AI agent correctly identified required dependencies.
