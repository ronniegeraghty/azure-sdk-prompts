# Azure Key Vault configuration provider

A Java 17 Maven sample with synchronous and asynchronous Key Vault secret
providers, in-memory caches, expiry-aware refresh, and soft-delete-aware secret
rotation.

## Authentication and permissions

The client uses `ManagedIdentityCredential`; no application secret or certificate
is stored in code. Assign the workload's managed identity the least-privilege Key
Vault data-plane permissions needed by the demo:

- secrets: get, set, delete, purge

`purge` is required because a soft-deleted secret name cannot be recreated until
the deleted object has been purged. Rotation by delete/recreate is incompatible
with purge protection until the retention period ends; for purge-protected
production vaults, prefer creating a new secret version instead.

Set these environment variables:

```text
AZURE_KEYVAULT_URL=https://your-vault.vault.azure.net
AZURE_CLIENT_ID=<optional user-assigned managed identity client ID>
DEMO_ROTATED_SECRET_VALUE=<new demo value>
```

`AZURE_CLIENT_ID` is omitted for a system-assigned managed identity.

## Build and run

```text
mvn clean package
mvn exec:java
```

The demo runs the synchronous flow first and then the asynchronous flow. Both
load required keys, read the cache without logging secret values, refresh one
key, refresh near-expiry entries, print expiry warnings, and rotate
`rotating-demo-secret`.

Missing required configuration secrets use the defaults declared in `Main`.
Authentication, authorization, throttling, and other service errors are not
treated as missing secrets and remain visible to the caller.

SDK reference:
https://learn.microsoft.com/java/api/overview/azure/security-keyvault-secrets-readme
