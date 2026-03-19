#!/usr/bin/env python3
"""Scan all .prompt.md files and regenerate manifest.yaml."""

import sys
from pathlib import Path
from datetime import datetime, timezone

try:
    import yaml
except ImportError:
    print("ERROR: PyYAML is required. Install with: pip install pyyaml")
    sys.exit(1)

REPO_ROOT = Path(__file__).resolve().parent.parent
PROMPTS_DIR = REPO_ROOT / "prompts"


def parse_frontmatter(path):
    """Extract YAML frontmatter from a .prompt.md file."""
    content = path.read_text(encoding="utf-8")
    if not content.startswith("---"):
        return None
    end = content.index("---", 3)
    return yaml.safe_load(content[3:end])


def main():
    prompts = []
    services = set()
    languages = set()
    categories = set()

    for prompt_file in sorted(PROMPTS_DIR.rglob("*.prompt.md")):
        meta = parse_frontmatter(prompt_file)
        if not meta:
            print(f"WARNING: No frontmatter in {prompt_file}")
            continue

        rel_path = str(prompt_file.relative_to(REPO_ROOT)).replace("\\", "/")
        meta["path"] = rel_path
        prompts.append(meta)
        services.add(meta.get("service", ""))
        languages.add(meta.get("language", ""))
        categories.add(meta.get("category", ""))

    manifest = {
        "generated_at": datetime.now(timezone.utc).isoformat(),
        "prompt_count": len(prompts),
        "services": sorted(services - {""}),
        "languages": sorted(languages - {""}),
        "categories": sorted(categories - {""}),
        "prompts": prompts,
    }

    manifest_path = REPO_ROOT / "manifest.yaml"
    with open(manifest_path, "w") as f:
        yaml.dump(manifest, f, default_flow_style=False, sort_keys=False)

    print(f"Generated manifest with {len(prompts)} prompts at {manifest_path}")


if __name__ == "__main__":
    main()
