package com.example;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.storage.StorageManager;
import com.azure.resourcemanager.storage.models.BlobServiceProperties;
import com.azure.resourcemanager.storage.models.StorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountSkuType;

public final class StorageAccountManagerApp {
    private StorageAccountManagerApp() {
    }

    public static void main(String[] args) {
        if (args.length != 3) {
            System.err.println(
                "Usage: mvn exec:java -Dexec.args=\"<subscription-id> <resource-group> <storage-account-name>\"");
            System.exit(2);
        }

        String subscriptionId = args[0];
        String resourceGroupName = args[1];
        String accountName = args[2];
        boolean accountCreated = false;
        int exitCode = 0;
        StorageManager storageManager = null;

        try {
            AzureProfile profile = new AzureProfile(
                null, subscriptionId, AzureEnvironment.AZURE);
            TokenCredential credential = new DefaultAzureCredentialBuilder()
                .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())
                .build();

            storageManager = StorageManager.authenticate(credential, profile);

            StorageAccount createdAccount = storageManager.storageAccounts()
                .define(accountName)
                .withRegion(Region.US_EAST)
                .withExistingResourceGroup(resourceGroupName)
                .withSku(StorageAccountSkuType.STANDARD_LRS)
                .withGeneralPurposeAccountKindV2()
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
                .getByResourceGroup(resourceGroupName, accountName);
            if (account == null) {
                throw new IllegalStateException(
                    "The created storage account could not be retrieved: " + accountName);
            }
            System.out.printf(
                "Properties: name=%s, location=%s, sku=%s, kind=%s, provisioningState=%s%n",
                account.name(),
                account.regionName(),
                account.skuType(),
                account.kind(),
                account.provisioningState());

            BlobServiceProperties blobProperties = storageManager.blobServices()
                .getServicePropertiesAsync(resourceGroupName, accountName)
                .block();
            if (blobProperties == null) {
                throw new IllegalStateException(
                    "Blob service properties were not returned for: " + accountName);
            }

            blobProperties.update()
                .withBlobVersioningEnabled()
                .apply();
            System.out.printf("Enabled blob versioning for: %s%n", accountName);

            storageManager.storageAccounts().deleteByResourceGroup(resourceGroupName, accountName);
            accountCreated = false;
            System.out.printf("Deleted storage account: %s%n", accountName);
        } catch (HttpResponseException e) {
            System.err.printf("Azure management request failed (status %d): %s%n",
                e.getResponse() == null ? -1 : e.getResponse().getStatusCode(),
                e.getMessage());
            exitCode = 1;
        } catch (RuntimeException e) {
            System.err.printf("Storage account operation failed: %s%n", e.getMessage());
            exitCode = 1;
        } finally {
            if (accountCreated && storageManager != null) {
                try {
                    storageManager.storageAccounts()
                        .deleteByResourceGroup(resourceGroupName, accountName);
                    System.err.printf(
                        "Cleaned up storage account after failure: %s%n", accountName);
                } catch (RuntimeException cleanupError) {
                    exitCode = 1;
                    System.err.printf(
                        "Cleanup failed for storage account '%s': %s%n",
                        accountName, cleanupError.getMessage());
                    System.err.printf(
                        "The account may still exist and require manual cleanup.%n");
                }
            }
        }

        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }
}
