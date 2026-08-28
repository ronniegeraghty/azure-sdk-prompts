$repo = "C:\github\hyoka-azure-skills-three-way-comparison"
$outputPath = Join-Path $PSScriptRoot "recovery-inventory.md"
$manifestPath = Join-Path $PSScriptRoot "rerun-manifest.csv"

$runs = [ordered]@{
    "python" = "20260827-143238"
    "js-ts" = "20260827-143332"
    "java" = "20260827-143433"
    "dotnet" = "20260827-143539"
}

$languageLabels = @{
    "python" = "Python"
    "js-ts" = "JavaScript/TypeScript"
    "java" = "Java"
    "dotnet" = ".NET"
}

$variantLabels = [ordered]@{
    "baseline" = "Baseline"
    "azure-skill-mcp" = "Azure Skill + MCP"
    "azure-skill-mcp-microsoft-skill" = "Azure Skill + MCP + Microsoft skill"
}

Push-Location $repo
try {
    $promptInventoryJson = go run .\hyoka list --json 2>$null | Out-String
    $promptInventory = ($promptInventoryJson | ConvertFrom-Json).prompts
} finally {
    Pop-Location
}

$reports = @{}
foreach ($language in $runs.Keys) {
    $runRoot = Join-Path $repo "reports\$($runs[$language])"
    foreach ($file in Get-ChildItem -LiteralPath $runRoot -Recurse -Filter report.json -File -ErrorAction SilentlyContinue) {
        $report = Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
        $key = "$($report.prompt_id)|$($report.config_name)"

        $promptPassed = 0
        $promptTotal = 0
        $languagePassed = 0
        $languageTotal = 0

        foreach ($grader in @($report.grader_results)) {
            foreach ($point in @($grader.points)) {
                if ($grader.source_type -eq "prompt_file") {
                    $promptTotal++
                    if ($point.pass -eq $true) {
                        $promptPassed++
                    }
                } elseif ($grader.source_type -eq "criteria_file" -and $grader.grader_type -notin @("workspace", "tool")) {
                    $languageTotal++
                    if ($point.pass -eq $true) {
                        $languagePassed++
                    }
                }
            }
        }

        $failureReason = [string]$report.failure_reason
        $invalidReason = $null
        if ($failureReason.StartsWith("tool_load_failure:")) {
            $invalidReason = "tool_load_failure"
        } elseif ($failureReason -match "context deadline exceeded|SDK evaluation error") {
            $invalidReason = "SDK timeout"
        }

        $reports[$key] = [pscustomobject]@{
            Valid = $null -eq $invalidReason
            InvalidReason = $invalidReason
            PromptPassed = $promptPassed
            PromptTotal = $promptTotal
            LanguagePassed = $languagePassed
            LanguageTotal = $languageTotal
        }
    }
}

$lines = [System.Collections.Generic.List[string]]::new()
$reruns = [System.Collections.Generic.List[object]]::new()
$summaryRows = [System.Collections.Generic.List[object]]::new()

$lines.Add("# Three-way evaluation recovery inventory")
$lines.Add("")
$lines.Add('This inventory preserves reports with valid grader outcomes, including ordinary check failures. Reports with `tool_load_failure`, SDK timeouts, or no report are excluded and listed for rerun.')
$lines.Add("")
$lines.Add('Score cells use `P passed/total; L passed/total`, where P is prompt-specific checks and L is language checks. `.NET` has no configured language checks.')
$lines.Add("")
$lines.Add("## Coverage summary")
$lines.Add("")
$lines.Add("| Language | Expected combinations | Valid results | Required reruns |")
$lines.Add("|---|---:|---:|---:|")

$languageSections = [System.Collections.Generic.List[object]]::new()

