# Microsoft Foundry project inventory

This console application uses the synchronous `azure-ai-projects` client and
`DefaultAzureCredential` to inspect a Microsoft Foundry project's connections
and model deployments. It reads configuration from the process environment and
never requests connection credentials.

## Install

Python 3.10 or later is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Authenticate locally with a credential supported by `DefaultAzureCredential`,
then set the project values:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://your-resource.services.ai.azure.com/api/projects/your-project"
$env:CONNECTION_NAME = "your-connection-name"
$env:DEPLOYMENT_NAME = "your-deployment-name"
```

## Run

```powershell
python app.py
```
