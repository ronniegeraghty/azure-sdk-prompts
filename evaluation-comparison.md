# Evaluation Results: Side-by-Side Comparison

> **Configs compared:**
> 1. **Baseline** — `baseline/claude-sonnet-4.5` (run `20260401-100850`)
> 2. **Azure MCP** — `azure-mcp/claude-sonnet-4.5` (run `20260401-115016`)
> 3. **Skills** — `baseline-skills/claude-sonnet-4.5` (run `20260402-134211`) — 26 azure-sdk-java generator skills

---

## 1. Score Comparison Table

| # | Prompt ID | Baseline | Azure MCP | Skills | Δ Skills vs Baseline | Verification (B/M/S) |
|---|-----------|----------|-----------|--------|----------------------|----------------------|
| 1 | `app-configuration-dp-java-crud` | **11/12** (92%) | **11/12** (92%) | **11/12** (92%) | +0.0% | ✅/✅/✅ |
| 2 | `app-configuration-dp-java-feature-flags` | **8/14** (57%) | **14/21** (67%) | **8/14** (57%) | +0.0% | ✅/✅/✅ |
| 3 | `cosmos-db-dp-java-crud` | **10/12** (83%) | **10/12** (83%) | **9/12** (75%) | -8.3% | ✅/✅/✅ |
| 4 | `cosmos-db-dp-java-todo-repository` | **24/32** (75%) | **13/19** (68%) | **13/19** (68%) | -6.6% | ✅/✅/✅ |
| 5 | `event-hubs-dp-java-streaming` | **11/12** (92%) | **11/12** (92%) | **12/13** (92%) | +0.6% | ✅/✅/✅ |
| 6 | `identity-dp-java-credential-chain` | **16/18** (89%) | **30/33** (91%) | **9/12** (75%) | -13.9% | ✅/✅/❌ |
| 7 | `identity-dp-java-default-credential` | **8/10** (80%) | **7/10** (70%) | **9/11** (82%) | +1.8% | ✅/✅/✅ |
| 8 | `identity-dp-java-managed-identity` | **8/11** (73%) | **8/11** (73%) | **9/11** (82%) | +9.1% | ✅/✅/✅ |
| 9 | `identity-dp-java-service-principal` | **9/10** (90%) | **9/10** (90%) | **8/10** (80%) | -10.0% | ✅/✅/✅ |
| 10 | `key-vault-dp-java-crud` | **8/10** (80%) | **9/10** (90%) | **9/10** (90%) | +10.0% | ✅/✅/✅ |
| 11 | `key-vault-dp-java-secret-config` | **18/20** (90%) | **12/15** (80%) | **24/25** (96%) | +6.0% | ✅/❌/✅ |
| 12 | `resource-manager-mp-java-rg-crud` | **10/12** (83%) | **11/12** (92%) | **11/12** (92%) | +8.3% | ✅/✅/✅ |
| 13 | `service-bus-dp-java-crud` | **10/12** (83%) | **11/12** (92%) | **11/12** (92%) | +8.3% | ✅/✅/✅ |
| 14 | `service-bus-dp-java-order-processor` | **11/17** (65%) | **11/17** (65%) | **11/17** (65%) | +0.0% | ❌/❌/❌ |
| 15 | `storage-dp-java-blob-event-notifier` | **8/15** (53%) | **12/15** (80%) | **18/22** (82%) | +28.5% | ❌/✅/✅ |
| 16 | `storage-dp-java-blob-manager` | **9/12** (75%) | **12/15** (80%) | **9/12** (75%) | +0.0% | ✅/❌/✅ |
| 17 | `storage-dp-java-crud` | **11/12** (92%) | **11/12** (92%) | **11/12** (92%) | +0.0% | ✅/✅/✅ |
| 18 | `storage-dp-java-encrypted-uploader` | **22/25** (88%) | **51/53** (96%) | **22/25** (88%) | +0.0% | ❌/❌/❌ |
| 19 | `storage-mp-java-account-mgmt` | **9/13** (69%) | **9/13** (69%) | **10/13** (77%) | +7.7% | ✅/✅/✅ |
|---|-----------|----------|-----------|--------|----------------------|----------------------|
| | **TOTALS** | **221/279** (79.2%) | **262/314** (83.4%) | **224/274** (81.8%) | +2.5% | 16/15/16 of 19 |

> **Note on max_score variation:** Different configs may produce different criteria counts (max_score) for
> the same prompt because the review rubric can expand based on generated code complexity. Percentages
> normalize for fair comparison.

---

## 2. Criteria Deep-Dive: Per-Prompt Pass/Fail Differences

Below shows criteria that **differ** across configs for each prompt.
Criteria passing identically in all configs are omitted for brevity.

### `app-configuration-dp-java-crud`

Scores: Baseline 11/12 | MCP 11/12 | Skills 11/12

*All criteria identical across configs.*

### `app-configuration-dp-java-feature-flags`

Scores: Baseline 8/14 | MCP 14/21 | Skills 8/14

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Best Practices | ❌ | ✅ | ✅ |
| Code Builds | ❌ | ✅ | ❌ |
| Detects sentinel value change and triggers full refresh | — | ✅ | — |
| Handles 304 Not Modified | — | ❌ | — |
| Implements conditional reads with matchConditions / setIfNoneMatch() using ETag | — | ❌ | — |
| Implements conditional reads with matchConditions/setIfNoneMatch using ETag | — | ❌ | — |
| Implements deterministic percentage rollout | — | ✅ | — |
| Parses JSON payload in feature flag setting values | — | ✅ | — |
| Retrieves settings with label using SettingSelector | — | ✅ | — |
| Uses .appconfig.featureflag/ prefix for feature flag keys | ✅ | ✅ | ❌ |

**Issues:**
- [Baseline] 304 Not Modified handled via fragile exception.getMessage().contains('304') instead of Response status code
- [Baseline] Async demo calls .block() defeating the purpose of async; does not demo async feature flags
- [Baseline] Code does not compile: getConfigurationSetting(String, String, null, MatchConditions) is not a valid SDK method signature
- [Baseline] Correct API for conditional reads is getConfigurationSettingWithResponse(ConfigurationSetting, OffsetDateTime, boolean ifChanged, Context)
- [Baseline] Does not use SDK's built-in FeatureFlagConfigurationSetting type
- [MCP] Does not use setIfNoneMatch() or matchConditions for conditional requests, so 304 Not Modified is not handled.
- [MCP] ETag caching is purely client-side; no setIfNoneMatch()/MatchConditions used for conditional HTTP requests, so all data is fully downloaded every time
- [MCP] Error handling swallows exceptions silently — no retries, no proper propagation
- [Skills] Errors silently caught and printed to console with no retry or propagation
- [Skills] Feature flag evaluator does not explicitly use the .appconfig.featureflag/ prefix.
- [MCP] No 304 Not Modified handling since conditional requests are never made
- [MCP] No HTTP-level conditional reads: ETags stored but never sent via setIfNoneMatch()/MatchConditions, so full payloads are always downloaded
- [MCP] No retry policies or timeout configurations on the clients
- [Skills] Not using the latest Azure SDK package versions.
- [Skills] azure-data-appconfiguration version 1.2.21 does not exist on Maven Central (latest is 1.8.0)
- [Baseline] azure-data-appconfiguration version 1.2.25 does not exist on Maven Central (latest is 1.8.0)
- [Baseline] azure-identity version 1.14.2 is outdated (latest 1.16.2)
- [MCP] azure-identity version 1.14.2 is outdated (latest is 1.16.2)
- [MCP] azure-identity version 1.14.2 is significantly outdated (latest stable is 1.18.2)
- [Skills] azure-identity version outdated (1.14.2 vs 1.16.2)
- [Skills] getConfigurationSettingWithResponse called with String where OffsetDateTime is expected - compilation error
- [MCP] getSettingsByPrefix always downloads all values despite ETag tracking (both branches of the if/else do the same thing)
- [MCP] getSettingsByPrefix with ETag tracking still downloads all values regardless of change status
- [Skills] onlyIfChanged parameter set to false, defeating the purpose of conditional reads
- [Skills] simulateChange method modifies local cache rather than demonstrating actual server-side changes

### `cosmos-db-dp-java-crud`

Scores: Baseline 10/12 | MCP 10/12 | Skills 9/12

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Best Practices | ✅ | ❌ | ✅ |
| Code Builds | ❌ | ✅ | ❌ |
| CosmosDatabase and CosmosContainer creation | ✅ | ✅ | ❌ |
| createItem readItem replaceItem deleteItem | ✅ | — | ✅ |
| createItem, readItem, replaceItem, deleteItem | — | ✅ | — |

