# Evaluation Report: key-vault-dp-dotnet-crud

**Config:** dotnet-azure-tools/with-azure-tools | **Result:** ✅ PASSED | **Duration:** 170.6s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `key-vault-dp-dotnet-crud` |
| Config | dotnet-azure-tools/with-azure-tools |
| Result | ✅ PASSED |
| Score | 5/5 |
| Duration | 170.6s |
| Timestamp | 2026-08-26T17:48:25Z |
| Files Generated | 3 |
| Event Count | 3745 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 142.9s |
| Review | 26.6s |
| **Total** | **170.6s** |

## Configuration

- **name:** dotnet-azure-tools/with-azure-tools
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | azure-ai-agents-persistent-dotnet, azure-ai-document-intelligence-dotnet, azure-ai-openai-dotnet, azure-ai-projects-dotnet, azure-ai-voicelive-dotnet, azure-eventgrid-dotnet, azure-eventhub-dotnet, azure-identity-dotnet, azure-maps-search-dotnet, azure-mgmt-apicenter-dotnet, azure-mgmt-apimanagement-dotnet, azure-mgmt-applicationinsights-dotnet, azure-mgmt-arizeaiobservabilityeval-dotnet, azure-mgmt-botservice-dotnet, azure-mgmt-fabric-dotnet, azure-mgmt-mongodbatlas-dotnet, azure-mgmt-weightsandbiases-dotnet, azure-resource-manager-cosmosdb-dotnet, azure-resource-manager-durabletask-dotnet, azure-resource-manager-mysql-dotnet, azure-resource-manager-playwright-dotnet, azure-resource-manager-postgresql-dotnet, azure-resource-manager-redis-dotnet, azure-resource-manager-sql-dotnet, azure-search-documents-dotnet, azure-security-keyvault-keys-dotnet, azure-servicebus-dotnet, m365-agents-dotnet, microsoft-azure-webjobs-extensions-authentication-events-dotnet, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-identity-dotnet |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=569964 out=3784 |
| Turn Count | 12 |

## Prompt Sent

```
Write a C# console application that performs
all four CRUD operations on Azure Key Vault secrets:
1. Create a new secret called "my-secret" with value "my-secret-value"
2. Read the secret back and print its value
3. Update the secret to a new value "updated-value"
4. Delete the secret and purge it (soft-delete enabled vault)

Use DefaultAzureCredential for authentication. Include proper error handling
and show required NuGet packages.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-identity-dotnet"}
```

### 🔧 2. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_best_practices","intent":"Create a local C# console application using Azure.Identity DefaultAzureCredential and Azure.Security.KeyVault.Secrets SecretClient to create, read, update, delete, and purge a Key Vault secret, with proper error handling. Do not provision or modify Azure resources during validation.","parameters":{"language":"csharp","operation":"code-generation","resource":"key-vault"}}
```

### 🔧 3. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\acceptance-criteria.md"}
```

### 🔧 4. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet\\examples.md"}
```

### 🔧 5. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 6. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Get Azure C# code-generation best practices for a local console sample using DefaultAzureCredential and Key Vault SecretClient CRUD operations.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 7. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find current official .NET Azure Key Vault Secrets SDK guidance for SecretClient SetSecretAsync, GetSecretAsync, StartDeleteSecretAsync, WaitForCompletionAsync, and PurgeDeletedSecretAsync, including exception handling and required NuGet packages.","learn":true}
```

### 🔧 8. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-dotnet\\skills\\azure-identity-dotnet"}
```

### 🔧 9. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Search current official documentation for Azure Key Vault secrets CRUD and purge with the .NET SDK.","parameters":{"query":"Azure Key Vault Secrets .NET SecretClient SetSecretAsync GetSecretAsync StartDeleteSecretAsync WaitForCompletionAsync PurgeDeletedSecretAsync DefaultAzureCredential"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find latest official C# code samples for Azure Key Vault secret CRUD and purge.","parameters":{"language":"csharp","query":"Azure.Security.KeyVault.Secrets SecretClient create read update delete purge deleted secret WaitForCompletionAsync"}}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"dotnet --version","description":"Check installed .NET SDK"}
```

### 🔧 12. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766574355-copilot-tool-output-8d5c01c34fcd4d3d8e9ae3a7460d24e5.txt","view_range":[1,40]}
```

### 🔧 13. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787766574511-copilot-tool-output-a36e9c94bfa44bcb82434143f89efcca.txt","view_range":[1,25]}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"dotnet new console --framework net8.0 --use-program-main --name KeyVaultSecretsCrud --output . --force \u0026\u0026 dotnet add package Azure.Identity \u0026\u0026 dotnet add package Azure.Security.KeyVault.Secrets","description":"Create console project and add Azure SDK packages","initial_wait":180,"mode":"sync"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"dotnet new console --framework net10.0 --use-program-main --name KeyVaultSecretsCrud --output . --force \u0026\u0026 dotnet add package Azure.Identity \u0026\u0026 dotnet add package Azure.Security.KeyVault.Secrets","description":"Create console project and add Azure SDK packages","initial_wait":180,"mode":"sync"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3340719505\\Program.cs"}
```

