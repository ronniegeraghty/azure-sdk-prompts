package com.example.orders.model;

import java.math.BigDecimal;
import java.util.Objects;

public class Order {
    private String orderId;
    private String customerName;
    private String product;
    private int quantity;
    private BigDecimal totalPrice;
    private OrderStatus status;

    public Order() {
    }

    public Order(
        String orderId,
        String customerName,
        String product,
        int quantity,
        BigDecimal totalPrice,
        OrderStatus status
    ) {
        this.orderId = requireText(orderId, "orderId");
        this.customerName = requireText(customerName, "customerName");
        this.product = requireText(product, "product");
        if (quantity <= 0) {
            throw new IllegalArgumentException("quantity must be positive");
        }
        this.quantity = quantity;
        this.totalPrice = Objects.requireNonNull(totalPrice, "totalPrice");
        if (totalPrice.signum() < 0) {
            throw new IllegalArgumentException("totalPrice must not be negative");
        }
        this.status = Objects.requireNonNull(status, "status");
    }

    private static String requireText(String value, String fieldName) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(fieldName + " must not be blank");
        }
        return value;
    }

    public String getOrderId() {
        return orderId;
    }

    public void setOrderId(String orderId) {
        this.orderId = orderId;
    }

    public String getCustomerName() {
        return customerName;
    }

    public void setCustomerName(String customerName) {
        this.customerName = customerName;
    }

    public String getProduct() {
        return product;
    }

    public void setProduct(String product) {
        this.product = product;
    }

    public int getQuantity() {
        return quantity;
    }

    public void setQuantity(int quantity) {
        this.quantity = quantity;
    }

    public BigDecimal getTotalPrice() {
        return totalPrice;
    }

    public void setTotalPrice(BigDecimal totalPrice) {
        this.totalPrice = totalPrice;
    }

    public OrderStatus getStatus() {
        return status;
    }

    public void setStatus(OrderStatus status) {
        this.status = status;
    }

    @Override
    public String toString() {
        return "Order{"
            + "orderId='" + orderId + '\''
            + ", customerName='" + customerName + '\''
            + ", product='" + product + '\''
            + ", quantity=" + quantity
            + ", totalPrice=" + totalPrice
            + ", status=" + status
            + '}';
    }
}
