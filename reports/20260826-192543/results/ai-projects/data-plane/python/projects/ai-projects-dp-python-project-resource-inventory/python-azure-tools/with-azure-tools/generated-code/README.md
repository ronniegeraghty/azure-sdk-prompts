# Microsoft Foundry Project Inventory

This synchronous Python console application lists a Microsoft Foundry project's
connections and model deployments, then retrieves one named connection and one
named model deployment. Connection credentials are never requested.

## Install

Python 3.10 or later is required.

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Authenticate locally with a supported `DefaultAzureCredential` source, such as
Azure CLI or Visual Studio Code. In production, use managed identity and set
`AZURE_TOKEN_CREDENTIALS=prod` to constrain the credential chain.

## Run

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
python app.py
```
