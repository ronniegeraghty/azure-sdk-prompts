# Azure AI Agents file-search console app

Requires Python 3.10 or later and Azure credentials supported by
`DefaultAzureCredential`.

Install:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Run:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
python app.py
```
