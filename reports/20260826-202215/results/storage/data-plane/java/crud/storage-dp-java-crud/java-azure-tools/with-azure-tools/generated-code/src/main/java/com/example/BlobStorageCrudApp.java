package com.example;

import com.azure.core.exception.ClientAuthenticationException;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobStorageException;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public final class BlobStorageCrudApp {
    private static final String ACCOUNT_URL_ENVIRONMENT_VARIABLE = "AZURE_STORAGE_ACCOUNT_URL";
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";
    private static final Path SOURCE_FILE = Path.of("data.txt");
    private static final Path DOWNLOAD_FILE = Path.of("data-downloaded.txt");

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) {
        try {
            runCrudOperations();
        } catch (BlobStorageException exception) {
            System.err.printf(
                "Blob Storage request failed (status=%d, errorCode=%s): %s%n",
                exception.getStatusCode(),
                exception.getErrorCode(),
                exception.getServiceMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            System.exit(1);
        } catch (IOException exception) {
            System.err.println("Local file operation failed: " + exception.getMessage());
            System.exit(1);
        } catch (IllegalArgumentException | IllegalStateException exception) {
            System.err.println("Invalid configuration: " + exception.getMessage());
            System.exit(1);
        }
    }

    private static void runCrudOperations() throws IOException {
        if (!Files.isRegularFile(SOURCE_FILE)) {
            throw new IOException("Source file does not exist: " + SOURCE_FILE.toAbsolutePath());
        }

        BlobServiceClient serviceClient = new BlobServiceClientBuilder()
            .endpoint(requiredEnvironmentVariable(ACCOUNT_URL_ENVIRONMENT_VARIABLE))
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        BlobContainerClient containerClient =
            serviceClient.getBlobContainerClient(CONTAINER_NAME);
        boolean containerCreated = containerClient.createIfNotExists();
        System.out.printf(
            "Container %s: %s%n",
            CONTAINER_NAME,
            containerCreated ? "created" : "already exists");

        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);
        blobClient.uploadFromFile(SOURCE_FILE.toString(), true);
        System.out.println("Uploaded blob: " + BLOB_NAME);

        System.out.println("Blobs in " + CONTAINER_NAME + ":");
        for (BlobItem blobItem : containerClient.listBlobs()) {
            Long size = blobItem.getProperties().getContentLength();
            System.out.printf("  %s (%s bytes)%n", blobItem.getName(), size);
        }

        Files.deleteIfExists(DOWNLOAD_FILE);
        blobClient.downloadToFile(DOWNLOAD_FILE.toString());
        System.out.println("Downloaded blob to: " + DOWNLOAD_FILE.toAbsolutePath());

        blobClient.delete();
        System.out.println("Deleted blob: " + BLOB_NAME);

        containerClient.delete();
        System.out.println("Deleted container: " + CONTAINER_NAME);
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " must be set");
        }
        return value;
    }
}
