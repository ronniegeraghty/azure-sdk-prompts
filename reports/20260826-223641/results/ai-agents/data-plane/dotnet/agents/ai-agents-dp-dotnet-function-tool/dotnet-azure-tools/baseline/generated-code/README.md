# Weather Agent

A .NET console application that uses an Azure AI Foundry persistent agent and a
local `get_weather` function tool to answer a deterministic Seattle weather
question.

## Run

Authenticate locally with an Azure credential supported by
`DefaultAzureCredential`, then set the required environment variables:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
dotnet restore
dotnet build --no-restore
dotnet run --no-build
```
