# Authenticate an Azure SDK client with `DefaultAzureCredential`

This example creates an Azure Key Vault `SecretClient`. The same credential
object can be passed to most Azure SDK clients that accept a `credential`
argument.

## 1. Install the packages

Create a virtual environment and install the dependencies:

```powershell
py -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

| Package | Purpose |
|---|---|
| `azure-identity` | Provides `DefaultAzureCredential` and the other Azure credentials. |
| `azure-identity-broker` | Enables brokered sign-in, including the current VS Code authentication path. |
| `azure-keyvault-secrets` | Provides the example `SecretClient`; replace it with the package for the Azure service being used. |

## 2. Create and use the credential

`default_azure_credential_example.py` creates one `DefaultAzureCredential`
instance and passes it to `SecretClient`. The credential does not authenticate
when it is constructed. It obtains and caches a token when the client first
makes a request.

Set the resource URL and run the example:

```powershell
$env:AZURE_KEY_VAULT_URL = "https://<vault-name>.vault.azure.net"
$env:AZURE_KEY_VAULT_SECRET_NAME = "example-secret"
python .\default_azure_credential_example.py
```

The signed-in identity must have permission to read secrets from that vault.
Authentication proves identity; Azure RBAC or an access policy separately
controls authorization.

Reuse a credential instead of constructing one for every request. Azure SDK
clients and credentials are designed to be long-lived and handle token caching
and refresh.

## 3. Credential chain order

With the installed packages and default options, current `azure-identity`
versions attempt these credentials in order:

1. **EnvironmentCredential** - service-principal or workload identity values in
   environment variables such as `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
   `AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`.
2. **WorkloadIdentityCredential** - federated workload identity configured by
   `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
   `AZURE_FEDERATED_TOKEN_FILE`, commonly in Azure Kubernetes Service.
3. **ManagedIdentityCredential** - a system-assigned or user-assigned managed
   identity exposed by the Azure hosting environment.
4. **SharedTokenCacheCredential** - cached Microsoft identity tokens where the
   platform supports the shared cache (primarily Windows).
5. **VisualStudioCodeCredential** - the Azure account signed into VS Code,
   enabled through `azure-identity-broker`.
6. **AzureCliCredential** - the account selected by `az login`.
7. **AzurePowerShellCredential** - the account selected by
   `Connect-AzAccount`.
8. **AzureDeveloperCliCredential** - the account selected by `azd auth login`.
9. **BrokerCredential** - brokered authentication through the operating
   system's account broker when `azure-identity-broker` is installed.

Interactive browser authentication is excluded by default. It can be enabled
with `DefaultAzureCredential(exclude_interactive_browser_credential=False)`,
but explicit developer-tool sign-in is usually more predictable.

The exact chain can vary by `azure-identity` version, operating system,
installed optional packages, and `exclude_*` constructor options. The package
ranges in `requirements.txt` use the chain described above. In recent versions,
developer credentials later in the chain are still attempted after an earlier
developer credential cannot obtain a token. A deployed-service credential that
is available but fails token acquisition causes authentication to stop so that
deployment configuration errors are not hidden.

## 4. Local development and Azure deployments

For local development, sign in once with one of the supported tools:

```powershell
az login
az account set --subscription "<subscription-id-or-name>"
```

Alternatively, sign into the Azure account extension in VS Code. The broker
package in `requirements.txt` allows `VisualStudioCodeCredential` to use that
session. Azure CLI and VS Code credentials are intended for development, not
production deployment.

In Azure, enable a managed identity on the App Service, Function App, virtual
machine, Container App, or other supported host and grant that identity the
minimum required role. `ManagedIdentityCredential` is then selected without
storing a password. For a user-assigned managed identity, set
`AZURE_CLIENT_ID` to its client ID. In AKS, prefer Microsoft Entra Workload ID;
the injected federated-token settings select `WorkloadIdentityCredential`.

Do not put client secrets in source control. When managed identity or workload
identity is unavailable, inject service-principal settings through the
deployment platform's secret/configuration facility.

## 5. Troubleshoot failures with logging

The example enables `DEBUG` logs for `azure.identity`. These logs show each
credential attempted, why unavailable credentials were skipped, and which
credential succeeded. Azure SDK HTTP logging can expose request metadata, so
keep it off unless needed and never publish logs without reviewing them.

Useful checks:

```powershell
az account show
az account get-access-token --scope https://vault.azure.net/.default
python .\default_azure_credential_example.py
```

When diagnosing a failure:

1. Read the complete final `ClientAuthenticationError`; it aggregates messages
   from the attempted credentials.
2. Confirm the intended local tool is signed in and using the correct tenant.
3. Check environment variables for stale or partially configured
   `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, `AZURE_CLIENT_SECRET`,
   `AZURE_CLIENT_CERTIFICATE_PATH`, or `AZURE_FEDERATED_TOKEN_FILE` values.
4. In Azure, verify that managed identity is enabled and, for a user-assigned
   identity, that `AZURE_CLIENT_ID` identifies the assigned identity.
5. Distinguish authentication errors (`ClientAuthenticationError`, token
   acquisition) from authorization errors (usually HTTP 403). A 403 normally
   means authentication succeeded but the identity lacks an RBAC role or access
   policy. Allow time for newly assigned roles to propagate.
6. If several developer accounts are available, make selection deterministic
   with `DefaultAzureCredential(tenant_id="...")`, a tool-specific tenant
   option, or `exclude_*` options. Avoid broadly excluding deployed identity
   credentials merely to conceal a configuration problem.

For deeper HTTP diagnostics, add the following temporarily after reviewing the
risk of sensitive metadata in logs:

```python
logging.getLogger("azure.core.pipeline.policies.http_logging_policy").setLevel(
    logging.DEBUG
)
```
