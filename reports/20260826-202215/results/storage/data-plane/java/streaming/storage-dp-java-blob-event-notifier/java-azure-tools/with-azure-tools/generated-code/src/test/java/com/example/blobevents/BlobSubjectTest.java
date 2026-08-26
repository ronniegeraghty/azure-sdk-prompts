package com.example.blobevents;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertThrows;

class BlobSubjectTest {
    @Test
    void parsesContainerAndNestedBlobName() {
        BlobSubject subject = BlobSubject.parse(
            "/blobServices/default/containers/documents/blobs/invoices/2026/invoice%201042.pdf");

        assertEquals("documents", subject.container());
        assertEquals("invoices/2026/invoice 1042.pdf", subject.blobName());
    }

    @Test
    void preservesPlusCharacters() {
        BlobSubject subject = BlobSubject.parse(
            "/blobServices/default/containers/documents/blobs/a+b.txt");

        assertEquals("a+b.txt", subject.blobName());
    }

    @Test
    void rejectsNonBlobSubject() {
        assertThrows(IllegalArgumentException.class, () -> BlobSubject.parse("/not/a/blob"));
    }
}
