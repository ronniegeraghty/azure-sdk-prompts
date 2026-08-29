package com.example.orders;

import com.azure.messaging.servicebus.ServiceBusClientBuilder;
import com.azure.messaging.servicebus.ServiceBusErrorContext;
import com.azure.messaging.servicebus.ServiceBusProcessorClient;
import com.azure.messaging.servicebus.ServiceBusReceivedMessageContext;
import com.azure.messaging.servicebus.models.SubQueue;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Objects;
import java.util.function.Function;

public final class AsyncDeadLetterQueueProcessor implements AutoCloseable {
    private static final Logger LOGGER = LoggerFactory.getLogger(AsyncDeadLetterQueueProcessor.class);

    private final ServiceBusProcessorClient processor;
    private final ObjectMapper objectMapper;
    private final Function<Order, reactor.core.publisher.Mono<Void>> resend;

    public AsyncDeadLetterQueueProcessor(
            ServiceBusClientBuilder clientBuilder,
            String queueName,
            ObjectMapper objectMapper,
            Function<Order, reactor.core.publisher.Mono<Void>> resend) {
        this.objectMapper = Objects.requireNonNull(objectMapper, "objectMapper");
        this.resend = resend;
        this.processor = Objects.requireNonNull(clientBuilder, "clientBuilder")
                .sessionProcessor()
                .queueName(Objects.requireNonNull(queueName, "queueName"))
                .subQueue(SubQueue.DEAD_LETTER_QUEUE)
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
            throw new IllegalStateException("Interrupted while reading the dead-letter queue", exception);
        } finally {
            processor.stop();
        }
    }

    private void process(ServiceBusReceivedMessageContext context) {
        LOGGER.info("Dead-letter message id={}, reason={}, description={}, body={}",
                context.getMessage().getMessageId(),
                context.getMessage().getDeadLetterReason(),
                context.getMessage().getDeadLetterErrorDescription(),
                context.getMessage().getBody());

        if (resend == null) {
            context.abandon();
            return;
        }

        try {
            Order order = objectMapper.readValue(context.getMessage().getBody().toBytes(), Order.class);
            resend.apply(order).block();
            context.complete();
            LOGGER.info("Re-enqueued dead-lettered order {}", order.getOrderId());
        } catch (Exception exception) {
            LOGGER.error("Could not reprocess dead-letter message {}: {}",
                    context.getMessage().getMessageId(), OrderProcessor.rootMessage(exception));
            context.abandon();
        }
    }

    private void processError(ServiceBusErrorContext context) {
        LOGGER.error("Dead-letter processor error for entity {}: {}",
                context.getEntityPath(), context.getException().getMessage(), context.getException());
    }

    @Override
    public void close() {
        processor.close();
    }
}
