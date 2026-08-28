package com.example.appconfig;

import java.util.Map;

public interface ConfigurationReader {
    String getSetting(String key, String label);

    Map<String, String> listSettings(String keyPrefix, String label);
}
