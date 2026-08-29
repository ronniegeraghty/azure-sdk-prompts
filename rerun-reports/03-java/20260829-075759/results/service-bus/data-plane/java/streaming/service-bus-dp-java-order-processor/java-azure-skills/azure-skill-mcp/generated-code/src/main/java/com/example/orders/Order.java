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
        Objects.requireNonNull(orderId, "orderId");
        Objects.requireNonNull(customerName, "customerName");
        Objects.requireNonNull(product, "product");
        Objects.requireNonNull(totalPrice, "totalPrice");
        Objects.requireNonNull(status, "status");
        if (orderId.isBlank() || customerName.isBlank() || product.isBlank()) {
            throw new IllegalArgumentException("Order ID, customer name, and product must not be blank");
        }
        if (quantity <= 0) {
            throw new IllegalArgumentException("Quantity must be greater than zero");
        }
        if (totalPrice.signum() < 0) {
            throw new IllegalArgumentException("Total price must not be negative");
        }
    }

    public Order withStatus(OrderStatus newStatus) {
        return new Order(orderId, customerName, product, quantity, totalPrice, newStatus);
    }
}
