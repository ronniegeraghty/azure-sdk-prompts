# Persistent agent file-search sample

This console application creates `contoso-trail-guide.txt`, uploads it to a Microsoft Foundry project, indexes it in a vector store, and asks a persistent agent a grounded question. It uses synchronous SDK clients and deletes all remote resources before exiting.

## Prerequisites

- Java 17 or later
- Maven 3.9 or later
- Azure credentials supported by `DefaultAzureCredential` (for example, Azure CLI login or service-principal environment variables)
- `PROJECT_ENDPOINT` set to the Microsoft Foundry project endpoint
- `MODEL_DEPLOYMENT_NAME` set to a model deployment that supports agents and file search

## Restore, build, and run

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"

mvn dependency:resolve
mvn package
mvn exec:java
```
