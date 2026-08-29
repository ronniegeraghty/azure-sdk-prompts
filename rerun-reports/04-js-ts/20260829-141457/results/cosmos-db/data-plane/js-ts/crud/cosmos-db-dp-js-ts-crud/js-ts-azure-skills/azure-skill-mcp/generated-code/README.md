# Azure Cosmos DB TypeScript CRUD sample

This sample uses `@azure/cosmos` to create `TestDB` and an `Items` container,
then creates, reads, queries, replaces, and deletes an item.

Install dependencies:

```powershell
npm install
```

Set the Azure Cosmos DB for NoSQL endpoint and key without storing credentials
in source:

```powershell
$env:COSMOS_ENDPOINT = "https://<account-name>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
npm start
```

The required runtime package is:

```powershell
npm install @azure/cosmos
```

For production Azure-hosted applications, prefer Microsoft Entra ID and managed
identity over account keys.

References:

- [Azure Cosmos DB JavaScript SDK](https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme?view=azure-node-latest)
- [Create and access items](https://learn.microsoft.com/azure/cosmos-db/how-to-javascript-create-item)
- [Official local development sample](https://learn.microsoft.com/azure/cosmos-db/development-loop#create-the-application)
