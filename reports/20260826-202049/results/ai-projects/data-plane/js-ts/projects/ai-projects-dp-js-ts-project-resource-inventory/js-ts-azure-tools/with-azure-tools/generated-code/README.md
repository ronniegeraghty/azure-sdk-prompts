# Foundry project resource inventory

A TypeScript console application that uses `@azure/ai-projects` to inspect
connection metadata and model deployments in a Microsoft Foundry project. It
uses `DefaultAzureCredential` and never requests connection credentials.

## Prerequisites

- Node.js 20 or later
- A Microsoft Entra identity with read access to the Foundry project
- Local authentication supported by `DefaultAzureCredential`, such as Azure
  CLI, Azure Developer CLI, or environment-based service principal credentials

## Configure

Set these environment variables in PowerShell:

```powershell
$env:FOUNDRY_PROJECT_ENDPOINT = "https://<resource>.services.ai.azure.com/api/projects/<project>"
$env:CONNECTION_NAME = "<connection-name>"
$env:DEPLOYMENT_NAME = "<deployment-name>"
```

`FOUNDRY_PROJECT_ENDPOINT` is the project endpoint shown in the Foundry portal.

## Install, build, and run

```powershell
npm install
npm run build
npm start
```
