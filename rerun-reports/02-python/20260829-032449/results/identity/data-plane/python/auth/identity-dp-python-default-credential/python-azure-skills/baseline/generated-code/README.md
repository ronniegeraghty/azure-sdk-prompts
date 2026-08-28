# Authenticate an Azure SDK client with `DefaultAzureCredential`

This example constructs an Azure Blob Storage client without storing a password,
secret, or access token in code. Client construction is offline. Azure Identity
requests a token only when an SDK operation, such as `list_containers()`, is
called.

## 1. Install the packages

Create and activate a virtual environment, then install the dependencies:

```text
python -m venv .venv
.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

| Package | Purpose |
|---|---|
| `azure-identity` | Provides `DefaultAzureCredential`. |
| `azure-storage-blob` | Provides the example `BlobServiceClient`; replace it with the package for the Azure service being used. |
| `azure-identity-broker` | Adds brokered developer authentication, including the current VS Code sign-in experience. |

## 2. Create and use the credential

`authenticate.py` creates one `DefaultAzureCredential` and passes it to
`BlobServiceClient`. Set the service endpoint and run it:

```text
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
python authenticate.py
```

The credential and client should be long-lived and reused instead of being
created for every request. Both expose `close()` and are closed in the example.
Creating the client does not prove that an identity has access. The first SDK
operation requests a token and the Azure service then checks that identity's
RBAC permissions.

For another Azure service, keep the same credential and pass it to that
service's client:

```text
credential = DefaultAzureCredential()
client = SomeAzureServiceClient(endpoint=endpoint, credential=credential)
```

Do not put client secrets in source control. If environment authentication is
required, inject `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
`AZURE_CLIENT_SECRET` through a secure local or deployment configuration.

## 3. Default credential chain

With all optional components available, `DefaultAzureCredential` tries these
credentials in order and uses the first one that can obtain a token:

1. `EnvironmentCredential` - service-principal or certificate settings in
   environment variables.
2. `WorkloadIdentityCredential` - federated workload identity configuration,
   commonly used by Azure Kubernetes Service.
3. `ManagedIdentityCredential` - a system-assigned or user-assigned managed
   identity exposed by the Azure host.
4. `SharedTokenCacheCredential` - cached user tokens where the platform and
   installed library version support it.
5. `VisualStudioCodeCredential` - the account signed in through VS Code's Azure
   tooling, where supported.
6. `AzureCliCredential` - the account selected by `az login`.
7. `AzurePowerShellCredential` - the account selected by `Connect-AzAccount`.
8. `AzureDeveloperCliCredential` - the account selected by `azd auth login`.
9. `BrokerCredential` - brokered sign-in supplied by
   `azure-identity-broker`, when supported.

The exact included credentials can vary by operating system, installed optional
packages, and `azure-identity` version. `InteractiveBrowserCredential` is
excluded by default; enable it explicitly only when interactive fallback is
appropriate. Constructor options such as `exclude_cli_credential=True` can
remove credentials from the chain.

Since `azure-identity` 1.14, failures from developer-tool credentials allow the
chain to continue to the next developer credential. A deployed credential that
is present but fails to authenticate raises an error instead of silently moving
on, which prevents a deployment misconfiguration from being masked.

## 4. Local development and Azure deployments

### Local development

- **Azure CLI:** run `az login`, select the intended subscription or tenant,
  and start the Python program. `AzureCliCredential` reuses that login. The CLI
  must be installed and available on `PATH`.
- **VS Code:** install the Azure Resources extension, sign in to the intended
  Azure account, and install `azure-identity-broker` as included here. Current
  Azure Identity versions use broker support for the VS Code account.
- **Environment variables:** use a development service principal only when a
  user login is unsuitable. Keep its secret outside `.env` files that might be
  committed.

The signed-in identity still needs the appropriate data-plane or management
role on the target resource. Being able to sign in does not grant resource
access.

### Azure deployments

Prefer a passwordless deployment identity:

- Enable a system-assigned or user-assigned managed identity on App Service,
  Functions, Virtual Machines, Container Apps, or another supported host.
- Grant that identity the least-privileged RBAC role required by the service.
- For a user-assigned managed identity, pass its client ID with
  `DefaultAzureCredential(managed_identity_client_id=...)` or set
  `AZURE_CLIENT_ID`.
- On AKS, configure Microsoft Entra Workload ID; the workload identity portion
  of the chain consumes the injected federated-token settings.

The application code does not change between local and deployed environments.
Locally, a developer credential normally wins. In Azure, workload identity or
managed identity normally wins, before the chain reaches developer tools.

## 5. Troubleshoot authentication

Run the example with Azure Identity debug logging:

```text
python authenticate.py --debug-auth
```

Because this example makes no SDK request, it only configures logging. Keep the
same logging setup while running the real SDK operation that fails; the log
then identifies every attempted credential and the reason it was unavailable
or failed. Avoid enabling HTTP body logging in production because request
details can contain sensitive data.

Check failures in this order:

1. Read the final `ClientAuthenticationError` and the preceding
   `azure.identity` entries to identify the credential that was selected.
2. Confirm the intended local login with `az account show`, or confirm that the
   expected managed/workload identity is enabled in Azure.
3. Check tenant selection. For CLI development, use `az login --tenant
   <tenant-id>` when the account belongs to multiple tenants.
4. Verify environment variable names and make sure a partially configured
   `EnvironmentCredential` is not taking precedence.
5. Distinguish authentication from authorization. HTTP 401 usually indicates a
   token, audience, or tenant problem; HTTP 403 usually means the identity
   authenticated but lacks the required RBAC role or access-policy permission.
6. Allow time for newly assigned RBAC roles to propagate, then retry.
7. Upgrade the identity packages together if logs show an unavailable broker
   or VS Code credential: `python -m pip install --upgrade azure-identity
   azure-identity-broker`.

For production troubleshooting, temporarily use `INFO` instead of `DEBUG` in
`configure_identity_logging()` when less detail is sufficient.
