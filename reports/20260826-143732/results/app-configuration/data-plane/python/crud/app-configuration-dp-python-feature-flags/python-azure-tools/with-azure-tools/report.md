# Evaluation Report: app-configuration-dp-python-feature-flags

**Config:** python-azure-tools/with-azure-tools | **Result:** ❌ FAILED | **Duration:** 331.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-python-feature-flags` |
| Config | python-azure-tools/with-azure-tools |
| Result | ❌ FAILED |
| Score | 13/16 |
| Duration | 331.9s |
| Timestamp | 2026-08-26T06:58:30Z |
| Files Generated | 8 |
| Event Count | 19761 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 234.1s |
| Review | 96.2s |
| **Total** | **331.9s** |

## Configuration

- **model:** gpt-5.6-sol
- **name:** python-azure-tools/with-azure-tools

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | agent-framework-azure-ai-py, azure-ai-contentsafety-py, azure-ai-contentunderstanding-py, azure-ai-language-conversations-py, azure-ai-ml-py, azure-ai-projects-py, azure-ai-textanalytics-py, azure-ai-transcription-py, azure-ai-translation-document-py, azure-ai-translation-text-py, azure-ai-vision-imageanalysis-py, azure-ai-voicelive-py, azure-appconfiguration-py, azure-containerregistry-py, azure-cosmos-db-py, azure-cosmos-py, azure-data-tables-py, azure-eventgrid-py, azure-eventhub-py, azure-identity-py, azure-keyvault-py, azure-messaging-webpubsubservice-py, azure-mgmt-apicenter-py, azure-mgmt-apimanagement-py, azure-mgmt-botservice-py, azure-mgmt-fabric-py, azure-monitor-ingestion-py, azure-monitor-opentelemetry-exporter-py, azure-monitor-opentelemetry-py, azure-monitor-query-py, azure-search-documents-py, azure-servicebus-py, azure-speech-to-text-rest-py, azure-storage-blob-py, azure-storage-file-datalake-py, azure-storage-file-share-py, azure-storage-queue-py, fastapi-router-py, m365-agents-py, pydantic-models-py, customize-cloud-agent, github-pr-media |
| Skills Invoked | azure-appconfiguration-py, azure-identity-py |
| MCP Servers | azure |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=715873 out=15739 |
| Turn Count | 14 |

## Prompt Sent

```
Create a Python project that implements a feature flag and configuration management system backed by Azure App Configuration.

**Write the code to files (use file-write tools, do not reply with code blocks).**

The project needs:

- A **configuration service module** (both sync and async versions) that retrieves settings from App Configuration. It should support fetching a single setting by key, fetching a setting with a specific label (to distinguish between environments like "production" vs "staging"), and listing all settings that match a key prefix (returned as a dictionary). It should also avoid re-downloading values that haven't changed since the last read — minimize unnecessary network traffic when polling for config changes.

- A **feature flag evaluator module** that reads feature flags from App Configuration. Feature flags in App Configuration use the `.appconfig.featureflag/` key prefix and store their state as a JSON payload. The evaluator should be able to check if a flag is enabled, and also support percentage-based rollout — if a flag is configured for a percentage rollout (e.g., 30% of users), the evaluator should deterministically decide whether a given user ID falls within the rollout percentage using a consistent hash, so the same user always gets the same result.

- A **configuration watcher module** that periodically polls for configuration changes. It should accept a list of "sentinel" keys to watch and a polling interval. When a sentinel key's value changes, the watcher should trigger a full refresh of all cached configuration. This is the recommended pattern for coordinating config updates in App Configuration.

- A **main script** that demos both implementations: connecting to App Configuration (endpoint from environment variable, authenticated with `DefaultAzureCredential`), reading some config values with labels, evaluating feature flags for a few sample user IDs with percentage rollout, and starting the config watcher to detect a change. Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.
```

## Tool Calls

### 🔧 1. skill

**Input:**

```json
{"skill":"azure-appconfiguration-py"}
```

### 🔧 2. skill

**Input:**

```json
{"skill":"azure-identity-py"}
```

### 🔧 3. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"intent":"Create a local-only Python application using Azure App Configuration and DefaultAzureCredential, with sync and async clients, ETag-aware caching, feature flags, and sentinel-based refresh.","learn":true}
```

### 🔧 4. glob

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-python-feature-flags-python-azure-tools-with-azure-tools-1382691877","pattern":"**/*"}
```

### 🔧 5. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py\\references\\capabilities.md"}
```

### 🔧 6. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py\\references\\non-hero-scenarios.md"}
```

### 🔧 7. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\capabilities.md"}
```

### 🔧 8. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py\\references\\non-hero-scenarios.md"}
```

### 🔧 9. azure-get_azure_bestpractices (via azure)

**Input:**

```json
{"command":"get_azure_bestpractices_get","intent":"Apply Azure best practices to a local-only Python Azure App Configuration sample using DefaultAzureCredential, sync/async SDK clients, conditional ETag reads, feature flags, and sentinel refresh.","parameters":{"action":"code-generation","resource":"general"}}
```

### 🔧 10. azure-documentation (via azure)

**Input:**

```json
{"intent":"Find official Python Azure App Configuration SDK documentation for conditional get_configuration_setting using ETag and MatchConditions.IfNotModified, ResourceNotModifiedError behavior, sync and async list_configuration_settings iteration, and feature flag JSON schema.","learn":true}
```

### 🔧 11. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_search","intent":"Verify Azure App Configuration Python SDK conditional ETag reads and feature flag schema.","parameters":{"query":"Azure App Configuration Python get_configuration_setting etag MatchConditions IfNotModified ResourceNotModifiedError list_configuration_settings async feature flag JSON percentage filter schema"}}
```

### 🔧 12. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_code_sample_search","intent":"Find current official Python code samples for Azure App Configuration sync and async clients with DefaultAzureCredential and conditional ETag requests.","parameters":{"language":"python","query":"AzureAppConfigurationClient get_configuration_setting etag MatchConditions.IfNotModified ResourceNotModifiedError DefaultAzureCredential async list_configuration_settings"}}
```

### 🔧 13. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727565021-copilot-tool-output-b733c827a0ee4b6d89468c8ed2a0853f.txt","pattern":"MatchConditions|ResourceNotModifiedError|etag|sourceUrl|url"}
```

### 🔧 14. rg

**Input:**

```json
{"head_limit":80,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727565018-copilot-tool-output-89af811801a84fcd835682d95972d015.txt","pattern":"MatchConditions|ResourceNotModifiedError|etag|sourceUrl|url"}
```

### 🔧 15. powershell

**Input:**

```json
{"command":"$docs = Get-Content -Raw 'C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727565021-copilot-tool-output-b733c827a0ee4b6d89468c8ed2a0853f.txt' | ConvertFrom-Json; $docs.results | Select-Object -First 8 title,contentUrl | Format-List; $docs.results | Where-Object { $_.content -match 'etag|MatchConditions|ResourceNotModified' } | Select-Object -First 3 -ExpandProperty content","description":"Extract relevant App Configuration docs","initial_wait":30,"mode":"sync"}
```

### 🔧 16. view

**Input:**

```json
{"path":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727565021-copilot-tool-output-b733c827a0ee4b6d89468c8ed2a0853f.txt","view_range":[1,10]}
```

### 🔧 17. glob

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py"],"pattern":"**/acceptance-criteria.md"}
```

### 🔧 18. glob

**Input:**

```json
{"paths":["C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-appconfiguration-py","C:\\Users\\weidxu\\.hyoka\\cache\\repos\\microsoft\\skills\\default\\.github\\plugins\\azure-sdk-python\\skills\\azure-identity-py"],"pattern":"**/examples.md"}
```

### 🔧 19. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the current complete Python Azure App Configuration client API reference identified by search.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient?view=azure-python"}}
```

