package com.example.orders;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

import java.util.Locale;

public enum OrderStatus {
    PENDING,
    PROCESSING,
    COMPLETED,
    FAILED;

    @JsonValue
    public String toJson() {
        return name().toLowerCase(Locale.ROOT);
    }

    @JsonCreator
    public static OrderStatus fromJson(String value) {
        return valueOf(value.toUpperCase(Locale.ROOT));
    }
}
