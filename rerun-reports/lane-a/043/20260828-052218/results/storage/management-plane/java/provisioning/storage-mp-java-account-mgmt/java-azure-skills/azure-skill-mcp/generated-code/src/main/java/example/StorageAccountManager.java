package example;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.core.models.AzureCloud;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.storage.StorageManager;
import com.azure.resourcemanager.storage.models.BlobServiceProperties;
import com.azure.resourcemanager.storage.models.StorageAccount;
import com.azure.resourcemanager.storage.models.StorageAccountSkuType;

public final class StorageAccountManager {
    private static final Region REGION = Region.US_EAST;

    private StorageAccountManager() {
    }

    public static void main(String[] args) {
        if (args.length != 2) {
            System.err.println(
                "Usage: mvn exec:java -Dexec.args=\"<resource-group> <globally-unique-storage-account-name>\"");
            System.exit(2);
        }

        String resourceGroupName = args[0];
        String storageAccountName = args[1];
        String subscriptionId;

        try {
            subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
            validateStorageAccountName(storageAccountName);
        } catch (IllegalArgumentException exception) {
            System.err.println("Invalid configuration: " + exception.getMessage());
            System.exit(2);
            return;
        }

        StorageManager storageManager = null;
        boolean accountCreated = false;
        int exitCode = 0;

        try {
            storageManager = createStorageManager(subscriptionId);

            StorageAccount createdAccount = storageManager.storageAccounts()
                .define(storageAccountName)
                .withRegion(REGION)
                .withExistingResourceGroup(resourceGroupName)
                .withSku(StorageAccountSkuType.STANDARD_LRS)
                .withGeneralPurposeAccountKindV2()
                .withOnlyHttpsTraffic()
                .create();
            accountCreated = true;
            System.out.printf("Created storage account: %s%n", createdAccount.id());

            System.out.printf("Storage accounts in resource group '%s':%n", resourceGroupName);
            for (StorageAccount account
                : storageManager.storageAccounts().listByResourceGroup(resourceGroupName)) {
                System.out.printf("  %s (%s)%n", account.name(), account.regionName());
            }

            StorageAccount account = storageManager.storageAccounts()
                .getByResourceGroup(resourceGroupName, storageAccountName);
            if (account == null) {
                throw new IllegalStateException(
                    "The created storage account could not be retrieved.");
            }

            System.out.println("Created storage account properties:");
            System.out.printf("  ID: %s%n", account.id());
            System.out.printf("  Name: %s%n", account.name());
            System.out.printf("  Region: %s%n", account.regionName());
            System.out.printf("  SKU: %s%n", account.skuType());
            System.out.printf("  Kind: %s%n", account.kind());

            BlobServiceProperties blobProperties = storageManager.blobServices()
                .getServicePropertiesAsync(resourceGroupName, storageAccountName)
                .block();
            if (blobProperties == null) {
                throw new IllegalStateException(
                    "The storage account's Blob service properties could not be retrieved.");
            }

            BlobServiceProperties updatedBlobProperties = blobProperties.update()
                .withBlobVersioningEnabled()
                .apply();
            if (!Boolean.TRUE.equals(updatedBlobProperties.isBlobVersioningEnabled())) {
                throw new IllegalStateException("Blob versioning was not enabled.");
            }
            System.out.println("Blob versioning enabled.");
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            exitCode = 1;
        } catch (HttpResponseException exception) {
            printHttpError("Azure Storage management operation failed", exception);
            exitCode = 1;
        } catch (IllegalArgumentException | IllegalStateException exception) {
            System.err.println("Storage account operation failed: " + exception.getMessage());
            exitCode = 1;
        } finally {
            if (accountCreated && storageManager != null) {
                try {
                    storageManager.storageAccounts()
                        .deleteByResourceGroup(resourceGroupName, storageAccountName);
                    System.out.printf("Deleted storage account: %s%n", storageAccountName);
                } catch (ClientAuthenticationException exception) {
                    System.err.println(
                        "Authentication failed while deleting the storage account: "
                            + exception.getMessage());
                    exitCode = 1;
                } catch (HttpResponseException exception) {
                    printHttpError("Failed to delete the storage account", exception);
                    exitCode = 1;
                }
            }
        }

        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }

    private static StorageManager createStorageManager(String subscriptionId) {
        AzureProfile profile = new AzureProfile(
            null,
            subscriptionId,
            AzureCloud.AZURE_PUBLIC_CLOUD);

        TokenCredential credential = new DefaultAzureCredentialBuilder()
            .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())
            .build();

        return StorageManager.authenticate(credential, profile);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Environment variable " + name + " must be set.");
        }
        return value;
    }

    private static void validateStorageAccountName(String name) {
        if (!name.matches("[a-z0-9]{3,24}")) {
            throw new IllegalArgumentException(
                "The storage account name must contain 3-24 lowercase letters or digits.");
        }
    }

    private static void printHttpError(String context, HttpResponseException exception) {
        String status = exception.getResponse() == null
            ? "unknown"
            : Integer.toString(exception.getResponse().getStatusCode());
        System.err.printf("%s (HTTP %s): %s%n", context, status, exception.getMessage());
    }
}
