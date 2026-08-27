package eval

import (
	"context"
	"strings"
	"testing"

	copilot "github.com/github/copilot-sdk/go"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/config"
	"github.com/ronniegeraghty/hyoka/hyoka/internal/prompt"
)

func TestBuildSessionConfig_EmptyToolsIsNil(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model:         "gpt-4",
			Tools:         []config.ToolEntry{},
			ExcludedTools: []string{},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.AvailableTools != nil {
		t.Errorf("expected AvailableTools nil (all tools), got %v", sc.AvailableTools)
	}
	if sc.ExcludedTools != nil {
		t.Errorf("expected ExcludedTools nil (no exclusions), got %v", sc.ExcludedTools)
	}
}

func TestBuildSessionConfig_NilAvailableToolsIsNil(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name:      "test",
		Generator: &config.GeneratorConfig{Model: "gpt-4"},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.AvailableTools != nil {
		t.Errorf("expected AvailableTools nil, got %v", sc.AvailableTools)
	}
	if sc.ExcludedTools != nil {
		t.Errorf("expected ExcludedTools nil, got %v", sc.ExcludedTools)
	}
}

func TestBuildSessionConfig_PopulatedAvailableToolsPreserved(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "create"},
				{Name: "edit"},
				{Name: "bash"},
			},
			ExcludedTools: []string{"web_fetch"},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if len(sc.AvailableTools) != 3 {
		t.Errorf("expected 3 AvailableTools, got %d", len(sc.AvailableTools))
	}
	if len(sc.ExcludedTools) != 1 {
		t.Errorf("expected 1 ExcludedTools, got %d", len(sc.ExcludedTools))
	}
}

func TestBuildSessionConfig_WorkingDirectory(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/eval-123", "", nil)
	if sc.WorkingDirectory != "/workspace/eval-123" {
		t.Errorf("expected WorkingDirectory '/workspace/eval-123', got %q", sc.WorkingDirectory)
	}
}

func TestBuildSessionConfig_ConfigDir(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/eval-123", "/isolated/config", nil)
	if sc.ConfigDirectory != "/isolated/config" {
		t.Errorf("expected ConfigDirectory '/isolated/config', got %q", sc.ConfigDirectory)
	}
}

func TestNewCopilotPromptRunnerCLIPath(t *testing.T) {
	runner := NewCopilotPromptRunner(PromptRunnerOptions{CLIPath: "C:\\tools\\copilot.exe"})

	connection, ok := runner.clientOpts.Connection.(copilot.StdioConnection)
	if !ok {
		t.Fatalf("Connection = %T, want copilot.StdioConnection", runner.clientOpts.Connection)
	}
	if connection.Path != "C:\\tools\\copilot.exe" {
		t.Errorf("Connection.Path = %q, want %q", connection.Path, "C:\\tools\\copilot.exe")
	}
}

func TestBuildSessionConfig_PermissionHandler(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{Name: "test", Generator: &config.GeneratorConfig{Model: "gpt-4"}}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.OnPermissionRequest == nil {
		t.Error("expected OnPermissionRequest to be set (ApproveAll)")
	}
}

func TestBuildSessionConfig_MCPServers(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{
					Name:    "azure",
					Type:    "mcp",
					Command: "npx",
					Args:    []string{"-y", "@azure/mcp@latest"},
				},
			},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if len(sc.MCPServers) != 1 {
		t.Errorf("expected 1 MCP server, got %d", len(sc.MCPServers))
	}
	azure, ok := sc.MCPServers["azure"]
	if !ok {
		t.Fatal("expected 'azure' MCP server")
	}
	azureConfig, ok := azure.(copilot.MCPStdioServerConfig)
	if !ok {
		t.Fatalf("expected stdio MCP config, got %T", azure)
	}
	if azureConfig.Command != "npx" {
		t.Errorf("expected MCP command 'npx', got %q", azureConfig.Command)
	}
}

// --- Tool filter resolution tests ---

