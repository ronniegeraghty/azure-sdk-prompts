package com.example.orders.model;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;

import java.util.Locale;

public enum OrderStatus {
    PENDING,
    PROCESSING,
    COMPLETED,
    FAILED;

    @JsonCreator
    public static OrderStatus fromJson(String value) {
        return OrderStatus.valueOf(value.toUpperCase(Locale.ROOT));
    }

    @JsonValue
    public String toJson() {
        return name().toLowerCase(Locale.ROOT);
    }
}
