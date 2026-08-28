# Azure Cosmos DB ToDo repository

This sample provides synchronous and asynchronous Python repositories for ToDo
items stored in the Azure Cosmos DB NoSQL API. It uses Microsoft Entra
authentication through `DefaultAzureCredential`; no account keys are accepted.

## Setup

1. Create and activate a Python 3.10 or newer virtual environment.
2. Install dependencies with `pip install -r requirements.txt`.
3. Grant the signed-in identity an appropriate Cosmos DB data-plane role, such
   as **Cosmos DB Built-in Data Contributor**.
4. Set `COSMOS_ENDPOINT` to the account endpoint. Optionally set
   `COSMOS_DATABASE` and `COSMOS_CONTAINER`.
5. Run `python main.py`.

The identity also needs permission to create the configured database and
container. The container uses `/category` as its partition key, a 90-day
default TTL, and an indexing policy that excludes `/description/?`.
