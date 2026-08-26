# Azure App Configuration feature flags

This project demonstrates synchronous and asynchronous configuration access,
ETag-aware caching, deterministic percentage feature rollouts, and
sentinel-coordinated cache refreshes.

## Setup

Use Python 3.10 or newer:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Authenticate locally with a credential supported by `DefaultAzureCredential`,
then set the App Configuration data-plane endpoint:

```powershell
$env:AZURE_APPCONFIGURATION_ENDPOINT = "https://<store>.azconfig.io"
python main.py
```

The identity needs the **App Configuration Data Reader** role. For an
Azure-hosted production deployment, use managed identity and set
`AZURE_TOKEN_CREDENTIALS=prod` to restrict the credential chain.

The demo defaults can be changed with these environment variables:

| Variable | Default |
| --- | --- |
| `DEMO_CONFIG_KEY` | `demo:message` |
| `DEMO_CONFIG_LABEL` | `production` |
| `DEMO_CONFIG_PREFIX` | `demo:` |
| `DEMO_FEATURE_FLAG` | `gradual-rollout` |
| `DEMO_USER_IDS` | `alice,bob,charlie,diana` |
| `DEMO_SENTINEL_KEY` | `demo:sentinel` |
| `DEMO_POLL_INTERVAL` | `5` |
| `DEMO_MAX_POLLS` | `3` |

The main script runs the sync demo first, followed by the async demo. During
each finite watch window, change the sentinel value to trigger a full cache
refresh.

Feature flags are read from `.appconfig.featureflag/<name>`. Percentage rollout
uses the `Microsoft.Percentage` filter and hashes `<flag-id>:<user-id>` with
SHA-256, giving each user a stable bucket.

## Tests

Tests use fake clients and do not connect to Azure:

```powershell
python -m unittest discover -s tests -v
```

References:

- [Azure App Configuration Python SDK](https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient)
- [Azure App Configuration Python quickstart](https://learn.microsoft.com/azure/azure-app-configuration/quickstart-python)
- [DefaultAzureCredential](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
