# Azure App Configuration Python Demo

This project demonstrates cached synchronous and asynchronous configuration reads,
feature-flag evaluation, percentage rollouts, and sentinel-based refresh with Azure
App Configuration.

## Setup

1. Create a virtual environment and install dependencies:

   `python -m venv .venv`

   `.venv\Scripts\python -m pip install -r requirements.txt`

2. Authenticate locally with a credential supported by `DefaultAzureCredential`,
   and grant that identity the **App Configuration Data Reader** role.

3. Set the endpoint and optional demo settings:

   `$env:AZURE_APPCONFIG_ENDPOINT = "https://your-store.azconfig.io"`

   `$env:AZURE_APPCONFIG_LABEL = "production"`

   `$env:CONFIG_WATCH_SECONDS = "15"`

4. Add `App:Message`, `App:Sentinel`, and a `BetaFeature` feature flag to the
   selected label, then run:

   `python main.py`

The sync demo runs first, followed by the async demo. Point reads use ETags and
conditional requests. Prefix reads first retrieve key/label/ETag metadata and only
download values whose ETags changed. Updating a sentinel causes all configuration
previously read by that service to be rebuilt.
