package com.example.appconfig;

import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncConfigurationService implements AsyncConfigurationReader {
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationAsyncClient client;
    private final Map<SettingKey, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixKey, Map<String, String>> prefixCache = new ConcurrentHashMap<>();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = client;
    }

    public Mono<String> getSetting(String key) {
        return getSetting(key, null);
    }

    @Override
    public Mono<String> getSetting(String key, String label) {
        return Mono.defer(() -> {
            SettingKey cacheKey = new SettingKey(key, label);
            ConfigurationSetting cached = settingCache.get(cacheKey);
            if (cached == null) {
                return client.getConfigurationSetting(key, label)
                    .doOnNext(setting -> settingCache.put(cacheKey, setting))
                    .map(ConfigurationSetting::getValue);
            }

            return client.getConfigurationSettingWithResponse(cached, null, true)
                .map(response -> {
                    if (response.getStatusCode() == 304) {
                        return cached.getValue();
                    }
                    ConfigurationSetting updated = response.getValue();
                    settingCache.put(cacheKey, updated);
                    return updated.getValue();
                });
        });
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    @Override
    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        return Mono.defer(() -> {
            PrefixKey cacheKey = new PrefixKey(keyPrefix, label);
            Map<String, String> cached = prefixCache.get(cacheKey);
            return cached == null ? loadPrefix(cacheKey) : Mono.just(cached);
        });
    }

    public Mono<Void> refreshAll() {
        return Mono.defer(() -> {
            List<SettingKey> settings = List.copyOf(settingCache.keySet());
            List<PrefixKey> prefixes = List.copyOf(prefixCache.keySet());
            settingCache.clear();
            prefixCache.clear();

            Mono<Void> reloadSettings = Flux.fromIterable(settings)
                .concatMap(key -> getSetting(key.key(), key.label()))
                .then();
            Mono<Void> reloadPrefixes = Flux.fromIterable(prefixes)
                .concatMap(key -> listSettings(key.prefix(), key.label()))
                .then();
            return reloadSettings.then(reloadPrefixes);
        });
    }

    private Mono<Map<String, String>> loadPrefix(PrefixKey key) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(escapeFilter(key.prefix()) + "*")
            .setLabelFilter(key.label() == null ? NULL_LABEL_FILTER : key.label());
        return client.listConfigurationSettings(selector)
            .doOnNext(setting -> settingCache.put(
                new SettingKey(setting.getKey(), setting.getLabel()), setting))
            .collectMap(ConfigurationSetting::getKey, ConfigurationSetting::getValue)
            .map(Map::copyOf)
            .doOnNext(values -> prefixCache.put(key, values));
    }

    private static String escapeFilter(String value) {
        return value.replace("\\", "\\\\").replace(",", "\\,").replace("*", "\\*");
    }

    private record SettingKey(String key, String label) {
    }

    private record PrefixKey(String prefix, String label) {
    }
}
