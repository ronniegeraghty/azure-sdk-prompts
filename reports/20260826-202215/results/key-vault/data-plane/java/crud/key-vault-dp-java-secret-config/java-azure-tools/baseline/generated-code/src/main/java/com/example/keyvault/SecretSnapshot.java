package com.example.keyvault;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;
import java.util.Optional;

public record SecretSnapshot(
        String name,
        String value,
        String version,
        OffsetDateTime expiresOn,
        boolean defaultValue) {

    public SecretSnapshot {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(value, "value");
    }

    public static SecretSnapshot missing(String name, String defaultValue) {
        return new SecretSnapshot(name, defaultValue, null, null, true);
    }

    public Optional<OffsetDateTime> expiry() {
        return Optional.ofNullable(expiresOn);
    }

    public boolean expiresWithin(Duration warningWindow, OffsetDateTime now) {
        Objects.requireNonNull(warningWindow, "warningWindow");
        Objects.requireNonNull(now, "now");
        if (warningWindow.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return expiresOn != null && !expiresOn.isAfter(now.plus(warningWindow));
    }
}
