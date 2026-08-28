# Evaluation Report: identity-dp-java-service-principal

**Config:** java-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 326.3s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `identity-dp-java-service-principal` |
| Config | java-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 16/17 |
| Duration | 326.3s |
| Timestamp | 2026-08-27T17:11:31Z |
| Files Generated | 0 |
| Event Count | 3573 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 95.5s |
| Review | 229.6s |
| **Total** | **326.3s** |

## Configuration

- **name:** java-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=122534 out=3027 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. web_search (via github-mcp-server) — 31253ms

**Input:**

```json
{"query":"As of August 2026, what are the current stable Maven versions and official usage examples for com.azure azure-identity ClientSecretCredentialBuilder and com.azure azure-storage-blob BlobServiceClientBuilder? Cite Microsoft Learn and Maven Central."}
```

**Output:**

```
{"type":"output_text","text":{"value":"### 1. Latest Stable Maven Versions (as of August 2026)\n\n**com.azure:azure-identity**  \n- **Latest version:** 1.13.2  \n- **Maven Central:**  \n  ```xml\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-identity</artifactId>\n    <version>1.13.2</version>\n  </dependency>\n  ```\n  Source: [Maven Central azure-identity](https://search.maven.org/artifact/com.azure/azure-identity) \n\n**com.azure:azure-storage-blob**  \n- **Latest version:** 12.35.0  \n- **Maven Central:**  \n  ```xml\n  <dependency>\n    <groupId>com.azure</groupId>\n    <artifactId>azure-storage-blob</artifactId>\n    <version>12.35.0</version>\n  </dependency>\n  ```\n  Source: [Maven Central azure-storage-blob](https://search.maven.org/artifact/com.azure/azure-storage-blob)【9:3†source】\n\n---\n\n### 2. Official Usage Examples\n\n#### a. **ClientSecretCredentialBuilder (com.azure:azure-identity)**\n\n```java\nimport com.azure.identity.ClientSecretCredential;\nimport com.azure.identity.ClientSecretCredentialBuilder;\n\nClientSecretCredential clientSecretCredential = new ClientSecretCredentialBuilder()\n    .clientId(\"<your-client-id>\")\n    .clientSecret(\"<your-client-secret>\")\n    .tenantId(\"<your-tenant-id>\")\n    .build();\n```\n- Microsoft Learn official documentation: [Authenticate with Azure Identity client library for Java](https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable#authenticate-with-client-secret) \n\n---\n\n#### b. **BlobServiceClientBuilder (com.azure:azure-storage-blob)**\n\n```java\nimport com.azure.identity.ClientSecretCredential;\nimport com.azure.identity.ClientSecretCredentialBuilder;\nimport com.azure.storage.blob.BlobServiceClient;\nimport com.azure.storage.blob.BlobServiceClientBuilder;\n\nClientSecretCredential credential = new ClientSecretCredentialBuilder()\n    .clientId(\"<your-client-id>\")\n    .clientSecret(\"<your-client-secret>\")\n    .tena
... (truncated)
```

### ✅ 2. web_fetch — 6782ms

**Input:**

```json
{"max_length":3000,"raw":true,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-java:
---
layout: Conceptual
title: 'Quickstart: Azure Blob Storage library - Java | Microsoft Learn'
canonicalUrl: https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-java
breadcrumb_path: /azure/bread/toc.json
feedback_help_link_url: https://learn.microsoft.com/answers/tags/125/azure-blob-storage/
feedback_help_link_type: get-help-at-qna
feedback_product_url: https://feedback.azure.com/d365community/forum/a8bb4a47-3525-ec11-b6e6-000d3a4f0f84
feedback_system: Standard
permissioned-type: public
recommendations: true
recommendation_types:
- Training
- Certification
uhfHeaderId: azure
ms.suite: office
adobe-target: true
zone_pivot_group_filename: zone-pivot-groups.json
learn_banner_products:
- azure
manager: akashdubey
ms.reviewer: akashdubey-ms
description: In this quickstart, you learn how to use the Azure Blob Storage client library for Java to create a container and a blob in Blob (object) storage. Next, you learn how to download the blob to your local computer, and how to list all of the blobs in a container.
author: stevenmatthew
ms.author: shaas
ms.date: 2024-09-13T00:00:00.0000000Z
ms.service: azure-blob-storage
ms.topic: quickstart
ms.devlang: java
zone_pivot_groups: azure-blob-storage-quickstart-options
ms.custom:
- devx-track-java
- mode-api
- passwordless-java
- devx-track-extended-java
- devx-track-extended-azdevcli
- sfi-image-nochange
- sfi-ropc-nochange
locale: en-us
document_id: d41d3443-0873-2f4c-da3e-9086663af3ac
document_version_independent_id: 26bf67f1-5847-904b-7cec-c6aecf9061e2
updated_at: 2026-05-24T11:44:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-pr/blob/live/articles/storage/blobs/storage-quickstart-blobs-java.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-pr/blob/e2ff1019f34ee74c3eb408b8d2e0a25f22f6ffbd/articles/storage/blobs/storage-quickstart-blobs-java.md
git_commit_id: e2f
... (truncated)
```

### ✅ 3. web_fetch — 6440ms

**Input:**

```json
{"max_length":3000,"raw":true,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable:
---
layout: Conceptual
monikers:
- azure-java-stable
defaultMoniker: azure-java-stable
versioningType: Ranged
title: Azure Identity client library for Java | Microsoft Learn
canonicalUrl: https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable
config_moniker_range: azure-java-stable
uhfHeaderId: MSDocsHeader-Java
feedback_system: OpenSource
feedback_product_url: https://github.com/Azure/azure-sdk-for-java/issues
ms.topic: reference
ms.devlang: java
ms.date: 2026-08-24T00:00:00.0000000Z
breadcrumb_path: /java/breadcrumb/toc.json
apiPlatform: java
ms.author: azsdkdocs
author: azure-sdk
products:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
keywords: Azure, java, SDK, API, azure-identity, entra-id
ms.service: entra-id
locale: en-us
document_id: a0c04cc9-4127-ee61-109d-2d672636a46e
document_version_independent_id: 0e5b5560-efe9-0c0a-9a63-38dc58e6af8b
updated_at: 2026-08-25T11:08:00.0000000Z
original_content_git_url: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/live/docs-ref-services/latest/identity-readme.md
gitcommit: https://github.com/MicrosoftDocs/azure-docs-sdk-java/blob/75b1775676c4b42b2562bb1c23a9a8a867297328/docs-ref-services/latest/identity-readme.md
git_commit_id: 75b1775676c4b42b2562bb1c23a9a8a867297328
default_moniker: azure-java-stable
site_name: Docs
depot_name: Azure.AzureJava
page_type: conceptual
toc_rel: ../../azure/toc.json
feedback_help_link_type: ''
feedback_help_link_url: ''
word_count: 2370
asset_id: api/overview/azure/identity-readme
moniker_range_name: 39dd779fb9abbd7cd992d58b73780133
monikers:
- azure-java-stable
item_type: Content
source_path: docs-ref-services/latest/identity-readme.md
cmProducts:
- https://authoring-docs-microsoft.poolparty.biz/devrel/68ec7f3a-2bc6-459f-b959-19beb729907d
spProducts:
- https://authoring-docs-microsoft
... (truncated)
```

### ✅ 4. web_fetch — 6111ms

**Input:**

```json
{"max_length":5000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable"}
```

**Output:**

```
Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-storage-blob</artifactId>
  <versioning>
    <latest>12.36.0-beta.1</latest>
    <release>12.36.0-beta.1</release>
    <versions>
      <version>12.7.0</version>
      <version>12.8.0-beta.1</version>
      <version>12.8.0</version>
      <version>12.9.0-beta.1</version>
      <version>12.9.0-beta.2</version>
      <version>12.9.0</version>
      <version>12.10.0-beta.1</version>
      <version>12.10.0</version>
      <version>12.10.1</version>
      <version>12.10.2</version>
      <version>12.11.0-beta.1</version>
      <version>12.11.0-beta.2</version>
      <version>12.11.0-beta.3</version>
      <version>12.11.0</version>
      <version>12.11.1</version>
      <version>12.12.0-beta.1</version>
      <version>12.12.0</version>
      <version>12.13.0</version>
      <version>12.14.0-beta.1</version>
      <version>12.14.0</version>
      <version>12.14.1</version>
      <version>12.14.2</version>
      <version>12.14.3</version>
      <version>12.14.4</version>
      <version>12.15.0-beta.1</version>
      <version>12.15.0-beta.2</version>
      <version>12.15.0-beta.3</version>
      <version>12.15.0</version>
      <version>12.16.0</version>
      <version>12.16.1</version>
      <version>12.17.0-beta.1</version>
      <version>12.17.0</version>
      <version>12.17.1</version>
      <version>12.18.0-beta.1</version>
      <version>12.18.0</version>
      <version>12.19.0</version>
      <version>12.19.1</version>
      <version>12.20.0-beta.1</version>
      <version>12.20.0</version>
      <version>12.20.1</version>
      <version>12.20.2</version>
      <version>12.20.3</version>
      <version>12.21.0-beta.1</version>
      <version>12.21.0</version>
      <version>12.21.1</version>
      <version>12.22.0-beta.1</version>
      <version>12.22.
... (truncated)
```

### ✅ 5. web_fetch — 5771ms

**Input:**

```json
{"max_length":4000,"url":"https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-java"}
```

**Output:**

```
Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<?xml version="1.0" encoding="UTF-8"?>
<metadata>
  <groupId>com.azure</groupId>
  <artifactId>azure-identity</artifactId>
  <versioning>
    <latest>1.19.0-beta.2</latest>
    <release>1.19.0-beta.2</release>
    <versions>
      <version>1.0.0</version>
      <version>1.0.0-preview.1</version>
      <version>1.0.0-preview.2</version>
      <version>1.0.0-preview.3</version>
      <version>1.0.0-preview.4</version>
      <version>1.0.0-preview.5</version>
      <version>1.0.1</version>
      <version>1.0.2</version>
      <version>1.0.3</version>
      <version>1.0.4</version>
      <version>1.0.5</version>
      <version>1.0.6</version>
      <version>1.0.7</version>
      <version>1.0.8</version>
      <version>1.0.9</version>
      <version>1.1.0-beta.1</version>
      <version>1.1.0-beta.2</version>
      <version>1.1.0-beta.3</version>
      <version>1.1.0-beta.4</version>
      <version>1.1.0-beta.5</version>
      <version>1.1.0-beta.6</version>
      <version>1.1.0-beta.7</version>
      <version>1.1.0</version>
      <version>1.1.1</version>
      <version>1.1.2</version>
      <version>1.1.3</version>
      <version>1.2.0-beta.1</version>
      <version>1.2.0-beta.2</version>
      <version>1.2.0</version>
      <version>1.2.1</version>
      <version>1.2.2</version>
      <version>1.2.3</version>
      <version>1.2.4</version>
      <version>1.2.5</version>
      <version>1.3.0-beta.1</version>
      <version>1.3.0-beta.2</version>
      <version>1.3.0</version>
      <version>1.3.1</version>
      <version>1.3.2</version>
      <version>1.3.3</version>
      <version>1.3.4</version>
      <version>1.3.5</version>
      <version>1.3.6</version>
      <version>1.3.7</version>
      <version>1.4.0-beta.1</version>
      <version>1.4.0</version>
      <version>1.4.1</version>
      <version>1.4.2</version>
      <version>1.4.3</version>
      <ve
... (truncated)
```

### ✅ 6. web_fetch — 2589ms

**Input:**

```json
{"max_length":3000,"raw":true,"start_index":6000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Contents of https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable:
or an Azure resource in [Enable managed identity for Azure resources](/en-us/azure/developer/java/sdk/identity-azure-hosted-auth#managed-identity-credential).

```java
/**
 * DefaultAzureCredential uses the user-assigned managed identity with the specified client ID.
 */
public void createDefaultAzureCredentialForUserAssignedManagedIdentity() {
    DefaultAzureCredential defaultCredential = new DefaultAzureCredentialBuilder()
        .managedIdentityClientId("<MANAGED_IDENTITY_CLIENT_ID>")
        .build();

    // Azure SDK client builders accept the credential as a parameter
    SecretClient client = new SecretClientBuilder()
        .vaultUrl("https://{YOUR_VAULT_NAME}.vault.azure.net")
        .credential(defaultCredential)
        .buildClient();
}
```

In addition to configuring the `managedIdentityClientId` via code, it can also be set using the `AZURE_CLIENT_ID` environment variable. These two approaches are equivalent when using `DefaultAzureCredential`.

### Authenticate a user in Azure Toolkit for IntelliJ with `DefaultAzureCredential`

To authenticate using IntelliJ, ensure that configuration instructions [here](/en-us/azure/developer/java/sdk/identity-dev-env-auth#sign-in-azure-toolkit-for-intellij-for-intellijcredential) have been successfully completed.

The below example demonstrates authenticating the `SecretClient` from the [azure-security-keyvault-secrets](https://github.com/Azure/azure-sdk-for-java/tree/com.azure+azure-identity_1.18.5/sdk/keyvault/azure-security-keyvault-secrets) client library using `DefaultAzureCredential`, on a workstation with IntelliJ IDEA installed, and the user has signed in with an Azure account to the Azure Toolkit for IntelliJ.

See more about how to configure your IntelliJ IDEA in [Sign in Azure Toolkit for IntelliJ for IntelliJCredential](/en-us/azure/developer/java/sdk/identity-dev-
... (truncated)
```

### ✅ 7. web_fetch — 2257ms

**Input:**

```json
{"max_length":3000,"raw":true,"start_index":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml"}
```

**Output:**

```
Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
<error>No more content available.</error>
```

### ✅ 8. web_fetch — 1917ms

**Input:**

```json
{"max_length":5000,"start_index":9000,"url":"https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable"}
```

**Output:**

```
Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-storage-blob/maven-metadata.xml:
1.0</version>
      <version>12.31.1</version>
      <version>12.31.2</version>
      <version>12.31.3</version>
      <version>12.32.0-beta.1</version>
      <version>12.32.0</version>
      <version>12.33.0-beta.1</version>
      <version>12.33.0</version>
      <version>12.33.1</version>
      <version>12.33.2</version>
      <version>12.33.3</version>
      <version>12.33.4</version>
      <version>12.34.0-beta.1</version>
      <version>12.34.0</version>
      <version>12.35.0-beta.1</version>
      <version>12.35.0</version>
      <version>12.35.1</version>
      <version>12.36.0-beta.1</version>
    </versions>
    <lastUpdated>20260818232411</lastUpdated>
  </versioning>
</metadata>

```

### ✅ 9. web_fetch — 1052ms

**Input:**

```json
{"max_length":3000,"raw":true,"start_index":3000,"url":"https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml"}
```

**Output:**

```
Here is the raw content:
Contents of https://repo1.maven.org/maven2/com/azure/azure-identity/maven-metadata.xml:
ion>
      <version>1.10.2</version>
      <version>1.10.3</version>
      <version>1.10.4</version>
      <version>1.11.0-beta.1</version>
      <version>1.11.0-beta.2</version>
      <version>1.11.0</version>
      <version>1.11.1</version>
      <version>1.11.2</version>
      <version>1.11.3</version>
      <version>1.11.4</version>
      <version>1.12.0-beta.1</version>
      <version>1.12.0</version>
      <version>1.12.1</version>
      <version>1.12.2</version>
      <version>1.13.0-beta.1</version>
      <version>1.13.0-beta.2</version>
      <version>1.13.0</version>
      <version>1.13.1</version>
      <version>1.13.2</version>
      <version>1.13.3</version>
      <version>1.14.0-beta.1</version>
      <version>1.14.0-beta.2</version>
      <version>1.14.0</version>
      <version>1.14.1</version>
      <version>1.14.2</version>
      <version>1.15.0</version>
      <version>1.15.1</version>
      <version>1.15.2</version>
      <version>1.15.3</version>
      <version>1.15.4</version>
      <version>1.16.0-beta.1</version>
      <version>1.16.0</version>
      <version>1.16.1</version>
      <version>1.16.2</version>
      <version>1.16.3</version>
      <version>1.17.0-beta.1</version>
      <version>1.17.0</version>
      <version>1.18.0</version>
      <version>1.18.1</version>
      <version>1.18.2</version>
      <version>1.18.3</version>
      <version>1.18.4</version>
      <version>1.18.5</version>
      <version>1.19.0-beta.1</version>
      <version>1.19.0-beta.2</version>
    </versions>
    <lastUpdated>20260824185316</lastUpdated>
  </versioning>
</metadata>

```

## Copilot Response

Below is a complete example using a service principal to authenticate to Azure Blob Storage. The same `ClientSecretCredential` can be passed to other Azure SDK client builders that accept a `TokenCredential`.

## 1. Maven dependencies

```xml
<project xmlns="http://maven.apache.org/POM/4.0.0"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://maven.apache.org/POM/4.0.0
                             https://maven.apache.org/xsd/maven-4.0.0.xsd">
    <modelVersion>4.0.0</modelVersion>

    <groupId>com.example</groupId>
    <artifactId>azure-service-principal-example</artifactId>
    <version>1.0.0</version>

    <properties>
        <maven.compiler.release>17</maven.compiler.release>
        <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
    </properties>

    <dependencies>
        <!-- Microsoft Entra ID credentials for Azure SDK clients -->
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-identity</artifactId>
            <version>1.18.5</version>
        </dependency>

        <!-- Example Azure SDK client -->
        <dependency>
            <groupId>com.azure</groupId>
            <artifactId>azure-storage-blob</artifactId>
            <version>12.35.1</version>
        </dependency>
    </dependencies>
</project>
```

These are the latest stable releases shown in Maven Central as of August 2026. For larger applications, prefer the [Azure SDK BOM](https://github.com/Azure/azure-sdk-for-java/tree/main/sdk/boms/azure-sdk-bom) to keep Azure SDK dependencies compatible.

## 2. Configuration

Supply the credentials and service endpoint through environment variables:

```text
AZURE_TENANT_ID=<Microsoft-Entra-tenant-ID>
AZURE_CLIENT_ID=<service-principal-application-client-ID>
AZURE_CLIENT_SECRET=<client-secret-value>
AZURE_STORAGE_BLOB_ENDPOINT=https://<account-name>.blob.core.windows.net
```

`AZURE_CLIENT_SECRET` must contain the **secret value**, not the secret's identifier.

The service principal must also have an appropriate data-plane role on the storage account, such as **Storage Blob Data Reader** for this read-only example.

## 3. Complete Java example

Save as `src/main/java/com/example/AzureBlobExample.java`:

```java
package com.example;

import com.azure.core.exception.AzureException;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.ClientSecretCredential;
import com.azure.identity.ClientSecretCredentialBuilder;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobContainerItem;
import com.azure.storage.blob.models.BlobStorageException;

public final class AzureBlobExample {
    private AzureBlobExample() {
    }

    public static void main(String[] args) {
        int exitCode = run();

        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }

    private static int run() {
        try {
            String tenantId = requiredEnvironmentVariable("AZURE_TENANT_ID");
            String clientId = requiredEnvironmentVariable("AZURE_CLIENT_ID");
            String clientSecret =
                    requiredEnvironmentVariable("AZURE_CLIENT_SECRET");
            String blobEndpoint =
                    requiredEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");

            ClientSecretCredential credential =
                    new ClientSecretCredentialBuilder()
                            .tenantId(tenantId)
                            .clientId(clientId)
                            .clientSecret(clientSecret)
                            .build();

            BlobServiceClient blobServiceClient =
                    new BlobServiceClientBuilder()
                            .endpoint(blobEndpoint)
                            .credential(credential)
                            .buildClient();

            /*
             * Building the credential and client does not authenticate.
             * Authentication occurs when the first Azure operation is made.
             */
            System.out.println("Containers:");

            for (BlobContainerItem container
                    : blobServiceClient.listBlobContainers()) {
                System.out.println("- " + container.getName());
            }

            return 0;
        } catch (IllegalStateException exception) {
            // Local configuration problem, such as a missing variable.
            System.err.println("Configuration error: "
                    + exception.getMessage());
            return 2;
        } catch (ClientAuthenticationException exception) {
            /*
             * Common causes:
             * - Incorrect tenant ID, client ID, or secret value
             * - Expired or disabled secret
             * - Deleted or disabled service principal
             */
            System.err.println(
                    "Azure authentication failed. Verify the tenant ID, "
                            + "client ID, and client-secret value. Also check "
                            + "whether the secret has expired.");
            return 3;
        } catch (BlobStorageException exception) {
            /*
             * Authentication may have succeeded while authorization or the
             * storage request failed.
             */
            if (exception.getStatusCode() == 403) {
                System.err.println(
                        "Azure authorization failed. The service principal "
                                + "does not have the required Blob Storage "
                                + "data-plane role.");
            } else {
                System.err.printf(
                        "Blob Storage request failed: HTTP %d, error code %s%n",
                        exception.getStatusCode(),
                        exception.getErrorCode());
            }

            return 4;
        } catch (AzureException exception) {
            // Covers other Azure SDK failures, including transport failures.
            System.err.println(
                    "The Azure SDK operation failed: "
                            + exception.getClass().getSimpleName());
            return 5;
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);

        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    "Required environment variable " + name + " is not set");
        }

        return value;
    }
}
```

An incorrect client secret is normally reported as `ClientAuthenticationException` when `listBlobContainers()` requests its first access token. A valid identity without the necessary Azure role usually receives HTTP `403`, represented here by `BlobStorageException`.

## 4. Secret-management best practices

- Never embed secrets in Java source, `pom.xml`, container images, logs, or committed configuration files.
- Inject secrets at runtime through a protected secret store, CI/CD secret facility, or environment variable.
- Prefer **managed identity** on Azure-hosted workloads and **workload identity federation** in CI/CD or Kubernetes. These avoid long-lived client secrets entirely.
- If a secret is unavoidable, use a short expiration, rotate it regularly, and support overlapping secrets during rotation.
- Grant only the required roles at the narrowest practical scope. For example, use **Storage Blob Data Reader** on one storage account or container instead of broad subscription-level access.
- Do not log the secret or include it in exception messages. Restrict access to process environments and diagnostic dumps.
- Store production secrets in Azure Key Vault or an equivalent secret manager. Prefer accessing the vault through managed/workload identity rather than using another embedded secret to retrieve the first one.

References: [Azure Identity for Java](https://learn.microsoft.com/en-us/java/api/overview/azure/identity-readme?view=azure-java-stable), [Blob Storage client library for Java](https://learn.microsoft.com/en-us/azure/storage/blobs/storage-quickstart-blobs-java), and [ClientSecretCredentialBuilder API](https://learn.microsoft.com/en-us/java/api/com.azure.identity.clientsecretcredentialbuilder?view=azure-java-stable).

## Reviewer Notes (LLM-as-Judge)

### Summary

The response satisfies the client resource-management criterion.

### Strengths

- Correctly uses BlobServiceClient without unnecessary resource-closing logic.

## Grader Results

- service-principal-auth.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (5/5)
      - Maven dependency for `azure-identity`: Pass
      - `ClientSecretCredentialBuilder` with tenantId, clientId, clientSecret: Pass
      - Passing credential to Azure SDK client builders: Pass
      - Environment variable patterns for secret storage: Pass
      - `AuthenticationException` handling: Pass
- java.yaml (criteria file):
  - Correct Dependencies (com.azure, not com.microsoft.azure) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Dependencies (com.azure, not com.microsoft.azure)**: Uses com.azure group ID for all Azure SDK packages. No com.microsoft.azure (legacy SDK) dependencies. Includes azure-identity for authentication.: Pass
  - Azure SDK BOM for Version Management (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Azure SDK BOM for Version Management**: Uses azure-sdk-bom in dependencyManagement to manage Azure SDK versions instead of hardcoding individual artifact versions. Dependencies should omit <version> tags when managed by the BOM.: Fail
  - Correct Imports (no legacy, no internal packages) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Imports (no legacy, no internal packages)**: All imports use com.azure.* packages. No com.microsoft.azure.* (legacy) or com.azure.*.implementation.* (internal API) imports.: Pass
  - DefaultAzureCredential Authentication (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Authentication**: Uses DefaultAzureCredential or another com.azure.identity credential. No hardcoded connection strings, account keys, SAS tokens, or secrets.: Pass
  - Client Builder Pattern (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Client Builder Pattern**: SDK clients constructed using *ClientBuilder classes with .endpoint() or .vaultUrl() and .credential(). No legacy constructors (CloudStorageAccount, DocumentClient, KeyVaultClient).: Pass
  - No Deprecated/Legacy Classes (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**No Deprecated/Legacy Classes**: No deprecated classes from the old SDK (CloudStorageAccount, CloudBlobClient, DocumentClient, QueueClient, ApplicationTokenCredentials, MSICredentials, ConnectionStringBuilder).: Pass
  - Pagination (PagedIterable/PagedFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Pagination (PagedIterable/PagedFlux)**: List/query operations return PagedIterable (sync) or PagedFlux (async). Does not flatten all pages into a raw List or Stream in memory.: Pass
  - LRO Pattern (SyncPoller/PollerFlux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**LRO Pattern (SyncPoller/PollerFlux)**: Long-running operations use SyncPoller (sync) or PollerFlux (async) with begin* method prefix. No Thread.sleep() polling loops.: Pass
  - Async Uses Project Reactor (Mono/Flux) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Uses Project Reactor (Mono/Flux)**: Async code uses Project Reactor types (Mono, Flux). Not CompletableFuture (wrong), not RxJava (wrong), not sync wrapped in ExecutorService. No .block() inside async service implementations.: Pass
  - Service-Specific Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Service-Specific Exception Handling**: Catches service-specific exceptions (BlobStorageException, CosmosException, ServiceBusException, HttpResponseException) with status code inspection. Not just generic Exception catches.: Pass
  - Code Compiles (mvn compile / gradle compileJava) (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Code Compiles (mvn compile / gradle compileJava)**: The generated code compiles without errors. Attempt build verification if build tools are available.: Pass
  - Try-With-Resources for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Try-With-Resources for Clients**: All Azure SDK client instances that implement AutoCloseable are used within try-with-resources blocks or explicitly closed in a finally block.: Pass

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Correct Dependencies (com.azure, not com.microsoft.azure)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Azure SDK BOM for Version Management` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Correct Imports (no legacy, no internal packages)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `DefaultAzureCredential Authentication` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Client Builder Pattern` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `No Deprecated/Legacy Classes` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Pagination (PagedIterable/PagedFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `LRO Pattern (SyncPoller/PollerFlux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Async Uses Project Reactor (Mono/Flux)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Service-Specific Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Code Compiles (mvn compile / gradle compileJava)` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| `Try-With-Resources for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 8.3% | ✅ |
| **Final** | | | **Σ 13.00** | **Σ 12.0000** | **92.3%** | |

## Re-run Command

```bash
hyoka run --prompt-id identity-dp-java-service-principal --config java-azure-skills/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
