# Azure service principal authentication with Python

This runnable example uses the OAuth 2.0 client credentials flow through
`ClientSecretCredential`, then passes that credential to the Azure Resource
Manager SDK to list resource groups. The Azure operation is read-only.

## Requirements

- Python 3.10 or later
- A Microsoft Entra service principal with a client secret
- The service principal assigned the least-privilege `Reader` role at the
  subscription or narrower scope that it must inspect

Install the required pip packages in a virtual environment:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

- `azure-identity` for `ClientSecretCredential`
- `azure-mgmt-resource` for `ResourceManagementClient`
- `python-dotenv` for optional local `.env` loading

## Configure and run

For local development, copy `.env.example` to `.env` and replace each
placeholder. The application calls `load_dotenv(override=False)`, so existing
environment variables take precedence over `.env` values.

Alternatively, set the values directly in the shell:

```powershell
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<application-client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret-value>"
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
python app.py
```

`create_credential` constructs the credential explicitly from `tenant_id`,
`client_id`, and `client_secret`. Before creating
`ResourceManagementClient`, the application requests an Azure Resource
Manager token so authentication failures are reported separately from
authorization, subscription, and network failures.

Run the local-only tests without making Azure requests:

```powershell
python -m unittest -v
```

## Secret-management practices

- Never hardcode or commit a client secret. `.env` is gitignored; commit only
  `.env.example` with placeholders.
- Treat `.env` as a local-development convenience, not a production secret
  store. Inject production secrets through the hosting platform or retrieve
  them from Azure Key Vault.
- Restrict access with least-privilege Azure RBAC at the narrowest practical
  scope, set secret expirations, and rotate secrets regularly.
- Do not log credentials or exception details that may contain sensitive
  request data. This example emits actionable messages without printing the
  secret.
- Prefer managed identity for Azure-hosted production workloads, or workload
  identity/certificate credentials where managed identity is unavailable, to
  avoid long-lived client secrets.
