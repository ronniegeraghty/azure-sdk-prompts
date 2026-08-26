# Contoso Trail Guide Agent

This console application uploads a generated trail guide, indexes it for file
search, asks a grounded question, prints the assistant response, and deletes all
remote resources before exiting.

## Install

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Authenticate with a credential supported by `DefaultAzureCredential`, then set
the project endpoint and deployed model name:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
```

## Run

```powershell
python .\trail_guide_agent.py
```
