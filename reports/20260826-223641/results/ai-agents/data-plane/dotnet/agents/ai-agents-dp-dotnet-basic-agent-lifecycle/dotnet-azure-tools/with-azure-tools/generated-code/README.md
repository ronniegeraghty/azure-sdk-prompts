# Basic Azure AI Agent Conversation

A .NET console application that creates a temporary Azure AI agent and thread,
asks the capital of France, prints the assistant response, and deletes the
created resources.

## Run

Set `PROJECT_ENDPOINT` to the Azure AI project endpoint and
`MODEL_DEPLOYMENT_NAME` to an existing model deployment.

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"

dotnet restore
dotnet build --no-restore
dotnet run --no-build
```

`DefaultAzureCredential` is used for authentication. For local development,
sign in with a supported developer credential before running the application.
