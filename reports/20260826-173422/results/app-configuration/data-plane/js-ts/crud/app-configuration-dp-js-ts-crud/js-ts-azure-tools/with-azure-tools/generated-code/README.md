# Azure App Configuration CRUD sample

This TypeScript program uses a connection string from the environment to set,
label, get, list, and delete configuration settings and to create a feature
flag.

## Install and build

```powershell
npm install
npm run build
```

Set the connection string without committing it:

```powershell
$env:AZURE_APPCONFIG_CONNECTION_STRING = "Endpoint=https://<store-name>.azconfig.io;Id=<id>;Secret=<secret>"
npm start
```

The delete operation removes the unlabeled `app:Settings:FontSize` setting.
The `Production`-labeled variant and `BetaFeature` feature flag remain in the
store.

Package reference:
[Azure App Configuration client library for JavaScript](https://learn.microsoft.com/javascript/api/overview/azure/app-configuration-readme).
