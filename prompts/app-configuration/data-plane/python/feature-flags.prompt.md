---
id: app-configuration-dp-python-feature-flags
properties:
  service: app-configuration
  plane: data-plane
  language: python
  category: crud
  difficulty: intermediate
  description: 'Can an agent generate a feature flag and configuration management system using Azure App Configuration with
    label-based settings, conditional reads with ETags, percentage-based rollout, and sentinel key watching?

    '
  sdk_package: azure-appconfiguration
  doc_url: https://learn.microsoft.com/en-us/python/api/overview/azure/appconfiguration-readme
  created: '2026-04-10'
  author: copilot
tags:
- app-configuration
- feature-flags
- etag
- conditional-reads
- labels
- sentinel-key
- percentage-rollout
- async
---

# Feature Flag Management: Azure App Configuration (Python)

## Prompt

Create a Python project that implements a feature flag and configuration management system backed by Azure App Configuration.

The project needs:

- A **configuration service module** (both sync and async versions) that retrieves settings from App Configuration. It should support fetching a single setting by key, fetching a setting with a specific label (to distinguish between environments like "production" vs "staging"), and listing all settings that match a key prefix (returned as a dictionary). It should also avoid re-downloading values that haven't changed since the last read — minimize unnecessary network traffic when polling for config changes.

- A **feature flag evaluator module** that reads feature flags from App Configuration. Feature flags in App Configuration use the `.appconfig.featureflag/` key prefix and store their state as a JSON payload. The evaluator should be able to check if a flag is enabled, and also support percentage-based rollout — if a flag is configured for a percentage rollout (e.g., 30% of users), the evaluator should deterministically decide whether a given user ID falls within the rollout percentage using a consistent hash, so the same user always gets the same result.

- A **configuration watcher module** that periodically polls for configuration changes. It should accept a list of "sentinel" keys to watch and a polling interval. When a sentinel key's value changes, the watcher should trigger a full refresh of all cached configuration. This is the recommended pattern for coordinating config updates in App Configuration.

- A **main script** that demos both implementations: connecting to App Configuration (endpoint from environment variable, authenticated with `DefaultAzureCredential`), reading some config values with labels, evaluating feature flags for a few sample user IDs with percentage rollout, and starting the config watcher to detect a change. Run the full demo with the sync implementation first, then repeat with the async implementation.

Include a `requirements.txt` with the necessary Azure SDK dependencies.

## Evaluation Criteria

### Scenario-Specific Patterns
- Retrieves settings with a specific label parameter using `label_filter`
- Lists settings filtered by key prefix using `key_filter`
- Implements conditional reads using `match_condition` with `MatchConditions.IfModified` and the setting's `etag`
- Handles 304 Not Modified (setting unchanged since last read)
- Uses `.appconfig.featureflag/` prefix for feature flag keys
- Parses the JSON payload in feature flag setting values
- Implements deterministic percentage rollout (consistent hash via `hashlib`, not `random`)
- Implements sentinel key watching with configurable polling interval
- Detects sentinel value change via ETag or value comparison and triggers full refresh
- Async version uses `azure.appconfiguration.aio.AzureAppConfigurationClient`

## Context

This goes beyond basic key-value reads (covered by `config-values.prompt.md`) to test three
advanced App Configuration patterns: label-based environment separation, conditional reads
with ETags for efficient polling, and the sentinel key watching pattern for coordinated config
refresh. The feature flag evaluator tests whether the agent understands App Configuration's
`.appconfig.featureflag/` key prefix convention and can implement deterministic percentage
rollout without using non-deterministic random values.
