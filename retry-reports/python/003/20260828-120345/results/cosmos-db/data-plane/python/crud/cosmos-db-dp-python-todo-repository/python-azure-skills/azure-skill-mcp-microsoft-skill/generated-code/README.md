# Azure Cosmos DB ToDo Repository

This project demonstrates synchronous and asynchronous CRUD repositories for
Azure Cosmos DB for NoSQL. It uses Microsoft Entra authentication, optimistic
concurrency through Cosmos DB etags, parameterized single-partition queries,
and page-by-page result processing.

## Setup

1. Create a virtual environment and install `requirements.txt`.
2. Set `COSMOS_ENDPOINT` to the account endpoint.
3. Optionally set `COSMOS_DATABASE_NAME` and `COSMOS_CONTAINER_NAME`.
4. Authenticate locally with a credential supported by
   `DefaultAzureCredential`, or use a managed identity in Azure.
5. Grant the identity sufficient Cosmos DB data-plane RBAC permissions to
   create and access the database and container.
6. Run `python main.py`.

The container is created with `/category` as its partition key, a 90-day
default TTL, and an indexing policy that excludes `description`.
