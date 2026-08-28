# Azure encrypted blob uploader

Java 17 sample implementing envelope encryption with Azure Blob Storage and the
Azure Key Vault Keys service. It generates a 256-bit AES data encryption key
(DEK) locally, encrypts content with AES-GCM, wraps the DEK with a versioned Key
Vault RSA key using RSA-OAEP-256, and stores only ciphertext and encryption
metadata in Blob Storage.

The managed identity needs data-plane permissions to read and write blobs in an
existing container and to read, wrap, and unwrap with the configured Key Vault
key. The sample does not create Azure resources.

Set these environment variables:

```text
AZURE_STORAGE_BLOB_ENDPOINT=https://<account>.blob.core.windows.net
AZURE_STORAGE_CONTAINER=<existing-container>
AZURE_KEY_VAULT_ENDPOINT=https://<vault>.vault.azure.net
AZURE_KEY_NAME=<rsa-key-name>
```

Build and run:

```text
mvn package
mvn exec:java
```

Both demos overwrite their respective sample blob:
`sync-encrypted-demo.bin` and `async-encrypted-demo.bin`.