**Issues:**
- [Baseline] Compilation failure: CosmosContainerResponse has no getContainer() method — should use database.getContainer(CONTAINER_NAME)
- [Baseline] Compilation failure: CosmosDatabaseResponse has no getDatabase() method — should use client.getDatabase(DATABASE_NAME)
- [Skills] Compilation failure: CosmosDatabaseResponse.getDatabase() and CosmosContainerResponse.getContainer() methods do not exist. Should use client.getDatabase(name) and database.getContainer(name) after createIfNotExists calls.
- [Skills] No mention of DefaultAzureCredential as a production best practice (minor, since prompt asked for key-based auth).
- [MCP] Uses key-based authentication instead of DefaultAzureCredential (though the prompt explicitly requested endpoint+key)
- [MCP] azure-cosmos version 4.45.0 is significantly outdated; latest stable is 4.71.0 (26+ versions behind)
- [Baseline] azure-cosmos version 4.50.0 is 21 minor versions behind latest 4.71.0
- [Skills] azure-cosmos version 4.53.0 is significantly outdated (latest stable is 4.71.0).

### `cosmos-db-dp-java-todo-repository`

Scores: Baseline 24/32 | MCP 13/19 | Skills 13/19

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Async query uses CosmosPagedFlux returning pages as a stream | ✅ | — | — |
| Configurable page size via QueryRequestOptions.setMaxItemCount | ❌ | — | ❌ |
| Configurable page size via setMaxItemCount | ❌ | ❌ | — |
| Correct partition key usage | ✅ | ✅ | — |
| Correct partition key usage: /category path, PartitionKey in all point operations | ✅ | — | ✅ |
| Does NOT flatten query results | ✅ | ✅ | — |
| Does NOT flatten query results (.stream() / .forEach() without page iteration) | ✅ | — | — |
| Does NOT flatten query results (no anti-pattern) | — | — | ✅ |
| ETag-based optimistic concurrency | ✅ | ✅ | — |
| ETag-based optimistic concurrency: captures ETag from read, passes ifMatchETag on update | ✅ | — | ✅ |
| Error Handling | ✅ | ❌ | ✅ |
| Handles 412 Precondition Failed | ✅ | — | — |
| Handles 412 Precondition Failed as a specific error case for conflicts | ✅ | — | — |
| Handles 412 Precondition Failed as specific error case | — | ✅ | ✅ |
| Page-by-page iteration using iterableByPage | ✅ | — | — |
| Page-by-page iteration using iterableByPage() or CosmosPagedIterable | ✅ | — | — |
| Parameterized queries using SqlQuerySpec | ❌ | — | — |
| Parameterized queries using SqlQuerySpec with SqlParameter | ❌ | ✅ | ✅ |
| Parameterized queries using SqlQuerySpec with SqlParameter (no string concatenation) | ✅ | — | — |
| RU cost extracted and logged per operation | ✅ | — | — |
| RU cost extracted from response via getRequestCharge() and logged per operation | ✅ | — | — |
| RU cost extracted via getRequestCharge() and logged | — | ✅ | — |
| RU cost extracted via getRequestCharge() and logged per operation | ✅ | — | ✅ |
| TTL configured at 90 days (7776000 seconds) | — | ✅ | ✅ |
| TTL configured at 90 days (7776000 seconds) via ContainerProperties.setDefaultTimeToLiveInSeconds() | ✅ | — | — |
| TTL configured at 90 days via setDefaultTimeToLiveInSeconds | ✅ | — | — |

**Issues:**
- [Baseline] 409 Conflict status code not handled
- [Skills] Bug: sync query sets item ETag to page continuation token (line 98), corrupting concurrency data for queried items
- [Baseline] Build fails: ExcludedPath() no-arg constructor does not exist - should use new ExcludedPath(String path)
- [Baseline] Build fails: SqlQuerySpec.setParameter(String,String) does not exist - should use SqlQuerySpec constructor with List<SqlParameter>
- [MCP] Catches generic Exception instead of CosmosException; uses fragile e.getMessage().contains("412") instead of getStatusCode()
- [MCP] Compilation error: ExcludedPath and IncludedPath require String constructor argument, not no-arg constructor with setPath()
- [Skills] Compilation fails: CosmosDatabaseResponse.getDatabase() method does not exist — should use cosmosClient.getDatabase() directly
- [Skills] Compilation fails: ExcludedPath requires a String constructor argument, not a no-arg constructor
- [Baseline] Compilation failure: ExcludedPath requires String path in constructor, no no-arg constructor available
- [Baseline] Compilation failure: SqlQuerySpec.setParameter(String,String) does not exist - should use SqlParameter objects
- [Baseline] Continuation token is not logged during page iteration.
- [Baseline] Continuation token not logged during pagination
- [Skills] Continuation token not logged per page — getContinuationToken() only used incorrectly as ETag
- [Baseline] No 409 conflict handling for create operations
- [Baseline] No configurable page size (setMaxItemCount not called)
- [Baseline] No configurable page size via setMaxItemCount
- [MCP] No continuation token logging per page
- [Skills] No handling for 404 (not found) or 409 (conflict on create) status codes
- [Baseline] No retry policy or timeout configuration on CosmosClient
- [Skills] No setMaxItemCount for configurable page size in either sync or async query
- [MCP] No setMaxItemCount() for configurable page size
- [MCP] No specific handling for 404 (not found) or 409 (conflict) status codes
- [Baseline] Outdated SDK versions: azure-cosmos 4.53.1 (latest 4.71.0), azure-identity 1.11.1 (latest 1.16.2)
- [Baseline] Package versions outdated: azure-cosmos 4.53.1 (latest 4.71.0), azure-identity 1.11.1 (latest 1.16.2)
- [Baseline] Query page size is not configurable via setMaxItemCount.
- [Skills] SDK versions significantly outdated (azure-cosmos 4.53.0 vs 4.71.0, azure-identity 1.11.1 vs 1.16.2)
- [MCP] SDK versions significantly outdated (azure-cosmos 4.55.0 vs 4.71.0+, azure-identity 1.11.4 vs 1.16.2+)
- [Baseline] azure-cosmos and azure-identity are not the latest versions.

### `event-hubs-dp-java-streaming`

Scores: Baseline 11/12 | MCP 11/12 | Skills 12/13

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| azure-messaging-eventhubs and azure-messaging-eventhubs-checkpointstore-blob Maven deps | — | — | ✅ |

**Issues:**
- [Skills] Azure SDK dependencies are not the latest versions.
- [MCP] EVENT_HUB_NAME constant declared but never used
- [MCP] Java source file at project root instead of src/main/java/ (non-standard Maven layout requiring manual reorganization)
- [Baseline] Java source file placed at project root instead of src/main/java (non-standard Maven layout)
- [MCP] Package versions significantly outdated: azure-messaging-eventhubs 5.18.0 (latest 5.20.3), checkpointstore-blob 1.19.0 (latest 1.20.7)
- [Skills] SDK versions are outdated: azure-messaging-eventhubs 5.19.0 vs latest 5.20.3, checkpointstore-blob 1.20.0 vs latest 1.20.7
- [Skills] Unused imports: EventPosition and Duration are imported but never used
- [Baseline] azure-messaging-eventhubs version 5.18.0 is outdated; latest stable is 5.20.3
- [Baseline] azure-messaging-eventhubs-checkpointstore-blob version 1.19.0 is outdated; latest stable is 1.20.7

### `identity-dp-java-credential-chain`

