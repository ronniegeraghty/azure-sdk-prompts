package com.example.appconfig;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Set;
import java.util.concurrent.ConcurrentHashMap;

public final class ConfigurationService {
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationClient client;
    private final Map<SettingQuery, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixQuery, Map<String, String>> prefixCache = new ConcurrentHashMap<>();
    private final Set<SettingQuery> trackedSettings = ConcurrentHashMap.newKeySet();
    private final Set<PrefixQuery> trackedPrefixes = ConcurrentHashMap.newKeySet();

    public ConfigurationService(ConfigurationClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public String getSetting(String key) {
        return getSetting(key, null);
    }

    public String getSetting(String key, String label) {
        SettingQuery query = new SettingQuery(requireText(key, "key"), label);
        trackedSettings.add(query);
        return fetchSetting(query).getValue();
    }

    public Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Map<String, String> listSettings(String keyPrefix, String label) {
        PrefixQuery query = new PrefixQuery(requireText(keyPrefix, "keyPrefix"), label);
        trackedPrefixes.add(query);
        Map<String, String> cached = prefixCache.get(query);
        return cached != null ? cached : refreshPrefix(query);
    }

    public void refreshAll() {
        trackedSettings.forEach(this::fetchSetting);
        trackedPrefixes.forEach(this::refreshPrefix);
    }

    private ConfigurationSetting fetchSetting(SettingQuery query) {
        ConfigurationSetting cached = settingCache.get(query);
        if (cached == null) {
            ConfigurationSetting loaded = client.getConfigurationSetting(query.key(), query.label());
            settingCache.put(query, loaded);
            return loaded;
        }

        Response<ConfigurationSetting> response =
            client.getConfigurationSettingWithResponse(cached, null, true, Context.NONE);
        if (response.getStatusCode() == 304) {
            return cached;
        }

        ConfigurationSetting updated = response.getValue();
        settingCache.put(query, updated);
        return updated;
    }

    private Map<String, String> refreshPrefix(PrefixQuery query) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(escapeFilter(query.prefix()) + "*")
            .setLabelFilter(query.label() == null ? NULL_LABEL_FILTER : escapeFilter(query.label()));

        Map<String, String> loaded = new LinkedHashMap<>();
        client.listConfigurationSettings(selector)
            .forEach(setting -> {
                loaded.put(setting.getKey(), setting.getValue());
                settingCache.put(new SettingQuery(setting.getKey(), setting.getLabel()), setting);
            });

        Map<String, String> snapshot = Collections.unmodifiableMap(loaded);
        prefixCache.put(query, snapshot);
        return snapshot;
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
