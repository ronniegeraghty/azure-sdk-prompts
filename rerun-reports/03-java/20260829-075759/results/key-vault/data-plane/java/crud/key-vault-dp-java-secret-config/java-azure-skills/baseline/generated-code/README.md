# Azure Key Vault configuration provider

A small Java 17 example with synchronous and Reactor-based asynchronous providers,
in-memory expiry-aware caches, managed-identity authentication, and safe
delete/purge/recreate rotation.

## Configuration

The application requires:

- `KEY_VAULT_URL`: an HTTPS Key Vault URL such as
  `https://my-vault.vault.azure.net`
- `ROTATED_SECRET_VALUE`: the new value used for the demo rotation

The Azure workload must have a managed identity with permissions to read, set,
delete, purge, and recover secrets. Purge permission and a vault configuration
that permits purging are required by the rotation demo. No credentials are
stored in this project.

## Build and run

```powershell
mvn test
mvn exec:java
```

`Main` runs the synchronous flow first and the asynchronous flow second. Both
flows rotate `demo-rotating-secret`; use a non-production vault and secret when
running the demo.
