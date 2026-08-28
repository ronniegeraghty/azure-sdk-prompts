# Evaluation Report: app-configuration-dp-python-crud

**Config:** python-azure-skills/baseline | **Result:** ❌ FAILED | **Duration:** 119.9s

## Overview

| Field | Value |
|-------|-------|
| Prompt ID | `app-configuration-dp-python-crud` |
| Config | python-azure-skills/baseline |
| Result | ❌ FAILED |
| Score | 8/13 |
| Duration | 119.9s |
| Timestamp | 2026-08-28T23:27:12Z |
| Files Generated | 2 |
| Event Count | 1411 |

## Phase Timing

| Phase | Duration |
|-------|----------|
| Generation | 40.4s |
| Review | 78.7s |
| **Total** | **119.9s** |

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
| Token Usage | in=102625 out=1350 |
| Turn Count | 5 |

## Tool Calls

### ✅ 1. glob — 8974ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511","pattern":"**/*"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 2. glob — 8603ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511","pattern":"**/requirements*.txt"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 3. glob — 8246ms

**Input:**

```json
{"paths":"C:\\Users\\weidxu\\AppData\\Local\\Temp\\hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511","pattern":"**/*.py"}
```

**Output:**

```
No files matched the pattern.
```

### ✅ 4. apply_patch — 764ms

**Input:**

```json
"*** Begin Patch\n*** Add File: app_configuration.py\n+import os\n+\n+from azure.appconfiguration import (\n+    AzureAppConfigurationClient,\n+    FeatureFlagConfigurationSetting,\n+)\n+from azure.core.exceptions import HttpResponseError\n+\n+\n+CONNECTION_STRING_ENV = \"AZURE_APP_CONFIG_CONNECTION_STRING\"\n+SETTING_KEY = \"app:Settings:FontSize\"\n+\n+\n+def main() -\u003e None:\n+    connection_string = os.environ.get(CONNECTION_STRING_ENV)\n+    if not connection_string:\n+        raise RuntimeError(\n+            f\"Set the {CONNECTION_STRING_ENV} environment variable before running.\"\n+        )\n+\n+    client = AzureAppConfigurationClient.from_connection_string(connection_string)\n+\n+    try:\n+        client.set_configuration_setting(key=SETTING_KEY, value=\"24\")\n+        client.set_configuration_setting(\n+            key=SETTING_KEY,\n+            label=\"Production\",\n+            value=\"24\",\n+        )\n+\n+        setting = client.get_configuration_setting(key=SETTING_KEY)\n+        print(setting.value)\n+\n+        for matching_setting in client.list_configuration_settings(\n+            key_filter=\"app:Settings:*\"\n+        ):\n+            print(\n+                f\"{matching_setting.key}={matching_setting.value} \"\n+                f\"(label={matching_setting.label!r})\"\n+            )\n+\n+        client.set_configuration_setting(\n+            FeatureFlagConfigurationSetting(\n+                feature_id=\"BetaFeature\",\n+                enabled=True,\n+            )\n+        )\n+\n+        client.delete_configuration_setting(key=SETTING_KEY)\n+    except HttpResponseError as error:\n+        print(\n+            f\"Azure App Configuration request failed \"\n+            f\"(status={error.status_code}): {error.message}\"\n+        )\n+        raise\n+\n+\n+if __name__ == \"__main__\":\n+    main()\n*** Add File: requirements.txt\n+azure-appconfiguration\u003e=1.7.0,\u003c2.0.0\n*** End Patch\n"
```

**Output:**

```

diff --git a/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511/app_configuration.py b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511/app_configuration.py
create file mode 100644
index 0000000..0000000
--- a/dev/null
+++ b/C:/Users/weidxu/AppData/Local/Temp/hyoka-app-configuration-dp-python-crud-python-azure-skills-baseline-3045344511/app_configuration.py
@@ -1,0 +1,60 @@
+import os
+
+from azure.appconfiguration import (
+    AzureAppConfigurationClient,
+    FeatureFlagConfigurationSetting,
+)
+from azure.core.exceptions import HttpResponseError
+
+
+CONNECTION_STRING_ENV = "AZURE_APP_CONFIG_CONNECTION_STRING"
+SETTING_KEY = "app:Settings:FontSize"
+
+
+def main() -> None:
+    connection_string = os.environ.get(CONNECTION_STRING_ENV)
+    if not connection_string:
+        raise RuntimeError(
+            f"Set the {CONNECTION_STRING_ENV} environment variable before running."
+        )
+
+    client = AzureAppConfigurationClient.from_connection_string(connection_string)
+
+    try:
+        client.set_configuration_setting(key=SETTING_KEY, value="24")
+        client.set_configuration_setting(
+            key=SETTING_KEY,
+            label="Production",
+            value="24",
+        )
+
+        setting = client.get_configuration_setting(key=SETTING_KEY)
+        print(setting.value)
+
+        for matching_setting in client.list_configuration_settings(
+            key_filter="app:Settings:*"
+        ):
+            print(
+                f"{matching_setting.key}={matching_setting.value} "
+                f"(label={matching_setting.label!r})"
+            )
+
+        client.set_configuration_setting(
+            FeatureFlagConfigurationSetting(
+                feature_id="BetaFeature",
+                enabled=True,
+            )
+        )
+
+        client.delete_configuration_setting(key=SETTING_KEY)
+    except Http
... (truncated)
```

