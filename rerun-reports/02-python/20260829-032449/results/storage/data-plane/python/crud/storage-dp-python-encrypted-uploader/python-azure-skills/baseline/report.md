# Evaluation Report: storage-dp-python-encrypted-uploader

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 381.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `storage-dp-python-encrypted-uploader` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 22/26 |
| Duration | 381.9s |
| Timestamp | 2026-08-28T22:23:09Z |
| Files Generated | 6 |
| Event Count | 14007 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 159.3s |
| Review | 221.4s |
| **Total** | **381.9s** |

## Configuration

- **name:** python-azure-skills/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=331897 out=8889 |
| Turn Count | 12 |

## Tool Calls

### ✅ 1. glob — 5595ms

**Input:**

```json
{"pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 4434ms

**Input:**

```json
{"pattern":"*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. apply_patch — 3247ms

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+azure-keyvault-keys\u003e=4.9.0,\u003c5.0.0\n+azure-storage-blob\u003e=12.22.0,\u003c13.0.0\n+cryptography\u003e=43.0.0,\u003c47.0.0\n*** Add File: config.py\n+from __future__ import annotations\n+\n+import os\n+from dataclasses import dataclass\n+\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+from azure.keyvault.keys import KeyClient\n+from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient\n+from azure.storage.blob import BlobServiceClient, ContainerClient\n+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient\n+from azure.storage.blob.aio import ContainerClient as AsyncContainerClient\n+\n+\n+class ConfigurationError(ValueError):\n+    \"\"\"Raised when required application configuration is missing.\"\"\"\n+\n+\n+def _required_environment_variable(name: str) -\u003e str:\n+    value = os.getenv(name)\n+    if not value:\n+        raise ConfigurationError(f\"Required environment variable {name!r} is not set\")\n+    return value\n+\n+\n+@dataclass(frozen=True)\n+class Settings:\n+    storage_account_url: str\n+    blob_container: str\n+    key_vault_url: str\n+    key_name: str\n+    key_version: str | None = None\n+\n+    @classmethod\n+    def from_environment(cls) -\u003e \"Settings\":\n+        return cls(\n+            storage_account_url=_required_environment_variable(\n+                \"AZURE_STORAGE_ACCOUNT_URL\"\n+            ),\n+            blob_container=_required_environment_variable(\"AZURE_BLOB_CONTAINER\"),\n+            key_vault_url=_required_environment_variable(\"AZURE_KEY_VAULT_URL\"),\n+            key_name=_required_environment_variable(\"AZURE_KEY_NAME\"),\n+            key_version=os.getenv(\"AZURE_KEY_VERSION\") or None,\n+        )\n+\n+\n+class SyncAzureConnections:\n+    \"\"\"Builds all synchronous Azure clients from one shared credential.\"\"\"\n+\n+    def __init__(\n+        self,\n+        settings: Settings,\n+        credential: DefaultAzureCredential | None = None,\n+    ) -\u003e None:\n+        self.settings = settings\n+        self.credential = credential or DefaultAzureCredential()\n+        self.blob_service_client = BlobServiceClient(\n+            account_url=settings.storage_account_url,\n+            credential=self.credential,\n+        )\n+        self.container_client: ContainerClient = (\n+            self.blob_service_client.get_container_client(settings.blob_container)\n+        )\n+        self.key_client = KeyClient(\n+            vault_url=settings.key_vault_url,\n+            credential=self.credential,\n+        )\n+\n+    def close(self) -\u003e None:\n+        self.key_client.close()\n+        self.blob_service_client.close()\n+        self.credential.close()\n+\n+    def __enter__(self) -\u003e \"SyncAzureConnections\":\n+        return self\n+\n+    def __exit__(self, *_: object) -\u003e None:\n+        self.close()\n+\n+\n+class AsyncAzureConnections:\n+    \"\"\"Builds all asynchronous Azure clients from one shared async credential.\"\"\"\n+\n+    def __init__(\n+        self,\n+        settings: Settings,\n+        credential: AsyncDefaultAzureCredential | None = None,\n+    ) -\u003e None:\n+        self.settings = settings\n+        self.credential = credential or AsyncDefaultAzureCredential()\n+        self.blob_service_client = AsyncBlobServiceClient(\n+            account_url=settings.storage_account_url,\n+            credential=self.credential,\n+        )\n+        self.container_client: AsyncContainerClient = (\n+            self.blob_service_client.get_container_client(settings.blob_container)\n+        )\n+        self.key_client = AsyncKeyClient(\n+            vault_url=settings.key_vault_url,\n+            credential=self.credential,\n+        )\n+\n+    async def close(self) -\u003e None:\n+        await self.key_client.close()\n+        await self.blob_service_client.close()\n+        await self.credential.close()\n+\n+    async def __aenter__(self) -\u003e \"AsyncAzureConnections\":\n+        return self\n+\n+    async def __aexit__(self, *_: object) -\u003e None:\n+        await self.close()\n*** Add File: key_management.py\n+from __future__ import annotations\n+\n+import os\n+from dataclasses import dataclass\n+\n+from azure.core.exceptions import HttpResponseError\n+from azure.keyvault.keys import KeyClient\n+from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient\n+from azure.keyvault.keys.crypto import CryptographyClient, KeyWrapAlgorithm\n+from azure.keyvault.keys.crypto.aio import CryptographyClient as AsyncCryptographyClient\n+from azure.core.credentials import TokenCredential\n+from azure.core.credentials_async import AsyncTokenCredential\n+\n+DATA_KEY_SIZE_BYTES = 32\n+KEY_WRAP_ALGORITHM = KeyWrapAlgorithm.rsa_oaep_256\n+\n+\n+class KeyManagementError(RuntimeError):\n+    \"\"\"Raised when a Key Vault key operation fails.\"\"\"\n+\n+\n+@dataclass(frozen=True)\n+class WrappedDataKey:\n+    key_id: str\n+    encrypted_key: bytes\n+    algorithm: str = KEY_WRAP_ALGORITHM.value\n+\n+\n+class KeyManager:\n+    def __init__(\n+        self,\n+        key_client: KeyClient,\n+        credential: TokenCredential,\n+        key_name: str,\n+        key_version: str | None = None,\n+    ) -\u003e None:\n+        self._key_client = key_client\n+        self._credential = credential\n+        self._key_name = key_name\n+        self._key_version = key_version\n+\n+    def generate_and_wrap_data_key(self) -\u003e tuple[bytes, WrappedDataKey]:\n+        data_key = os.urandom(DATA_KEY_SIZE_BYTES)\n+        try:\n+            key = self._key_client.get_key(self._key_name, self._key_version)\n+            if not key.id:\n+                raise KeyManagementError(\"Key Vault returned a key without an ID\")\n+            crypto_client = CryptographyClient(key, credential=self._credential)\n+            result = crypto_client.wrap_key(KEY_WRAP_ALGORITHM, data_key)\n+            return data_key, WrappedDataKey(\n+                key_id=key.id,\n+                encrypted_key=result.encrypted_key,\n+            )\n+        except HttpResponseError as exc:\n+            raise KeyManagementError(\n+                f\"Key Vault could not wrap the data key with {self._key_name!r}: {exc}\"\n+            ) from exc\n+\n+    def unwrap_data_key(self, wrapped_key: WrappedDataKey) -\u003e bytes:\n+        if wrapped_key.algorithm != KEY_WRAP_ALGORITHM.value:\n+            raise KeyManagementError(\n+                f\"Unsupported key wrap algorithm: {wrapped_key.algorithm!r}\"\n+            )\n+        try:\n+            crypto_client = CryptographyClient(\n+                wrapped_key.key_id,\n+                credential=self._credential,\n+            )\n+            result = crypto_client.unwrap_key(\n+                KEY_WRAP_ALGORITHM,\n+                wrapped_key.encrypted_key,\n+            )\n+            return result.key\n+        except HttpResponseError as exc:\n+            raise KeyManagementError(\n+                f\"Key Vault could not unwrap the data key with \"\n+                f\"{wrapped_key.key_id!r}: {exc}\"\n+            ) from exc\n+\n+\n+class AsyncKeyManager:\n+    def __init__(\n+        self,\n+        key_client: AsyncKeyClient,\n+        credential: AsyncTokenCredential,\n+        key_name: str,\n+        key_version: str | None = None,\n+    ) -\u003e None:\n+        self._key_client = key_client\n+        self._credential = credential\n+        self._key_name = key_name\n+        self._key_version = key_version\n+\n+    async def generate_and_wrap_data_key(self) -\u003e tuple[bytes, WrappedDataKey]:\n+        data_key = os.urandom(DATA_KEY_SIZE_BYTES)\n+        try:\n+            key = await self._key_client.get_key(self._key_name, self._key_version)\n+            if not key.id:\n+                raise KeyManagementError(\"Key Vault returned a key without an ID\")\n+            crypto_client = AsyncCryptographyClient(\n+                key,\n+                credential=self._credential,\n+            )\n+            result = await crypto_client.wrap_key(KEY_WRAP_ALGORITHM, data_key)\n+            return data_key, WrappedDataKey(\n+                key_id=key.id,\n+                encrypted_key=result.encrypted_key,\n+            )\n+        except HttpResponseError as exc:\n+            raise KeyManagementError(\n+                f\"Key Vault could not wrap the data key with {self._key_name!r}: {exc}\"\n+            ) from exc\n+\n+    async def unwrap_data_key(self, wrapped_key: WrappedDataKey) -\u003e bytes:\n+        if wrapped_key.algorithm != KEY_WRAP_ALGORITHM.value:\n+            raise KeyManagementError(\n+                f\"Unsupported key wrap algorithm: {wrapped_key.algorithm!r}\"\n+            )\n+        try:\n+            crypto_client = AsyncCryptographyClient(\n+                wrapped_key.key_id,\n+                credential=self._credential,\n+            )\n+            result = await crypto_client.unwrap_key(\n+                KEY_WRAP_ALGORITHM,\n+                wrapped_key.encrypted_key,\n+            )\n+            return result.key\n+        except HttpResponseError as exc:\n+            raise KeyManagementError(\n+                f\"Key Vault could not unwrap the data key with \"\n+                f\"{wrapped_key.key_id!r}: {exc}\"\n+            ) from exc\n*** Add File: blob_crypto.py\n+from __future__ import annotations\n+\n+import base64\n+import binascii\n+import os\n+from dataclasses import dataclass\n+from pathlib import Path\n+from typing import Mapping\n+\n+from azure.core.exceptions import HttpResponseError, ResourceNotFoundError\n+from azure.storage.blob import ContainerClient\n+from azure.storage.blob.aio import ContainerClient as AsyncContainerClient\n+from cryptography.exceptions import InvalidTag\n+from cryptography.hazmat.primitives.ciphers.aead import AESGCM\n+\n+from key_management import AsyncKeyManager, KeyManager, WrappedDataKey\n+\n+AES_NONCE_SIZE_BYTES = 12\n+ENCRYPTION_ALGORITHM = \"AES-256-GCM\"\n+METADATA_VERSION = \"1\"\n+\n+\n+class BlobEncryptionError(RuntimeError):\n+    \"\"\"Raised when encrypted blob storage or cryptographic processing fails.\"\"\"\n+\n+\n+class EncryptedBlobNotFoundError(BlobEncryptionError):\n+    \"\"\"Raised when the requested encrypted blob does not exist.\"\"\"\n+\n+\n+@dataclass(frozen=True)\n+class UploadResult:\n+    blob_name: str\n+    key_id: str\n+    wrapped_data_key: bytes\n+\n+\n+def _b64encode(value: bytes) -\u003e str:\n+    return base64.b64encode(value).decode(\"ascii\")\n+\n+\n+def _b64decode(metadata: Mapping[str, str], name: str) -\u003e bytes:\n+    value = metadata.get(name)\n+    if not value:\n+        raise BlobEncryptionError(f\"Encrypted blob metadata is missing {name!r}\")\n+    try:\n+        return base64.b64decode(value, validate=True)\n+    except (ValueError, binascii.Error) as exc:\n+        raise BlobEncryptionError(\n+            f\"Encrypted blob metadata {name!r} is not valid base64\"\n+        ) from exc\n+\n+\n+def _encryption_metadata(wrapped_key: WrappedDataKey, nonce: bytes) -\u003e dict[str, str]:\n+    return {\n+        \"encryption-version\": METADATA_VERSION,\n+        \"encryption-algorithm\": ENCRYPTION_ALGORITHM,\n+        \"key-wrap-algorithm\": wrapped_key.algorithm,\n+        \"key-id\": wrapped_key.key_id,\n+        \"wrapped-data-key\": _b64encode(wrapped_key.encrypted_key),\n+        \"nonce\": _b64encode(nonce),\n+    }\n+\n+\n+def _parse_metadata(metadata: Mapping[str, str]) -\u003e tuple[WrappedDataKey, bytes]:\n+    if metadata.get(\"encryption-version\") != METADATA_VERSION:\n+        raise BlobEncryptionError(\"Encrypted blob has an unsupported metadata version\")\n+    if metadata.get(\"encryption-algorithm\") != ENCRYPTION_ALGORITHM:\n+        raise BlobEncryptionError(\"Encrypted blob uses an unsupported cipher\")\n+\n+    key_id = metadata.get(\"key-id\")\n+    key_wrap_algorithm = metadata.get(\"key-wrap-algorithm\")\n+    if not key_id or not key_wrap_algorithm:\n+        raise BlobEncryptionError(\"Encrypted blob metadata is incomplete\")\n+\n+    nonce = _b64decode(metadata, \"nonce\")\n+    if len(nonce) != AES_NONCE_SIZE_BYTES:\n+        raise BlobEncryptionError(\"Encrypted blob contains an invalid AES-GCM nonce\")\n+\n+    return (\n+        WrappedDataKey(\n+            key_id=key_id,\n+            encrypted_key=_b64decode(metadata, \"wrapped-data-key\"),\n+            algorithm=key_wrap_algorithm,\n+        ),\n+        nonce,\n+    )\n+\n+\n+def _decrypt(ciphertext: bytes, data_key: bytes, nonce: bytes) -\u003e bytes:\n+    try:\n+        return AESGCM(data_key).decrypt(nonce, ciphertext, None)\n+    except (InvalidTag, ValueError) as exc:\n+        raise BlobEncryptionError(\n+            \"Blob decryption failed; the ciphertext, key, or metadata is invalid\"\n+        ) from exc\n+\n+\n+class EncryptedBlobClient:\n+    def __init__(\n+        self,\n+        container_client: ContainerClient,\n+        key_manager: KeyManager,\n+    ) -\u003e None:\n+        self._container_client = container_client\n+        self._key_manager = key_manager\n+\n+    def upload_bytes(\n+        self,\n+        blob_name: str,\n+        plaintext: bytes,\n+        *,\n+        overwrite: bool = False,\n+    ) -\u003e UploadResult:\n+        data_key, wrapped_key = self._key_manager.generate_and_wrap_data_key()\n+        nonce = os.urandom(AES_NONCE_SIZE_BYTES)\n+        ciphertext = AESGCM(data_key).encrypt(nonce, plaintext, None)\n+        metadata = _encryption_metadata(wrapped_key, nonce)\n+\n+        try:\n+            self._container_client.upload_blob(\n+                name=blob_name,\n+                data=ciphertext,\n+                metadata=metadata,\n+                overwrite=overwrite,\n+            )\n+        except HttpResponseError as exc:\n+            raise BlobEncryptionError(\n+                f\"Blob Storage could not upload {blob_name!r}: {exc}\"\n+            ) from exc\n+\n+        return UploadResult(\n+            blob_name=blob_name,\n+            key_id=wrapped_key.key_id,\n+            wrapped_data_key=wrapped_key.encrypted_key,\n+        )\n+\n+    def download_bytes(self, blob_name: str) -\u003e bytes:\n+        try:\n+            downloader = self._container_client.download_blob(blob_name)\n+            ciphertext = downloader.readall()\n+            metadata = downloader.properties.metadata or {}\n+        except ResourceNotFoundError as exc:\n+            raise EncryptedBlobNotFoundError(\n+                f\"Encrypted blob {blob_name!r} does not exist\"\n+            ) from exc\n+        except HttpResponseError as exc:\n+            raise BlobEncryptionError(\n+                f\"Blob Storage could not download {blob_name!r}: {exc}\"\n+            ) from exc\n+\n+        wrapped_key, nonce = _parse_metadata(metadata)\n+        data_key = self._key_manager.unwrap_data_key(wrapped_key)\n+        return _decrypt(ciphertext, data_key, nonce)\n+\n+    def upload_file(\n+        self,\n+        source: str | Path,\n+        blob_name: str,\n+        *,\n+        overwrite: bool = False,\n+    ) -\u003e UploadResult:\n+        try:\n+            plaintext = Path(source).read_bytes()\n+        except OSError as exc:\n+            raise BlobEncryptionError(f\"Could not read {str(source)!r}: {exc}\") from exc\n+        return self.upload_bytes(blob_name, plaintext, overwrite=overwrite)\n+\n+    def download_file(self, blob_name: str, destination: str | Path) -\u003e None:\n+        plaintext = self.download_bytes(blob_name)\n+        try:\n+            Path(destination).write_bytes(plaintext)\n+        except OSError as exc:\n+            raise BlobEncryptionError(\n+                f\"Could not write {str(destination)!r}: {exc}\"\n+            ) from exc\n+\n+\n+class AsyncEncryptedBlobClient:\n+    def __init__(\n+        self,\n+        container_client: AsyncContainerClient,\n+        key_manager: AsyncKeyManager,\n+    ) -\u003e None:\n+        self._container_client = container_client\n+        self._key_manager = key_manager\n+\n+    async def upload_bytes(\n+        self,\n+        blob_name: str,\n+        plaintext: bytes,\n+        *,\n+        overwrite: bool = False,\n+    ) -\u003e UploadResult:\n+        data_key, wrapped_key = await self._key_manager.generate_and_wrap_data_key()\n+        nonce = os.urandom(AES_NONCE_SIZE_BYTES)\n+        ciphertext = AESGCM(data_key).encrypt(nonce, plaintext, None)\n+        metadata = _encryption_metadata(wrapped_key, nonce)\n+\n+        try:\n+            await self._container_client.upload_blob(\n+                name=blob_name,\n+                data=ciphertext,\n+                metadata=metadata,\n+                overwrite=overwrite,\n+            )\n+        except HttpResponseError as exc:\n+            raise BlobEncryptionError(\n+                f\"Blob Storage could not upload {blob_name!r}: {exc}\"\n+            ) from exc\n+\n+        return UploadResult(\n+            blob_name=blob_name,\n+            key_id=wrapped_key.key_id,\n+            wrapped_data_key=wrapped_key.encrypted_key,\n+        )\n+\n+    async def download_bytes(self, blob_name: str) -\u003e bytes:\n+        try:\n+            downloader = await self._container_client.download_blob(blob_name)\n+            ciphertext = await downloader.readall()\n+            metadata = downloader.properties.metadata or {}\n+        except ResourceNotFoundError as exc:\n+            raise EncryptedBlobNotFoundError(\n+                f\"Encrypted blob {blob_name!r} does not exist\"\n+            ) from exc\n+        except HttpResponseError as exc:\n+            raise BlobEncryptionError(\n+                f\"Blob Storage could not download {blob_name!r}: {exc}\"\n+            ) from exc\n+\n+        wrapped_key, nonce = _parse_metadata(metadata)\n+        data_key = await self._key_manager.unwrap_data_key(wrapped_key)\n+        return _decrypt(ciphertext, data_key, nonce)\n+\n+    async def upload_file(\n+        self,\n+        source: str | Path,\n+        blob_name: str,\n+        *,\n+        overwrite: bool = False,\n+    ) -\u003e UploadResult:\n+        try:\n+            plaintext = Path(source).read_bytes()\n+        except OSError as exc:\n+            raise BlobEncryptionError(f\"Could not read {str(source)!r}: {exc}\") from exc\n+        return await self.upload_bytes(blob_name, plaintext, overwrite=overwrite)\n+\n+    async def download_file(\n+        self,\n+        blob_name: str,\n+        destination: str | Path,\n+    ) -\u003e None:\n+        plaintext = await self.download_bytes(blob_name)\n+        try:\n+            Path(destination).write_bytes(plaintext)\n+        except OSError as exc:\n+            raise BlobEncryptionError(\n+                f\"Could not write {str(destination)!r}: {exc}\"\n+            ) from exc\n*** Add File: main.py\n+from __future__ import annotations\n+\n+import asyncio\n+import base64\n+\n+from blob_crypto import AsyncEncryptedBlobClient, EncryptedBlobClient\n+from config import AsyncAzureConnections, Settings, SyncAzureConnections\n+from key_management import AsyncKeyManager, KeyManager\n+\n+SYNC_BLOB_NAME = \"encrypted-demo-sync.bin\"\n+ASYNC_BLOB_NAME = \"encrypted-demo-async.bin\"\n+\n+\n+def print_result(label: str, key_id: str, wrapped_key: bytes, plaintext: bytes) -\u003e None:\n+    print(f\"{label} vault key ID: {key_id}\")\n+    print(\n+        f\"{label} wrapped DEK (base64): \"\n+        f\"{base64.b64encode(wrapped_key).decode('ascii')}\"\n+    )\n+    print(f\"{label} decrypted output: {plaintext.decode('utf-8')}\")\n+\n+\n+def run_sync(settings: Settings) -\u003e None:\n+    with SyncAzureConnections(settings) as connections:\n+        key_manager = KeyManager(\n+            key_client=connections.key_client,\n+            credential=connections.credential,\n+            key_name=settings.key_name,\n+            key_version=settings.key_version,\n+        )\n+        encrypted_blobs = EncryptedBlobClient(\n+            connections.container_client,\n+            key_manager,\n+        )\n+        result = encrypted_blobs.upload_bytes(\n+            SYNC_BLOB_NAME,\n+            b\"Hello from the synchronous encrypted uploader.\",\n+            overwrite=True,\n+        )\n+        plaintext = encrypted_blobs.download_bytes(SYNC_BLOB_NAME)\n+        print_result(\"Sync\", result.key_id, result.wrapped_data_key, plaintext)\n+\n+\n+async def run_async(settings: Settings) -\u003e None:\n+    async with AsyncAzureConnections(settings) as connections:\n+        key_manager = AsyncKeyManager(\n+            key_client=connections.key_client,\n+            credential=connections.credential,\n+            key_name=settings.key_name,\n+            key_version=settings.key_version,\n+        )\n+        encrypted_blobs = AsyncEncryptedBlobClient(\n+            connections.container_client,\n+            key_manager,\n+        )\n+        result = await encrypted_blobs.upload_bytes(\n+            ASYNC_BLOB_NAME,\n+            b\"Hello from the asynchronous encrypted uploader.\",\n+            overwrite=True,\n+        )\n+        plaintext = await encrypted_blobs.download_bytes(ASYNC_BLOB_NAME)\n+        print_result(\"Async\", result.key_id, result.wrapped_data_key, plaintext)\n+\n+\n+def main() -\u003e None:\n+    settings = Settings.from_environment()\n+    run_sync(settings)\n+    asyncio.run(run_async(settings))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: .env.example\n+AZURE_STORAGE_ACCOUNT_URL=https://your-account.blob.core.windows.net\n+AZURE_BLOB_CONTAINER=encrypted-files\n+AZURE_KEY_VAULT_URL=https://your-vault.vault.azure.net\n+AZURE_KEY_NAME=blob-encryption-key\n+# Optional: pin a specific RSA key version instead of using the current version.\n+# AZURE_KEY_VERSION=\n*** Add File: README.md\n+# Azure client-side encrypted blob uploader\n+\n+This project encrypts each payload locally with a fresh AES-256-GCM data\n+encryption key (DEK). An RSA key in Azure Key Vault's Keys service wraps the\n+DEK with RSA-OAEP-256. Only the wrapped DEK, nonce, algorithms, version, and\n+Key Vault key ID are persisted as blob metadata.\n+\n+## Prerequisites\n+\n+- An existing Blob Storage account and container.\n+- An existing RSA key in Azure Key Vault.\n+- An identity supported by `DefaultAzureCredential` with blob data access and\n+  Key Vault `keys/get`, `keys/wrapKey`, and `keys/unwrapKey` permissions.\n+\n+The project does not create or modify Azure resources.\n+\n+## Run\n+\n+Create and activate a virtual environment, install `requirements.txt`, and set\n+the variables shown in `.env.example` in your shell. Then run:\n+\n+    python main.py\n+\n+The demo uploads two blobs: one through the synchronous clients and one through\n+the asynchronous clients. It then downloads and decrypts both.\n+\n+The sync Azure clients share one synchronous `DefaultAzureCredential` instance.\n+The async clients share one asynchronous `DefaultAzureCredential` instance;\n+the Azure SDK requires separate sync and async credential protocols.\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/requirements.txt b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/requirements.txt
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/requirements.txt
@@ -1,0 +1,5 @@
+azure-identity>=1.17.0,<2.0.0
+azure-keyvault-keys>=4.9.0,<5.0.0
+azure-storage-blob>=12.22.0,<13.0.0
+cryptography>=43.0.0,<47.0.0
+



diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/config.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/config.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/config.py
@@ -1,0 +1,113 @@
+from __future__ import annotations
+
+import os
+from dataclasses import dataclass
+
+from azure.identity import DefaultAzureCredential
+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
+from azure.keyvault.keys import KeyClient
+from azure.keyvault.keys.aio import KeyClient as AsyncKeyClient
+from azure.storage.blob import BlobServiceClient, ContainerClient
+from azure.storage.blob.aio import BlobServiceClient as AsyncBlobServiceClient
+from azure.storage.blob.aio import ContainerClient as AsyncContainerClient
+
+
+class ConfigurationError(ValueError):
+    """Raised when required application configuration is missing."""
+
+
+def _required_environment_variable(name: str) -> str:
+    value = os.getenv(name)
+    if not value:
+        raise ConfigurationError(f"Required environment variable {name!r} is not set")
+    retur
... (truncated)
```

### ✅ 4. apply_patch — 2970ms

**Input:**

```json
"*** Begin Patch\n*** Update File: blob_crypto.py\n@@\n import base64\n import binascii\n+import asyncio\n import os\n@@\n     return {\n-        \"encryption-version\": METADATA_VERSION,\n-        \"encryption-algorithm\": ENCRYPTION_ALGORITHM,\n-        \"key-wrap-algorithm\": wrapped_key.algorithm,\n-        \"key-id\": wrapped_key.key_id,\n-        \"wrapped-data-key\": _b64encode(wrapped_key.encrypted_key),\n+        \"encryption_version\": METADATA_VERSION,\n+        \"encryption_algorithm\": ENCRYPTION_ALGORITHM,\n+        \"key_wrap_algorithm\": wrapped_key.algorithm,\n+        \"key_id\": wrapped_key.key_id,\n+        \"wrapped_data_key\": _b64encode(wrapped_key.encrypted_key),\n         \"nonce\": _b64encode(nonce),\n     }\n@@\n-    if metadata.get(\"encryption-version\") != METADATA_VERSION:\n+    if metadata.get(\"encryption_version\") != METADATA_VERSION:\n         raise BlobEncryptionError(\"Encrypted blob has an unsupported metadata version\")\n-    if metadata.get(\"encryption-algorithm\") != ENCRYPTION_ALGORITHM:\n+    if metadata.get(\"encryption_algorithm\") != ENCRYPTION_ALGORITHM:\n         raise BlobEncryptionError(\"Encrypted blob uses an unsupported cipher\")\n \n-    key_id = metadata.get(\"key-id\")\n-    key_wrap_algorithm = metadata.get(\"key-wrap-algorithm\")\n+    key_id = metadata.get(\"key_id\")\n+    key_wrap_algorithm = metadata.get(\"key_wrap_algorithm\")\n@@\n             key_id=key_id,\n-            encrypted_key=_b64decode(metadata, \"wrapped-data-key\"),\n+            encrypted_key=_b64decode(metadata, \"wrapped_data_key\"),\n@@\n-            plaintext = Path(source).read_bytes()\n+            plaintext = await asyncio.to_thread(Path(source).read_bytes)\n@@\n-            Path(destination).write_bytes(plaintext)\n+            await asyncio.to_thread(Path(destination).write_bytes, plaintext)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
@@ -2,6 +2,7 @@
 
 import base64
 import binascii
+import asyncio
 import os
 from dataclasses import dataclass
 from pathlib import Path
@@ -53,23 +54,23 @@
 
 def _encryption_metadata(wrapped_key: WrappedDataKey, nonce: bytes) -> dict[str, str]:
     return {
-        "encryption-version": METADATA_VERSION,
-        "encryption-algorithm": ENCRYPTION_ALGORITHM,
-        "key-wrap-algorithm": wrapped_key.algorithm,
-        "key-id": wrapped_key.key_id,
-        "wrapped-data-key": _b64encode(wrapped_key.encrypted_key),
+        "encryption_version": METADATA_VERSION,
+        "encryption_algorithm": ENCRYPTION_ALGORITHM,
+        "key_wrap_algorithm": wrapped_key.algorithm,
+        "key_id": wrapped_key.key_id,
+        "wrapped_data_key": _b64encode(wrapped_key.encrypted_key),
         "nonce": _b64encode(nonce),
     }
 
 
 def _parse_metadata(metadata: Mapping[str, str]) -> tuple[WrappedDataKey, bytes]:
-    if metadata.get("encryption-version") != METADATA_VERSION:
+    if metadata.get("encryption_version") != METADATA_VERSION:
         raise BlobEncryptionError("Encrypted blob has an unsupported metadata version")
-    if metadata.get("encryption-algorithm") != ENCRYPTION_ALGORITHM:
+    if metadata.get("encryption_algorithm") != ENCRYPTION_ALGORITHM:
         raise BlobEncryptionError("Encrypted blob uses an unsupported cipher")
 
-    key_id = metadata.get("key-id")
-    key_wr
... (truncated)
```

### ✅ 5. powershell — 8280ms

**Input:**

```json
{"command":"python -m compileall -q .","description":"Compile all Python source files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
*** Error compiling '.\\blob_crypto.py'...
  File ".\blob_crypto.py", line 165
    plaintext = await asyncio.to_thread(Path(source).read_bytes)
                ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
SyntaxError: 'await' outside async function
<shellId: 0 completed with exit code 1>
```

### ✅ 6. rg — 9608ms

**Input:**

```json
{"-n":true,"glob":"*.py","output_mode":"content","paths":".","pattern":"encryption-|key-wrap|key-id|wrapped-data"}
```

**Output:**

```
.\config.py:6:from azure.identity import DefaultAzureCredential
.\config.py:7:from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential
.\config.py:53:        credential: DefaultAzureCredential | None = None,
.\config.py:56:        self.credential = credential or DefaultAzureCredential()
.\config.py:87:        credential: AsyncDefaultAzureCredential | None = None,
.\config.py:90:        self.credential = credential or AsyncDefaultAzureCredential()
```

### ✅ 7. rg — 7005ms

**Input:**

```json
{"-n":true,"glob":"*.py","output_mode":"content","paths":".","pattern":"DefaultAzureCredential"}
```

**Output:**

```
No matches found.
```

### ✅ 8. view — 2815ms

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345\\blob_crypto.py","view_range":[145,290]}
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
@@ -145,146 +145,146 @@
             raise EncryptedBlobNotFoundError(
                 f"Encrypted blob {blob_name!r} does not exist"
             ) from exc
         except HttpResponseError as exc:
             raise BlobEncryptionError(
                 f"Blob Storage could not download {blob_name!r}: {exc}"
             ) from exc
 
         wrapped_key, nonce = _parse_metadata(metadata)
         data_key = self._key_manager.unwrap_data_key(wrapped_key)
         return _decrypt(ciphertext, data_key, nonce)
 
     def upload_file(
         self,
         source: str | Path,
         blob_name: str,
         *,
         overwrite: bool = False,
     ) -> UploadResult:
         try:
             plaintext = await asyncio.to_thread(Path(source).read_bytes)
         except OSError as exc:
             raise BlobEncryptionError(f"Could not read {str(source)!r}: {exc}") from exc
         return self.upload_bytes(blob_name, plaintext, overwrite=overwrite)
 
     def download_file(self, blob_name: str, destination: str | Path) -> None:
         plaintext = self.download_bytes(blob_name)
         try:
             await asyncio.to_thread(Path(destination).write_bytes, plaintext)
         except OSError as exc:
             raise BlobEncryptionError(
                 f"Could not write {str(destination)!r}: {exc}"
             ) from exc
 
 
 class AsyncEncryptedBlobClient:
     def __init__(
 
... (truncated)
```

### ✅ 9. apply_patch — 2933ms

**Input:**

```json
"*** Begin Patch\n*** Update File: blob_crypto.py\n@@\n-import base64\n-import binascii\n import asyncio\n+import base64\n+import binascii\n@@\n     ) -\u003e UploadResult:\n         try:\n-            plaintext = await asyncio.to_thread(Path(source).read_bytes)\n+            plaintext = Path(source).read_bytes()\n@@\n         plaintext = self.download_bytes(blob_name)\n         try:\n-            await asyncio.to_thread(Path(destination).write_bytes, plaintext)\n+            Path(destination).write_bytes(plaintext)\n@@\n     ) -\u003e UploadResult:\n         try:\n-            plaintext = Path(source).read_bytes()\n+            plaintext = await asyncio.to_thread(Path(source).read_bytes)\n@@\n         plaintext = await self.download_bytes(blob_name)\n         try:\n-            Path(destination).write_bytes(plaintext)\n+            await asyncio.to_thread(Path(destination).write_bytes, plaintext)\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
index 0000000..0000000 100644
--- a/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-storage-dp-python-encrypted-uploader-python-azure-skills-baseline-2463496345/blob_crypto.py
@@ -1,8 +1,8 @@
 from __future__ import annotations
 
+import asyncio
 import base64
 import binascii
-import asyncio
 import os
 from dataclasses import dataclass
 from pathlib import Path
@@ -162,7 +162,7 @@
         overwrite: bool = False,
     ) -> UploadResult:
         try:
-            plaintext = await asyncio.to_thread(Path(source).read_bytes)
+            plaintext = Path(source).read_bytes()
         except OSError as exc:
             raise BlobEncryptionError(f"Could not read {str(source)!r}: {exc}") from exc
         return self.upload_bytes(blob_name, plaintext, overwrite=overwrite)
@@ -170,7 +170,7 @@
     def download_file(self, blob_name: str, destination: str | Path) -> None:
         plaintext = self.download_bytes(blob_name)
         try:
-            await asyncio.to_thread(Path(destination).write_bytes, plaintext)
+            Path(destination).write_bytes(plaintext)
         except OSError as exc:
             raise BlobEncryptionError(
                 f"Could not write {str(destination)!r}: {exc}"
@@ -242,7 +242,7 @@
         overwrite: bool = False,
     ) -> UploadResult:
         try:
-            plaintext = Path(source).read_bytes()
+            plaintext = await asyncio.to_thread(Path(source).read_bytes)
         except OSError as exc:
             raise BlobEncryptionError(f"Could not read {str(source)!r}: {exc}") from exc
         return await self.uplo
... (truncated)
```

### ✅ 10. powershell — 5256ms

**Input:**

```json
{"command":"python -m compileall -q .","description":"Compile all Python source files","initial_wait":30,"mode":"sync"}
```

**Output:**

```
dependencies available
<shellId: 2 completed with exit code 0>
```

### ✅ 11. powershell — 4929ms

**Input:**

```json
{"command":"python -c \"import azure.identity, azure.keyvault.keys, azure.storage.blob, cryptography; print('dependencies available')\"","description":"Check required Python dependencies","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

### ✅ 12. powershell — 2277ms

**Input:**

```json
{"command":"python -c \"import blob_crypto, config, key_management, main; print(key_management.KEY_WRAP_ALGORITHM.value)\"","description":"Import project modules","initial_wait":30,"mode":"sync"}
```

**Output:**

```
RSA-OAEP-256
<shellId: 3 completed with exit code 0>
```

### ✅ 13. powershell — 2217ms

**Input:**

```json
{"command":"@'\nimport asyncio\nimport os\nfrom types import SimpleNamespace\n\nfrom blob_crypto import AsyncEncryptedBlobClient, EncryptedBlobClient\nfrom key_management import WrappedDataKey\n\nclass Keys:\n    def __init__(self): self.keys = {}\n    def generate_and_wrap_data_key(self):\n        key = os.urandom(32); wrapped = os.urandom(48); self.keys[wrapped] = key\n        return key, WrappedDataKey('https://vault/keys/test/version', wrapped)\n    def unwrap_data_key(self, wrapped): return self.keys[wrapped.encrypted_key]\n\nclass Container:\n    def __init__(self): self.blobs = {}\n    def upload_blob(self, name, data, metadata, overwrite): self.blobs[name] = (data, metadata)\n    def download_blob(self, name):\n        data, metadata = self.blobs[name]\n        return SimpleNamespace(readall=lambda: data, properties=SimpleNamespace(metadata=metadata))\n\nclass AsyncKeys(Keys):\n    async def generate_and_wrap_data_key(self): return super().generate_and_wrap_data_key()\n    async def unwrap_data_key(self, wrapped): return super().unwrap_data_key(wrapped)\n\nclass AsyncDownloader:\n    def __init__(self, data, metadata):\n        self.data = data; self.properties = SimpleNamespace(metadata=metadata)\n    async def readall(self): return self.data\n\nclass AsyncContainer(Container):\n    async def upload_blob(self, name, data, metadata, overwrite): self.blobs[name] = (data, metadata)\n    async def download_blob(self, name): return AsyncDownloader(*self.blobs[name])\n\nplain = b'client-side encryption test'\nsync_client = EncryptedBlobClient(Container(), Keys())\nsync_client.upload_bytes('sync', plain)\nassert sync_client.download_bytes('sync') == plain\n\nasync def test_async():\n    client = AsyncEncryptedBlobClient(AsyncContainer(), AsyncKeys())\n    await client.upload_bytes('async', plain)\n    assert await client.download_bytes('async') == plain\n\nasyncio.run(test_async())\nprint('sync and async in-memory round trips passed')\n'@ | python -","description":"Test encrypted round trips locally","initial_wait":30,"mode":"sync"}
```

**Output:**

```
sync and async in-memory round trips passed
<shellId: 4 completed with exit code 0>
```

### ✅ 14. glob — 696ms

**Input:**

```json
{"pattern":"**/__pycache__/**"}
```

**Output:**

```
.\__pycache__\main.cpython-313.pyc
.\__pycache__\key_management.cpython-313.pyc
.\__pycache__\config.cpython-313.pyc
.\__pycache__\blob_crypto.cpython-313.pyc
```

### ✅ 15. powershell — 999ms

**Input:**

```json
{"command":"Remove-Item -LiteralPath '.\\__pycache__' -Recurse -Force","description":"Remove generated Python cache","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 5 completed with exit code 0>
```

## Generated Files

- `README.md`
- `blob_crypto.py`
- `config.py`
- `key_management.py`
- `main.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then create the sync/async modules and a locally testable configuration without provisioning Azure resources.

The design will use RSA-OAEP-256 for Key Vault key wrapping and AES-256-GCM locally, with metadata versioning and explicit service-specific errors. Sync clients share one sync credential; async clients share one async credential because Azure’s SDK uses distinct credential protocols.

The initial implementation is in place. I’m tightening Blob metadata naming for Azure compatibility and ensuring the async file helpers do not block the event loop before running project checks.

The modules now compile and load against the installed Azure SDK. I’m exercising both round trips with in-memory Blob/Key Vault stand-ins so encryption, metadata serialization, async behavior, and decryption are verified without contacting Azure.

Created the complete Python project with:

- Sync and async Azure Key Vault **Keys** envelope encryption using RSA-OAEP-256
- Local AES-256-GCM blob encryption/decryption
- Wrapped DEK, nonce, algorithms, key ID, and metadata version stored as Blob metadata
- Sync and async Blob Storage upload/download APIs for bytes and files
- Shared `DefaultAzureCredential` instances per sync/async SDK protocol
- Service-specific error handling for missing blobs, Key Vault failures, and authentication-tag failures
- Round-trip demo in `main.py`
- `requirements.txt`, `.env.example`, and setup documentation in `README.md`

The project does not provision or modify Azure resources. Configure the environment variables shown in `.env.example`, install dependencies, and run `python main.py`.

## Reviewer Notes (LLM-as-Judge)

### Summary

The generated project fully satisfies the exception-handling criterion for Azure SDK service errors.

### Strengths

- Consistent handling across synchronous and asynchronous service calls.
- Specific missing-blob exception handling.
- Exception chaining preserves underlying Azure SDK error context.

## Grader Results

- encrypted-uploader.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (18/19)
      - Uses the Key Vault Keys SDK (NOT Secrets) — critical distinction: Pass
      - Uses a cryptography library for local AES-GCM encryption: Pass
      - Uses Key Vault's cryptographic client for key wrap and unwrap operations (NOT a secrets client): Pass
      - Specifies an RSA key wrap algorithm for protecting the DEK: Pass
      - Key material never leaves Key Vault — wrap/unwrap happens server-side: Fail
      - Generates a random 256-bit data encryption key (DEK) locally: Pass
      - Encrypts data locally with AES-GCM using the DEK: Pass
      - Wraps the DEK via Key Vault before storing: Pass
      - Stores the wrapped DEK, nonce/IV, and vault key identifier as blob metadata: Pass
      - Decryption retrieves metadata, unwraps the DEK via Key Vault, and decrypts locally: Pass
      - Uses AES-GCM specifically (not CBC, ECB, or other modes): Pass
      - Generates a fresh random nonce for each encryption: Pass
      - Handles Key Vault errors (key disabled, key not found): Pass
      - Handles blob not found: Pass
      - Code must build and run without import errors or runtime crashes: Pass
      - Async versions use the async variants of Blob Storage and Key Vault clients: Pass
      - Does NOT use a secrets client instead of a keys/crypto client: Pass
      - Does NOT encrypt data directly with the vault key (must be envelope encryption): Pass
      - Does NOT store the raw DEK in plaintext: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Fail
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Pass
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: any tool (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 95% | 1.00 | 0.9474 | 19.1% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.2% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.2% | ✅ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.2% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 20.2% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 4.9474** | **61.8%** | |

## Re-run Command

```bash
hyoka run --prompt-id storage-dp-python-encrypted-uploader --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
