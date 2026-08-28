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
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";
    private static final Path SOURCE_FILE = Path.of("data.txt");
    private static final Path DOWNLOADED_FILE = Path.of("data-downloaded.txt");

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) {
        String accountUrl = System.getenv("AZURE_STORAGE_ACCOUNT_URL");
        if (accountUrl == null || accountUrl.isBlank()) {
            System.err.println(
                "AZURE_STORAGE_ACCOUNT_URL must be set, for example "
                    + "https://<account-name>.blob.core.windows.net.");
            System.exit(1);
        }

        if (!Files.isRegularFile(SOURCE_FILE)) {
            System.err.println("Source file does not exist: " + SOURCE_FILE.toAbsolutePath());
            System.exit(1);
        }

        try {
            runCrudOperations(accountUrl);
        } catch (BlobStorageException exception) {
            System.err.printf(
                "Azure Blob Storage request failed (status=%d, errorCode=%s): %s%n",
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
        }
    }

    private static void runCrudOperations(String accountUrl) throws IOException {
        BlobServiceClient serviceClient = new BlobServiceClientBuilder()
            .endpoint(accountUrl)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        BlobContainerClient containerClient =
            serviceClient.getBlobContainerClient(CONTAINER_NAME);
        boolean containerCreated = containerClient.createIfNotExists();
        System.out.printf(
            "Container '%s' %s.%n",
            CONTAINER_NAME,
            containerCreated ? "created" : "already exists");

        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);
        blobClient.uploadFromFile(SOURCE_FILE.toString(), true);
        System.out.printf("Uploaded '%s' as '%s'.%n", SOURCE_FILE, BLOB_NAME);

        System.out.println("Blobs in container:");
        for (BlobItem item : containerClient.listBlobs()) {
            Long size = item.getProperties() == null
                ? null
                : item.getProperties().getContentLength();
            System.out.printf(
                "  %s (%s bytes)%n",
                item.getName(),
                size == null ? "unknown" : size);
        }

        Files.deleteIfExists(DOWNLOADED_FILE);
        blobClient.downloadToFile(DOWNLOADED_FILE.toString());
        System.out.printf("Downloaded '%s' to '%s'.%n", BLOB_NAME, DOWNLOADED_FILE);

        blobClient.delete();
        System.out.printf("Deleted blob '%s'.%n", BLOB_NAME);

        containerClient.delete();
        System.out.printf("Deleted container '%s'.%n", CONTAINER_NAME);
    }
}
