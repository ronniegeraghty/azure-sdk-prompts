# Basic Azure AI Agent Conversation

Java 17 console application using the synchronous Azure AI Agents Persistent SDK.
It creates an agent and thread, asks for the capital of France, prints assistant
text responses, and deletes the thread and agent.

## Prerequisites

- Java 17 or later
- Apache Maven 3.9 or later
- An Azure identity available to `DefaultAzureCredential`, such as an Azure CLI
  login, and permission to use the Azure AI project

## Configure

In PowerShell:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
```

## Restore, build, and run

```powershell
mvn dependency:resolve
mvn package
mvn exec:java
```
