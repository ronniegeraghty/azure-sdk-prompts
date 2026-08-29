package com.example.appconfig;

record SelectorId(String prefix, String label) {
    SelectorId {
        if (prefix == null || prefix.isBlank()) {
            throw new IllegalArgumentException("Configuration key prefix must not be blank");
        }
    }
}
