package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.storage.StorageManager;
import com.azure.resourcemanager.storage.models.BlobServiceProperties;
import com.azure.resourcemanager.storage.models.StorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountSkuType;

public final class StorageAccountManager {
    private static final String REGION = "eastus";

    private StorageAccountManager() {
    }

    public static void main(String[] args) {
        String storageAccountName = null;
        String resourceGroupName = null;
        StorageManager storageManager = null;
        boolean accountCreated = false;
        int exitCode = 0;

        try {
            String subscriptionId = requiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
            resourceGroupName = requiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
            storageAccountName = requiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
            AzureProfile profile = new AzureProfile(
                System.getenv("AZURE_TENANT_ID"),
                subscriptionId,
                AzureEnvironment.AZURE);
            storageManager = StorageManager.authenticate(credential, profile);

            StorageAccount createdAccount = storageManager.storageAccounts()
                .define(storageAccountName)
                .withRegion(REGION)
                .withExistingResourceGroup(resourceGroupName)
                .withSku(StorageAccountSkuType.STANDARD_LRS)
                .create();
            accountCreated = true;
            System.out.printf("Created storage account: %s%n", createdAccount.id());

            System.out.printf("Storage accounts in resource group '%s':%n", resourceGroupName);
            for (StorageAccount account
                : storageManager.storageAccounts().listByResourceGroup(resourceGroupName)) {
                System.out.printf("  %s (%s, %s)%n",
                    account.name(), account.regionName(), account.skuType());
            }

            StorageAccount account = storageManager.storageAccounts()
                .getByResourceGroup(resourceGroupName, storageAccountName);
            if (account == null) {
                throw new IllegalStateException(
                    "Created storage account could not be retrieved: " + storageAccountName);
            }
            System.out.printf(
                "Properties: id=%s, region=%s, sku=%s, kind=%s, status=%s%n",
                account.id(),
                account.regionName(),
                account.skuType(),
                account.kind(),
                account.accountStatuses());

            BlobServiceProperties blobProperties = storageManager.blobServices()
                .getServicePropertiesAsync(resourceGroupName, storageAccountName)
                .block();
            if (blobProperties == null) {
                throw new IllegalStateException(
                    "Blob service properties could not be retrieved: " + storageAccountName);
            }
            blobProperties.update()
                .withBlobVersioningEnabled()
                .apply();
            System.out.printf("Enabled blob versioning for: %s%n", storageAccountName);

            storageManager.storageAccounts()
                .deleteByResourceGroup(resourceGroupName, storageAccountName);
            accountCreated = false;
            System.out.printf("Deleted storage account: %s%n", storageAccountName);
        } catch (CredentialUnavailableException exception) {
            System.err.println("No Azure credential was available: " + exception.getMessage());
            exitCode = 2;
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            exitCode = 2;
        } catch (IllegalArgumentException exception) {
            System.err.println("Invalid configuration: " + exception.getMessage());
            exitCode = 4;
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure management request failed (HTTP %d): %s%n",
                statusCode,
                exception.getMessage());
            exitCode = 3;
        } catch (RuntimeException exception) {
            System.err.println("Storage account operation failed: " + exception.getMessage());
            exitCode = 1;
        } finally {
            if (accountCreated
                && storageManager != null
                && resourceGroupName != null
                && storageAccountName != null) {
                try {
                    storageManager.storageAccounts()
                        .deleteByResourceGroup(resourceGroupName, storageAccountName);
                    System.err.printf(
                        "Deleted storage account during error cleanup: %s%n",
                        storageAccountName);
                } catch (RuntimeException cleanupException) {
                    System.err.printf(
                        "Cleanup failed; storage account '%s' may still exist: %s%n",
                        storageAccountName,
                        cleanupException.getMessage());
                }
            }
        }

        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }
}
