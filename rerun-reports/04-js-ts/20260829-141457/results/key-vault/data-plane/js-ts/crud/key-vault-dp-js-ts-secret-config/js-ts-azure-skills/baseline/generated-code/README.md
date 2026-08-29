# Azure Key Vault configuration provider

A TypeScript configuration provider with managed-identity authentication,
in-memory caching, expiry warnings, versioned reads, and safe secret rotation.

## Run the offline demo

```powershell
npm install
npm start
```

The demo uses `InMemorySecretStore` and never contacts Azure. Production code
can call `createKeyVaultConfiguration()` after setting:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
```

`DefaultAzureCredential` automatically uses the application's managed identity
in Azure. No client secret or certificate is stored in this project. Assign the
identity only the Key Vault data-plane roles it needs.

## Rotation and cleanup

`SecretRotationHelper.rotate()` calls `setSecret`, which creates a new version
under the existing name. Key Vault cannot purge one version independently.
`deleteAndPurge()` is therefore an explicit destructive full-name cleanup: it
starts the long-running deletion, waits for completion, and only then purges
the deleted secret. It removes every version and requires purge permissions.
