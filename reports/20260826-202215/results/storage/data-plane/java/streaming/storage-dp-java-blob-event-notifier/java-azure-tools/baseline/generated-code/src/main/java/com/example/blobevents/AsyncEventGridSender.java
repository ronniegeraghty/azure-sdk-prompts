package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;
import reactor.core.publisher.Mono;

import java.util.List;

@FunctionalInterface
public interface AsyncEventGridSender {
    Mono<Void> send(List<EventGridEvent> events);
}
