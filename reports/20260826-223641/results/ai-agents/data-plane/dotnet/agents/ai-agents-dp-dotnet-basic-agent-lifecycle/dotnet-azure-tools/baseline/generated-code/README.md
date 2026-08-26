# Hyoka Basic Azure AI Agent

A .NET 8 console application that creates an Azure AI Agent, asks it for the
capital of France, prints the assistant response, and deletes the created
thread and agent.

## Configuration

Set `PROJECT_ENDPOINT` to the Azure AI Foundry project endpoint and
`MODEL_DEPLOYMENT_NAME` to the model deployment name. Authentication uses
`DefaultAzureCredential`.

PowerShell example:

    $env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
    $env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"

## Restore, build, and run

    dotnet restore
    dotnet build --no-restore
    dotnet run --no-build
