# Azure credential chain demo

This Python 3.9+ project chooses an explicit Azure Identity credential chain for
local development, CI/CD, or production and tests it against the Azure Resource
Manager token scope. It never prints access tokens.

## Setup and run

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
python main.py
python main.py --enable-cae
```

`AZURE_ENABLE_CAE=true` also enables Continuous Access Evaluation requests.
CAE is requested through `get_token(enable_cae=True)`; whether it is honored
depends on the selected credential and target resource.

## Environment selection

Set `AZURE_AUTH_ENVIRONMENT` to `dev`, `ci`, or `production` to override
automatic detection. Otherwise, CI markers take precedence over managed
identity endpoint, workload identity, and Azure hosting markers. A process with
none of those markers is considered local development.

| Environment | Credential order |
|---|---|
| `dev` | Azure CLI, Azure Developer CLI, Azure PowerShell, VS Code |
| `ci` | `EnvironmentCredential`, then Azure Pipelines workload identity when fully configured |
| `production` | Managed identity, then Kubernetes workload identity when fully configured |

For generic CI service principals, configure `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and either `AZURE_CLIENT_SECRET` or
`AZURE_CLIENT_CERTIFICATE_PATH`.

For an Azure Pipelines workload identity service connection, configure
`AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_SERVICE_CONNECTION_ID`, and map the
pipeline OAuth token to `SYSTEM_ACCESSTOKEN`.

Production uses a system-assigned managed identity by default. Set
`AZURE_MANAGED_IDENTITY_CLIENT_ID` to select a user-assigned identity. The
workload identity fallback is added when `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,
and `AZURE_FEDERATED_TOKEN_FILE` are all present.

## References

- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)
- [Credential chains in Azure Identity](https://aka.ms/azsdk/python/identity/credential-chains)
- [Continuous Access Evaluation in Azure Identity](https://learn.microsoft.com/python/api/overview/azure/identity-readme#continuous-access-evaluation)
