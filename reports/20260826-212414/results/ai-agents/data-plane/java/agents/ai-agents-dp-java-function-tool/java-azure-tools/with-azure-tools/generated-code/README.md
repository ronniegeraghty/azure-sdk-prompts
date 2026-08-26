# Persistent Weather Agent

Java 17 console application using the synchronous Azure AI Agents Persistent SDK and a local `get_weather` function.

Set the required environment variables in PowerShell:

    $env:PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
    $env:MODEL_DEPLOYMENT_NAME = "<model-deployment-name>"

`DefaultAzureCredential` must also be able to obtain an Azure credential from the local environment.

Restore dependencies:

    mvn dependency:go-offline

Build:

    mvn package

Run:

    mvn exec:java
