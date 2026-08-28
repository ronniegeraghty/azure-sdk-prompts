# Authenticate an Azure SDK client with `DefaultAzureCredential`

This sample creates an authenticated Azure Blob Storage client without storing
keys or secrets in source code. It is intentionally offline-safe: client
construction and `get_container_client` do not contact Azure. An access token is
requested only when an SDK operation sends a service request.

## 1. Install the packages

Python 3.9 or later is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

| Package | Purpose |
|---|---|
| `azure-identity` | Provides `DefaultAzureCredential` and other Microsoft Entra credentials. |
| `azure-storage-blob` | Provides the example `BlobServiceClient`. Replace it with the package for the Azure service you use. |
| `azure-identity-broker` | Enables brokered authentication and current VS Code authentication support. It is optional if neither is needed. |

## 2. Create and use the credential

Run the offline sample:

```powershell
$env:AZURE_STORAGE_ACCOUNT = "your-storage-account"
python .\default_credential_example.py
```

`default_credential_example.py` creates one `DefaultAzureCredential`, passes it
to `BlobServiceClient`, and reuses it for the lifetime of the client. Both
objects are context managers so their transports are closed cleanly.

The credential does not authenticate in its constructor. The first Azure SDK
operation that needs authorization asks the credential for a token with the
service's Microsoft Entra scope. The SDK caches and refreshes tokens. In a real
application, an operation such as `list_containers()` would cause this network
authentication and service request; this sample intentionally does not issue
one.

The signed-in identity must also have an appropriate Azure RBAC role. Successful
authentication proves identity; it does not grant access. For listing blobs, for
example, a data-plane role such as **Storage Blob Data Reader** is required at
the appropriate scope.

## 3. Default credential chain

With current `azure-identity`, the default chain tries these credentials in
order and stops when one obtains a token:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service principal values such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or certificate settings. |
| 2 | `WorkloadIdentityCredential` | Federated token settings, commonly injected into an AKS pod. |
| 3 | `ManagedIdentityCredential` | A system-assigned or user-assigned identity exposed by the Azure host. |
| 4 | `SharedTokenCacheCredential` | Windows shared sign-in cache, commonly populated by Visual Studio. |
| 5 | `VisualStudioCodeCredential` | The account signed in through the VS Code Azure Resources extension; current support uses `azure-identity-broker`. |
| 6 | `AzureCliCredential` | The account selected by `az login`. |
| 7 | `AzurePowerShellCredential` | The account selected by `Connect-AzAccount`. |
| 8 | `AzureDeveloperCliCredential` | The account selected by `azd auth login`. |
| 9 | `BrokerCredential` | On Windows or WSL, the Web Account Manager account when `azure-identity-broker` is installed. |
| 10 | `InteractiveBrowserCredential` | Browser sign-in, **disabled by default**; enable with `exclude_interactive_browser_credential=False`. |

Exact contents can vary by `azure-identity` version, operating system,
installed optional packages, and constructor exclusions. Since version 1.14.0,
developer credentials continue to the next developer credential after an
authentication failure. Deployed-service credentials stop when they can attempt
authentication but fail, surfacing configuration errors instead of silently
falling through.

For predictable production behavior, either use the specific production
credential (often `ManagedIdentityCredential`) or constrain
`DefaultAzureCredential`. Current releases support
`AZURE_TOKEN_CREDENTIALS=prod` to retain only deployed-service credentials,
`AZURE_TOKEN_CREDENTIALS=dev` to retain developer credentials, or a supported
credential name to select one credential. Never hardcode tenant IDs, client
secrets, certificates, or access tokens.

## 4. Local development and Azure deployments

### Azure CLI

Authenticate locally with `az login`, select the intended subscription with
`az account set`, and run the sample. `DefaultAzureCredential` reaches
`AzureCliCredential` after earlier unavailable credentials. Azure CLI is a
developer dependency only; the application does not shell out to Azure CLI once
another earlier credential succeeds.

### VS Code

Install the VS Code **Azure Resources** extension, sign in to Azure from VS
Code, and install `azure-identity-broker` as included here. The chain can then
use `VisualStudioCodeCredential`. If several accounts or tenants are available,
make the tenant selection explicit rather than relying on an unintended cached
account.

### Azure-hosted applications

Enable a managed identity on App Service, Functions, a VM, or another supported
host, then assign that identity the least-privileged Azure RBAC role needed by
the service. The same application code uses `ManagedIdentityCredential` within
the chain; no client secret is deployed.

For a user-assigned managed identity, set `AZURE_CLIENT_ID` to its client ID or
pass `managed_identity_client_id` to `DefaultAzureCredential`. On AKS, prefer
Microsoft Entra Workload ID; the workload identity webhook supplies
`AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE`.

In production, set `AZURE_TOKEN_CREDENTIALS=prod` when supported by the installed
`azure-identity` version, or replace the chain with the one specific credential
expected by the host. This reduces latency and prevents accidental use of a
developer login.

## 5. Troubleshoot authentication failures

Enable identity-only debug logging for this sample:

```powershell
$env:AZURE_IDENTITY_DEBUG = "true"
python .\default_credential_example.py
```

The offline sample will not request a token, so it will not emit a credential
attempt trace. Enable the same logger in the real application and invoke the
failing SDK operation. The trace identifies each attempted credential and why it
was unavailable or failed. Do not leave debug logging enabled in production:
logs can contain tenant IDs, client IDs, object IDs, request URLs, and other
sensitive metadata.

Use the failure text to check:

1. **Wrong local account or tenant:** confirm the VS Code or CLI login and tenant.
2. **Incomplete environment credential:** either set all required service
   principal variables or remove stale partial variables.
3. **Managed identity unavailable:** verify the identity is enabled on the host;
   for a user-assigned identity, verify `AZURE_CLIENT_ID`.
4. **Workload identity misconfigured:** verify tenant ID, client ID, federated
   token file, service account, and federated identity subject/audience.
5. **Authentication succeeds but the service returns 403:** assign the identity
   the correct data-plane or management-plane RBAC role and allow time for role
   assignment propagation.
6. **Network or authority errors:** verify DNS, proxy/firewall rules, the Azure
   authority host, and access to Microsoft Entra endpoints.

`CredentialUnavailableError` means a credential could not be used in the
current environment. `ClientAuthenticationError` means authentication was
attempted but failed; its message normally contains the chain's diagnostics.

## References

- [DefaultAzureCredential API reference](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
- [Credential chains in Azure Identity for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains)
- [Authenticate Python apps to Azure services](https://learn.microsoft.com/azure/developer/python/sdk/authentication-overview)
- [Azure SDK for Python logging](https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-logging)
