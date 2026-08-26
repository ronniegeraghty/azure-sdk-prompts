# Hyoka Basic Azure AI Agent

## Prerequisites

- Node.js 20 or later
- Azure credentials supported by `DefaultAzureCredential`
- An Azure AI project and model deployment

Set the required environment variables in PowerShell:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
```

Restore, build, and run:

```powershell
npm install
npm run build
npm start
```
