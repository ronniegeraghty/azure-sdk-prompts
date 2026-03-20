package main

import (
"context"
"fmt"
"os"
"strings"
"time"

"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/config"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/eval"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/manifest"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/prompt"
"github.com/ronniegeraghty/azure-sdk-prompts/tool/internal/validate"
"github.com/spf13/cobra"
)

var version = "0.1.0"

func main() {
if err := rootCmd().Execute(); err != nil {
os.Exit(1)
}
}

func rootCmd() *cobra.Command {
root := &cobra.Command{
Use:   "sdk-eval",
Short: "SDK Evaluation Tool — test AI agent code generation quality",
Long:  "A tool for evaluating AI agent code generation quality by running prompts through the Copilot SDK, verifying builds, and generating reports.",
}

root.AddCommand(runCmd())
root.AddCommand(listCmd())
root.AddCommand(configsCmd())
root.AddCommand(versionCmd())
root.AddCommand(manifestCmd())
root.AddCommand(validateCmd())

return root
}

// resolvePathFlag returns the flag value if explicitly set by the user,
// otherwise tries the candidate paths in order, falling back to the default.
func resolvePathFlag(cmd *cobra.Command, flagName string, candidates []string) string {
if cmd.Flags().Changed(flagName) {
val, _ := cmd.Flags().GetString(flagName)
return val
}
for _, c := range candidates {
if info, err := os.Stat(c); err == nil && info.IsDir() || err == nil {
return c
}
}
val, _ := cmd.Flags().GetString(flagName)
return val
}

// resolvePromptsDir resolves the --prompts flag with auto-detection.
func resolvePromptsDir(cmd *cobra.Command) string {
return resolvePathFlag(cmd, "prompts", []string{"./prompts", "../prompts"})
}

// resolveConfigFile resolves the --config-file flag with auto-detection.
func resolveConfigFile(cmd *cobra.Command) string {
return resolvePathFlag(cmd, "config-file", []string{
"./configs/all.yaml", "../configs/all.yaml",
"./configs.yaml", "../configs.yaml",
})
}

// resolveOutputFile resolves the --output flag with auto-detection for file paths.
func resolveOutputFile(cmd *cobra.Command, candidates []string) string {
return resolvePathFlag(cmd, "output", candidates)
}

type runFlags struct {
prompts    string
service    string
language   string
plane      string
category   string
tags       string
promptID   string
configName string
configFile string
workers    int
timeout    int
model      string
output     string
skipTests  bool
skipReview bool
debug      bool
dryRun     bool
}

func addFilterFlags(cmd *cobra.Command, f *runFlags) {
cmd.Flags().StringVar(&f.prompts, "prompts", "./prompts", "Path to prompt library directory")
cmd.Flags().StringVar(&f.service, "service", "", "Filter by Azure service")
cmd.Flags().StringVar(&f.language, "language", "", "Filter by programming language")
cmd.Flags().StringVar(&f.plane, "plane", "", "Filter by data-plane/management-plane")
cmd.Flags().StringVar(&f.category, "category", "", "Filter by use-case category")
cmd.Flags().StringVar(&f.tags, "tags", "", "Filter by tags (comma-separated)")
cmd.Flags().StringVar(&f.promptID, "prompt-id", "", "Run a single prompt by ID")
cmd.Flags().StringVar(&f.configName, "config", "", "Config name(s) from config file (comma-separated)")
cmd.Flags().StringVar(&f.configFile, "config-file", "./configs/all.yaml", "Path to configuration YAML")
cmd.Flags().IntVar(&f.workers, "workers", 4, "Parallel workers")
cmd.Flags().IntVar(&f.timeout, "timeout", 300, "Per-prompt timeout in seconds")
cmd.Flags().StringVar(&f.model, "model", "", "Override model for all configs")
cmd.Flags().StringVar(&f.output, "output", "./reports", "Report output directory")
cmd.Flags().BoolVar(&f.skipTests, "skip-tests", false, "Skip test generation")
cmd.Flags().BoolVar(&f.skipReview, "skip-review", false, "Skip code review")
cmd.Flags().BoolVar(&f.debug, "debug", false, "Verbose output")
cmd.Flags().BoolVar(&f.dryRun, "dry-run", false, "List matching prompts without running")
}

func buildFilter(f *runFlags) prompt.Filter {
var tags []string
if f.tags != "" {
tags = strings.Split(f.tags, ",")
for i := range tags {
tags[i] = strings.TrimSpace(tags[i])
}
}
return prompt.Filter{
Service:  f.service,
Plane:    f.plane,
Language: f.language,
Category: f.category,
Tags:     tags,
PromptID: f.promptID,
}
}

