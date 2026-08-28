$ErrorActionPreference = "Stop"

$repo = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot "..\.."))
$outputDir = Join-Path $repo "reports\azure-skills-three-way-comparison"
$selectedPath = Join-Path $PSScriptRoot "selected-reports.json"

$runs = [ordered]@{
    "python" = "20260827-143238"
    "js-ts" = "20260827-143332"
    "java" = "20260827-143433"
    "dotnet" = "20260827-143539"
}
$languageLabels = [ordered]@{
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
$statusDirs = @(
    "statuses",
    "timeout-statuses",
    "timeout-java-second-statuses",
    "timeout-python-statuses",
    "timeout-python-second-statuses",
    "grading-python-third-statuses",
    "timeout-js-ts-statuses",
    "timeout-js-ts-batch-statuses",
    "timeout-js-ts-second-statuses"
)

function Get-Key($report) {
    "$($report.prompt_id)|$($report.config_name)"
}

function Get-Variant([string]$configName) {
    ($configName -split "/")[-1]
}

function Get-Score($report) {
    $score = [ordered]@{
        prompt_passed = 0
        prompt_total = 0
        language_passed = 0
        language_total = 0
        workspace_passed = 0
        workspace_total = 0
        tool_passed = 0
        tool_total = 0
    }
    foreach ($grader in @($report.grader_results)) {
        foreach ($point in @($grader.points)) {
            $passed = $point.pass -eq $true
            if ($grader.source_type -eq "prompt_file") {
                $score.prompt_total++
                if ($passed) { $score.prompt_passed++ }
            } elseif ($grader.grader_type -eq "workspace") {
                $score.workspace_total++
                if ($passed) { $score.workspace_passed++ }
            } elseif ($grader.grader_type -eq "tool") {
                $score.tool_total++
                if ($passed) { $score.tool_passed++ }
            } elseif ($grader.source_type -eq "criteria_file") {
                $score.language_total++
                if ($passed) { $score.language_passed++ }
            }
        }
    }
    [pscustomobject]$score
}

function Get-Rate([int]$passed, [int]$total) {
    if ($total -eq 0) { return $null }
    [math]::Round(100 * $passed / $total, 1)
}

function Format-Score($aggregate) {
    if ($aggregate.total -eq 0) { return "Not configured" }
    "$($aggregate.passed)/$($aggregate.total) ($('{0:N1}' -f (Get-Rate $aggregate.passed $aggregate.total))%)"
}

function Format-Delta($value) {
    if ($null -eq $value) { return "N/A" }
    if ($value -gt 0) { return "+$('{0:N1}' -f $value) pp" }
    "$('{0:N1}' -f $value) pp"
}

function Add-Candidate([hashtable]$selected, [string]$path, [string]$source, [datetime]$completedAt) {
    if (-not (Test-Path -LiteralPath $path)) { return }
    $report = Get-Content -Raw -LiteralPath $path | ConvertFrom-Json
    $failure = [string]$report.failure_reason
    if ($failure.StartsWith("tool_load_failure:") -or
        $failure -match "context deadline exceeded|SDK evaluation error|session\.idle" -or
        @($report.grader_results).Count -eq 0) {
        return
    }
    $key = Get-Key $report
    if (-not $selected.ContainsKey($key) -or $completedAt -ge $selected[$key].completed_at) {
        $selected[$key] = [pscustomobject]@{
            key = $key
            language = [string]$report.prompt_metadata.language
            prompt_id = [string]$report.prompt_id
            config = [string]$report.config_name
            variant = Get-Variant ([string]$report.config_name)
            report_path = [IO.Path]::GetRelativePath($repo, $path)
            source = $source
            completed_at = $completedAt
            report = $report
            score = Get-Score $report
        }
    }
}

$selected = @{}
foreach ($language in $runs.Keys) {
    $root = Join-Path $repo "reports\$($runs[$language])"
    foreach ($file in Get-ChildItem -LiteralPath $root -Recurse -Filter "report.json" -File -ErrorAction SilentlyContinue) {
        Add-Candidate $selected $file.FullName "original/$($runs[$language])" $file.LastWriteTimeUtc
    }
}

foreach ($statusDirName in $statusDirs) {
    $statusDir = Join-Path $PSScriptRoot $statusDirName
    foreach ($file in Get-ChildItem -LiteralPath $statusDir -Filter "*.json" -File -ErrorAction SilentlyContinue) {
        $status = Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
        if ($status.status -ne "valid" -or -not $status.report_path) { continue }
        $completedAt = if ($status.completed_at) { [datetime]$status.completed_at } else { $file.LastWriteTimeUtc }
        Add-Candidate $selected ([string]$status.report_path) $status.id $completedAt
    }
}

Push-Location $repo
try {
    $promptInventory = ((go run .\hyoka list --json 2>$null | Out-String | ConvertFrom-Json).prompts)
} finally {
    Pop-Location
}

$excluded = [Collections.Generic.List[object]]::new()
$completePrompts = @{}
foreach ($language in $runs.Keys) {
    $completePrompts[$language] = [Collections.Generic.List[string]]::new()
    foreach ($prompt in @($promptInventory | Where-Object { $_.properties.language -eq $language } | Sort-Object id)) {
        $missing = @()
        foreach ($variant in $variantLabels.Keys) {
            $key = "$($prompt.id)|$language-azure-skills/$variant"
            if (-not $selected.ContainsKey($key)) { $missing += "$language-azure-skills/$variant" }
        }
        if ($missing.Count -eq 0) {
            $completePrompts[$language].Add([string]$prompt.id)
        } else {
            $excluded.Add([pscustomobject]@{
                language = $language
                prompt_id = [string]$prompt.id
                missing_configs = $missing
                reason = "Copilot SDK session.idle timeout after controlled retries"
            })
        }
    }
}

New-Item -ItemType Directory -Force -Path $outputDir | Out-Null

$allLanguageResults = [Collections.Generic.List[object]]::new()
foreach ($language in $runs.Keys) {
    $promptIDs = @($completePrompts[$language])
    $variantResults = [ordered]@{}
    foreach ($variant in $variantLabels.Keys) {
        $records = @($promptIDs | ForEach-Object { $selected["$_|$language-azure-skills/$variant"] })
        $promptPassed = ($records | Measure-Object -Property { $_.score.prompt_passed } -Sum).Sum
        $promptTotal = ($records | Measure-Object -Property { $_.score.prompt_total } -Sum).Sum
        $languagePassed = ($records | Measure-Object -Property { $_.score.language_passed } -Sum).Sum
        $languageTotal = ($records | Measure-Object -Property { $_.score.language_total } -Sum).Sum
        $workspacePassed = ($records | Measure-Object -Property { $_.score.workspace_passed } -Sum).Sum
        $workspaceTotal = ($records | Measure-Object -Property { $_.score.workspace_total } -Sum).Sum
        $toolPassed = ($records | Measure-Object -Property { $_.score.tool_passed } -Sum).Sum
        $toolTotal = ($records | Measure-Object -Property { $_.score.tool_total } -Sum).Sum
        $variantResults[$variant] = [pscustomobject]@{
            label = $variantLabels[$variant]
            prompt = [pscustomobject]@{ passed = $promptPassed; total = $promptTotal; rate = Get-Rate $promptPassed $promptTotal }
            language = [pscustomobject]@{ passed = $languagePassed; total = $languageTotal; rate = Get-Rate $languagePassed $languageTotal }
            workspace = [pscustomobject]@{ passed = $workspacePassed; total = $workspaceTotal }
            tool = [pscustomobject]@{ passed = $toolPassed; total = $toolTotal }
        }
    }

    $pairs = @(
        [pscustomobject]@{ id = "azure_vs_baseline"; label = "Azure Skill + MCP vs baseline"; left = "azure-skill-mcp"; right = "baseline" },
        [pscustomobject]@{ id = "microsoft_vs_baseline"; label = "Microsoft skill vs baseline"; left = "azure-skill-mcp-microsoft-skill"; right = "baseline" },
        [pscustomobject]@{ id = "microsoft_vs_azure"; label = "Microsoft skill vs Azure Skill + MCP"; left = "azure-skill-mcp-microsoft-skill"; right = "azure-skill-mcp" }
    )
    $pairResults = [ordered]@{}
    foreach ($pair in $pairs) {
        $improved = 0
        $regressed = 0
        $tied = 0
        foreach ($promptID in $promptIDs) {
            $left = $selected["$promptID|$language-azure-skills/$($pair.left)"].score.prompt_passed
            $right = $selected["$promptID|$language-azure-skills/$($pair.right)"].score.prompt_passed
            if ($left -gt $right) { $improved++ }
            elseif ($left -lt $right) { $regressed++ }
            else { $tied++ }
        }
        $pairResults[$pair.id] = [pscustomobject]@{
            label = $pair.label
            improved = $improved
            regressed = $regressed
            tied = $tied
        }
    }

    $perPrompt = foreach ($promptID in $promptIDs) {
        $arms = [ordered]@{}
        foreach ($variant in $variantLabels.Keys) {
            $item = $selected["$promptID|$language-azure-skills/$variant"]
            $arms[$variant] = [pscustomobject]@{
                prompt_passed = $item.score.prompt_passed
                prompt_total = $item.score.prompt_total
                language_passed = $item.score.language_passed
                language_total = $item.score.language_total
                workspace_passed = $item.score.workspace_passed
                workspace_total = $item.score.workspace_total
                tool_passed = $item.score.tool_passed
                tool_total = $item.score.tool_total
                report_path = $item.report_path
                source = $item.source
            }
        }
        [pscustomobject]@{ prompt_id = $promptID; arms = $arms }
    }

    $languageResult = [pscustomobject]@{
        language = $language
        label = $languageLabels[$language]
        complete_triplets = $promptIDs.Count
        valid_evaluations = $promptIDs.Count * 3
        excluded = @($excluded | Where-Object language -eq $language)
        variants = $variantResults
        pairwise_prompt_outcomes = $pairResults
        prompts = @($perPrompt)
    }
    $allLanguageResults.Add($languageResult)
    $languageResult | ConvertTo-Json -Depth 12 | Set-Content -LiteralPath (Join-Path $outputDir "$language.json") -Encoding utf8

    $baseline = $variantResults["baseline"]
    $azure = $variantResults["azure-skill-mcp"]
    $microsoft = $variantResults["azure-skill-mcp-microsoft-skill"]
    $lines = [Collections.Generic.List[string]]::new()
    $lines.Add("# $($languageLabels[$language]) three-way comparison")
    $lines.Add("")
    $lines.Add("The report includes **$($promptIDs.Count) complete prompt triplets** and **$($promptIDs.Count * 3) valid evaluations**.")
    $lines.Add("")
    $lines.Add("## Prompt checks")
    $lines.Add("")
    $lines.Add("| Arm | Passed | Rate | Difference from baseline |")
    $lines.Add("|---|---:|---:|---:|")
    $lines.Add("| Baseline | $($baseline.prompt.passed)/$($baseline.prompt.total) | $($baseline.prompt.rate)% | - |")
    $lines.Add("| Azure Skill + MCP | $($azure.prompt.passed)/$($azure.prompt.total) | $($azure.prompt.rate)% | $(Format-Delta ($azure.prompt.rate - $baseline.prompt.rate)) |")
    $lines.Add("| Azure Skill + MCP + Microsoft skill | $($microsoft.prompt.passed)/$($microsoft.prompt.total) | $($microsoft.prompt.rate)% | $(Format-Delta ($microsoft.prompt.rate - $baseline.prompt.rate)) |")
    $lines.Add("")
    $lines.Add("| Pairwise prompt outcome | Improved | Regressed | Tied |")
    $lines.Add("|---|---:|---:|---:|")
    foreach ($pair in $pairResults.Values) {
        $lines.Add("| $($pair.label) | $($pair.improved) | $($pair.regressed) | $($pair.tied) |")
    }
    $lines.Add("")
    $lines.Add("Adding Azure Skill + MCP changed the prompt-check rate by **$(Format-Delta ($azure.prompt.rate - $baseline.prompt.rate))**. Adding the Microsoft language skill changed it by **$(Format-Delta ($microsoft.prompt.rate - $azure.prompt.rate))** relative to Azure Skill + MCP.")
    $lines.Add("")
    $lines.Add("## Language checks")
    $lines.Add("")
    if ($baseline.language.total -eq 0) {
        $lines.Add("No generic $($languageLabels[$language]) language criteria are configured.")
    } else {
        $lines.Add("| Arm | Passed | Rate | Difference from baseline |")
        $lines.Add("|---|---:|---:|---:|")
        $lines.Add("| Baseline | $($baseline.language.passed)/$($baseline.language.total) | $($baseline.language.rate)% | - |")
        $lines.Add("| Azure Skill + MCP | $($azure.language.passed)/$($azure.language.total) | $($azure.language.rate)% | $(Format-Delta ($azure.language.rate - $baseline.language.rate)) |")
        $lines.Add("| Azure Skill + MCP + Microsoft skill | $($microsoft.language.passed)/$($microsoft.language.total) | $($microsoft.language.rate)% | $(Format-Delta ($microsoft.language.rate - $baseline.language.rate)) |")
    }
    $lines.Add("")
    $lines.Add("## Excluded diagnostics")
    $lines.Add("")
    $lines.Add("| Diagnostic | Baseline | Azure Skill + MCP | Microsoft skill |")
    $lines.Add("|---|---:|---:|---:|")
    $lines.Add("| Workspace checks | $(Format-Score $baseline.workspace) | $(Format-Score $azure.workspace) | $(Format-Score $microsoft.workspace) |")
    $lines.Add("| Azure MCP usage checks | $(Format-Score $baseline.tool) | $(Format-Score $azure.tool) | $(Format-Score $microsoft.tool) |")
    $lines.Add("")
    $lines.Add("Workspace and tool checks are excluded from scored aggregates because they measure report collection or observed tool invocation rather than generated-code correctness.")
    if (@($languageResult.excluded).Count -gt 0) {
        $lines.Add("")
        $lines.Add("## Execution exclusions")
        $lines.Add("")
        $lines.Add("| Prompt ID | Timed-out config | Attempts | Reason |")
        $lines.Add("|---|---|---:|---|")
        foreach ($item in $languageResult.excluded) {
            foreach ($config in $item.missing_configs) {
                $lines.Add("| ``$($item.prompt_id)`` | ``$config`` | 3 | Copilot SDK ``session.idle`` timeout |")
            }
        }
    }
    $lines.Add("")
    $lines.Add("## Per-prompt prompt checks")
    $lines.Add("")
    $lines.Add("| Prompt ID | Baseline | Azure Skill + MCP | Microsoft skill |")
    $lines.Add("|---|---:|---:|---:|")
    foreach ($prompt in $perPrompt) {
        $lines.Add("| ``$($prompt.prompt_id)`` | $($prompt.arms.baseline.prompt_passed)/$($prompt.arms.baseline.prompt_total) | $($prompt.arms.'azure-skill-mcp'.prompt_passed)/$($prompt.arms.'azure-skill-mcp'.prompt_total) | $($prompt.arms.'azure-skill-mcp-microsoft-skill'.prompt_passed)/$($prompt.arms.'azure-skill-mcp-microsoft-skill'.prompt_total) |")
    }
    $lines -join "`n" | Set-Content -LiteralPath (Join-Path $outputDir "$language.md") -Encoding utf8
}

$rollupVariants = [ordered]@{}
foreach ($variant in $variantLabels.Keys) {
    $passed = ($allLanguageResults | ForEach-Object { $_.variants.$variant.prompt.passed } | Measure-Object -Sum).Sum
    $total = ($allLanguageResults | ForEach-Object { $_.variants.$variant.prompt.total } | Measure-Object -Sum).Sum
    $rollupVariants[$variant] = [pscustomobject]@{ passed = $passed; total = $total; rate = Get-Rate $passed $total }
}
$summary = [pscustomobject]@{
    generated_at = [datetime]::UtcNow.ToString("o")
    complete_triplets = ($allLanguageResults | Measure-Object complete_triplets -Sum).Sum
    valid_evaluations = ($allLanguageResults | Measure-Object valid_evaluations -Sum).Sum
    excluded = @($excluded)
    variants = $rollupVariants
    languages = @($allLanguageResults)
}
$summary | ConvertTo-Json -Depth 14 | Set-Content -LiteralPath (Join-Path $outputDir "summary.json") -Encoding utf8

$summaryLines = [Collections.Generic.List[string]]::new()
$summaryLines.Add("# Azure skills three-way comparison")
$summaryLines.Add("")
$summaryLines.Add("Prompt checks are the primary task-correctness measure. Language checks are supplemental. Workspace and tool/MCP checks are excluded from scored aggregates.")
$summaryLines.Add("")
$summaryLines.Add("## Prompt checks")
$summaryLines.Add("")
$summaryLines.Add("| Language | Complete triplets | Baseline | Azure Skill + MCP | Difference | Microsoft skill | Difference vs baseline | Difference vs Azure Skill + MCP |")
$summaryLines.Add("|---|---:|---:|---:|---:|---:|---:|---:|")
foreach ($result in $allLanguageResults) {
    $b = $result.variants.baseline.prompt
    $a = $result.variants.'azure-skill-mcp'.prompt
    $m = $result.variants.'azure-skill-mcp-microsoft-skill'.prompt
    $summaryLines.Add("| [$($result.label)](./$($result.language).md) | $($result.complete_triplets) | $($b.passed)/$($b.total) ($($b.rate)%) | $($a.passed)/$($a.total) ($($a.rate)%) | $(Format-Delta ($a.rate - $b.rate)) | $($m.passed)/$($m.total) ($($m.rate)%) | $(Format-Delta ($m.rate - $b.rate)) | $(Format-Delta ($m.rate - $a.rate)) |")
}
$rb = $rollupVariants.baseline
$ra = $rollupVariants.'azure-skill-mcp'
$rm = $rollupVariants.'azure-skill-mcp-microsoft-skill'
$summaryLines.Add("| **Informational rollup** | **$($summary.complete_triplets)** | **$($rb.passed)/$($rb.total) ($($rb.rate)%)** | **$($ra.passed)/$($ra.total) ($($ra.rate)%)** | **$(Format-Delta ($ra.rate - $rb.rate))** | **$($rm.passed)/$($rm.total) ($($rm.rate)%)** | **$(Format-Delta ($rm.rate - $rb.rate))** | **$(Format-Delta ($rm.rate - $ra.rate))** |")
$summaryLines.Add("")
$summaryLines.Add("The informational rollup combines equivalent prompt checks only. It is not a language ranking.")
$summaryLines.Add("")
$summaryLines.Add("## Language checks")
$summaryLines.Add("")
$summaryLines.Add("| Language | Baseline | Azure Skill + MCP | Microsoft skill |")
$summaryLines.Add("|---|---:|---:|---:|")
foreach ($result in $allLanguageResults) {
    $summaryLines.Add("| $($result.label) | $(Format-Score $result.variants.baseline.language) | $(Format-Score $result.variants.'azure-skill-mcp'.language) | $(Format-Score $result.variants.'azure-skill-mcp-microsoft-skill'.language) |")
}
$summaryLines.Add("")
$summaryLines.Add("## Interpretation limits")
$summaryLines.Add("")
$summaryLines.Add("- This is a single trial per prompt/config combination and is subject to model variance.")
$summaryLines.Add("- Cross-language scores are not directly comparable because prompt inventories and criteria differ.")
$summaryLines.Add("- Loaded skills might not be invoked for every prompt.")
$summaryLines.Add("- MCP invocation and workspace checks are diagnostics, not generated-code correctness checks.")
if ($excluded.Count -gt 0) {
    $summaryLines.Add("- One JS/TS prompt triplet is excluded because the Microsoft-skill arm repeatedly hit a Copilot SDK ``session.idle`` timeout.")
}
$summaryLines -join "`n" | Set-Content -LiteralPath (Join-Path $outputDir "summary.md") -Encoding utf8

$selection = @($selected.Values | Sort-Object language, prompt_id, config | ForEach-Object {
    [pscustomobject]@{
        language = $_.language
        prompt_id = $_.prompt_id
        config = $_.config
        report_path = $_.report_path
        source = $_.source
        selected = $completePrompts[$_.language] -contains $_.prompt_id
    }
})
$selection | ConvertTo-Json -Depth 5 | Set-Content -LiteralPath $selectedPath -Encoding utf8

Write-Output "selected=$($selected.Count)"
Write-Output "complete_triplets=$($summary.complete_triplets)"
Write-Output "valid_evaluations=$($summary.valid_evaluations)"
Write-Output "excluded=$($excluded.Count)"
Write-Output "output=$outputDir"
