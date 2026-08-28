# Python Managed Identity Azure SDK Example

This runnable example selects an Azure credential, creates Key Vault and Blob
Storage SDK clients, and handles common authentication and authorization
failures. Its default command is an offline dry run: credentials and clients are
constructed, but no token or Azure resource is requested.

## System-assigned and user-assigned identities

| | System-assigned | User-assigned |
|---|---|---|
| Lifecycle | Created and deleted with one Azure resource | Independent Azure resource |
| Attachment | Exactly one identity on its host | Can be attached to multiple hosts |
| Credential | `ManagedIdentityCredential()` | `ManagedIdentityCredential(client_id=...)` |
| Configuration | No identity selector | Identity **client ID** in `AZURE_CLIENT_ID` |
| Typical use | One workload with host-bound permissions | Shared or pre-authorized identity, stable across host replacement |

Both types need Azure RBAC or a service-specific access policy. Enabling an
identity authenticates the workload but does not grant it access to data.

## Setup and offline run

Python 3.10 or newer is recommended.

```text
python -m venv .venv
.venv\Scripts\activate
python -m pip install -r requirements.txt
python -m managed_identity_demo.main
```

The final command uses local mode by default and does not contact Azure.

## Azure-hosted examples

For a host with its system-assigned identity enabled:

```text
set APP_ENV=azure
set MANAGED_IDENTITY_TYPE=system
python -m managed_identity_demo.main
```

For a host with a user-assigned identity attached:

```text
set APP_ENV=azure
set MANAGED_IDENTITY_TYPE=user
set AZURE_CLIENT_ID=<managed-identity-client-id>
python -m managed_identity_demo.main
```

Use the user-assigned identity's **client ID**, not its object/principal ID or
Azure resource ID. These commands remain dry runs unless a network option is
added.

To request a token:

```text
python -m managed_identity_demo.main --check-auth
```

To demonstrate authenticated Azure SDK operations, set one or both endpoints
and request resource listing:

```text
set AZURE_KEY_VAULT_URL=https://<vault>.vault.azure.net
set AZURE_STORAGE_ACCOUNT_URL=https://<account>.blob.core.windows.net
python -m managed_identity_demo.main --list-resources
```

The identity needs suitable least-privilege roles, such as **Key Vault Secrets
User** to read secret metadata and **Storage Blob Data Reader** to list
containers. The code never reads secret values.

## Local development fallback

Set `APP_ENV=local` or leave it unset. The project then uses
`DefaultAzureCredential(exclude_managed_identity_credential=True)`, which can
use developer sign-ins from Azure CLI, Azure Developer CLI, Azure PowerShell,
or supported IDE tooling. Excluding managed identity prevents a slow or
misleading managed identity endpoint probe on a developer machine.

Use `ManagedIdentityCredential` directly in Azure rather than
`DefaultAzureCredential`. This keeps production authentication deterministic
and avoids accidentally selecting a developer or environment credential.
Never add client secrets as a managed identity fallback.

## Troubleshooting

| Symptom | Check |
|---|---|
| `CredentialUnavailableError` | Identity is enabled and attached to the Azure host; local developer tooling is signed in |
| `ClientAuthenticationError` | `AZURE_CLIENT_ID` is the user-assigned identity client ID; the selected identity is attached to the host |
| HTTP 401 | Token audience and service endpoint are correct |
| HTTP 403 | Authentication worked, but RBAC is missing, scoped incorrectly, or still propagating |
| Timeout reaching identity endpoint | The app is really running on a supported Azure host; proxies/firewalls are not intercepting the platform endpoint |
| Multiple user-assigned identities | Always provide `AZURE_CLIENT_ID` to select one unambiguously |

Add `--verbose` to enable Azure Identity diagnostic logging. Review logs before
sharing them because diagnostic output can contain tenant, endpoint, and
account metadata. Do not log access tokens.

Run the local unit tests with:

```text
python -m unittest discover -s tests -v
```
