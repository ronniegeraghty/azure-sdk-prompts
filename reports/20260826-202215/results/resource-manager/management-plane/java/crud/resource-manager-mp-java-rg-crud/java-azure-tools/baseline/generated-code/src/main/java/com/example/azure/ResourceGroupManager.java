package com.example.azure;

import com.azure.core.credential.TokenCredential;
import com.azure.core.exception.ClientAuthenticationException;
import com.azure.core.exception.HttpResponseException;
import com.azure.core.management.AzureEnvironment;
import com.azure.core.management.Region;
import com.azure.core.management.profile.AzureProfile;
import com.azure.identity.CredentialUnavailableException;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.resourcemanager.AzureResourceManager;
import com.azure.resourcemanager.resources.models.ResourceGroup;

import java.time.Instant;

public final class ResourceGroupManager {
    private static final String SUBSCRIPTION_ID_ENV = "AZURE_SUBSCRIPTION_ID";
    private static final String TAG_NAME = "managed-by";
    private static final String TAG_VALUE = "azure-resourcemanager-java";

    private ResourceGroupManager() {
    }

    public static void main(String[] args) {
        try {
            String subscriptionId = requiredEnvironmentVariable(SUBSCRIPTION_ID_ENV);
            String resourceGroupName = args.length > 0
                ? args[0]
                : "java-sdk-rg-" + Instant.now().getEpochSecond();

            manageResourceGroup(subscriptionId, resourceGroupName);
        } catch (CredentialUnavailableException e) {
            System.err.println("No credential was available to DefaultAzureCredential: " + e.getMessage());
            System.exit(2);
        } catch (ClientAuthenticationException e) {
            System.err.println("Azure authentication failed: " + e.getMessage());
            System.exit(3);
        } catch (HttpResponseException e) {
            int statusCode = e.getResponse() == null ? -1 : e.getResponse().getStatusCode();
            System.err.printf("Azure Resource Manager request failed (HTTP %d): %s%n",
                statusCode, e.getMessage());
            System.exit(4);
        } catch (IllegalArgumentException | IllegalStateException e) {
            System.err.println("Configuration or resource state error: " + e.getMessage());
            System.exit(5);
        } catch (RuntimeException e) {
            System.err.println("Unexpected failure: " + e.getMessage());
            e.printStackTrace(System.err);
            System.exit(1);
        }
    }

    private static void manageResourceGroup(String subscriptionId, String resourceGroupName) {
        TokenCredential credential = new DefaultAzureCredentialBuilder().build();
        AzureProfile profile = new AzureProfile(AzureEnvironment.AZURE);

        AzureResourceManager azure = AzureResourceManager
            .authenticate(credential, profile)
            .withSubscription(subscriptionId);

        boolean resourceGroupCreated = false;
        RuntimeException operationFailure = null;

        try {
            ResourceGroup created = azure.resourceGroups()
                .define(resourceGroupName)
                .withRegion(Region.US_EAST)
                .create();
            resourceGroupCreated = true;
            System.out.printf("Created resource group '%s' in '%s'.%n",
                created.name(), created.regionName());

            System.out.printf("Resource groups in subscription %s:%n", subscriptionId);
            for (ResourceGroup resourceGroup : azure.resourceGroups().list()) {
                System.out.printf("- %s (%s)%n",
                    resourceGroup.name(), resourceGroup.regionName());
            }

            ResourceGroup details = azure.resourceGroups().getByName(resourceGroupName);
            if (details == null) {
                throw new IllegalStateException(
                    "Created resource group could not be retrieved: " + resourceGroupName);
            }
            System.out.printf("Details: id=%s, name=%s, region=%s, tags=%s%n",
                details.id(), details.name(), details.regionName(), details.tags());

            ResourceGroup tagged = details.update()
                .withTag(TAG_NAME, TAG_VALUE)
                .apply();
            System.out.printf("Added tag %s=%s. Current tags: %s%n",
                TAG_NAME, TAG_VALUE, tagged.tags());
        } catch (RuntimeException e) {
            operationFailure = e;
            throw e;
        } finally {
            if (resourceGroupCreated) {
                try {
                    azure.resourceGroups().deleteByName(resourceGroupName);
                    System.out.printf("Deleted resource group '%s'.%n", resourceGroupName);
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

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                "Set the " + name + " environment variable before running the application.");
        }
        return value;
    }
}
