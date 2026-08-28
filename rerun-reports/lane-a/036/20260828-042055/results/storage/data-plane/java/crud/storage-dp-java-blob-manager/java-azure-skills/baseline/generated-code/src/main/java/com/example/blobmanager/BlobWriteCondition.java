package com.example.blobmanager;

import com.azure.storage.blob.models.BlobRequestConditions;

public record BlobWriteCondition(String expectedETag, String leaseId, boolean requireNewBlob) {
    public BlobWriteCondition {
        int selectedConditions = (expectedETag == null ? 0 : 1)
                + (leaseId == null ? 0 : 1)
                + (requireNewBlob ? 1 : 0);
        if (selectedConditions != 1) {
            throw new IllegalArgumentException(
                    "Exactly one of expectedETag, leaseId, or createOnly must be supplied");
        }
    }

    public static BlobWriteCondition createOnly() {
        return new BlobWriteCondition(null, null, true);
    }

    public static BlobWriteCondition ifUnchanged(String expectedETag) {
        return new BlobWriteCondition(requireText(expectedETag, "expectedETag"), null, false);
    }

    public static BlobWriteCondition withLease(String leaseId) {
        return new BlobWriteCondition(null, requireText(leaseId, "leaseId"), false);
    }

    BlobRequestConditions toRequestConditions() {
        BlobRequestConditions conditions = new BlobRequestConditions();
        if (requireNewBlob) {
            conditions.setIfNoneMatch("*");
        } else if (expectedETag != null) {
            conditions.setIfMatch(expectedETag);
        } else {
            conditions.setLeaseId(leaseId);
        }
        return conditions;
    }

    private static String requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }
}
