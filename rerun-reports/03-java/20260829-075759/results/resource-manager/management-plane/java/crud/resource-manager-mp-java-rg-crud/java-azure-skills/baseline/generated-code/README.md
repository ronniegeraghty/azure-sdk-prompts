# Azure Resource Group Manager

This Maven application uses `DefaultAzureCredential` and the modern
`azure-resourcemanager` SDK to create, list, retrieve, tag, and delete an Azure
Resource Group. Deletion runs in a `finally` block so the sample attempts to
clean up a group that it successfully created even if a later operation fails.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An Azure identity supported by `DefaultAzureCredential`
- Permission to manage Resource Groups in the selected subscription

For service-principal authentication, set placeholder-based environment
variables in your local shell:

```powershell
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_RESOURCE_GROUP_NAME = "<unique-resource-group-name>"
```

Running without arguments is a local-only dry run:

```powershell
mvn compile exec:java
```

To deliberately execute the management operations against Azure:

```powershell
mvn compile exec:java -Dexec.args="--execute"
```

The required dependencies are declared in `pom.xml`:

- `com.azure.resourcemanager:azure-resourcemanager:2.63.0`
- `com.azure:azure-identity:1.18.5`

Use a disposable Resource Group name. The delete step can fail because of
authorization, resource locks, or transient service errors; if that happens,
the application reports the failure and the group must be inspected manually.
