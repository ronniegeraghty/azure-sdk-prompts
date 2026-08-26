# Authenticate an Azure SDK client with `DefaultAzureCredential`

This example creates an Azure Blob Storage client without account keys or
connection strings. `DefaultAzureCredential` obtains a Microsoft Entra access
token from the environment where the application runs, and the Blob client
automatically requests and refreshes tokens when it sends requests.

## 1. Install the packages

Python 3.9 or later is required.

```powershell
py -m venv .venv
.\.venv\Scripts\Activate.ps1
py -m pip install -r requirements.txt
```

The example needs:

- `azure-identity`: provides `DefaultAzureCredential`.
- `azure-storage-blob`: provides the example `BlobServiceClient`.
- `azure-identity-broker` (optional): enables brokered authentication and is
  also required for `VisualStudioCodeCredential` in current Azure Identity
  versions.

Install the optional integration with:

```powershell
py -m pip install azure-identity-broker
```

`azure-core` is installed transitively by the Azure packages; it does not need
to be listed separately.

## 2. Create and use the credential

Set the non-secret Blob service endpoint:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<storage-account-name>.blob.core.windows.net"
```

Run the offline-safe path, which constructs and closes both the credential and
client but sends no request:

```powershell
py app.py
```

To force token acquisition and make a read-only request:

```powershell
py app.py --list-containers
```

The signed-in identity needs an appropriate **data-plane** role, such as
`Storage Blob Data Reader`, scoped as narrowly as practical. Management roles
such as `Contributor` do not automatically grant access to blob data.

`app.py` creates one `DefaultAzureCredential` and passes it to
`BlobServiceClient`. Both are context managers so their underlying transports
are closed. Reuse a credential instance across clients in a long-running
application instead of constructing one per request.

## 3. Default credential chain

With the current `azure-identity` package, credentials are attempted in this
order and the chain stops when one returns a token:

| Order | Credential | What it uses |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal environment variables, such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or certificate settings |
| 2 | `WorkloadIdentityCredential` | Federated workload identity configuration, commonly on AKS |
| 3 | `ManagedIdentityCredential` | A system-assigned or user-assigned Azure managed identity |
| 4 | `SharedTokenCacheCredential` | A cached Microsoft application/Visual Studio sign-in on Windows |
| 5 | `VisualStudioCodeCredential` | The account selected by the VS Code Azure Resources extension; requires `azure-identity-broker` |
| 6 | `AzureCliCredential` | The account selected by `az login` |
| 7 | `AzurePowerShellCredential` | The account selected by `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | The account selected by `azd auth login` |
| 9 | `BrokerCredential` | The Windows/WSL Web Account Manager account when `azure-identity-broker` is installed |

`InteractiveBrowserCredential` is excluded by default. It can be appended by
constructing `DefaultAzureCredential(
exclude_interactive_browser_credential=False)`, but interactive authentication
is generally unsuitable for services.

Since Azure Identity 1.14.0, failures from developer-tool credentials do not
prevent later developer credentials from being attempted. Deployed-service
credentials have stricter behavior: if one is configured and attempts token
acquisition but authentication fails, the chain stops and reports that failure.
An unavailable credential is skipped.

The chain can be narrowed without changing the code:

- `AZURE_TOKEN_CREDENTIALS=dev` keeps developer-tool credentials.
- `AZURE_TOKEN_CREDENTIALS=prod` keeps deployed-service credentials
  (`EnvironmentCredential`, workload identity, and managed identity).
- A specific name, such as `AzureCliCredential` or
  `ManagedIdentityCredential`, keeps only that credential.

Use `prod` in Azure deployments to prevent an accidental developer sign-in
from becoming the service identity.

## 4. Local development and Azure deployments

### Local development

For Azure CLI authentication:

```powershell
az login
az account show
py app.py --list-containers
```

If more than one subscription or tenant is available, select the intended
subscription with `az account set`. The CLI login only establishes identity;
that identity must still have the required Storage data-plane role.

For VS Code, install the Azure Resources extension, sign in to Azure from VS
Code, and install `azure-identity-broker`. `DefaultAzureCredential` can then use
the selected VS Code account. If both VS Code and Azure CLI are signed in, VS
Code is earlier in the chain.

Avoid setting service-principal environment variables on a developer machine
unless they are intentional: `EnvironmentCredential` is first and can take
precedence over the developer login.

### Azure deployment

Enable a system-assigned managed identity on the Azure host and grant it the
least-privileged data-plane role needed by the application. The same code then
uses `ManagedIdentityCredential` through the default chain.

For a user-assigned managed identity, set:

```text
AZURE_CLIENT_ID=<managed-identity-client-id>
AZURE_TOKEN_CREDENTIALS=prod
```

AKS commonly uses Microsoft Entra Workload ID instead; the admission webhook
supplies the tenant, client ID, and federated token file variables used by
`WorkloadIdentityCredential`.

Do not deploy developer logins, account keys, or client secrets when managed or
workload identity is available. For a tightly controlled production service,
using `ManagedIdentityCredential` or `WorkloadIdentityCredential` directly is
even more explicit than a chain.

## 5. Troubleshoot authentication failures

Enable Azure Identity diagnostics for this example:

```powershell
$env:AZURE_SDK_LOG_LEVEL = "DEBUG"
py app.py --list-containers
```

The logs show each attempted credential, why it was unavailable or failed, and
which credential succeeded. DEBUG logs can contain tenant IDs, client IDs,
object IDs, request URLs, and other sensitive metadata; use them temporarily
and sanitize them before sharing.

Check failures in this order:

1. Confirm which credential actually succeeded or stopped the chain in the
   `azure.identity` logs.
2. For CLI auth, run `az account show` and verify the tenant and subscription.
   For VS Code, verify the selected Azure account and broker package.
3. Remove incomplete `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, or
   `AZURE_CLIENT_SECRET` settings that accidentally activate
   `EnvironmentCredential`.
4. In Azure, confirm managed/workload identity is enabled and that
   `AZURE_CLIENT_ID` identifies the intended user-assigned identity.
5. Distinguish authentication from authorization. A `401` generally indicates
   token or tenant problems; a `403` generally means the identity authenticated
   but lacks a Storage data-plane role. New role assignments can take time to
   propagate.
6. Confirm `AZURE_STORAGE_ACCOUNT_URL` has the expected account and cloud
   suffix, and verify proxy/firewall access to Microsoft Entra and Storage
   endpoints.

`app.py` surfaces `CredentialUnavailableError`, `ClientAuthenticationError`,
and Storage `HttpResponseError` separately instead of silently falling back.

## References

- [Credential chains in Azure Identity for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains)
- [`DefaultAzureCredential` API reference](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
- [Azure SDK for Python authentication overview](https://learn.microsoft.com/azure/developer/python/sdk/authentication/overview)
- [Azure SDK for Python logging](https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-logging)
- [Blob Storage Python client library](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)
