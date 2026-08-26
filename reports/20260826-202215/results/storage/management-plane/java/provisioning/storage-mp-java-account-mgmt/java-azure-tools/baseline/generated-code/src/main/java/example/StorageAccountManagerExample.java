package example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.storage.StorageManager;
import com.azure.resourcemanager.storage.models.BlobServiceProperties;
import com.azure.resourcemanager.storage.models.StorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountSkuType;

public final class StorageAccountManagerExample {
    private StorageAccountManagerExample() {
    }

    public static void main(String[] args) {
        try {
            manageStorageAccount();
        } catch (CredentialUnavailableException e) {
            System.err.println("No DefaultAzureCredential source is available: " + e.getMessage());
            System.exit(1);
        } catch (ClientAuthenticationException e) {
            System.err.println("Azure authentication failed: " + e.getMessage());
            System.exit(1);
        } catch (HttpResponseException e) {
            int statusCode = e.getResponse() == null
                ? -1
                : e.getResponse().getStatusCode();
            System.err.printf(
                "Azure Storage management request failed (HTTP %d): %s%n",
                statusCode,
                e.getMessage());
            for (Throwable suppressed : e.getSuppressed()) {
                System.err.println("Cleanup also failed: " + suppressed.getMessage());
            }
            System.exit(1);
        } catch (IllegalArgumentException e) {
            System.err.println("Invalid configuration: " + e.getMessage());
            System.exit(2);
        } catch (RuntimeException e) {
            System.err.println("Unexpected Azure SDK failure: " + e.getMessage());
            for (Throwable suppressed : e.getSuppressed()) {
                System.err.println("Cleanup also failed: " + suppressed.getMessage());
            }
            System.exit(1);
        }
    }

    private static void manageStorageAccount() {
        String subscriptionId = requiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
        String resourceGroupName = requiredEnvironmentVariable("AZURE_RESOURCE_GROUP");
        String storageAccountName = requiredEnvironmentVariable("AZURE_STORAGE_ACCOUNT_NAME");

        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        AzureProfile profile = new AzureProfile(
            null,
            subscriptionId,
            AzureEnvironment.AZURE);
        StorageManager storageManager = StorageManager.authenticate(credential, profile);

        StorageAccount createdAccount = null;
        RuntimeException operationFailure = null;

        try {
            createdAccount = storageManager.storageAccounts()
                .define(storageAccountName)
                .withRegion(Region.US_EAST)
                .withExistingResourceGroup(resourceGroupName)
                .withSku(StorageAccountSkuType.STANDARD_LRS)
                .create();
            System.out.println("Created storage account: " + createdAccount.name());

            System.out.printf("Storage accounts in resource group '%s':%n", resourceGroupName);
            for (StorageAccount account
                : storageManager.storageAccounts().listByResourceGroup(resourceGroupName)) {
                System.out.printf("- %s (%s)%n", account.name(), account.regionName());
            }

            StorageAccount account = storageManager.storageAccounts()
                .getByResourceGroup(resourceGroupName, storageAccountName);
            if (account == null) {
                throw new IllegalStateException(
                    "The created storage account could not be retrieved.");
            }

            System.out.printf(
                "Properties: id=%s, location=%s, sku=%s, provisioningState=%s%n",
                account.id(),
                account.regionName(),
                account.skuType(),
                account.provisioningState());

            BlobServiceProperties blobServiceProperties = storageManager.blobServices()
                .getServicePropertiesAsync(resourceGroupName, storageAccountName)
                .block();
            if (blobServiceProperties == null) {
                throw new IllegalStateException(
                    "Blob service properties could not be retrieved.");
            }

            blobServiceProperties.update()
                .withBlobVersioningEnabled()
                .apply();
            System.out.println("Blob versioning enabled.");
        } catch (RuntimeException e) {
            operationFailure = e;
            throw e;
        } finally {
            if (createdAccount != null) {
                try {
                    storageManager.storageAccounts()
                        .deleteByResourceGroup(resourceGroupName, storageAccountName);
                    System.out.println("Deleted storage account: " + storageAccountName);
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

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Environment variable " + name + " must be set.");
        }
        return value;
    }
}
