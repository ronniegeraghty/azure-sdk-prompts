# Weather Agent

A .NET 8 console application that uses `Azure.AI.Agents.Persistent` and a local
`get_weather` function tool to answer a deterministic Seattle weather question.

Authenticate locally with a credential supported by `DefaultAzureCredential`,
then run these PowerShell commands:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
dotnet restore
dotnet build --no-restore
dotnet run --no-build
```

The application creates an agent and thread for the run and deletes both before
it exits.
