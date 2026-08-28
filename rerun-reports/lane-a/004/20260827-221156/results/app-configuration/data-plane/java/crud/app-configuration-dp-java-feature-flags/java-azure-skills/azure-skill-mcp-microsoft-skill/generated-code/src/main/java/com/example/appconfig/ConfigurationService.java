package com.example.appconfig;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;

public final class ConfigurationService implements ConfigurationReader {
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationClient client;
    private final Map<SettingKey, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final Map<PrefixKey, Map<String, String>> prefixCache = new ConcurrentHashMap<>();

    public ConfigurationService(ConfigurationClient client) {
        this.client = client;
    }

    public String getSetting(String key) {
        return getSetting(key, null);
    }

    @Override
    public String getSetting(String key, String label) {
        SettingKey cacheKey = new SettingKey(key, label);
        ConfigurationSetting cached = settingCache.get(cacheKey);
        if (cached == null) {
            ConfigurationSetting loaded = client.getConfigurationSetting(key, label);
            settingCache.put(cacheKey, loaded);
            return loaded.getValue();
        }

        Response<ConfigurationSetting> response = client.getConfigurationSettingWithResponse(
            cached, null, true, Context.NONE);
        if (response.getStatusCode() == 304) {
            return cached.getValue();
        }

        ConfigurationSetting updated = response.getValue();
        settingCache.put(cacheKey, updated);
        return updated.getValue();
    }

    public Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    @Override
    public Map<String, String> listSettings(String keyPrefix, String label) {
        PrefixKey cacheKey = new PrefixKey(keyPrefix, label);
        return prefixCache.computeIfAbsent(cacheKey, this::loadPrefix);
    }

    public synchronized void refreshAll() {
        List<SettingKey> settings = List.copyOf(settingCache.keySet());
        List<PrefixKey> prefixes = List.copyOf(prefixCache.keySet());
        settingCache.clear();
        prefixCache.clear();

        settings.forEach(key -> getSetting(key.key(), key.label()));
        prefixes.forEach(key -> listSettings(key.prefix(), key.label()));
    }

    private Map<String, String> loadPrefix(PrefixKey key) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(escapeFilter(key.prefix()) + "*")
            .setLabelFilter(key.label() == null ? NULL_LABEL_FILTER : key.label());
        Map<String, String> values = new LinkedHashMap<>();
        client.listConfigurationSettings(selector).forEach(setting -> {
            values.put(setting.getKey(), setting.getValue());
            settingCache.put(new SettingKey(setting.getKey(), setting.getLabel()), setting);
        });
        return Map.copyOf(values);
    }

    private static String escapeFilter(String value) {
        return value.replace("\\", "\\\\").replace(",", "\\,").replace("*", "\\*");
    }

    private record SettingKey(String key, String label) {
    }

    private record PrefixKey(String prefix, String label) {
    }
}
