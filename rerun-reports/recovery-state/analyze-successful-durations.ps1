$ErrorActionPreference = "Stop"
$repo = "C:\github\hyoka-azure-skills-three-way-comparison"
$stateRoot = "C:\Users\weidxu\.copilot\session-state\21a834dd-17c4-416d-bbad-ebecd7fb9297\files\three-way-runs"
$runs = @(
    "20260827-143238",
    "20260827-143332",
    "20260827-143433",
    "20260827-143539"
)

$records = @{}

foreach ($run in $runs) {
    $runPath = Join-Path $repo "reports\$run"
    foreach ($file in Get-ChildItem -LiteralPath $runPath -Recurse -Filter report.json -File -ErrorAction SilentlyContinue) {
        $report = Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
        $failure = [string]$report.failure_reason
        if ($failure.StartsWith("tool_load_failure:") -or $failure -match "context deadline exceeded|SDK evaluation error|session\.idle") {
            continue
        }

        $records["$($report.prompt_id)|$($report.config_name)"] = [pscustomobject]@{
            Prompt = $report.prompt_id
            Config = $report.config_name
            Language = $report.prompt_metadata.language
            Duration = [double]$report.duration_seconds
            Generation = [double]$report.generation_duration_seconds
            Review = [double]$report.review_duration_seconds
        }
    }
}

$statusDir = Join-Path $stateRoot "statuses"
foreach ($file in Get-ChildItem -LiteralPath $statusDir -Filter "lane-*.json" -File) {
    $status = Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
    if ($status.status -ne "valid" -or -not (Test-Path -LiteralPath $status.report_path)) {
        continue
    }

    $report = Get-Content -Raw -LiteralPath $status.report_path | ConvertFrom-Json
    $records["$($report.prompt_id)|$($report.config_name)"] = [pscustomobject]@{
        Prompt = $report.prompt_id
        Config = $report.config_name
        Language = $report.prompt_metadata.language
        Duration = [double]$report.duration_seconds
        Generation = [double]$report.generation_duration_seconds
        Review = [double]$report.review_duration_seconds
    }
}

$suffixes = @(
    "baseline",
    "azure-skill-mcp",
    "azure-skill-mcp-microsoft-skill"
)
$complete = @()

foreach ($group in $records.Values | Group-Object Prompt) {
    $bySuffix = @{}
    foreach ($record in $group.Group) {
        $bySuffix[($record.Config -split "/")[-1]] = $record
    }
    if (@($suffixes | Where-Object { -not $bySuffix.ContainsKey($_) }).Count -ne 0) {
        continue
    }

    $complete += [pscustomobject]@{
        Language = $group.Group[0].Language
        Prompt = $group.Name
        Baseline = $bySuffix["baseline"]
        MCP = $bySuffix["azure-skill-mcp"]
        Microsoft = $bySuffix["azure-skill-mcp-microsoft-skill"]
    }
}

Write-Output "COMPLETE_TRIPLETS=$($complete.Count)"
Write-Output "SAMPLE"

$sample = $complete |
    Sort-Object Language, Prompt |
    Group-Object Language |
    ForEach-Object { $_.Group | Select-Object -First 2 }

$sample | ForEach-Object {
    [pscustomobject]@{
        Language = $_.Language
        Prompt = $_.Prompt
        BaselineMin = [math]::Round($_.Baseline.Duration / 60, 1)
        MCPMin = [math]::Round($_.MCP.Duration / 60, 1)
        MicrosoftMin = [math]::Round($_.Microsoft.Duration / 60, 1)
        MCPGenerationMin = [math]::Round($_.MCP.Generation / 60, 1)
        MCPReviewMin = [math]::Round($_.MCP.Review / 60, 1)
        MicrosoftGenerationMin = [math]::Round($_.Microsoft.Generation / 60, 1)
        MicrosoftReviewMin = [math]::Round($_.Microsoft.Review / 60, 1)
    }
} | Format-Table -AutoSize

Write-Output "ARM_SUMMARY"
$armDefinitions = @(
    @{ Name = "Baseline"; Property = "Baseline" },
    @{ Name = "Azure Skill + MCP"; Property = "MCP" },
    @{ Name = "Microsoft skill"; Property = "Microsoft" }
)

& {
    foreach ($arm in $armDefinitions) {
        $values = @($complete | ForEach-Object { $_.($arm.Property).Duration } | Sort-Object)
        if ($values.Count -eq 0) {
            continue
        }

        $middle = [int][math]::Floor($values.Count / 2)
        $median = if ($values.Count % 2) {
            $values[$middle]
        } else {
            ($values[$middle - 1] + $values[$middle]) / 2
        }
        $p90Index = [math]::Min($values.Count - 1, [int][math]::Floor($values.Count * 0.9))

        [pscustomobject]@{
            Arm = $arm.Name
            Samples = $values.Count
            MedianMin = [math]::Round($median / 60, 1)
            P90Min = [math]::Round($values[$p90Index] / 60, 1)
            MaxMin = [math]::Round($values[-1] / 60, 1)
        }
    }
} | Format-Table -AutoSize

Write-Output "LONGEST_SUCCESSFUL"
$complete |
    ForEach-Object { @($_.Baseline, $_.MCP, $_.Microsoft) } |
    Sort-Object Duration -Descending |
    Select-Object -First 8 `
        Language,
        Prompt,
        @{ Name = "Arm"; Expression = { ($_.Config -split "/")[-1] } },
        @{ Name = "TotalMin"; Expression = { [math]::Round($_.Duration / 60, 1) } },
        @{ Name = "GenerationMin"; Expression = { [math]::Round($_.Generation / 60, 1) } },
        @{ Name = "ReviewMin"; Expression = { [math]::Round($_.Review / 60, 1) } } |
    Format-Table -AutoSize
