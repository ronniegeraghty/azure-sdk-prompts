# Azure Resource Group Manager (Java)

This Maven application uses the modern Azure management-plane SDK to create,
list, retrieve, tag, and delete an Azure resource group. It uses
`DefaultAzureCredential`, so it can authenticate with a service principal,
managed identity, Azure CLI login, or another supported credential source.

## Requirements

- Java 17 or later
- Apache Maven 3.9 or later
- An Azure identity with permission to manage resource groups
- `AZURE_SUBSCRIPTION_ID` set to the target subscription

For service-principal authentication, also set `AZURE_TENANT_ID`,
`AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. Do not store credentials in source
control.

The required SDK dependencies are declared in `pom.xml`:

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

Compile the application:

```powershell
mvn compile
```

Run locally without contacting Azure:

```powershell
mvn exec:java
```

To execute the real Azure CRUD sequence, set the subscription and explicitly
opt in:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:RESOURCE_GROUP_NAME = "rg-java-sdk-demo" # Optional
mvn exec:java '-Dexec.args=--execute'
```

The application deletes the resource group after tagging it. If a later
operation fails after creation, the `finally` block attempts cleanup.
