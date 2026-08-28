package com.example.orders;

import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.function.Supplier;

final class FluxSupport {
    private FluxSupport() {
    }

    static <T> Flux<T> repeatSequentially(int count, Supplier<Mono<T>> operation) {
        return Flux.range(0, count).concatMap(ignored -> Mono.defer(operation));
    }
}
