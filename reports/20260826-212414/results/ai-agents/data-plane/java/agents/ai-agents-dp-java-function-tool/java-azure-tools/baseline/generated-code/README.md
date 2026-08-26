# Persistent weather agent

Java 17 console application using the synchronous Azure AI Persistent Agents SDK and a local `get_weather` function tool.

## Restore and build

```powershell
mvn dependency:go-offline
mvn clean package
```

## Run

Authenticate locally with a credential supported by `DefaultAzureCredential`, then set the project endpoint and model deployment:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
mvn exec:java
```
