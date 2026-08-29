package com.example.appconfig;

import java.util.Objects;

public record Sentinel(String key, String label) {
    public Sentinel {
        Objects.requireNonNull(key, "key");
        if (key.isBlank()) {
            throw new IllegalArgumentException("Sentinel key must not be blank");
        }
    }
}