### 🔧 17. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-key-vault-dp-dotnet-crud-dotnet-azure-tools-with-azure-tools-3340719505\\KeyVaultSecretsCrud.csproj"}
```

### 🔧 18. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: Program.cs\n@@\n-﻿namespace KeyVaultSecretsCrud;\n-\n-class Program\n-{\n-    static void Main(string[] args)\n-    {\n-        Console.WriteLine(\"Hello, World!\");\n-    }\n-}\n+﻿using Azure;\n+using Azure.Core;\n+using Azure.Identity;\n+using Azure.Security.KeyVault.Secrets;\n+\n+namespace KeyVaultSecretsCrud;\n+\n+internal static class Program\n+{\n+    private const string SecretName = \"my-secret\";\n+    private const string InitialValue = \"my-secret-value\";\n+    private const string UpdatedValue = \"updated-value\";\n+\n+    private static async Task\u003cint\u003e Main()\n+    {\n+        string? vaultUrl = Environment.GetEnvironmentVariable(\"KEY_VAULT_URL\");\n+        if (!Uri.TryCreate(vaultUrl, UriKind.Absolute, out Uri? vaultUri)\n+            || vaultUri.Scheme != Uri.UriSchemeHttps)\n+        {\n+            Console.Error.WriteLine(\n+                \"Set KEY_VAULT_URL to a valid HTTPS URI, for example \" +\n+                \"https://my-vault.vault.azure.net/.\");\n+            return 1;\n+        }\n+\n+        using var cancellationSource = new CancellationTokenSource();\n+        Console.CancelKeyPress += (_, eventArgs) =\u003e\n+        {\n+            eventArgs.Cancel = true;\n+            cancellationSource.Cancel();\n+        };\n+\n+        var options = new SecretClientOptions\n+        {\n+            Retry =\n+            {\n+                Mode = RetryMode.Exponential,\n+                Delay = TimeSpan.FromSeconds(1),\n+                MaxDelay = TimeSpan.FromSeconds(8),\n+                MaxRetries = 5\n+            }\n+        };\n+\n+        var client = new SecretClient(\n+            vaultUri,\n+            new DefaultAzureCredential(),\n+            options);\n+\n+        try\n+        {\n+            Console.WriteLine($\"Creating secret '{SecretName}'...\");\n+            await client.SetSecretAsync(\n+                SecretName,\n+                InitialValue,\n+                cancellationSource.Token);\n+\n+            KeyVaultSecret secret = await client.GetSecretAsync(\n+                SecretName,\n+                cancellationToken: cancellationSource.Token);\n+            Console.WriteLine($\"Read secret value: {secret.Value}\");\n+\n+            Console.WriteLine($\"Updating secret '{SecretName}'...\");\n+            await client.SetSecretAsync(\n+                SecretName,\n+                UpdatedValue,\n+                cancellationSource.Token);\n+\n+            KeyVaultSecret updatedSecret = await client.GetSecretAsync(\n+                SecretName,\n+                cancellationToken: cancellationSource.Token);\n+            Console.WriteLine($\"Updated secret value: {updatedSecret.Value}\");\n+\n+            Console.WriteLine($\"Deleting secret '{SecretName}'...\");\n+            DeleteSecretOperation deleteOperation =\n+                await client.StartDeleteSecretAsync(\n+                    SecretName,\n+                    cancellationSource.Token);\n+\n+            // Purge is valid only after the soft-delete operation completes.\n+            await deleteOperation.WaitForCompletionAsync(cancellationSource.Token);\n+\n+            Console.WriteLine($\"Purging secret '{SecretName}'...\");\n+            await client.PurgeDeletedSecretAsync(\n+                SecretName,\n+                cancellationSource.Token);\n+\n+            Console.WriteLine(\"CRUD operations completed successfully.\");\n+            return 0;\n+        }\n+        catch (CredentialUnavailableException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"No credential was available for DefaultAzureCredential: {exception.Message}\");\n+            return 2;\n+        }\n+        catch (AuthenticationFailedException exception)\n+        {\n+            Console.Error.WriteLine($\"Azure authentication failed: {exception.Message}\");\n+            return 2;\n+        }\n+        catch (RequestFailedException exception)\n+        {\n+            Console.Error.WriteLine(\n+                $\"Key Vault request failed ({exception.Status}, {exception.ErrorCode}): \" +\n+                exception.Message);\n+\n+            if (exception.Status == 403)\n+            {\n+                Console.Error.WriteLine(\n+                    \"Verify that the identity has secret get, set, delete, and purge permissions. \" +\n+                    \"Purge also fails when purge protection is enabled.\");\n+            }\n+\n+            return 3;\n+        }\n+        catch (OperationCanceledException)\n+        {\n+            Console.Error.WriteLine(\"Operation canceled.\");\n+            return 4;\n+        }\n+    }\n+}\n*** Add File: README.md\n+# Azure Key Vault Secrets CRUD\n+\n+This .NET console application uses `DefaultAzureCredential` to create, read,\n+update, delete, and purge the `my-secret` secret.\n+\n+## Required NuGet packages\n+\n+```powershell\n+dotnet add package Azure.Identity\n+dotnet add package Azure.Security.KeyVault.Secrets\n+```\n+\n+The project currently resolves these package versions:\n+\n+- `Azure.Identity` 1.21.0\n+- `Azure.Security.KeyVault.Secrets` 4.11.0\n+\n+## Run\n+\n+Set the vault URL without putting credentials in source code:\n+\n+```powershell\n+$env:KEY_VAULT_URL = \"https://\u003cvault-name\u003e.vault.azure.net/\"\n+dotnet run\n+```\n+\n+`DefaultAzureCredential` can use local developer credentials and automatically\n+use managed identity when hosted in Azure. The authenticated identity needs Key\n+Vault data-plane permissions to get, set, delete, and purge secrets. The vault\n+must have soft delete enabled and purge protection disabled for immediate purge.\n+\n+Updating the value with `SetSecretAsync` creates a new version of the existing\n+secret. The application waits for the delete operation before purging it.\n+\n+Reference: [Azure Key Vault secret client library quickstart for\n+.NET](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-net)\n*** End Patch\n"
```

### 🔧 19. powershell

**Input:**

```json
{"command":"dotnet build --no-restore --configuration Release","description":"Build Key Vault console app","initial_wait":180,"mode":"sync"}
```

### 🔧 20. powershell

**Input:**

```json
{"command":"Remove-Item Env:KEY_VAULT_URL -ErrorAction SilentlyContinue; dotnet run --no-build --configuration Release; if ($LASTEXITCODE -ne 1) { throw \"Expected exit code 1 but got $LASTEXITCODE\" }","description":"Check missing configuration handling","initial_wait":30,"mode":"sync"}
```

### 🔧 21. powershell

**Input:**

```json
{"command":"dotnet format --no-restore --verify-no-changes","description":"Check C# formatting","initial_wait":180,"mode":"sync"}
```

## Generated Files

- `KeyVaultSecretsCrud.csproj`
- `Program.cs`
- `README.md`

## Copilot Response

I’m grounding the sample in the current Azure SDK patterns, then I’ll build and locally validate the console app without contacting Azure.

The official SDK confirms that changing a secret’s value uses `SetSecretAsync` to create a new version, and purge must wait for the soft-delete operation to complete. I’m implementing that sequence with bounded cancellation and service-specific errors.

Created a complete C# console application:

- `Program.cs` — asynchronous create, read, update, delete, and purge operations using `DefaultAzureCredential`
- `KeyVaultSecretsCrud.csproj` — includes:
  - `Azure.Identity` 1.21.0
  - `Azure.Security.KeyVault.Secrets` 4.11.0
- `README.md` — setup, permissions, and execution instructions

Run with:

```powershell
$env:KEY_VAULT_URL = "https://<vault-name>.vault.azure.net/"
dotnet run
```

The authenticated identity needs secret **get**, **set**, **delete**, and **purge** permissions. Immediate purge requires soft delete enabled and purge protection disabled.

Reference: [Azure Key Vault secret client library quickstart for .NET](https://learn.microsoft.com/azure/key-vault/secrets/quick-create-net)

## Reviewer Notes (LLM-as-Judge)

### Summary

All 5 criteria passed. The application fully implements asynchronous Azure Key Vault secret CRUD, waits for soft deletion before purging, uses DefaultAzureCredential, documents the required packages, and handles Azure request failures.

### Strengths

- Complete asynchronous CRUD and purge workflow
- Correct soft-delete completion handling before purge
- Explicit RequestFailedException handling with actionable diagnostics
- Secure vault configuration through an environment variable
- Clear package and permission documentation

## Grader Results

- crud-secrets.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Installing `Azure.Security.KeyVault.Secrets` and `Azure.Identity` NuGet packages: Pass
      - Creating a `SecretClient` with vault URI and credential: Pass
      - `SetSecret()`, `GetSecret()`, `StartDeleteSecret()`, `PurgeDeletedSecret()`: Pass
      - Handling soft-delete (polling `DeleteSecretOperation` to completion before purge): Pass
      - Exception handling for `RequestFailedException`: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 100.0% | ✅ |
| **Final** | | | **Σ 1.00** | **Σ 1.0000** | **100.0%** | |

## Re-run Command

```bash
hyoka run --prompt-id key-vault-dp-dotnet-crud --config dotnet-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
