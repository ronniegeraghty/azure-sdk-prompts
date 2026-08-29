# Azure service principal authentication

This TypeScript sample uses `ClientSecretCredential` to authenticate a service
principal and creates an Azure Key Vault `SecretClient`. It requests one page of
secret metadata to verify both the credential and the service principal's Key
Vault data-plane authorization. It does not retrieve or print secret values.

## Run

Use Node.js 20 or newer:

```powershell
npm install
Copy-Item .env.example .env
# Edit .env with the service principal and Key Vault values.
npm run dev
```

The service principal needs permission to list secret metadata. With Azure RBAC,
assign the narrowest role that permits the required operation at the appropriate
vault scope. Never commit `.env`, and rotate the client secret regularly.

To build and run the compiled program:

```powershell
npm run build
npm start
```

## References

- [Authenticate JavaScript apps to Azure with a service principal](https://learn.microsoft.com/azure/developer/javascript/sdk/authentication/service-principal)
- [Azure Key Vault Secrets client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/keyvault-secrets-readme)
