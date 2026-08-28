# Azure Storage account management sample

This Java 11 sample uses `DefaultAzureCredential` and the Azure Storage
management-plane SDK to create, list, inspect, update, and delete a storage
account. The resource group must already exist.

Set `AZURE_SUBSCRIPTION_ID`, authenticate with a supported local developer
credential, and run:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
mvn compile exec:java -Dexec.args="<resource-group> <globally-unique-storage-account-name>"
```

The signed-in identity needs permission to manage storage accounts in the
target resource group. The program deletes the created account in a `finally`
block, including when an intermediate operation fails.

`DefaultAzureCredential` is convenient for local development. For an
Azure-hosted production application, use a specific managed identity
credential and least-privilege Azure RBAC instead.