Scores: Baseline 16/18 | MCP 30/33 | Skills 9/12

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Anti-Pattern: NOT using DefaultAzureCredential as CI credential | — | ✅ | — |
| Anti-Pattern: NOT using DefaultAzureCredential for CI | ✅ | — | — |
| Anti-Patterns | — | — | ✅ |
| Anti-Patterns (scenario-specific) | — | ✅ | — |
| Anti-pattern: NOT using DefaultAzureCredential as CI credential | — | ✅ | — |
| Async tester uses reactive Mono<AccessToken> | ✅ | — | — |
| Async tester uses reactive getToken returning Mono | — | ✅ | — |
| CAE Support | ✅ | ❌ | ❌ |
| CAE Support via setCaeEnabled or enableCae on builders | — | ❌ | — |
| CI chain uses EnvironmentCredential not DefaultAzureCredential | — | ✅ | — |
| CI chain uses EnvironmentCredential or AzurePipelinesCredential | ✅ | — | — |
| Calls getToken and prints token expiry | — | ✅ | — |
| Code Builds | ❌ | ✅ | ❌ |
| Credential Chain Construction - ChainedTokenCredentialBuilder | — | ✅ | — |
| Credentials added via .addLast() | — | ✅ | — |
| Dev chain includes AzureCliCredential and others | — | ✅ | — |
| Dev chain includes CLI and IDE credentials | — | ✅ | — |
| Dev chain includes developer tool credentials | ✅ | — | — |
| Environment Detection | — | ✅ | ✅ |
| Environment Detection - CI | ✅ | ✅ | — |
| Environment Detection - Dev fallback | ✅ | — | — |
| Environment Detection - Falls back to dev | — | ✅ | — |
| Environment Detection - Production | — | ✅ | — |
| Environment Detection - Production/Managed Identity | ✅ | — | — |
| Environment Detection CI | — | ✅ | — |
| Environment Detection Production | — | ✅ | — |
| Environment Detection fallback to dev | — | ✅ | — |
| Environment-Specific Chains | — | ✅ | ✅ |
| Failure handling with specific exception info | ✅ | — | — |
| Handles failure with specific exception info | — | ✅ | — |
| Prints token expiry from AccessToken.getExpiresAt() | — | ✅ | — |
| Production chain ManagedIdentity first with WorkloadIdentity fallback | — | ✅ | — |
| Production chain with ManagedIdentity first and WorkloadIdentity fallback | ✅ | ✅ | — |
| Scenario-Specific Async | — | ✅ | ✅ |
| Token Request & Testing | — | ✅ | ✅ |
| Token Request with correct scope | ✅ | ✅ | — |
| Token getToken and expiry printing | ✅ | — | — |

**Issues:**
- [MCP] Azure SDK dependencies are not the latest versions.
- [Skills] Build fails: EnvironmentCredentialBuilder.additionallyAllowedTenants() does not exist in azure-identity 1.13.0 (2 compilation errors)
- [Skills] CAE enableCae flag stored but never used — no setCaeEnabled(true) call on TokenRequestContext and no enableCae() on credential builders
- [MCP] CAE enablement uses setClaims() workaround instead of TokenRequestContext.setCaeEnabled(true)
- [MCP] CAE is implemented via setClaims with raw xms_cc JSON rather than using TokenRequestContext.setCaeEnabled(true)
- [Baseline] CAE on credential builders incorrectly uses tokenCachePersistenceOptions instead of enableCae()
- [Baseline] Code does not compile: tokenCachePersistenceOptions() doesn't exist on EnvironmentCredentialBuilder, ManagedIdentityCredentialBuilder, or WorkloadIdentityCredentialBuilder (4 compilation errors)
- [MCP] CredentialFactory.enableCae field is stored but never applied to any credential builder or token request
- [Baseline] Outdated SDK versions: azure-identity 1.13.2 (latest 1.16.2), azure-core 1.51.0 (latest 1.55.4)
- [MCP] The enableCae field on CredentialFactory is never used to configure credentials - only used for display strings
- [Baseline] Unused import javax.naming.ServiceUnavailableException in ConnectivityTester.java
- [Skills] azure-core 1.49.0 is significantly outdated (latest stable: 1.57.1)
- [MCP] azure-core 1.51.0 is 4 minor versions behind latest 1.55.4
- [MCP] azure-core version 1.51.0 is outdated; latest is 1.55.4
- [Skills] azure-identity 1.13.0 is significantly outdated (latest stable: 1.18.2)
- [MCP] azure-identity 1.13.2 is 3 minor versions behind latest 1.16.2
- [MCP] azure-identity version 1.13.2 is outdated; latest is 1.16.2

### `identity-dp-java-default-credential`

Scores: Baseline 8/10 | MCP 7/10 | Skills 9/11

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Credential chain order in Java SDK | ✅ | ❌ | ✅ |
| Passing credential to client builders (e.g., SecretClientBuilder) | — | — | ✅ |

**Issues:**
- [Skills] BOM version 1.2.29 is outdated; latest stable is 1.3.5
- [Skills] Build fails: excludeEnvironmentCredential() does not exist on DefaultAzureCredentialBuilder in azure-identity 1.14.0
- [MCP] Code does not compile — excludeEnvironmentCredential() method does not exist on Java's DefaultAzureCredentialBuilder (verified via Maven build with both v1.11.0 and v1.16.2)
- [Skills] Credential chain order has minor inaccuracies for Java SDK (IntelliJ position, VS Code may be deprecated)
- [MCP] InteractiveBrowserCredential incorrectly listed in Java DefaultAzureCredential chain
- [Baseline] Java source and logback.xml files placed in root directory instead of Maven-standard src/main/java/ and src/main/resources/
- [Baseline] README also references non-existent excludeEnvironmentCredential() and excludeAzurePowerShellCredential() methods in Exclude Credentials section
- [MCP] README troubleshooting section references non-existent exclude methods throughout (.NET patterns applied to Java)
- [Baseline] azure-core-http-netty version 1.13.7 is outdated, latest is 1.15.12
- [Skills] azure-identity direct version 1.15.0 mentioned in README is outdated; latest stable is 1.18.2
- [Baseline] azure-identity version 1.11.0 is 5+ minor versions behind latest 1.16.2
- [MCP] azure-identity version 1.11.0 is outdated (latest: 1.16.2)
- [Skills] azure-sdk-bom version 1.2.29 is outdated; latest versions should be used for azure-identity and azure-security-keyvault-secrets.
- [Baseline] azure-security-keyvault-secrets version 4.6.0 is 4 minor versions behind latest 4.10.0
- [MCP] azure-security-keyvault-secrets version 4.7.0 is outdated (latest: 4.10.0)
- [Baseline] excludeAzurePowerShellCredential() method does not exist on DefaultAzureCredentialBuilder — causes compilation failure (confirmed by build attempts with both v1.11.0 and v1.16.2)
- [MCP] retryTimeout() method does not exist on Java's DefaultAzureCredentialBuilder (correct method is credentialProcessTimeout())

### `identity-dp-java-managed-identity`

Scores: Baseline 8/11 | MCP 8/11 | Skills 9/11

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| CredentialUnavailableException when not in Azure | ❌ | ❌ | ✅ |

**Issues:**
- [Baseline] All Azure SDK package versions are significantly outdated (e.g., azure-identity 1.11.0 vs 1.16.2, azure-cosmos 4.52.0 vs 4.71.0)
- [Baseline] Build fails with 3 compilation errors: incorrect Cosmos DB API (getDatabase/getContainer on response objects) and non-existent BlobRetryOptions class in com.azure.storage.blob.models
- [MCP] Build fails: imports internal class com.azure.identity.implementation.IdentityClientException that is not accessible
- [MCP] Build fails: missing 'new' keyword in AzureSDKClientExamples.java line 187 (java.io.ByteArrayInputStream)
- [Skills] Build failure: KeyVaultErrorException referenced from com.azure.security.keyvault.secrets.models but only exists in com.azure.security.keyvault.secrets.implementation.models
- [Skills] Build failure: Missing import for com.azure.core.credential.TokenCredential in ErrorHandlingExample.java (line 157)
- [Skills] Build failure: Missing import for com.azure.core.credential.TokenCredential in LocalDevelopmentFallbackExample.java (line 134)
- [MCP] ClientAuthenticationException.getResponse() may return null but is dereferenced without null check in detailedErrorHandling
- [MCP] CredentialUnavailableException is never referenced - this is the specific exception thrown when managed identity is unavailable outside Azure
- [Baseline] CredentialUnavailableException is never referenced despite being an explicit prompt-specific requirement
- [MCP] README hardcodes azure-identity version 1.11.0 (latest stable is 1.18.2)
- [Skills] README references azure-identity 1.15.0 when latest is 1.16.2
- [MCP] azure-sdk-bom version 1.2.18 is outdated; latest stable is 1.3.5
- [Skills] azure-sdk-bom version 1.2.29 is significantly outdated; latest is 1.3.5
- [Baseline] retryTimeout lambda usage on ManagedIdentityCredentialBuilder may not match actual API signature

### `identity-dp-java-service-principal`

Scores: Baseline 9/10 | MCP 9/10 | Skills 8/10

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Code Builds | ✅ | ✅ | ❌ |

