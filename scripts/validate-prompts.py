#!/usr/bin/env python3
"""Validate all prompt files have correct frontmatter."""

import sys
from pathlib import Path

try:
    import yaml
except ImportError:
    print("ERROR: PyYAML is required. Install with: pip install pyyaml")
    sys.exit(1)

REPO_ROOT = Path(__file__).resolve().parent.parent
PROMPTS_DIR = REPO_ROOT / "prompts"

REQUIRED_FIELDS = [
    "id", "service", "plane", "language", "category",
    "difficulty", "description", "created", "author",
]

VALID_SERVICES = [
    "storage", "key-vault", "cosmos-db", "event-hubs",
    "app-configuration", "purview", "digital-twins",
    "identity", "resource-manager", "service-bus",
]
VALID_PLANES = ["data-plane", "management-plane"]
VALID_LANGUAGES = ["dotnet", "java", "js-ts", "python", "go", "rust", "cpp"]
VALID_CATEGORIES = [
    "authentication", "pagination", "polling", "retries",
    "error-handling", "crud", "batch", "streaming",
    "auth", "provisioning",
]
VALID_DIFFICULTIES = ["basic", "intermediate", "advanced"]

ENUM_VALIDATORS = {
    "service": VALID_SERVICES,
    "plane": VALID_PLANES,
    "language": VALID_LANGUAGES,
    "category": VALID_CATEGORIES,
    "difficulty": VALID_DIFFICULTIES,
}


def parse_frontmatter(path):
    """Extract YAML frontmatter from a .prompt.md file."""
    content = path.read_text(encoding="utf-8")
    if not content.startswith("---"):
        return None, "File does not start with YAML frontmatter (---)"
    try:
        end = content.index("---", 3)
    except ValueError:
        return None, "No closing --- for frontmatter block"
    try:
        meta = yaml.safe_load(content[3:end])
    except yaml.YAMLError as e:
        return None, f"Invalid YAML in frontmatter: {e}"
    return meta, None


def has_prompt_section(path):
    """Check if the file has a ## Prompt section with content."""
    content = path.read_text(encoding="utf-8")
    in_prompt = False
    for line in content.split("\n"):
        if line.strip().startswith("## Prompt"):
            in_prompt = True
            continue
        if in_prompt and line.strip().startswith("## "):
            break
        if in_prompt and line.strip():
            return True
    return False


def validate_prompt(path):
    """Validate a single prompt file. Returns list of error strings."""
    errors = []
    rel = path.relative_to(REPO_ROOT)

    meta, parse_err = parse_frontmatter(path)
    if parse_err:
        return [f"{rel}: {parse_err}"]
    if not meta:
        return [f"{rel}: Empty frontmatter"]

    # Check required fields
    for field in REQUIRED_FIELDS:
        if field not in meta or meta[field] is None:
            errors.append(f"{rel}: Missing required field '{field}'")

    # Validate enum fields
    for field, valid_values in ENUM_VALIDATORS.items():
        val = meta.get(field)
        if val and val not in valid_values:
            errors.append(
                f"{rel}: Invalid {field} '{val}'. "
                f"Must be one of: {', '.join(valid_values)}"
            )

    # Check ID naming convention
    id_val = meta.get("id", "")
    if id_val and meta.get("service") and meta.get("plane") and meta.get("language"):
        plane_short = "dp" if meta["plane"] == "data-plane" else "mp"
        expected_prefix = f"{meta['service']}-{plane_short}-{meta['language']}-"
        if not id_val.startswith(expected_prefix):
            errors.append(
                f"{rel}: ID '{id_val}' should start with '{expected_prefix}'"
            )

    # Check for ## Prompt section
    if not has_prompt_section(path):
        errors.append(f"{rel}: Missing or empty '## Prompt' section")

    return errors


def main():
    prompt_files = sorted(PROMPTS_DIR.rglob("*.prompt.md"))

    if not prompt_files:
        print("No .prompt.md files found in prompts/")
        sys.exit(1)

    all_errors = []
    for pf in prompt_files:
        errs = validate_prompt(pf)
        all_errors.extend(errs)

    if all_errors:
        print(f"Validation failed with {len(all_errors)} error(s):\n")
        for err in all_errors:
            print(f"  ✗ {err}")
        sys.exit(1)
    else:
        print(f"✓ All {len(prompt_files)} prompt(s) are valid")


if __name__ == "__main__":
    main()
