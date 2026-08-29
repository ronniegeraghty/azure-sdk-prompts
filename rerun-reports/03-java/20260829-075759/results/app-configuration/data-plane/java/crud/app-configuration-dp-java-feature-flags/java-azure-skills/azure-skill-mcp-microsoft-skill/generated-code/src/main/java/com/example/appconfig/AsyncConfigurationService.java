package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncConfigurationService implements AsyncSettingProvider {
    private static final int NOT_MODIFIED = 304;
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationAsyncClient client;
    private final Map<SettingId, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixQuery, Map<String, String>> prefixCache = new ConcurrentHashMap<>();
    private final Map<SettingId, Optional<String>> sentinelValues = new ConcurrentHashMap<>();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<String> getSetting(String key) {
        return getSetting(key, null);
    }

    @Override
    public Mono<String> getSetting(String key, String label) {
        return Mono.defer(() -> {
            SettingId id = new SettingId(key, label);
            ConfigurationSetting cached = settingCache.get(id);
            if (cached == null) {
                return loadSetting(id).map(ConfigurationSetting::getValue);
            }

            return client.getConfigurationSettingWithResponse(cached, null, true)
                .flatMap(response -> {
                    if (response.getStatusCode() == NOT_MODIFIED) {
                        return Mono.justOrEmpty(cached.getValue());
                    }
                    ConfigurationSetting updated = response.getValue();
                    settingCache.put(id, updated);
                    return Mono.justOrEmpty(updated.getValue());
                })
                .onErrorResume(ResourceNotFoundException.class, exception -> {
                    settingCache.remove(id);
                    return Mono.empty();
                });
        });
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        return Mono.defer(() -> {
            PrefixQuery query = new PrefixQuery(keyPrefix, label);
            Map<String, String> cached = prefixCache.get(query);
            if (cached != null) {
                return Mono.just(cached);
            }
            return loadPrefix(query).doOnNext(values -> prefixCache.put(query, values));
        });
    }

    Mono<Boolean> checkForUpdate(Sentinel sentinel) {
        return Mono.defer(() -> {
            SettingId id = new SettingId(sentinel.key(), sentinel.label());
            return getSetting(id.key(), id.label())
                .map(Optional::of)
                .defaultIfEmpty(Optional.empty())
                .map(current -> {
                    Optional<String> previous = sentinelValues.put(id, current);
                    return previous != null && !previous.equals(current);
                });
        });
    }

    public Mono<Void> refreshAll() {
        return Mono.defer(() -> {
            Flux<Void> settings = Flux.fromIterable(settingCache.keySet())
                .concatMap(id -> loadSetting(id).then());
            Flux<Void> prefixes = Flux.fromIterable(prefixCache.keySet())
                .concatMap(query -> loadPrefix(query)
                    .doOnNext(values -> prefixCache.put(query, values))
                    .then());
            return Flux.concat(settings, prefixes).then();
        });
    }

    private Mono<ConfigurationSetting> loadSetting(SettingId id) {
        return client.getConfigurationSetting(id.key(), id.label())
            .doOnNext(setting -> settingCache.put(id, setting))
            .onErrorResume(ResourceNotFoundException.class, exception -> {
                settingCache.remove(id);
                return Mono.empty();
            });
    }

    private Mono<Map<String, String>> loadPrefix(PrefixQuery query) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(query.prefix() + "*")
            .setLabelFilter(query.label() == null ? NULL_LABEL_FILTER : query.label());

        return client.listConfigurationSettings(selector)
            .collect(
                LinkedHashMap<String, String>::new,
                (values, setting) -> {
                    values.put(setting.getKey(), setting.getValue());
                    settingCache.put(new SettingId(setting.getKey(), setting.getLabel()), setting);
                })
            .map(Map::copyOf);
    }

    private record SettingId(String key, String label) {
        private SettingId {
            Objects.requireNonNull(key, "key");
            if (key.isBlank()) {
                throw new IllegalArgumentException("Setting key must not be blank");
            }
        }
    }

    private record PrefixQuery(String prefix, String label) {
        private PrefixQuery {
            Objects.requireNonNull(prefix, "prefix");
        }
    }
}
