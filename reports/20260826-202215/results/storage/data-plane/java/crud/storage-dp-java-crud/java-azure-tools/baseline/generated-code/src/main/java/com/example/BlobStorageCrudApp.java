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

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

public final class BlobStorageCrudApp {
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";

    private BlobStorageCrudApp() {
    }

    public static void main(String[] args) throws IOException {
        String accountUrl = requireEnvironmentVariable("AZURE_STORAGE_ACCOUNT_URL");
        Path uploadPath = Path.of("data.txt");
        Path downloadPath = Path.of("data-downloaded.txt");

        if (!Files.isRegularFile(uploadPath)) {
            throw new IOException("Upload file does not exist or is not a regular file: "
                    + uploadPath.toAbsolutePath());
        }

        DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
        BlobServiceClient serviceClient = new BlobServiceClientBuilder()
                .endpoint(accountUrl)
                .credential(credential)
                .buildClient();

        try {
            BlobContainerClient containerClient =
                    serviceClient.getBlobContainerClient(CONTAINER_NAME);
            containerClient.createIfNotExists();
            System.out.printf("Container ready: %s%n", CONTAINER_NAME);

            BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);
            blobClient.uploadFromFile(uploadPath.toString(), true);
            System.out.printf("Uploaded %s as %s%n", uploadPath, BLOB_NAME);

            System.out.println("Blobs:");
            for (BlobItem blob : containerClient.listBlobs()) {
                Long size = blob.getProperties().getContentLength();
                System.out.printf("  %s (%s bytes)%n",
                        blob.getName(), size == null ? "unknown" : size);
            }

            blobClient.downloadToFile(downloadPath.toString(), true);
            System.out.printf("Downloaded %s to %s%n", BLOB_NAME, downloadPath);

            if (!blobClient.deleteIfExists()) {
                throw new IllegalStateException("Blob disappeared before it could be deleted: " + BLOB_NAME);
            }
            System.out.printf("Deleted blob: %s%n", BLOB_NAME);

            if (!containerClient.deleteIfExists()) {
                throw new IllegalStateException(
                        "Container disappeared before it could be deleted: " + CONTAINER_NAME);
            }
            System.out.printf("Deleted container: %s%n", CONTAINER_NAME);
        } catch (BlobStorageException exception) {
            System.err.printf(
                    "Azure Blob Storage request failed (HTTP %d, error code %s): %s%n",
                    exception.getStatusCode(),
                    exception.getErrorCode(),
                    exception.getServiceMessage());
            throw exception;
        } catch (ClientAuthenticationException exception) {
            System.err.println("Azure authentication failed: " + exception.getMessage());
            throw exception;
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                    name + " must be set, for example https://<account-name>.blob.core.windows.net");
        }
        return value;
    }
}
