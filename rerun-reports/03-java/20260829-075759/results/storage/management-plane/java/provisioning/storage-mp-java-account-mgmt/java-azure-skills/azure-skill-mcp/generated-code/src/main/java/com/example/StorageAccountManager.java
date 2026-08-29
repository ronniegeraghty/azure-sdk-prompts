package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.storage.StorageManager;
import com.azure.resourcemanager.storage.models.BlobServiceProperties;
import com.azure.resourcemanager.storage.models.StorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountSkuType;

import java.util.Objects;
import java.util.regex.Pattern;

public final class StorageAccountManager {
    private static final Pattern STORAGE_ACCOUNT_NAME = Pattern.compile("[a-z0-9]{3,24}");

    private StorageAccountManager() {
    }

    public static void main(String[] args) {
        try {
            run();
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure request failed (HTTP %d): %s%n",
                statusCode,
                exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.println("Invalid configuration: " + exception.getMessage());
            System.exit(1);
        } catch (RuntimeException exception) {
            System.err.println("Storage account operation failed: " + exception.getMessage());
            System.exit(1);
        }
    }

    private static void run() {
        String subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
        String resourceGroupName = requireEnvironmentVariable("AZURE_RESOURCE_GROUP");
        String storageAccountName = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

        if (!STORAGE_ACCOUNT_NAME.matcher(storageAccountName).matches()) {
            throw new IllegalArgumentException(
                "AZURE_STORAGE_ACCOUNT_NAME must contain 3-24 lowercase letters or digits.");
        }

        if (!Boolean.parseBoolean(System.getenv("AZURE_ENABLE_RESOURCE_CHANGES"))) {
            System.out.println(
                "Dry run: set AZURE_ENABLE_RESOURCE_CHANGES=true to execute the Azure operations.");
            return;
        }

        var credential = new DefaultAzureCredentialBuilder().build();
        var profile = new AzureProfile(null, subscriptionId, AzureEnvironment.AZURE);
        StorageManager storageManager = StorageManager.authenticate(credential, profile);

        boolean accountCreated = false;
        boolean accountDeleted = false;
        RuntimeException operationFailure = null;
        try {
            StorageAccount createdAccount = storageManager.storageAccounts()
                .define(storageAccountName)
                .withRegion(Region.US_EAST)
                .withExistingResourceGroup(resourceGroupName)
                .withSku(StorageAccountSkuType.STANDARD_LRS)
                .create();
            accountCreated = true;
            System.out.println("Created storage account: " + createdAccount.id());

            System.out.println("Storage accounts in resource group " + resourceGroupName + ":");
            for (StorageAccount account
                : storageManager.storageAccounts().listByResourceGroup(resourceGroupName)) {
                System.out.printf(
                    "- %s (region=%s, sku=%s)%n",
                    account.name(),
                    account.regionName(),
                    account.skuType());
            }

            StorageAccount account = storageManager.storageAccounts()
                .getByResourceGroup(resourceGroupName, storageAccountName);
            if (account == null) {
                throw new IllegalStateException(
                    "The newly created storage account could not be retrieved.");
            }
            System.out.printf(
                "Properties: id=%s, region=%s, kind=%s, sku=%s%n",
                account.id(),
                account.regionName(),
                account.kind(),
                account.skuType());

            BlobServiceProperties blobService = Objects.requireNonNull(
                storageManager.blobServices()
                    .getServicePropertiesAsync(resourceGroupName, storageAccountName)
                    .block(),
                "Blob service properties were not returned.");
            BlobServiceProperties updatedBlobService = blobService.update()
                .withBlobVersioningEnabled()
                .apply();
            System.out.println(
                "Blob versioning enabled: " + updatedBlobService.isBlobVersioningEnabled());

            storageManager.storageAccounts()
                .deleteByResourceGroup(resourceGroupName, storageAccountName);
            accountDeleted = true;
            System.out.println("Deleted storage account: " + storageAccountName);
        } catch (RuntimeException exception) {
            operationFailure = exception;
            throw exception;
        } finally {
            if (accountCreated && !accountDeleted) {
                try {
                    storageManager.storageAccounts()
                        .deleteByResourceGroup(resourceGroupName, storageAccountName);
                    System.out.println(
                        "Deleted storage account during failure cleanup: " + storageAccountName);
                } catch (RuntimeException cleanupFailure) {
                    if (operationFailure != null) {
                        operationFailure.addSuppressed(cleanupFailure);
                    } else {
                        throw cleanupFailure;
                    }
                }
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " is required.");
        }
        return value;
    }
}
