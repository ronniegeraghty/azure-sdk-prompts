# Basic Azure AI Agent conversation

This TypeScript console application creates an Azure AI Agent, asks it for the
capital of France, prints every assistant text response, and deletes the thread
and agent.

Requires Node.js 20 or later and credentials supported by
`DefaultAzureCredential`.

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
npm install
npm run build
npm start
```
