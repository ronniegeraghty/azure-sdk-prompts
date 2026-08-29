# Azure credential chains for Java

Small Java 17 sample that selects a deliberately narrow Azure Identity credential chain for local
development, CI/CD, or production. It requests an Azure Resource Manager token synchronously and
asynchronously without creating or modifying Azure resources.

## Credential strategies

| Environment | Detection examples | Credential order |
|---|---|---|
| Development | No CI or Azure-hosting markers | Azure CLI, Azure Developer CLI, IntelliJ |
| CI/CD | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, `PIPELINE_WORKSPACE` | Azure Pipelines workload federation when configured, then `EnvironmentCredential` |
| Production | `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, `AZURE_FEDERATED_TOKEN_FILE` | Managed identity, then workload identity |

Set `APP_ENVIRONMENT=dev`, `ci`, or `production` to override detection. For a user-assigned managed
identity, set `AZURE_MANAGED_IDENTITY_CLIENT_ID`; otherwise the production chain uses the
system-assigned identity.

Generic CI uses the standard `EnvironmentCredential` variables such as `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`. Azure Pipelines
workload federation uses:

- `AZURE_TENANT_ID`
- `AZURE_CLIENT_ID`
- `AZURE_SERVICE_CONNECTION_ID`
- `SYSTEM_ACCESSTOKEN`
- `SYSTEM_OIDCREQUESTURI`

CAE is requested by default through `TokenRequestContext.setCaeEnabled(true)`. Set
`AZURE_CAE_ENABLED=false` to disable the request. The output says **CAE requested** because
`AccessToken` does not expose a definitive CAE-capable flag; the target resource decides whether it
honors the request.

## Build and run

```text
mvn test
mvn exec:java
```

The run command performs real token acquisition. Sign in with one of the configured developer
tools or provide the environment variables appropriate for the detected environment.

References:

- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)
- [AzurePipelinesCredentialBuilder](https://learn.microsoft.com/java/api/com.azure.identity.azurepipelinescredentialbuilder)
- [TokenRequestContext and CAE](https://learn.microsoft.com/java/api/com.azure.core.credential.tokenrequestcontext)