foreach ($language in $runs.Keys) {
    $prompts = @($promptInventory | Where-Object { $_.properties.language -eq $language } | Sort-Object id)
    $validCount = 0
    $languageRerunCount = 0
    $matrixRows = [System.Collections.Generic.List[string]]::new()

    foreach ($prompt in $prompts) {
        $cells = [System.Collections.Generic.List[string]]::new()
        foreach ($variant in $variantLabels.Keys) {
            $configName = "$language-azure-skills/$variant"
            $key = "$($prompt.id)|$configName"
            $report = $reports[$key]

            if ($null -eq $report) {
                $cells.Add("RERUN (missing)")
                $reruns.Add([pscustomobject]@{
                    Language = $languageLabels[$language]
                    LanguageID = $language
                    PromptID = $prompt.id
                    Variant = $configName
                    Reason = "missing report"
                })
                $languageRerunCount++
            } elseif (-not $report.Valid) {
                $cells.Add("RERUN ($($report.InvalidReason))")
                $reruns.Add([pscustomobject]@{
                    Language = $languageLabels[$language]
                    LanguageID = $language
                    PromptID = $prompt.id
                    Variant = $configName
                    Reason = $report.InvalidReason
                })
                $languageRerunCount++
            } else {
                $languageScore = if ($report.LanguageTotal -gt 0) {
                    "$($report.LanguagePassed)/$($report.LanguageTotal)"
                } else {
                    "n/a"
                }
                $cells.Add("P $($report.PromptPassed)/$($report.PromptTotal); L $languageScore")
                $validCount++
            }
        }

        $matrixRows.Add(('| `{0}` | {1} | {2} | {3} |' -f $prompt.id, $cells[0], $cells[1], $cells[2]))
    }

    $expectedCount = $prompts.Count * $variantLabels.Count
    $summaryRows.Add([pscustomobject]@{
        Language = $languageLabels[$language]
        Expected = $expectedCount
        Valid = $validCount
        Reruns = $languageRerunCount
    })
    $languageSections.Add([pscustomobject]@{
        Language = $languageLabels[$language]
        RunID = $runs[$language]
        Rows = $matrixRows
    })
}

foreach ($row in $summaryRows) {
    $lines.Add("| $($row.Language) | $($row.Expected) | $($row.Valid) | $($row.Reruns) |")
}

$totalExpected = ($summaryRows | Measure-Object Expected -Sum).Sum
$totalValid = ($summaryRows | Measure-Object Valid -Sum).Sum
$totalReruns = ($summaryRows | Measure-Object Reruns -Sum).Sum
$lines.Add("| **Total** | **$totalExpected** | **$totalValid** | **$totalReruns** |")

foreach ($section in $languageSections) {
    $lines.Add("")
    $lines.Add("## $($section.Language) valid-data matrix")
    $lines.Add("")
    $lines.Add(('Run: `{0}`' -f $section.RunID))
    $lines.Add("")
    $lines.Add("| Prompt ID | Baseline | Azure Skill + MCP | Azure Skill + MCP + Microsoft skill |")
    $lines.Add("|---|---|---|---|")
    foreach ($row in $section.Rows) {
        $lines.Add($row)
    }
}

$lines.Add("")
$lines.Add("## Required reruns")
$lines.Add("")
$lines.Add("| Language | Prompt ID | Variant | Reason |")
$lines.Add("|---|---|---|---|")
foreach ($rerun in $reruns | Sort-Object Language, PromptID, Variant) {
    $lines.Add(('| {0} | `{1}` | `{2}` | {3} |' -f $rerun.Language, $rerun.PromptID, $rerun.Variant, $rerun.Reason))
}

$lines.Add("")
$lines.Add("Total required reruns: **$totalReruns**.")

$lines | Set-Content -LiteralPath $outputPath -Encoding utf8

$laneByLanguage = @{
    "java" = "a"
    "js-ts" = "a"
    "dotnet" = "b"
    "python" = "b"
}
$phaseByLanguage = @{
    "java" = 1
    "js-ts" = 2
    "dotnet" = 1
    "python" = 2
}
$laneCounters = @{
    "a" = 0
    "b" = 0
}

$manifest = foreach ($rerun in $reruns | Sort-Object `
    @{ Expression = { $laneByLanguage[$_.LanguageID] } }, `
    @{ Expression = { $phaseByLanguage[$_.LanguageID] } }, `
    PromptID, Variant) {
    $lane = $laneByLanguage[$rerun.LanguageID]
    $laneCounters[$lane]++
    $laneOrder = $laneCounters[$lane]

    [pscustomobject]@{
        id = "lane-$lane-$('{0:d3}' -f $laneOrder)"
        lane = $lane
        lane_order = $laneOrder
        language = $rerun.LanguageID
        prompt_id = $rerun.PromptID
        config = $rerun.Variant
        reason = $rerun.Reason
        output_folder = "rerun-reports\lane-$lane\$('{0:d3}' -f $laneOrder)"
    }
}

$manifest | Export-Csv -LiteralPath $manifestPath -NoTypeInformation -Encoding utf8

Write-Output "output=$outputPath"
Write-Output "manifest=$manifestPath"
Write-Output "expected=$totalExpected"
Write-Output "valid=$totalValid"
Write-Output "reruns=$totalReruns"
foreach ($row in $summaryRows) {
    Write-Output "$($row.Language): valid=$($row.Valid), reruns=$($row.Reruns)"
}
