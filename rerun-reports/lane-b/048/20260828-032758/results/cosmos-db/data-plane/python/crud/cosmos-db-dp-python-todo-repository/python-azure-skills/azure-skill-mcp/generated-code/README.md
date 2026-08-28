# Azure Cosmos DB ToDo Repository

This project demonstrates synchronous and asynchronous ToDo repositories using
the Azure Cosmos DB for NoSQL Python SDK. Authentication uses
`DefaultAzureCredential`; account keys are not supported.

## Setup

1. Create and activate a Python 3.10 or newer virtual environment.
2. Install dependencies with `python -m pip install -r requirements.txt`.
3. Grant your identity a Cosmos DB data-plane role that can create databases,
   containers, and items.
4. Set `AZURE_COSMOS_ENDPOINT` to the account endpoint. Optionally set
   `AZURE_COSMOS_DATABASE` and `AZURE_COSMOS_CONTAINER`.
5. Run `python main.py`.

The factory creates the database and container when absent. The container uses
`/category` as its partition key, a 90-day default TTL, and an indexing policy
that excludes the `description` property.
