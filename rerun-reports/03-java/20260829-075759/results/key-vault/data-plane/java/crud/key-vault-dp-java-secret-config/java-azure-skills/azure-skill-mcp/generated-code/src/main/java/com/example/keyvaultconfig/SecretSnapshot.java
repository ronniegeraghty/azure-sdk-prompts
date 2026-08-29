package com.example.keyvaultconfig;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

public record SecretSnapshot(
        String name,
        String value,
        String version,
        OffsetDateTime expiresOn,
        boolean found) {

    public SecretSnapshot {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(value, "value");
    }

    public boolean expiresWithin(Duration warningWindow, Clock clock) {
        Objects.requireNonNull(warningWindow, "warningWindow");
        Objects.requireNonNull(clock, "clock");
        if (warningWindow.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }
        return expiresOn != null
                && !expiresOn.isAfter(OffsetDateTime.now(clock).plus(warningWindow));
    }
}
