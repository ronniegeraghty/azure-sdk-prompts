// Package serve provides a local web server for browsing evaluation reports.
//
// # Security
//
// The serve command is designed for local, single-user inspection of
// evaluation reports. It has no authentication, no authorization, and no
// rate limiting. The default listener binds to all interfaces on the chosen
// port (e.g. 0.0.0.0:8080), which means anyone who can reach the host on
// that port can:
//
//   - Read every evaluation report, prompt, and doc the server is configured
//     to expose, including prompt text, generated source, and model output.
//   - Enumerate run IDs and download raw JSON via /api/runs and /reports/.
//   - Access the embedded SPA without any credential.
//
// CORS is intentionally permissive (Access-Control-Allow-Origin: *) so the
// dev-mode Vite server can hit the API; this is safe for localhost use but
// should NOT be exposed on a shared network.
//
// Operator guidance for shared machines:
//
//   - Prefer binding to 127.0.0.1 via a reverse proxy or SSH tunnel rather
//     than exposing the port directly.
//   - Treat the reports directory as sensitive: it may contain prompt text,
//     agent output, and evaluation results that reveal internal grading
//     rubrics or proprietary prompts.
//   - Do not run `hyoka serve` on an untrusted network without first placing
//     an authenticating proxy (nginx, caddy) in front of it.
//
// These expectations are also printed as a banner at startup.
package serve

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/fs"
	"log/slog"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

// Options configures the serve command.
type Options struct {
	ReportsDir string
	DocsDir    string
	SiteDir    string
	PromptsDir string
	Port       int
}

// DocInfo holds metadata about a documentation file.
type DocInfo struct {
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	Content string `json:"content,omitempty"`
}

// internalDocs lists documentation files that should be excluded from the API.
// Developer docs (architecture, contributing) belong in the repo, not the user-facing site.
var internalDocs = map[string]bool{
	"architecture": true,
}

// Start launches a local HTTP server for browsing reports.
func Start(opts Options) error {
	if opts.Port == 0 {
		opts.Port = 8080
	}

	abs, err := filepath.Abs(opts.ReportsDir)
	if err != nil {
		return fmt.Errorf("resolving reports dir: %w", err)
	}
	opts.ReportsDir = abs

	if opts.DocsDir != "" {
		if d, err := filepath.Abs(opts.DocsDir); err == nil {
			opts.DocsDir = d
		}
	}
	if opts.SiteDir != "" {
		if d, err := filepath.Abs(opts.SiteDir); err == nil {
			opts.SiteDir = d
		}
	}
	if opts.PromptsDir != "" {
		if d, err := filepath.Abs(opts.PromptsDir); err == nil {
			opts.PromptsDir = d
		}
	}

	mux := buildMux(opts)

	addr := fmt.Sprintf(":%d", opts.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen on %s: %w", addr, err)
	}

	actualPort := listener.Addr().(*net.TCPAddr).Port
	url := fmt.Sprintf("http://localhost:%d", actualPort)
	fmt.Printf("🌐 Serving evaluation reports at %s\n", url)
	fmt.Printf("   Reports directory: %s\n", opts.ReportsDir)
	if opts.SiteDir != "" {
		fmt.Printf("   Site directory:    %s (overriding embedded site)\n", opts.SiteDir)
	} else {
		fmt.Printf("   Site:             embedded (use --site-dir to override)\n")
	}
	if opts.DocsDir != "" {
		fmt.Printf("   Docs directory:    %s\n", opts.DocsDir)
	}
	if opts.PromptsDir != "" {
		fmt.Printf("   Prompts directory: %s\n", opts.PromptsDir)
	}
	fmt.Printf("\n")
	fmt.Printf("   ⚠  No authentication. Reports are readable by anyone who\n")
	fmt.Printf("      can reach this port. Do not expose on untrusted networks —\n")
	fmt.Printf("      use an SSH tunnel or authenticating reverse proxy instead.\n")
	fmt.Printf("\n")
	fmt.Printf("   Press Ctrl+C to stop\n\n")

	return http.Serve(listener, corsMiddleware(mux))
}

