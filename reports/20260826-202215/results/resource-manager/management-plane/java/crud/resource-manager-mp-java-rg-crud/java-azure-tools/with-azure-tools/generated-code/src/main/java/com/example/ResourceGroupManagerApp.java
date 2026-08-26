package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.exception.ManagementException;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;

import java.util.UUID;
import java.util.logging.Level;
import java.util.logging.Logger;

public final class ResourceGroupManagerApp {
    private static final Logger LOGGER = Logger.getLogger(ResourceGroupManagerApp.class.getName());
    private static final Region REGION = Region.US_EAST;
    private static final String TAG_KEY = "environment";
    private static final String TAG_VALUE = "sdk-demo";

    private ResourceGroupManagerApp() {
    }

    public static void main(String[] args) {
        try {
            manageResourceGroup();
        } catch (CredentialUnavailableException exception) {
            LOGGER.log(Level.SEVERE,
                "DefaultAzureCredential could not find a usable credential. Configure a supported "
                    + "developer credential or managed identity.", exception);
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            LOGGER.log(Level.SEVERE,
                "Microsoft Entra authentication failed. Verify the credential and tenant configuration.",
                exception);
            System.exit(1);
        } catch (ManagementException exception) {
            LOGGER.log(Level.SEVERE,
                "Azure Resource Manager rejected an operation. Verify the subscription, RBAC permissions, "
                    + "resource name, and regional policy.", exception);
            System.exit(1);
        } catch (RuntimeException exception) {
            LOGGER.log(Level.SEVERE, "The resource group workflow failed unexpectedly.", exception);
            System.exit(1);
        }
    }

    private static void manageResourceGroup() {
        AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);
        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder()
            .authorityHost(profile.getEnvironment().getActiveDirectoryEndpoint())
            .build();

        AzureResourceManager azure = AzureResourceManager
            .authenticate(credential, profile)
            .withDefaultSubscription();

        String resourceGroupName = "rg-java-sdk-demo-" + UUID.randomUUID()
            .toString()
            .replace("-", "")
            .substring(0, 12);

        boolean created = false;
        RuntimeException operationFailure = null;

        try {
            LOGGER.info(() -> "Creating resource group " + resourceGroupName + " in " + REGION.name());
            ResourceGroup createdGroup = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(REGION)
                .create();
            created = true;
            logResourceGroup("Created", createdGroup);

            LOGGER.info("Listing resource groups in subscription " + azure.subscriptionId());
            for (ResourceGroup group : azure.resourceGroups().list()) {
                LOGGER.info(() -> String.format("Resource group: name=%s, region=%s",
                    group.name(), group.regionName()));
            }

            LOGGER.info(() -> "Getting details for " + resourceGroupName);
            ResourceGroup retrievedGroup = azure.resourceGroups().getByName(resourceGroupName);
            if (retrievedGroup == null) {
                throw new IllegalStateException(
                    "The newly created resource group could not be retrieved: " + resourceGroupName);
            }
            logResourceGroup("Retrieved", retrievedGroup);

            LOGGER.info(() -> String.format("Adding tag %s=%s to %s",
                TAG_KEY, TAG_VALUE, resourceGroupName));
            ResourceGroup taggedGroup = retrievedGroup.update()
                .withTag(TAG_KEY, TAG_VALUE)
                .apply();
            logResourceGroup("Tagged", taggedGroup);
        } catch (RuntimeException exception) {
            operationFailure = exception;
            throw exception;
        } finally {
            if (created) {
                try {
                    LOGGER.info(() -> "Deleting resource group " + resourceGroupName);
                    azure.resourceGroups().deleteByName(resourceGroupName);
                    LOGGER.info(() -> "Deleted resource group " + resourceGroupName);
                } catch (RuntimeException deletionFailure) {
                    if (operationFailure != null) {
                        operationFailure.addSuppressed(deletionFailure);
                    } else {
                        throw deletionFailure;
                    }
                }
            }
        }
    }

    private static void logResourceGroup(String action, ResourceGroup group) {
        LOGGER.info(() -> String.format(
            "%s resource group: name=%s, id=%s, region=%s, provisioningState=%s, tags=%s",
            action,
            group.name(),
            group.id(),
            group.regionName(),
            group.provisioningState(),
            group.tags()));
    }
}
