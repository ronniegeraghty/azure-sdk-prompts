# Azure Cosmos DB CRUD with TypeScript

This sample creates a database and container, then creates, reads, queries,
replaces, and deletes an item with the Azure Cosmos DB for NoSQL SDK.

## Requirements

- Node.js 20 or later
- An Azure Cosmos DB for NoSQL endpoint and account key

Install the required runtime package:

```powershell
npm install @azure/cosmos
```

Install all project dependencies and compile:

```powershell
npm install
npm run build
```

Set credentials without placing them in source control, then run the sample:

```powershell
$env:COSMOS_ENDPOINT = "https://<account>.documents.azure.com:443/"
$env:COSMOS_KEY = "<account-key>"
npm start
```

The program uses `createIfNotExists` for `TestDB` and `Items`, whose partition
key is `/category`. It creates a uniquely identified item and deletes it after
the CRUD sequence. For production workloads, prefer Microsoft Entra ID and a
managed identity over account-key authentication.

SDK reference:
https://learn.microsoft.com/javascript/api/overview/azure/cosmos-readme
