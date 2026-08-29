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
import com.azure.resourcemanager.resources.models.ResourceGroups;

import java.util.Locale;
import java.util.Map;
import java.util.UUID;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class ResourceGroupManagerApp {
    private static final Logger LOGGER = Logger.getLogger(ResourceGroupManagerApp.class.getName());
    private static final String REGION = "eastus";
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "azure-java-sdk";

    private ResourceGroupManagerApp() {
    }

    public static void main(String[] args) {
        try {
            String subscriptionId = requiredEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
            String resourceGroupName = resourceGroupName();

            AzureResourceManager azure = createClient(subscriptionId);
            manageResourceGroup(azure.resourceGroups(), resourceGroupName);
        } catch (IllegalArgumentException exception) {
            LOGGER.log(Level.SEVERE, "Invalid configuration: {0}", exception.getMessage());
            System.exit(2);
        } catch (ClientAuthenticationException exception) {
            LOGGER.log(Level.SEVERE,
                "Azure authentication failed. Check the DefaultAzureCredential configuration.", exception);
            System.exit(3);
        } catch (HttpResponseException exception) {
            LOGGER.log(Level.SEVERE, () -> String.format(
                "Azure Resource Manager request failed with status %d: %s",
                exception.getResponse().getStatusCode(),
                exception.getMessage()));
            System.exit(4);
        } catch (AzureException exception) {
            LOGGER.log(Level.SEVERE, "Azure SDK operation failed.", exception);
            System.exit(5);
        }
    }

    private static AzureResourceManager createClient(String subscriptionId) {
        AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);
        TokenCredential credential = new DefaultAzureCredentialBuilder()
            .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())
            .build();

        return AzureResourceManager.authenticate(credential, profile)
            .withSubscription(subscriptionId);
    }

    private static void manageResourceGroup(ResourceGroups resourceGroups, String resourceGroupName) {
        boolean created = false;
        boolean deleted = false;

        try {
            LOGGER.log(Level.INFO, "Creating resource group {0} in {1}.",
                new Object[] {resourceGroupName, REGION});
            ResourceGroup createdGroup = resourceGroups.define(resourceGroupName)
                .withRegion(REGION)
                .create();
            created = true;
            logDetails("Created", createdGroup);

            LOGGER.info("Resource groups in the subscription:");
            resourceGroups.list().forEach(group ->
                LOGGER.log(Level.INFO, "  {0} ({1})",
                    new Object[] {group.name(), group.regionName()}));

            ResourceGroup fetchedGroup = resourceGroups.getByName(resourceGroupName);
            if (fetchedGroup == null) {
                throw new AzureException("Created resource group could not be retrieved: " + resourceGroupName);
            }
            logDetails("Fetched", fetchedGroup);

            LOGGER.log(Level.INFO, "Adding tag {0}={1}.", new Object[] {TAG_NAME, TAG_VALUE});
            ResourceGroup taggedGroup = fetchedGroup.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            logDetails("Updated", taggedGroup);

            LOGGER.log(Level.INFO, "Deleting resource group {0}.", resourceGroupName);
            resourceGroups.deleteByName(resourceGroupName);
            deleted = true;
            LOGGER.log(Level.INFO, "Deleted resource group {0}.", resourceGroupName);
        } finally {
            if (created && !deleted) {
                cleanupResourceGroup(resourceGroups, resourceGroupName);
            }
        }
    }

    private static void cleanupResourceGroup(ResourceGroups resourceGroups, String resourceGroupName) {
        try {
            LOGGER.log(Level.WARNING,
                "An earlier operation failed; attempting to delete resource group {0}.", resourceGroupName);
            resourceGroups.deleteByName(resourceGroupName);
        } catch (AzureException cleanupException) {
            LOGGER.log(Level.SEVERE,
                "Cleanup failed. Delete the resource group manually: " + resourceGroupName,
                cleanupException);
        }
    }

    private static void logDetails(String operation, ResourceGroup resourceGroup) {
        Map<String, String> tags = resourceGroup.tags();
        LOGGER.log(Level.INFO,
            "{0} resource group: name={1}, region={2}, provisioningState={3}, tags={4}",
            new Object[] {
                operation,
                resourceGroup.name(),
                resourceGroup.regionName(),
                resourceGroup.provisioningState(),
                tags
            });
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must be set.");
        }
        return value.trim();
    }

    private static String resourceGroupName() {
        String configuredName = System.getenv("RESOURCE_GROUP_NAME");
        if (configuredName != null && !configuredName.isBlank()) {
            return configuredName.trim();
        }

        String suffix = UUID.randomUUID().toString().substring(0, 8).toLowerCase(Locale.ROOT);
        return "java-sdk-rg-" + suffix;
    }
}
