package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;

public class ConfigurationService {
    private final ConfigurationClient client;
    private final Map<CacheKey, ConfigurationSetting> cache = new ConcurrentHashMap<>();

    public ConfigurationService(ConfigurationClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Optional<String> getSetting(String key) {
        return getSetting(key, null);
    }

    public Optional<String> getSetting(String key, String label) {
        Objects.requireNonNull(key, "key");
        CacheKey cacheKey = new CacheKey(key, label);
        ConfigurationSetting cached = cache.get(cacheKey);
        ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);

        if (cached != null) {
            request.setETag(cached.getETag());
        }

        try {
            Response<ConfigurationSetting> response =
                client.getConfigurationSettingWithResponse(request, null, cached != null, Context.NONE);
            if (response.getStatusCode() == 304) {
                return Optional.ofNullable(cached.getValue());
            }

            ConfigurationSetting current = response.getValue();
            cache.put(cacheKey, current);
            return Optional.ofNullable(current.getValue());
        } catch (ResourceNotFoundException exception) {
            cache.remove(cacheKey);
            return Optional.empty();
        }
    }

    public Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Map<String, String> listSettings(String keyPrefix, String label) {
        Objects.requireNonNull(keyPrefix, "keyPrefix");
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(keyPrefix + "*")
            .setLabelFilter(label == null ? ConfigurationSetting.NO_LABEL : label);

        Map<String, String> result = new LinkedHashMap<>();
        for (ConfigurationSetting setting : client.listConfigurationSettings(selector)) {
            result.put(setting.getKey(), setting.getValue());
            cache.put(new CacheKey(setting.getKey(), setting.getLabel()), setting);
        }
        return Map.copyOf(result);
    }

    public void refreshAll() {
        List<CacheKey> cachedKeys = List.copyOf(cache.keySet());
        for (CacheKey cachedKey : cachedKeys) {
            getSetting(cachedKey.key(), cachedKey.label());
        }
    }

    private record CacheKey(String key, String label) {
    }
}
