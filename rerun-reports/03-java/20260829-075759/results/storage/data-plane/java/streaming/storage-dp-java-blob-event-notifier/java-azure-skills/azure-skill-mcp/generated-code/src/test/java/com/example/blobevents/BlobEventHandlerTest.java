package com.example.blobevents;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class BlobEventHandlerTest {
    @Test
    void parsesAndDecodesBlobSubject() {
        BlobEventHandler.BlobLocation location = BlobEventHandler.parseSubject(
            "/blobServices/default/containers/my-container/blobs/folder%20one/report%2Bfinal.pdf");

        assertEquals("my-container", location.container());
        assertEquals("folder one/report+final.pdf", location.name());
    }
}
