# Azure service principal Python example

This project authenticates with a service principal and client secret, creates an
Azure Resource Management SDK client, and prints the resource groups in a
subscription.

## Setup

Create and activate a virtual environment, then install the required packages:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Set the following environment variables:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
python .\app.py
```

For local development, you can instead copy `.env.example` to `.env` and replace
its placeholders. `python-dotenv` loads that file without overriding variables
already present in the process environment.

## Secret management

- Prefer environment variables supplied by the deployment platform or a secret
  store such as Azure Key Vault. Do not hard-code credentials.
- Use `.env` only for local development. It is ignored by Git; never commit it,
  paste it into logs, or build it into a container image.
- Commit only `.env.example`, containing placeholders rather than real secrets.
- Grant the service principal only the minimum Azure RBAC role and scope needed.
- Rotate client secrets regularly and immediately after suspected disclosure.
  Prefer workload identity or managed identity when the hosting environment
  supports one, because those approaches avoid long-lived client secrets.

## Error behavior

The program exits with status `2` when configuration is missing, `3` when Azure
authentication fails, and `4` when Azure returns another HTTP error. SDK error
messages are logged, but credential values are never logged.
