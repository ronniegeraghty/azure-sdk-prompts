package com.example.appconfig;

import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;
import com.azure.core.exception.ResourceNotFoundException;

import java.util.ArrayList;
import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentMap;

public final class SyncConfigurationService {
    private final ConfigurationClient client;
    private final ConcurrentMap<SettingId, ConfigurationSetting> settingCache = new ConcurrentHashMap<>();
    private final ConcurrentMap<SelectorId, Map<String, String>> prefixCache = new ConcurrentHashMap<>();

    public SyncConfigurationService(ConfigurationClient client) {
        this.client = java.util.Objects.requireNonNull(client, "client");
    }

    public Optional<String> getSetting(String key) {
        return getSetting(key, null);
    }

    public Optional<String> getSetting(String key, String label) {
        SettingId id = SettingId.of(key, label);
        ConfigurationSetting cached = settingCache.get(id);
        try {
            ConfigurationSetting request = new ConfigurationSetting().setKey(key).setLabel(label);
            if (cached != null) {
                request.setETag(cached.getETag());
            }

            Response<ConfigurationSetting> response = client.getConfigurationSettingWithResponse(
                request, null, cached != null, Context.NONE);
            ConfigurationSetting resolved = response.getStatusCode() == 304 ? cached : response.getValue();
            if (resolved != null) {
                settingCache.put(id, resolved);
            }
            return Optional.ofNullable(resolved).map(ConfigurationSetting::getValue);
        } catch (ResourceNotFoundException exception) {
            settingCache.remove(id);
            return Optional.empty();
        }
    }

    public Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public Map<String, String> listSettings(String keyPrefix, String label) {
        SelectorId id = new SelectorId(keyPrefix, label);
        return prefixCache.computeIfAbsent(id, this::loadSettings);
    }

    public boolean hasSettingChanged(String key, String label) {
        SettingId id = SettingId.of(key, label);
        ConfigurationSetting cached = settingCache.get(id);
        if (cached == null) {
            getSetting(key, label);
            return false;
        }

        try {
            Response<ConfigurationSetting> response = client.getConfigurationSettingWithResponse(
                new ConfigurationSetting()
                    .setKey(key)
                    .setLabel(label)
                    .setETag(cached.getETag()),
                null,
                true,
                Context.NONE);
            if (response.getStatusCode() == 304) {
                return false;
            }
            settingCache.put(id, response.getValue());
            return true;
        } catch (ResourceNotFoundException exception) {
            settingCache.remove(id);
            return true;
        }
    }

    public void refreshAll() {
        List<SettingId> settings = new ArrayList<>(settingCache.keySet());
        List<SelectorId> selectors = new ArrayList<>(prefixCache.keySet());
        settingCache.clear();
        prefixCache.clear();

        settings.forEach(id -> getSetting(id.key(), id.label()));
        selectors.forEach(id -> listSettings(id.prefix(), id.label()));
    }

    private Map<String, String> loadSettings(SelectorId id) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(ConfigurationFilters.keyPrefix(id.prefix()))
            .setLabelFilter(ConfigurationFilters.label(id.label()));
        Map<String, String> settings = new LinkedHashMap<>();
        client.listConfigurationSettings(selector)
            .forEach(setting -> settings.put(setting.getKey(), setting.getValue()));
        return Collections.unmodifiableMap(settings);
    }
}
