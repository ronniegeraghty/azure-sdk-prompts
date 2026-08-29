package com.example.keyvaultconfig;

public interface SecretProvider {
    SecretSnapshot getSecret(String name, String defaultValue);

    SecretSnapshot getSecret(String name, String version, String defaultValue);
}
