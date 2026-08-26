# Azure AI Agents document-grounded console app

## Install, build, and run

```powershell
npm install
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
npm run build
npm start
```

Authentication uses `DefaultAzureCredential`. Configure any supported local credential, such as Azure CLI login, before running.