### 🔧 20. azure-documentation (via azure)

**Input:**

```json
{"command":"microsoft_docs_fetch","intent":"Read the current complete async Python Azure App Configuration client API reference identified by search.","parameters":{"url":"https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.aio.azureappconfigurationclient?view=azure-python"}}
```

### 🔧 21. rg

**Input:**

```json
{"-A":24,"-B":3,"head_limit":120,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727616720-copilot-tool-output-dab752e7b7d346a7ae0b231a2a6282f1.txt","pattern":"list_configuration_settings|fields|SettingFields|get_configuration_setting\\("}
```

### 🔧 22. rg

**Input:**

```json
{"-A":24,"-B":3,"head_limit":120,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727616851-copilot-tool-output-cf29d887e36f4fa2bbdc3b4e77c262e5.txt","pattern":"list_configuration_settings|fields|SettingFields|get_configuration_setting\\("}
```

### 🔧 23. rg

**Input:**

```json
{"-n":true,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727616720-copilot-tool-output-dab752e7b7d346a7ae0b231a2a6282f1.txt","pattern":"^### list_configuration_settings|^list_configuration_settings"}
```

### 🔧 24. rg

**Input:**

```json
{"-n":true,"output_mode":"content","paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\1787727616851-copilot-tool-output-cf29d887e36f4fa2bbdc3b4e77c262e5.txt","pattern":"^### list_configuration_settings|^list_configuration_settings"}
```

