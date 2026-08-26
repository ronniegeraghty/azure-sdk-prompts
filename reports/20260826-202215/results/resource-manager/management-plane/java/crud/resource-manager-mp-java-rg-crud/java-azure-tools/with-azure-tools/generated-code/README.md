# Azure Resource Group Manager (Java)

This sample uses the modern Azure management-plane SDK to create, list, read, tag, and
delete an Azure resource group in the `eastus` region. The generated name is unique, and
the application attempts deletion in a `finally` block after a successful create.

## Prerequisites

- JDK 17 or later
- Maven 3.9 or later
- An Azure subscription
- An identity with permission to manage resource groups, scoped as narrowly as practical

`DefaultAzureCredential` supports managed identity when the application runs in Azure.
For local development, sign in with a supported developer tool or configure these
environment variables for a service principal:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
```

Do not store credentials in source code. Use managed identity for Azure-hosted workloads.

## Maven dependencies

The complete dependencies are in `pom.xml`:

```xml
<dependency>
    <groupId>com.azure.resourcemanager</groupId>
    <artifactId>azure-resourcemanager</artifactId>
    <version>2.63.0</version>
</dependency>
<dependency>
    <groupId>com.azure</groupId>
    <artifactId>azure-identity</artifactId>
    <version>1.18.2</version>
</dependency>
```

## Build and run

```powershell
mvn clean compile
mvn exec:java
```

Running the application performs real management-plane operations and deletes the resource
group at the end. The source can be built without connecting to Azure.

## References

- [Azure Resource Manager client library for Java](https://learn.microsoft.com/en-us/java/api/overview/azure/resourcemanager-readme?view=azure-java-stable)
- [Authenticate Java apps with Azure Identity](https://learn.microsoft.com/en-us/azure/developer/java/sdk/authentication/overview)
- [AzureResourceManager Java API](https://learn.microsoft.com/en-us/java/api/com.azure.resourcemanager.azureresourcemanager?view=azure-java-stable)
