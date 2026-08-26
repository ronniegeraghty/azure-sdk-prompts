# Azure Key Vault configuration provider

Small Java 17 sample with synchronous and asynchronous Key Vault secret providers, in-memory
caches, expiry-aware refresh, managed identity authentication, and delete-and-recreate rotation.
Secret values are never printed.

## Configure and run

The managed identity needs Key Vault data-plane permissions to read, delete, purge, and create
secrets. Delete-and-recreate rotation requires purge protection to be disabled; if organizational
policy enables purge protection, use normal Key Vault version rotation (`setSecret`) instead.

Set these environment variables:

```text
AZURE_KEY_VAULT_URL=https://<vault-name>.vault.azure.net
AZURE_MANAGED_IDENTITY_CLIENT_ID=<optional-user-assigned-managed-identity-client-id>
DEMO_SYNC_ROTATION_SECRET_NAME=<existing-secret-name>
DEMO_SYNC_ROTATION_NEW_VALUE=<new-secret-value>
DEMO_ASYNC_ROTATION_SECRET_NAME=<different-existing-secret-name>
DEMO_ASYNC_ROTATION_NEW_VALUE=<new-secret-value>
```

Run:

```text
mvn test
mvn exec:java
```

The demo bulk-loads `database-connection-string`, `api-base-url`, and `feature-flags`, reads each
from cache, refreshes `api-base-url`, refreshes secrets within the seven-day warning window, prints
expiry warnings without exposing values, and rotates one sync and one async demo secret.

## Design notes

- Only `ResourceNotFoundException` is converted to a caller-provided default. Authentication,
  authorization, throttling, and service failures remain visible.
- A cached secret inside the warning window is automatically fetched again when read. The
  `refreshExpiring` methods support proactive sweeps as well.
- Rotation waits for the delete long-running operation, purges the soft-deleted secret, polls until
  the deleted name is no longer visible, and only then creates the replacement. Polling has a
  two-minute timeout.
- Azure SDK versions are managed by the current stable Azure SDK BOM (`1.3.8`).

## References

- [Azure Key Vault Secret client library for Java quickstart](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-java)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
- [Azure SDK for Java Key Vault Secrets samples](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/keyvault/azure-security-keyvault-secrets/src/samples/java/com/azure/security/keyvault/secrets)