func TestBuildSessionConfig_ToolEntryResolution(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "create"},
				{Name: "edit"},
				{Name: "azure_mcp", When: map[string]string{"language": "python"}},
			},
		},
	}

	// Python prompt should get all 3 tools
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{"language": "python", "service": "identity"})
	if len(sc.AvailableTools) != 3 {
		t.Fatalf("expected 3 AvailableTools for python, got %d: %v", len(sc.AvailableTools), sc.AvailableTools)
	}

	// Dotnet prompt should get only 2 tools (azure_mcp excluded)
	sc = e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{"language": "dotnet"})
	if len(sc.AvailableTools) != 2 {
		t.Fatalf("expected 2 AvailableTools for dotnet, got %d: %v", len(sc.AvailableTools), sc.AvailableTools)
	}
	for _, tool := range sc.AvailableTools {
		if tool == "azure_mcp" {
			t.Error("azure_mcp should not be included for dotnet")
		}
	}
}

func TestBuildSessionConfig_ToolEntryAllConditionalNoneMatch(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "azure_mcp", When: map[string]string{"language": "python"}},
			},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{"language": "dotnet"})
	if sc.AvailableTools != nil {
		t.Errorf("expected nil AvailableTools when no tools match, got %v", sc.AvailableTools)
	}
}

func TestBuildSessionConfig_ExcludedToolsWithToolEntries(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model:         "gpt-4",
			Tools:         []config.ToolEntry{{Name: "create"}, {Name: "edit"}},
			ExcludedTools: []string{"web_fetch"},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if len(sc.AvailableTools) != 2 {
		t.Errorf("expected 2 AvailableTools, got %d", len(sc.AvailableTools))
	}
	if len(sc.ExcludedTools) != 1 || sc.ExcludedTools[0] != "web_fetch" {
		t.Errorf("expected ExcludedTools [web_fetch], got %v", sc.ExcludedTools)
	}
}

func TestMergePromptProperties(t *testing.T) {
	tests := []struct {
		name string
		p    *prompt.Prompt
		want map[string]string
	}{
		{
			name: "all fields populated",
			p: &prompt.Prompt{
				Properties: map[string]string{
					"service": "identity", "language": "python", "plane": "data-plane",
					"category": "auth", "difficulty": "medium",
				},
			},
			want: map[string]string{
				"service": "identity", "language": "python", "plane": "data-plane",
				"category": "auth", "difficulty": "medium",
			},
		},
		{
			name: "partial fields",
			p:    &prompt.Prompt{Properties: map[string]string{"service": "storage", "language": "dotnet"}},
			want: map[string]string{"service": "storage", "language": "dotnet"},
		},
		{
			name: "empty prompt",
			p:    &prompt.Prompt{},
			want: map[string]string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergePromptProperties(tt.p)
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("key %q: got %q, want %q", k, got[k], v)
				}
			}
			for k := range got {
				if _, ok := tt.want[k]; !ok {
					t.Errorf("unexpected key %q=%q in merged properties", k, got[k])
				}
			}
		})
	}
}

func TestBuildSessionConfig_CustomSystemPrompt(t *testing.T) {
	e := &CopilotPromptRunner{}

	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model:        "gpt-4",
			SystemPrompt: "You are a helpful code generator.",
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{"language": "python"})

	if sc.Model != "gpt-4" {
		t.Errorf("expected model 'gpt-4', got %q", sc.Model)
	}
}

func TestBuildSessionConfig_EmptySystemPrompt(t *testing.T) {
	e := &CopilotPromptRunner{}

	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{})

	if sc.Model != "gpt-4" {
		t.Errorf("expected model 'gpt-4', got %q", sc.Model)
	}
}

// --- Integration tests: YAML config → prompt properties → session config tools ---

