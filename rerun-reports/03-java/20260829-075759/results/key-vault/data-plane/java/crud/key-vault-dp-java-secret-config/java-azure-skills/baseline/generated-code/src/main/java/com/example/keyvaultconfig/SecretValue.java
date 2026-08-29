package com.example.keyvaultconfig;

import java.time.OffsetDateTime;
import java.util.Objects;
import java.util.Optional;

public record SecretValue(String name, String value, OffsetDateTime expiresOn) {
    public SecretValue {
        Objects.requireNonNull(name, "name");
        Objects.requireNonNull(value, "value");
    }

    public Optional<OffsetDateTime> expiry() {
        return Optional.ofNullable(expiresOn);
    }
}
