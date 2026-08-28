# Azure App Configuration Java demo

This Java 17 project demonstrates synchronous and asynchronous configuration reads, labels,
prefix queries, ETag-aware caching, feature flags with deterministic percentage rollout, and
sentinel-driven cache refresh.

## App Configuration data

Create these settings in an existing App Configuration store:

| Key | Label | Example value |
| --- | --- | --- |
| `demo:message` | *(no label)* | `Hello from the default configuration` |
| `demo:message` | `production` | `Hello from production` |
| `demo:sentinel` | `production` | `1` |
| `.appconfig.featureflag/beta-dashboard` | `production` | See below |

Example feature flag payload:

```json
{
  "id": "beta-dashboard",
  "description": "Gradual dashboard rollout",
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

Assign the managed identity the **App Configuration Data Reader** role, then set the endpoint:

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
mvn compile exec:java
```

The demo runs the synchronous flow first, watches for ten seconds, and then repeats with the
asynchronous client. Change `demo:sentinel` while either watcher is running to trigger a full
refresh of all entries that the application has cached.
