package com.example.blobevents;

import com.azure.core.credential.TokenCredential;
import com.azure.core.util.BinaryData;
import com.azure.messaging.eventgrid.EventGridEvent;
import com.azure.messaging.eventgrid.EventGridPublisherClient;
import com.azure.messaging.eventgrid.EventGridPublisherClientBuilder;

import java.util.List;

public final class EventPublisher {
    private final EventGridSender sender;

    public EventPublisher(String topicEndpoint, TokenCredential credential) {
        EventGridPublisherClient<EventGridEvent> client = new EventGridPublisherClientBuilder()
                .endpoint(topicEndpoint)
                .credential(credential)
                .buildEventGridEventPublisherClient();
        this.sender = client::sendEvents;
    }

    public EventPublisher(EventGridSender sender) {
        this.sender = sender;
    }

    public void publish(List<CustomEvent> events) {
        sender.send(events.stream().map(EventPublisher::toEventGridEvent).toList());
    }

    static EventGridEvent toEventGridEvent(CustomEvent event) {
        return new EventGridEvent(
                event.subject(),
                event.eventType(),
                BinaryData.fromObject(event.data()),
                event.dataVersion());
    }
}
