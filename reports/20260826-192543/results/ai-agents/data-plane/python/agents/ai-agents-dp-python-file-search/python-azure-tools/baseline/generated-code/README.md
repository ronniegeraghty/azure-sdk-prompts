# Azure AI Agents document-grounded console app

Requires Python 3.9 or later, an Azure AI Foundry project endpoint, a deployed
model, and credentials available to `DefaultAzureCredential`.

Install and run from PowerShell:

```powershell
py -m venv .venv
.\.venv\Scripts\python -m pip install -r requirements.txt
$env:PROJECT_ENDPOINT = "https://your-project-endpoint"
$env:MODEL_DEPLOYMENT_NAME = "your-model-deployment"
.\.venv\Scripts\python app.py
```

The app writes `contoso_trail_guide.txt`, uploads and indexes it, prints the
assistant's grounded answer, and deletes all remote resources it created.
