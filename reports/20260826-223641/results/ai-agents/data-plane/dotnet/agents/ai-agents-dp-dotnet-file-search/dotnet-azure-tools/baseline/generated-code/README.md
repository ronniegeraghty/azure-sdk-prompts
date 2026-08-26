# Hyoka Trail Guide

Authenticate with Azure using `DefaultAzureCredential`, then set the required environment variables.

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
dotnet restore
dotnet build --no-restore
dotnet run --no-build
```
