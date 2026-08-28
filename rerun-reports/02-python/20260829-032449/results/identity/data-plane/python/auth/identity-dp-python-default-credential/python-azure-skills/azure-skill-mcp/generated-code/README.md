# Authenticate an Azure SDK client with `DefaultAzureCredential`

This Python example creates an Azure Blob Storage `BlobServiceClient` with
passwordless Microsoft Entra authentication. It does not contain a client
secret, storage key, or connection string.

## 1. Install the pip packages

Create and activate a virtual environment, then install:

```powershell
py -m venv .venv
.\.venv\Scripts\Activate.ps1
py -m pip install -r requirements.txt
```

The required packages are:

- `azure-identity`: provides `DefaultAzureCredential`.
- `azure-storage-blob`: provides `BlobServiceClient`.

For Visual Studio Code single sign-on and broker authentication, install the
optional broker package instead:

```powershell
py -m pip install -r requirements-vscode.txt
```

`requirements-vscode.txt` includes both required packages plus
`azure-identity-broker`.

## 2. Create and use `DefaultAzureCredential`

`app.py` performs these steps:

1. Creates one `DefaultAzureCredential` instance.
2. Passes it to `BlobServiceClient` through the `credential` argument.
3. Reuses that client for service operations.
4. Closes both the client and credential when finished.

Authentication is lazy. Constructing the credential and client does not request
a token. The first Azure SDK operation requests a token for the service scope,
and the credential caches tokens for reuse.

Set the account endpoint in the current PowerShell session:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://your-storage-account.blob.core.windows.net"
```

Create the objects without making a network request:

```powershell
py app.py
```

After signing in and receiving the appropriate Azure RBAC role, trigger a
read-only operation:

```powershell
py app.py --list-containers
```

Listing containers normally requires a Blob Storage data-plane role such as
**Storage Blob Data Reader** at the narrowest practical scope. A management role
such as Contributor does not automatically grant access to blob data.

## 3. Default credential chain order

With current `azure-identity` versions, `DefaultAzureCredential` tries the
following credentials in order. It stops after one successfully gets a token.

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service-principal settings such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or certificate settings |
| 2 | `WorkloadIdentityCredential` | Federated workload identity configuration, commonly in Kubernetes |
| 3 | `ManagedIdentityCredential` | A system- or user-assigned managed identity on an Azure host |
| 4 | `SharedTokenCacheCredential` | A cached Visual Studio sign-in on Windows |
| 5 | `VisualStudioCodeCredential` | VS Code Azure sign-in; requires `azure-identity-broker` |
| 6 | `AzureCliCredential` | The account signed in with `az login` |
| 7 | `AzurePowerShellCredential` | The account signed in with `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | The account signed in with `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser sign-in; disabled by default |
| 10 | Broker credential | The OS broker account; requires `azure-identity-broker` |

The exact chain can change between `azure-identity` releases. Developer-tool
credentials continue to the next developer credential when token acquisition
fails. Deployed-service credentials are stricter: once one is configured and
attempts authentication, its authentication failure is surfaced instead of
silently falling through to a developer identity.

To deliberately limit local runs to developer-tool credentials, PowerShell can
set:

```powershell
$env:AZURE_TOKEN_CREDENTIALS = "dev"
```

For deterministic behavior, the constructor can enforce that setting with
`DefaultAzureCredential(require_envvar=True)`. This feature requires
`azure-identity` 1.23.0 or newer. Version 1.24.0 or newer can also set
`AZURE_TOKEN_CREDENTIALS` to one credential name, such as
`AzureCliCredential`.

## 4. Local development

Choose one supported developer sign-in:

- **Azure CLI:** run `az login`.
- **Azure Developer CLI:** run `azd auth login`.
- **Azure PowerShell:** run `Connect-AzAccount`.
- **VS Code:** install the Azure Resources extension, run **Azure: Sign In**,
  and install `azure-identity-broker`.

The app then runs as that developer account. The account must have the required
data-plane RBAC role on the target storage account or container. For teams,
assign least-privilege roles through a Microsoft Entra group rather than
granting broad permissions separately to every developer.

Environment credential variables are checked before local tool credentials.
Stale or partially configured `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, or
`AZURE_CLIENT_SECRET` values can therefore cause surprising behavior. Remove
unintended values or constrain the chain with `AZURE_TOKEN_CREDENTIALS`.

For Azure-hosted production workloads, prefer a specific
`ManagedIdentityCredential` instead of the broad default chain. This makes the
production identity deterministic and reduces fallback latency.

## 5. Troubleshoot authentication failures

Run the sample with identity logging:

```powershell
py app.py --list-containers --debug-auth
```

The `azure.identity` DEBUG log reports each unavailable credential and identifies
the credential that acquired a token. Treat DEBUG logs as sensitive because
they can contain tenant IDs, client IDs, object IDs, and account metadata.

Use the failure type to narrow the problem:

| Symptom | Likely cause |
|---|---|
| `ClientAuthenticationError` | No credential worked, a sign-in expired, the wrong tenant was used, or an earlier configured deployed credential failed |
| HTTP 403 | A token was acquired, but that identity lacks the required data-plane RBAC role |
| HTTP 404 or DNS failure | `AZURE_STORAGE_ACCOUNT_URL` is incorrect |
| Unexpected identity | An earlier chain entry, often environment variables or cached sign-in state, succeeded first |

Do not enable HTTP body logging routinely. Azure SDK HTTP DEBUG logging can
expose sensitive headers and request data.

## References

- [Credential chains in the Azure Identity library for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains)
- [Authenticate Python apps during local development](https://learn.microsoft.com/azure/developer/python/sdk/authentication/local-development-dev-accounts)
- [Configure logging in the Azure libraries for Python](https://learn.microsoft.com/azure/developer/python/sdk/azure-sdk-logging)
- [Azure Blob Storage client library for Python](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)
