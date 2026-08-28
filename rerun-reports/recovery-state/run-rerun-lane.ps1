param(
    [Parameter(Mandatory = $true)]
    [ValidateSet("a", "b")]
    [string]$Lane,

    [int]$StartDelaySeconds = 0,

    [string]$HyokaPath = (Join-Path $env:TEMP "hyoka-pr-656-rerun.exe"),

    [string]$ManifestPath = "",

    [string]$StatusDir = "",

    [string]$LogDir = ""
)

$ErrorActionPreference = "Stop"
$repo = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot "..\.."))
$stateRoot = $PSScriptRoot
if (-not $ManifestPath) {
    $ManifestPath = Join-Path $stateRoot "rerun-manifest.csv"
}
if (-not $StatusDir) {
    $StatusDir = Join-Path $stateRoot "statuses"
}
if (-not $LogDir) {
    $LogDir = Join-Path $stateRoot "logs\lane-$Lane"
}

New-Item -ItemType Directory -Force -Path $StatusDir, $LogDir | Out-Null

function Write-RunStatus {
    param(
        [Parameter(Mandatory = $true)]
        [hashtable]$Data
    )

    $path = Join-Path $StatusDir "$($Data.id).json"
    $tempPath = "$path.tmp"
    $Data | ConvertTo-Json -Depth 8 | Set-Content -LiteralPath $tempPath -Encoding utf8
    Move-Item -LiteralPath $tempPath -Destination $path -Force
}

function Get-SanityProblems {
    param(
        [Parameter(Mandatory = $true)]
        [pscustomobject]$ManifestRow,

        [Parameter(Mandatory = $true)]
        [int]$ExitCode,

        [Parameter(Mandatory = $true)]
        [AllowEmptyCollection()]
        [System.IO.FileInfo[]]$ReportFiles
    )

    $problems = [System.Collections.Generic.List[string]]::new()
    if ($ExitCode -ne 0) {
        $problems.Add("command_exit_$ExitCode")
    }
    if ($ReportFiles.Count -eq 0) {
        $problems.Add("missing_report")
        return [pscustomobject]@{
            Problems = @($problems)
            Report = $null
            ReportPath = $null
        }
    }
    if ($ReportFiles.Count -ne 1) {
        $problems.Add("multiple_reports_$($ReportFiles.Count)")
    }

    $reportFile = $ReportFiles | Sort-Object LastWriteTimeUtc -Descending | Select-Object -First 1
    try {
        $report = Get-Content -Raw -LiteralPath $reportFile.FullName | ConvertFrom-Json
    } catch {
        $problems.Add("invalid_report_json")
        return [pscustomobject]@{
            Problems = @($problems)
            Report = $null
            ReportPath = $reportFile.FullName
        }
    }

    if ([string]$report.prompt_id -ne [string]$ManifestRow.prompt_id) {
        $problems.Add("prompt_mismatch")
    }
    if ([string]$report.config_name -ne [string]$ManifestRow.config) {
        $problems.Add("config_mismatch")
    }

    $failureReason = [string]$report.failure_reason
    if ($failureReason.StartsWith("tool_load_failure:")) {
        $problems.Add("tool_load_failure")
    }
    if ($failureReason -match "context deadline exceeded|SDK evaluation error|session\.idle") {
        $problems.Add("sdk_timeout")
    }
    if (@($report.grader_results).Count -eq 0) {
        $problems.Add("missing_graders")
    }

    foreach ($tool in @($report.session_setup.skills) + @($report.session_setup.mcp_servers)) {
        if ($null -ne $tool -and $null -ne $tool.status -and [string]$tool.status -ne "loaded") {
            $kind = if ($null -ne $tool.kind) { [string]$tool.kind } else { "tool" }
            $name = if ($null -ne $tool.name) { [string]$tool.name } else { "unknown" }
            $problems.Add("setup_${kind}_${name}_$([string]$tool.status)")
        }
    }

    return [pscustomobject]@{
        Problems = @($problems | Sort-Object -Unique)
        Report = $report
        ReportPath = $reportFile.FullName
    }
}

if ($StartDelaySeconds -gt 0) {
    Start-Sleep -Seconds $StartDelaySeconds
}

if (-not (Test-Path -LiteralPath $HyokaPath)) {
    Push-Location $repo
    try {
        go build -o $HyokaPath .\hyoka
        if ($LASTEXITCODE -ne 0) {
            throw "Failed to build Hyoka executable"
        }
    } finally {
        Pop-Location
    }
}

$env:npm_config_registry = "https://pkgs.dev.azure.com/azure-sdk/public/_packaging/azure-sdk-for-js/npm/registry/"
$env:NPM_CONFIG_REGISTRY = $env:npm_config_registry

$rows = @(Import-Csv -LiteralPath $ManifestPath |
    Where-Object lane -eq $Lane |
    Sort-Object { [int]$_.lane_order })

Set-Location $repo

foreach ($row in $rows) {
    $statusPath = Join-Path $StatusDir "$($row.id).json"
    if (Test-Path -LiteralPath $statusPath) {
        $existing = Get-Content -Raw -LiteralPath $statusPath | ConvertFrom-Json
        if ($existing.status -in @("valid", "problem")) {
            continue
        }
    }

    $startedAt = [DateTime]::UtcNow
    $consoleLog = Join-Path $LogDir "$($row.id)-console.log"
    $diagnosticLog = Join-Path $LogDir "$($row.id)-hyoka.log"
    $outputPath = Join-Path $repo $row.output_folder

    Write-RunStatus @{
        id = $row.id
        lane = $row.lane
        lane_order = [int]$row.lane_order
        language = $row.language
        prompt_id = $row.prompt_id
        config = $row.config
        output_folder = $row.output_folder
        status = "running"
        started_at = $startedAt.ToString("o")
        completed_at = $null
        exit_code = $null
        problems = @()
        report_path = $null
    }

    $exitCode = -1
    $sanity = $null
    try {
        & $HyokaPath run `
            --prompt-id $row.prompt_id `
            --config $row.config `
            --criteria-dir ".\criteria" `
            --workers 1 `
            --output $row.output_folder `
            --progress off `
            --log-level info `
            --log-file $diagnosticLog `
            --yes *> $consoleLog
        $exitCode = $LASTEXITCODE

        $reportFiles = @(Get-ChildItem -LiteralPath $outputPath -Recurse -Filter report.json -File -ErrorAction SilentlyContinue)
        $sanity = Get-SanityProblems -ManifestRow $row -ExitCode $exitCode -ReportFiles $reportFiles
    } catch {
        $sanity = [pscustomobject]@{
            Problems = @("runner_exception", $_.Exception.Message)
            Report = $null
            ReportPath = $null
        }
    }

    $completedAt = [DateTime]::UtcNow
    $status = if (@($sanity.Problems).Count -eq 0) { "valid" } else { "problem" }
    $failureReason = if ($null -ne $sanity.Report) { [string]$sanity.Report.failure_reason } else { $null }

    Write-RunStatus @{
        id = $row.id
        lane = $row.lane
        lane_order = [int]$row.lane_order
        language = $row.language
        prompt_id = $row.prompt_id
        config = $row.config
        output_folder = $row.output_folder
        status = $status
        started_at = $startedAt.ToString("o")
        completed_at = $completedAt.ToString("o")
        duration_seconds = [math]::Round(($completedAt - $startedAt).TotalSeconds, 1)
        exit_code = $exitCode
        problems = @($sanity.Problems)
        failure_reason = $failureReason
        report_path = $sanity.ReportPath
    }
}
