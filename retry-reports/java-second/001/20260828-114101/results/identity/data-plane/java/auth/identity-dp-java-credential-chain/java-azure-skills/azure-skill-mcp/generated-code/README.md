# Azure credential chains for Java

A Java 17 sample that selects an explicit Azure Identity credential chain for local development,
CI/CD, or production. It requests an Azure Resource Manager token without calling any management API.

## Credential strategies

| Environment | Detection examples | Credential order |
|---|---|---|
| Development | Default when no hosted marker is present | Azure CLI, Azure Developer CLI, Azure PowerShell, IntelliJ |
| CI/CD | `TF_BUILD`, `BUILD_BUILDID`, `GITHUB_ACTIONS`, `CI` | Environment credential, then Azure Pipelines workload identity when fully configured |
| Production | `IDENTITY_ENDPOINT`, `MSI_ENDPOINT`, `WEBSITE_INSTANCE_ID`, `KUBERNETES_SERVICE_HOST` | Managed identity, then workload identity when fully configured |

Set `APP_DEPLOYMENT_ENVIRONMENT` to `dev`, `ci`, or `production` to override detection.

For a user-assigned managed identity, set `AZURE_MANAGED_IDENTITY_CLIENT_ID`. If it is absent,
the production chain uses system-assigned managed identity.

The Azure Pipelines workload identity fallback expects:

- `AZURESUBSCRIPTION_CLIENT_ID`
- `AZURESUBSCRIPTION_TENANT_ID`
- `AZURESUBSCRIPTION_SERVICE_CONNECTION_ID`
- `SYSTEM_ACCESSTOKEN`
- `SYSTEM_OIDCREQUESTURI` (provided by Azure Pipelines)

The Kubernetes workload identity fallback expects `AZURE_CLIENT_ID`, `AZURE_TENANT_ID`, and
`AZURE_FEDERATED_TOKEN_FILE`.

CAE is requested by default. Set `AZURE_ENABLE_CAE=false` to disable it. A token request can ask
for CAE, but the target resource and tenant decide whether the returned token carries the `cp1`
capability; the tester reports both the request and the observable token claim.

## Run

```shell
mvn clean test
mvn exec:java
```

For local development, first sign in with one of the configured developer tools, for example
`az login` or `azd auth login`. Authentication values are read only from the process environment;
the sample contains no credentials and creates or modifies no Azure resources.
