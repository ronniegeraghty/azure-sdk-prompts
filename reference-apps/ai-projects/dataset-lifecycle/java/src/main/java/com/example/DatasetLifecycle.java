package com.example;

import com.azure.ai.projects.AIProjectClientBuilder;
import com.azure.ai.projects.DatasetsClient;
import com.azure.ai.projects.models.BlobReference;
import com.azure.ai.projects.models.DatasetCredential;
import com.azure.ai.projects.models.DatasetVersion;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.azure.storage.blob.BlobClient;
import com.azure.storage.blob.BlobClientBuilder;
import java.io.IOException;
import java.net.URI;
import java.net.URISyntaxException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.Arrays;

public final class DatasetLifecycle {
    private DatasetLifecycle() {
    }

    public static void main(String[] args)
        throws IOException, URISyntaxException {
        String endpoint = requireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
        String datasetName = requireEnvironmentVariable("DATASET_NAME");
        String datasetVersion = requireEnvironmentVariable("DATASET_VERSION");
        Path inputPath = Path.of(requireEnvironmentVariable("DATA_FILE_PATH"));
        Path downloadPath = Path.of(requireEnvironmentVariable("DOWNLOAD_FILE_PATH"));
        byte[] expectedBytes = Files.readAllBytes(inputPath);

        DatasetsClient datasets = new AIProjectClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildDatasetsClient();
        boolean uploaded = false;

        try {
            datasets.createDatasetWithFile(datasetName, datasetVersion, inputPath);
            uploaded = true;

            DatasetVersion dataset =
                datasets.getDatasetVersion(datasetName, datasetVersion);
            System.out.printf(
                "Dataset: name=%s version=%s id=%s type=%s dataUri=%s%n",
                dataset.getName(),
                dataset.getVersion(),
                dataset.getId(),
                dataset.getType(),
                dataset.getDataUrl());

            DatasetCredential credential =
                datasets.getCredentials(datasetName, datasetVersion);
            BlobReference reference = credential.getBlobReference();
            URI blobUri = URI.create(reference.getBlobUrl());
            URI sasUri = URI.create(reference.getCredential().getSasUrl());
            URI authorizedBlobUri = new URI(
                sasUri.getScheme(),
                sasUri.getAuthority(),
                blobUri.getPath(),
                sasUri.getQuery(),
                null);
            BlobClient blobClient = new BlobClientBuilder()
                .endpoint(authorizedBlobUri.toString())
                .buildClient();

            Path parent = downloadPath.toAbsolutePath().getParent();
            if (parent != null) {
                Files.createDirectories(parent);
            }
            blobClient.downloadToFile(downloadPath.toString(), true);

            byte[] downloadedBytes = Files.readAllBytes(downloadPath);
            if (!Arrays.equals(expectedBytes, downloadedBytes)) {
                throw new IllegalStateException(
                    "Downloaded bytes don't match the source file.");
            }
            System.out.println(
                "Downloaded bytes verified: " + downloadedBytes.length);
        } finally {
            if (uploaded) {
                datasets.deleteDatasetVersion(datasetName, datasetVersion);
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
