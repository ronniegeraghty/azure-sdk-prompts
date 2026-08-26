# Trail Guide Agent

This console app uploads a generated trail-guide document, indexes it for file search, asks a grounded question, prints the assistant response, and deletes every created Azure resource.

Set `PROJECT_ENDPOINT` and `MODEL_DEPLOYMENT_NAME`, then authenticate with a credential supported by `DefaultAzureCredential` (for example, Azure CLI or environment credentials).

```powershell
dotnet restore
dotnet build --no-restore
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<deployment-name>"
dotnet run --no-build
```