func runCmd() *cobra.Command {
f := &runFlags{}
cmd := &cobra.Command{
Use:   "run",
Short: "Run evaluations",
Long:  "Run evaluations with optional filters against the prompt library.",
RunE: func(cmd *cobra.Command, args []string) error {
f.prompts = resolvePromptsDir(cmd)
f.configFile = resolveConfigFile(cmd)

// Load config
cfgFile, err := config.Load(f.configFile)
if err != nil {
return fmt.Errorf("loading config: %w", err)
}

// Get selected configs
var configNames []string
if f.configName != "" {
configNames = strings.Split(f.configName, ",")
for i := range configNames {
configNames[i] = strings.TrimSpace(configNames[i])
}
}
configs, err := cfgFile.GetConfigs(configNames)
if err != nil {
return err
}

// Override model if specified
if f.model != "" {
for i := range configs {
configs[i].Model = f.model
}
}

// Load and filter prompts
prompts, err := prompt.LoadPrompts(f.prompts)
if err != nil {
return fmt.Errorf("loading prompts: %w", err)
}

filter := buildFilter(f)
filtered := prompt.FilterPrompts(prompts, filter)

if len(filtered) == 0 {
fmt.Println("No prompts matched the given filters.")
return nil
}

fmt.Printf("Found %d prompt(s), %d config(s) → %d evaluation(s)\n",
len(filtered), len(configs), len(filtered)*len(configs))

// Create and run engine
engine := eval.NewEngine(&eval.StubEvaluator{}, eval.EngineOptions{
Workers:    f.workers,
Timeout:    time.Duration(f.timeout) * time.Second,
OutputDir:  f.output,
SkipTests:  f.skipTests,
SkipReview: f.skipReview,
Debug:      f.debug,
DryRun:     f.dryRun,
})

summary, err := engine.Run(context.Background(), filtered, configs)
if err != nil {
return fmt.Errorf("evaluation failed: %w", err)
}

fmt.Printf("\nRun Summary:\n")
fmt.Printf("  Run ID:      %s\n", summary.RunID)
fmt.Printf("  Evaluations: %d\n", summary.TotalEvals)
fmt.Printf("  Passed:      %d\n", summary.Passed)
fmt.Printf("  Failed:      %d\n", summary.Failed)
fmt.Printf("  Errors:      %d\n", summary.Errors)
fmt.Printf("  Duration:    %.2fs\n", summary.Duration)

return nil
},
}

addFilterFlags(cmd, f)
return cmd
}

func listCmd() *cobra.Command {
f := &runFlags{}
cmd := &cobra.Command{
Use:   "list",
Short: "List matching prompts",
Long:  "List prompts matching the given filters (dry-run equivalent).",
RunE: func(cmd *cobra.Command, args []string) error {
f.prompts = resolvePromptsDir(cmd)

prompts, err := prompt.LoadPrompts(f.prompts)
if err != nil {
return fmt.Errorf("loading prompts: %w", err)
}

filter := buildFilter(f)
filtered := prompt.FilterPrompts(prompts, filter)

if len(filtered) == 0 {
fmt.Println("No prompts matched the given filters.")
return nil
}

fmt.Printf("Found %d prompt(s):\n\n", len(filtered))
for _, p := range filtered {
fmt.Printf("  %-30s %s/%s/%s [%s]\n", p.ID, p.Service, p.Plane, p.Language, p.Category)
if p.Description != "" {
fmt.Printf("  %-30s %s\n", "", p.Description)
}
}
return nil
},
}

addFilterFlags(cmd, f)
return cmd
}

func configsCmd() *cobra.Command {
var configFile string

cmd := &cobra.Command{
Use:   "configs",
Short: "List available configurations",
RunE: func(cmd *cobra.Command, args []string) error {
configFile = resolveConfigFile(cmd)

cfgFile, err := config.Load(configFile)
if err != nil {
return fmt.Errorf("loading config: %w", err)
}

fmt.Printf("Available configurations (%d):\n\n", len(cfgFile.Configs))
for _, c := range cfgFile.Configs {
fmt.Printf("  %-20s %s (model: %s)\n", c.Name, c.Description, c.Model)
if len(c.MCPServers) > 0 {
fmt.Printf("  %-20s MCP servers: ", "")
var names []string
for name := range c.MCPServers {
names = append(names, name)
}
fmt.Println(strings.Join(names, ", "))
}
}
return nil
},
}

cmd.Flags().StringVar(&configFile, "config-file", "./configs/all.yaml", "Path to configuration YAML")
return cmd
}

func manifestCmd() *cobra.Command {
var promptsDir string
var outputPath string

cmd := &cobra.Command{
Use:   "manifest",
Short: "Regenerate manifest.yaml from prompt files",
Long:  "Scan the prompts directory and generate a manifest.yaml with metadata about all prompts.",
RunE: func(cmd *cobra.Command, args []string) error {
promptsDir = resolvePromptsDir(cmd)
outputPath = resolveOutputFile(cmd, []string{"./manifest.yaml", "../manifest.yaml"})

m, err := manifest.Generate(promptsDir)
if err != nil {
return fmt.Errorf("generating manifest: %w", err)
}

data, err := m.Marshal()
if err != nil {
return fmt.Errorf("marshaling manifest: %w", err)
}

if err := os.WriteFile(outputPath, data, 0644); err != nil {
return fmt.Errorf("writing manifest: %w", err)
}

fmt.Printf("Generated %s with %d prompt(s)\n", outputPath, m.PromptCount)
return nil
},
}

cmd.Flags().StringVar(&promptsDir, "prompts", "./prompts", "Path to prompt library directory")
cmd.Flags().StringVar(&outputPath, "output", "./manifest.yaml", "Output path for manifest.yaml")
return cmd
}

func validateCmd() *cobra.Command {
var promptsDir string

cmd := &cobra.Command{
Use:   "validate",
Short: "Validate prompt frontmatter",
Long:  "Validate all prompt files against schema rules and naming conventions.",
RunE: func(cmd *cobra.Command, args []string) error {
promptsDir = resolvePromptsDir(cmd)

result, err := validate.Validate(promptsDir)
if err != nil {
return fmt.Errorf("validation: %w", err)
}

fmt.Print(validate.FormatResult(result))

if !result.OK() {
os.Exit(1)
}
return nil
},
}

cmd.Flags().StringVar(&promptsDir, "prompts", "./prompts", "Path to prompt library directory")
return cmd
}

func versionCmd() *cobra.Command {
return &cobra.Command{
Use:   "version",
Short: "Print version",
Run: func(cmd *cobra.Command, args []string) {
fmt.Printf("sdk-eval version %s\n", version)
},
}
}
