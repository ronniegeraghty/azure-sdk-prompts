#!/usr/bin/env python3
"""
Run doc-agent evaluate against all (or filtered) prompts in the repo.

Usage:
  python scripts/run-evals.py                                    # All prompts
  python scripts/run-evals.py --service storage                  # All Storage prompts
  python scripts/run-evals.py --language dotnet                  # All .NET prompts
  python scripts/run-evals.py --service storage --language dotnet # Storage + .NET
  python scripts/run-evals.py --category authentication          # All auth prompts
  python scripts/run-evals.py --plane data-plane                 # All data-plane prompts
  python scripts/run-evals.py --tags identity                    # Filter by tag
  python scripts/run-evals.py --prompt-id storage-dp-dotnet-auth # Single by ID
  python scripts/run-evals.py --prompt prompts/storage/.../x.prompt.md  # Single by path
  python scripts/run-evals.py --service storage --dry-run        # List without running
"""

import argparse
import datetime
import os
import shutil
import subprocess
import sys
from pathlib import Path

try:
    import yaml
except ImportError:
    print("ERROR: PyYAML is required. Install with: pip install pyyaml")
    sys.exit(1)

REPO_ROOT = Path(__file__).resolve().parent.parent
PROMPTS_DIR = REPO_ROOT / "prompts"
REPORTS_DIR = REPO_ROOT / "reports" / "runs"
MANIFEST_PATH = REPO_ROOT / "manifest.yaml"


def load_manifest():
    """Load the central manifest."""
    if not MANIFEST_PATH.exists():
        print(f"ERROR: Manifest not found at {MANIFEST_PATH}")
        print("Run: python scripts/generate-manifest.py")
        sys.exit(1)
    with open(MANIFEST_PATH) as f:
        return yaml.safe_load(f)


def filter_prompts(manifest, args):
    """Apply all filter flags. Filters compose with AND logic."""
    prompts = manifest.get("prompts", [])

    for key in ["service", "language", "plane", "category"]:
        val = getattr(args, key, None)
        if val:
            prompts = [p for p in prompts if p.get(key) == val]

    if args.tags:
        prompts = [p for p in prompts if args.tags in p.get("tags", [])]

    if args.prompt_id:
        prompts = [p for p in prompts if p["id"] == args.prompt_id]

    if args.prompt:
        normalized = args.prompt.replace("\\", "/")
        prompts = [p for p in prompts if p["path"] == normalized]

    return prompts


def extract_prompt_text(prompt_path):
    """Extract the prompt text from the ## Prompt section of the markdown file."""
    content = Path(prompt_path).read_text(encoding="utf-8")
    in_prompt = False
    lines = []
    for line in content.split("\n"):
        if line.strip().startswith("## Prompt"):
            in_prompt = True
            continue
        if in_prompt and line.strip().startswith("## "):
            break
        if in_prompt:
            lines.append(line)
    return "\n".join(lines).strip()


def run_single_eval(prompt_text, output_dir, args):
    """Run doc-agent evaluate for one prompt."""
    cmd = ["doc-agent", "evaluate", prompt_text, "-o", str(output_dir)]
    if args.model:
        cmd += ["-m", args.model]
    if args.timeout:
        cmd += ["-t", str(args.timeout)]
    if args.verbose:
        cmd.append("-v")

    try:
        result = subprocess.run(
            cmd,
            capture_output=True,
            text=True,
            timeout=args.timeout or 3600,
        )
        return {
            "exit_code": result.returncode,
            "stdout": result.stdout,
            "stderr": result.stderr,
        }
    except FileNotFoundError:
        return {
            "exit_code": 127,
            "stdout": "",
            "stderr": "ERROR: doc-agent command not found. "
                      "Install from https://github.com/coreai-microsoft/doc-review-agent",
        }
    except subprocess.TimeoutExpired:
        return {
            "exit_code": 124,
            "stdout": "",
            "stderr": f"ERROR: Evaluation timed out after {args.timeout or 3600}s",
        }


def reorganize_eval_output(raw_output_dir, target_dir):
    """
    doc-agent writes to output/eval-<timestamp>-<slug>/.
    Move those files into our structured report directory.
    """
    eval_dirs = sorted(raw_output_dir.glob("eval-*"))
    if not eval_dirs:
        return

    src = eval_dirs[-1]
    target_dir.mkdir(parents=True, exist_ok=True)

    for item in src.iterdir():
        dest = target_dir / item.name
        if item.is_dir():
            shutil.copytree(str(item), str(dest), dirs_exist_ok=True)
        else:
            shutil.move(str(item), str(dest))

    # Clean up the source eval directory
    shutil.rmtree(str(src), ignore_errors=True)