**Issues:**
- [Skills] COMPILATION ERROR: credential.getToken(requestContext) returns Mono<AccessToken>, needs .block() call to get AccessToken synchronously
- [MCP] Java source file not placed in Maven conventional directory structure (src/main/java/com/example/azure/) — 'mvn compile' finds no sources as delivered
- [MCP] SLF4J version 2.0.9 is also slightly outdated
- [Baseline] Unused import java.util.Map in ServicePrincipalAuthExample.java line 14
- [Baseline] azure-identity resolves to 1.11.1 via BOM, latest is 1.16.2
- [MCP] azure-identity version 1.11.0 is outdated; latest stable is 1.16.2
- [Baseline] azure-resourcemanager version 2.33.0 is significantly behind latest 2.51.0 (18 versions behind)
- [Baseline] azure-sdk-bom version 1.2.19 is significantly behind latest 1.2.35 (16 versions behind)
- [Skills] azure-sdk-bom version 1.2.29 is outdated, latest stable is 1.3.5
- [MCP] azure-security-keyvault-secrets version 4.7.0 is outdated; latest stable is 4.10.0
- [Baseline] slf4j 2.0.9 is not the latest stable release
- [Skills] slf4j-simple version 2.0.9 may not be latest

### `key-vault-dp-java-crud`

Scores: Baseline 8/10 | MCP 9/10 | Skills 9/10

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Exception handling for HttpResponseException | ❌ | ✅ | ✅ |

**Issues:**
- [Skills] Extra unrelated managed identity example files in tmp/ directory were not requested
- [Baseline] HttpResponseException is not explicitly caught; only its subclass ResourceNotFoundException is handled
- [Baseline] Package versions are outdated: azure-security-keyvault-secrets 4.8.0 should be 4.10.0, azure-identity 1.13.0 should be 1.16.2
- [MCP] azure-identity version 1.13.0 is outdated; latest stable is 1.16.2
- [Skills] azure-identity version 1.15.0 is outdated; latest is 1.16.2
- [MCP] azure-security-keyvault-secrets version 4.8.0 is outdated; latest stable is 4.10.0
- [Skills] azure-security-keyvault-secrets version 4.9.0 is outdated; latest is 4.10.0

### `key-vault-dp-java-secret-config`

Scores: Baseline 18/20 | MCP 12/15 | Skills 24/25

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Async uses PollerFlux to wait for delete completion | ✅ | ❌ | ✅ |
| Code Builds | ❌ | ❌ | ✅ |
| Configurable warning window for near-expiry | — | — | ✅ |
| Creates new secret only after delete completes | ✅ | — | ✅ |
| In-memory caching (e.g., ConcurrentHashMap) with bulk-load and single-key refresh | ✅ | — | ✅ |
| In-memory caching with ConcurrentHashMap, bulk-load and single-key refresh | — | ✅ | ✅ |
| In-memory caching with bulk-load and single-key refresh | ✅ | — | ✅ |
| NOT using fire-and-forget deleteSecret() | — | — | ✅ |
| NOT using fire-and-forget deleteSecret() without waiting | ✅ | — | — |
| Returns default value when secret not found | ✅ | — | ✅ |
| Secret expiry via properties().getExpiresOn() | — | — | ✅ |
| Secret expiry: accesses properties().getExpiresOn() | ✅ | — | ✅ |
| Secret rotation uses beginDeleteSecret() as LRO | — | — | ✅ |
| Secret versioning via getSecret(name, version) | — | — | ✅ |

**Issues:**
- [MCP] Async demo in Main.java skips rotation step, violating the prompt requirement to repeat the full demo
- [Baseline, Skills] Azure SDK BOM version is outdated; should be updated to the latest stable version.
- [MCP] Azure SDK packages are significantly outdated (keyvault-secrets latest is 4.10.6, identity latest is 1.18.2)
- [MCP] Compilation error: beginDeleteSecret().getFinalResult() returns Void, not DeletedSecret
- [MCP] No async secret rotation helper — PollerFlux pattern is completely absent
- [Baseline] SLF4J version 2.0.9 is outdated
- [Baseline] SecretRotationHelper.java:51 - waitForCompletion() returns PollResponse<DeletedSecret>, not DeletedSecret, causing compilation failure
- [MCP] SecretRotationHelper.waitForDeletionComplete has redundant logic — getFinalResult() already blocks until deletion completes
- [Baseline] azure-sdk-bom version 1.14.2 does not exist; latest is 1.2.35
- [Skills] azure-sdk-bom version 1.2.19 is outdated; latest stable is 1.2.35
- [Skills] azure-sdk-bom version 1.2.19 is outdated; latest stable is 1.3.5
- [MCP] azure-security-keyvault-secrets version 4.8.8 does not exist on Maven Central — build fails entirely

### `resource-manager-mp-java-rg-crud`

Scores: Baseline 10/12 | MCP 11/12 | Skills 11/12

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Code Builds | ❌ | ✅ | ✅ |

**Issues:**
- [Baseline] Compilation failure: ResourceGroup.location() does not exist; should use regionName() (line 151, verified by Maven build)
- [Baseline] Outdated azure-identity version 1.13.0 (latest is 1.16.2)
- [Baseline] Outdated azure-resourcemanager version 2.40.0 (latest is 2.51.0)
- [Baseline] Uses withTags() (map variant) instead of withTag() (single key-value) as specified in criteria, though functionally equivalent
- [Skills] Uses withTags() (plural, replaces all tags) instead of withTag() (singular, adds individual tag) — functionally equivalent but not an exact API match to the criterion
- [MCP] Uses withTags(Map) instead of withTag(key,value) for tag additions, though functionally equivalent
- [Skills] azure-identity 1.13.2 is several versions behind latest 1.16.2
- [MCP] azure-identity version 1.13.3 is behind latest 1.16.2 (3 minor versions behind)
- [Skills] azure-resourcemanager 2.43.0 is 8 versions behind latest 2.51.0
- [MCP] azure-resourcemanager version 2.43.0 is significantly behind latest 2.51.0 (8 versions behind)

### `service-bus-dp-java-crud`

Scores: Baseline 10/12 | MCP 11/12 | Skills 11/12

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Error Handling | ❌ | ✅ | ✅ |

**Issues:**
- [Baseline] CountDownLatch is created but never counted down; unconventional for a simple delay
- [MCP] Java source file placed in project root rather than src/main/java (standard Maven layout)
- [Baseline] No retry/timeout configuration on clients
- [Baseline] No try-catch exception handling around Service Bus send/receive operations
- [Skills] Source file placed in project root instead of src/main/java/ (Maven standard layout)
- [Skills] Unused import: java.util.List is imported but never used
- [Skills] azure-identity dependency included but never used in code
- [Baseline] azure-identity version 1.10.0 is significantly behind latest stable 1.16.2
- [Skills] azure-identity version 1.12.0 is significantly outdated; latest stable is 1.16.2
- [MCP] azure-messaging-servicebus version 7.14.0 is outdated; latest stable is 7.17.11
- [Baseline] azure-messaging-servicebus version 7.14.0 is significantly behind latest stable 7.17.11
- [Skills] azure-messaging-servicebus version 7.17.0 is outdated; latest stable is 7.17.11

### `service-bus-dp-java-order-processor`

Scores: Baseline 11/17 | MCP 11/17 | Skills 11/17

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Best Practices | ❌ | ✅ | ✅ |
| Correlation sets order ID via setCorrelationId or application properties | ✅ | — | — |
| Correlation sets order ID via setCorrelationId() | — | ✅ | — |
| Correlation sets order ID via setCorrelationId() or application properties | — | — | ✅ |
| Dead-letter explicitly dead-letters failed messages with reason string | — | — | ✅ |
| Dead-letter explicitly dead-letters with reason string | ✅ | ✅ | — |
| Error Handling | ✅ | ❌ | ❌ |
| Error handler logs entity path and error source | ✅ | ❌ | ❌ |
| Handles message not fitting in current batch | — | ✅ | — |
| Handles message that doesn't fit in current batch | ✅ | — | ✅ |
| Processor uses .processor().queueName().processMessage().processError() chain | ✅ | ✅ | ❌ |
| Scheduled delivery uses setScheduledEnqueueTime with ~30s delay | ✅ | — | — |
| Scheduled delivery uses setScheduledEnqueueTime() with ~30s delay | — | ✅ | ✅ |
| Sender uses .sender().queueName().buildClient() chain | ❌ | ✅ | ✅ |
| Session-aware processing uses .sessionProcessor() or session-enabled receiver | ❌ | — | — |
| Session-aware processing uses sessionProcessor or session-enabled receiver | — | ✅ | — |
| Session-aware processing uses sessionProcessor() or session-enabled receiver | — | — | ✅ |

