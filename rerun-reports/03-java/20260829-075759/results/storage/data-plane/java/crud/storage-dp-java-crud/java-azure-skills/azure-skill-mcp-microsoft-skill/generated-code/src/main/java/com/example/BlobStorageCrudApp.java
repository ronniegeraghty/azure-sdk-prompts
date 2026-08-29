package com.example;

import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobContainerClient;
import com.azure.storage.blob.BlobServiceClient;
import com.azure.storage.blob.BlobServiceClientBuilder;
import com.azure.storage.blob.models.BlobItem;
import com.azure.storage.blob.models.BlobStorageException;

import java.nio.file.Files;
import java.nio.file.Path;

public final class BlobStorageCrudApp {
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";
    private static final Path UPLOAD_PATH = Path.of("data.txt");
    private static final Path DOWNLOAD_PATH = Path.of("data-downloaded.txt");

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) {
        String accountUrl = System.getenv("AZURE_STORAGE_ACCOUNT_URL");
        if (accountUrl == null || accountUrl.isBlank()) {
            System.err.println(
                "AZURE_STORAGE_ACCOUNT_URL must be set, for example "
                    + "https://<account-name>.blob.core.windows.net");
            System.exit(2);
        }

        if (!Files.isRegularFile(UPLOAD_PATH)) {
            System.err.printf("Upload file does not exist: %s%n", UPLOAD_PATH.toAbsolutePath());
            System.exit(2);
        }

        BlobServiceClient serviceClient = new BlobServiceClientBuilder()
            .endpoint(accountUrl)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        BlobContainerClient containerClient =
            serviceClient.getBlobContainerClient(CONTAINER_NAME);
        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);

        try {
            boolean containerCreated = containerClient.createIfNotExists();
            System.out.printf(
                "Container %s: %s%n",
                CONTAINER_NAME,
                containerCreated ? "created" : "already exists");

            blobClient.uploadFromFile(UPLOAD_PATH.toString(), true);
            System.out.printf("Uploaded %s as %s%n", UPLOAD_PATH, BLOB_NAME);

            System.out.println("Blobs:");
            for (BlobItem item : containerClient.listBlobs()) {
                System.out.printf(
                    "  %s (%d bytes)%n",
                    item.getName(),
                    item.getProperties().getContentLength());
            }

            blobClient.downloadToFile(DOWNLOAD_PATH.toString(), true);
            System.out.printf("Downloaded %s to %s%n", BLOB_NAME, DOWNLOAD_PATH);

            blobClient.delete();
            System.out.printf("Deleted blob %s%n", BLOB_NAME);

            containerClient.delete();
            System.out.printf("Deleted container %s%n", CONTAINER_NAME);
        } catch (BlobStorageException exception) {
            System.err.printf(
                "Azure Blob Storage request failed: status=%d, errorCode=%s, message=%s%n",
                exception.getStatusCode(),
                exception.getErrorCode(),
                exception.getServiceMessage());
            System.exit(1);
        } catch (RuntimeException exception) {
            System.err.printf("Application failed: %s%n", exception.getMessage());
            System.exit(1);
        }
    }
}