// buildMux creates the HTTP handler with all routes.
//
// Each call constructs a dedicated in-memory file cache (see cache.go). The
// cache is scoped to the returned mux so tests and concurrent Start calls
// don't share state.
func buildMux(opts Options) *http.ServeMux {
	mux := http.NewServeMux()
	cache := newFileCache()

	// --- API: runs ---
	mux.HandleFunc("/api/runs", func(w http.ResponseWriter, r *http.Request) {
		handleAPIRuns(w, r, opts.ReportsDir, cache)
	})
	mux.HandleFunc("/api/runs/", func(w http.ResponseWriter, r *http.Request) {
		handleAPIRunDetail(w, r, opts.ReportsDir, cache)
	})

	// --- API: docs ---
	if opts.DocsDir != "" {
		mux.HandleFunc("/api/docs", func(w http.ResponseWriter, r *http.Request) {
			handleAPIDocs(w, r, opts.DocsDir)
		})
		mux.HandleFunc("/api/docs/", func(w http.ResponseWriter, r *http.Request) {
			handleAPIDocDetail(w, r, opts.DocsDir)
		})
	}

	// --- API: prompts ---
	if opts.PromptsDir != "" {
		mux.HandleFunc("/api/prompts", func(w http.ResponseWriter, r *http.Request) {
			handleAPIPrompts(w, r, opts.PromptsDir)
		})
	}
	// Always register /api/prompts/ — history works without PromptsDir.
	mux.HandleFunc("/api/prompts/", func(w http.ResponseWriter, r *http.Request) {
		handleAPIPromptDetail(w, r, opts.PromptsDir, opts.ReportsDir)
	})

	// --- Dashboard routes (comparison, trends, drill-down) ---
	registerDashboardRoutes(mux, opts, cache)

	// --- Static file serving for raw report files ---
	reportFS := http.FileServer(http.Dir(opts.ReportsDir))
	mux.Handle("/reports/", http.StripPrefix("/reports/", reportFS))

	// --- SPA fallback / site serving ---
	mux.HandleFunc("/", spaHandler(opts.SiteDir))

	return mux
}

// corsMiddleware adds CORS headers to all responses for dev-server compatibility.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --- Runs API handlers ---

func handleAPIRuns(w http.ResponseWriter, _ *http.Request, reportsDir string, cache *fileCache) {
	runs, err := listRunSummaries(reportsDir, cache)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, runs)
}

func handleAPIRunDetail(w http.ResponseWriter, r *http.Request, reportsDir string, cache *fileCache) {
	// Route: /api/runs/{runId} or /api/runs/{runId}/eval?path=...
	rest := strings.TrimPrefix(r.URL.Path, "/api/runs/")
	if rest == "" {
		http.NotFound(w, r)
		return
	}

	parts := strings.SplitN(rest, "/", 2)
	runID := parts[0]

	// Prevent directory traversal via runID
	if strings.Contains(runID, "..") || strings.Contains(runID, "/") || strings.Contains(runID, string(filepath.Separator)) {
		http.Error(w, "invalid run ID", http.StatusBadRequest)
		return
	}

	// Sub-resource dispatch for /api/runs/{runId}/{sub}
	if len(parts) == 2 {
		switch parts[1] {
		case "eval":
			handleAPIEval(w, r, reportsDir, runID, cache)
		case "graders":
			handleAPIGraders(w, r, reportsDir, runID, cache)
		case "timeline":
			handleAPITimeline(w, r, reportsDir, runID, cache)
		case "score-breakdown":
			handleAPIScoreBreakdown(w, r, reportsDir, runID, cache)
		case "comparisons":
			handleAPIRunComparisons(w, r, reportsDir, runID)
		case "pairwise":
			handleAPIPairwise(w, r, reportsDir, runID, cache)
		default:
			http.NotFound(w, r)
		}
		return
	}

	// /api/runs/{runId} — return full summary.json
	if len(parts) == 1 {
		summaryPath := filepath.Join(reportsDir, runID, "summary.json")
		serveJSONFile(w, r, summaryPath, cache)
		return
	}

	http.NotFound(w, r)
}

// handleAPIPairwise returns the pairwise.json contents for a run, or 404 if
// the run has no pairwise analysis. This makes pairwise results a first-class
// part of the run detail API rather than requiring the site to fish them out
// of summary.json (#360 R140).
func handleAPIPairwise(w http.ResponseWriter, r *http.Request, reportsDir, runID string, cache *fileCache) {
	pairwisePath := filepath.Join(reportsDir, runID, "pairwise.json")
	serveJSONFile(w, r, pairwisePath, cache)
}

func handleAPIEval(w http.ResponseWriter, r *http.Request, reportsDir, runID string, cache *fileCache) {
	relPath := r.URL.Query().Get("path")
	if relPath == "" {
		http.Error(w, `missing "path" query parameter`, http.StatusBadRequest)
		return
	}

	// Prevent directory traversal
	cleaned := filepath.Clean(relPath)
	if strings.Contains(cleaned, "..") {
		http.Error(w, "invalid path", http.StatusBadRequest)
		return
	}

	fullPath := filepath.Join(reportsDir, runID, cleaned)
	serveJSONFile(w, r, fullPath, cache)
}

// --- Docs API handlers ---

