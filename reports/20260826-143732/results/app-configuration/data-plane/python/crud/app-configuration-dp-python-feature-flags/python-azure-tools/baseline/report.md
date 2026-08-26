# Evaluation Report: app-configuration-dp-python-feature-flags

**Config:** python-azure-tools/baseline | **Result:** ❌ FAILED | **Duration:** 303.8s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-python-feature-flags` |
| Config | python-azure-tools/baseline |
| Result | ❌ FAILED |
| Score | 13/16 |
| Duration | 303.8s |
| Timestamp | 2026-08-26T06:53:26Z |
| Files Generated | 9 |
| Event Count | 17783 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 209.5s |
| Review | 93.2s |
| **Total** | **303.8s** |

## Configuration

- **name:** python-azure-tools/baseline
- **model:** gpt-5.6-sol

## Environment & Configuration

| Setting | Value |
|---------|-------|
| Model | gpt-5.6-sol |
| Skills Loaded | customize-cloud-agent, github-pr-media |
| Safety Boundaries | ✅ Active |
| Cloud Access | ❌ Denied |
| Token Usage | in=386637 out=12410 |
| Turn Count | 13 |

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

### 🔧 1. glob

**Input:**

```json
{"pattern":"**/*"}
```

### 🔧 2. powershell

**Input:**

```json
{"command":"Get-Location; Get-ChildItem -Force | Select-Object Name,Mode,Length","description":"Inspect workspace contents"}
```

### 🔧 3. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Add File: requirements.txt\n+azure-appconfiguration\u003e=1.7.0,\u003c2.0.0\n+azure-identity\u003e=1.17.0,\u003c2.0.0\n+\n*** Add File: appconfig_manager/__init__.py\n+\"\"\"Azure App Configuration helpers.\"\"\"\n+\n+from .configuration import AsyncConfigurationService, ConfigurationService\n+from .feature_flags import AsyncFeatureFlagEvaluator, FeatureFlagEvaluator\n+from .watcher import AsyncConfigurationWatcher, ConfigurationWatcher\n+\n+__all__ = [\n+    \"AsyncConfigurationService\",\n+    \"AsyncConfigurationWatcher\",\n+    \"AsyncFeatureFlagEvaluator\",\n+    \"ConfigurationService\",\n+    \"ConfigurationWatcher\",\n+    \"FeatureFlagEvaluator\",\n+]\n+\n*** Add File: appconfig_manager/configuration.py\n+\"\"\"Cached synchronous and asynchronous Azure App Configuration clients.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+from dataclasses import dataclass\n+from threading import RLock\n+from typing import Any\n+\n+from azure.core import MatchConditions\n+from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError\n+\n+_NULL_LABEL_FILTER = \"\\0\"\n+\n+\n+@dataclass(frozen=True)\n+class _CachedSetting:\n+    value: str | None\n+    etag: str\n+\n+\n+def _cache_key(key: str, label: str | None) -\u003e tuple[str, str | None]:\n+    return key, label\n+\n+\n+def _label_filter(label: str | None) -\u003e str:\n+    return _NULL_LABEL_FILTER if label is None else label\n+\n+\n+def _etag(setting: Any) -\u003e str:\n+    return str(setting.etag)\n+\n+\n+class ConfigurationService:\n+    \"\"\"Retrieve and cache settings with conditional ETag requests.\"\"\"\n+\n+    def __init__(self, client: Any) -\u003e None:\n+        self._client = client\n+        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}\n+        self._key_queries: set[tuple[str, str | None]] = set()\n+        self._prefix_queries: set[tuple[str, str | None]] = set()\n+        self._prefix_members: dict[\n+            tuple[str, str | None], set[tuple[str, str | None]]\n+        ] = {}\n+        self._lock = RLock()\n+\n+    def get_setting(\n+        self, key: str, label: str | None = None, *, refresh: bool = False\n+    ) -\u003e str | None:\n+        \"\"\"Return a setting value, optionally checking Azure for a newer ETag.\"\"\"\n+        with self._lock:\n+            self._key_queries.add((key, label))\n+            value, _ = self._fetch_setting(key, label, refresh=refresh)\n+            return value\n+\n+    def get_setting_with_label(self, key: str, label: str) -\u003e str | None:\n+        \"\"\"Return a setting for a specific environment label.\"\"\"\n+        return self.get_setting(key, label)\n+\n+    def list_settings(\n+        self, prefix: str, label: str | None = None, *, refresh: bool = False\n+    ) -\u003e dict[str, str | None]:\n+        \"\"\"Return no-label or specifically labeled settings matching a key prefix.\"\"\"\n+        with self._lock:\n+            query = (prefix, label)\n+            self._prefix_queries.add(query)\n+            members = self._prefix_members.get(query)\n+            if members is not None and not refresh:\n+                return {key: self._cache[(key, item_label)].value for key, item_label in members}\n+\n+            if members is None:\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{prefix}*\",\n+                    label_filter=_label_filter(label),\n+                )\n+                new_members: set[tuple[str, str | None]] = set()\n+                for setting in settings:\n+                    item = _cache_key(setting.key, setting.label)\n+                    self._cache[item] = _CachedSetting(setting.value, _etag(setting))\n+                    new_members.add(item)\n+            else:\n+                settings = self._client.list_configuration_settings(\n+                    key_filter=f\"{prefix}*\",\n+                    label_filter=_label_filter(label),\n+                    fields=[\"key\", \"label\", \"etag\"],\n+                )\n+                new_members = set()\n+                for setting in settings:\n+                    item = _cache_key(setting.key, setting.label)\n+                    new_members.add(item)\n+                    cached = self._cache.get(item)\n+                    if cached is None or cached.etag != _etag(setting):\n+                        self._fetch_setting(setting.key, setting.label, refresh=True)\n+\n+                for deleted in members - new_members:\n+                    self._cache.pop(deleted, None)\n+\n+            self._prefix_members[query] = new_members\n+            return {key: self._cache[(key, item_label)].value for key, item_label in new_members}\n+\n+    def check_for_update(self, key: str, label: str | None = None) -\u003e bool:\n+        \"\"\"Conditionally poll a setting and report whether it changed or was deleted.\"\"\"\n+        with self._lock:\n+            item = _cache_key(key, label)\n+            if item not in self._cache:\n+                self.get_setting(key, label)\n+                return False\n+            try:\n+                _, changed = self._fetch_setting(key, label, refresh=True)\n+                return changed\n+            except ResourceNotFoundError:\n+                self._cache.pop(item, None)\n+                return True\n+\n+    def refresh_all(self) -\u003e None:\n+        \"\"\"Invalidate and reload all keys and prefixes requested so far.\"\"\"\n+        with self._lock:\n+            key_queries = tuple(self._key_queries)\n+            prefix_queries = tuple(self._prefix_queries)\n+            self._cache.clear()\n+            self._prefix_members.clear()\n+            for key, label in key_queries:\n+                try:\n+                    self._fetch_setting(key, label, refresh=False)\n+                except ResourceNotFoundError:\n+                    pass\n+            for prefix, label in prefix_queries:\n+                self.list_settings(prefix, label)\n+\n+    def _fetch_setting(\n+        self, key: str, label: str | None, *, refresh: bool\n+    ) -\u003e tuple[str | None, bool]:\n+        item = _cache_key(key, label)\n+        cached = self._cache.get(item)\n+        if cached is not None and not refresh:\n+            return cached.value, False\n+\n+        kwargs: dict[str, Any] = {\"key\": key, \"label\": label}\n+        if cached is not None:\n+            kwargs.update(\n+                etag=cached.etag,\n+                match_condition=MatchConditions.IfModified,\n+            )\n+        try:\n+            setting = self._client.get_configuration_setting(**kwargs)\n+        except ResourceNotModifiedError:\n+            return cached.value, False\n+\n+        current = _CachedSetting(setting.value, _etag(setting))\n+        self._cache[item] = current\n+        return current.value, cached is not None and current.etag != cached.etag\n+\n+\n+class AsyncConfigurationService:\n+    \"\"\"Asynchronous counterpart to :class:`ConfigurationService`.\"\"\"\n+\n+    def __init__(self, client: Any) -\u003e None:\n+        self._client = client\n+        self._cache: dict[tuple[str, str | None], _CachedSetting] = {}\n+        self._key_queries: set[tuple[str, str | None]] = set()\n+        self._prefix_queries: set[tuple[str, str | None]] = set()\n+        self._prefix_members: dict[\n+            tuple[str, str | None], set[tuple[str, str | None]]\n+        ] = {}\n+        self._lock = asyncio.Lock()\n+\n+    async def get_setting(\n+        self, key: str, label: str | None = None, *, refresh: bool = False\n+    ) -\u003e str | None:\n+        \"\"\"Return a setting value, optionally checking Azure for a newer ETag.\"\"\"\n+        async with self._lock:\n+            self._key_queries.add((key, label))\n+            value, _ = await self._fetch_setting(key, label, refresh=refresh)\n+            return value\n+\n+    async def get_setting_with_label(self, key: str, label: str) -\u003e str | None:\n+        \"\"\"Return a setting for a specific environment label.\"\"\"\n+        return await self.get_setting(key, label)\n+\n+    async def list_settings(\n+        self, prefix: str, label: str | None = None, *, refresh: bool = False\n+    ) -\u003e dict[str, str | None]:\n+        \"\"\"Return no-label or specifically labeled settings matching a key prefix.\"\"\"\n+        async with self._lock:\n+            return await self._list_settings_locked(prefix, label, refresh=refresh)\n+\n+    async def check_for_update(\n+        self, key: str, label: str | None = None\n+    ) -\u003e bool:\n+        \"\"\"Conditionally poll a setting and report whether it changed or was deleted.\"\"\"\n+        async with self._lock:\n+            item = _cache_key(key, label)\n+            if item not in self._cache:\n+                await self._fetch_setting(key, label, refresh=False)\n+                return False\n+            try:\n+                _, changed = await self._fetch_setting(key, label, refresh=True)\n+                return changed\n+            except ResourceNotFoundError:\n+                self._cache.pop(item, None)\n+                return True\n+\n+    async def refresh_all(self) -\u003e None:\n+        \"\"\"Invalidate and reload all keys and prefixes requested so far.\"\"\"\n+        async with self._lock:\n+            key_queries = tuple(self._key_queries)\n+            prefix_queries = tuple(self._prefix_queries)\n+            self._cache.clear()\n+            self._prefix_members.clear()\n+            for key, label in key_queries:\n+                try:\n+                    await self._fetch_setting(key, label, refresh=False)\n+                except ResourceNotFoundError:\n+                    pass\n+            for prefix, label in prefix_queries:\n+                await self._list_settings_locked(prefix, label, refresh=False)\n+\n+    async def _list_settings_locked(\n+        self, prefix: str, label: str | None, *, refresh: bool\n+    ) -\u003e dict[str, str | None]:\n+        query = (prefix, label)\n+        self._prefix_queries.add(query)\n+        members = self._prefix_members.get(query)\n+        if members is not None and not refresh:\n+            return {key: self._cache[(key, item_label)].value for key, item_label in members}\n+\n+        if members is None:\n+            settings = self._client.list_configuration_settings(\n+                key_filter=f\"{prefix}*\",\n+                label_filter=_label_filter(label),\n+            )\n+            new_members: set[tuple[str, str | None]] = set()\n+            async for setting in settings:\n+                item = _cache_key(setting.key, setting.label)\n+                self._cache[item] = _CachedSetting(setting.value, _etag(setting))\n+                new_members.add(item)\n+        else:\n+            settings = self._client.list_configuration_settings(\n+                key_filter=f\"{prefix}*\",\n+                label_filter=_label_filter(label),\n+                fields=[\"key\", \"label\", \"etag\"],\n+            )\n+            new_members = set()\n+            async for setting in settings:\n+                item = _cache_key(setting.key, setting.label)\n+                new_members.add(item)\n+                cached = self._cache.get(item)\n+                if cached is None or cached.etag != _etag(setting):\n+                    await self._fetch_setting(setting.key, setting.label, refresh=True)\n+\n+            for deleted in members - new_members:\n+                self._cache.pop(deleted, None)\n+\n+        self._prefix_members[query] = new_members\n+        return {key: self._cache[(key, item_label)].value for key, item_label in new_members}\n+\n+    async def _fetch_setting(\n+        self, key: str, label: str | None, *, refresh: bool\n+    ) -\u003e tuple[str | None, bool]:\n+        item = _cache_key(key, label)\n+        cached = self._cache.get(item)\n+        if cached is not None and not refresh:\n+            return cached.value, False\n+\n+        kwargs: dict[str, Any] = {\"key\": key, \"label\": label}\n+        if cached is not None:\n+            kwargs.update(\n+                etag=cached.etag,\n+                match_condition=MatchConditions.IfModified,\n+            )\n+        try:\n+            setting = await self._client.get_configuration_setting(**kwargs)\n+        except ResourceNotModifiedError:\n+            return cached.value, False\n+\n+        current = _CachedSetting(setting.value, _etag(setting))\n+        self._cache[item] = current\n+        return current.value, cached is not None and current.etag != cached.etag\n+\n*** Add File: appconfig_manager/feature_flags.py\n+\"\"\"Feature flag evaluation, including deterministic percentage rollout.\"\"\"\n+\n+from __future__ import annotations\n+\n+import hashlib\n+import json\n+from typing import Any\n+\n+from .configuration import AsyncConfigurationService, ConfigurationService\n+\n+FEATURE_FLAG_PREFIX = \".appconfig.featureflag/\"\n+_PERCENTAGE_FILTER_NAMES = {\"Microsoft.Percentage\", \"Percentage\"}\n+\n+\n+def _rollout_percentage(payload: dict[str, Any]) -\u003e float | None:\n+    conditions = payload.get(\"conditions\") or {}\n+    for item in conditions.get(\"client_filters\") or []:\n+        if item.get(\"name\") not in _PERCENTAGE_FILTER_NAMES:\n+            continue\n+        parameters = item.get(\"parameters\") or {}\n+        raw_value = next(\n+            (value for key, value in parameters.items() if key.lower() == \"value\"),\n+            None,\n+        )\n+        if raw_value is None:\n+            raise ValueError(\"Percentage feature filter is missing its Value parameter\")\n+        try:\n+            percentage = float(raw_value)\n+        except (TypeError, ValueError) as exc:\n+            raise ValueError(\"Percentage feature filter Value must be numeric\") from exc\n+        if not 0 \u003c= percentage \u003c= 100:\n+            raise ValueError(\"Percentage feature filter Value must be between 0 and 100\")\n+        return percentage\n+    return None\n+\n+\n+def _parse_flag(flag_name: str, raw_value: str | None, user_id: str | None) -\u003e bool:\n+    if raw_value is None:\n+        raise ValueError(f\"Feature flag {flag_name!r} has no JSON value\")\n+    try:\n+        payload = json.loads(raw_value)\n+    except json.JSONDecodeError as exc:\n+        raise ValueError(f\"Feature flag {flag_name!r} contains invalid JSON\") from exc\n+    if not isinstance(payload, dict):\n+        raise ValueError(f\"Feature flag {flag_name!r} must contain a JSON object\")\n+    if payload.get(\"enabled\") is not True:\n+        return False\n+\n+    percentage = _rollout_percentage(payload)\n+    if percentage is None:\n+        return True\n+    if user_id is None:\n+        raise ValueError(\"A user_id is required for percentage-based feature flags\")\n+\n+    flag_id = str(payload.get(\"id\") or flag_name)\n+    digest = hashlib.sha256(f\"{flag_id}:{user_id}\".encode(\"utf-8\")).digest()\n+    bucket = int.from_bytes(digest[:8], \"big\") * 100 / 2**64\n+    return bucket \u003c percentage\n+\n+\n+class FeatureFlagEvaluator:\n+    \"\"\"Evaluate feature flag JSON obtained through a configuration service.\"\"\"\n+\n+    def __init__(self, configuration: ConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    def is_enabled(\n+        self,\n+        flag_name: str,\n+        *,\n+        user_id: str | None = None,\n+        label: str | None = None,\n+    ) -\u003e bool:\n+        raw_value = self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_name}\", label\n+        )\n+        return _parse_flag(flag_name, raw_value, user_id)\n+\n+\n+class AsyncFeatureFlagEvaluator:\n+    \"\"\"Asynchronously evaluate feature flag JSON.\"\"\"\n+\n+    def __init__(self, configuration: AsyncConfigurationService) -\u003e None:\n+        self._configuration = configuration\n+\n+    async def is_enabled(\n+        self,\n+        flag_name: str,\n+        *,\n+        user_id: str | None = None,\n+        label: str | None = None,\n+    ) -\u003e bool:\n+        raw_value = await self._configuration.get_setting(\n+            f\"{FEATURE_FLAG_PREFIX}{flag_name}\", label\n+        )\n+        return _parse_flag(flag_name, raw_value, user_id)\n+\n*** Add File: appconfig_manager/watcher.py\n+\"\"\"Sentinel-based configuration watchers.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import inspect\n+from collections.abc import Awaitable, Callable, Sequence\n+from threading import Event, Thread\n+from typing import Any\n+\n+from .configuration import AsyncConfigurationService, ConfigurationService\n+\n+\n+class ConfigurationWatcher:\n+    \"\"\"Poll sentinel keys in a background thread and refresh on changes.\"\"\"\n+\n+    def __init__(\n+        self,\n+        configuration: ConfigurationService,\n+        sentinel_keys: Sequence[str],\n+        polling_interval: float = 30.0,\n+        on_refresh: Callable[[], None] | None = None,\n+        label: str | None = None,\n+    ) -\u003e None:\n+        if not sentinel_keys:\n+            raise ValueError(\"At least one sentinel key is required\")\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        self._configuration = configuration\n+        self._sentinel_keys = tuple(sentinel_keys)\n+        self._polling_interval = polling_interval\n+        self._on_refresh = on_refresh\n+        self._label = label\n+        self._stop_event = Event()\n+        self._thread: Thread | None = None\n+\n+    def start(self) -\u003e None:\n+        \"\"\"Prime sentinels and start polling in a daemon thread.\"\"\"\n+        if self._thread is not None and self._thread.is_alive():\n+            raise RuntimeError(\"Configuration watcher is already running\")\n+        for key in self._sentinel_keys:\n+            self._configuration.get_setting(key, self._label)\n+        self._stop_event.clear()\n+        self._thread = Thread(target=self.run, name=\"appconfig-watcher\", daemon=True)\n+        self._thread.start()\n+\n+    def run(self) -\u003e None:\n+        \"\"\"Poll until stopped.\"\"\"\n+        while not self._stop_event.wait(self._polling_interval):\n+            changed = any(\n+                self._configuration.check_for_update(key, self._label)\n+                for key in self._sentinel_keys\n+            )\n+            if changed:\n+                self._configuration.refresh_all()\n+                if self._on_refresh is not None:\n+                    self._on_refresh()\n+\n+    def stop(self, timeout: float | None = None) -\u003e None:\n+        \"\"\"Stop polling and wait for the background thread.\"\"\"\n+        self._stop_event.set()\n+        if self._thread is not None:\n+            self._thread.join(timeout)\n+            if self._thread.is_alive():\n+                raise TimeoutError(\"Configuration watcher did not stop in time\")\n+            self._thread = None\n+\n+    def __enter__(self) -\u003e ConfigurationWatcher:\n+        self.start()\n+        return self\n+\n+    def __exit__(self, *args: Any) -\u003e None:\n+        self.stop()\n+\n+\n+class AsyncConfigurationWatcher:\n+    \"\"\"Poll sentinel keys in an asyncio task and refresh on changes.\"\"\"\n+\n+    def __init__(\n+        self,\n+        configuration: AsyncConfigurationService,\n+        sentinel_keys: Sequence[str],\n+        polling_interval: float = 30.0,\n+        on_refresh: Callable[[], None | Awaitable[None]] | None = None,\n+        label: str | None = None,\n+    ) -\u003e None:\n+        if not sentinel_keys:\n+            raise ValueError(\"At least one sentinel key is required\")\n+        if polling_interval \u003c= 0:\n+            raise ValueError(\"polling_interval must be greater than zero\")\n+        self._configuration = configuration\n+        self._sentinel_keys = tuple(sentinel_keys)\n+        self._polling_interval = polling_interval\n+        self._on_refresh = on_refresh\n+        self._label = label\n+        self._stop_event = asyncio.Event()\n+        self._task: asyncio.Task[None] | None = None\n+\n+    async def start(self) -\u003e None:\n+        \"\"\"Prime sentinels and start polling in an asyncio task.\"\"\"\n+        if self._task is not None and not self._task.done():\n+            raise RuntimeError(\"Configuration watcher is already running\")\n+        for key in self._sentinel_keys:\n+            await self._configuration.get_setting(key, self._label)\n+        self._stop_event.clear()\n+        self._task = asyncio.create_task(self.run(), name=\"appconfig-watcher\")\n+\n+    async def run(self) -\u003e None:\n+        \"\"\"Poll until stopped.\"\"\"\n+        while True:\n+            try:\n+                await asyncio.wait_for(\n+                    self._stop_event.wait(), timeout=self._polling_interval\n+                )\n+                return\n+            except TimeoutError:\n+                pass\n+\n+            changed = False\n+            for key in self._sentinel_keys:\n+                if await self._configuration.check_for_update(key, self._label):\n+                    changed = True\n+            if changed:\n+                await self._configuration.refresh_all()\n+                if self._on_refresh is not None:\n+                    result = self._on_refresh()\n+                    if inspect.isawaitable(result):\n+                        await result\n+\n+    async def stop(self) -\u003e None:\n+        \"\"\"Stop polling and wait for the asyncio task.\"\"\"\n+        self._stop_event.set()\n+        if self._task is not None:\n+            await self._task\n+            self._task = None\n+\n+    async def __aenter__(self) -\u003e AsyncConfigurationWatcher:\n+        await self.start()\n+        return self\n+\n+    async def __aexit__(self, *args: Any) -\u003e None:\n+        await self.stop()\n+\n*** Add File: main.py\n+\"\"\"Run synchronous and asynchronous Azure App Configuration demos.\"\"\"\n+\n+from __future__ import annotations\n+\n+import asyncio\n+import os\n+import time\n+\n+from azure.appconfiguration import AzureAppConfigurationClient\n+from azure.appconfiguration.aio import AzureAppConfigurationClient as AsyncClient\n+from azure.identity import DefaultAzureCredential\n+from azure.identity.aio import DefaultAzureCredential as AsyncDefaultAzureCredential\n+\n+from appconfig_manager import (\n+    AsyncConfigurationService,\n+    AsyncConfigurationWatcher,\n+    AsyncFeatureFlagEvaluator,\n+    ConfigurationService,\n+    ConfigurationWatcher,\n+    FeatureFlagEvaluator,\n+)\n+\n+SAMPLE_USERS = (\"alice\", \"bob\", \"charlie\", \"diana\")\n+\n+\n+def _endpoint() -\u003e str:\n+    try:\n+        return os.environ[\"AZURE_APPCONFIG_ENDPOINT\"]\n+    except KeyError as exc:\n+        raise RuntimeError(\n+            \"Set AZURE_APPCONFIG_ENDPOINT to an App Configuration endpoint\"\n+        ) from exc\n+\n+\n+def run_sync_demo(endpoint: str, watch_seconds: float) -\u003e None:\n+    print(\"\\n--- Synchronous demo ---\")\n+    with DefaultAzureCredential() as credential:\n+        with AzureAppConfigurationClient(endpoint, credential) as client:\n+            configuration = ConfigurationService(client)\n+            flags = FeatureFlagEvaluator(configuration)\n+\n+            print(\n+                \"production API URL:\",\n+                configuration.get_setting_with_label(\"Demo:ApiUrl\", \"production\"),\n+            )\n+            print(\n+                \"staging settings:\",\n+                configuration.list_settings(\"Demo:\", label=\"staging\"),\n+            )\n+            for user_id in SAMPLE_USERS:\n+                enabled = flags.is_enabled(\n+                    \"NewCheckout\", user_id=user_id, label=\"production\"\n+                )\n+                print(f\"NewCheckout for {user_id}: {enabled}\")\n+\n+            watcher = ConfigurationWatcher(\n+                configuration,\n+                sentinel_keys=[\"Demo:Sentinel\"],\n+                polling_interval=5,\n+                label=\"production\",\n+                on_refresh=lambda: print(\"Sync configuration cache refreshed\"),\n+            )\n+            watcher.start()\n+            print(f\"Watching for sentinel changes for {watch_seconds:g} seconds...\")\n+            time.sleep(watch_seconds)\n+            watcher.stop()\n+\n+\n+async def run_async_demo(endpoint: str, watch_seconds: float) -\u003e None:\n+    print(\"\\n--- Asynchronous demo ---\")\n+    async with AsyncDefaultAzureCredential() as credential:\n+        async with AsyncClient(endpoint, credential) as client:\n+            configuration = AsyncConfigurationService(client)\n+            flags = AsyncFeatureFlagEvaluator(configuration)\n+\n+            print(\n+                \"production API URL:\",\n+                await configuration.get_setting_with_label(\n+                    \"Demo:ApiUrl\", \"production\"\n+                ),\n+            )\n+            print(\n+                \"staging settings:\",\n+                await configuration.list_settings(\"Demo:\", label=\"staging\"),\n+            )\n+            for user_id in SAMPLE_USERS:\n+                enabled = await flags.is_enabled(\n+                    \"NewCheckout\", user_id=user_id, label=\"production\"\n+                )\n+                print(f\"NewCheckout for {user_id}: {enabled}\")\n+\n+            watcher = AsyncConfigurationWatcher(\n+                configuration,\n+                sentinel_keys=[\"Demo:Sentinel\"],\n+                polling_interval=5,\n+                label=\"production\",\n+                on_refresh=lambda: print(\"Async configuration cache refreshed\"),\n+            )\n+            await watcher.start()\n+            print(f\"Watching for sentinel changes for {watch_seconds:g} seconds...\")\n+            await asyncio.sleep(watch_seconds)\n+            await watcher.stop()\n+\n+\n+def main() -\u003e None:\n+    endpoint = _endpoint()\n+    watch_seconds = float(os.getenv(\"DEMO_WATCH_SECONDS\", \"15\"))\n+    if watch_seconds \u003c 0:\n+        raise ValueError(\"DEMO_WATCH_SECONDS cannot be negative\")\n+    run_sync_demo(endpoint, watch_seconds)\n+    asyncio.run(run_async_demo(endpoint, watch_seconds))\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n+\n*** Add File: README.md\n+# Azure App Configuration demo\n+\n+This project provides cached synchronous and asynchronous configuration services,\n+deterministic percentage-based feature flags, and sentinel-driven configuration\n+watchers.\n+\n+## Setup\n+\n+Create the demo entries in Azure App Configuration:\n+\n+- `Demo:ApiUrl`, with `production` and `staging` labels\n+- One or more `Demo:` settings with a `staging` label\n+- `Demo:Sentinel` with a `production` label\n+- `.appconfig.featureflag/NewCheckout` with a `production` label and a\n+  `Microsoft.Percentage` client filter\n+\n+Then install and run:\n+\n+```powershell\n+python -m venv .venv\n+.\\.venv\\Scripts\\Activate.ps1\n+pip install -r requirements.txt\n+$env:AZURE_APPCONFIG_ENDPOINT = \"https://\u003cstore\u003e.azconfig.io\"\n+$env:DEMO_WATCH_SECONDS = \"30\"\n+python main.py\n+```\n+\n+`DefaultAzureCredential` uses local developer credentials or managed identity.\n+The demo runs the synchronous flow first and then the asynchronous flow.\n*** End Patch\n"
```

### 🔧 4. powershell

**Input:**

```json
{"command":"python -m compileall -q appconfig_manager main.py; if ($LASTEXITCODE -eq 0) { Write-Output 'compileall passed' }","description":"Compile all Python source files"}
```

### 🔧 5. powershell

**Input:**

```json
{"command":"python -c \"import azure.appconfiguration, azure.identity; print('Azure SDK available')\"","description":"Check Azure SDK availability"}
```

### 🔧 6. rg

**Input:**

```json
{"-n":true,"glob":"*.py","head_limit":100,"output_mode":"content","paths":["appconfig_manager","main.py"],"pattern":"^.{89,}$"}
```

### 🔧 7. powershell

**Input:**

```json
{"command":"python -m pip install --quiet -r requirements.txt","description":"Install declared Azure SDK dependencies","initial_wait":180,"mode":"sync"}
```

### 🔧 8. powershell

**Input:**

```json
{"command":"python -c \"import inspect; from azure.appconfiguration import AzureAppConfigurationClient; from azure.appconfiguration.aio import AzureAppConfigurationClient as A; from azure.core import MatchConditions; print(inspect.signature(AzureAppConfigurationClient.get_configuration_setting)); print(inspect.signature(AzureAppConfigurationClient.list_configuration_settings)); print(inspect.signature(A.get_configuration_setting)); print(list(MatchConditions))\"","description":"Inspect Azure SDK method signatures"}
```

### 🔧 9. powershell

**Input:**

```json
{"command":"python -c \"import appconfig_manager; import main; print('project imports passed')\"","description":"Import project modules"}
```

### 🔧 10. powershell

**Input:**

```json
{"command":"python -m pip show azure-appconfiguration azure-identity | Select-String '^(Name|Version):'","description":"Show installed Azure SDK versions"}
```

### 🔧 11. powershell

**Input:**

```json
{"command":"python -c \"from azure.appconfiguration import AzureAppConfigurationClient; print(AzureAppConfigurationClient.list_configuration_settings.__doc__)\"","description":"Inspect list settings SDK documentation"}
```

### 🔧 12. powershell

**Input:**

```json
{"command":"python -c \"from azure.core import MatchConditions; print(MatchConditions.IfModified, MatchConditions.IfNotModified)\"","description":"Confirm match condition values"}
```

### 🔧 13. powershell

**Input:**

```json
{"command":"python -c \"from azure.appconfiguration import ConfigurationSetting; import inspect; print(inspect.signature(ConfigurationSetting))\"","description":"Inspect SDK setting model"}
```

### 🔧 14. powershell

**Input:**

```json
{"command":"python -c \"import inspect; from azure.appconfiguration import AzureAppConfigurationClient; print(inspect.getsource(AzureAppConfigurationClient.list_configuration_settings))\"","description":"Verify list settings keyword support"}
```

### 🔧 15. apply_patch

**Input:**

```json
"*** Begin Patch\n*** Update File: appconfig_manager/configuration.py\n@@\n             self._prefix_queries.add(query)\n             members = self._prefix_members.get(query)\n             if members is not None and not refresh:\n-                return {key: self._cache[(key, item_label)].value for key, item_label in members}\n+                return {\n+                    key: self._cache[(key, item_label)].value\n+                    for key, item_label in members\n+                }\n@@\n \n             self._prefix_members[query] = new_members\n-            return {key: self._cache[(key, item_label)].value for key, item_label in new_members}\n+            return {\n+                key: self._cache[(key, item_label)].value\n+                for key, item_label in new_members\n+            }\n@@\n-        \"\"\"Conditionally poll a setting and report whether it changed or was deleted.\"\"\"\n+        \"\"\"Report whether a conditional poll finds a change or deletion.\"\"\"\n@@\n-        \"\"\"Conditionally poll a setting and report whether it changed or was deleted.\"\"\"\n+        \"\"\"Report whether a conditional poll finds a change or deletion.\"\"\"\n@@\n         self._prefix_queries.add(query)\n         members = self._prefix_members.get(query)\n         if members is not None and not refresh:\n-            return {key: self._cache[(key, item_label)].value for key, item_label in members}\n+            return {\n+                key: self._cache[(key, item_label)].value\n+                for key, item_label in members\n+            }\n@@\n \n         self._prefix_members[query] = new_members\n-        return {key: self._cache[(key, item_label)].value for key, item_label in new_members}\n+        return {\n+            key: self._cache[(key, item_label)].value\n+            for key, item_label in new_members\n+        }\n*** Update File: appconfig_manager/feature_flags.py\n@@\n         )\n         if raw_value is None:\n-            raise ValueError(\"Percentage feature filter is missing its Value parameter\")\n+            raise ValueError(\n+                \"Percentage feature filter is missing its Value parameter\"\n+            )\n@@\n         try:\n             percentage = float(raw_value)\n         except (TypeError, ValueError) as exc:\n-            raise ValueError(\"Percentage feature filter Value must be numeric\") from exc\n+            raise ValueError(\n+                \"Percentage feature filter Value must be numeric\"\n+            ) from exc\n         if not 0 \u003c= percentage \u003c= 100:\n-            raise ValueError(\"Percentage feature filter Value must be between 0 and 100\")\n+            raise ValueError(\n+                \"Percentage feature filter Value must be between 0 and 100\"\n+            )\n*** Add File: tests/__init__.py\n+\n*** Add File: tests/test_appconfig_manager.py\n+from __future__ import annotations\n+\n+import asyncio\n+import json\n+import unittest\n+from types import SimpleNamespace\n+from typing import Any\n+\n+from azure.core.exceptions import ResourceNotFoundError, ResourceNotModifiedError\n+\n+from appconfig_manager.configuration import (\n+    AsyncConfigurationService,\n+    ConfigurationService,\n+)\n+from appconfig_manager.feature_flags import (\n+    AsyncFeatureFlagEvaluator,\n+    FeatureFlagEvaluator,\n+)\n+\n+\n+def _setting(\n+    key: str, value: str | None, etag: str, label: str | None = None\n+) -\u003e SimpleNamespace:\n+    return SimpleNamespace(key=key, value=value, etag=etag, label=label)\n+\n+\n+class FakeClient:\n+    def __init__(self, settings: list[SimpleNamespace]) -\u003e None:\n+        self.settings = {(item.key, item.label): item for item in settings}\n+        self.get_calls = 0\n+        self.list_calls: list[list[str] | None] = []\n+\n+    def get_configuration_setting(self, **kwargs: Any) -\u003e SimpleNamespace:\n+        self.get_calls += 1\n+        item = self.settings.get((kwargs[\"key\"], kwargs[\"label\"]))\n+        if item is None:\n+            raise ResourceNotFoundError(\"missing\")\n+        if kwargs.get(\"etag\") == item.etag:\n+            raise ResourceNotModifiedError(\"unchanged\")\n+        return item\n+\n+    def list_configuration_settings(self, **kwargs: Any) -\u003e list[SimpleNamespace]:\n+        fields = kwargs.get(\"fields\")\n+        self.list_calls.append(fields)\n+        prefix = kwargs[\"key_filter\"][:-1]\n+        label = None if kwargs[\"label_filter\"] == \"\\0\" else kwargs[\"label_filter\"]\n+        matches = [\n+            item\n+            for (key, item_label), item in self.settings.items()\n+            if key.startswith(prefix) and item_label == label\n+        ]\n+        if fields:\n+            return [\n+                _setting(item.key, None, item.etag, item.label) for item in matches\n+            ]\n+        return matches\n+\n+\n+class AsyncItems:\n+    def __init__(self, items: list[SimpleNamespace]) -\u003e None:\n+        self._items = items\n+\n+    def __aiter__(self) -\u003e AsyncItems:\n+        self._iterator = iter(self._items)\n+        return self\n+\n+    async def __anext__(self) -\u003e SimpleNamespace:\n+        try:\n+            return next(self._iterator)\n+        except StopIteration as exc:\n+            raise StopAsyncIteration from exc\n+\n+\n+class FakeAsyncClient(FakeClient):\n+    async def get_configuration_setting(self, **kwargs: Any) -\u003e SimpleNamespace:\n+        return super().get_configuration_setting(**kwargs)\n+\n+    def list_configuration_settings(self, **kwargs: Any) -\u003e AsyncItems:\n+        return AsyncItems(super().list_configuration_settings(**kwargs))\n+\n+\n+def _percentage_flag(value: float) -\u003e str:\n+    return json.dumps(\n+        {\n+            \"id\": \"rollout\",\n+            \"enabled\": True,\n+            \"conditions\": {\n+                \"client_filters\": [\n+                    {\n+                        \"name\": \"Microsoft.Percentage\",\n+                        \"parameters\": {\"Value\": value},\n+                    }\n+                ]\n+            },\n+        }\n+    )\n+\n+\n+class ConfigurationServiceTests(unittest.TestCase):\n+    def test_get_uses_cache_and_conditional_refresh(self) -\u003e None:\n+        client = FakeClient([_setting(\"Api:Url\", \"one\", \"1\")])\n+        service = ConfigurationService(client)\n+\n+        self.assertEqual(service.get_setting(\"Api:Url\"), \"one\")\n+        self.assertEqual(service.get_setting(\"Api:Url\"), \"one\")\n+        self.assertEqual(client.get_calls, 1)\n+        self.assertFalse(service.check_for_update(\"Api:Url\"))\n+        self.assertEqual(client.get_calls, 2)\n+\n+        client.settings[(\"Api:Url\", None)] = _setting(\"Api:Url\", \"two\", \"2\")\n+        self.assertTrue(service.check_for_update(\"Api:Url\"))\n+        self.assertEqual(service.get_setting(\"Api:Url\"), \"two\")\n+\n+    def test_prefix_refresh_fetches_only_changed_values(self) -\u003e None:\n+        client = FakeClient(\n+            [_setting(\"App:A\", \"a\", \"1\"), _setting(\"App:B\", \"b\", \"1\")]\n+        )\n+        service = ConfigurationService(client)\n+\n+        self.assertEqual(service.list_settings(\"App:\"), {\"App:A\": \"a\", \"App:B\": \"b\"})\n+        client.settings[(\"App:B\", None)] = _setting(\"App:B\", \"new\", \"2\")\n+        self.assertEqual(\n+            service.list_settings(\"App:\", refresh=True),\n+            {\"App:A\": \"a\", \"App:B\": \"new\"},\n+        )\n+        self.assertEqual(client.list_calls, [None, [\"key\", \"label\", \"etag\"]])\n+        self.assertEqual(client.get_calls, 1)\n+\n+    def test_percentage_rollout_is_deterministic(self) -\u003e None:\n+        client = FakeClient(\n+            [_setting(\".appconfig.featureflag/Test\", _percentage_flag(30), \"1\")]\n+        )\n+        evaluator = FeatureFlagEvaluator(ConfigurationService(client))\n+\n+        first = evaluator.is_enabled(\"Test\", user_id=\"alice\")\n+        self.assertEqual(first, evaluator.is_enabled(\"Test\", user_id=\"alice\"))\n+        outcomes = {\n+            evaluator.is_enabled(\"Test\", user_id=f\"user-{number}\")\n+            for number in range(100)\n+        }\n+        self.assertEqual(outcomes, {False, True})\n+\n+\n+class AsyncConfigurationServiceTests(unittest.IsolatedAsyncioTestCase):\n+    async def test_async_cache_refresh_and_rollout(self) -\u003e None:\n+        client = FakeAsyncClient(\n+            [\n+                _setting(\"Api:Url\", \"one\", \"1\", \"production\"),\n+                _setting(\n+                    \".appconfig.featureflag/Test\",\n+                    _percentage_flag(100),\n+                    \"1\",\n+                    \"production\",\n+                ),\n+            ]\n+        )\n+        service = AsyncConfigurationService(client)\n+        evaluator = AsyncFeatureFlagEvaluator(service)\n+\n+        self.assertEqual(\n+            await service.get_setting_with_label(\"Api:Url\", \"production\"), \"one\"\n+        )\n+        self.assertFalse(await service.check_for_update(\"Api:Url\", \"production\"))\n+        self.assertTrue(\n+            await evaluator.is_enabled(\n+                \"Test\", user_id=\"alice\", label=\"production\"\n+            )\n+        )\n+\n+        client.settings[(\"Api:Url\", \"production\")] = _setting(\n+            \"Api:Url\", \"two\", \"2\", \"production\"\n+        )\n+        self.assertTrue(await service.check_for_update(\"Api:Url\", \"production\"))\n+        self.assertEqual(\n+            await service.get_setting(\"Api:Url\", \"production\"), \"two\"\n+        )\n+\n+\n+if __name__ == \"__main__\":\n+    unittest.main()\n*** End Patch\n"
```

### 🔧 16. powershell

**Input:**

```json
{"command":"python -m unittest discover -v","description":"Run project unit tests","initial_wait":120,"mode":"sync"}
```

### 🔧 17. powershell

**Input:**

```json
{"command":"python -m compileall -q appconfig_manager tests main.py; if ($LASTEXITCODE -eq 0) { Write-Output 'compileall passed' }","description":"Compile project and tests"}
```

### 🔧 18. rg

**Input:**

```json
{"-n":true,"glob":"*.py","head_limit":100,"output_mode":"content","paths":["appconfig_manager","tests","main.py"],"pattern":"^.{89,}$"}
```

### 🔧 19. glob

**Input:**

```json
{"pattern":"**/__pycache__"}
```

### 🔧 20. powershell

**Input:**

```json
{"command":"Get-ChildItem -Recurse -Directory -Filter __pycache__ | Select-Object -ExpandProperty FullName","description":"Find Python cache directories"}
```

### 🔧 21. powershell

**Input:**

```json
{"command":"Remove-Item -Recurse -Force '.\\__pycache__'; Remove-Item -Recurse -Force '.\\appconfig_manager\\__pycache__'; Remove-Item -Recurse -Force '.\\tests\\__pycache__'","description":"Remove generated Python cache directories"}
```

## Generated Files

- `README.md`
- `appconfig_manager\__init__.py`
- `appconfig_manager\configuration.py`
- `appconfig_manager\feature_flags.py`
- `appconfig_manager\watcher.py`
- `main.py`
- `requirements.txt`
- `tests\__init__.py`
- `tests\test_appconfig_manager.py`

## Copilot Response

I’ll scaffold the project, implement the sync/async services and watcher, then run local syntax and unit-level checks without contacting Azure.

The design will use ETag-based conditional reads for individual settings and metadata-only prefix scans on refresh, so unchanged values remain cached. Sentinel changes invalidate and repopulate every previously requested key and prefix.

The source compiles, but the Azure SDK is not installed in this workspace. I’m installing the declared dependencies so the real SDK imports and request signatures can be validated.

The installed SDK confirms support for `fields=["key", "label", "etag"]` and `MatchConditions.IfModified`. I’m adding focused fake-client tests to exercise caching, ETag polling, prefix refreshes, and identical sync/async rollout behavior without any Azure connection.

Created the Python project with:

- Sync and async Azure App Configuration services with ETag caching and labeled/prefix lookups
- Deterministic percentage-based feature flag evaluation
- Sync and async sentinel watchers that refresh cached configuration
- `main.py` demonstrating both implementations with `DefaultAzureCredential`
- Azure SDK dependencies in `requirements.txt`
- Setup instructions in `README.md`
- Offline unit tests using fake Azure clients

Run with:

```powershell
$env:AZURE_APPCONFIG_ENDPOINT = "https://<store>.azconfig.io"
python main.py
```

## Reviewer Notes (LLM-as-Judge)

### Summary

Azure SDK exception handling is only partial and does not fully satisfy the criterion.

### Strengths

- Conditional requests correctly treat ResourceNotModifiedError as an unchanged cached value.
- Sentinel deletion and refresh-time missing settings are explicitly handled via ResourceNotFoundError.

### Issues

- No handling for general HttpResponseError failures such as authentication, authorization, throttling, or service errors.
- Watcher polling has no exception boundary, so an Azure request failure can silently stop background monitoring.
- ResourceNotFoundError from a normal get_setting call is exposed directly rather than being deliberately handled or translated.

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
hyoka run --prompt-id app-configuration-dp-python-feature-flags --config python-azure-tools/baseline --pairwise-variant baseline
```

---

[← Back to Summary](../../../../../../summary.md)
