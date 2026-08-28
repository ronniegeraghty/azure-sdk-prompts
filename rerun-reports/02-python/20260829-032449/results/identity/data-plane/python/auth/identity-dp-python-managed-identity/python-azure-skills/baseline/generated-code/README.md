# Azure Managed Identity with Python

This runnable project authenticates an Azure Blob Storage SDK client without
passwords, keys, or connection strings. It only accesses an existing account;
it does not create or modify Azure resources.

## Identity types

| Type | Lifecycle | Azure host relationship | Credential construction |
|---|---|---|---|
| System-assigned | Created and deleted with the host resource | Exactly one identity belongs to that host | `ManagedIdentityCredential()` |
| User-assigned | Independent Azure resource with its own lifecycle | One identity can be attached to multiple hosts; a host can have several | `ManagedIdentityCredential(client_id="...")` |

The `client_id` for a user-assigned identity disambiguates which attached
identity to use. It is the application's client ID, not the identity's object
(principal) ID or full Azure resource ID. Both identity types need an Azure
RBAC data-plane role appropriate for the operation; being attached to a host
does not grant access by itself.

## Install

Python 3.9 or newer is required.

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install -e ".[dev]"
```

Set the URL of an existing Blob Storage account:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://your-account.blob.core.windows.net"
```

## Run on Azure

Enable the identity on the Azure compute host and grant it a suitable role,
such as **Storage Blob Data Reader**, scoped as narrowly as practical. Role
assignments can take several minutes to propagate.

System-assigned identity:

```powershell
.\.venv\Scripts\python -m managed_identity_demo.cli --identity system
```

User-assigned identity:

```powershell
$env:AZURE_CLIENT_ID = "00000000-0000-0000-0000-000000000000"
.\.venv\Scripts\python -m managed_identity_demo.cli --identity user
```

`managed_identity_demo/auth.py` contains the credential factories.
`managed_identity_demo/storage.py` demonstrates passing the resulting
`TokenCredential` directly to `BlobServiceClient`. The same pattern works with
other modern Azure SDK clients that accept a `credential` argument.

## Local development

The managed identity endpoint exists only in supported Azure hosting
environments, so `ManagedIdentityCredential` normally cannot authenticate on a
developer workstation. Sign in using a developer credential:

```powershell
az login
$env:AZURE_ALLOW_LOCAL_CREDENTIALS = "1"
.\.venv\Scripts\python -m managed_identity_demo.cli --identity local
```

Local mode uses `DefaultAzureCredential` but explicitly excludes managed
identity, environment secrets, the shared token cache, and interactive browser
login. It can use supported developer tools such as Azure CLI, Azure Developer
CLI, or an IDE credential. The opt-in prevents a deployed application from
silently falling back to a developer identity. For automated local tests, mock
the `TokenCredential` as this project's tests do; do not store a client secret
in source control.

Production code should select `system` or `user` explicitly rather than put
`DefaultAzureCredential` in the deployed authentication path. This produces
predictable failures if managed identity is misconfigured.

## Errors and troubleshooting

The CLI returns distinct nonzero exit codes and writes details to stderr:

| Exit | Meaning | Checks |
|---|---|---|
| 2 | Invalid or missing configuration | Account URL, identity type, and user-assigned client ID |
| 3 | Authentication failure | Identity enabled and attached to this host; correct client ID; hosting service supports managed identity |
| 4 | Azure authorization or request rejection | Correct data-plane RBAC role and scope; allow time for role propagation |
| 5 | Network failure | DNS, proxy, firewall, TLS, private endpoint, and outbound access to identity/Azure endpoints |

Additional diagnostics:

1. Verify the user-assigned value is the **client ID**, especially when multiple
   identities are attached.
2. Separate authentication from authorization: obtaining a token can succeed
   while Blob Storage returns HTTP 403 because an RBAC role is missing.
3. Enable Azure SDK logging only while diagnosing because logs may contain
   identifiers and request metadata:

   ```powershell
   $env:AZURE_LOG_LEVEL = "info"
   ```

4. Do not retry configuration or authentication errors indefinitely. Azure SDK
   clients already retry appropriate transient HTTP failures.
5. Run offline tests with:

   ```powershell
   .\.venv\Scripts\python -m pytest
   ```
