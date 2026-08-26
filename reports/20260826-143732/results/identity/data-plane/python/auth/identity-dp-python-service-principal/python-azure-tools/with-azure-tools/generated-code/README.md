# Azure service principal authentication (Python)

This example creates an explicit `ClientSecretCredential`, verifies it can
obtain an Azure Storage token, and uses the credential with
`BlobServiceClient` to list up to 10 containers. It performs read-only
operations and reports configuration, authentication, authorization/service,
and network failures separately.

## Requirements

- Python 3.9 or later
- A Microsoft Entra service principal with an unexpired client secret
- An Azure Storage account
- The service principal assigned the least-privileged role needed for this
  example, normally **Storage Blob Data Reader**, at the narrowest practical
  scope

Install the required pip packages:

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

## Configure and run

For local development, copy `.env.example` to `.env` and replace the
placeholders. `python-dotenv` loads that file without overriding variables
already supplied by the process environment.

```powershell
Copy-Item .env.example .env
python main.py
```

The credential is created from `AZURE_TENANT_ID`, `AZURE_CLIENT_ID`, and
`AZURE_CLIENT_SECRET`. `AZURE_STORAGE_ACCOUNT_URL` must be an HTTPS Blob
service URL such as `https://example.blob.core.windows.net`.

## Secret-management practices

- Never hardcode or commit a client secret. `.env` is ignored by Git and is
  intended only for local development.
- Set environment variables through the deployment platform or CI/CD secret
  store. Restrict who can read or update them and prevent them from appearing
  in logs.
- Store production secrets in Azure Key Vault or the platform's managed secret
  facility, rotate them regularly, and monitor expiration.
- Prefer workload identity federation or managed identity for Azure-hosted
  production workloads because they avoid long-lived client secrets. This
  project uses a client secret because that authentication method is the
  explicit subject of the example.
- Grant the service principal only the data-plane role and scope it needs.

Run the offline tests (Azure calls are mocked):

```powershell
python -m unittest -v
```

## References

- [ClientSecretCredential API](https://learn.microsoft.com/python/api/azure-identity/azure.identity.clientsecretcredential)
- [Azure Identity client-secret authentication](https://learn.microsoft.com/azure/developer/python/sdk/authentication-on-premises-apps)
- [BlobServiceClient API](https://learn.microsoft.com/python/api/azure-storage-blob/azure.storage.blob.blobserviceclient)
