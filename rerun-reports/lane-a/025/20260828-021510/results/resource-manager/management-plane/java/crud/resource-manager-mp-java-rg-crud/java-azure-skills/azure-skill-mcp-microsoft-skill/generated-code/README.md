# Azure Resource Group Manager

This Maven application uses `DefaultAzureCredential` and the modern
`azure-resourcemanager` management-plane SDK to:

1. Create a resource group in `eastus`.
2. List all resource groups in the subscription.
3. Retrieve and display the created resource group.
4. Add a `managed-by=azure-sdk-for-java` tag.
5. Delete the resource group.

The required Maven dependencies are declared in `pom.xml`:

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

## Run

Use Java 17 or later. Configure `DefaultAzureCredential` for your environment
(for example, Azure CLI credentials for local development or managed identity
in Azure), then set the subscription:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "<optional-resource-group-name>"
mvn compile exec:java
```

If `AZURE_RESOURCE_GROUP_NAME` is omitted, the application generates a unique
name. The signed-in identity needs permission to read, create, update, and
delete resource groups in the subscription. The application deletes the
created resource group during normal execution and attempts cleanup if a later
operation fails.
