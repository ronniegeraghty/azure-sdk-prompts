package com.example.appconfig;

final class ConfigurationFilters {
    private ConfigurationFilters() {
    }

    static String keyPrefix(String prefix) {
        return escape(prefix) + "*";
    }

    static String label(String label) {
        return label == null ? "\\0" : escape(label);
    }

    private static String escape(String value) {
        return value.replace("\\", "\\\\")
            .replace("*", "\\*")
            .replace(",", "\\,");
    }
}