func TestIntegration_YAMLConfigToSessionTools(t *testing.T) {
	yamlData := `
configs:
  - name: integration-test
    description: "Integration test config"
    generator:
      model: gpt-4
      tools:
        - name: create
        - name: edit
        - name: azure_mcp
          when:
            language: python
        - name: dotnet_tools
          when:
            language: dotnet
            service: identity
      excluded_tools:
        - web_fetch
`
	cf, err := config.Parse([]byte(yamlData))
	if err != nil {
		t.Fatalf("Parse failed: %v", err)
	}
	if len(cf.Configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(cf.Configs))
	}
	cfg := &cf.Configs[0]

	tests := []struct {
		name         string
		prompt       *prompt.Prompt
		wantAvail    []string
		wantExcluded []string
	}{
		{
			name:         "python prompt gets azure_mcp",
			prompt:       &prompt.Prompt{Properties: map[string]string{"language": "python", "service": "identity"}},
			wantAvail:    []string{"create", "edit", "azure_mcp"},
			wantExcluded: []string{"web_fetch"},
		},
		{
			name:         "dotnet+identity prompt gets dotnet_tools",
			prompt:       &prompt.Prompt{Properties: map[string]string{"language": "dotnet", "service": "identity"}},
			wantAvail:    []string{"create", "edit", "dotnet_tools"},
			wantExcluded: []string{"web_fetch"},
		},
		{
			name:         "dotnet+storage prompt gets only unconditional",
			prompt:       &prompt.Prompt{Properties: map[string]string{"language": "dotnet", "service": "storage"}},
			wantAvail:    []string{"create", "edit"},
			wantExcluded: []string{"web_fetch"},
		},
		{
			name:         "go prompt gets only unconditional",
			prompt:       &prompt.Prompt{Properties: map[string]string{"language": "go", "service": "key-vault"}},
			wantAvail:    []string{"create", "edit"},
			wantExcluded: []string{"web_fetch"},
		},
	}

	e := &CopilotPromptRunner{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			props := mergePromptProperties(tt.prompt)
			sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", props)

			if len(sc.AvailableTools) != len(tt.wantAvail) {
				t.Fatalf("AvailableTools: got %v, want %v", sc.AvailableTools, tt.wantAvail)
			}
			for i, tool := range tt.wantAvail {
				if sc.AvailableTools[i] != tool {
					t.Errorf("AvailableTools[%d] = %q, want %q", i, sc.AvailableTools[i], tool)
				}
			}

			if len(sc.ExcludedTools) != len(tt.wantExcluded) {
				t.Fatalf("ExcludedTools: got %v, want %v", sc.ExcludedTools, tt.wantExcluded)
			}
			for i, tool := range tt.wantExcluded {
				if sc.ExcludedTools[i] != tool {
					t.Errorf("ExcludedTools[%d] = %q, want %q", i, sc.ExcludedTools[i], tool)
				}
			}
		})
	}
}

func TestIntegration_DuplicateToolEntries(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "create"},
				{Name: "create", When: map[string]string{"language": "python"}},
				{Name: "edit"},
			},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", map[string]string{"language": "python"})
	if len(sc.AvailableTools) != 2 {
		t.Fatalf("expected 2 tools after dedup, got %d: %v", len(sc.AvailableTools), sc.AvailableTools)
	}
	if sc.AvailableTools[0] != "create" || sc.AvailableTools[1] != "edit" {
		t.Errorf("expected [create edit], got %v", sc.AvailableTools)
	}
}

func TestBuildSessionConfig_SystemPromptSet(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model:        "gpt-4",
			SystemPrompt: "You are a helpful Azure SDK assistant.",
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.SystemMessage == nil {
		t.Fatal("expected SystemMessage to be set when generator.system_prompt is configured")
	}
	if sc.SystemMessage.Mode != "append" {
		t.Errorf("expected SystemMessage.Mode 'append', got %q", sc.SystemMessage.Mode)
	}
	// With allowCloud=false (default), safety boundaries are appended.
	if !strings.Contains(sc.SystemMessage.Content, "You are a helpful Azure SDK assistant.") {
		t.Errorf("expected SystemMessage.Content to contain system_prompt, got %q", sc.SystemMessage.Content)
	}
	if !strings.Contains(sc.SystemMessage.Content, "SAFETY BOUNDARIES") {
		t.Errorf("expected safety boundaries in system message when allowCloud=false, got %q", sc.SystemMessage.Content)
	}
}

