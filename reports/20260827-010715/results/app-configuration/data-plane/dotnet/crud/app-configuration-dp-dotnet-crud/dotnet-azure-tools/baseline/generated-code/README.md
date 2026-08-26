# Azure App Configuration CRUD sample

Required NuGet package:

```powershell
dotnet add package Azure.Data.AppConfiguration
```

The project already references this package. Set a connection string, then run the sample:

```powershell
$env:AZURE_APP_CONFIGURATION_CONNECTION_STRING = "<your-connection-string>"
dotnet run
```

The connection string needs permission to read and write configuration settings. The sample creates an
unlabeled setting, a `Production`-labeled setting, and an enabled `BetaFeature` feature flag. It deletes the
unlabeled setting at the end; the labeled setting and feature flag remain in the store.