**Issues:**
- [Baseline] Async batch sender has concurrency issue with shared mutable messageBatch in flatMap
- [Skills] Async batch sending creates a new batch per message instead of accumulating messages into batches
- [MCP] Build fails with 6 compilation errors: missing imports for SubQueue and DeadLetterOptions from com.azure.messaging.servicebus.models package
- [Skills] Does not use .processor()/.sessionProcessor() pattern with processMessage()/processError() callbacks
- [Baseline] Fabricated APIs: ServiceBusSessionReceiverClientBuilder, ServiceBusReceiverClientBuilder, serviceBusClient.createReceiver/createSessionReceiver
- [Baseline] Missing imports for DeadLetterOptions and SubQueue (in com.azure.messaging.servicebus.models subpackage)
- [Skills] No error handler logging entity path or error source from ServiceBusErrorContext
- [Skills] No retry configuration or timeout configuration on the client builder
- [Baseline] No transient vs non-transient error discrimination in error handler
- [MCP, Skills] No transient vs non-transient error distinction via isTransient() or getReason()
- [MCP] OrderProcessor.deadLetter() uses non-existent 3-arg signature (message, String, String) — should use DeadLetterOptions object
- [Skills] OrderProcessorAsync.java fails to compile with 9 errors - wrong types for session receiver, getMessage() on ServiceBusReceivedMessage, missing complete()/deadLetter() methods
- [MCP] Processor error handlers don't log entity path or error source from ServiceBusErrorContext
- [Baseline] SDK versions significantly outdated (servicebus 7.17.0 vs 7.17.17, identity 1.13.0 vs 1.18.2)
- [Skills] SDK versions significantly outdated: servicebus 7.17.1 vs 7.17.11, identity 1.13.1 vs 1.16.2
- [Baseline] Sender does not use the correct .sender().queueName().buildClient() builder chain — uses fabricated createSender() method
- [Baseline] ServiceBusClient class does not exist in the Azure SDK — causes 17 compile errors across all files
- [Baseline] ServiceBusReceiverAsyncClient.receiveMessages(int) does not exist — takes no arguments
- [Baseline] Uses .processor().sessionEnabled(true) instead of .sessionProcessor() for session-aware processing
- [MCP] azure-identity 1.13.0 significantly outdated (latest stable: 1.16.2)
- [MCP] azure-messaging-servicebus 7.17.0 outdated (latest stable: 7.17.17)

### `storage-dp-java-blob-event-notifier`

Scores: Baseline 8/15 | MCP 12/15 | Skills 18/22

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Best Practices | ❌ | ✅ | ✅ |
| Catches Event Grid-specific exceptions for publishing errors | ✅ | ❌ | ❌ |
| Code Builds | ❌ | ❌ | ✅ |
| Does NOT manually parse JSON without SDK deserialization helpers | ❌ | ✅ | ✅ |
| Does NOT manually parse JSON without the SDK's deserialization helpers | — | — | ✅ |
| Handles CloudEvents 1.0 schema via CloudEvent.fromString() | ❌ | ✅ | ✅ |
| Handles CloudEvents 1.0 schema via CloudEvent.fromString() deserialization | — | — | ✅ |
| Handles Event Grid native schema via EventGridEvent.fromString() | ❌ | ✅ | ✅ |
| Handles Event Grid native schema via EventGridEvent.fromString() deserialization | — | — | ✅ |
| Handles race condition with BlobStorageException 404 | ✅ | — | ✅ |
| Handles race condition: blob may no longer exist (catches BlobStorageException with 404 status) | — | — | ✅ |
| Handles race condition: blob may no longer exist (catches BlobStorageException with 404) | — | ✅ | ✅ |
| Logs a warning for unrecognized event types | — | — | ✅ |
| Logs warning for unrecognized event types | ✅ | ✅ | — |
| Parses container and blob name from event subject | ❌ | — | — |
| Parses container name and blob name from event subject | — | ✅ | ❌ |
| Parses container name and blob name from event subject (/blobServices/default/containers/{container}/blobs/{blob}) | — | — | ✅ |
| Publishes custom events with subject hierarchy for filtering | ✅ | ✅ | ❌ |
| Routes events based on event type string (Microsoft.Storage.BlobCreated, Microsoft.Storage.BlobDeleted) | — | — | ✅ |

**Issues:**
- [Skills] Azure SDK version is not the latest.
- [MCP] BOM version 1.2.29 is outdated (latest is 1.2.35)
- [MCP] Build fails: CloudEvent constructor 4th param must be CloudEventDataFormat enum, not String
- [MCP] Build fails: EventGridPublisherClientBuilder.buildClient()/buildAsyncClient() are private; must use buildCloudEventPublisherClient()/buildCloudEventPublisherAsyncClient()
- [MCP] Build fails: onErrorMap lambda type mismatch in AsyncEventPublisher
- [Skills] CloudEvent constructor parameters are swapped: subject is passed as 'type' and eventType as 'dataContentType', setSubject() is never called
- [Skills] CloudEvent constructor parameters swapped in EventPublisher: subject passed as type, eventType passed as dataContentType
- [Baseline] Code fails to compile with 6 errors: EventGridEvent data parameter requires BinaryData not Object; EventGridPublisherClientBuilder.buildClient()/buildAsyncClient() require Class<T> parameter
- [Skills] Container/blob name parsed from blob URL instead of event subject path
- [Skills] Container/blob parsed from eventData.getUrl() instead of event subject path
- [Baseline] Generic Exception catch in publishers instead of Event Grid-specific exceptions like HttpResponseException
- [Baseline] Manual JSON parsing with Jackson ObjectMapper instead of SDK's EventGridEvent.fromString() and CloudEvent.fromString()
- [Skills] No Event Grid-specific exception handling in publisher classes
- [Skills] No explicit catch for Event Grid publishing exceptions.
- [Skills] No try-catch for Event Grid publishing exceptions in EventPublisher or EventPublisherAsync
- [MCP] Publishing error handling catches generic Exception instead of Event Grid-specific HttpResponseException
- [Baseline] Subject parsing fallback is incorrect: extracts 'blobs' as container name instead of actual container from /blobServices/default/containers/{container}/blobs/{blob}
- [Baseline] Unused imports in EventReceiver: EventGridEvent, StorageBlobCreatedEventData, StorageBlobDeletedEventData, ArrayList, List
- [Skills] azure-sdk-bom version 1.2.19 is outdated (latest is 1.3.5)
- [Skills] azure-sdk-bom version 1.2.19 is outdated (latest: 1.3.5)
- [Baseline] azure-sdk-bom version 1.2.27 is outdated; latest stable is 1.3.5

### `storage-dp-java-blob-manager`

Scores: Baseline 9/12 | MCP 12/15 | Skills 9/12

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| Code Builds | ❌ | ✅ | ❌ |
| Configures custom retry policy (exponential backoff, max retries, delay) | ❌ | ✅ | ✅ |
| Implements blob lease acquisition before overwrite | — | ✅ | — |
| Implements parallel/block upload for large files (ParallelTransferOptions, not manual chunking) | — | ✅ | — |
| Properly composes reactive chains in demo | ✅ | — | — |
| Properly composes reactive chains in the demo | — | ❌ | ❌ |
| Sets blob index tags on upload (Map<String, String> via upload options) | ✅ | — | — |
| Sets blob index tags on upload (not just metadata) | — | — | ✅ |
| Sets blob index tags on upload (not just metadata) — Map<String, String> via upload options | — | ✅ | — |
| Sets blob index tags on upload via upload options | — | ✅ | — |
| Sets per-request or per-operation timeout | ✅ | ❌ | ✅ |

