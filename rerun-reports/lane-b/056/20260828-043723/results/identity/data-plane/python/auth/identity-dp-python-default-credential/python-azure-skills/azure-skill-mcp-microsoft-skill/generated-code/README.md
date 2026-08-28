# Authenticate an Azure SDK client with `DefaultAzureCredential`

This sample authenticates a Python `BlobServiceClient` with Microsoft Entra ID.
It uses no account keys or connection strings and works with a developer identity
locally or a workload identity in Azure.

## 1. Install the packages

Python 3.9 or later is required.

For Azure CLI, Azure PowerShell, Azure Developer CLI, managed identity, workload
identity, or environment-based authentication:

```powershell
python -m pip install -r requirements.txt
```

The packages are:

- `azure-identity`: supplies `DefaultAzureCredential`.
- `azure-storage-blob`: supplies the example `BlobServiceClient`. Replace this
  with the package for the Azure service your application uses.

For sign-in through the VS Code Azure Resources extension or the Windows/WSL
authentication broker, install the optional broker dependency:

```powershell
python -m pip install -r requirements-vscode.txt
```

`azure-core` is installed transitively by Azure SDK packages.

## 2. Create and use the credential

`authenticate.py` creates one `DefaultAzureCredential`, passes it to
`BlobServiceClient`, and reuses it for the lifetime of that client. The Azure SDK
requests and refreshes access tokens automatically; application code should not
read or store tokens.

Set the Blob service endpoint, then run the example:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
python authenticate.py
```

The authenticated identity needs an appropriate data-plane role, such as
**Storage Blob Data Reader**, on the storage account or container. Authentication
proves identity; Azure RBAC separately determines authorization.

Both the credential and client are context managers so their transports are
closed predictably. A single credential can be shared by multiple Azure SDK
clients in the same process.

## 3. Default credential chain

For `azure-identity` 1.24 or later, `DefaultAzureCredential` attempts the
following credentials in order and stops when one returns a token:

| Order | Credential | When it is usable |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal variables such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or certificate settings are complete. |
| 2 | `WorkloadIdentityCredential` | Federated workload identity variables and token file are configured, commonly by AKS workload identity. |
| 3 | `ManagedIdentityCredential` | The Azure host has a system-assigned or user-assigned managed identity. |
| 4 | `SharedTokenCacheCredential` | On Windows, a supported Microsoft application such as Visual Studio has cached a signed-in user. |
| 5 | `VisualStudioCodeCredential` | The Azure Resources extension is signed in and `azure-identity-broker` is installed. |
| 6 | `AzureCliCredential` | Azure CLI has an active `az login` session. |
| 7 | `AzurePowerShellCredential` | Azure PowerShell has an active `Connect-AzAccount` session. |
| 8 | `AzureDeveloperCliCredential` | Azure Developer CLI has an active `azd auth login` session. |
| 9 | `InteractiveBrowserCredential` | Only when explicitly enabled with `exclude_interactive_browser_credential=False`; it is disabled by default. |
| 10 | Broker credential | On Windows or WSL, `azure-identity-broker` can use the account known to the OS broker. |

The exact chain can evolve with `azure-identity`; consult the API reference for
the installed version. Constructor `exclude_*` options can remove entries.
Starting with `azure-identity` 1.24, `AZURE_TOKEN_CREDENTIALS` can narrow the
chain without changing code:

```powershell
# Use only developer-tool credentials locally.
$env:AZURE_TOKEN_CREDENTIALS = "dev"

# Use only deployed-service credentials in Azure.
$env:AZURE_TOKEN_CREDENTIALS = "prod"

# Or require one exact credential.
$env:AZURE_TOKEN_CREDENTIALS = "ManagedIdentityCredential"
```

Set `require_envvar=True` on `DefaultAzureCredential` if the application must
fail unless `AZURE_TOKEN_CREDENTIALS` is explicitly configured.

## 4. Local development and Azure deployment

### Azure CLI

Sign in with `az login`; `AzureCliCredential` then obtains tokens for that
developer account. If several tenants or subscriptions are available, select the
intended context in the CLI. The user must have the same data-plane permissions
the application operation requires.

For deterministic local behavior, set:

```powershell
$env:AZURE_TOKEN_CREDENTIALS = "AzureCliCredential"
```

### VS Code

Install the VS Code **Azure Resources** extension, sign in from its Azure view,
and install `requirements-vscode.txt`. `VisualStudioCodeCredential` then uses
that signed-in account through the broker package.

For deterministic VS Code behavior, set:

```powershell
$env:AZURE_TOKEN_CREDENTIALS = "VisualStudioCodeCredential"
```

### Azure-hosted application

Enable a managed identity on App Service, Functions, a VM, or another supported
host, then grant that identity the least-privileged Azure RBAC role needed by the
application. No secret is placed in code or configuration.

- A system-assigned identity needs no identity selector.
- For a user-assigned managed identity, set `AZURE_CLIENT_ID` to its client ID
  or pass `managed_identity_client_id` to `DefaultAzureCredential`.
- On AKS with Microsoft Entra Workload ID, the webhook supplies the tenant,
  client, and federated-token settings used by `WorkloadIdentityCredential`.

Use `AZURE_TOKEN_CREDENTIALS=prod` to exclude developer credentials in an Azure
deployment, or select `ManagedIdentityCredential`/`WorkloadIdentityCredential`
explicitly for the most deterministic production behavior. The application code
and SDK client construction remain otherwise unchanged.

## 5. Troubleshoot failures with logging

Run the sample with Azure Identity debug logging:

```powershell
python authenticate.py --debug-auth
```

The logs show each attempted credential, why it was unavailable, and which
credential succeeded. The sample distinguishes:

- `CredentialUnavailableError`: no chain member had usable configuration.
- `ClientAuthenticationError`: a credential attempted sign-in but token
  acquisition failed; its message normally includes per-credential details.
- HTTP 403 from Blob Storage: authentication succeeded, but the identity lacks
  the required Azure RBAC data-plane role or the role assignment has not
  propagated.

Check these common causes:

1. Confirm the relevant developer tool is signed in, or that managed/workload
   identity is enabled on the Azure host.
2. Remove stale or incomplete `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,
   `AZURE_CLIENT_SECRET`, and `AZURE_FEDERATED_TOKEN_FILE` variables. An
   accidentally configured earlier credential can change chain behavior.
3. Confirm the tenant is correct and the selected identity has the necessary
   Azure RBAC role at the correct scope.
4. Set `AZURE_TOKEN_CREDENTIALS` to the expected credential to isolate the
   failing authentication path.
5. Allow time for new role assignments to propagate.

Identity-only logging is enabled by the sample. Avoid enabling full Azure HTTP
logging in routine use: `logging_enable=True` at DEBUG level can expose sensitive
headers and request data. If it is temporarily necessary, protect and delete the
logs after diagnosis.

## References

- [Credential chains in Azure Identity for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains)
- [`DefaultAzureCredential` API reference](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
- [Local developer-account authentication](https://learn.microsoft.com/azure/developer/python/sdk/authentication/local-development-dev-accounts)
- [Azure SDK logging for Python](https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-logging)
