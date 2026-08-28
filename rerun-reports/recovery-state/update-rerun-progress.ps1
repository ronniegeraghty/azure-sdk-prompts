$ErrorActionPreference = "Stop"
$stateRoot = $PSScriptRoot
$manifestPath = Join-Path $stateRoot "rerun-manifest.csv"
$statusDir = Join-Path $stateRoot "statuses"
$commentID = "5439613722"

$retained = @{
    "python" = 18
    "js-ts" = 14
    "java" = 13
    "dotnet" = 20
}
$labels = @{
    "python" = "Python"
    "js-ts" = "JavaScript/TypeScript"
    "java" = "Java"
    "dotnet" = ".NET"
}
$expected = @{
    "python" = 57
    "js-ts" = 42
    "java" = 57
    "dotnet" = 60
}

$manifest = @(Import-Csv -LiteralPath $manifestPath)
$statuses = @()
if (Test-Path -LiteralPath $statusDir) {
    foreach ($file in Get-ChildItem -LiteralPath $statusDir -Filter "lane-*.json" -File) {
        try {
            $statuses += Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
        } catch {
            # Atomic status writes should make this rare; count an unreadable file as running.
            $statuses += [pscustomobject]@{ id = $file.BaseName; status = "running"; language = "unknown" }
        }
    }
}

$validReruns = @($statuses | Where-Object status -eq "valid")
$problems = @($statuses | Where-Object status -eq "problem")
$running = @($statuses | Where-Object status -eq "running")
$finished = $validReruns.Count + $problems.Count
$pending = $manifest.Count - $finished - $running.Count

$retryStatuses = @()
foreach ($retryDirName in @(
    "timeout-statuses",
    "timeout-java-second-statuses",
    "timeout-python-statuses",
    "timeout-python-second-statuses",
    "timeout-js-ts-statuses",
    "timeout-js-ts-batch-statuses",
    "timeout-js-ts-second-statuses"
)) {
    $retryDir = Join-Path $stateRoot $retryDirName
    if (-not (Test-Path -LiteralPath $retryDir)) {
        continue
    }
    foreach ($file in Get-ChildItem -LiteralPath $retryDir -Filter "*.json" -File) {
        try {
            $retryStatuses += Get-Content -Raw -LiteralPath $file.FullName | ConvertFrom-Json
        } catch {
            # Ignore unreadable snapshots; the next atomic status update will replace them.
        }
    }
}

$validRetryKeys = @{}
foreach ($retry in $retryStatuses | Where-Object status -eq "valid") {
    $key = "$($retry.language)|$($retry.prompt_id)|$($retry.config)"
    $validRetryKeys[$key] = $true
}

$recovered = @($problems | Where-Object {
    $key = "$($_.language)|$($_.prompt_id)|$($_.config)"
    $validRetryKeys.ContainsKey($key)
})
$unresolved = @($problems | Where-Object {
    $key = "$($_.language)|$($_.prompt_id)|$($_.config)"
    -not $validRetryKeys.ContainsKey($key)
})
$activeRetries = @($retryStatuses | Where-Object status -eq "running")
$validTotal = 65 + $validReruns.Count + $recovered.Count

$lines = [System.Collections.Generic.List[string]]::new()
$lines.Add("**Progress: $validTotal / 216 valid**")
$lines.Add("")
$lines.Add("Main reruns: **$finished / 151 finished** · Controlled recoveries: **$($recovered.Count)** · Persistent problems: **$($unresolved.Count)**")
$lines.Add("")
$lines.Add("| Language | Valid | Problems | Running | Expected |")
$lines.Add("|---|---:|---:|---:|---:|")

foreach ($language in @("python", "js-ts", "java", "dotnet")) {
    $languageValid = $retained[$language] +
        @($validReruns | Where-Object language -eq $language).Count +
        @($recovered | Where-Object language -eq $language).Count
    $languageProblems = @($unresolved | Where-Object language -eq $language).Count
    $languageRunning = @($running | Where-Object language -eq $language).Count +
        @($activeRetries | Where-Object language -eq $language).Count
    $lines.Add("| $($labels[$language]) | $languageValid | $languageProblems | $languageRunning | $($expected[$language]) |")
}

$lines.Add("")
$lines.Add("_Updated: $([DateTime]::UtcNow.ToString("yyyy-MM-dd HH:mm:ss 'UTC'"))_")
$body = $lines -join "`n"

gh api "repos/ronniegeraghty/hyoka/issues/comments/$commentID" -X PATCH -f body="$body" | Out-Null

[pscustomobject]@{
    Valid = $validTotal
    FinishedReruns = $finished
    ValidReruns = $validReruns.Count
    Recovered = $recovered.Count
    Problems = $unresolved.Count
    Running = $running.Count + $activeRetries.Count
    Pending = $pending
    Complete = ($finished -eq $manifest.Count -and $running.Count -eq 0)
} | ConvertTo-Json -Compress
