# Azure Resource Group Manager

This Java 17 Maven sample uses `DefaultAzureCredential` and the modern
`com.azure.resourcemanager:azure-resourcemanager` management-plane SDK to:

1. Create a resource group in `eastus`.
2. List the subscription's resource groups.
3. Fetch the created resource group.
4. Add a `managed-by=java-sdk-sample` tag.
5. Delete the resource group.

The application also attempts to delete the created resource group if a later
operation fails.

## Dependencies

The `pom.xml` imports Azure SDK BOM `1.3.8`, which currently selects:

- `com.azure.resourcemanager:azure-resourcemanager:2.63.0`
- `com.azure:azure-identity:1.18.4`

Using the BOM keeps shared Azure SDK dependencies on compatible versions.

## Configuration

Set `AZURE_SUBSCRIPTION_ID`. Configure any identity supported by
`DefaultAzureCredential`; for example, sign in locally with a supported
developer credential, or use managed identity when running in Azure. The
identity needs permission to read, create, tag, and delete resource groups at
the subscription scope.

No credentials are stored in the source code.

## Build and run

```powershell
mvn clean package
mvn exec:java -Dexec.args="example-java-sdk-rg"
```

Running the second command creates and then deletes a real Azure resource
group. Resource group names must be unique within the subscription.

## References

- [Azure Resource Manager SDK for Java](https://learn.microsoft.com/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable)
- [Azure authentication with Java and DefaultAzureCredential](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview)
