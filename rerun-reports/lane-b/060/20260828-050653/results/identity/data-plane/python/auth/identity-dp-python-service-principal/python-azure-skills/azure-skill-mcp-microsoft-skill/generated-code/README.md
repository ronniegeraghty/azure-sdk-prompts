# Azure service principal authentication (Python)

This example creates an `azure.identity.ClientSecretCredential` from environment
variables and passes it to an Azure Blob Storage SDK client. Its default mode is
offline-safe: it validates configuration and constructs the SDK objects without
sending a request. `--check-auth` explicitly performs authentication and queries
the storage account.

## Requirements

- Python 3.9 or later
- A Microsoft Entra service principal with a client secret
- For the live check, an existing storage account and an appropriate Blob Storage
  data role, such as **Storage Blob Data Reader**, assigned to the service principal

Create and activate a virtual environment, then install the required pip packages:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

The packages are:

- `azure-identity`: provides `ClientSecretCredential`
- `azure-storage-blob`: provides `BlobServiceClient`
- `python-dotenv`: loads a local `.env` file during development

## Configuration

Copy `.env.example` to `.env` and replace the placeholders:

```powershell
Copy-Item .env.example .env
```

`app.py` reads these settings:

| Variable | Purpose |
|---|---|
| `AZURE_TENANT_ID` | Microsoft Entra tenant ID |
| `AZURE_CLIENT_ID` | Application (client) ID |
| `AZURE_CLIENT_SECRET` | Service principal client secret |
| `AZURE_STORAGE_ACCOUNT_URL` | Blob endpoint, such as `https://myaccount.blob.core.windows.net` |

Run the offline-safe configuration check:

```powershell
python app.py
```

Explicitly contact Microsoft Entra ID and Blob Storage:

```powershell
python app.py --check-auth
```

The live check first calls `credential.get_token()` so invalid, expired, or
revoked credentials are reported as authentication failures. It then uses the
same credential with `BlobServiceClient.get_account_information()`. The process
returns a nonzero exit code for configuration, authentication, network, permission,
or service failures.

## Secret management

- Never hardcode or commit client secrets. `.env` is ignored by this project and
  `.env.example` contains placeholders only.
- Use `.env` only for local development. Prefer a secure secret store supplied by
  your CI/CD platform for automation.
- In Azure-hosted production workloads, prefer managed identity instead of a client
  secret. If a secret is unavoidable, store it in Azure Key Vault, restrict access,
  rotate it regularly, and monitor its expiration.
- Give the service principal only the roles it needs and scope assignments as
  narrowly as practical.
- Avoid verbose identity logging in production because diagnostic output can expose
  security-sensitive metadata.
