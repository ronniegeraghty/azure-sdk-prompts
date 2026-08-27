# Contoso Trail Guide Agent

Install:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Run after authenticating with a credential supported by
`DefaultAzureCredential`:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
python .\app.py
```
