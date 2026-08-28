package com.example.blobevents;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

import com.example.blobevents.BlobEventHandler.BlobAddress;
import org.junit.jupiter.api.Test;

class BlobEventHandlerTest {
    @Test
    void parsesEncodedNestedBlobNameWithoutTreatingPlusAsSpace() {
        BlobAddress address = BlobEventHandler.parseAddress(
            "/blobServices/default/containers/documents/blobs/invoices%2F2026%2Ftotal+tax.pdf"
        );

        assertEquals("documents", address.container());
        assertEquals("invoices/2026/total+tax.pdf", address.blobName());
    }

    @Test
    void rejectsMalformedSubject() {
        assertThrows(
            IllegalArgumentException.class,
            () -> BlobEventHandler.parseAddress("/containers/documents")
        );
    }
}
