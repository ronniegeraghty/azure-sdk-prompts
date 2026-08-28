# Azure App Configuration CRUD sample

This .NET 8 console application uses `Azure.Data.AppConfiguration` 1.11.1. No
additional NuGet package is required for the requested operations.

Install the package manually with:

```powershell
dotnet add package Azure.Data.AppConfiguration --version 1.11.1
```

Set the connection string without placing it in source code, then run the app:

```powershell
$env:AZURE_APPCONFIG_CONNECTION_STRING = "<your-app-configuration-connection-string>"
dotnet run
```

The connection string must grant read/write access. The sample creates the
default and `Production` versions of `app:Settings:FontSize`, lists matching
settings, creates the enabled `BetaFeature` feature flag, and deletes the
default-label font-size setting.

API reference:

- <https://learn.microsoft.com/dotnet/api/overview/azure/data.appconfiguration-readme>
- <https://learn.microsoft.com/dotnet/api/azure.data.appconfiguration.configurationclient.getconfigurationsettings>
- <https://learn.microsoft.com/dotnet/api/azure.data.appconfiguration.featureflagconfigurationsetting.-ctor>
