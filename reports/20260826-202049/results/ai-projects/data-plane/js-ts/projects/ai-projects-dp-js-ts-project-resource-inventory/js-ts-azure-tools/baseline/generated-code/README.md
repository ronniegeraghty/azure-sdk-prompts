# Foundry project resource inventory

A TypeScript console application that lists a Microsoft Foundry project's
connections and model deployments, then retrieves one named connection and one
named model deployment.

Authentication uses `DefaultAzureCredential`. Set these environment variables
before running:

- `FOUNDRY_PROJECT_ENDPOINT`: Foundry project endpoint
- `CONNECTION_NAME`: connection to retrieve without credentials
- `DEPLOYMENT_NAME`: model deployment to retrieve

## Install, build, and run

```powershell
npm install
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
npm run build
npm start
```
