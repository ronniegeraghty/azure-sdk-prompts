---
id: key-vault-dp-dotnet-certificates
service: key-vault
plane: data-plane
language: dotnet
category: crud
difficulty: intermediate
description: >
  Can a developer create, import, and retrieve certificates from
  Azure Key Vault using the .NET SDK?
sdk_package: Azure.Security.KeyVault.Certificates
doc_url: https://learn.microsoft.com/en-us/dotnet/api/overview/azure/security.keyvault.certificates-readme
tags:
  - certificates
  - tls
  - x509
  - crud
created: 2026-04-09
author: JonathanCrd
---

# CRUD Certificates: Azure Key Vault (.NET)

## Prompt

I need to manage TLS certificates in Azure Key Vault using .NET.
How do I:
1. Create a self-signed certificate for testing
2. Wait for the certificate creation to complete (it's a long-running operation)
3. Download the certificate with its private key so I can use it in my app
4. Import an existing PFX certificate into Key Vault

Authenticate securely using identity-based credentials.

## Evaluation Criteria

- Uses the Key Vault Certificates SDK to create, download, and import certificates
- Handles the long-running operation pattern for certificate creation
- Downloads the certificate with its private key in a usable format
- Imports an existing PFX certificate
- Authenticates with `DefaultAzureCredential` or identity-based credentials
- Handles errors appropriately

## Context

Certificate management in Key Vault is critical for TLS and code-signing
scenarios. Unlike secrets, certificates involve async creation (LRO) and
the distinction between the certificate, key, and secret portions is
a common source of confusion for developers.
