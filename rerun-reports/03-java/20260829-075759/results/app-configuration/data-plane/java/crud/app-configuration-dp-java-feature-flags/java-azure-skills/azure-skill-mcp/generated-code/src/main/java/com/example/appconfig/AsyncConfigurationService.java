package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.http.HttpHeaderName;
import com.azure.core.http.MatchConditions;
import com.azure.core.http.rest.PagedResponse;
import com.azure.data.appconfiguration.ConfigurationAsyncClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingFields;
import com.azure.data.appconfiguration.models.SettingSelector;
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public final class AsyncConfigurationService {
    private final ConfigurationAsyncClient client;
    private final Map<SettingKey, CachedSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixQuery, CachedPrefix> prefixCache = new ConcurrentHashMap<>();

    public AsyncConfigurationService(ConfigurationAsyncClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Mono<Optional<String>> getSetting(String key) {
        return getSetting(key, null);
    }

    public Mono<Optional<String>> getSetting(String key, String label) {
        return readSetting(new SettingKey(requireKey(key), label), false);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Mono<Map<String, String>> listSettings(String keyPrefix, String label) {
        PrefixQuery query = new PrefixQuery(requirePrefix(keyPrefix), label);
        CachedPrefix cached = prefixCache.get(query);
        if (cached == null) {
            return loadPrefix(query);
        }
        return hasPrefixChanged(query, cached.pageEtags())
            .flatMap(changed -> changed ? loadPrefix(query) : Mono.just(cached.values()));
    }

    public Mono<Void> refreshAll() {
        List<SettingKey> settings = List.copyOf(settingCache.keySet());
        List<PrefixQuery> prefixes = List.copyOf(prefixCache.keySet());
        return Flux.concat(
                Flux.fromIterable(settings).concatMap(key -> readSetting(key, true)).then(),
                Flux.fromIterable(prefixes).concatMap(this::loadPrefix).then())
            .then();
    }

    private Mono<Optional<String>> readSetting(SettingKey key, boolean forceRefresh) {
        CachedSetting cached = settingCache.get(key);
        ConfigurationSetting request = new ConfigurationSetting()
            .setKey(key.key())
            .setLabel(key.label());
        boolean conditional = !forceRefresh && cached != null;
        if (conditional) {
            request.setETag(cached.etag());
        }

        return client.getConfigurationSettingWithResponse(request, null, conditional)
            .map(response -> {
                if (response.getStatusCode() == 304) {
                    return Optional.ofNullable(cached.value());
                }
                ConfigurationSetting setting = response.getValue();
                CachedSetting updated = new CachedSetting(setting.getValue(), setting.getETag());
                settingCache.put(key, updated);
                return Optional.ofNullable(updated.value());
            })
            .onErrorResume(ResourceNotFoundException.class, exception -> {
                settingCache.remove(key);
                return Mono.just(Optional.empty());
            });
    }

    private Mono<Boolean> hasPrefixChanged(PrefixQuery query, List<String> pageEtags) {
        if (pageEtags.isEmpty()) {
            return Mono.just(true);
        }

        SettingSelector selector = selectorFor(query).setMatchConditions(
            pageEtags.stream()
                .map(etag -> new MatchConditions().setIfNoneMatch(etag))
                .toList());

        return client.checkConfigurationSettings(selector)
            .byPage()
            .map(PagedResponse::getStatusCode)
            .collectList()
            .map(statuses -> statuses.size() != pageEtags.size()
                || statuses.stream().anyMatch(status -> status != 304));
    }

    private Mono<Map<String, String>> loadPrefix(PrefixQuery query) {
        return client.listConfigurationSettings(selectorFor(query))
            .byPage()
            .collectList()
            .map(pages -> {
                Map<String, String> values = new LinkedHashMap<>();
                List<String> pageEtags = pages.stream()
                    .map(page -> page.getHeaders().getValue(HttpHeaderName.ETAG))
                    .filter(Objects::nonNull)
                    .toList();
                pages.stream()
                    .flatMap(page -> page.getValue().stream())
                    .forEach(setting -> values.put(setting.getKey(), setting.getValue()));

                Map<String, String> immutableValues =
                    Collections.unmodifiableMap(new LinkedHashMap<>(values));
                prefixCache.put(query, new CachedPrefix(immutableValues, pageEtags));
                return immutableValues;
            });
    }

    private static SettingSelector selectorFor(PrefixQuery query) {
        String labelFilter = query.label() == null ? ConfigurationSetting.NO_LABEL : query.label();
        return new SettingSelector()
            .setKeyFilter(escapeFilter(query.prefix()) + "*")
            .setLabelFilter(labelFilter)
            .setFields(SettingFields.KEY, SettingFields.VALUE, SettingFields.ETAG);
    }

    private static String escapeFilter(String value) {
        return value.replace("\\", "\\\\")
            .replace("*", "\\*")
            .replace(",", "\\,");
    }

    private static String requireKey(String key) {
        if (key == null || key.isBlank()) {
            throw new IllegalArgumentException("Configuration key must not be blank");
        }
        return key;
    }

    private static String requirePrefix(String prefix) {
        return Objects.requireNonNull(prefix, "keyPrefix");
    }

    private record SettingKey(String key, String label) {
    }

    private record PrefixQuery(String prefix, String label) {
    }

    private record CachedSetting(String value, String etag) {
    }

    private record CachedPrefix(Map<String, String> values, List<String> pageEtags) {
    }
}
