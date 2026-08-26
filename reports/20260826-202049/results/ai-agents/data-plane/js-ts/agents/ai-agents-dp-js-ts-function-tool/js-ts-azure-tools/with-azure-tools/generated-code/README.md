# Hyoka weather agent

Set `PROJECT_ENDPOINT` and `MODEL_DEPLOYMENT_NAME`, then authenticate with a
credential supported by `DefaultAzureCredential` (for example, Azure CLI
authentication).

```powershell
npm install
npm run build
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
npm start
```
