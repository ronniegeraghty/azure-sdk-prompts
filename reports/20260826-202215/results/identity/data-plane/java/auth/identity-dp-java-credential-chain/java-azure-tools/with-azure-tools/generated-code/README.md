# Azure credential chains for Java

This Java 17 sample selects a deliberately narrow Azure Identity credential chain for local development,
CI/CD, or production. It requests an Azure Resource Manager token synchronously and asynchronously without
creating or changing Azure resources.

## Run

```powershell
mvn test
mvn compile exec:java
```

CAE requests are enabled by default. Set `AZURE_ENABLE_CAE=false` to disable them. The output says
`CAE requested` because CAE is negotiated with the resource when the token is requested; an access token
does not expose a reliable client-side `isCaeEnabled` property. Applications that enable CAE must also
handle claims challenges from resource APIs.

## Environment-specific configuration

| Environment | Detection examples | Credential strategy |
| --- | --- | --- |
| Development | No CI or Azure-hosting signal | Azure CLI, Azure Developer CLI, Azure PowerShell, IntelliJ, then VS Code |
| CI/CD | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, `BUILD_BUILDID` | Environment-based service principal, then Azure Pipelines OIDC when fully configured |
| Production | `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, `WEBSITE_SITE_NAME`, `AZURE_FEDERATED_TOKEN_FILE`, Kubernetes | Managed identity first, then workload identity when fully configured |

CI service principals use `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and either `AZURE_CLIENT_SECRET` or
`AZURE_CLIENT_CERTIFICATE_PATH`. Azure Pipelines workload identity federation additionally uses
`AZURESUBSCRIPTION_SERVICE_CONNECTION_ID`, `SYSTEM_ACCESSTOKEN`, and `SYSTEM_OIDCREQUESTURI`.

Production uses the system-assigned managed identity unless `AZURE_MANAGED_IDENTITY_CLIENT_ID` names a
user-assigned identity. The Kubernetes fallback requires `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
`AZURE_FEDERATED_TOKEN_FILE`.

## References

- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)
- [AzurePipelinesCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.azurepipelinescredentialbuilder)
- [TokenRequestContext and CAE](https://learn.microsoft.com/java/api/com.azure.core.credential.tokenrequestcontext)
