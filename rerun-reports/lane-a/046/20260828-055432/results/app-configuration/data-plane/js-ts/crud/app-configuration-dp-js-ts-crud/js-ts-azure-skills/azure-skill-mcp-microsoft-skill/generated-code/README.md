# Azure App Configuration CRUD sample

Install the required Azure SDK package and the package that exports `RestError`:

```powershell
npm install @azure/app-configuration @azure/core-rest-pipeline
```

Install all project dependencies:

```powershell
npm install
```

Set the connection string without committing it:

```powershell
$env:AZURE_APPCONFIG_CONNECTION_STRING = "Endpoint=https://<store>.azconfig.io;Id=<id>;Secret=<secret>"
npm start
```

The program creates an unlabeled setting and a `Production`-labeled setting,
reads and lists settings, creates the `BetaFeature` feature flag, and deletes
the unlabeled setting. The labeled setting and feature flag remain in the
store.
