package com.example.azure;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.core.models.AzureCloud;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;

import java.util.Arrays;
import java.util.Locale;
import java.util.Map;
import java.util.UUID;

public final class ResourceGroupManagerApp {
    private static final String EXECUTE_ARGUMENT = "--execute";
    private static final String RESOURCE_GROUP_NAME_VARIABLE = "RESOURCE_GROUP_NAME";
    private static final String SUBSCRIPTION_ID_VARIABLE = "AZURE_SUBSCRIPTION_ID";
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "java-sdk";

    private ResourceGroupManagerApp() {
    }

    public static void main(String[] args) {
        if (!Arrays.asList(args).contains(EXECUTE_ARGUMENT)) {
            printDryRunMessage();
            return;
        }

        try {
            manageResourceGroup();
        } catch (ClientAuthenticationException exception) {
            System.err.printf("Azure authentication failed: %s%n", exception.getMessage());
            System.exit(1);
        } catch (HttpResponseException exception) {
            int statusCode = exception.getResponse() == null
                ? -1
                : exception.getResponse().getStatusCode();
            System.err.printf(
                "Azure Resource Manager request failed (HTTP %s): %s%n",
                statusCode < 0 ? "unknown" : Integer.toString(statusCode),
                exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException | IllegalStateException exception) {
            System.err.printf("Invalid configuration or state: %s%n", exception.getMessage());
            System.exit(1);
        } catch (RuntimeException exception) {
            System.err.printf("Unexpected failure: %s%n", exception.getMessage());
            System.exit(1);
        }
    }

    private static void manageResourceGroup() {
        String subscriptionId = requireEnvironmentVariable(SUBSCRIPTION_ID_VARIABLE);
        String resourceGroupName = getResourceGroupName();

        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        AzureProfile profile = new AzureProfile(AzureCloud.AZURE_PUBLIC_CLOUD);
        AzureResourceManager azure = AzureResourceManager
            .authenticate(credential, profile)
            .withSubscription(subscriptionId);

        boolean resourceGroupExists = false;
        try {
            System.out.printf(
                "Creating resource group '%s' in %s...%n",
                resourceGroupName,
                Region.US_EAST.name());
            ResourceGroup created = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(Region.US_EAST)
                .create();
            resourceGroupExists = true;
            printResourceGroup("Created", created);

            System.out.println("Resource groups in the subscription:");
            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {
                System.out.printf(
                    "  - %s (%s)%n",
                    resourceGroup.name(),
                    resourceGroup.regionName());
            }

            ResourceGroup details = azure.resourceGroups().getByName(resourceGroupName);
            if (details == null) {
                throw new IllegalStateException(
                    "The newly created resource group could not be retrieved.");
            }
            printResourceGroup("Retrieved", details);

            ResourceGroup tagged = details.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            System.out.printf(
                "Added tag %s=%s. Current tags: %s%n",
                TAG_NAME,
                TAG_VALUE,
                tagged.tags());

            System.out.printf("Deleting resource group '%s'...%n", resourceGroupName);
            azure.resourceGroups().deleteByName(resourceGroupName);
            resourceGroupExists = false;
            System.out.println("Resource group deleted.");
        } finally {
            if (resourceGroupExists) {
                cleanupResourceGroup(azure, resourceGroupName);
            }
        }
    }

    private static String getResourceGroupName() {
        String configuredName = System.getenv(RESOURCE_GROUP_NAME_VARIABLE);
        if (configuredName != null && !configuredName.isBlank()) {
            return configuredName.trim();
        }

        return "rg-java-sdk-demo-"
            + UUID.randomUUID().toString().substring(0, 8).toLowerCase(Locale.ROOT);
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Set the " + name + " environment variable before using --execute.");
        }
        return value.trim();
    }

    private static void printResourceGroup(String action, ResourceGroup resourceGroup) {
        Map<String, String> tags = resourceGroup.tags();
        System.out.printf(
            "%s resource group: name=%s, id=%s, region=%s, provisioningState=%s, tags=%s%n",
            action,
            resourceGroup.name(),
            resourceGroup.id(),
            resourceGroup.regionName(),
            resourceGroup.provisioningState(),
            tags);
    }

    private static void cleanupResourceGroup(
        AzureResourceManager azure,
        String resourceGroupName
    ) {
        try {
            System.err.printf(
                "Cleaning up resource group '%s' after the failed operation...%n",
                resourceGroupName);
            azure.resourceGroups().deleteByName(resourceGroupName);
            System.err.println("Cleanup completed.");
        } catch (RuntimeException cleanupException) {
            System.err.printf(
                "Cleanup failed for resource group '%s': %s%n",
                resourceGroupName,
                cleanupException.getMessage());
        }
    }

    private static void printDryRunMessage() {
        System.out.println("Dry run: no Azure requests were made.");
        System.out.println("The --execute option performs this sequence:");
        System.out.println("  1. Authenticate with DefaultAzureCredential.");
        System.out.println("  2. Create an eastus resource group.");
        System.out.println("  3. List all resource groups.");
        System.out.println("  4. Retrieve the created resource group.");
        System.out.printf("  5. Add the tag %s=%s.%n", TAG_NAME, TAG_VALUE);
        System.out.println("  6. Delete the resource group.");
    }
}
