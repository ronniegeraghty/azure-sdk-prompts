package com.example.orders;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonValue;

import java.math.BigDecimal;
import java.util.Objects;

public final class Order {
    public enum Status {
        PENDING,
        PROCESSING,
        COMPLETED,
        FAILED;

        @JsonCreator
        public static Status fromJson(String value) {
            return Status.valueOf(value.toUpperCase(java.util.Locale.ROOT));
        }

        @JsonValue
        public String toJson() {
            return name().toLowerCase(java.util.Locale.ROOT);
        }
    }

    private final String orderId;
    private final String customerName;
    private final String product;
    private final int quantity;
    private final BigDecimal totalPrice;
    private final Status status;

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

    public Order withStatus(Status newStatus) {
        return new Order(orderId, customerName, product, quantity, totalPrice, newStatus);
    }

    @Override
    public String toString() {
        return "Order{" +
                "orderId='" + orderId + '\'' +
                ", customerName='" + customerName + '\'' +
                ", product='" + product + '\'' +
                ", quantity=" + quantity +
                ", totalPrice=" + totalPrice +
                ", status=" + status +
                '}';
    }
}
