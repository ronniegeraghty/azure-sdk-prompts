"""Detect the runtime, build Azure credentials, and test ARM authentication."""

from __future__ import annotations

import argparse
import asyncio
import os

from connectivity_tester import test_credential_async, test_credential_sync
from credential_factory import build_async_credential, build_sync_credential
from environment_detector import detect_environment

ARM_SCOPE = "https://management.azure.com/.default"


def _environment_flag(name: str) -> bool:
    return os.environ.get(name, "").strip().lower() in {"1", "true", "yes", "on"}


def _parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test an environment-specific Azure credential chain."
    )
    parser.add_argument(
        "--enable-cae",
        action="store_true",
        default=_environment_flag("AZURE_ENABLE_CAE"),
        help="request a Continuous Access Evaluation-capable token",
    )
    return parser.parse_args()


async def _run_async(environment, enable_cae: bool) -> bool:
    selection = build_async_credential(environment, enable_cae=enable_cae)
    print(f"Async credential strategy: {selection.strategy}")
    try:
        return await test_credential_async(
            selection.credential,
            ARM_SCOPE,
            enable_cae=selection.enable_cae,
        )
    finally:
        await selection.credential.close()


def main() -> int:
    args = _parse_args()
    environment = detect_environment()
    print(f"Detected environment: {environment.value}")

    sync_selection = build_sync_credential(
        environment, enable_cae=args.enable_cae
    )
    print(f"Sync credential strategy: {sync_selection.strategy}")
    try:
        sync_ok = test_credential_sync(
            sync_selection.credential,
            ARM_SCOPE,
            enable_cae=sync_selection.enable_cae,
        )
    finally:
        sync_selection.credential.close()

    async_ok = asyncio.run(_run_async(environment, args.enable_cae))
    return 0 if sync_ok and async_ok else 1


if __name__ == "__main__":
    raise SystemExit(main())
