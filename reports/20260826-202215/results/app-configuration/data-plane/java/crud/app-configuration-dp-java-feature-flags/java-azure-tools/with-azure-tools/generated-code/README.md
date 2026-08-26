# Azure App Configuration Java demo

Java 17 Maven sample with synchronous and asynchronous configuration services, ETag-aware caching,
percentage feature flags, and sentinel-based refresh watchers.

## App Configuration data

Create these settings in an existing App Configuration store:

| Key | Label | Example value |
|---|---|---|
| `application:name` | *(no label)* | `Demo application` |
| `application:message` | `production` | `Hello from production` |
| `application:sentinel` | `production` | `1` |
| `.appconfig.featureflag/BetaDashboard` | `production` | See below |

```json
{
  "id": "BetaDashboard",
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

Assign the managed identity the **App Configuration Data Reader** role on the store. The demo uses a
system-assigned identity by default. Set `AZURE_CLIENT_ID` to select a user-assigned managed identity.

## Run

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store-name>.azconfig.io"
$env:CONFIG_POLL_SECONDS = "10"
mvn compile exec:java
```

The demo runs the synchronous flow first and then the asynchronous flow. While either watcher is
running, update the sentinel value to cause all settings previously read by that service to refresh.

SDK references:

- https://learn.microsoft.com/java/api/overview/azure/data-appconfiguration-readme
- https://learn.microsoft.com/java/api/overview/azure/identity-readme
