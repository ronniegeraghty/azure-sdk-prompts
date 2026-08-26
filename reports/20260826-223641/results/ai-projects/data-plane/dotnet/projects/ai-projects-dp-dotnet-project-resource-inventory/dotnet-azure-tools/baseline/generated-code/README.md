# Foundry Project Inventory

This console application uses `Azure.AI.Projects` and `DefaultAzureCredential` to
list and retrieve connections and model deployments from a Microsoft Foundry
project.

Set the required environment variables in PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
```

Restore, build, and run:

```powershell
dotnet restore
dotnet build --no-restore
dotnet run --no-build
```

Authenticate locally with a credential supported by `DefaultAzureCredential`,
such as Azure CLI login, Visual Studio, or environment-based service principal
credentials. The application requests connection metadata without credentials.
