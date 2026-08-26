# Foundry Project Inventory

Java 17 console application that uses the synchronous `azure-ai-projects` clients to inspect a Microsoft Foundry project's connections and model deployments. Authentication uses `DefaultAzureCredential`; sign in locally with a supported developer credential before running.

## Configuration

Set these environment variables in PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
```

## Restore, build, and run

```powershell
mvn dependency:resolve
mvn package
mvn exec:java
```
