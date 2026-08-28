"""Run sync and async Azure Resource Manager authentication checks."""

from __future__ import annotations

import argparse
import asyncio
import os

from connectivity_tester import test_credential, test_credential_async
from credential_factory import build_async_credential, build_sync_credential
from environment_detector import detect_environment


ARM_SCOPE = "https://management.azure.com/.default"


def _cae_default() -> bool:
    value = os.getenv("AZURE_ENABLE_CAE", "").strip().lower()
    return value in {"1", "true", "yes", "on"}


def _parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test environment-specific Azure credential chains."
    )
    parser.add_argument(
        "--cae",
        action=argparse.BooleanOptionalAction,
        default=_cae_default(),
        help="request a CAE-enabled token (default: AZURE_ENABLE_CAE)",
    )
    return parser.parse_args()


async def _run_async(environment, enable_cae: bool) -> None:
    selection = build_async_credential(environment, enable_cae=enable_cae)
    print(f"Async credential strategy: {selection.strategy}")
    async with selection.credential:
        await test_credential_async(
            selection.credential,
            ARM_SCOPE,
            enable_cae=selection.enable_cae,
        )


def main() -> None:
    args = _parse_args()
    environment = detect_environment()
    print(f"Detected environment: {environment.value}")

    selection = build_sync_credential(environment, enable_cae=args.cae)
    print(f"Sync credential strategy: {selection.strategy}")
    with selection.credential:
        test_credential(
            selection.credential,
            ARM_SCOPE,
            enable_cae=selection.enable_cae,
        )

    asyncio.run(_run_async(environment, args.cae))


if __name__ == "__main__":
    main()
