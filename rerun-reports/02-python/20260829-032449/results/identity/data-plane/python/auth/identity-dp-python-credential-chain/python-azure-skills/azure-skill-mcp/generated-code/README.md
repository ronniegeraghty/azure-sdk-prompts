# Environment-specific Azure credential chains

This sample builds explicit Azure Identity chains instead of using one broad
credential everywhere. It only requests an Azure Resource Manager token; it
does not create or modify Azure resources.

## Setup and run

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
python main.py
```

Set `AZURE_ENABLE_CAE=true` to request a CAE-enabled token. Azure Resource
Manager and the selected identity flow must support CAE for the resulting token
to have CAE semantics.

Use `APP_ENV=dev`, `APP_ENV=ci`, or `APP_ENV=production` to override automatic
detection.

## Credential configuration

| Environment | Credential order | Configuration |
| --- | --- | --- |
| Development | VS Code, Azure CLI, Azure PowerShell, Azure Developer CLI | Sign in with the corresponding developer tool. VS Code authentication requires the Azure Resources extension and `azure-identity-broker`. |
| CI | Azure Pipelines service connection when fully configured; otherwise environment credential | For Azure Pipelines, expose `SYSTEM_ACCESSTOKEN`, `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_SERVICE_CONNECTION_ID` as secret pipeline variables. Other CI systems can configure the service principal variables supported by `EnvironmentCredential`. |
| Production | Managed identity, then configured Kubernetes workload identity | Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` for a user-assigned managed identity; omit it for system-assigned. Workload identity uses `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE`. |

Run the offline tests with:

```powershell
python -m unittest discover -s tests -v
```

References:

- https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains
- https://learn.microsoft.com/python/api/azure-identity/azure.identity.azurepipelinescredential
- https://learn.microsoft.com/python/api/azure-identity/azure.identity.workloadidentitycredential
