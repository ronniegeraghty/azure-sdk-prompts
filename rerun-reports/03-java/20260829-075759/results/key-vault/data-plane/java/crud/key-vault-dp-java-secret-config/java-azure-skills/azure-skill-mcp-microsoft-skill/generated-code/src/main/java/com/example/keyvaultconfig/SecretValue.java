package com.example.keyvaultconfig;

import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

public record SecretValue(
        String name,
        String value,
        String version,
        OffsetDateTime expiresOn,
        boolean defaultValue) {

    public SecretValue {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(value, "value");
    }

    public boolean expiresWithin(Duration warningWindow, OffsetDateTime now) {
        Objects.requireNonNull(warningWindow, "warningWindow");
        Objects.requireNonNull(now, "now");
        return expiresOn != null && !expiresOn.isAfter(now.plus(warningWindow));
    }
}
