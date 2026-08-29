package com.example.appconfig;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class ConfigurationFiltersTest {
    @Test
    void noLabelUsesAzureNullLabelFilter() {
        assertEquals("\\0", ConfigurationFilters.label(null));
    }

    @Test
    void prefixEscapesReservedFilterCharacters() {
        assertEquals("Demo\\*\\,Path\\\\Name*", ConfigurationFilters.keyPrefix("Demo*,Path\\Name"));
    }
}
