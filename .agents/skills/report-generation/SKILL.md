---
name: "report-generation"
description: "JSON/HTML/Markdown reports, template system"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/report/*.go"
---

## Context

Hyoka generates multi-format evaluation reports (JSON, HTML, Markdown) from a unified internal representation. Reports capture the entire evaluation timeline: generation, build, grading, and review phases.

## Report Formats

### JSON Report

The canonical format — contains all raw data and is used to generate other formats:

```json
{
  "metadata": {
    "run_id": "eval-2024-04-06-123456",
    "timestamp": "2024-04-06T12:34:56Z",
    "prompt": {
      "id": "identity-dp-python-default-credential",
      "service": "identity",
      "language": "python"
    },
    "config": {
      "name": "baseline/claude-opus-4.6"
    }
  },
  "generation": {
    "status": "success",
    "duration_ms": 45000,
    "code": "# Generated Python code...",
    "action_timeline": {
      "events": [...],
      "summary": {...}
    }
  },
  "build": {
    "status": "success",
    "duration_ms": 5000,
    "stdout": "Compilation successful",
    "stderr": ""
  },
  "graders": [
    {
      "kind": "behavior",
      "name": "tool_compliance",
      "pass": true,
      "score": 1.0,
      "message": "All required tools used"
    }
  ],
  "review": {
    "status": "success",
    "reviewers": [
      {
        "model": "claude-opus-4.6",
        "score": 0.9,
        "findings": "Code is mostly correct..."
      }
    ],
    "consensus_score": 0.85
  }
}
```

### HTML Report

Browser-viewable report with interactive elements:

```html
<!DOCTYPE html>
<html>
<head>
  <title>Evaluation Report: identity-dp-python</title>
  <style>...</style>
</head>
<body>
  <div class="report-container">
    <h1>Evaluation Report</h1>
    <section class="generation">
      <h2>Generation Phase</h2>
      <code>{{ .Code }}</code>
      <div class="timeline">...</div>
    </section>
    <section class="graders">
      <h2>Grading Results</h2>
      ...
    </section>
  </div>
  <script>
    // Interactive features (expand/collapse, filtering)
  </script>
</body>
</html>
```

### Markdown Report

Human-readable text format for documentation and sharing:

```markdown
# Evaluation Report: identity-dp-python

## Generation
- Status: Success
- Duration: 45s

## Code
\`\`\`python
# Generated code...
\`\`\`

## Graders
| Grader | Status | Score | Message |
|--------|--------|-------|---------|
| behavior | Pass | 1.0 | All required tools used |
| lint | Pass | 0.95 | 1 warning |

## Review Panel
- Claude Opus 4.6: 0.9
- GPT-5.4: 0.85
- Consensus: 0.875
```

## Report Generation Pipeline

1. **Evaluation Completion:** Engine returns `EvaluationResult` with all phase outputs
2. **JSON Marshaling:** Convert result to JSON bytes
3. **File Write:** Save JSON to `reports/{run_id}/report.json`
4. **Template Rendering:** Apply HTML/Markdown templates to JSON
5. **Asset Linking:** Copy static assets (CSS, JS) to report directory

## Template System

Reports use Go's `text/template` package:

```go
// report.go
type ReportTemplate struct {
    HTML      *template.Template
    Markdown  *template.Template
}

func (rt *ReportTemplate) RenderHTML(data EvaluationResult) (string, error) {
    var buf bytes.Buffer
    if err := rt.HTML.Execute(&buf, data); err != nil {
        return "", err
    }
    return buf.String(), nil
}
```

### Template Variables

Templates can access all fields from `EvaluationResult`:

```handlebars
{{- .Metadata.RunID }}
{{- .Generation.Status }}
{{- range .Graders }}
  {{- .Kind }}: {{- .Score }}
{{- end }}
```

## Report Storage

Reports are written to:
```
reports/
  {run_id}/
    report.json           # Canonical data
    report.html           # Browser view
    report.md             # Text format
    assets/
      style.css
      script.js
```

## Multi-Report Aggregation

For batch runs (multiple prompts × configs):

```json
{
  "batch_metadata": {
    "run_date": "2024-04-06",
    "prompts": 5,
    "configs": 2,
    "total_evals": 10
  },
  "summary": {
    "avg_score": 0.82,
    "pass_count": 8,
    "fail_count": 2
  },
  "reports": [
    { "prompt_id": "...", "config": "...", "result": {...} }
  ]
}
```

## Re-rendering from JSON

The `rerender` command regenerates HTML/Markdown from saved JSON:

```bash
go run ./hyoka rerender --report-id eval-2024-04-06-123456 --format html
```

This allows updating templates without re-running evaluations.

## Code Locations

- **Report structure:** `hyoka/internal/report/report.go`
- **JSON marshaling:** `hyoka/internal/report/json.go`
- **HTML/Markdown rendering:** `hyoka/internal/report/render.go`
- **Templates:** `hyoka/templates/`

## Anti-Patterns

- Hardcoding paths in templates (use config for asset paths)
- Not including action timeline in JSON (required for reproducibility)
- Rendering HTML/Markdown before generation completes (wait for all phases)
- Losing build/grader details in summary format (preserve full data in JSON)
- Assuming report directory exists (create it before writing)
