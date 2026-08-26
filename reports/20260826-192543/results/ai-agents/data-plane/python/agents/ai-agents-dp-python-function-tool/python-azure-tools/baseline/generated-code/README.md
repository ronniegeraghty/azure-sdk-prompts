# Azure AI Agents weather console app

Install the dependencies:

```powershell
python -m pip install -r requirements.txt
```

Set the Azure AI Foundry project endpoint and model deployment, then run:

```powershell
$env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"
python .\weather_agent.py
```

Authentication uses `DefaultAzureCredential`. Sign in with a supported local
developer credential before running the application.
