package com.example.azure;

import com.azure.core.exception.AzureException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;

import java.util.Arrays;

public final class ResourceGroupManager {
    private static final String LOCATION = "eastus";
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "azure-resourcemanager-java";

    private ResourceGroupManager() {
    }

    public static void main(String[] args) {
        if (!Arrays.asList(args).contains("--execute")) {
            printDryRun();
            return;
        }

        String resourceGroupName = System.getenv("AZURE_RESOURCE_GROUP_NAME");
        if (resourceGroupName == null || resourceGroupName.isBlank()) {
            System.err.println(
                "AZURE_RESOURCE_GROUP_NAME must be set when --execute is used.");
            System.exit(2);
        }

        AzureResourceManager azure = null;
        ResourceGroup createdResourceGroup = null;
        boolean operationFailed = false;

        try {
            DefaultAzureCredential credential =
                new DefaultAzureCredentialBuilder().build();
            AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);

            azure = AzureResourceManager
                .authenticate(credential, profile)
                .withDefaultSubscription();

            System.out.printf(
                "Creating resource group '%s' in '%s'...%n",
                resourceGroupName,
                LOCATION);
            createdResourceGroup = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(LOCATION)
                .create();

            System.out.println("Resource groups in the subscription:");
            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {
                System.out.printf(
                    "- %s (%s)%n",
                    resourceGroup.name(),
                    resourceGroup.regionName());
            }

            ResourceGroup details = azure.resourceGroups()
                .getByName(resourceGroupName);
            if (details == null) {
                throw new IllegalStateException(
                    "Created resource group could not be retrieved: "
                        + resourceGroupName);
            }

            System.out.printf(
                "Created resource group: id=%s, name=%s, region=%s, tags=%s%n",
                details.id(),
                details.name(),
                details.regionName(),
                details.tags());

            ResourceGroup updated = details.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            System.out.printf("Updated tags: %s%n", updated.tags());
        } catch (AzureException exception) {
            operationFailed = true;
            System.err.printf(
                "Azure Resource Manager operation failed: %s%n",
                exception.getMessage());
            exception.printStackTrace(System.err);
        } catch (RuntimeException exception) {
            operationFailed = true;
            System.err.printf(
                "Authentication or application error: %s%n",
                exception.getMessage());
            exception.printStackTrace(System.err);
        } finally {
            if (azure != null && createdResourceGroup != null) {
                try {
                    System.out.printf(
                        "Deleting resource group '%s'...%n",
                        createdResourceGroup.name());
                    azure.resourceGroups().deleteByName(createdResourceGroup.name());
                    System.out.println("Resource group deleted.");
                } catch (RuntimeException exception) {
                    operationFailed = true;
                    System.err.printf(
                        "Cleanup failed for resource group '%s': %s%n",
                        createdResourceGroup.name(),
                        exception.getMessage());
                    exception.printStackTrace(System.err);
                }
            }
        }

        if (operationFailed) {
            System.exit(1);
        }
    }

    private static void printDryRun() {
        System.out.println("Dry run only; no Azure operations were performed.");
        System.out.printf(
            "Planned flow: authenticate, create a resource group in '%s', list, "
                + "get details, tag, and delete.%n",
            LOCATION);
        System.out.println(
            "Set AZURE_RESOURCE_GROUP_NAME and pass --execute to run against Azure.");
    }
}
