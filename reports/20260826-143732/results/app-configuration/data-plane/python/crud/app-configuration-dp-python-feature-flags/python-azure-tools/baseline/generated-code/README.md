# Azure App Configuration demo

This project provides cached synchronous and asynchronous configuration services,
deterministic percentage-based feature flags, and sentinel-driven configuration
watchers.

## Setup

Create the demo entries in Azure App Configuration:

- `Demo:ApiUrl`, with `production` and `staging` labels
- One or more `Demo:` settings with a `staging` label
- `Demo:Sentinel` with a `production` label
- `.appconfig.featureflag/NewCheckout` with a `production` label and a
  `Microsoft.Percentage` client filter

Then install and run:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store>.azconfig.io"
$env:DEMO_WATCH_SECONDS = "30"
python main.py
```

`DefaultAzureCredential` uses local developer credentials or managed identity.
The demo runs the synchronous flow first and then the asynchronous flow.
