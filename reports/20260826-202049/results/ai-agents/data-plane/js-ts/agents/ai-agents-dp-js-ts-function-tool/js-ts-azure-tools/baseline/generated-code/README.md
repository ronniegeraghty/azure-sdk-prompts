# Hyoka Weather Agent

A TypeScript console application that uses an Azure AI Agent and a local
`get_weather` function tool to answer a deterministic Seattle weather question.

## Restore, build, and run

```powershell
npm install
npm run build
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
npm start
```

Authentication uses `DefaultAzureCredential`. Configure any supported local
credential before running.
