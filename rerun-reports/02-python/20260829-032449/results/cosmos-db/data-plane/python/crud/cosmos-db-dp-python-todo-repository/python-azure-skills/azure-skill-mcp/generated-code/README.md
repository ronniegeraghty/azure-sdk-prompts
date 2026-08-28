# Azure Cosmos DB ToDo Repository

This sample provides synchronous and asynchronous CRUD repositories for the
Azure Cosmos DB for NoSQL Python SDK. It uses Microsoft Entra authentication
through `DefaultAzureCredential`; no account keys or connection strings are
accepted.

## Run

Use Python 3.10 or later:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
$env:COSMOS_ENDPOINT = "https://<account-name>.documents.azure.com:443/"
python main.py
```

`COSMOS_DATABASE_NAME` and `COSMOS_CONTAINER_NAME` are optional and default to
`todo-db` and `items`. The signed-in identity needs Cosmos DB data-plane
permissions to create databases and containers and to read and write items.

New containers use `/category` as the partition key, a 90-day default TTL, and
an indexing policy that excludes `/description/?`. Existing containers are
returned unchanged by the SDK's `create_container_if_not_exists` operation.

Updates require the ETag returned by a repository read. Cosmos DB applies that
ETag as an `If-Match` condition, and stale updates raise
`ConcurrencyConflictError`.

## References

- [Azure Cosmos DB Python quickstart](https://learn.microsoft.com/azure/cosmos-db/quickstart-python)
- [Azure Cosmos DB Python SDK API](https://learn.microsoft.com/python/api/azure-cosmos/)
- [Passwordless connections for Azure services](https://learn.microsoft.com/azure/developer/intro/passwordless-overview)
