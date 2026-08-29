package com.example.appconfig;

import java.util.Optional;

@FunctionalInterface
public interface SettingProvider {
    Optional<String> getSetting(String key, String label);
}
