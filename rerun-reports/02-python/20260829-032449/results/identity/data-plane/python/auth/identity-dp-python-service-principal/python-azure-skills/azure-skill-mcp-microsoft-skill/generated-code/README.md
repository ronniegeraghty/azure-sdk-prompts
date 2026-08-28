# Azure service principal authentication (Python)

This example authenticates a non-interactive Python application with a Microsoft
Entra service principal and client secret, then uses that credential with the
Azure Blob Storage SDK to list containers.

## Requirements

- Python 3.9 or newer
- An existing Microsoft Entra app registration/service principal
- A client secret for that service principal
- An existing Azure Storage account
- The service principal assigned the least-privilege **Storage Blob Data Reader**
  role at the required scope

This project does not create or modify Azure resources.

## Setup

Create and activate a virtual environment, then install the required packages:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

For local development, copy `.env.example` to `.env` and replace every
placeholder:

```powershell
Copy-Item .env.example .env
python app.py
```

The application constructs the credential explicitly:

```text
ClientSecretCredential(
    tenant_id=AZURE_TENANT_ID,
    client_id=AZURE_CLIENT_ID,
    client_secret=AZURE_CLIENT_SECRET,
)
```

It passes that credential to `BlobServiceClient` and lists the account's blob
containers. Successful authentication does not imply authorization; the
service principal must also have an appropriate Azure RBAC data role.

## Secret-management practices

- Never hardcode or commit client secrets. `.env` is ignored by Git and is only
  for local development.
- In CI/CD, store the values in the platform's encrypted secret store and inject
  them as environment variables.
- In production on Azure, prefer managed identity or workload identity so no
  client secret is needed. If a secret is unavoidable, store it in Azure Key
  Vault, grant least-privilege access, rotate it regularly, and monitor expiry.
- Restrict the service principal to the smallest necessary RBAC role and scope.
- Do not log environment variables, access tokens, or secret values.

## Error handling

The program returns a nonzero exit code and logs a focused message for:

| Exit code | Failure |
|---:|---|
| 2 | Missing or invalid environment configuration |
| 3 | Credential unavailable or Microsoft Entra authentication rejected |
| 4 | Azure Storage network/transport failure |
| 5 | Azure Storage HTTP error, including authorization failures |

The Azure SDK retries eligible transient service failures according to its
built-in retry policy.

## Tests

The tests mock Azure clients and do not contact Azure:

```powershell
python -m unittest -v
```

## References

- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)
- [Authenticate Python apps to Azure services by using service principals](https://learn.microsoft.com/azure/developer/python/sdk/authentication-on-premises-apps)
- [Azure Blob Storage client library for Python](https://learn.microsoft.com/python/api/overview/azure/storage-blob-readme)
