# Python Managed Identity with Azure SDK clients

This runnable example authenticates `BlobServiceClient` with Microsoft Entra ID.
It uses no account keys or connection strings and creates no Azure resources.

## System-assigned and user-assigned identities

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created on and deleted with one Azure resource | Independent Azure resource; can be shared |
| Identity selection | Azure provides the resource's only system identity | The client must select one when multiple identities are available |
| Credential | `ManagedIdentityCredential()` | `ManagedIdentityCredential(client_id="...")` |
| Best fit | One workload, simple lifecycle | Shared identity, stable permissions, or identity reuse |

Enabling an identity does not grant data access. Assign the identity an appropriate
least-privilege role at the required scope. This sample needs a Blob data-plane role,
such as **Storage Blob Data Reader**, not merely a management-plane Reader role.

## Install

Python 3.9 or newer is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -e .
```

Set the endpoint to an existing storage account:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account>.blob.core.windows.net"
```

## Run in Azure

Use the system-assigned identity attached to the host:

```powershell
python -m managed_identity_demo --identity system
```

Use a particular user-assigned identity attached to the host. Use its **client ID**,
not its object/principal ID:

```powershell
$env:AZURE_CLIENT_ID = "<user-assigned-managed-identity-client-id>"
python -m managed_identity_demo --identity user
```

The credential is passed directly to `BlobServiceClient`; the Azure Identity library
acquires and refreshes access tokens automatically.

## Local development fallback

Managed Identity endpoints exist only on supported Azure hosts, so direct
`ManagedIdentityCredential` calls normally cannot authenticate on a developer machine.
The default mode uses `DefaultAzureCredential`, allowing the same command to use a
developer sign-in locally and managed identity after deployment:

```powershell
# Authenticate with one supported developer tool first, for example Azure CLI,
# Azure Developer CLI, VS Code, or Azure PowerShell.
python -m managed_identity_demo --identity default
```

When `AZURE_CLIENT_ID` is set, default mode selects that user-assigned identity in
Azure. Leave it unset for a system-assigned identity. In production, consider setting
`AZURE_TOKEN_CREDENTIALS=prod` to constrain newer `azure-identity` versions to
deployment-safe credentials. An explicit `ChainedTokenCredential` with
`ManagedIdentityCredential` followed by `AzureCliCredential` is another option, but
`DefaultAzureCredential` is generally simpler and avoids maintaining a custom chain.
Never put client secrets in source code.

## Troubleshooting

Run with `--debug` to enable Azure Identity and HTTP pipeline diagnostics. Review logs
before sharing them because request metadata can be sensitive.

| Symptom | Likely cause and action |
|---|---|
| `CredentialUnavailableError` | Managed identity is disabled/not attached, the code is running locally with `--identity system` or `user`, or the host does not support managed identity. Enable/attach it or use `--identity default` locally. |
| `ClientAuthenticationError` | Wrong user-assigned client ID, tenant mismatch, or identity endpoint failure. Confirm `AZURE_CLIENT_ID` is the managed identity's client ID. |
| HTTP 403 | Authentication worked but authorization failed. Add the required data-plane RBAC role and allow time for role propagation. |
| HTTP 404 or DNS/connectivity error | Check `AZURE_STORAGE_ACCOUNT_URL`, private endpoint DNS, firewall rules, and outbound network access. |
| Slow local authentication | Avoid direct managed identity mode locally. Use default mode and exclude unused credentials in application-specific configurations if needed. |

The CLI returns distinct exit codes: `2` for configuration, `3` for unavailable
credentials, `4` for authentication, `5` for Azure service/authorization errors, and
`6` for network failures.

## Test offline

Tests mock credential creation and Azure service calls; they do not contact Azure:

```powershell
python -m unittest discover -s tests -v
```
