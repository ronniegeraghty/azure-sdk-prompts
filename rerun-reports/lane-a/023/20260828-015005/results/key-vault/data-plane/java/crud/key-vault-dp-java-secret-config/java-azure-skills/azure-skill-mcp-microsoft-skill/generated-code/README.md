# Key Vault configuration provider

A small Java 17 sample with synchronous and asynchronous Azure Key Vault secret
providers, expiry-aware in-memory caches, and soft-delete-aware secret rotation.

## Prerequisites

- Java 17 and Maven 3.9+
- An Azure-hosted workload with a system-assigned or user-assigned managed identity
- `AZURE_KEYVAULT_URL`, for example `https://my-vault.vault.azure.net`
- Optional `AZURE_CLIENT_ID` for a user-assigned managed identity

The identity needs secret read/list permissions. The rotation demo additionally
needs set, delete, get-deleted, and purge permissions. Purge protection prevents
immediate same-name delete-and-recreate rotation; in that case the helper fails
clearly rather than creating a false success.

## Run

```text
mvn compile exec:java
```

The demo expects `database-url`, `api-key`, and `feature-flag` to exist. Missing
configuration secrets use the defaults shown in `Main`. The
`rotating-demo-secret` must already exist because rotation deliberately deletes
the previous secret. Use a non-production vault intended for this demo.
