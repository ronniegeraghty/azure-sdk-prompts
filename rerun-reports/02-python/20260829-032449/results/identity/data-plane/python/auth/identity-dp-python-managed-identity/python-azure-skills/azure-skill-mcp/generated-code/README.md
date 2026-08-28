# Azure Managed Identity with Python

This project shows how to authenticate an Azure SDK client without secrets. It
uses `ManagedIdentityCredential` on Azure and developer credentials only during
local development. The sample client lists Blob Storage containers, an operation
that needs a data-plane role such as **Storage Blob Data Reader**.

The CLI is safe to run offline by default: it constructs the credential and SDK
client but makes no request unless `--execute` is supplied.

## System-assigned and user-assigned identities

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created on, and deleted with, one Azure host resource | Independent Azure resource |
| Sharing | Used by its single host | Can be attached to multiple hosts |
| Selection in code | No identity selector | Pass its **client ID** |
| Best fit | Simple one-host ownership and cleanup | Shared permissions, stable identity across host replacement |

Both types obtain tokens from the Azure managed identity endpoint; neither
stores a secret in application configuration. The identity must be attached to
the Azure host, and its service principal must have an appropriate Azure RBAC
role. Prefer resource or container scope over resource-group or subscription
scope.

## Install

Python 3.10 or newer is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -e .
```

Copy `.env.example` values into your process environment. This project does not
load `.env` automatically, which avoids accidentally treating a checked-in file
as a credential source.

## Run the examples

Construct a system-assigned credential and Blob client without network access:

```powershell
managed-identity-demo --auth system
```

Use a system-assigned identity on an Azure host:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
managed-identity-demo --auth system --execute
```

Use a user-assigned identity attached to the Azure host:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
$env:AZURE_MANAGED_IDENTITY_CLIENT_ID = "<managed-identity-client-id>"
managed-identity-demo --auth user --execute
```

`src/managed_identity_demo/auth.py` contains the essential distinction:
`ManagedIdentityCredential()` selects the system-assigned identity, while
`ManagedIdentityCredential(client_id=...)` selects a user-assigned identity.
The credential is passed directly to `BlobServiceClient`; the SDK acquires and
caches access tokens as needed.

## Local development fallback

Managed identity endpoints exist only on supported Azure hosts. Do not emulate a
managed identity or copy production secrets locally. Use one of these explicit
development modes:

```powershell
# Uses signed-in developer tools while skipping hosted credential probes.
managed-identity-demo --auth local-default --execute

# Deterministic fallback for a team standardized on Azure CLI authentication.
managed-identity-demo --auth local-cli --execute
```

Before running, sign in using Azure CLI, Azure Developer CLI, Azure PowerShell,
or the VS Code Azure extension as appropriate. The developer identity needs the
same data-plane access as the application. `local-default` deliberately excludes
environment, workload identity, and managed identity credentials so production
authentication cannot be selected accidentally. `local-cli` narrows the choice
to `AzureCliCredential`.

For a wider standard `DefaultAzureCredential` development chain, recent
`azure-identity` versions also support setting `AZURE_TOKEN_CREDENTIALS=dev` and
constructing `DefaultAzureCredential(require_envvar=True)`. A service principal
is suitable for automated local integration tests, but use certificate or
federated authentication and a narrowly scoped identity rather than a client
secret.

## Troubleshooting

Run with `--verbose` to see credential-selection diagnostics. The sample keeps
HTTP body logging disabled so tokens and service data are not written to logs.

| Symptom | Likely cause and fix |
|---|---|
| Credential unavailable | Managed identity is not enabled/attached, the code is not running on a supported Azure host, or the selected local tool is not signed in |
| Authentication failed | For user-assigned identity, verify the **client ID** (not object/principal ID), attachment to the host, and tenant |
| HTTP 403 | Authentication worked but authorization failed; add the minimum Blob **data-plane** RBAC role and allow time for propagation |
| HTTP 404 or wrong account | Check `AZURE_STORAGE_ACCOUNT_URL`, cloud suffix, container/resource name, and private endpoint DNS |
| Timeout or connection error | Check proxy, firewall, DNS, private endpoint routing, and whether the Azure host can reach its managed identity endpoint |
| Works locally, fails on Azure | Local developer and managed identity are different principals; assign the Azure role to the managed identity |
| Multiple user-assigned identities | Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` explicitly so credential selection is deterministic |

The process exits with code 2 for configuration errors, 3 for credential or
authentication errors, 4 for Storage HTTP errors, and 5 for network failures.

## References

- [System-assigned managed identity authentication for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/system-assigned-managed-identity)
- [User-assigned managed identity authentication for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/user-assigned-managed-identity)
- [Local development authentication for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/local-development-dev-accounts)
- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)
- [Azure Blob Storage client library for Python](https://learn.microsoft.com/azure/storage/blobs/storage-quickstart-blobs-python)
