package com.example;

import com.azure.core.exception.AzureException;
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

public final class BlobStorageCrudApplication {
    private static final String ACCOUNT_URL_ENV = "AZURE_STORAGE_ACCOUNT_URL";
    private static final String CONTAINER_NAME = "my-container";
    private static final String BLOB_NAME = "uploads/data.txt";
    private static final Path UPLOAD_PATH = Path.of("data.txt");
    private static final Path DOWNLOAD_PATH = Path.of("data-downloaded.txt");

    private BlobStorageCrudApplication() {
    }

    public static void main(String[] args) {
        try {
            validateLocalInput();

            String accountUrl = requireEnvironmentVariable(ACCOUNT_URL_ENV);
            DefaultAzureCredential credential = new DefaultAzureCredentialBuilder().build();
            BlobServiceClient serviceClient = new BlobServiceClientBuilder()
                    .endpoint(accountUrl)
                    .credential(credential)
                    .buildClient();

            runCrudOperations(serviceClient);
        } catch (BlobStorageException exception) {
            System.err.printf(
                    "Blob Storage request failed (status=%d, errorCode=%s): %s%n",
                    exception.getStatusCode(),
                    exception.getErrorCode(),
                    exception.getServiceMessage());
            System.exit(1);
        } catch (AzureException exception) {
            System.err.println("Azure authentication or client operation failed: "
                    + exception.getMessage());
            System.exit(1);
        } catch (IOException | IllegalArgumentException exception) {
            System.err.println("Application error: " + exception.getMessage());
            System.exit(1);
        }
    }

    private static void runCrudOperations(BlobServiceClient serviceClient) {
        BlobContainerClient containerClient =
                serviceClient.getBlobContainerClient(CONTAINER_NAME);

        boolean created = containerClient.createIfNotExists();
        System.out.printf("Container %s: %s%n",
                CONTAINER_NAME, created ? "created" : "already exists");

        BlobClient blobClient = containerClient.getBlobClient(BLOB_NAME);
        blobClient.uploadFromFile(UPLOAD_PATH.toString(), true);
        System.out.println("Uploaded blob: " + BLOB_NAME);

        System.out.println("Blobs in " + CONTAINER_NAME + ":");
        for (BlobItem blob : containerClient.listBlobs()) {
            Long size = blob.getProperties().getContentLength();
            System.out.printf("  %s (%d bytes)%n", blob.getName(), size == null ? 0L : size);
        }

        blobClient.downloadToFile(DOWNLOAD_PATH.toString(), true);
        System.out.println("Downloaded blob to: " + DOWNLOAD_PATH);

        blobClient.delete();
        System.out.println("Deleted blob: " + BLOB_NAME);

        containerClient.delete();
        System.out.println("Deleted container: " + CONTAINER_NAME);
    }

    private static void validateLocalInput() throws IOException {
        if (!Files.isRegularFile(UPLOAD_PATH)) {
            throw new IOException("Upload file does not exist or is not a regular file: "
                    + UPLOAD_PATH.toAbsolutePath());
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(
                    "Required environment variable is not set: " + name);
        }
        return value;
    }
}
