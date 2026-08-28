$ErrorActionPreference = "Stop"

$repo = [IO.Path]::GetFullPath((Join-Path $PSScriptRoot "..\.."))
$manifest = @(Import-Csv -LiteralPath (Join-Path $PSScriptRoot "rerun-manifest.csv"))
$statusDir = Join-Path $PSScriptRoot "statuses"
$target = Join-Path $repo "rerun-reports\run-ledger.md"

$lines = [Collections.Generic.List[string]]::new()
$lines.Add("# Three-way rerun ledger")
$lines.Add("")
$lines.Add("Recovered from the interrupted Copilot CLI session. Pre-rerun reports are backed up in ``..\hyoka-azure-skills-three-way-comparison-report-backup-20260827-210750\``.")
$lines.Add("")
$lines.Add("- Expected evaluations: **216**")
$lines.Add("- Retained valid results from original runs: **65**")
$lines.Add("- Required reruns: **151**")
$lines.Add("")
$lines.Add("| ID | Language | Prompt ID | Variant | Original reason | Status | Output folder | Report |")
$lines.Add("|---|---|---|---|---|---|---|---|")

foreach ($row in $manifest) {
    $status = "pending"
    $report = ""
    $statusPath = Join-Path $statusDir "$($row.id).json"
    if (Test-Path -LiteralPath $statusPath) {
        try {
            $runStatus = Get-Content -Raw -LiteralPath $statusPath | ConvertFrom-Json
            $status = [string]$runStatus.status
            if ($runStatus.report_path) {
                $report = [IO.Path]::GetRelativePath($repo, [string]$runStatus.report_path)
            }
        } catch {
            $status = "status unreadable"
        }
    }

    $values = @(
        $row.id,
        $row.language,
        $row.prompt_id,
        $row.config,
        $row.reason,
        $status,
        $row.output_folder,
        $report
    )
    $lines.Add(('| ``{0}`` | {1} | ``{2}`` | ``{3}`` | {4} | **{5}** | ``{6}`` | ``{7}`` |' -f $values))
}

$lines.Add("")
$lines.Add("_Generated from ``recovery-state\rerun-manifest.csv`` and ``recovery-state\statuses\``._")
$lines -join "`n" | Set-Content -LiteralPath $target -Encoding utf8
