# Azure credential chain demo

Small Java 17 sample that selects an explicit Azure Identity credential chain for local development,
CI/CD, or production and tests it against the Azure Resource Manager scope. It only requests tokens;
it does not create or modify Azure resources.

## Credential strategies

| Environment | Detection | Credential order |
| --- | --- | --- |
| Development | Default when no hosted-environment marker is present | Azure CLI, IntelliJ, Visual Studio Code, Azure Developer CLI, Azure PowerShell |
| CI/CD | `CI`, `TF_BUILD`, `BUILD_BUILDID`, `PIPELINE_WORKSPACE`, or another supported CI marker | Azure Pipelines workload identity service connection when fully configured, then `EnvironmentCredential` |
| Production | A managed identity endpoint variable or complete workload identity configuration | System- or user-assigned managed identity, then Kubernetes workload identity when configured |

Set `APP_ENVIRONMENT` to `dev`, `ci`, or `production` to override auto-detection.

## Configuration

- Pipeline service principal: `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and either
  `AZURE_CLIENT_SECRET` or `AZURE_CLIENT_CERTIFICATE_PATH`.
- Azure Pipelines workload identity service connection: `TF_BUILD`, `SYSTEM_OIDCREQUESTURI`,
  `SYSTEM_ACCESSTOKEN`, `AZURE_SERVICE_CONNECTION_ID`, `AZURE_TENANT_ID`, and `AZURE_CLIENT_ID`.
  Map the pipeline's `System.AccessToken` secret into `SYSTEM_ACCESSTOKEN`.
- User-assigned managed identity: `AZURE_MANAGED_IDENTITY_CLIENT_ID`.
  Leave it unset to use the system-assigned identity.
- Kubernetes workload identity fallback: `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
  `AZURE_FEDERATED_TOKEN_FILE`.
- CAE: set `AZURE_ENABLE_CAE=true`. CAE is requested through `TokenRequestContext`; Microsoft Entra ID
  and the target resource decide whether a CAE token is issued. Azure Identity developer credentials
  don't support CAE, so normally leave CAE disabled for local development.

Never commit secret values or pipeline access tokens.

## Build and run

```text
mvn clean package
mvn exec:java
```

The program performs a synchronous token request and then an asynchronous request for
`https://management.azure.com/.default`. It prints no token contents.

## References

- [Credential chains in Azure Identity for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/credential-chains)
- [Azure Identity client library for Java](https://learn.microsoft.com/java/api/overview/azure/identity-readme)
- [Continuous Access Evaluation](https://learn.microsoft.com/entra/identity/conditional-access/concept-continuous-access-evaluation)
