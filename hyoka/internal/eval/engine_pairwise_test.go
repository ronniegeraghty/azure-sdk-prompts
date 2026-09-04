package eval

import (
"testing"
)

func TestExtractPairwiseVariant(t *testing.T) {
tests := []struct {
name       string
configName string
want       string
}{
{
name:       "baseline variant",
configName: "python-pairwise/baseline/claude-opus-4.6",
want:       "baseline",
},
{
name:       "simple ablation (without-azure)",
configName: "python-pairwise/without-azure/claude-opus-4.6",
want:       "without-azure",
},
{
name:       "deep MCP variant",
configName: "python-pairwise/without-azure/storage_blob_list/claude-opus-4.6",
want:       "without-azure/storage_blob_list",
},
{
name:       "baseline without model suffix",
configName: "python-pairwise/baseline",
want:       "baseline",
},
{
name:       "ablation without model suffix",
configName: "python-pairwise/without-cosmos",
want:       "without-cosmos",
},
{
name:       "non-pairwise config",
configName: "baseline",
want:       "",
},
{
name:       "multi-model non-pairwise",
configName: "python-base/claude-opus-4.6",
want:       "",
},
{
name:       "empty string",
configName: "",
want:       "",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
got := extractPairwiseVariant(tt.configName)
if got != tt.want {
t.Errorf("extractPairwiseVariant(%q) = %q, want %q", tt.configName, got, tt.want)
}
})
}
}
