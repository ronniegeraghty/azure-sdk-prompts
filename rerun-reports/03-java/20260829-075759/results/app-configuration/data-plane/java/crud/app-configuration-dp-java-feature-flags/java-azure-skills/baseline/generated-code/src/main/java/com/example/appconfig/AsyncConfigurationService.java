package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

public final class AsyncConfigurationService {
    private final ConfigurationAsyncClient client;
    private final ConcurrentMap<SettingId, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final ConcurrentMap<SelectorId, Map<String, String>> prefixCache = new ConcurrentHashMap<>();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = java.util.Objects.requireNonNull(client, "client");
    }

    public Mono<Optional<String>> getSetting(String key) {
        return getSetting(key, null);
    }

    public Mono<Optional<String>> getSetting(String key, String label) {
        return Mono.defer(() -> {
            SettingId id = SettingId.of(key, label);
            ConfigurationSetting cached = settingCache.get(id);
            ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);
            if (cached != null) {
                request.setETag(cached.getETag());
            }

            return client.getConfigurationSettingWithResponse(request, null, cached != null)
                .map(response -> {
                    ConfigurationSetting resolved = response.getStatusCode() == 304 ? cached : response.getValue();
                    if (resolved != null) {
                        settingCache.put(id, resolved);
                    }
                    return Optional.ofNullable(resolved).map(ConfigurationSetting::getValue);
                })
                .onErrorResume(ResourceNotFoundException.class, exception -> {
                    settingCache.remove(id);
                    return Mono.just(Optional.empty());
                });
        });
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        return Mono.defer(() -> {
            SelectorId id = new SelectorId(keyPrefix, label);
            Map<String, String> cached = prefixCache.get(id);
            return cached == null ? loadSettings(id) : Mono.just(cached);
        });
    }

    public Mono<Boolean> hasSettingChanged(String key, String label) {
        return Mono.defer(() -> {
            SettingId id = SettingId.of(key, label);
            ConfigurationSetting cached = settingCache.get(id);
            if (cached == null) {
                return getSetting(key, label).thenReturn(false);
            }

            return client.getConfigurationSettingWithResponse(
                    new ConfigurationSetting()
                        .setKey(key)
                        .setLabel(label)
                        .setETag(cached.getETag()),
                    null,
                    true)
                .map(response -> {
                    if (response.getStatusCode() == 304) {
                        return false;
                    }
                    settingCache.put(id, response.getValue());
                    return true;
                })
                .onErrorResume(ResourceNotFoundException.class, exception -> {
                    settingCache.remove(id);
                    return Mono.just(true);
                });
        });
    }

    public Mono<Void> refreshAll() {
        return Mono.defer(() -> {
            List<SettingId> settings = new ArrayList<>(settingCache.keySet());
            List<SelectorId> selectors = new ArrayList<>(prefixCache.keySet());
            settingCache.clear();
            prefixCache.clear();

            Flux<?> settingRefreshes = Flux.fromIterable(settings)
                .concatMap(id -> getSetting(id.key(), id.label()));
            Flux<?> selectorRefreshes = Flux.fromIterable(selectors)
                .concatMap(id -> listSettings(id.prefix(), id.label()));
            return Flux.concat(settingRefreshes, selectorRefreshes).then();
        });
    }

    private Mono<Map<String, String>> loadSettings(SelectorId id) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(ConfigurationFilters.keyPrefix(id.prefix()))
            .setLabelFilter(ConfigurationFilters.label(id.label()));
        return client.listConfigurationSettings(selector)
            .collect(
                LinkedHashMap<String, String>::new,
                (settings, setting) -> settings.put(setting.getKey(), setting.getValue()))
            .map(Collections::unmodifiableMap)
            .doOnNext(settings -> prefixCache.put(id, settings));
    }
}
