# Environment-specific Azure credential chains

This local sample selects an explicit Azure Identity credential chain for
development, CI/CD, or production, then requests an Azure Resource Manager
token with both the synchronous and asynchronous APIs.

## Set up and run

Python 3.9 or later is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
pip install -r requirements.txt
python main.py
python main.py --enable-cae
```

The sample only requests a token. It does not create, update, or delete Azure
resources.

## Environment selection

`environment_detector.py` checks CI markers first, then Azure managed identity,
Azure hosting, and Kubernetes workload identity markers. It defaults to local
development. Set `APP_ENVIRONMENT` to `dev`, `ci`, or `production` to override
automatic detection.

| Environment | Credential order |
|---|---|
| Development | Azure CLI, Azure Developer CLI, Azure PowerShell, VS Code |
| CI/CD | Azure Pipelines workload identity service connection, then environment credential |
| Production | Managed identity, then Kubernetes workload identity when configured |

For an Azure Pipelines workload identity service connection, expose
`AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_SERVICE_CONNECTION_ID`, and
`SYSTEM_ACCESSTOKEN`. Otherwise, CI uses `EnvironmentCredential`, configured
with the standard Azure Identity service-principal secret or certificate
variables.

Production uses a system-assigned managed identity by default. Set
`AZURE_MANAGED_IDENTITY_CLIENT_ID` to select a user-assigned managed identity.
The workload identity fallback is added when `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE` are all present.

CAE is requested per token acquisition by passing `--enable-cae`. Whether the
issued token is CAE-capable is determined by the identity provider and resource.

Azure Identity reference:
https://learn.microsoft.com/python/api/overview/azure/identity-readme
