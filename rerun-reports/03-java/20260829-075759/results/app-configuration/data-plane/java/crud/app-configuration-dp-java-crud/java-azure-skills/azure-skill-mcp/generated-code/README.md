# Azure App Configuration Java CRUD sample

This sample uses `com.azure:azure-data-appconfiguration` to create, read, list,
and delete configuration settings and to create a feature flag.

Set the connection string in an environment variable rather than placing it in
source code:

```powershell
$env:AZURE_APPCONFIGURATION_CONNECTION_STRING = "<connection-string>"
mvn compile
mvn exec:java -Dexec.mainClass=com.example.AppConfigurationManager
```

The program deletes only the unlabeled `app:Settings:FontSize` setting. The
`Production`-labeled setting and `BetaFeature` feature flag remain in the store.

For production workloads, prefer Microsoft Entra authentication with managed
identity over a connection string.

## References

- [Azure App Configuration client library for Java](https://learn.microsoft.com/java/api/overview/azure/data-appconfiguration-readme?view=azure-java-stable)
- [Exception handling in the Azure SDK for Java](https://learn.microsoft.com/azure/developer/java/sdk/troubleshooting-overview#exception-handling-in-the-azure-sdk-for-java)
- [Maven Central artifact](https://central.sonatype.com/artifact/com.azure/azure-data-appconfiguration/1.10.1)
