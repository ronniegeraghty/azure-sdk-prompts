package com.example.keyvault;

public interface SyncSecretProvider {
    SecretSnapshot get(String name, String defaultValue);

    SecretSnapshot getVersion(String name, String version, String defaultValue);
}
