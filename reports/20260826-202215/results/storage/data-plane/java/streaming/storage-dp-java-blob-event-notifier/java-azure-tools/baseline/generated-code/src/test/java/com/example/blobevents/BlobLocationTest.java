package com.example.blobevents;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class BlobLocationTest {
    @Test
    void parsesNestedAndEncodedBlobName() {
        BlobLocation location = BlobLocation.fromSubject(
                "/blobServices/default/containers/documents/blobs/invoices/2026/Invoice%20A%2BB.pdf");

        assertEquals("documents", location.container());
        assertEquals("invoices/2026/Invoice A+B.pdf", location.blobName());
    }

    @Test
    void rejectsInvalidSubject() {
        assertThrows(IllegalArgumentException.class, () -> BlobLocation.fromSubject("/not/a/blob"));
    }
}
