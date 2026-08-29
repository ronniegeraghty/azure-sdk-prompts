# Azure App Configuration CRUD example

Install the required Azure App Configuration SDK and supporting packages:

```powershell
npm install
```

Set the connection string without storing it in source control, then run the
example:

```powershell
$env:AZURE_APPCONFIG_CONNECTION_STRING = "Endpoint=https://<store>.azconfig.io;Id=<id>;Secret=<secret>"
npm start
```

The program creates an unlabeled setting and a `Production`-labeled variant,
reads and lists settings, creates the `BetaFeature` feature flag, and deletes
the unlabeled setting. Azure App Configuration identifies a setting by both
key and label, so the labeled variant remains after the final delete.

`@azure/core-rest-pipeline` is installed directly because it exports
`RestError`, which the example uses for service-specific error handling.

SDK references:

- https://learn.microsoft.com/javascript/api/@azure/app-configuration/appconfigurationclient
- https://learn.microsoft.com/javascript/api/@azure/core-rest-pipeline/resterror
