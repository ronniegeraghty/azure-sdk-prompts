# Hyoka Basic Azure AI Agent

This console application uses the synchronous `azure-ai-agents` client to ask
an Azure AI Agent for the capital of France, print its response, and delete the
created thread and agent.

## Restore and build

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install -e .
.\.venv\Scripts\python -m compileall -q hyoka_basic_agent.py
```

## Run

Authenticate `DefaultAzureCredential` using a supported local identity, then
set the project endpoint and model deployment name:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
.\.venv\Scripts\hyoka-basic-agent.exe
```
