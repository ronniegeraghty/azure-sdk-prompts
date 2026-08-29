# Azure Resource Group Manager (Java)

This Maven console application uses the modern Azure management-plane SDK to:

1. Authenticate with `DefaultAzureCredential`.
2. Create a resource group in `eastus`.
3. List the subscription's resource groups.
4. Retrieve the created resource group.
5. add the `managed-by=azure-java-sdk` tag.
6. Delete the resource group.

If an operation fails after creation, the application attempts to delete the resource group in a
`finally` block and logs any cleanup failure.

## Requirements

- JDK 17+
- Maven 3.9+
- An Azure identity with permission to manage resource groups in the target subscription
- `AZURE_SUBSCRIPTION_ID` set to the target subscription ID

`DefaultAzureCredential` supports local developer credentials and workload/managed identity
credentials. For service-principal authentication, set `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
`AZURE_CLIENT_SECRET`. Do not store credentials in source control.

Optionally set `RESOURCE_GROUP_NAME`. If omitted, the application generates a name such as
`java-sdk-rg-a1b2c3d4`.

## Dependencies

```xml
<dependency>
    <groupId>com.azure.resourcemanager</groupId>
    <artifactId>azure-resourcemanager</artifactId>
    <version>2.63.0</version>
</dependency>
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.18.5</version>
</dependency>
```

## Build and run

PowerShell:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:RESOURCE_GROUP_NAME = "my-java-sdk-rg" # optional
mvn compile
mvn exec:java
```

Running the application creates and deletes a real Azure resource group. Review the subscription
and resource-group name before executing it.

## References

- [Azure Resource Manager client library for Java](https://learn.microsoft.com/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable)
- [Authentication with Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/authentication/overview)
