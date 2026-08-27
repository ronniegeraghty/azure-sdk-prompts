# Microsoft Foundry project inventory

This synchronous Python console application lists a Microsoft Foundry project's
connections and model deployments, then retrieves one named connection and one
named model deployment. Connection credentials are never requested or printed.

## Prerequisites

- Python 3.10 or later
- A Microsoft Entra identity with permission to read the Foundry project
- Local authentication available to `DefaultAzureCredential`, such as an Azure
  CLI or Visual Studio Code sign-in

## Install and run (PowerShell)

```powershell
python -m venv .venv
.\.venv\Scripts\python -m pip install -r requirements.txt

$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
.\.venv\Scripts\python app.py
```

For production-hosted execution, configure managed identity or workload identity
and set `AZURE_TOKEN_CREDENTIALS=prod` to constrain `DefaultAzureCredential` to
production-safe credential types.