**Issues:**
- [MCP] Async demo uses .block() on each call instead of reactive chain composition
- [Skills] Async demo uses .block() on each step individually instead of composing a reactive chain with .then()/.flatMap()
- [MCP] Async demo uses sequential .block() calls instead of reactive chain composition (flatMap/then)
- [Baseline] Async uploadFromFile creates ParallelTransferOptions but never uses them (dead code)
- [Skills] BOM version 1.2.24 is outdated; latest stable is 1.3.5
- [Baseline] BlobParallelUploadOptions(Flux<ByteBuffer>, long) constructor doesn't exist; Flux variant takes only Flux<ByteBuffer> without length — causes compilation failure
- [Baseline] BlobUploadFromFileOptions imported from com.azure.storage.blob.models instead of com.azure.storage.blob.options — causes compilation failure
- [Skills] Compilation error: RetryPolicy(ExponentialBackoffOptions) is not a valid constructor — needs RetryOptions wrapper
- [MCP] Main demo acquires a lease then calls uploadBlobWithLease which acquires a second lease — would fail at runtime with 409 Conflict
- [Baseline] RetryOptions() no-arg constructor does not exist; must use RetryOptions(ExponentialBackoffOptions) or RetryOptions(FixedDelayOptions) — causes compilation failure
- [Baseline] RetryOptions.setMaxRetries() and setTryTimeout() are not valid methods on RetryOptions class
- [MCP] SDK packages are several major versions behind (blob 12.25.1 vs 12.30.0, identity 1.11.2 vs 1.16.2)
- [MCP] SDK versions are not the latest as of 2026.
- [MCP] Unused imports in BlobStorageConfiguration (HttpClient, HttpPipeline, NettyAsyncHttpClientBuilder, etc.)
- [Baseline] azure-identity 1.11.1 is significantly behind latest stable 1.16.2
- [MCP] azure-identity 1.11.2 is 5 minor versions behind latest stable 1.16.2
- [Baseline] azure-storage-blob 12.25.0 is 5 versions behind latest stable 12.30.0
- [MCP] azure-storage-blob 12.25.1 is 5 minor versions behind latest stable 12.30.0
- [Skills] createSyncClient/createAsyncClient hardcode retry values (5, 1s, 30s) instead of using builder-configured parameters
- [Skills] exponentialBackoffRetry builder method creates RetryPolicy that is never actually used by client creation methods
- [MCP] perRequestTimeout field is a dead config value - stored but never passed to builder or API calls
- [MCP] perRequestTimeout field is stored but never passed to BlobServiceClientBuilder or individual operations
- [MCP] perRequestTimeout is not applied to client or operations.

### `storage-dp-java-crud`

Scores: Baseline 11/12 | MCP 11/12 | Skills 11/12

*All criteria identical across configs.*

### `storage-dp-java-encrypted-uploader`

Scores: Baseline 22/25 | MCP 51/53 | Skills 22/25

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| AES-GCM | — | ✅ | — |
| AES-GCM - Generates random IV | — | — | ✅ |
| AES-GCM - Uses AES-GCM mode | — | — | ✅ |
| AES-GCM: Generates random 12-byte IV | ✅ | ✅ | — |
| AES-GCM: Uses AES-GCM (not CBC/ECB) | ✅ | — | — |
| AES-GCM: Uses AES-GCM mode | — | ✅ | — |
| Anti-Pattern - NOT encrypting data directly with vault key | — | — | ✅ |
| Anti-Pattern - NOT storing raw DEK in plaintext | — | — | ✅ |
| Anti-Pattern - NOT using SecretClient | — | — | ✅ |
| Anti-Pattern: NOT encrypting data directly with vault key | — | ✅ | — |
| Anti-Pattern: NOT storing raw DEK in plaintext | — | ✅ | — |
| Anti-Pattern: NOT using SecretClient | — | ✅ | — |
| Anti-Patterns (scenario-specific) | — | ✅ | — |
| Anti-pattern: NOT encrypting data directly with vault key | ✅ | — | — |
| Anti-pattern: NOT storing raw DEK in plaintext | ✅ | — | — |
| Anti-pattern: NOT using SecretClient | ✅ | — | — |
| Async: Uses BlobAsyncClient and CryptographyAsyncClient | ✅ | ✅ | — |
| Best Practices | ❌ | ✅ | ✅ |
| Client Construction (scenario-specific) | — | ✅ | — |
| Client Construction - Uses KeyClient/CryptographyClient builder (NOT SecretClient) | — | — | ✅ |
| Client Construction: Uses KeyClient/CryptographyClient (not SecretClient) | ✅ | — | — |
| Client Construction: Uses KeyClient/CryptographyClient builders | — | ✅ | — |
| Decryption retrieves wrapped DEK, unwraps via KV, decrypts locally | — | ✅ | — |
| Dependencies (scenario-specific) | — | ✅ | — |
| Dependencies - Uses azure-security-keyvault-keys | — | — | ✅ |
| Dependencies - Uses javax.crypto for local AES-GCM | — | — | ✅ |
| Dependencies: Uses azure-security-keyvault-keys | — | ✅ | — |
| Dependencies: Uses azure-security-keyvault-keys (not Secrets) | ✅ | — | — |
| Dependencies: Uses javax.crypto/java.security for local AES-GCM | — | ✅ | — |
| Dependencies: Uses javax.crypto/java.security for local encryption | ✅ | — | — |
| Encrypts data with AES-GCM locally using DEK | — | ✅ | — |
| Envelope Encryption - Decryption retrieves and unwraps correctly | — | — | ✅ |
| Envelope Encryption - Encrypts data with AES-GCM locally | — | — | ✅ |
| Envelope Encryption - Generates random AES-256 DEK locally | — | — | ✅ |
| Envelope Encryption - Stores IV in blob metadata | — | — | ✅ |
| Envelope Encryption - Stores vault key identifier in blob metadata | — | — | ✅ |
| Envelope Encryption - Stores wrapped DEK as blob metadata | — | — | ✅ |
| Envelope Encryption - Wraps DEK via Key Vault wrapKey | — | — | ✅ |
| Envelope Encryption Patterns (critical) | — | ✅ | — |
| Envelope Encryption: Decryption retrieves metadata, unwraps DEK, decrypts locally | — | ✅ | — |
| Envelope Encryption: Decryption retrieves wrapped DEK, unwraps, decrypts locally | ✅ | — | — |
| Envelope Encryption: Encrypts data with AES-GCM locally | ✅ | ✅ | — |
| Envelope Encryption: Generates random AES-256 DEK locally | ✅ | ✅ | — |
| Envelope Encryption: Stores IV in blob metadata | ✅ | ✅ | — |
| Envelope Encryption: Stores vault key identifier in blob metadata | ✅ | ✅ | — |
| Envelope Encryption: Stores wrapped DEK as blob metadata | ✅ | ✅ | — |
| Envelope Encryption: Wraps DEK via Key Vault wrapKey | ✅ | ✅ | — |
| Error Handling: Handles Key Vault errors | ✅ | ✅ | — |
| Generates random AES-256 DEK locally | — | ✅ | — |
| Generates random IV for each encryption (12 bytes) | — | ✅ | — |
| Handles Key Vault errors | — | ✅ | — |
| KV Keys: Key material never leaves Key Vault | — | ✅ | — |
| KV Keys: Specifies RSA key wrap algorithm | — | ✅ | — |
| KV Keys: Uses CryptographyClient for wrapKey/unwrapKey | — | ✅ | — |
| Key Vault Keys - Key material never leaves Key Vault | — | — | ✅ |
| Key Vault Keys - Specifies RSA key wrap algorithm | — | — | ✅ |
| Key Vault Keys - Uses CryptographyClient wrapKey/unwrapKey | — | — | ✅ |
| Key Vault Keys Patterns (critical) | — | ✅ | — |
| Key Vault Keys: Key material never leaves Key Vault | ✅ | — | — |
| Key Vault Keys: Specifies RSA key wrap algorithm | ✅ | — | — |
| Key Vault Keys: Uses CryptographyClient for wrapKey/unwrapKey | ✅ | — | — |
| Key material never leaves Key Vault | — | ✅ | — |
| NOT encrypting data directly with vault key | — | ✅ | — |
| NOT storing raw DEK in plaintext | — | ✅ | — |
| NOT using SecretClient | — | ✅ | — |
| Scenario-Specific Async | — | ✅ | — |
| Scenario-Specific Async - Uses BlobAsyncClient and CryptographyAsyncClient | — | — | ✅ |
| Scenario-Specific Error Handling | — | ✅ | — |
| Scenario-Specific Error Handling - Key Vault errors | — | — | ❌ |
| Specifies RSA key wrap algorithm | — | ✅ | — |
| Stores IV in blob metadata | — | ✅ | — |
| Stores vault key identifier in blob metadata | — | ✅ | — |
| Stores wrapped DEK as blob metadata | — | ✅ | — |
| Uses AES-GCM (not CBC/ECB) | — | ✅ | — |
| Uses BlobAsyncClient and CryptographyAsyncClient for async | — | ✅ | — |
| Uses CryptographyClient for wrapKey() and unwrapKey() | — | ✅ | — |
| Uses KeyClient/CryptographyClient builder (NOT SecretClient) | — | ✅ | — |
| Uses azure-security-keyvault-keys (not Secrets) | — | ✅ | — |
| Uses javax.crypto or java.security for local AES-GCM | — | ✅ | — |
| Wraps DEK via Key Vault wrapKey() | — | ✅ | — |

