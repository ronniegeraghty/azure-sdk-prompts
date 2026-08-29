package com.example.orders;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;

import java.math.BigDecimal;
import java.util.Locale;
import java.util.Objects;

public final class Order {
    public enum Status {
        PENDING,
        PROCESSING,
        COMPLETED,
        FAILED;

        @JsonCreator
        public static Status fromJson(String value) {
            return Status.valueOf(value.toUpperCase(Locale.ROOT));
        }

        @JsonValue
        public String toJson() {
            return name().toLowerCase(Locale.ROOT);
        }
    }

    private final String orderId;
    private final String customerName;
    private final String product;
    private final int quantity;
    private final BigDecimal totalPrice;
    private Status status;

    @JsonCreator
    public Order(
            @JsonProperty("orderId") String orderId,
            @JsonProperty("customerName") String customerName,
            @JsonProperty("product") String product,
            @JsonProperty("quantity") int quantity,
            @JsonProperty("totalPrice") BigDecimal totalPrice,
            @JsonProperty("status") Status status) {
        this.orderId = Objects.requireNonNull(orderId, "orderId");
        this.customerName = Objects.requireNonNull(customerName, "customerName");
        this.product = Objects.requireNonNull(product, "product");
        if (quantity <= 0) {
            throw new IllegalArgumentException("quantity must be positive");
        }
        this.quantity = quantity;
        this.totalPrice = Objects.requireNonNull(totalPrice, "totalPrice");
        this.status = Objects.requireNonNull(status, "status");
    }

    public String getOrderId() {
        return orderId;
    }

    public String getCustomerName() {
        return customerName;
    }

    public String getProduct() {
        return product;
    }

    public int getQuantity() {
        return quantity;
    }

    public BigDecimal getTotalPrice() {
        return totalPrice;
    }

    public Status getStatus() {
        return status;
    }

    public void setStatus(Status status) {
        this.status = Objects.requireNonNull(status, "status");
    }

    @Override
    public String toString() {
        return "Order{orderId='%s', customerName='%s', product='%s', quantity=%d, totalPrice=%s, status=%s}"
                .formatted(orderId, customerName, product, quantity, totalPrice, status);
    }
}
