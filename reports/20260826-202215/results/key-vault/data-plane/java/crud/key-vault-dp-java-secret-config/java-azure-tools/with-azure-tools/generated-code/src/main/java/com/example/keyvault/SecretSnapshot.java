package com.example.keyvault;

import java.time.OffsetDateTime;
import java.util.Objects;

public record SecretSnapshot(
    String name,
    String value,
    String version,
    OffsetDateTime expiresOn,
    boolean defaultValue
) {
    public SecretSnapshot {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(value, "value");
    }
}
