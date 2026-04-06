---
name: "serve-patterns"
description: "API endpoints, SPA dashboard, path traversal protection"
domain: "architecture"
confidence: "high"
source: "hyoka/internal/serve/serve.go"
---

## Context

The `hyoka serve` command launches a local web server for browsing evaluation reports. The server provides RESTful endpoints for retrieving reports and serves a single-page application (SPA) dashboard. Security focuses on path traversal protection and sandboxing report access.

## Server Architecture

```
┌─────────────────────┐
│  SPA Dashboard      │
│  (React/JS)         │
└──────────┬──────────┘
           │ HTTP requests
           ↓
┌─────────────────────────────────────┐
│  REST API (hyoka serve)             │
│  :8080                              │
│  ├─ GET /api/reports                │
│  ├─ GET /api/reports/{id}           │
│  ├─ GET /api/reports/{id}/json      │
│  └─ GET /static/*                   │
└──────────┬──────────────────────────┘
           │
           ↓
      reports/
      ├─ {run_id1}/
      │  ├─ report.json
      │  └─ report.html
      └─ {run_id2}/
         ├─ report.json
         └─ report.html
```

## API Endpoints

### List Reports
```
GET /api/reports

Response:
{
  "reports": [
    {
      "id": "eval-2024-04-06-123456",
      "prompt": "identity-dp-python-default-credential",
      "config": "baseline/claude-opus-4.6",
      "timestamp": "2024-04-06T12:34:56Z",
      "status": "success"
    }
  ]
}
```

### Get Report Metadata
```
GET /api/reports/{id}

Response:
{
  "id": "eval-2024-04-06-123456",
  "prompt": {...},
  "config": {...},
  "generation": { "status": "success", "duration_ms": 45000 },
  "build": { "status": "success" },
  "graders": [...],
  "review": {...}
}
```

### Get Full Report JSON
```
GET /api/reports/{id}/json

Response:
(Full report.json from disk)
```

### Serve Static Files
```
GET /static/{path}

Serves from reports/{id}/assets/ with path traversal protection
```

## Path Traversal Protection

Critical security feature: prevent access to files outside `reports/` directory.

**Vulnerable pattern:**
```go
// ✗ WRONG: Directly concatenates user input
filePath := filepath.Join(reportDir, userInput)
// If userInput = "../../etc/passwd", could read arbitrary files!
```

**Safe pattern:**
```go
// ✓ CORRECT: Validate path is within sandbox
filePath := filepath.Join(reportDir, filepath.Clean(userInput))
if !strings.HasPrefix(filePath, reportDir) {
    return fmt.Errorf("path traversal attempt")
}
// Read file safely
```

### Implementation

```go
// hyoka/internal/serve/serve.go
func (s *Server) serveReport(w http.ResponseWriter, r *http.Request) {
    reportID := chi.URLParam(r, "id")
    
    // Validate report ID format
    if !isValidRunID(reportID) {
        http.Error(w, "Invalid report ID", http.StatusBadRequest)
        return
    }
    
    // Construct safe path
    reportPath := filepath.Join(s.reportDir, reportID, "report.json")
    
    // Verify path is within sandbox
    absReportPath, _ := filepath.Abs(reportPath)
    absSandbox, _ := filepath.Abs(s.reportDir)
    if !strings.HasPrefix(absReportPath, absSandbox) {
        http.Error(w, "Forbidden", http.StatusForbidden)
        return
    }
    
    // Read and serve
    data, err := os.ReadFile(reportPath)
    // ...
}
```

## Routing

Use a router (e.g., chi) for clean, type-safe route handling:

```go
import "github.com/go-chi/chi/v5"

func (s *Server) setupRoutes() *chi.Mux {
    r := chi.NewRouter()
    
    // API routes
    r.Get("/api/reports", s.listReports)
    r.Get("/api/reports/{id}", s.getReport)
    r.Get("/api/reports/{id}/json", s.getReportJSON)
    
    // Static files (SPA assets)
    r.Get("/static/*", s.serveStatic)
    
    // SPA fallback (serve index.html for unmatched routes)
    r.NotFound(s.serveSPA)
    
    return r
}
```

## SPA (Single-Page Application) Serving

The dashboard is a static SPA (HTML + JS) that:
1. Loads list of reports via `/api/reports`
2. Fetches detailed report via `/api/reports/{id}`
3. Renders interactive UI in browser

```go
// Serve index.html for all non-API routes
func (s *Server) serveSPA(w http.ResponseWriter, r *http.Request) {
    if r.RequestURI == "/" {
        // Home page
        http.ServeFile(w, r, filepath.Join(s.staticDir, "index.html"))
        return
    }
    
    // Route not found
    http.NotFound(w, r)
}
```

## CORS (Cross-Origin Resource Sharing)

If SPA is on different domain than API:

```go
func (s *Server) setupCORS() func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            next.ServeHTTP(w, r)
        })
    }
}
```

## Error Handling

```go
func (s *Server) getReport(w http.ResponseWriter, r *http.Request) {
    reportID := chi.URLParam(r, "id")
    
    // Validate ID
    if reportID == "" {
        http.Error(w, "Report ID required", http.StatusBadRequest)
        return
    }
    
    // Read report
    data, err := s.readReport(reportID)
    if err != nil {
        if os.IsNotExist(err) {
            http.Error(w, "Report not found", http.StatusNotFound)
            return
        }
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        slog.Error("Failed to read report", "id", reportID, "error", err)
        return
    }
    
    // Serve
    w.Header().Set("Content-Type", "application/json")
    w.Write(data)
}
```

## Server Startup

```bash
# Default (localhost:8080)
go run ./hyoka serve

# Custom port
go run ./hyoka serve --port 9000

# Custom report directory
go run ./hyoka serve --report-dir ./my-reports
```

Implementation:

```go
var serveCmd = &cobra.Command{
    Use:   "serve",
    Short: "Start local web server for browsing reports",
    RunE: func(cmd *cobra.Command, args []string) error {
        port, _ := cmd.Flags().GetInt("port")
        reportDir, _ := cmd.Flags().GetString("report-dir")
        
        server, err := NewServer(reportDir, port)
        if err != nil {
            return err
        }
        
        slog.Info("Server starting", "url", fmt.Sprintf("http://localhost:%d", port))
        return server.Start()
    },
}
```

## Testing

```go
func TestPathTraversal(t *testing.T) {
    server := NewServer("./reports", 8080)
    
    tests := []struct {
        reportID string
        expect   int  // HTTP status
    }{
        {"eval-valid-123", 200},
        {"../../etc/passwd", 400},  // Bad format
        {"./../secret", 400},       // Path traversal attempt
    }
    
    for _, tt := range tests {
        req := httptest.NewRequest("GET", fmt.Sprintf("/api/reports/%s", tt.reportID), nil)
        w := httptest.NewRecorder()
        server.getReport(w, req)
        
        if w.Code != tt.expect {
            t.Errorf("reportID %q: expected %d, got %d", tt.reportID, tt.expect, w.Code)
        }
    }
}
```

## Code Locations

- **Server implementation:** `hyoka/internal/serve/serve.go`
- **API handlers:** `hyoka/internal/serve/handlers.go` (if split)
- **Serve command:** `hyoka/cmd/serve.go`
- **Static assets:** `hyoka/site/` or `hyoka/templates/`

## Anti-Patterns

- Not validating report IDs (enables directory traversal)
- Serving files outside `reports/` directory
- Not escaping HTML in JSON responses (XSS risk)
- Hardcoding port in code (use flags)
- Not logging API errors (makes debugging hard)
- Assuming localhost is secure (anyone with network access can reach port 8080)
