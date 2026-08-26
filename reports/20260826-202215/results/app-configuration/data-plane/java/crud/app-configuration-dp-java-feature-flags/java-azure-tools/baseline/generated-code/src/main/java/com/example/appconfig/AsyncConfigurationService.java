package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public class AsyncConfigurationService {
    private final ConfigurationAsyncClient client;
    private final Map<CacheKey, ConfigurationSetting> cache = new ConcurrentHashMap<>();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<Optional<String>> getSetting(String key) {
        return getSetting(key, null);
    }

    public Mono<Optional<String>> getSetting(String key, String label) {
        Objects.requireNonNull(key, "key");
        return Mono.defer(() -> {
            CacheKey cacheKey = new CacheKey(key, label);
            ConfigurationSetting cached = cache.get(cacheKey);
            ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);
            if (cached != null) {
                request.setETag(cached.getETag());
            }

            return client.getConfigurationSettingWithResponse(request, null, cached != null)
                .map(response -> {
                    if (response.getStatusCode() == 304) {
                        return Optional.ofNullable(cached.getValue());
                    }

                    ConfigurationSetting current = response.getValue();
                    cache.put(cacheKey, current);
                    return Optional.ofNullable(current.getValue());
                })
                .onErrorResume(ResourceNotFoundException.class, exception -> {
                    cache.remove(cacheKey);
                    return Mono.just(Optional.empty());
                });
        });
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        Objects.requireNonNull(keyPrefix, "keyPrefix");
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(keyPrefix + "*")
            .setLabelFilter(label == null ? ConfigurationSetting.NO_LABEL : label);

        return client.listConfigurationSettings(selector)
            .collect(LinkedHashMap<String, String>::new, (result, setting) -> {
                result.put(setting.getKey(), setting.getValue());
                cache.put(new CacheKey(setting.getKey(), setting.getLabel()), setting);
            })
            .map(Map::copyOf);
    }

    public Mono<Void> refreshAll() {
        return Mono.defer(() -> Flux.fromIterable(List.copyOf(cache.keySet()))
            .concatMap(cachedKey -> getSetting(cachedKey.key(), cachedKey.label()))
            .then());
    }

    private record CacheKey(String key, String label) {
    }
}
