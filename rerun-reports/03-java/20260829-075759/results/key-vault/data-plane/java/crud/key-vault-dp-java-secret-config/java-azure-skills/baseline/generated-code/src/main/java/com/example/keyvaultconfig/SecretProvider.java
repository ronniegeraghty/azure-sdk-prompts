package com.example.keyvaultconfig;

public interface SecretProvider {
    SecretValue getSecret(String name, String defaultValue);

    SecretValue getSecret(String name, String version, String defaultValue);
}
