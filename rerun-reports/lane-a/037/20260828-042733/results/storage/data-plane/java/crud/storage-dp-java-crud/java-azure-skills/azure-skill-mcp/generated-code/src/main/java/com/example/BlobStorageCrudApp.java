package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.DefaultAzureCredential;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobStorageException;

import java.io.UncheckedIOException;

public final class BlobStorageCrudApp {
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";
    private static final String SOURCE_FILE = "data.txt";
    private static final String DOWNLOAD_FILE = "data-downloaded.txt";

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) {
        String operation = "initializing the Blob Storage client";

        try {
            String endpoint = requireEnvironmentVariable("AZURE_STORAGE_BLOB_ENDPOINT");
            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
            BlobServiceClient serviceClient = new BlobServiceClientBuilder()
                    .endpoint(endpoint)
                    .credential(credential)
                    .buildClient();

            operation = "creating container " + CONTAINER_NAME;
            BlobContainerClient containerClient =
                    serviceClient.getBlobContainerClient(CONTAINER_NAME);
            boolean containerCreated = containerClient.createIfNotExists();
            System.out.printf("Container %s: %s%n", CONTAINER_NAME,
                    containerCreated ? "created" : "already exists");

            operation = "uploading blob " + BLOB_NAME;
            BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);
            blobClient.uploadFromFile(SOURCE_FILE, true);
            System.out.printf("Uploaded %s as %s%n", SOURCE_FILE, BLOB_NAME);

            operation = "listing blobs in container " + CONTAINER_NAME;
            System.out.println("Blobs:");
            for (BlobItem blob : containerClient.listBlobs()) {
                System.out.printf("  %s (%d bytes)%n",
                        blob.getName(), blob.getProperties().getContentLength());
            }

            operation = "downloading blob " + BLOB_NAME;
            blobClient.downloadToFile(DOWNLOAD_FILE, true);
            System.out.printf("Downloaded %s to %s%n", BLOB_NAME, DOWNLOAD_FILE);

            operation = "deleting blob " + BLOB_NAME;
            boolean blobDeleted = blobClient.deleteIfExists();
            System.out.printf("Blob %s: %s%n", BLOB_NAME,
                    blobDeleted ? "deleted" : "not found");

            operation = "deleting container " + CONTAINER_NAME;
            boolean containerDeleted = containerClient.deleteIfExists();
            System.out.printf("Container %s: %s%n", CONTAINER_NAME,
                    containerDeleted ? "deleted" : "not found");
        } catch (BlobStorageException exception) {
            System.err.printf(
                    "Azure Blob Storage failed while %s. Status: %d, error code: %s, message: %s%n",
                    operation,
                    exception.getStatusCode(),
                    exception.getErrorCode(),
                    exception.getServiceMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.printf("Azure authentication failed: %s%n", exception.getMessage());
            System.exit(1);
        } catch (UncheckedIOException exception) {
            System.err.printf("Local file operation failed: %s%n", exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException exception) {
            System.err.printf("Configuration error: %s%n", exception.getMessage());
            System.exit(1);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                    name + " must be set, for example https://<account>.blob.core.windows.net");
        }
        return value;
    }
}
