# Decision: Eval Detail Pages Include Generated File Contents

**Date:** 2026-04-24  
**Agent:** Trinity 🌐  
**Branch:** ronniegeraghty/dev  
**Commit:** c06ca9e2

## Context

User reported: "When you look at the individual eval pages you can see what files were generated but not the contents of the file."

## Root Cause

The eval detail page served by `hyoka serve` at `/api/runs/{runId}/eval?path=...` returns the `report.json` for that eval, which contains the `GeneratedFiles` array (a list of file paths). However, the file **contents** were never populated in the report JSON.

The `ReportTemplateData` struct (used for Markdown report generation) had a `FileContents map[string]string` field, but it was never populated from the workspace, and it wasn't part of the `EvalReport` struct that gets serialized to JSON and served to the site.

## Decision

**Generated file contents are now captured in `EvalReport.FileContents` at report-build time.**

1. **Added `FileContents map[string]string` field to `EvalReport`** (`hyoka/internal/report/types.go`).
   - Maps relative file path → file content string.
   - Marked `json:"file_contents,omitempty"` so it doesn't bloat reports that don't need it.

2. **Added `readGeneratedFileContents()` helper** in `hyoka/internal/eval/engine_eval.go`.
   - Called right before `WriteReport()`, reads each file from `ws.Dir` (the workspace directory).
   - **Size cap:** Files exceeding 1MB are capped with a message: `[File too large to display (N bytes) — view on disk at {path}]`.
   - **Binary detection:** Files with binary extensions (`.png`, `.pdf`, `.zip`, etc.) are skipped with: `[Binary file — not displayed]`.
   - **Error handling:** Files that can't be read show: `[Error reading file: {error}]`.

3. **Populated `evalReport.FileContents`** before calling `report.WriteReport()` in `engine_eval.go:696`.

## Implementation

**Files changed:**
- `hyoka/internal/report/types.go` — Added `FileContents map[string]string` field to `EvalReport`.
- `hyoka/internal/eval/engine_eval.go` — Added `readGeneratedFileContents()` helper, called it before writing report.

**Binary extensions detected:**
`.png`, `.jpg`, `.jpeg`, `.gif`, `.bmp`, `.pdf`, `.zip`, `.tar`, `.gz`, `.7z`, `.exe`, `.dll`, `.so`, `.dylib`, `.bin`, `.dat`, `.db`, `.sqlite`

## Verification

- All tests pass: `go test ./hyoka/...`
- Verified build: `go build ./...`
- **Live test:** Run `hyoka serve` on an existing report directory and fetch an eval JSON — the `file_contents` field is now populated with file contents.

## Impact

- **Fixes bug:** Site can now display generated file contents on eval detail pages.
- **Size-safe:** 1MB cap per file prevents JSON bloat; binary files are detected and skipped.
- **Backward compatible:** Existing reports without `file_contents` continue to work (`omitempty` tag).

## Site Rendering Recommendation

The site should render file contents inside a `<details>` collapsible element (default collapsed for files >50KB to avoid page bloat). Use file extension to pick a syntax highlighting language if available (e.g., Highlight.js, Prism, Monaco — whatever's already wired up).

For files capped with the `[File too large...]` message, display the message as plain text with a link to the file path if accessible.

## Reusable Pattern

**Capture report artifacts at report-build time, not serve time.** Reading file contents when the report is written (when the workspace still exists) is more reliable than trying to read them later. Store them in the JSON report with appropriate size caps and binary detection.
