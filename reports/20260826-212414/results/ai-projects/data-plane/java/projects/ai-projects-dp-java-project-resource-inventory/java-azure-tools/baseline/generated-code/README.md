# Foundry Project Inventory

A Java 17 console application that uses the synchronous Azure AI Projects SDK to inspect the connections and model deployments in a Microsoft Foundry project.

Authentication uses `DefaultAzureCredential`. Configure one of its supported local credential sources, such as an Azure CLI login, and set:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://your-resource.services.ai.azure.com/api/projects/your-project"
$env:CONNECTION_NAME = "your-connection"
$env:DEPLOYMENT_NAME = "your-model-deployment"
```

Restore dependencies:

```powershell
mvn dependency:go-offline
```

Build the executable JAR:

```powershell
mvn clean package
```

Run:

```powershell
java -jar target\foundry-project-inventory-1.0.0-all.jar
```

The application lists all connections through the SDK pageable API, retrieves the requested connection without credential values, lists all deployments while printing typed model deployment details, and validates that the requested deployment is a model deployment.
