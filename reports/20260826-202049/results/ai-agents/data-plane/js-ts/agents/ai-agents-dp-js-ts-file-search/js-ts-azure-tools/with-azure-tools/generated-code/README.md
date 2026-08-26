# Hyoka Trail Guide Agent

This TypeScript console application uploads a generated trail-guide document to Azure AI
Agents, indexes it for file search, asks a grounded question, prints the assistant response,
and deletes all remote resources it creates.

Requires Node.js 20 or later and Azure credentials supported by `DefaultAzureCredential`.

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
npm install
npm run build
npm start
```
