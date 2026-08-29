# Azure App Configuration demo

A small Java 17 Maven application showing synchronous and asynchronous configuration reads,
ETag-based caching, feature-flag percentage rollout, and sentinel-driven cache refresh.

## Expected App Configuration entries

| Key | Label | Example value |
| --- | --- | --- |
| `Demo:Message` | no label | `Hello` |
| `Demo:Message` | `staging` | `Hello staging` |
| `Demo:Sentinel` | `staging` | `1` |
| `.appconfig.featureflag/BetaCheckout` | `staging` | JSON below |

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

The feature flag content type should be
`application/vnd.microsoft.appconfig.ff+json;charset=utf-8`.

## Run

Assign the managed identity the **App Configuration Data Reader** role, then set:

```powershell
$env:AZURE_APP_CONFIGURATION_ENDPOINT = "https://your-store.azconfig.io"
# Optional for a user-assigned managed identity:
$env:AZURE_CLIENT_ID = "00000000-0000-0000-0000-000000000000"
mvn compile exec:java
```

`CONFIG_POLL_INTERVAL_SECONDS` and `DEMO_WATCH_SECONDS` can override the demo's
two-second polling interval and five-second watcher duration.