func TestBuildSessionConfig_SystemPromptEmpty(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name:      "test",
		Generator: &config.GeneratorConfig{Model: "gpt-4"},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	// With allowCloud=false (default), safety boundaries are appended even without system_prompt.
	if sc.SystemMessage == nil {
		t.Fatal("expected SystemMessage with safety boundaries when allowCloud=false")
	}
	if !strings.Contains(sc.SystemMessage.Content, "SAFETY BOUNDARIES") {
		t.Errorf("expected safety boundaries when allowCloud=false, got %q", sc.SystemMessage.Content)
	}
}

func TestBuildSessionConfig_AllowCloudSkipsSafetyBoundaries(t *testing.T) {
	e := &CopilotPromptRunner{allowCloud: true}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model:        "gpt-4",
			SystemPrompt: "You are a helpful Azure SDK assistant.",
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.SystemMessage == nil {
		t.Fatal("expected SystemMessage to be set when generator.system_prompt is configured")
	}
	if strings.Contains(sc.SystemMessage.Content, "SAFETY BOUNDARIES") {
		t.Error("expected no safety boundaries when allowCloud=true")
	}
	if sc.SystemMessage.Content != "You are a helpful Azure SDK assistant." {
		t.Errorf("expected only system_prompt when allowCloud=true, got %q", sc.SystemMessage.Content)
	}
}

func TestBuildSessionConfig_AllowCloudNoSystemPrompt(t *testing.T) {
	e := &CopilotPromptRunner{allowCloud: true}
	cfg := &config.ToolConfig{
		Name:      "test",
		Generator: &config.GeneratorConfig{Model: "gpt-4"},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if sc.SystemMessage != nil {
		t.Errorf("expected SystemMessage nil when allowCloud=true and no system_prompt, got %+v", sc.SystemMessage)
	}
}

// ---------------------------------------------------------------------------
// Workspace containment tests (#346)
// ---------------------------------------------------------------------------

func TestExtractAbsPathsFromCommand(t *testing.T) {
	tests := []struct {
		cmd  string
		want int
	}{
		{cmd: "ls -la", want: 0},
		{cmd: "cat /etc/passwd", want: 1},
		{cmd: "cp /home/user/file.txt /tmp/out.txt", want: 2},
		{cmd: "echo hello", want: 0},
		{cmd: "cd /workspace/test && ls", want: 1},
	}
	for _, tt := range tests {
		paths := extractAbsPathsFromCommand(tt.cmd)
		if len(paths) != tt.want {
			t.Errorf("extractAbsPathsFromCommand(%q) returned %d paths, want %d: %v",
				tt.cmd, len(paths), tt.want, paths)
		}
	}
}

func TestExtractAbsPathsFromCommandNormalizesTraversal(t *testing.T) {
	paths := extractAbsPathsFromCommand("cat /workspace/../etc/passwd")
	if len(paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(paths))
	}
	// filepath.Abs normalizes the ../
	if strings.Contains(paths[0], "..") {
		t.Errorf("expected normalized path without .., got %q", paths[0])
	}
}

func TestBuildSessionConfig_RemoteMCP(t *testing.T) {
	e := &CopilotPromptRunner{}
	cfg := &config.ToolConfig{
		Name: "test",
		Generator: &config.GeneratorConfig{
			Model: "gpt-4",
			Tools: []config.ToolEntry{
				{Name: "remote-server", Type: "mcp", MCPType: "remote", URL: "https://mcp.example.com"},
			},
		},
	}
	sc := e.buildSessionConfig(context.Background(), cfg, "/workspace/test", "", nil)
	if len(sc.MCPServers) != 1 {
		t.Fatalf("expected 1 MCP server, got %d", len(sc.MCPServers))
	}
	serverCfg := sc.MCPServers["remote-server"]
	remoteConfig, ok := serverCfg.(copilot.MCPHTTPServerConfig)
	if !ok {
		t.Fatalf("expected HTTP MCP config, got %T", serverCfg)
	}
	if remoteConfig.URL != "https://mcp.example.com" {
		t.Errorf("expected url=https://mcp.example.com, got %q", remoteConfig.URL)
	}
}
