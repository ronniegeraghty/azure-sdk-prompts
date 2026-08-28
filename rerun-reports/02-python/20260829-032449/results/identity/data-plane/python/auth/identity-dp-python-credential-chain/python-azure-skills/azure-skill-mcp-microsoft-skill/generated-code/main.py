"""Run synchronous and asynchronous Azure credential connectivity checks."""

from __future__ import annotations

import argparse
import asyncio
import os

from async_connectivity_tester import test_connectivity_async
from connectivity_tester import test_connectivity
from credential_factory import build_async_credential, build_sync_credential
from environment_detector import detect_environment

ARM_SCOPE = "https://management.azure.com/.default"


def _environment_flag(name: str, default: bool = False) -> bool:
    value = os.getenv(name)
    if value is None:
        return default
    normalized = value.strip().lower()
    if normalized in {"1", "true", "yes", "on"}:
        return True
    if normalized in {"0", "false", "no", "off"}:
        return False
    raise ValueError(f"{name} must be a boolean value")


def _parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test an environment-specific Azure credential chain."
    )
    parser.add_argument(
        "--scope",
        default=ARM_SCOPE,
        help=f"Azure token scope (default: {ARM_SCOPE})",
    )
    parser.add_argument(
        "--enable-cae",
        action="store_true",
        default=_environment_flag("AZURE_ENABLE_CAE"),
        help="request a Continuous Access Evaluation-capable token",
    )
    return parser.parse_args()


async def _run_async(scope: str, environment, enable_cae: bool) -> bool:
    selection = build_async_credential(environment, enable_cae=enable_cae)
    print("\nAsync connectivity test")
    print(f"  Strategy: {selection.strategy}")
    try:
        return await test_connectivity_async(
            selection.credential,
            scope,
            enable_cae=selection.enable_cae,
        )
    finally:
        await selection.credential.close()


def main() -> int:
    args = _parse_args()
    detection = detect_environment()
    print(f"Detected environment: {detection.environment.value}")
    print(f"Detection reason: {detection.reason}")

    sync_selection = build_sync_credential(
        detection.environment, enable_cae=args.enable_cae
    )
    print("\nSync connectivity test")
    print(f"  Strategy: {sync_selection.strategy}")
    try:
        sync_succeeded = test_connectivity(
            sync_selection.credential,
            args.scope,
            enable_cae=sync_selection.enable_cae,
        )
    finally:
        sync_selection.credential.close()

    async_succeeded = asyncio.run(
        _run_async(args.scope, detection.environment, args.enable_cae)
    )
    return 0 if sync_succeeded and async_succeeded else 1


if __name__ == "__main__":
    raise SystemExit(main())