### 🔧 25. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: appconfig_demo/__init__.py\n+\"\"\"Azure App Configuration caching and feature flag demo.\"\"\"\n+\n+from .configuration_service import AsyncConfigurationService, ConfigurationService\n+from .feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator\n+from .watcher import AsyncConfigurationWatcher, ConfigurationWatcher, SentinelKey\n+\n+__all__ = [\n+    \"AsyncConfigurationService\",\n+    \"AsyncConfigurationWatcher\",\n+    \"AsyncFeatureFlagEvaluator\",\n+    \"ConfigurationService\",\n+    \"ConfigurationWatcher\",\n+    \"FeatureFlagEvaluator\",\n+    \"SentinelKey\",\n+]\n*** Add File: appconfig_demo/configuration_service.py\n+\"\"\"Cached sync and async access to Azure App Configuration.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+from dataclasses import dataclass\n+from threading import RLock\n+from typing import TypeAlias\n+\n+from azure.appconfiguration import AzureAppConfigurationClient\n+from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient\n+from azure.core import MatchConditions\n+from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError\n+\n+logger = logging.getLogger(__name__)\n+\n+SettingIdentity: TypeAlias = tuple[str, str | None]\n+PrefixQuery: TypeAlias = tuple[str, str | None]\n+\n+\n+@dataclass(frozen=True)\n+class _CachedSetting:\n+    value: str | None\n+    etag: str\n+\n+\n+def _identity(key: str, label: str | None) -\u003e SettingIdentity:\n+    return key, label\n+\n+\n+def _label_filter(label: str | None) -\u003e str:\n+    # App Configuration uses the NUL label filter to select only unlabeled values.\n+    return label if label is not None else \"\\0\"\n+\n+\n+class ConfigurationService:\n+    \"\"\"Retrieve and cache settings using conditional ETag requests.\"\"\"\n+\n+    def __init__(self, client: AzureAppConfigurationClient) -\u003e None:\n+        self._client = client\n+        self._cache: dict[SettingIdentity, _CachedSetting] = {}\n+        self._exact_queries: set[SettingIdentity] = set()\n+        self._prefix_members: dict[PrefixQuery, set[SettingIdentity]] = {}\n+        self._lock = RLock()\n+\n+    def get_setting(\n+        self,\n+        key: str,\n+        label: str | None = None,\n+        *,\n+        force_refresh: bool = False,\n+    ) -\u003e str | None:\n+        \"\"\"Get one value, optionally selecting an environment label.\"\"\"\n+        if not key:\n+            raise ValueError(\"key must not be empty\")\n+\n+        identity = _identity(key, label)\n+        with self._lock:\n+            self._exact_queries.add(identity)\n+            return self._get_setting_locked(identity, force_refresh=force_refresh)\n+\n+    def _get_setting_locked(\n+        self,\n+        identity: SettingIdentity,\n+        *,\n+        force_refresh: bool,\n+    ) -\u003e str | None:\n+        key, label = identity\n+        cached = self._cache.get(identity)\n+        request_options: dict[str, object] = {}\n+        if cached is not None and not force_refresh:\n+            # If-None-Match returns no payload when the cached ETag is current.\n+            request_options = {\n+                \"etag\": cached.etag,\n+                \"match_condition\": MatchConditions.IfModified,\n+            }\n+\n+        try:\n+            setting = self._client.get_configuration_setting(\n+                key=key,\n+                label=label,\n+                **request_options,\n+            )\n+        except ResourceNotModifiedError:\n+            if cached is None:\n+                raise RuntimeError(\"received a not-modified response without a cached value\")\n+            return cached.value\n+        except ResourceNotFoundError:\n+            self._cache.pop(identity, None)\n+            raise\n+\n+        # Some transports represent HTTP 304 as None instead of raising.\n+        if setting is None:\n+            if cached is None:\n+                raise RuntimeError(\"App Configuration returned no setting\")\n+            return cached.value\n+\n+        self._cache[identity] = _CachedSetting(setting.value, str(setting.etag))\n+        return setting.value\n+\n+    def list_settings(\n+        self,\n+        key_prefix: str,\n+        label: str | None = None,\n+        *,\n+        force_refresh: bool = False,\n+    ) -\u003e dict[str, str | None]:\n+        \"\"\"List values under a key prefix while only downloading changed values.\"\"\"\n+        if not key_prefix:\n+            raise ValueError(\"key_prefix must not be empty\")\n+\n+        query = (key_prefix, label)\n+        with self._lock:\n+            known_members = self._prefix_members.get(query)\n+            if known_members is None or force_refresh:\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{key_prefix}*\",\n+                    label_filter=_label_filter(label),\n+                )\n+                current_members: set[SettingIdentity] = set()\n+                result: dict[str, str | None] = {}\n+                for setting in settings:\n+                    identity = _identity(setting.key, setting.label)\n+                    current_members.add(identity)\n+                    self._cache[identity] = _CachedSetting(\n+                        setting.value,\n+                        str(setting.etag),\n+                    )\n+                    result[setting.key] = setting.value\n+            else:\n+                # Fetch only identity and ETag metadata, then download changed values.\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{key_prefix}*\",\n+                    label_filter=_label_filter(label),\n+                    fields=[\"key\", \"label\", \"etag\"],\n+                )\n+                current_members = set()\n+                result = {}\n+                for setting in settings:\n+                    identity = _identity(setting.key, setting.label)\n+                    current_members.add(identity)\n+                    cached = self._cache.get(identity)\n+                    if cached is None or cached.etag != str(setting.etag):\n+                        self._get_setting_locked(identity, force_refresh=True)\n+                    result[setting.key] = self._cache[identity].value\n+\n+            self._prefix_members[query] = current_members\n+            self._remove_unreferenced(known_members or set(), current_members)\n+            return result\n+\n+    def _remove_unreferenced(\n+        self,\n+        previous_members: set[SettingIdentity],\n+        current_members: set[SettingIdentity],\n+    ) -\u003e None:\n+        for identity in previous_members - current_members:\n+            referenced_by_prefix = any(\n+                identity in members for members in self._prefix_members.values()\n+            )\n+            if identity not in self._exact_queries and not referenced_by_prefix:\n+                self._cache.pop(identity, None)\n+\n+    def refresh_all(self) -\u003e None:\n+        \"\"\"Fully reload every exact key and prefix previously requested.\"\"\"\n+        with self._lock:\n+            exact_queries = tuple(self._exact_queries)\n+            prefix_queries = tuple(self._prefix_members)\n+\n+        for identity in exact_queries:\n+            try:\n+                self.get_setting(*identity, force_refresh=True)\n+            except ResourceNotFoundError:\n+                logger.info(\"Cached setting was deleted: key=%s label=%s\", *identity)\n+\n+        for prefix, label in prefix_queries:\n+            self.list_settings(prefix, label, force_refresh=True)\n+\n+    def clear_cache(self) -\u003e None:\n+        \"\"\"Clear values and remembered queries.\"\"\"\n+        with self._lock:\n+            self._cache.clear()\n+            self._exact_queries.clear()\n+            self._prefix_members.clear()\n+\n+\n+class AsyncConfigurationService:\n+    \"\"\"Async counterpart to :class:`ConfigurationService`.\"\"\"\n+\n+    def __init__(self, client: AsyncAzureAppConfigurationClient) -\u003e None:\n+        self._client = client\n+        self._cache: dict[SettingIdentity, _CachedSetting] = {}\n+        self._exact_queries: set[SettingIdentity] = set()\n+        self._prefix_members: dict[PrefixQuery, set[SettingIdentity]] = {}\n+        self._lock = asyncio.Lock()\n+\n+    async def get_setting(\n+        self,\n+        key: str,\n+        label: str | None = None,\n+        *,\n+        force_refresh: bool = False,\n+    ) -\u003e str | None:\n+        \"\"\"Get one value, optionally selecting an environment label.\"\"\"\n+        if not key:\n+            raise ValueError(\"key must not be empty\")\n+\n+        identity = _identity(key, label)\n+        async with self._lock:\n+            self._exact_queries.add(identity)\n+            return await self._get_setting_locked(identity, force_refresh=force_refresh)\n+\n+    async def _get_setting_locked(\n+        self,\n+        identity: SettingIdentity,\n+        *,\n+        force_refresh: bool,\n+    ) -\u003e str | None:\n+        key, label = identity\n+        cached = self._cache.get(identity)\n+        request_options: dict[str, object] = {}\n+        if cached is not None and not force_refresh:\n+            request_options = {\n+                \"etag\": cached.etag,\n+                \"match_condition\": MatchConditions.IfModified,\n+            }\n+\n+        try:\n+            setting = await self._client.get_configuration_setting(\n+                key=key,\n+                label=label,\n+                **request_options,\n+            )\n+        except ResourceNotModifiedError:\n+            if cached is None:\n+                raise RuntimeError(\"received a not-modified response without a cached value\")\n+            return cached.value\n+        except ResourceNotFoundError:\n+            self._cache.pop(identity, None)\n+            raise\n+\n+        if setting is None:\n+            if cached is None:\n+                raise RuntimeError(\"App Configuration returned no setting\")\n+            return cached.value\n+\n+        self._cache[identity] = _CachedSetting(setting.value, str(setting.etag))\n+        return setting.value\n+\n+    async def list_settings(\n+        self,\n+        key_prefix: str,\n+        label: str | None = None,\n+        *,\n+        force_refresh: bool = False,\n+    ) -\u003e dict[str, str | None]:\n+        \"\"\"List values under a key prefix while only downloading changed values.\"\"\"\n+        if not key_prefix:\n+            raise ValueError(\"key_prefix must not be empty\")\n+\n+        query = (key_prefix, label)\n+        async with self._lock:\n+            known_members = self._prefix_members.get(query)\n+            if known_members is None or force_refresh:\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{key_prefix}*\",\n+                    label_filter=_label_filter(label),\n+                )\n+                current_members: set[SettingIdentity] = set()\n+                result: dict[str, str | None] = {}\n+                async for setting in settings:\n+                    identity = _identity(setting.key, setting.label)\n+                    current_members.add(identity)\n+                    self._cache[identity] = _CachedSetting(\n+                        setting.value,\n+                        str(setting.etag),\n+                    )\n+                    result[setting.key] = setting.value\n+            else:\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{key_prefix}*\",\n+                    label_filter=_label_filter(label),\n+                    fields=[\"key\", \"label\", \"etag\"],\n+                )\n+                current_members = set()\n+                result = {}\n+                async for setting in settings:\n+                    identity = _identity(setting.key, setting.label)\n+                    current_members.add(identity)\n+                    cached = self._cache.get(identity)\n+                    if cached is None or cached.etag != str(setting.etag):\n+                        await self._get_setting_locked(identity, force_refresh=True)\n+                    result[setting.key] = self._cache[identity].value\n+\n+            self._prefix_members[query] = current_members\n+            self._remove_unreferenced(known_members or set(), current_members)\n+            return result\n+\n+    def _remove_unreferenced(\n+        self,\n+        previous_members: set[SettingIdentity],\n+        current_members: set[SettingIdentity],\n+    ) -\u003e None:\n+        for identity in previous_members - current_members:\n+            referenced_by_prefix = any(\n+                identity in members for members in self._prefix_members.values()\n+            )\n+            if identity not in self._exact_queries and not referenced_by_prefix:\n+                self._cache.pop(identity, None)\n+\n+    async def refresh_all(self) -\u003e None:\n+        \"\"\"Fully reload every exact key and prefix previously requested.\"\"\"\n+        async with self._lock:\n+            exact_queries = tuple(self._exact_queries)\n+            prefix_queries = tuple(self._prefix_members)\n+\n+        for identity in exact_queries:\n+            try:\n+                await self.get_setting(*identity, force_refresh=True)\n+            except ResourceNotFoundError:\n+                logger.info(\"Cached setting was deleted: key=%s label=%s\", *identity)\n+\n+        for prefix, label in prefix_queries:\n+            await self.list_settings(prefix, label, force_refresh=True)\n+\n+    async def clear_cache(self) -\u003e None:\n+        \"\"\"Clear values and remembered queries.\"\"\"\n+        async with self._lock:\n+            self._cache.clear()\n+            self._exact_queries.clear()\n+            self._prefix_members.clear()\n*** Add File: appconfig_demo/feature_flags.py\n+\"\"\"Feature flag evaluation for Azure App Configuration JSON payloads.\"\"\"\n+\n+from __future__ import annotations\n+\n+import hashlib\n+import json\n+from collections.abc import Mapping\n+from typing import Any\n+\n+from .configuration_service import AsyncConfigurationService, ConfigurationService\n+\n+FEATURE_FLAG_PREFIX = \".appconfig.featureflag/\"\n+PERCENTAGE_FILTER_NAMES = frozenset(\n+    {\n+        \"microsoft.percentage\",\n+        \"percentage\",\n+    }\n+)\n+\n+\n+class FeatureFlagFormatError(ValueError):\n+    \"\"\"Raised when an App Configuration feature flag payload is malformed.\"\"\"\n+\n+\n+def _parse_flag(payload: str | None, flag_name: str) -\u003e Mapping[str, Any]:\n+    if payload is None:\n+        raise FeatureFlagFormatError(f\"feature flag {flag_name!r} has no JSON payload\")\n+    try:\n+        parsed = json.loads(payload)\n+    except json.JSONDecodeError as error:\n+        raise FeatureFlagFormatError(\n+            f\"feature flag {flag_name!r} contains invalid JSON\"\n+        ) from error\n+    if not isinstance(parsed, dict):\n+        raise FeatureFlagFormatError(\n+            f\"feature flag {flag_name!r} must contain a JSON object\"\n+        )\n+    return parsed\n+\n+\n+def _percentage(parameters: object, flag_name: str) -\u003e float:\n+    if not isinstance(parameters, dict):\n+        raise FeatureFlagFormatError(\n+            f\"percentage filter for {flag_name!r} must have parameters\"\n+        )\n+    raw_value = parameters.get(\"Value\", parameters.get(\"value\"))\n+    if isinstance(raw_value, bool):\n+        raise FeatureFlagFormatError(\n+            f\"percentage for {flag_name!r} must be a number from 0 to 100\"\n+        )\n+    try:\n+        value = float(raw_value)\n+    except (TypeError, ValueError) as error:\n+        raise FeatureFlagFormatError(\n+            f\"percentage for {flag_name!r} must be a number from 0 to 100\"\n+        ) from error\n+    if not 0 \u003c= value \u003c= 100:\n+        raise FeatureFlagFormatError(\n+            f\"percentage for {flag_name!r} must be between 0 and 100\"\n+        )\n+    return value\n+\n+\n+def _is_in_rollout(flag_id: str, user_id: str, percentage: float) -\u003e bool:\n+    digest = hashlib.sha256(f\"{flag_id}:{user_id}\".encode(\"utf-8\")).digest()\n+    bucket = int.from_bytes(digest[:8], byteorder=\"big\") % 10_000\n+    return bucket \u003c round(percentage * 100)\n+\n+\n+def _evaluate(payload: str | None, requested_name: str, user_id: str | None) -\u003e bool:\n+    flag = _parse_flag(payload, requested_name)\n+    if flag.get(\"enabled\") is not True:\n+        return False\n+\n+    flag_id = flag.get(\"id\", requested_name)\n+    if not isinstance(flag_id, str):\n+        raise FeatureFlagFormatError(f\"feature flag {requested_name!r} has an invalid id\")\n+\n+    conditions = flag.get(\"conditions\", {})\n+    if not isinstance(conditions, dict):\n+        raise FeatureFlagFormatError(\n+            f\"feature flag {requested_name!r} has invalid conditions\"\n+        )\n+    filters = conditions.get(\"client_filters\", [])\n+    if not isinstance(filters, list):\n+        raise FeatureFlagFormatError(\n+            f\"feature flag {requested_name!r} has invalid client_filters\"\n+        )\n+    if not filters:\n+        return True\n+\n+    for client_filter in filters:\n+        if not isinstance(client_filter, dict):\n+            raise FeatureFlagFormatError(\n+                f\"feature flag {requested_name!r} contains an invalid filter\"\n+            )\n+        name = client_filter.get(\"name\")\n+        if isinstance(name, str) and name.casefold() in PERCENTAGE_FILTER_NAMES:\n+            if user_id is None:\n+                return False\n+            rollout = _percentage(client_filter.get(\"parameters\"), requested_name)\n+            if _is_in_rollout(flag_id, user_id, rollout):\n+                return True\n+\n+    # Client filters have OR semantics; unsupported filters fail closed.\n+    return False\n+\n+\n+class FeatureFlagEvaluator:\n+    \"\"\"Evaluate feature flags through the sync configuration service.\"\"\"\n+\n+    def __init__(self, configuration: ConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    def is_enabled(\n+        self,\n+        flag_name: str,\n+        user_id: str | None = None,\n+        label: str | None = None,\n+    ) -\u003e bool:\n+        if not flag_name:\n+            raise ValueError(\"flag_name must not be empty\")\n+        payload = self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_name}\",\n+            label,\n+        )\n+        return _evaluate(payload, flag_name, user_id)\n+\n+\n+class AsyncFeatureFlagEvaluator:\n+    \"\"\"Evaluate feature flags through the async configuration service.\"\"\"\n+\n+    def __init__(self, configuration: AsyncConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    async def is_enabled(\n+        self,\n+        flag_name: str,\n+        user_id: str | None = None,\n+        label: str | None = None,\n+    ) -\u003e bool:\n+        if not flag_name:\n+            raise ValueError(\"flag_name must not be empty\")\n+        payload = await self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_name}\",\n+            label,\n+        )\n+        return _evaluate(payload, flag_name, user_id)\n*** Add File: appconfig_demo/watcher.py\n+\"\"\"Sentinel-based configuration watchers.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+from collections.abc import Awaitable, Callable, Sequence\n+from dataclasses import dataclass\n+from threading import Event\n+\n+from azure.core.exceptions import ResourceNotFoundError\n+\n+from .configuration_service import AsyncConfigurationService, ConfigurationService\n+\n+logger = logging.getLogger(__name__)\n+\n+_UNSET = object()\n+_MISSING = object()\n+\n+\n+@dataclass(frozen=True)\n+class SentinelKey:\n+    \"\"\"A sentinel key and its optional App Configuration label.\"\"\"\n+\n+    key: str\n+    label: str | None = None\n+\n+    def __post_init__(self) -\u003e None:\n+        if not self.key:\n+            raise ValueError(\"sentinel key must not be empty\")\n+\n+\n+class ConfigurationWatcher:\n+    \"\"\"Poll sentinel keys and fully refresh the sync cache after a change.\"\"\"\n+\n+    def __init__(\n+        self,\n+        configuration: ConfigurationService,\n+        sentinels: Sequence[SentinelKey],\n+        polling_interval: float,\n+        on_refresh: Callable[[set[SentinelKey]], None] | None = None,\n+    ) -\u003e None:\n+        if not sentinels:\n+            raise ValueError(\"at least one sentinel key is required\")\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        self._configuration = configuration\n+        self._sentinels = tuple(sentinels)\n+        self._polling_interval = polling_interval\n+        self._on_refresh = on_refresh\n+        self._last_values: dict[SentinelKey, object] = {\n+            sentinel: _UNSET for sentinel in sentinels\n+        }\n+\n+    def poll_once(self) -\u003e set[SentinelKey]:\n+        \"\"\"Check all sentinels once and return those whose values changed.\"\"\"\n+        changed: set[SentinelKey] = set()\n+        for sentinel in self._sentinels:\n+            try:\n+                value: object = self._configuration.get_setting(\n+                    sentinel.key,\n+                    sentinel.label,\n+                )\n+            except ResourceNotFoundError:\n+                value = _MISSING\n+\n+            previous = self._last_values[sentinel]\n+            if previous is not _UNSET and previous != value:\n+                changed.add(sentinel)\n+            self._last_values[sentinel] = value\n+\n+        if changed:\n+            logger.info(\"Sentinel change detected; refreshing all cached configuration\")\n+            self._configuration.refresh_all()\n+            if self._on_refresh is not None:\n+                self._on_refresh(changed)\n+        return changed\n+\n+    def run(\n+        self,\n+        stop_event: Event | None = None,\n+        *,\n+        max_polls: int | None = None,\n+    ) -\u003e None:\n+        \"\"\"Poll until stopped, or until max_polls is reached for finite demos.\"\"\"\n+        if max_polls is not None and max_polls \u003c= 0:\n+            raise ValueError(\"max_polls must be greater than zero\")\n+        stop_event = stop_event or Event()\n+        polls = 0\n+        while not stop_event.is_set():\n+            self.poll_once()\n+            polls += 1\n+            if max_polls is not None and polls \u003e= max_polls:\n+                return\n+            stop_event.wait(self._polling_interval)\n+\n+\n+AsyncRefreshCallback = Callable[[set[SentinelKey]], Awaitable[None] | None]\n+\n+\n+class AsyncConfigurationWatcher:\n+    \"\"\"Poll sentinel keys and fully refresh the async cache after a change.\"\"\"\n+\n+    def __init__(\n+        self,\n+        configuration: AsyncConfigurationService,\n+        sentinels: Sequence[SentinelKey],\n+        polling_interval: float,\n+        on_refresh: AsyncRefreshCallback | None = None,\n+    ) -\u003e None:\n+        if not sentinels:\n+            raise ValueError(\"at least one sentinel key is required\")\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        self._configuration = configuration\n+        self._sentinels = tuple(sentinels)\n+        self._polling_interval = polling_interval\n+        self._on_refresh = on_refresh\n+        self._last_values: dict[SentinelKey, object] = {\n+            sentinel: _UNSET for sentinel in sentinels\n+        }\n+\n+    async def poll_once(self) -\u003e set[SentinelKey]:\n+        \"\"\"Check all sentinels once and return those whose values changed.\"\"\"\n+        changed: set[SentinelKey] = set()\n+        for sentinel in self._sentinels:\n+            try:\n+                value: object = await self._configuration.get_setting(\n+                    sentinel.key,\n+                    sentinel.label,\n+                )\n+            except ResourceNotFoundError:\n+                value = _MISSING\n+\n+            previous = self._last_values[sentinel]\n+            if previous is not _UNSET and previous != value:\n+                changed.add(sentinel)\n+            self._last_values[sentinel] = value\n+\n+        if changed:\n+            logger.info(\"Sentinel change detected; refreshing all cached configuration\")\n+            await self._configuration.refresh_all()\n+            if self._on_refresh is not None:\n+                callback_result = self._on_refresh(changed)\n+                if callback_result is not None:\n+                    await callback_result\n+        return changed\n+\n+    async def run(\n+        self,\n+        stop_event: asyncio.Event | None = None,\n+        *,\n+        max_polls: int | None = None,\n+    ) -\u003e None:\n+        \"\"\"Poll until stopped, or until max_polls is reached for finite demos.\"\"\"\n+        if max_polls is not None and max_polls \u003c= 0:\n+            raise ValueError(\"max_polls must be greater than zero\")\n+        stop_event = stop_event or asyncio.Event()\n+        polls = 0\n+        while not stop_event.is_set():\n+            await self.poll_once()\n+            polls += 1\n+            if max_polls is not None and polls \u003e= max_polls:\n+                return\n+            try:\n+                await asyncio.wait_for(\n+                    stop_event.wait(),\n+                    timeout=self._polling_interval,\n+                )\n+            except TimeoutError:\n+                pass\n*** Add File: main.py\n+\"\"\"Run sync and async Azure App Configuration demonstrations.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import logging\n+import os\n+\n+from azure.appconfiguration import AzureAppConfigurationClient\n+from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncAzureAppConfigurationClient\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+\n+from appconfig_demo import (\n+    AsyncConfigurationService,\n+    AsyncConfigurationWatcher,\n+    AsyncFeatureFlagEvaluator,\n+    ConfigurationService,\n+    ConfigurationWatcher,\n+    FeatureFlagEvaluator,\n+    SentinelKey,\n+)\n+\n+logging.basicConfig(level=os.getenv(\"LOG_LEVEL\", \"INFO\"))\n+logger = logging.getLogger(__name__)\n+\n+\n+def _endpoint() -\u003e str:\n+    endpoint = os.getenv(\"AZURE_APPCONFIGURATION_ENDPOINT\")\n+    if not endpoint:\n+        raise RuntimeError(\"AZURE_APPCONFIGURATION_ENDPOINT must be set\")\n+    return endpoint\n+\n+\n+def _demo_options() -\u003e tuple[str, str, str, str, list[str], float, int]:\n+    key = os.getenv(\"DEMO_CONFIG_KEY\", \"demo:message\")\n+    label = os.getenv(\"DEMO_CONFIG_LABEL\", \"production\")\n+    prefix = os.getenv(\"DEMO_CONFIG_PREFIX\", \"demo:\")\n+    flag = os.getenv(\"DEMO_FEATURE_FLAG\", \"gradual-rollout\")\n+    users = [\n+        user.strip()\n+        for user in os.getenv(\"DEMO_USER_IDS\", \"alice,bob,charlie,diana\").split(\",\")\n+        if user.strip()\n+    ]\n+    interval = float(os.getenv(\"DEMO_POLL_INTERVAL\", \"5\"))\n+    polls = int(os.getenv(\"DEMO_MAX_POLLS\", \"3\"))\n+    return key, label, prefix, flag, users, interval, polls\n+\n+\n+def run_sync_demo(endpoint: str) -\u003e None:\n+    \"\"\"Run the synchronous client, evaluator, and watcher.\"\"\"\n+    key, label, prefix, flag, users, interval, polls = _demo_options()\n+    sentinel = SentinelKey(os.getenv(\"DEMO_SENTINEL_KEY\", \"demo:sentinel\"), label)\n+\n+    print(\"\\n=== Synchronous demo ===\")\n+    with DefaultAzureCredential() as credential:\n+        with AzureAppConfigurationClient(\n+            base_url=endpoint,\n+            credential=credential,\n+        ) as client:\n+            configuration = ConfigurationService(client)\n+            flags = FeatureFlagEvaluator(configuration)\n+\n+            print(f\"{key} (unlabeled): {configuration.get_setting(key)}\")\n+            print(f\"{key} ({label}): {configuration.get_setting(key, label)}\")\n+            print(\n+                f\"{prefix}* ({label}): \"\n+                f\"{configuration.list_settings(prefix, label)}\"\n+            )\n+            for user_id in users:\n+                enabled = flags.is_enabled(flag, user_id=user_id, label=label)\n+                print(f\"feature {flag!r} for {user_id}: {enabled}\")\n+\n+            def refreshed(changed: set[SentinelKey]) -\u003e None:\n+                print(f\"sync refresh triggered by: {sorted(s.key for s in changed)}\")\n+\n+            print(\n+                f\"Watching {sentinel.key!r} for {polls} polls; \"\n+                \"change its value to trigger a refresh.\"\n+            )\n+            ConfigurationWatcher(\n+                configuration,\n+                [sentinel],\n+                interval,\n+                on_refresh=refreshed,\n+            ).run(max_polls=polls)\n+\n+\n+async def run_async_demo(endpoint: str) -\u003e None:\n+    \"\"\"Run the asynchronous client, evaluator, and watcher.\"\"\"\n+    key, label, prefix, flag, users, interval, polls = _demo_options()\n+    sentinel = SentinelKey(os.getenv(\"DEMO_SENTINEL_KEY\", \"demo:sentinel\"), label)\n+\n+    print(\"\\n=== Asynchronous demo ===\")\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncAzureAppConfigurationClient(\n+            base_url=endpoint,\n+            credential=credential,\n+        ) as client:\n+            configuration = AsyncConfigurationService(client)\n+            flags = AsyncFeatureFlagEvaluator(configuration)\n+\n+            print(f\"{key} (unlabeled): {await configuration.get_setting(key)}\")\n+            print(f\"{key} ({label}): {await configuration.get_setting(key, label)}\")\n+            print(\n+                f\"{prefix}* ({label}): \"\n+                f\"{await configuration.list_settings(prefix, label)}\"\n+            )\n+            for user_id in users:\n+                enabled = await flags.is_enabled(flag, user_id=user_id, label=label)\n+                print(f\"feature {flag!r} for {user_id}: {enabled}\")\n+\n+            async def refreshed(changed: set[SentinelKey]) -\u003e None:\n+                print(f\"async refresh triggered by: {sorted(s.key for s in changed)}\")\n+\n+            print(\n+                f\"Watching {sentinel.key!r} for {polls} polls; \"\n+                \"change its value to trigger a refresh.\"\n+            )\n+            await AsyncConfigurationWatcher(\n+                configuration,\n+                [sentinel],\n+                interval,\n+                on_refresh=refreshed,\n+            ).run(max_polls=polls)\n+\n+\n+def main() -\u003e None:\n+    endpoint = _endpoint()\n+    run_sync_demo(endpoint)\n+    asyncio.run(run_async_demo(endpoint))\n+\n+\n+if __name__ == \"__main__\":\n+    try:\n+        main()\n+    except (ValueError, RuntimeError):\n+        logger.exception(\"Demo configuration is invalid\")\n+        raise\n*** Add File: tests/test_appconfig_demo.py\n+from __future__ import annotations\n+\n+import json\n+import unittest\n+from types import SimpleNamespace\n+\n+from azure.core import MatchConditions\n+\n+from appconfig_demo.configuration_service import (\n+    AsyncConfigurationService,\n+    ConfigurationService,\n+)\n+from appconfig_demo.feature_flags import (\n+    AsyncFeatureFlagEvaluator,\n+    FeatureFlagEvaluator,\n+    FeatureFlagFormatError,\n+)\n+from appconfig_demo.watcher import (\n+    AsyncConfigurationWatcher,\n+    ConfigurationWatcher,\n+    SentinelKey,\n+)\n+\n+\n+def setting(key: str, value: str | None, etag: str, label: str | None = None):\n+    return SimpleNamespace(key=key, value=value, etag=etag, label=label)\n+\n+\n+class FakeSyncClient:\n+    def __init__(self) -\u003e None:\n+        self.settings = {\n+            (\"app:one\", None): setting(\"app:one\", \"1\", \"e1\"),\n+            (\"app:two\", None): setting(\"app:two\", \"2\", \"e2\"),\n+        }\n+        self.get_calls: list[dict[str, object]] = []\n+        self.list_calls: list[dict[str, object]] = []\n+\n+    def get_configuration_setting(self, **kwargs):\n+        self.get_calls.append(kwargs)\n+        current = self.settings[(kwargs[\"key\"], kwargs.get(\"label\"))]\n+        if (\n+            kwargs.get(\"match_condition\") is MatchConditions.IfModified\n+            and kwargs.get(\"etag\") == current.etag\n+        ):\n+            return None\n+        return current\n+\n+    def list_configuration_settings(self, **kwargs):\n+        self.list_calls.append(kwargs)\n+        prefix = str(kwargs[\"key_filter\"])[:-1]\n+        label_filter = kwargs[\"label_filter\"]\n+        selected = [\n+            item\n+            for item in self.settings.values()\n+            if item.key.startswith(prefix)\n+            and (item.label if item.label is not None else \"\\0\") == label_filter\n+        ]\n+        if \"fields\" in kwargs:\n+            return [\n+                setting(item.key, None, item.etag, item.label) for item in selected\n+            ]\n+        return selected\n+\n+\n+class SyncTests(unittest.TestCase):\n+    def test_single_setting_uses_conditional_request_after_first_read(self) -\u003e None:\n+        client = FakeSyncClient()\n+        service = ConfigurationService(client)\n+\n+        self.assertEqual(\"1\", service.get_setting(\"app:one\"))\n+        self.assertEqual(\"1\", service.get_setting(\"app:one\"))\n+\n+        self.assertEqual(2, len(client.get_calls))\n+        self.assertEqual(\"e1\", client.get_calls[1][\"etag\"])\n+        self.assertIs(\n+            MatchConditions.IfModified,\n+            client.get_calls[1][\"match_condition\"],\n+        )\n+\n+    def test_prefix_listing_only_downloads_changed_values(self) -\u003e None:\n+        client = FakeSyncClient()\n+        service = ConfigurationService(client)\n+\n+        self.assertEqual(\n+            {\"app:one\": \"1\", \"app:two\": \"2\"},\n+            service.list_settings(\"app:\"),\n+        )\n+        client.settings[(\"app:two\", None)] = setting(\"app:two\", \"updated\", \"e3\")\n+        self.assertEqual(\n+            {\"app:one\": \"1\", \"app:two\": \"updated\"},\n+            service.list_settings(\"app:\"),\n+        )\n+\n+        self.assertEqual([\"key\", \"label\", \"etag\"], client.list_calls[1][\"fields\"])\n+        self.assertEqual([\"app:two\"], [call[\"key\"] for call in client.get_calls])\n+\n+    def test_percentage_rollout_is_deterministic(self) -\u003e None:\n+        client = FakeSyncClient()\n+        payload = json.dumps(\n+            {\n+                \"id\": \"gradual\",\n+                \"enabled\": True,\n+                \"conditions\": {\n+                    \"client_filters\": [\n+                        {\n+                            \"name\": \"Microsoft.Percentage\",\n+                            \"parameters\": {\"Value\": 30},\n+                        }\n+                    ]\n+                },\n+            }\n+        )\n+        client.settings[(\".appconfig.featureflag/gradual\", None)] = setting(\n+            \".appconfig.featureflag/gradual\",\n+            payload,\n+            \"flag1\",\n+        )\n+        evaluator = FeatureFlagEvaluator(ConfigurationService(client))\n+\n+        first = evaluator.is_enabled(\"gradual\", user_id=\"alice\")\n+        self.assertEqual(first, evaluator.is_enabled(\"gradual\", user_id=\"alice\"))\n+        outcomes = {\n+            evaluator.is_enabled(\"gradual\", user_id=f\"user-{index}\")\n+            for index in range(100)\n+        }\n+        self.assertEqual({False, True}, outcomes)\n+\n+    def test_invalid_percentage_fails_explicitly(self) -\u003e None:\n+        client = FakeSyncClient()\n+        payload = json.dumps(\n+            {\n+                \"id\": \"bad\",\n+                \"enabled\": True,\n+                \"conditions\": {\n+                    \"client_filters\": [\n+                        {\n+                            \"name\": \"Microsoft.Percentage\",\n+                            \"parameters\": {\"Value\": 101},\n+                        }\n+                    ]\n+                },\n+            }\n+        )\n+        client.settings[(\".appconfig.featureflag/bad\", None)] = setting(\n+            \".appconfig.featureflag/bad\",\n+            payload,\n+            \"flag1\",\n+        )\n+        evaluator = FeatureFlagEvaluator(ConfigurationService(client))\n+\n+        with self.assertRaises(FeatureFlagFormatError):\n+            evaluator.is_enabled(\"bad\", user_id=\"alice\")\n+\n+    def test_sentinel_change_triggers_full_refresh(self) -\u003e None:\n+        values = iter([\"v1\", \"v2\"])\n+\n+        class FakeService:\n+            refreshes = 0\n+\n+            def get_setting(self, key, label=None):\n+                return next(values)\n+\n+            def refresh_all(self):\n+                self.refreshes += 1\n+\n+        service = FakeService()\n+        watcher = ConfigurationWatcher(service, [SentinelKey(\"sentinel\")], 0.01)\n+\n+        self.assertEqual(set(), watcher.poll_once())\n+        self.assertEqual({SentinelKey(\"sentinel\")}, watcher.poll_once())\n+        self.assertEqual(1, service.refreshes)\n+\n+\n+class AsyncPage:\n+    def __init__(self, items) -\u003e None:\n+        self._items = iter(items)\n+\n+    def __aiter__(self):\n+        return self\n+\n+    async def __anext__(self):\n+        try:\n+            return next(self._items)\n+        except StopIteration as error:\n+            raise StopAsyncIteration from error\n+\n+\n+class FakeAsyncClient(FakeSyncClient):\n+    async def get_configuration_setting(self, **kwargs):\n+        return super().get_configuration_setting(**kwargs)\n+\n+    def list_configuration_settings(self, **kwargs):\n+        return AsyncPage(super().list_configuration_settings(**kwargs))\n+\n+\n+class AsyncTests(unittest.IsolatedAsyncioTestCase):\n+    async def test_async_service_and_evaluator(self) -\u003e None:\n+        client = FakeAsyncClient()\n+        payload = json.dumps(\n+            {\n+                \"id\": \"always\",\n+                \"enabled\": True,\n+                \"conditions\": {\"client_filters\": []},\n+            }\n+        )\n+        client.settings[(\".appconfig.featureflag/always\", \"production\")] = setting(\n+            \".appconfig.featureflag/always\",\n+            payload,\n+            \"f1\",\n+            \"production\",\n+        )\n+        service = AsyncConfigurationService(client)\n+        evaluator = AsyncFeatureFlagEvaluator(service)\n+\n+        self.assertTrue(await evaluator.is_enabled(\"always\", label=\"production\"))\n+        self.assertTrue(await evaluator.is_enabled(\"always\", label=\"production\"))\n+        self.assertIs(\n+            MatchConditions.IfModified,\n+            client.get_calls[1][\"match_condition\"],\n+        )\n+\n+    async def test_async_watcher_refreshes(self) -\u003e None:\n+        class FakeService:\n+            def __init__(self) -\u003e None:\n+                self.values = iter([\"v1\", \"v2\"])\n+                self.refreshes = 0\n+\n+            async def get_setting(self, key, label=None):\n+                return next(self.values)\n+\n+            async def refresh_all(self):\n+                self.refreshes += 1\n+\n+        service = FakeService()\n+        watcher = AsyncConfigurationWatcher(\n+            service,\n+            [SentinelKey(\"sentinel\")],\n+            0.01,\n+        )\n+\n+        self.assertEqual(set(), await watcher.poll_once())\n+        self.assertEqual({SentinelKey(\"sentinel\")}, await watcher.poll_once())\n+        self.assertEqual(1, service.refreshes)\n*** Add File: requirements.txt\n+azure-appconfiguration\u003e=1.7.1,\u003c2.0.0\n+azure-identity\u003e=1.23.0,\u003c2.0.0\n*** Add File: README.md\n+# Azure App Configuration feature flags\n+\n+This project demonstrates synchronous and asynchronous configuration access,\n+ETag-aware caching, deterministic percentage feature rollouts, and\n+sentinel-coordinated cache refreshes.\n+\n+## Setup\n+\n+Use Python 3.10 or newer:\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+python -m pip install -r requirements.txt\n+```\n+\n+Authenticate locally with a credential supported by `DefaultAzureCredential`,\n+then set the App Configuration data-plane endpoint:\n+\n+```powershell\n+$env:AZURE_APPCONFIGURATION_ENDPOINT = \"https://\u003cstore\u003e.azconfig.io\"\n+python main.py\n+```\n+\n+The identity needs the **App Configuration Data Reader** role. For an\n+Azure-hosted production deployment, use managed identity and set\n+`AZURE_TOKEN_CREDENTIALS=prod` to restrict the credential chain.\n+\n+The demo defaults can be changed with these environment variables:\n+\n+| Variable | Default |\n+| --- | --- |\n+| `DEMO_CONFIG_KEY` | `demo:message` |\n+| `DEMO_CONFIG_LABEL` | `production` |\n+| `DEMO_CONFIG_PREFIX` | `demo:` |\n+| `DEMO_FEATURE_FLAG` | `gradual-rollout` |\n+| `DEMO_USER_IDS` | `alice,bob,charlie,diana` |\n+| `DEMO_SENTINEL_KEY` | `demo:sentinel` |\n+| `DEMO_POLL_INTERVAL` | `5` |\n+| `DEMO_MAX_POLLS` | `3` |\n+\n+The main script runs the sync demo first, followed by the async demo. During\n+each finite watch window, change the sentinel value to trigger a full cache\n+refresh.\n+\n+Feature flags are read from `.appconfig.featureflag/\u003cname\u003e`. Percentage rollout\n+uses the `Microsoft.Percentage` filter and hashes `\u003cflag-id\u003e:\u003cuser-id\u003e` with\n+SHA-256, giving each user a stable bucket.\n+\n+## Tests\n+\n+Tests use fake clients and do not connect to Azure:\n+\n+```powershell\n+python -m unittest discover -s tests -v\n+```\n+\n+References:\n+\n+- [Azure App Configuration Python SDK](https://learn.microsoft.com/python/api/azure-appconfiguration/azure.appconfiguration.azureappconfigurationclient)\n+- [Azure App Configuration Python quickstart](https://learn.microsoft.com/azure/azure-app-configuration/quickstart-python)\n+- [DefaultAzureCredential](https://learn.microsoft.com/python/api/azure-identity/azure.identity.defaultazurecredential)\n*** End Patch\n"
```

### 🔧 26. powershell

**Input:**

```json
{"command":"python -m compileall -q appconfig_demo main.py tests; if ($LASTEXITCODE -ne 0) { exit $LASTEXITCODE }; python -m unittest discover -s tests -v","description":"Compile project and run unit tests","initial_wait":120,"mode":"sync"}
```

### 🔧 27. powershell

**Input:**

```json
{"command":"python -c \"import inspect, importlib.metadata as m; from azure.appconfiguration import AzureAppConfigurationClient; from azure.appconfiguration.aio import AzureAppConfigurationClient as A; from azure.core import MatchConditions; print('azure-appconfiguration', m.version('azure-appconfiguration')); print('azure-identity', m.version('azure-identity')); print(inspect.signature(AzureAppConfigurationClient.list_configuration_settings)); print(inspect.signature(A.list_configuration_settings)); print(MatchConditions.IfModified.value)\"","description":"Check installed Azure SDK compatibility","initial_wait":30,"mode":"sync"}
```

## Generated Files

- `README.md`
- `appconfig_demo\__init__.py`
- `appconfig_demo\configuration_service.py`
- `appconfig_demo\feature_flags.py`
- `appconfig_demo\watcher.py`
- `main.py`
- `requirements.txt`
- `tests\test_appconfig_demo.py`

## Copilot Response

I’m grounding the implementation in the SDK lifecycle and refresh semantics, then I’ll build the sync/async modules and local tests together.

The design will use `If-None-Match` semantics (`MatchConditions.IfModified`) for cached key reads, ETags for prefix cache reconciliation, and sentinel value transitions to coordinate full refreshes.

I’ve finalized the cache contract and watcher behavior; I’m now writing the package, executable demo, dependency manifest, and focused offline tests.

The offline suite passes. I’m doing the final SDK-compatibility pass now, including installed signatures/version bounds and generated-artifact cleanup.

## Reviewer Notes (LLM-as-Judge)

### Summary

The exception-handling criterion is not fully met because common Azure SDK HTTP failures are not handled at an appropriate application boundary.

### Strengths

- ResourceNotModifiedError is correctly used to return cached values, and ResourceNotFoundError is handled meaningfully for deleted settings and missing sentinels.

### Issues

- HttpResponseError is not imported or caught in main.py or watcher run loops, so most Azure service failures terminate the demo without contextual handling.

## Grader Results

- feature-flags.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Pass (9/9)
      - Lists settings filtered by key prefix using `key_filter`: Pass
      - Implements conditional reads using `match_condition` with `MatchConditions.IfModified` and the setting's `etag`: Pass
      - Handles 304 Not Modified (setting unchanged since last read): Pass
      - Uses `.appconfig.featureflag/` prefix for feature flag keys: Pass
      - Parses the JSON payload in feature flag setting values: Pass
      - Implements deterministic percentage rollout (consistent hash via `hashlib`, not `random`): Pass
      - Implements sentinel key watching with configurable polling interval: Pass
      - Detects sentinel value change via ETag or value comparison and triggers full refresh: Pass
      - Async version uses `azure.appconfiguration.aio.AzureAppConfigurationClient`: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Pass
  - Context Manager for Clients (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Context Manager for Clients**: Azure SDK clients that support context managers are used with `with` statements or explicitly closed.: Pass
  - Async Client Usage (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Async Client Usage**: If async operations are requested, code uses the async client variant with proper await patterns.: Pass
  - Proper Exception Handling (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**Proper Exception Handling**: Azure SDK exceptions (HttpResponseError and subclasses) are caught and handled appropriately.: Fail
  - Output Files Exist (workspace): Fail (0/1)
      - file: *.py (state=present): Fail
  - Tool Usage Verification (tool): Fail (0/1)
      - tool_used: azure (source=mcp, server=azure): Fail

## Score Breakdown

**Formula:** `Final Score = Σ(grader_score × weight) / Σ(weights)`

| Grader | Type | Score | Weight | Weighted | Contribution | Status |
|--------|------|-------|--------|----------|--------------|--------|
| `Criteria from prompt file` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Context Manager for Clients` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 20.0% | ✅ |
| `Proper Exception Handling` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 5.0000** | **62.5%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-python-feature-flags --config python-azure-tools/with-azure-tools
```

---

[← Back to Summary](../../../../../../summary.md)
