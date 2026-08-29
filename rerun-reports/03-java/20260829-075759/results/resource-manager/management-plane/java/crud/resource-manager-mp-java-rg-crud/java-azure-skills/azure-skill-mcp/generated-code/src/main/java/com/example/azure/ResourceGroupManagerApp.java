package com.example.azure;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.AzureException;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;
import com.azure.core.management.Region;

import java.util.Optional;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class ResourceGroupManagerApp {
    private static final Logger LOGGER = Logger.getLogger(ResourceGroupManagerApp.class.getName());
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "java-sdk-sample";

    private ResourceGroupManagerApp() {
    }

    public static void main(String[] args) {
        int exitCode = run(args);
        if (exitCode != 0) {
            System.exit(exitCode);
        }
    }

    static int run(String[] args) {
        String resourceGroupName;
        try {
            resourceGroupName = getResourceGroupName(args);
        } catch (IllegalArgumentException exception) {
            LOGGER.severe(exception.getMessage());
            return 2;
        }

        AzureResourceManager azure = null;
        boolean resourceGroupCreated = false;

        try {
            AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);
            TokenCredential credential = new DefaultAzureCredentialBuilder()
                .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())
                .build();

            azure = AzureResourceManager.authenticate(credential, profile)
                .withDefaultSubscription();

            LOGGER.info(() -> "Creating resource group '" + resourceGroupName + "' in eastus...");
            ResourceGroup created = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(Region.US_EAST)
                .create();
            resourceGroupCreated = true;
            logDetails("Created resource group", created);

            LOGGER.info("Resource groups in the subscription:");
            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {
                LOGGER.info(() -> String.format("  %s (%s)", resourceGroup.name(), resourceGroup.regionName()));
            }

            ResourceGroup fetched = azure.resourceGroups().getByName(resourceGroupName);
            if (fetched == null) {
                throw new IllegalStateException("The created resource group could not be retrieved.");
            }
            logDetails("Fetched resource group", fetched);

            ResourceGroup tagged = fetched.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            LOGGER.info(() -> String.format(
                "Added tag %s=%s. Current tags: %s",
                TAG_NAME,
                TAG_VALUE,
                tagged.tags()));

            azure.resourceGroups().deleteByName(resourceGroupName);
            resourceGroupCreated = false;
            LOGGER.info(() -> "Deleted resource group '" + resourceGroupName + "'.");
            return 0;
        } catch (ClientAuthenticationException exception) {
            LOGGER.log(Level.SEVERE,
                "Authentication failed. Check the DefaultAzureCredential configuration and Azure login.", exception);
            return 1;
        } catch (HttpResponseException exception) {
            String status = Optional.ofNullable(exception.getResponse())
                .map(response -> Integer.toString(response.getStatusCode()))
                .orElse("unavailable");
            LOGGER.log(Level.SEVERE, "Azure Resource Manager request failed with HTTP status " + status + ".", exception);
            return 1;
        } catch (AzureException exception) {
            LOGGER.log(Level.SEVERE, "An Azure SDK operation failed.", exception);
            return 1;
        } catch (RuntimeException exception) {
            LOGGER.log(Level.SEVERE, "The application failed.", exception);
            return 1;
        } finally {
            if (resourceGroupCreated && azure != null) {
                try {
                    LOGGER.warning(() -> "Cleaning up resource group '" + resourceGroupName + "' after failure.");
                    azure.resourceGroups().deleteByName(resourceGroupName);
                } catch (AzureException cleanupException) {
                    LOGGER.log(Level.SEVERE,
                        "Cleanup failed. Delete resource group '" + resourceGroupName + "' manually.",
                        cleanupException);
                }
            }
        }
    }

    private static String getResourceGroupName(String[] args) {
        if (args.length != 1 || args[0].isBlank()) {
            throw new IllegalArgumentException(
                "Usage: mvn exec:java -Dexec.args=\"<resource-group-name>\"");
        }
        return args[0].trim();
    }

    private static void logDetails(String action, ResourceGroup resourceGroup) {
        LOGGER.info(() -> String.format(
            "%s: name=%s, id=%s, region=%s, tags=%s",
            action,
            resourceGroup.name(),
            resourceGroup.id(),
            resourceGroup.regionName(),
            resourceGroup.tags()));
    }
}