**Issues:**
- [MCP] 10 compilation errors prevent the code from building
- [MCP] 10+ compilation errors prevent the code from building
- [Baseline] All Azure SDK versions are outdated (BOM 1.2.28 vs 1.2.35, keyvault-keys 4.9.1 vs 4.10.0, identity 1.14.2 vs 1.16.2, storage-blob 12.28.1 vs latest)
- [Baseline] Anti-pattern: specifying explicit dependency versions while also using the BOM — should let BOM manage versions
- [MCP] BlobAsyncClient.upload() called with (Flux<ByteBuffer>, int, boolean) but needs (Flux<ByteBuffer>, ParallelTransferOptions, boolean)
- [MCP] BlobAsyncClient.upload(Flux,int,boolean) signature mismatch - second param should be ParallelTransferOptions
- [Baseline] CRITICAL: Code does not compile — getHttpPipeline() is package-private in KeyClient/KeyAsyncClient and CryptographyClientBuilder.credential() expects TokenCredential, not HttpPipeline (9 compilation errors)
- [Skills] CRITICAL: CryptographyClientBuilder.credential() receives HttpPipeline instead of TokenCredential — code does not compile (9 errors)
- [MCP] CryptographyClient.getKeyId() does not exist in the SDK
- [MCP] CryptographyClientBuilder constructed incorrectly — tries to pass HttpPipeline to credential() instead of using pipeline() method or passing the shared TokenCredential
- [MCP] CryptographyClientBuilder.credential() is passed HttpPipeline instead of TokenCredential
- [MCP] KeyClient.getHttpPipeline() is not a public API method
- [Skills] KeyClient.getHttpPipeline() is not a public API — cannot be accessed from outside package
- [Skills] Mono<SecretKeySpec> cannot be assigned to Mono<SecretKey> in AsyncKeyManager.unwrapDataEncryptionKey
- [Skills] No specific error handling for Key Vault key disabled or key not found scenarios (generic catch only)
- [MCP] Type inference issues in async code (Mono<SecretKeySpec> vs Mono<SecretKey>)
- [Baseline] Type mismatch in KeyManagerAsync: Mono<SecretKeySpec> cannot convert to Mono<SecretKey>
- [Baseline] Unused imports in EncryptedBlobUploader (IvParameterSpec, ByteArrayInputStream)
- [Skills] Upload and setMetadata are two separate calls — metadata could be lost if setMetadata fails after upload succeeds
- [MCP] azure-sdk-bom 1.2.27 is outdated; latest is 1.3.5
- [Skills] azure-sdk-bom version 1.2.24 is outdated; latest stable is 1.3.5
- [MCP] azure-sdk-bom version 1.2.27 is outdated; latest is 1.3.5
- [MCP] azure-sdk-bom version is not the latest (should be 1.2.28).
- [MCP] getKeyId() called on CryptographyClient/CryptographyAsyncClient but method doesn't exist — should use getKey() and then getId()

### `storage-mp-java-account-mgmt`

Scores: Baseline 9/13 | MCP 9/13 | Skills 10/13

| Criterion | Baseline | Azure MCP | Skills |
|-----------|----------|-----------|--------|
| update() or service properties update for blob versioning | ❌ | — | — |
| update().withBlobAccessTier() or service properties update | — | ❌ | ✅ |

**Issues:**
- [Baseline] All Azure SDK dependencies are significantly outdated (21+ versions behind for storage)
- [MCP] All three Azure SDK packages are significantly outdated (e.g., storage 2.43.0 vs 2.51.0, identity 1.13.2 vs 1.16.2)
- [Baseline] Java source file placed in root instead of src/main/java (Maven standard layout)
- [MCP] Subscription ID read from environment and validated but never passed to AzureProfile constructor (works via env var fallback but is misleading)
- [Skills] azure-core 1.51.0 is outdated (latest 1.55.4)
- [Skills] azure-identity 1.13.2 is outdated (latest 1.16.2)
- [Skills] azure-resourcemanager-storage 2.43.0 is outdated (latest 2.51.0)
- [Skills] azure-resourcemanager-storage is not the latest version.
- [Baseline] blobServices().getServiceProperties() method doesn't exist on BlobServices interface
- [Skills] inner() method called on StorageAccount fluent interface should be innerModel() - 3 compilation errors
- [MCP] withBlobVersioningEnabled() does not exist on StorageAccount.Update interface; blob versioning must be configured via BlobServiceProperties — confirmed by compilation
- [Baseline] withIsVersioningEnabled() method doesn't exist on BlobServiceProperties.Update
- [Baseline] withSku() expects StorageAccountSkuType not Sku - should use StorageAccountSkuType.STANDARD_LRS
- [MCP] withSku() expects StorageAccountSkuType.STANDARD_LRS, not new Sku().withName(SkuName.STANDARD_LRS) — confirmed by compilation
- [Skills] withSku() uses new Sku().withName(SkuName.STANDARD_LRS) instead of StorageAccountSkuType.STANDARD_LRS - compilation error
- [Baseline] withSubscription() method does not exist on StorageManager - subscription should be set via AzureProfile constructor

---

## 3. Common Failures vs. Fixed by Skills

### 3a. Criteria Failing in ALL Three Configs

These criteria failed regardless of config — systemic gaps:

| Prompt | Criterion |
|--------|-----------|
| `app-configuration-dp-java-crud` | Latest Package Versions |
| `app-configuration-dp-java-feature-flags` | Error Handling |
| `app-configuration-dp-java-feature-flags` | Handles 304 Not Modified (setting unchanged since last read) |
| `app-configuration-dp-java-feature-flags` | Implements conditional reads with matchConditions / setIfNoneMatch() using the setting's ETag |
| `app-configuration-dp-java-feature-flags` | Latest Package Versions |
| `cosmos-db-dp-java-crud` | Latest Package Versions |
| `cosmos-db-dp-java-todo-repository` | Catches CosmosException with status code checks (404, 409, 412) |
| `cosmos-db-dp-java-todo-repository` | Code Builds |
| `cosmos-db-dp-java-todo-repository` | Latest Package Versions |
| `cosmos-db-dp-java-todo-repository` | Logs continuation token and item count per page |
| `event-hubs-dp-java-streaming` | Latest Package Versions |
| `identity-dp-java-credential-chain` | Latest Package Versions |
| `identity-dp-java-default-credential` | Code Builds |
| `identity-dp-java-default-credential` | Latest Package Versions |
| `identity-dp-java-managed-identity` | Code Builds |
| `identity-dp-java-managed-identity` | Latest Package Versions |
| `identity-dp-java-service-principal` | Latest Package Versions |
| `key-vault-dp-java-crud` | Latest Package Versions |
| `key-vault-dp-java-secret-config` | Latest Package Versions |
| `resource-manager-mp-java-rg-crud` | Latest Package Versions |
| `service-bus-dp-java-crud` | Latest Package Versions |
| `service-bus-dp-java-order-processor` | Code Builds |
| `service-bus-dp-java-order-processor` | Distinguishes transient vs non-transient errors |
| `service-bus-dp-java-order-processor` | Latest Package Versions |
| `storage-dp-java-blob-event-notifier` | Latest Package Versions |
| `storage-dp-java-blob-manager` | Latest Package Versions |
| `storage-dp-java-crud` | Latest Package Versions |
| `storage-dp-java-encrypted-uploader` | Code Builds |
| `storage-dp-java-encrypted-uploader` | Latest Package Versions |
| `storage-mp-java-account-mgmt` | Code Builds |
| `storage-mp-java-account-mgmt` | Latest Package Versions |
| `storage-mp-java-account-mgmt` | storageAccounts().define().withRegion().withExistingResourceGroup().withSku().create() |

### 3b. Criteria Fixed by Skills (Failed Baseline → Passed Skills)

These criteria failed in baseline but passed with skills — demonstrating skill impact:

