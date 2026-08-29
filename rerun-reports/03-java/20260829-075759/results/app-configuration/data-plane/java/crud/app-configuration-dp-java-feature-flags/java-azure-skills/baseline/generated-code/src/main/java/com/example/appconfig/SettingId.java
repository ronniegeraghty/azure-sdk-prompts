package com.example.appconfig;

record SettingId(String key, String label) {
    SettingId {
        if (key == null || key.isBlank()) {
            throw new IllegalArgumentException("Configuration key must not be blank");
        }
    }

    static SettingId of(String key, String label) {
        return new SettingId(key, label);
    }

    @Override
    public String toString() {
        return label == null ? key : key + " [" + label + "]";
    }
}
