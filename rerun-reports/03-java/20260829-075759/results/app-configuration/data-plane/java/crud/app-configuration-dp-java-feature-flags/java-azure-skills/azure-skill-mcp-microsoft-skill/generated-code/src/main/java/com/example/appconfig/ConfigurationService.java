package com.example.appconfig;

import com.azure.core.exception.ResourceNotFoundException;
import com.azure.core.http.rest.Response;
import com.azure.core.util.Context;
import com.azure.data.appconfiguration.ConfigurationClient;
import com.azure.data.appconfiguration.models.ConfigurationSetting;
import com.azure.data.appconfiguration.models.SettingSelector;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;

public final class ConfigurationService implements SettingProvider {
    private static final int NOT_MODIFIED = 304;
    private static final String NULL_LABEL_FILTER = "\0";

    private final ConfigurationClient client;
    private final Map<SettingId, ConfigurationSetting> settingCache = new HashMap<>();
    private final Map<PrefixQuery, Map<String, String>> prefixCache = new HashMap<>();
    private final Map<SettingId, Optional<String>> sentinelValues = new HashMap<>();

    public ConfigurationService(ConfigurationClient client) {
        this.client = Objects.requireNonNull(client, "client");
    }

    public Optional<String> getSetting(String key) {
        return getSetting(key, null);
    }

    @Override
    public synchronized Optional<String> getSetting(String key, String label) {
        SettingId id = new SettingId(key, label);
        ConfigurationSetting cached = settingCache.get(id);
        if (cached == null) {
            return loadSetting(id).map(ConfigurationSetting::getValue);
        }

        try {
            Response<ConfigurationSetting> response =
                client.getConfigurationSettingWithResponse(cached, null, true, Context.NONE);
            if (response.getStatusCode() == NOT_MODIFIED) {
                return Optional.ofNullable(cached.getValue());
            }

            ConfigurationSetting updated = response.getValue();
            settingCache.put(id, updated);
            return Optional.ofNullable(updated.getValue());
        } catch (ResourceNotFoundException exception) {
            settingCache.remove(id);
            return Optional.empty();
        }
    }

    public synchronized Map<String, String> listSettings(String keyPrefix) {
        return listSettings(keyPrefix, null);
    }

    public synchronized Map<String, String> listSettings(String keyPrefix, String label) {
        PrefixQuery query = new PrefixQuery(keyPrefix, label);
        Map<String, String> cached = prefixCache.get(query);
        if (cached != null) {
            return cached;
        }

        Map<String, String> loaded = loadPrefix(query);
        prefixCache.put(query, loaded);
        return loaded;
    }

    synchronized boolean checkForUpdate(Sentinel sentinel) {
        SettingId id = new SettingId(sentinel.key(), sentinel.label());
        Optional<String> current = getSetting(id.key(), id.label());
        Optional<String> previous = sentinelValues.put(id, current);
        return previous != null && !previous.equals(current);
    }

    public synchronized void refreshAll() {
        var settingIds = settingCache.keySet().toArray(SettingId[]::new);
        var prefixQueries = prefixCache.keySet().toArray(PrefixQuery[]::new);

        for (SettingId id : settingIds) {
            loadSetting(id);
        }
        for (PrefixQuery query : prefixQueries) {
            prefixCache.put(query, loadPrefix(query));
        }
    }

    private Optional<ConfigurationSetting> loadSetting(SettingId id) {
        try {
            ConfigurationSetting setting = client.getConfigurationSetting(id.key(), id.label());
            settingCache.put(id, setting);
            return Optional.of(setting);
        } catch (ResourceNotFoundException exception) {
            settingCache.remove(id);
            return Optional.empty();
        }
    }

    private Map<String, String> loadPrefix(PrefixQuery query) {
        SettingSelector selector = new SettingSelector()
            .setKeyFilter(query.prefix() + "*")
            .setLabelFilter(query.label() == null ? NULL_LABEL_FILTER : query.label());

        Map<String, String> values = new LinkedHashMap<>();
        client.listConfigurationSettings(selector)
            .forEach(setting -> {
                values.put(setting.getKey(), setting.getValue());
                settingCache.put(new SettingId(setting.getKey(), setting.getLabel()), setting);
            });
        return Map.copyOf(values);
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
