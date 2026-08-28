package com.example.azureidentity;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertFalse;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.junit.jupiter.api.Assertions.assertTrue;

class CredentialConnectivityTesterTest {
    @Test
    void setsCaeOnTokenRequest() {
        assertTrue(CredentialConnectivityTester.request("https://management.azure.com/.default", true)
            .isCaeEnabled());
        assertFalse(CredentialConnectivityTester.request("https://management.azure.com/.default", false)
            .isCaeEnabled());
    }

    @Test
    void rejectsBlankScope() {
        assertThrows(
            IllegalArgumentException.class,
            () -> CredentialConnectivityTester.request(" ", true)
        );
    }
}
