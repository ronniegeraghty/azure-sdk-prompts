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

I need to manage TLS certificates in Azure Key Vault using C#.
How do I use the CertificateClient to:
1. Create a self-signed certificate for testing
2. Wait for the certificate creation to complete (it's an LRO)
3. Download the certificate with its private key as an X509Certificate2
4. Import an existing PFX certificate into Key Vault

Use DefaultAzureCredential for authentication. Show required NuGet packages.

## Evaluation Criteria

- `Azure.Security.KeyVault.Certificates` NuGet package
- `CertificateClient` creation with vault URI and credential
- `StartCreateCertificateAsync()` with `CertificatePolicy.Default`
- `WaitForCompletionAsync()` on the certificate operation
- `CertificateClient.DownloadCertificateAsync()` returning `X509Certificate2`
- `ImportCertificateAsync()` with `ImportCertificateOptions`
- `RequestFailedException` handling

## Context

Certificate management in Key Vault is critical for TLS and code-signing
scenarios. Unlike secrets, certificates involve async creation (LRO) and
the distinction between the certificate, key, and secret portions is
a common source of confusion for developers.
