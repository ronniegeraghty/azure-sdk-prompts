package com.example.appconfig;

import reactor.core.publisher.Mono;

import java.util.Map;

public interface AsyncConfigurationReader {
    Mono<String> getSetting(String key, String label);

    Mono<Map<String, String>> listSettings(String keyPrefix, String label);
}
