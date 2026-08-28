# Azure service principal authentication (Python)

This example uses the OAuth 2.0 client credentials flow to authenticate a
non-interactive application with `ClientSecretCredential`. It passes that
credential to `ResourceManagementClient` and performs the read-only operation
of listing resource groups.

## Requirements

- Python 3.10 or later
- A Microsoft Entra service principal
- An Azure subscription
- A least-privilege Azure RBAC assignment that permits resource-group reads

The pip packages are declared in `requirements.txt`:

- `azure-identity` provides `ClientSecretCredential`.
- `azure-mgmt-resource` provides `ResourceManagementClient`.
- `python-dotenv` loads a local `.env` file during development.

## Setup

Create and activate a virtual environment:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

For local development, copy `.env.example` to `.env` and replace its
placeholders:

```powershell
Copy-Item .env.example .env
python main.py
```

Alternatively, set `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,
`AZURE_CLIENT_SECRET`, and `AZURE_SUBSCRIPTION_ID` in the process environment.
Process environment variables override values from `.env`.

The client secret must be the secret **value**, which is displayed only once
when the credential is created. It is not the secret ID.

## Secret-management practices

- Never hardcode or commit credentials. `.env` is ignored by Git and should be
  used only for local development.
- Store production secrets in a dedicated secret store, such as Azure Key
  Vault, or inject them securely through the deployment platform.
- Prefer managed identity or workload identity for Azure-hosted production
  workloads so no client secret is required.
- If a client secret is required, give the service principal only the minimum
  Azure RBAC permissions at the narrowest scope, set an expiration, rotate the
  secret regularly, and monitor service-principal sign-ins.
- Do not log secrets or enable credential/request-body logging in production.

The program returns exit code `2` for missing configuration, `3` for
authentication failures, and `4` for authorization or Azure API failures.

## References

- [Authenticate Azure-hosted Python apps](https://learn.microsoft.com/azure/developer/python/sdk/authentication-overview)
- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)
- [ClientSecretCredential class](https://learn.microsoft.com/python/api/azure-identity/azure.identity.clientsecretcredential)
- [Azure Resource Management client library](https://learn.microsoft.com/python/api/overview/azure/mgmt-resource-readme)
