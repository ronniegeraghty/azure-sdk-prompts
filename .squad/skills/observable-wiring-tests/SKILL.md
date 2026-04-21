# Observable Wiring Tests

> Prove the wire is hot, not just that it's connected.

## Problem

You add a plugin point — interface + registry + config field. Tests cover:

- Config YAML parses ✅
- Registry stores the thing ✅
- Interface satisfied ✅

…and the user-visible feature still doesn't work, because **nothing in production calls into the registered thing**. The plumbing is parseable but not actually invoked at runtime.

This is the **#587 trap**: config exists, code compiles, tests are green, behavior never happens. We hit it on plugins (#587), and Neo guarded against it explicitly in WI-027/PR #605 with `TestCustomFetcherInvokedAtRuntime`.

## The Pattern

Write at least one test per extension point that:

1. **Registers a real mock** against the **real production registry** (not a test-only fake registry).
2. **Invokes the real entry point** the user's code path would hit (e.g., `ResolveSkills`, not the registry's `Lookup` method directly).
3. **Asserts on observable runtime evidence**:
   - A call counter that increments (`atomic.Int64` works well).
   - The payload the mock received matches what the config promised (e.g., resolved version, expanded args).
4. **Cleans up** with `t.Cleanup(func() { Registry.Unregister(...) })` so the global registry stays clean for sibling tests.

```go
func TestCustomFetcherInvokedAtRuntime(t *testing.T) {
    mock := &mockFetcher{name: "test-runtime", matchRepo: "acme/widgets"}
    if err := DefaultRegistry.Register(mock); err != nil {
        t.Fatalf("register: %v", err)
    }
    t.Cleanup(func() { DefaultRegistry.Unregister(mock.name) })

    _, err := ResolveSkills([]Entry{
        {Type: TypeSkill, Source: SourceRemote, Repo: "acme/widgets", Version: "v1.2.3"},
    }, t.TempDir())
    if err != nil { t.Fatal(err) }

    if got := mock.calls.Load(); got != 1 {
        t.Fatalf("mock.Fetch was never called — wiring is dead")
    }
    if req := mock.lastReq.Load(); req.Version != "v1.2.3" {
        t.Errorf("mock saw version %q, config said v1.2.3", req.Version)
    }
}
```

## Why It Works

- **Real registry** catches "we registered into the wrong global."
- **Real entry point** catches "no production code path actually calls Lookup."
- **Payload assertion** catches "we dispatched, but with wrong/stale data" (e.g., config field never threaded through).
- **Counter == 1** (not `>= 1`) catches "we called the fallback *and* the custom" or "we called twice by mistake."

## When To Use

Any time you add:

- A plugin/registry/strategy interface.
- A new config field that's supposed to change runtime behavior.
- A "users can extend X" extension point.

If the only way to know your code ran is to `slog.Info` and squint at logs, you owe yourself an observable wiring test.

## When NOT To Use

- Pure-function unit tests already cover the dispatch (no global registry, no async, no I/O).
- The "wire" is a single function call you can grep in 10 seconds.

## Anti-pattern this replaces

```go
// ❌ Tests config parses, but never proves the parsed value reaches the runtime.
func TestConfigParsesCustomFetcher(t *testing.T) {
    cfg, _ := Parse(yaml)
    if cfg.Fetchers[0].Name != "custom" { t.Fail() }
}
```

Green test, broken feature.

## Origin

- **#587** (plugin trap): config wired up, plugin never invoked at runtime.
- **PR #605 / WI-027** (Neo): pattern formalized via `TestCustomFetcherInvokedAtRuntime`.
- **Promoted by:** Morpheus 🕶️ during PR #605 architectural review.
