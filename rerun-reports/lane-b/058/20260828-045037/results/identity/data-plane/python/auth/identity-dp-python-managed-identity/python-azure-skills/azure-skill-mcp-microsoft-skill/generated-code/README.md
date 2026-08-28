# Azure Managed Identity with Python

This runnable project shows how to authenticate `BlobServiceClient` with
Microsoft Entra tokens. It does not use storage keys, connection strings, or
embedded secrets.

## System-assigned and user-assigned identities

| Characteristic | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created and deleted with one Azure resource | Independent Azure resource |
| Sharing | Belongs to one host resource | Can be attached to multiple hosts |
| Selection | `ManagedIdentityCredential()` | `ManagedIdentityCredential(client_id=...)` |
| Best fit | One workload with matching lifecycle | Shared permissions, stable identity, or multiple identities on one host |

Both identity types must be enabled or attached to the Azure compute resource,
and both require an Azure RBAC data-plane role on the target storage resource.
For this read-only example, use a role such as **Storage Blob Data Reader**.
Role assignments can take several minutes to propagate.

## Install and run

Python 3.9 or newer is required.

```text
python -m venv .venv
.venv\Scripts\activate
python -m pip install -e .
```

Set an account URL. PowerShell example:

```text
$env:AZURE_STORAGE_ACCOUNT_URL = "https://your-account.blob.core.windows.net"
```

The default `inspect` command constructs and closes the selected credential and
SDK client without requesting a token or contacting Azure:

```text
managed-identity-demo inspect --mode system
managed-identity-demo inspect --mode user --client-id <managed-identity-client-id>
```

After the identity is attached and authorized, make a real SDK request:

```text
managed-identity-demo list-containers --mode system
managed-identity-demo list-containers --mode user --client-id <managed-identity-client-id>
```

The user-assigned selector is the managed identity's **client ID**, not its
object/principal ID. It can instead be supplied through
`AZURE_MANAGED_IDENTITY_CLIENT_ID`.

## Local development fallback

Managed identity endpoints exist only on supported Azure hosts, so strict
`system` and `user` modes normally fail on a developer machine. Use one of
these modes locally:

| Mode | Behavior |
|---|---|
| `local` | `DefaultAzureCredential` with managed identity disabled; uses developer credentials such as Azure CLI, VS Code, Azure PowerShell, or Azure Developer CLI |
| `auto-system` | One code path for local developer credentials and system-assigned identity on Azure |
| `auto-user` | One code path for local developer credentials and the selected user-assigned identity on Azure |

Examples:

```text
managed-identity-demo list-containers --mode local
managed-identity-demo list-containers --mode auto-system
managed-identity-demo list-containers --mode auto-user --client-id <client-id>
```

Interactive browser authentication is disabled by default. Add
`--allow-interactive-browser` only for local interactive use. In production,
prefer strict `system` or `user` mode. If using an auto mode, set
`AZURE_TOKEN_CREDENTIALS=prod` where supported to constrain
`DefaultAzureCredential` to deployment-safe credentials.

## Using the credential with other Azure SDK clients

Azure SDK clients that accept a `TokenCredential` use the same pattern as
`BlobServiceClient`: create one credential, pass it as the client's
`credential` argument, reuse it across clients, then close clients and the
credential. The sample implements this in `credentials.py` and `storage.py`.
The identity needs the service-specific RBAC data-plane role; a management-plane
role does not automatically grant access to blob data.

## Troubleshooting

Run with `--debug` to enable Azure Identity diagnostics. Review logs before
sharing them because they can contain tenant, endpoint, and account metadata.

| Symptom | Likely cause and action |
|---|---|
| `CredentialUnavailableError` | Managed identity is not enabled/attached, the code is not on a supported Azure host, or no local developer credential is signed in |
| `ClientAuthenticationError` | Wrong user-assigned client ID, identity endpoint failure, tenant mismatch, or token acquisition failure |
| HTTP 403 | Authentication worked, but the identity lacks a Blob data-plane role or RBAC propagation is still in progress |
| Connection/timeout failure | Check account URL, DNS, proxy/firewall rules, private endpoints, and managed identity endpoint availability |
| Multiple user-assigned identities | Pass the intended identity's client ID explicitly; do not rely on implicit selection |

The CLI returns exit code `2` for configuration, authentication, authorization,
and connectivity failures. It preserves exception causes in the Python API
while printing concise remediation guidance from the command line.

## Tests

Tests are local and make no Azure or managed identity endpoint calls:

```text
python -m unittest discover -s tests -v
```
