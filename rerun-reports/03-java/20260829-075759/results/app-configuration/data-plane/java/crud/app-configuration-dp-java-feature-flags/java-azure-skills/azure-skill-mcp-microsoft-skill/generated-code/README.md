# Azure App Configuration Java demo

This Java 17 sample provides synchronous and Reactor-based asynchronous configuration
services, deterministic percentage feature flags, and sentinel-driven cache refresh.
Point reads use App Configuration ETags and conditional requests; prefix queries remain
cached until a watched sentinel changes.

## App Configuration data

Create these example settings in an existing App Configuration store:

| Key | Label | Example value |
|---|---|---|
| `app:greeting` | `production` | `Hello from production` |
| `app:greeting` | `staging` | `Hello from staging` |
| `app:sentinel` | `production` | `1` |
| `.appconfig.featureflag/beta-dashboard` | `production` | See below |

Example feature flag payload:

```json
{
  "id": "beta-dashboard",
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

The managed identity running the demo needs the **App Configuration Data Reader** role
on the store. No access keys or connection strings are used.

## Run

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
$env:DEMO_WATCH_SECONDS = "30"
mvn clean test exec:java
```

Change the `app:sentinel` value while each watcher is running to trigger a complete
refresh of all keys and prefix queries currently held in that implementation's cache.

SDK reference: [Azure App Configuration Java client library](https://learn.microsoft.com/java/api/overview/azure/data-appconfiguration-readme)
