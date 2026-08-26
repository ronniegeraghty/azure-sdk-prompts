"""Local developer credential-chain example."""

from managed_identity_demo.cli import main

raise SystemExit(main(["--identity", "local", "--list-containers"]))
