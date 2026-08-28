package com.example.blobmanager;

import com.azure.core.util.BinaryData;
import com.azure.storage.blob.models.BlobRequestConditions;
import com.azure.storage.blob.models.ParallelTransferOptions;
import com.azure.storage.blob.options.BlobParallelUploadOptions;

import java.nio.file.Path;
import java.util.Map;

final class BlobUploadOptionsFactory {
    private static final long BLOCK_SIZE = 8L * 1024 * 1024;
    private static final long SINGLE_UPLOAD_SIZE = 32L * 1024 * 1024;
    private static final int MAX_CONCURRENCY = 4;

    private BlobUploadOptionsFactory() {
    }

    static BlobParallelUploadOptions create(
            Path source,
            Map<String, String> metadata,
            Map<String, String> tags,
            String expectedETag,
            String leaseId
    ) {
        return new BlobParallelUploadOptions(BinaryData.fromFile(source))
                .setParallelTransferOptions(new ParallelTransferOptions()
                        .setBlockSizeLong(BLOCK_SIZE)
                        .setMaxSingleUploadSizeLong(SINGLE_UPLOAD_SIZE)
                        .setMaxConcurrency(MAX_CONCURRENCY))
                .setMetadata(metadata == null ? Map.of() : metadata)
                .setTags(tags == null ? Map.of() : tags)
                .setRequestConditions(writeConditions(expectedETag, leaseId));
    }

    private static BlobRequestConditions writeConditions(String expectedETag, String leaseId) {
        BlobRequestConditions conditions = new BlobRequestConditions();
        if (expectedETag == null || expectedETag.isBlank()) {
            conditions.setIfNoneMatch("*");
        } else {
            conditions.setIfMatch(expectedETag);
        }
        if (leaseId != null && !leaseId.isBlank()) {
            conditions.setLeaseId(leaseId);
        }
        return conditions;
    }
}