### ✅ 5. powershell — 1604ms

**Input:**

```json
{"command":"python -m py_compile .\\app_configuration.py","description":"Check Python script syntax","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 0 completed with exit code 0>
```

### ✅ 6. powershell — 1299ms

**Input:**

```json
{"command":"if (Test-Path .\\__pycache__) { Remove-Item -Recurse -Force .\\__pycache__ }","description":"Remove Python cache directory","initial_wait":30,"mode":"sync"}
```

**Output:**

```

<shellId: 1 completed with exit code 0>
```

## Generated Files

- `app_configuration.py`
- `requirements.txt`

## Copilot Response

I’ll inspect the workspace, then add a locally runnable example and its required dependency manifest.

The files are in place; I’m checking syntax without connecting to or modifying any Azure resource.

Created:

- `app_configuration.py` — implements setting, labeled setting, retrieval, filtered listing, enabled `BetaFeature` feature flag, deletion, and `HttpResponseError` handling.
- `requirements.txt` — includes `azure-appconfiguration>=1.7.0,<2.0.0`.

The script reads the connection string from `AZURE_APP_CONFIG_CONNECTION_STRING` and does not contain credentials.

## Reviewer Notes (LLM-as-Judge)

### Summary

1 of 1 criteria passed.

### Strengths

- Uses the correct Azure SDK exception type and handles request failures explicitly while preserving failure propagation.

## Grader Results

- config-values.prompt.md (prompt file):
  - Criteria from prompt file (prompt): Fail (5/6)
      - `azure-appconfiguration` pip package: Pass
      - `set_configuration_setting()` with `ConfigurationSetting` objects: Fail
      - `get_configuration_setting()` by key: Pass
      - `list_configuration_settings()` with key_filter: Pass
      - `FeatureFlagConfigurationSetting` for feature flags: Pass
      - `delete_configuration_setting()` and `HttpResponseError` handling: Pass
- python.yaml (criteria file):
  - Correct Package Imports (prompt): Pass (1/1)
      - ### Attribute-Matched Criteria

**Correct Package Imports**: Imports use the latest azure-sdk-for-python package structure (azure.*), not deprecated packages.: Pass
  - DefaultAzureCredential Usage (prompt): Fail (0/1)
      - ### Attribute-Matched Criteria

**DefaultAzureCredential Usage**: Authentication matches what the prompt asks for. If the prompt explicitly requires a connection string (or other key-based auth), using `from_connection_string()` / connection-string-based clients is correct and should pass. Otherwise, authentication must use DefaultAzureCredential from azure-identity (or another `azure.identity` credential), not connection strings or hardcoded keys. Hardcoded secrets/keys/connection strings in source code always fail — required values should come from environment variables or a secret store.: Fail
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
| `Criteria from prompt file` | prompt_review | 83% | 1.00 | 0.8333 | 21.7% | ❌ |
| `Correct Package Imports` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `DefaultAzureCredential Usage` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Context Manager for Clients` | prompt_review | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Async Client Usage` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `Proper Exception Handling` | prompt_review | 100% | 1.00 | 1.0000 | 26.1% | ✅ |
| `Output Files Exist` | workspace | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| `Tool Usage Verification` | tool | 0% | 1.00 | 0.0000 | 0.0% | ❌ |
| **Final** | | | **Σ 8.00** | **Σ 3.8333** | **47.9%** | |

## Re-run Command

```bash
hyoka run --prompt-id app-configuration-dp-python-crud --config python-azure-skills/baseline --pairwise-variant baseline --monitor-resources
```

---

[← Back to Summary](../../../../../../summary.md)
