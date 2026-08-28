# Azure credential chains for Java

This Java 17 sample chooses an explicit Azure Identity credential chain for local development, CI, or production.
It only requests an Azure Resource Manager token; it does not create or modify Azure resources.

| Environment | Detection examples | Credential strategy |
|---|---|---|
| Development | No CI or Azure hosting marker | Azure CLI, Azure Developer CLI, Azure PowerShell, then IntelliJ |
| CI | `CI`, `TF_BUILD`, `GITHUB_ACTIONS`, `BUILD_BUILDID` | Azure Pipelines service connection when configured, then `EnvironmentCredential` |
| Production | Managed identity endpoint, App Service, Container Apps, Kubernetes, or federated token marker | System/user-assigned managed identity, then configured Kubernetes workload identity |

## Run

```shell
mvn clean test
mvn exec:java
mvn exec:java -Dexec.args="--cae"
```

CAE can also be enabled with `AZURE_ENABLE_CAE=true`. Azure Identity applies CAE through
`TokenRequestContext.setCaeEnabled(true)`; the wrapper produced by the factory applies that setting to every request.
Developer credentials do not support CAE, so use the flag with CI or production credentials.

## Configuration

For a CI service principal using environment credentials, set `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and either
`AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`.

For an Azure Pipelines workload-identity service connection, set:

- `AZURE_TENANT_ID`
- `AZURE_CLIENT_ID`
- `AZURE_SERVICE_CONNECTION_ID` (the service connection resource ID)
- `SYSTEM_ACCESSTOKEN` (map `System.AccessToken` into this environment variable)
- `SYSTEM_OIDCREQUESTURI` (provided by Azure Pipelines)

Production uses a system-assigned managed identity by default. Set `AZURE_MANAGED_IDENTITY_CLIENT_ID` to select a
user-assigned managed identity. Kubernetes workload identity is added as a fallback when `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and `AZURE_FEDERATED_TOKEN_FILE` are all set.

The identity needs an appropriate Azure RBAC role to use Azure Resource Manager. No secret or token is printed.
