# Microsoft Foundry project inventory

This console application uses the synchronous `azure-ai-projects` client to list
project connections and model deployments, then retrieve one named connection
and one named model deployment. Authentication uses `DefaultAzureCredential`.

## Install

Requires Python 3.10 or later.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

## Configure and run

Set the project endpoint shown on the Microsoft Foundry project's overview page,
plus an existing connection name and deployment name:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
python app.py
```

`DefaultAzureCredential` supports local Azure CLI sign-in and other standard
Azure Identity credential sources. The application never requests connection
credentials.
