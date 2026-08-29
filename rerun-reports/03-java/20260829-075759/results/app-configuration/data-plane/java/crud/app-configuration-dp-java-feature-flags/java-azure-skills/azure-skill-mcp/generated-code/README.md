# Azure App Configuration Java demo

Java 17 sample with synchronous and asynchronous Azure App Configuration clients, ETag-aware
caching, percentage feature flags, and sentinel-based refresh watchers.

## App Configuration data

Create these sample entries in an App Configuration store:

| Key | Label | Value |
|---|---|---|
| `App:Title` | `production` | `Production app` |
| `App:Title` | `staging` | `Staging app` |
| `Demo:Sentinel` | `production` | `1` |
| `.appconfig.featureflag/BetaCheckout` | `production` | JSON below |

```json
{
  "id": "BetaCheckout",
  "enabled": true,
  "conditions": {
    "client_filters": [
      {
        "name": "Microsoft.Percentage",
        "parameters": {
          "Value": 30
        }
      }
    ]
  }
}
```

Assign the managed identity the **App Configuration Data Reader** role on the store. Set
`AZURE_APPCONFIG_ENDPOINT` to the store endpoint. For a user-assigned managed identity, also set
`AZURE_CLIENT_ID`.

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
$env:AZURE_CLIENT_ID = "<optional-user-assigned-managed-identity-client-id>"
mvn test
mvn exec:java
```

The demo runs the synchronous flow first and then the asynchronous flow. Each watcher runs for
20 seconds and polls every 10 seconds; update the `Demo:Sentinel` value to trigger a complete
refresh of entries already held in the service cache.

## References

- [Azure App Configuration Java client library](https://learn.microsoft.com/java/api/overview/azure/data-appconfiguration-readme)
- [Authenticate Azure-hosted Java applications](https://learn.microsoft.com/azure/developer/java/sdk/authentication/azure-hosted-apps)
- [Azure App Configuration feature-management schema](https://learn.microsoft.com/azure/azure-app-configuration/feature-management-reference)