| Prompt | Criterion |
|--------|-----------|
| `app-configuration-dp-java-feature-flags` | Best Practices |
| `cosmos-db-dp-java-todo-repository` | Parameterized queries using SqlQuerySpec with SqlParameter |
| `identity-dp-java-managed-identity` | CredentialUnavailableException when not in Azure |
| `key-vault-dp-java-crud` | Exception handling for HttpResponseException |
| `key-vault-dp-java-secret-config` | Code Builds |
| `resource-manager-mp-java-rg-crud` | Code Builds |
| `service-bus-dp-java-crud` | Error Handling |
| `service-bus-dp-java-order-processor` | Best Practices |
| `service-bus-dp-java-order-processor` | Sender uses .sender().queueName().buildClient() chain |
| `storage-dp-java-blob-event-notifier` | Best Practices |
| `storage-dp-java-blob-event-notifier` | Code Builds |
| `storage-dp-java-blob-event-notifier` | Does NOT manually parse JSON without SDK deserialization helpers |
| `storage-dp-java-blob-event-notifier` | Handles CloudEvents 1.0 schema via CloudEvent.fromString() |
| `storage-dp-java-blob-event-notifier` | Handles Event Grid native schema via EventGridEvent.fromString() |
| `storage-dp-java-blob-manager` | Configures custom retry policy (exponential backoff, max retries, delay) |
| `storage-dp-java-encrypted-uploader` | Best Practices |

### 3c. Criteria Regressed by Skills (Passed Baseline → Failed Skills)

These criteria passed in baseline but failed with skills — potential regressions:

| Prompt | Criterion |
|--------|-----------|
| `app-configuration-dp-java-feature-flags` | Uses .appconfig.featureflag/ prefix for feature flag keys |
| `cosmos-db-dp-java-crud` | CosmosDatabase and CosmosContainer creation |
| `identity-dp-java-credential-chain` | CAE Support |
| `identity-dp-java-service-principal` | Code Builds |
| `service-bus-dp-java-order-processor` | Error Handling |
| `service-bus-dp-java-order-processor` | Error handler logs entity path and error source |
| `service-bus-dp-java-order-processor` | Processor uses .processor().queueName().processMessage().processError() chain |
| `storage-dp-java-blob-event-notifier` | Catches Event Grid-specific exceptions for publishing errors |
| `storage-dp-java-blob-event-notifier` | Publishes custom events with subject hierarchy for filtering |

### 3d. Most Common Universal Failures (by criterion name)

| Criterion | # Prompts Failing |
|-----------|-------------------|
| Latest Package Versions | 19 |
| Code Builds | 6 |
| Error Handling | 1 |
| Handles 304 Not Modified (setting unchanged since last read) | 1 |
| Implements conditional reads with matchConditions / setIfNoneMatch() using the setting's ETag | 1 |
| Catches CosmosException with status code checks (404, 409, 412) | 1 |
| Logs continuation token and item count per page | 1 |
| Distinguishes transient vs non-transient errors | 1 |
| storageAccounts().define().withRegion().withExistingResourceGroup().withSku().create() | 1 |

---

## 4. Duration Comparison

| # | Prompt ID | Baseline (s) | Azure MCP (s) | Skills (s) |
|---|-----------|-------------|---------------|------------|
| 1 | `app-configuration-dp-java-crud` | 48.3 | 60.2 | 56.5 |
| 2 | `app-configuration-dp-java-feature-flags` | 233.1 | 236.7 | 178.4 |
| 3 | `cosmos-db-dp-java-crud` | 69.5 | 61.0 | 72.5 |
| 4 | `cosmos-db-dp-java-todo-repository` | 201.3 | 211.0 | 178.7 |
| 5 | `event-hubs-dp-java-streaming` | 56.1 | 67.2 | 99.1 |
| 6 | `identity-dp-java-credential-chain` | 187.0 | 364.7 | 219.9 |
| 7 | `identity-dp-java-default-credential` | 124.7 | 94.7 | 313.9 |
| 8 | `identity-dp-java-managed-identity` | 242.5 | 199.9 | 274.3 |
| 9 | `identity-dp-java-service-principal` | 116.7 | 85.4 | 184.7 |
| 10 | `key-vault-dp-java-crud` | 67.5 | 60.7 | 125.0 |
| 11 | `key-vault-dp-java-secret-config` | 252.8 | 197.2 | 331.9 |
| 12 | `resource-manager-mp-java-rg-crud` | 97.0 | 102.3 | 84.8 |
| 13 | `service-bus-dp-java-crud` | 54.1 | 65.9 | 101.7 |
| 14 | `service-bus-dp-java-order-processor` | 204.5 | 204.1 | 212.0 |
| 15 | `storage-dp-java-blob-event-notifier` | 197.8 | 263.7 | 600.2 |
| 16 | `storage-dp-java-blob-manager` | 216.8 | 526.8 | 371.9 |
| 17 | `storage-dp-java-crud` | 85.6 | 101.5 | 120.3 |
| 18 | `storage-dp-java-encrypted-uploader` | 220.9 | 227.4 | 232.5 |
| 19 | `storage-mp-java-account-mgmt` | 62.3 | 75.7 | 72.8 |
|---|-----------|-------------|---------------|------------|
| | **Average** | **144.1** | **168.7** | **201.6** |
| | **Total** | **2738.5** | **3206.1** | **3831.0** |

> Skills config is on average **40%** slower than baseline,
> while Azure MCP is **17%** slower than baseline.

---

## 5. Summary Statistics

| Metric | Baseline | Azure MCP | Skills |
|--------|----------|-----------|--------|
| Total Score | 221/279 | 262/314 | 224/274 |
| Score % | 79.2% | 83.4% | 81.8% |
| Verification Pass | 16/19 (84%) | 15/19 (79%) | 16/19 (84%) |
| Avg Duration (s) | 144.1 | 168.7 | 201.6 |
| Criteria Passed | 221/279 (79.2%) | 263/314 (83.8%) | 225/274 (82.1%) |
| Total Files Generated | 105 | 136 | 139 |
| Universal Failures | 32 criteria | 32 criteria | 32 criteria |
| Fixed by Skills | — | — | 16 criteria |
| Regressions (Skills) | — | — | 9 criteria |

### Per-Prompt Best Config

| # | Prompt | Baseline % | MCP % | Skills % | Best |
|---|--------|-----------|-------|----------|------|
| 1 | `app-configuration/crud` | 92% | 92% | 92% | **Tie** |
| 2 | `app-configuration/feature-flags` | 57% | 67% | 57% | **Azure MCP** |
| 3 | `cosmos-db/crud` | 83% | 83% | 75% | **Baseline / Azure MCP** |
| 4 | `cosmos-db/todo-repository` | 75% | 68% | 68% | **Baseline** |
| 5 | `event-hubs/streaming` | 92% | 92% | 92% | **Skills** |
| 6 | `identity/credential-chain` | 89% | 91% | 75% | **Azure MCP** |
| 7 | `identity/default-credential` | 80% | 70% | 82% | **Skills** |
| 8 | `identity/managed-identity` | 73% | 73% | 82% | **Skills** |
| 9 | `identity/service-principal` | 90% | 90% | 80% | **Baseline / Azure MCP** |
| 10 | `key-vault/crud` | 80% | 90% | 90% | **Azure MCP / Skills** |
| 11 | `key-vault/secret-config` | 90% | 80% | 96% | **Skills** |
| 12 | `resource-manager/mgmt/rg-crud` | 83% | 92% | 92% | **Azure MCP / Skills** |
| 13 | `service-bus/crud` | 83% | 92% | 92% | **Azure MCP / Skills** |
| 14 | `service-bus/order-processor` | 65% | 65% | 65% | **Tie** |
| 15 | `storage/blob-event-notifier` | 53% | 80% | 82% | **Skills** |
| 16 | `storage/blob-manager` | 75% | 80% | 75% | **Azure MCP** |
| 17 | `storage/crud` | 92% | 92% | 92% | **Tie** |
| 18 | `storage/encrypted-uploader` | 88% | 96% | 88% | **Azure MCP** |
| 19 | `storage/mgmt/account-mgmt` | 69% | 69% | 77% | **Skills** |

**Win counts** (highest score %): Baseline=1, Azure MCP=4, Skills=6, Ties=8

### Verification Differences

Prompts where verification status differs across configs:

| Prompt | Baseline | MCP | Skills |
|--------|----------|-----|--------|
| `identity-dp-java-credential-chain` | ✅ | ✅ | ❌ |
| `key-vault-dp-java-secret-config` | ✅ | ❌ | ✅ |
| `storage-dp-java-blob-event-notifier` | ❌ | ✅ | ✅ |
| `storage-dp-java-blob-manager` | ✅ | ❌ | ✅ |