func handleAPIDocs(w http.ResponseWriter, _ *http.Request, docsDir string) {
	docs, err := listDocs(docsDir)
	if err != nil {
		slog.Error("listing docs", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, docs)
}

func handleAPIDocDetail(w http.ResponseWriter, r *http.Request, docsDir string) {
	slug := strings.TrimPrefix(r.URL.Path, "/api/docs/")
	if slug == "" || strings.Contains(slug, "/") || strings.Contains(slug, "..") {
		http.NotFound(w, r)
		return
	}

	if internalDocs[slug] {
		http.NotFound(w, r)
		return
	}

	filePath := filepath.Join(docsDir, slug+".md")
	content, err := os.ReadFile(filePath)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	title := extractMarkdownTitle(string(content))
	doc := DocInfo{
		Slug:    slug,
		Title:   title,
		Content: string(content),
	}
	writeJSON(w, doc)
}

// --- Prompts API handlers ---

func handleAPIPrompts(w http.ResponseWriter, _ *http.Request, promptsDir string) {
	prompts, err := prompt.LoadPrompts(promptsDir)
	if err != nil {
		slog.Error("loading prompts", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, prompts)
}

func handleAPIPromptDetail(w http.ResponseWriter, r *http.Request, promptsDir, reportsDir string) {
	slug := strings.TrimPrefix(r.URL.Path, "/api/prompts/")
	if slug == "" || strings.Contains(slug, "..") {
		http.NotFound(w, r)
		return
	}

	// Handle /api/prompts/{promptId}/history
	if strings.HasSuffix(slug, "/history") {
		handleAPIPromptHistory(w, r, reportsDir)
		return
	}

	if promptsDir == "" {
		http.NotFound(w, r)
		return
	}

	prompts, err := prompt.LoadPrompts(promptsDir)
	if err != nil {
		slog.Error("loading prompts", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, p := range prompts {
		if p.ID == slug {
			writeJSON(w, p)
			return
		}
	}

	http.NotFound(w, r)
}

// --- SPA handler ---

func spaHandler(siteDir string) http.HandlerFunc {
	var siteFS fs.FS
	if siteDir != "" {
		siteFS = os.DirFS(siteDir)
	} else {
		// embeddedSite is already the dist/ subtree, use it directly
		siteFS = embeddedSite
	}

	fileServer := http.FileServerFS(siteFS)

	return func(w http.ResponseWriter, r *http.Request) {
		// Strip leading slash for fs.FS operations
		fsPath := strings.TrimPrefix(r.URL.Path, "/")

		// Try to serve the static file directly
		if fsPath != "" {
			if f, err := siteFS.Open(fsPath); err == nil {
				stat, statErr := f.Stat()
				f.Close()
				if statErr == nil && !stat.IsDir() {
					fileServer.ServeHTTP(w, r)
					return
				}
			}
		}

		// SPA fallback: serve index.html for client-side routing
		indexData, err := fs.ReadFile(siteFS, "index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(indexData)
	}
}

// --- Data helpers ---

// listRunSummaries reads run directories and returns their full summary.json
// content. Summaries are served through the file cache so repeated list
// requests don't re-read unchanged summary.json files from disk.
func listRunSummaries(reportsDir string, cache *fileCache) ([]json.RawMessage, error) {
	entries, err := os.ReadDir(reportsDir)
	if err != nil {
		return nil, fmt.Errorf("reading reports dir: %w", err)
	}

	type summaryEntry struct {
		runID string
		data  json.RawMessage
	}

	var items []summaryEntry
	for _, e := range entries {
		if !e.IsDir() || e.Name() == "trends" {
			continue
		}

		summaryPath := filepath.Join(reportsDir, e.Name(), "summary.json")
		data, err := cache.ReadFile(summaryPath)
		if err != nil {
			// Include a minimal entry for runs without summary.json
			minimal, _ := json.Marshal(map[string]string{"run_id": e.Name()})
			items = append(items, summaryEntry{runID: e.Name(), data: minimal})
			continue
		}
		items = append(items, summaryEntry{runID: e.Name(), data: data})
	}

	// Sort newest first
	sort.Slice(items, func(i, j int) bool {
		return items[i].runID > items[j].runID
	})

	result := make([]json.RawMessage, len(items))
	for i, item := range items {
		result[i] = item.data
	}
	return result, nil
}

// listDocs reads the docs directory and returns metadata for each public doc.
func listDocs(docsDir string) ([]DocInfo, error) {
	entries, err := os.ReadDir(docsDir)
	if err != nil {
		return nil, fmt.Errorf("reading docs dir: %w", err)
	}

	var docs []DocInfo
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}

		slug := strings.TrimSuffix(e.Name(), ".md")
		if internalDocs[slug] {
			continue
		}

		content, err := os.ReadFile(filepath.Join(docsDir, e.Name()))
		if err != nil {
			continue
		}

		docs = append(docs, DocInfo{
			Slug:  slug,
			Title: extractMarkdownTitle(string(content)),
		})
	}

	return docs, nil
}

// extractMarkdownTitle returns the text of the first `# ` heading, or the slug.
func extractMarkdownTitle(content string) string {
	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}
	return ""
}

// --- JSON helpers ---

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("encoding JSON response", "error", err)
	}
}

// serveJSONFile reads the named JSON file through the cache and writes it to
// the response, or responds 404 on any read error. Using the cache makes
// repeated reads of unchanged report files (summary.json, report.json, etc.)
// a memory hit rather than a disk + JSON re-parse round-trip.
func serveJSONFile(w http.ResponseWriter, _ *http.Request, path string, cache *fileCache) {
	data, err := cache.ReadFile(path)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
