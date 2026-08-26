"""System-assigned managed identity example."""

from managed_identity_demo.cli import main

raise SystemExit(main(["--identity", "system", "--list-containers"]))
