package com.example.appconfig;

import com.azure.core.http.rest.Response;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncConfigurationService {
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationAsyncClient client;
    private final Map<SettingQuery, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixQuery, Map<String, String>> prefixCache = new ConcurrentHashMap<>();
    private final Set<SettingQuery> trackedSettings = ConcurrentHashMap.newKeySet();
    private final Set<PrefixQuery> trackedPrefixes = ConcurrentHashMap.newKeySet();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<String> getSetting(String key) {
        return getSetting(key, null);
    }

    public Mono<String> getSetting(String key, String label) {
        SettingQuery query = new SettingQuery(requireText(key, "key"), label);
        trackedSettings.add(query);
        return fetchSetting(query).map(ConfigurationSetting::getValue);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        PrefixQuery query = new PrefixQuery(requireText(keyPrefix, "keyPrefix"), label);
        trackedPrefixes.add(query);
        Map<String, String> cached = prefixCache.get(query);
        return cached != null ? Mono.just(cached) : refreshPrefix(query);
    }

    public Mono<Void> refreshAll() {
        Mono<Void> settingsRefresh = Flux.fromIterable(trackedSettings)
            .concatMap(this::fetchSetting)
            .then();
        Mono<Void> prefixesRefresh = Flux.fromIterable(trackedPrefixes)
            .concatMap(this::refreshPrefix)
            .then();
        return settingsRefresh.then(prefixesRefresh);
    }

    private Mono<ConfigurationSetting> fetchSetting(SettingQuery query) {
        ConfigurationSetting cached = settingCache.get(query);
        Mono<Response<ConfigurationSetting>> request = cached == null
            ? client.getConfigurationSettingWithResponse(
                new ConfigurationSetting().setKey(query.key()).setLabel(query.label()), null, false)
            : client.getConfigurationSettingWithResponse(cached, null, true);

        return request.map(response -> {
            if (response.getStatusCode() == 304) {
                return cached;
            }
            ConfigurationSetting updated = response.getValue();
            settingCache.put(query, updated);
            return updated;
        });
    }

    private Mono<Map<String, String>> refreshPrefix(PrefixQuery query) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(escapeFilter(query.prefix()) + "*")
            .setLabelFilter(query.label() == null ? NULL_LABEL_FILTER : escapeFilter(query.label()));

        return client.listConfigurationSettings(selector)
            .collect(
                LinkedHashMap<String, String>::new,
                (loaded, setting) -> {
                    loaded.put(setting.getKey(), setting.getValue());
                    settingCache.put(new SettingQuery(setting.getKey(), setting.getLabel()), setting);
                })
            .map(loaded -> {
                Map<String, String> snapshot = Collections.unmodifiableMap(loaded);
                prefixCache.put(query, snapshot);
                return snapshot;
            });
    }

    private static String escapeFilter(String value) {
        return value.replace("\\", "\\\\").replace("*", "\\*").replace(",", "\\,");
    }

    private static String requireText(String value, String name) {
        if (value == null || value.isBlank()) {
            throw new IllegalArgumentException(name + " must not be blank");
        }
        return value;
    }

    private record SettingQuery(String key, String label) {
    }

    private record PrefixQuery(String prefix, String label) {
    }
}
