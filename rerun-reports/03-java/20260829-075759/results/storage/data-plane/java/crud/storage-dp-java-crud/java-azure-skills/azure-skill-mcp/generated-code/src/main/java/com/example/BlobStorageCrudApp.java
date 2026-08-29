package com.example;

import com.azure.core.exception.ClientAuthenticationException;
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
    private static final Path UPLOAD_FILE = Path.of("data.txt");
    private static final Path DOWNLOAD_FILE = Path.of("data-downloaded.txt");

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) {
        String accountUrl = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL");

        if (!Files.isRegularFile(UPLOAD_FILE)) {
            System.err.printf("Upload file does not exist or is not a regular file: %s%n",
                    UPLOAD_FILE.toAbsolutePath());
            System.exit(2);
        }

        BlobServiceClient serviceClient = new BlobServiceClientBuilder()
                .endpoint(accountUrl)
                .credential(new DefaultAzureCredentialBuilder().build())
                .buildClient();

        BlobContainerClient containerClient = serviceClient.getBlobContainerClient(CONTAINER_NAME);
        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);

        try {
            containerClient.createIfNotExists();
            System.out.printf("Container ready: %s%n", CONTAINER_NAME);

            blobClient.uploadFromFile(UPLOAD_FILE.toString(), true);
            System.out.printf("Uploaded %s as %s%n", UPLOAD_FILE, BLOB_NAME);

            System.out.println("Blobs:");
            for (BlobItem item : containerClient.listBlobs()) {
                System.out.printf("  %s (%d bytes)%n",
                        item.getName(), item.getProperties().getContentLength());
            }

            blobClient.downloadToFile(DOWNLOAD_FILE.toString(), true);
            System.out.printf("Downloaded %s to %s%n", BLOB_NAME, DOWNLOAD_FILE);

            blobClient.delete();
            System.out.printf("Deleted blob: %s%n", BLOB_NAME);

            containerClient.delete();
            System.out.printf("Deleted container: %s%n", CONTAINER_NAME);
        } catch (BlobStorageException exception) {
            System.err.printf(
                    "Azure Blob Storage request failed: status=%d, errorCode=%s, message=%s%n",
                    exception.getStatusCode(),
                    exception.getErrorCode(),
                    exception.getServiceMessage() == null
                            ? exception.getMessage()
                            : exception.getServiceMessage());
            System.exit(1);
        } catch (ClientAuthenticationException exception) {
            System.err.printf("Azure authentication failed: %s%n", exception.getMessage());
            System.exit(1);
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            System.err.printf(
                    "Required environment variable %s is not set. "
                            + "Example: https://<account-name>.blob.core.windows.net%n",
                    name);
            System.exit(2);
        }
        return value;
    }
}
