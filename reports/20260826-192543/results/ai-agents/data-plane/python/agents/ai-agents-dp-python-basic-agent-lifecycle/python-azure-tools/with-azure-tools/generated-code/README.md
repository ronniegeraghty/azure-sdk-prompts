# Hyoka Basic Azure AI Agent

A synchronous Python console application that creates an Azure AI agent, asks it
for the capital of France, prints every assistant text response, and deletes the
created thread and agent.

## Prerequisites

- Python 3.11 or later
- Access to an existing Microsoft Foundry project and model deployment
- A local identity supported by `DefaultAzureCredential`, such as an Azure CLI
  sign-in, with permission to use the project

## Restore, build, and run

PowerShell:

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install --upgrade pip
.\.venv\Scripts\python -m pip install -e .
.\.venv\Scripts\python -m pip wheel --no-deps --wheel-dir dist .
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
.\.venv\Scripts\hyoka-basic-agent.exe
```
