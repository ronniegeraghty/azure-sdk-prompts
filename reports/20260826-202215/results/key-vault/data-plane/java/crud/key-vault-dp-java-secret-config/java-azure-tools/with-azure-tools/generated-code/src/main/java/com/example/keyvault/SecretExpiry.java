package com.example.keyvault;

import java.time.Clock;
import java.time.Duration;
import java.time.OffsetDateTime;
import java.util.Objects;

final class SecretExpiry {
    private SecretExpiry() {
    }

    static boolean isWithin(SecretSnapshot secret, Duration warningWindow, Clock clock) {
        Objects.requireNonNull(secret, "secret");
        Objects.requireNonNull(warningWindow, "warningWindow");
        Objects.requireNonNull(clock, "clock");
        if (warningWindow.isNegative()) {
            throw new IllegalArgumentException("warningWindow must not be negative");
        }

        OffsetDateTime expiresOn = secret.expiresOn();
        return expiresOn != null
            && !expiresOn.isAfter(OffsetDateTime.now(clock).plus(warningWindow));
    }
}
