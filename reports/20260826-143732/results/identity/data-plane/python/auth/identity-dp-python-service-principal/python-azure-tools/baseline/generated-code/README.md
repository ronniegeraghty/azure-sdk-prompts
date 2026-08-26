# Azure service principal Python example

This project creates an `azure.identity.ClientSecretCredential` and uses it
with `azure.mgmt.resource.ResourceManagementClient`. Its default command only
constructs the objects, so it is safe to run offline. The optional resource
group command performs a read-only Azure request.

## Setup

Create and activate a virtual environment, then install the required packages:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Set the four required values in the process environment:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
python .\azure_sp_example.py
```

For local development, copy `.env.example` to `.env` and replace its
placeholders. `python-dotenv` loads that file without overriding values already
present in the process environment. `.env` is ignored by Git.

To authenticate and list resource groups through the Azure SDK:

```powershell
python .\azure_sp_example.py --list-resource-groups
```

The service principal needs only the least-privileged read role required for
the target subscription or resource scope. The example never creates, changes,
or deletes Azure resources.

## Secret management

- Prefer workload identity or managed identity when the runtime supports it.
- In CI/CD or production, inject the secret from the platform's secret store;
  do not use a checked-in `.env` file.
- Never log the secret, include it in command-line arguments, or commit `.env`.
- Scope permissions narrowly, rotate secrets regularly, use short expiration
  periods, and revoke unused credentials.

The CLI reports missing configuration separately from authentication and Azure
HTTP failures. It does not print access tokens or the client secret.

## Tests

Tests use mocks and make no network requests:

```powershell
python -m unittest -v
```
