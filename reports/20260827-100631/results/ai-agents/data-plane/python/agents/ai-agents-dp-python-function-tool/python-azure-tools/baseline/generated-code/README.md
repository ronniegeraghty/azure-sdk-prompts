# Hyoka weather agent

## Install

```powershell
py -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

## Run

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
python .\weather_agent.py
```

Authentication uses `DefaultAzureCredential`. Configure one of its supported
credential sources before running the application.
