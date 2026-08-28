package com.example.orders;

import java.math.BigDecimal;
import java.util.Objects;

public record Order(
        String orderId,
        String customerName,
        String product,
        int quantity,
        BigDecimal totalPrice,
        OrderStatus status) {

    public Order {
        if (orderId == null || orderId.isBlank()) {
            throw new IllegalArgumentException("orderId is required");
        }
        if (customerName == null || customerName.isBlank()) {
            throw new IllegalArgumentException("customerName is required");
        }
        if (product == null || product.isBlank()) {
            throw new IllegalArgumentException("product is required");
        }
        if (quantity <= 0) {
            throw new IllegalArgumentException("quantity must be positive");
        }
        Objects.requireNonNull(totalPrice, "totalPrice is required");
        if (totalPrice.signum() < 0) {
            throw new IllegalArgumentException("totalPrice cannot be negative");
        }
        Objects.requireNonNull(status, "status is required");
    }
}