def main():
    parser = argparse.ArgumentParser(
        description="Run doc-agent evaluate against prompts.",
        epilog="No arguments = run ALL prompts. Flags compose with AND logic.",
    )

    # Filter flags
    parser.add_argument("--service", help="Filter by service (e.g., storage, key-vault)")
    parser.add_argument("--language", help="Filter by language (e.g., dotnet, python)")
    parser.add_argument("--plane", help="Filter by plane (data-plane, management-plane)")
    parser.add_argument("--category", help="Filter by category (e.g., authentication, crud)")
    parser.add_argument("--tags", help="Filter by tag")
    parser.add_argument("--prompt-id", help="Run a single prompt by its ID")
    parser.add_argument("--prompt", help="Run a single prompt by file path")

    # Execution options
    parser.add_argument("--dry-run", action="store_true",
                        help="List matching prompts without running evaluations")
    parser.add_argument("--model", "-m", help="Override doc-agent model")
    parser.add_argument("--timeout", "-t", type=int,
                        help="Timeout per evaluation in seconds (default: 3600)")
    parser.add_argument("--verbose", "-v", action="store_true",
                        help="Verbose doc-agent output")

    args = parser.parse_args()

    manifest = load_manifest()
    prompts = filter_prompts(manifest, args)

    if not prompts:
        print("No prompts matched the given filters.")
        sys.exit(1)

    print(f"Found {len(prompts)} prompt(s) to evaluate")

    if args.dry_run:
        for p in prompts:
            print(f"  [{p['id']}] {p['path']}")
        return

    # Create timestamped run directory
    timestamp = datetime.datetime.utcnow().strftime("%Y-%m-%dT%H-%M-%SZ")
    run_dir = REPORTS_DIR / timestamp
    run_dir.mkdir(parents=True, exist_ok=True)

    results = []
    for i, prompt_meta in enumerate(prompts, 1):
        print(f"\n[{i}/{len(prompts)}] Evaluating {prompt_meta['id']}...")

        prompt_path = REPO_ROOT / prompt_meta["path"]
        prompt_text = extract_prompt_text(prompt_path)

        if not prompt_text:
            print(f"  WARNING: No prompt text found in {prompt_meta['path']}, skipping")
            results.append({
                "prompt_id": prompt_meta["id"],
                "status": "skipped",
                "report_path": "",
                "error": "No ## Prompt section found",
            })
            continue

        # Determine per-prompt report directory
        prompt_rel = prompt_meta["path"] \
            .replace("prompts/", "") \
            .replace(".prompt.md", "")
        report_subdir = run_dir / prompt_rel

        # Run doc-agent evaluate into a temp output dir, then reorganize
        tmp_output = run_dir / "_tmp_eval"
        tmp_output.mkdir(exist_ok=True)

        eval_result = run_single_eval(prompt_text, tmp_output, args)
        reorganize_eval_output(tmp_output, report_subdir)

        status = "pass" if eval_result["exit_code"] == 0 else "fail"
        entry = {
            "prompt_id": prompt_meta["id"],
            "status": status,
            "report_path": str(report_subdir.relative_to(run_dir)),
        }
        if eval_result["exit_code"] != 0:
            entry["error"] = eval_result["stderr"][:500]

        results.append(entry)
        print(f"  Result: {status}")

    # Clean up temp dir
    tmp_output = run_dir / "_tmp_eval"
    if tmp_output.exists():
        shutil.rmtree(str(tmp_output), ignore_errors=True)

    # Gather active filters for metadata
    active_filters = {}
    for key in ["service", "language", "plane", "category", "tags",
                "prompt_id", "prompt"]:
        val = getattr(args, key.replace("-", "_"), None)
        if val:
            active_filters[key] = val

    pass_count = sum(1 for r in results if r["status"] == "pass")
    fail_count = sum(1 for r in results if r["status"] == "fail")
    skip_count = sum(1 for r in results if r["status"] == "skipped")

    run_meta = {
        "timestamp": timestamp,
        "prompt_count": len(prompts),
        "pass_count": pass_count,
        "fail_count": fail_count,
        "skip_count": skip_count,
        "filters": active_filters if active_filters else "none (all prompts)",
        "results": results,
    }

    with open(run_dir / "run-metadata.yaml", "w") as f:
        yaml.dump(run_meta, f, default_flow_style=False, sort_keys=False)

    # Update latest symlink
    latest = REPORTS_DIR / "latest"
    if latest.is_symlink() or latest.exists():
        latest.unlink()
    os.symlink(timestamp, str(latest))

    print(f"\nRun complete: {pass_count}/{len(prompts)} passed", end="")
    if skip_count:
        print(f", {skip_count} skipped", end="")
    if fail_count:
        print(f", {fail_count} failed", end="")
    print(f"\nReports: {run_dir}")


if __name__ == "__main__":
    main()
