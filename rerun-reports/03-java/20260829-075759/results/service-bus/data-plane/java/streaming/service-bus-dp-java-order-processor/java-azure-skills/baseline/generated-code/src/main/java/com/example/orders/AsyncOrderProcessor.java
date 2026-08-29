package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusErrorContext;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
import com.azure.messaging.servicebus.models.DeadLetterOptions;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Objects;

public final class AsyncOrderProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncOrderProcessor.class);

    private final ServiceBusProcessorClient processor;
    private final ObjectMapper objectMapper;

    public AsyncOrderProcessor(ServiceBusClientBuilder clientBuilder, String queueName, ObjectMapper objectMapper) {
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
        this.processor = Objects.requireNonNull(clientBuilder, "clientBuilder")
                .sessionProcessor()
                .queueName(Objects.requireNonNull(queueName, "queueName"))
                .maxConcurrentSessions(1)
                .maxConcurrentCalls(1)
                .disableAutoComplete()
                .processMessage(this::process)
                .processError(this::processError)
                .buildProcessorClient();
    }

    public void processFor(Duration duration) {
        processor.start();
        try {
            Thread.sleep(duration.toMillis());
        } catch (InterruptedException exception) {
            Thread.currentThread().interrupt();
            throw new IllegalStateException("Interrupted while processing orders", exception);
        } finally {
            processor.stop();
        }
    }

    private void process(ServiceBusReceivedMessageContext context) {
        try {
            Order order = objectMapper.readValue(context.getMessage().getBody().toBytes(), Order.class);
            LOGGER.info("Asynchronously processed {}", order);
            context.complete();
        } catch (Exception exception) {
            String description = OrderProcessor.rootMessage(exception);
            LOGGER.error("Dead-lettering message {}: {}", context.getMessage().getMessageId(), description);
            context.deadLetter(new DeadLetterOptions()
                    .setDeadLetterReason("OrderProcessingFailed")
                    .setDeadLetterErrorDescription(description));
        }
    }

    private void processError(ServiceBusErrorContext context) {
        LOGGER.error("Service Bus processor error for namespace {} and entity {}: {}",
                context.getFullyQualifiedNamespace(),
                context.getEntityPath(),
                context.getException().getMessage(),
                context.getException());
    }

    @Override
    public void close() {
        processor.close();
    }
}
