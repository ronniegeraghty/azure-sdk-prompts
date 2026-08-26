# Azure credential chains for Java

This Java 17 sample selects an Azure Identity credential chain for local
development, CI/CD, or production, then requests an Azure Resource Manager token
with both synchronous and asynchronous SDK APIs. It does not create or modify
Azure resources.

## Credential strategies

| Environment | Detection signals | Credential order |
| --- | --- | --- |
| Development | No CI or Azure-hosting signals | Azure CLI, Azure Developer CLI, Visual Studio Code, IntelliJ, Azure PowerShell |
| CI/CD | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, pipeline workspace variables, and similar | `EnvironmentCredential`, then Azure Pipelines workload-identity service connection when configured |
| Production | App Service/Functions managed-identity variables, Kubernetes workload-identity variables, or reachable IMDS | Managed identity, then Kubernetes workload identity when configured |

Production uses a system-assigned managed identity by default. Set
`AZURE_CLIENT_ID` to select a user-assigned managed identity.

For an Azure Pipelines workload-identity service connection, expose:

- `AZURE_TENANT_ID`
- `AZURE_CLIENT_ID`
- `AZURE_SERVICE_CONNECTION_ID`
- `SYSTEM_ACCESSTOKEN`
- `SYSTEM_OIDCREQUESTURI` (provided by Azure Pipelines)

For Kubernetes workload identity, expose `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`,
and `AZURE_FEDERATED_TOKEN_FILE`.

CAE is enabled by default. Set `AZURE_ENABLE_CAE=false` to disable it. CAE is
requested through `TokenRequestContext`; the resource and tenant must also
support CAE.

## Run

```shell
mvn test
mvn compile exec:java
```

The second command performs real authentication against Microsoft Entra ID but
does not call Azure Resource Manager or change any Azure resources.
