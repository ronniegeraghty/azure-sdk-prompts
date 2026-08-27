# Hyoka Basic Azure AI Agent

A synchronous Python console application that creates an Azure AI Agent, asks it
for the capital of France, prints every assistant text response, and deletes the
temporary thread and agent.

## Prerequisites

- Python 3.9 or later
- An Azure AI Foundry project and model deployment
- A Microsoft Entra identity with access to the project
- Local authentication available to `DefaultAzureCredential`

## Restore

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install --upgrade pip
.\.venv\Scripts\python -m pip install -e .
```

## Build

```powershell
.\.venv\Scripts\python -m pip wheel . --no-deps --wheel-dir dist
```

## Run

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
.\.venv\Scripts\hyoka-basic-agent.exe
```
