# Azure App Configuration demo

Java 17 example with synchronous and Reactor-based asynchronous configuration access, ETag-aware
conditional reads, percentage feature flags, and sentinel-driven cache refresh.

Expected sample data:

| Key | Label | Example value |
| --- | --- | --- |
| `app:message` | `production` | `Hello from production` |
| `app:message` | `staging` | `Hello from staging` |
| `app:sentinel` | no label | `1` |
| `.appconfig.featureflag/BetaFeature` | `production` | JSON below |

```json
{
  "id": "BetaFeature",
  "enabled": true,
  "conditions": {
    "client_filters": [
      {
        "name": "Microsoft.Percentage",
        "parameters": {
          "Value": "30"
        }
      }
    ]
  }
}
```

Build and run:

```powershell
mvn verify
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
mvn exec:java
```

The runtime environment must provide a managed identity with App Configuration Data Reader access.
The demo does not create or modify Azure resources.
