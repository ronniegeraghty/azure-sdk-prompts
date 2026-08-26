package com.example.blobevents;

import com.azure.messaging.eventgrid.EventGridEvent;

import java.util.List;

@FunctionalInterface
public interface EventGridSender {
    void send(List<EventGridEvent> events);
}
