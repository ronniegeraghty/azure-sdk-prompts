# Azure Resource Group Manager

This Maven application uses `DefaultAzureCredential` and the modern
`azure-resourcemanager` fluent SDK to create, list, retrieve, tag, and delete an
Azure resource group.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- An Azure identity with permission to manage resource groups
- `AZURE_SUBSCRIPTION_ID` set to the target subscription

`DefaultAzureCredential` checks supported credential sources in order. For
local development, authenticate with a supported developer credential or set
`AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and `AZURE_CLIENT_SECRET`. In Azure, use
a managed identity where possible.

## Build and run

```powershell
mvn compile
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
mvn exec:java -Dexec.args="java-sdk-resource-group"
```

Omit `-Dexec.args` to generate a unique resource group name. The application
creates the group in `eastus` and deletes it in a `finally` block, including
when a later management operation fails.
