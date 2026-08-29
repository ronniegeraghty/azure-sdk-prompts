package com.example.appconfig;

import reactor.core.publisher.Mono;

@FunctionalInterface
public interface AsyncSettingProvider {
    Mono<String> getSetting(String key, String label);
}
