package com.example.azure;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.exception.ManagementException;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;

import java.util.UUID;

public final class ResourceGroupManager {
    private static final String LOCATION = "eastus";
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "azure-sdk-for-java";

    private ResourceGroupManager() {
    }

    public static void main(String[] args) {
        String subscriptionId = requireEnvironmentVariable("AZURE_SUBSCRIPTION_ID");
        String resourceGroupName = getResourceGroupName();

        AzureResourceManager azure = null;
        ResourceGroup createdGroup = null;
        boolean deleted = false;

        try {
            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
            AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);

            azure = AzureResourceManager
                .authenticate(credential, profile)
                .withSubscription(subscriptionId);

            System.out.printf("Creating resource group '%s' in %s...%n",
                resourceGroupName, LOCATION);
            createdGroup = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(LOCATION)
                .create();

            System.out.println("\nResource groups in the subscription:");
            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {
                System.out.printf("- %s (%s)%n",
                    resourceGroup.name(), resourceGroup.regionName());
            }

            ResourceGroup fetchedGroup = azure.resourceGroups()
                .getByName(resourceGroupName);
            if (fetchedGroup == null) {
                throw new IllegalStateException(
                    "The created resource group could not be retrieved: " + resourceGroupName);
            }

            System.out.println("\nCreated resource group details:");
            printResourceGroup(fetchedGroup);

            ResourceGroup taggedGroup = fetchedGroup.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            System.out.printf("%nAdded tag %s=%s. Current tags: %s%n",
                TAG_NAME, TAG_VALUE, taggedGroup.tags());

            System.out.printf("%nDeleting resource group '%s'...%n", resourceGroupName);
            azure.resourceGroups().deleteByName(resourceGroupName);
            deleted = true;
            System.out.println("Resource group deleted.");
        } catch (CredentialUnavailableException exception) {
            System.err.println("No credential was available to DefaultAzureCredential: "
                + exception.getMessage());
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
        } catch (ManagementException exception) {
            System.err.printf("Azure Resource Manager request failed (status %d): %s%n",
                exception.getResponse().getStatusCode(), exception.getMessage());
        } finally {
            if (azure != null && createdGroup != null && !deleted) {
                deleteAfterFailure(azure, resourceGroupName);
            }
        }
    }

    private static void printResourceGroup(ResourceGroup resourceGroup) {
        System.out.println("Name: " + resourceGroup.name());
        System.out.println("ID: " + resourceGroup.id());
        System.out.println("Region: " + resourceGroup.regionName());
        System.out.println("Provisioning state: " + resourceGroup.provisioningState());
        System.out.println("Tags: " + resourceGroup.tags());
    }

    private static void deleteAfterFailure(
        AzureResourceManager azure,
        String resourceGroupName
    ) {
        try {
            System.err.printf("Cleaning up resource group '%s' after failure...%n",
                resourceGroupName);
            azure.resourceGroups().deleteByName(resourceGroupName);
        } catch (ClientAuthenticationException | ManagementException cleanupException) {
            System.err.printf("Cleanup failed for resource group '%s': %s%n",
                resourceGroupName, cleanupException.getMessage());
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Required environment variable is not set: " + name);
        }
        return value;
    }

    private static String getResourceGroupName() {
        String configuredName = System.getenv("AZURE_RESOURCE_GROUP_NAME");
        if (configuredName != null && !configuredName.isBlank()) {
            return configuredName;
        }
        return "java-sdk-rg-" + UUID.randomUUID().toString().substring(0, 8);
    }
}
