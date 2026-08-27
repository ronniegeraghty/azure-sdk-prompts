# Hyoka Basic Azure AI Agent

A synchronous Python console application that creates an Azure AI agent, asks it
for the capital of France, prints all assistant text responses, and deletes the
created thread and agent.

## Prerequisites

- Python 3.10 or later
- An Azure AI project with a deployed model
- Credentials supported by `DefaultAzureCredential` with access to the project

Set the required environment variables in PowerShell:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
```

## Restore, build, and run

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install --upgrade pip
.\.venv\Scripts\python -m pip install -e ".[dev]"
.\.venv\Scripts\python -m build
.\.venv\Scripts\python -m hyoka_basic_agent
```

After an editable install, the final command can also be:

```powershell
.\.venv\Scripts\hyoka-basic-agent.exe
```
