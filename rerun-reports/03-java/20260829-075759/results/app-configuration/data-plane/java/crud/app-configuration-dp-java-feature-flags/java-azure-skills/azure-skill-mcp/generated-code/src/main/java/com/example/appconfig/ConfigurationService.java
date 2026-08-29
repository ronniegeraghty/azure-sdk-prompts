package com.example.appconfig;

import com.azure.core.http.HttpHeaderName;
import com.azure.core.http.MatchConditions;
import com.azure.core.http.rest.PagedResponse;
import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import com.azure.data.appconfiguration.models.SettingFields;
import com.azure.core.exception.ResourceNotFoundException;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public final class ConfigurationService {
    private final ConfigurationClient client;
    private final Map<SettingKey, CachedSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixQuery, CachedPrefix> prefixCache = new ConcurrentHashMap<>();

    public ConfigurationService(ConfigurationClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Optional<String> getSetting(String key) {
        return getSetting(key, null);
    }

    public Optional<String> getSetting(String key, String label) {
        return readSetting(new SettingKey(requireKey(key), label), false);
    }

    public Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Map<String, String> listSettings(String keyPrefix, String label) {
        PrefixQuery query = new PrefixQuery(requirePrefix(keyPrefix), label);
        CachedPrefix cached = prefixCache.get(query);
        if (cached != null && !hasPrefixChanged(query, cached.pageEtags())) {
            return cached.values();
        }
        return loadPrefix(query);
    }

    public void refreshAll() {
        List<SettingKey> settings = List.copyOf(settingCache.keySet());
        List<PrefixQuery> prefixes = List.copyOf(prefixCache.keySet());
        settings.forEach(key -> readSetting(key, true));
        prefixes.forEach(this::loadPrefix);
    }

    private Optional<String> readSetting(SettingKey key, boolean forceRefresh) {
        CachedSetting cached = settingCache.get(key);
        ConfigurationSetting request = new ConfigurationSetting()
            .setKey(key.key())
            .setLabel(key.label());
        boolean conditional = !forceRefresh && cached != null;
        if (conditional) {
            request.setETag(cached.etag());
        }

        try {
            Response<ConfigurationSetting> response =
                client.getConfigurationSettingWithResponse(request, null, conditional, Context.NONE);
            if (response.getStatusCode() == 304) {
                return Optional.ofNullable(cached.value());
            }

            ConfigurationSetting setting = response.getValue();
            CachedSetting updated = new CachedSetting(setting.getValue(), setting.getETag());
            settingCache.put(key, updated);
            return Optional.ofNullable(updated.value());
        } catch (ResourceNotFoundException exception) {
            settingCache.remove(key);
            return Optional.empty();
        }
    }

    private boolean hasPrefixChanged(PrefixQuery query, List<String> pageEtags) {
        if (pageEtags.isEmpty()) {
            return true;
        }

        SettingSelector selector = selectorFor(query).setMatchConditions(
            pageEtags.stream()
                .map(etag -> new MatchConditions().setIfNoneMatch(etag))
                .toList());

        int checkedPages = 0;
        for (PagedResponse<ConfigurationSetting> page
            : client.checkConfigurationSettings(selector).iterableByPage()) {
            checkedPages++;
            if (page.getStatusCode() != 304) {
                return true;
            }
        }
        return checkedPages == pageEtags.size();
    }

    private Map<String, String> loadPrefix(PrefixQuery query) {
        Map<String, String> values = new LinkedHashMap<>();
        List<String> pageEtags = new ArrayList<>();

        for (PagedResponse<ConfigurationSetting> page
            : client.listConfigurationSettings(selectorFor(query)).iterableByPage()) {
            String etag = page.getHeaders().getValue(HttpHeaderName.ETAG);
            if (etag != null) {
                pageEtags.add(etag);
            }
            page.getValue().forEach(setting -> values.put(setting.getKey(), setting.getValue()));
        }

        Map<String, String> immutableValues =
            Collections.unmodifiableMap(new LinkedHashMap<>(values));
        prefixCache.put(query, new CachedPrefix(immutableValues, List.copyOf(pageEtags)));
        return immutableValues;
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
