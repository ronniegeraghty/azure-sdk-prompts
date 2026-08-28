# Authenticate an Azure SDK client with `DefaultAzureCredential`

This example creates an Azure Blob Storage client and supplies a
`DefaultAzureCredential`. The default run only constructs the client, so it works
offline. Azure SDK clients request a token lazily, when the first service operation
is made.

## 1. Install the pip packages

Create and activate a virtual environment, then install:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

- `azure-identity`: provides `DefaultAzureCredential`.
- `azure-storage-blob`: provides the example `BlobServiceClient`. Replace this
  package with the client library for the Azure service you use.
- `azure-identity-broker` (optional): enables VS Code authentication and
  brokered Windows/WSL sign-in. Install `requirements-broker.txt` instead of
  `requirements.txt` when those flows are needed.

## 2. Create and use the credential

`default_credential_example.py` creates one `DefaultAzureCredential` and passes it
to `BlobServiceClient`:

```python
credential = DefaultAzureCredential()
client = BlobServiceClient(account_url=account_url, credential=credential)
```

Run the offline construction example:

```powershell
python .\default_credential_example.py
```

For a real storage account, set its HTTPS endpoint. The identity selected by the
credential must also have an appropriate Azure RBAC data-plane role:

```powershell
$env:AZURE_STORAGE_ACCOUNT_URL = "https://<account-name>.blob.core.windows.net"
```

To perform the example's read-only service request, explicitly enable it:

```powershell
$env:AZURE_RUN_LIVE_REQUEST = "1"
python .\default_credential_example.py
```

The first service request causes the client to ask the credential for a token.
Reuse the credential and service client rather than constructing them for every
request. Close them during application shutdown, as the example does.

## 3. Default credential chain order

With `azure-identity` 1.23 or later, the default chain tries these credentials in
order and stops when one gets a token:

| Order | Credential | Source |
|---:|---|---|
| 1 | `EnvironmentCredential` | Service principal configured through environment variables |
| 2 | `WorkloadIdentityCredential` | Federated workload identity configuration, commonly in AKS |
| 3 | `ManagedIdentityCredential` | System-assigned or user-assigned Azure managed identity |
| 4 | `SharedTokenCacheCredential` | Windows shared cache, such as a Visual Studio sign-in |
| 5 | `VisualStudioCodeCredential` | VS Code Azure Resources extension; requires `azure-identity-broker` |
| 6 | `AzureCliCredential` | Account signed in with `az login` |
| 7 | `AzurePowerShellCredential` | Account signed in with `Connect-AzAccount` |
| 8 | `AzureDeveloperCliCredential` | Account signed in with `azd auth login` |
| 9 | `InteractiveBrowserCredential` | Browser sign-in; **disabled by default** |
| 10 | Broker credential | Windows/WSL Web Account Manager; requires `azure-identity-broker` |

Unavailable credentials are skipped. Interactive browser authentication can be
enabled with
`DefaultAzureCredential(exclude_interactive_browser_credential=False)`, but it is
usually better for local development to sign in through an approved developer
tool.

The chain can be narrowed with constructor exclusion options. In
`azure-identity` 1.23 or later, `AZURE_TOKEN_CREDENTIALS=dev` keeps developer
credentials and `AZURE_TOKEN_CREDENTIALS=prod` keeps deployed-service
credentials. Narrowing the chain makes the selected identity more predictable.

## 4. Local development and Azure deployments

### Azure CLI

Sign in with `az login`. If multiple subscriptions or tenants are available,
select the intended subscription with `az account set`. The Python process then
uses the CLI account when earlier credentials in the chain are unavailable.

### VS Code

Install the Azure Resources extension, sign in to Azure from VS Code, and install
the optional broker requirements:

```powershell
python -m pip install -r requirements-broker.txt
```

`VisualStudioCodeCredential` is before Azure CLI in the chain. If both are signed
in as different users, VS Code can therefore win. Logging shows which credential
actually supplied the token.

### Azure-hosted applications

Enable a managed identity on App Service, Functions, Container Apps, a VM, or
another supported host, and grant that identity the least-privilege Azure RBAC
role on the target resource. Do not deploy developer credentials or client
secrets.

- A system-assigned identity requires no client ID configuration.
- For a user-assigned identity, set `AZURE_CLIENT_ID` to that identity's client
  ID, or pass
  `managed_identity_client_id="<client-id>"` to `DefaultAzureCredential`.
- AKS and other federated environments can use workload identity, which appears
  before managed identity in the chain.

The same `DefaultAzureCredential()` code can move between local development and
Azure. For stricter production behavior, use `AZURE_TOKEN_CREDENTIALS=prod`, or
replace the chain with the specific `ManagedIdentityCredential` or
`WorkloadIdentityCredential` expected by the deployment.

## 5. Troubleshoot authentication with logging

Enable identity-chain logging without changing the code:

```powershell
$env:AZURE_IDENTITY_LOG_LEVEL = "INFO"
python .\default_credential_example.py
```

Use `DEBUG` only for short diagnostic sessions. Debug-level HTTP logging in Azure
SDKs can expose sensitive request details, so do not enable it permanently or
publish raw logs.

Check the following:

1. Read the complete `DefaultAzureCredential` error. It lists every attempted
   credential and why each failed. `CredentialUnavailableError` means a method
   was not configured for this environment; `ClientAuthenticationError` means an
   available method tried and failed to authenticate.
2. Find the successful credential in the `azure.identity` logs. If it is the
   wrong developer account, sign out of that tool or exclude that credential.
3. For CLI authentication, confirm `az account show` reports the expected tenant
   and subscription, and renew an expired session with `az login`.
4. For environment authentication, verify `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,
   and `AZURE_CLIENT_SECRET` are all present and that the secret value, not its
   identifier, was supplied. Prefer managed identity over a secret in Azure.
5. For managed identity, verify the identity is enabled on the host. For a
   user-assigned identity, verify the client ID is correct.
6. Distinguish authentication from authorization: HTTP 401 usually indicates an
   invalid or unsuitable token; HTTP 403 usually means authentication succeeded
   but the identity lacks the required RBAC role. RBAC changes can take time to
   propagate.
7. Check Microsoft Entra sign-in logs and the error's `AADSTS` code for tenant,
   Conditional Access, consent, or credential failures.

## References

- [DefaultAzureCredential API reference](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)
- [Credential chains for the Azure Identity library for Python](https://learn.microsoft.com/azure/developer/python/sdk/authentication/credential-chains)
- [Authenticate Python apps during local development](https://learn.microsoft.com/azure/developer/python/sdk/authentication/local-development-dev-accounts)
- [Azure Identity troubleshooting guide](https://github.com/Azure/azure-sdk-for-python/blob/main/sdk/identity/azure-identity/TROUBLESHOOTING.md)
