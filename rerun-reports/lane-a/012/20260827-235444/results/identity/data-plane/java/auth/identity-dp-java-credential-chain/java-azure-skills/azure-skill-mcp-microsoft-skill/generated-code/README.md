# Environment-specific Azure credential chains for Java

This Java 17 sample chooses an explicit Azure Identity credential chain for local development, CI, or production, then requests an Azure Resource Manager token synchronously and asynchronously.

## Credential strategies

| Environment | Detection examples | Credential order |
|---|---|---|
| Development | No CI or Azure hosting markers | Azure CLI, Azure Developer CLI, Azure PowerShell, IntelliJ |
| CI | `TF_BUILD`, `PIPELINE_WORKSPACE`, `GITHUB_ACTIONS`, `CI` | Azure Pipelines workload identity service connection when fully configured, then `EnvironmentCredential` |
| Production | `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, `IMDS_ENDPOINT`, `AZURE_FEDERATED_TOKEN_FILE` | Managed identity, then AKS workload identity when its three standard variables are present |

For a user-assigned managed identity, set `AZURE_MANAGED_IDENTITY_CLIENT_ID`. If it is absent, the production chain uses the system-assigned identity.

The workload identity fallback is added when `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE` are all present, as they are in a correctly configured AKS workload identity pod.

Generic CI authentication uses the standard `EnvironmentCredential` variables:

```text
AZURE_TENANT_ID
AZURE_CLIENT_ID
AZURE_CLIENT_SECRET
```

An Azure Pipelines workload identity service connection additionally uses:

```text
AZURE_PIPELINES_SERVICE_CONNECTION_ID
SYSTEM_ACCESSTOKEN
SYSTEM_OIDCREQUESTURI
```

Map the pipeline's `System.AccessToken` into `SYSTEM_ACCESSTOKEN` and enable scripts to access the OAuth token. `SYSTEM_OIDCREQUESTURI` is supplied by Azure Pipelines for OIDC-enabled jobs.

## Build and run

```powershell
mvn test
mvn exec:java
```

CAE is enabled by default. Set `AZURE_ENABLE_CAE=false` to disable it. Azure Identity enables CAE on `TokenRequestContext`; developer credentials do not support CAE, and `AccessToken` is intentionally opaque, so the tester reports whether CAE was requested rather than claiming that the resource issued a CAE-capable token.

The sample only requests a token. It does not create, modify, or delete Azure resources.
