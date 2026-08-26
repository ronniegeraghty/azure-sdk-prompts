# Foundry Project Inventory

Read-only .NET console application that inventories connections and model
deployments in a Microsoft Foundry project by using `Azure.AI.Projects`.
Authentication uses `DefaultAzureCredential`; sign in with a supported developer
credential or configure a workload identity before running.

## Configure

PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
```

## Restore, build, and run

```powershell
dotnet restore
dotnet build --no-restore
dotnet run --no-build
```
