# Hyoka Basic Azure AI Agent

Java 17 console application using the synchronous Azure AI Persistent Agents client.

Set `PROJECT_ENDPOINT` to the Azure AI project endpoint and
`MODEL_DEPLOYMENT_NAME` to the model deployment name. `DefaultAzureCredential`
must be able to obtain credentials in the current environment.

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
```

Restore dependencies:

```powershell
mvn dependency:go-offline
```

Build:

```powershell
mvn package
```

Run:

```powershell
mvn exec:java
```
