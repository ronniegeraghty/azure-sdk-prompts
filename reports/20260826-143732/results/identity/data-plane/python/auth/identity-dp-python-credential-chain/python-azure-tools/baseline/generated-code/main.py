"""Detect the environment and test its Azure credential chain."""

from __future__ import annotations

import argparse
import asyncio

from connectivity_tester import test_credential, test_credential_async
from credential_factory import build_async_credential, build_credential
from environment_detector import detect_environment


ARM_SCOPE = "https://management.azure.com/.default"


def _arguments() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test an environment-specific Azure credential chain."
    )
    parser.add_argument(
        "--enable-cae",
        action="store_true",
        help="Request Continuous Access Evaluation when acquiring tokens.",
    )
    parser.add_argument(
        "--scope",
        default=ARM_SCOPE,
        help=f"Azure token scope (default: {ARM_SCOPE}).",
    )
    return parser.parse_args()


async def _run_async(scope: str, enable_cae: bool, deployment_environment) -> bool:
    selection = build_async_credential(
        deployment_environment, enable_cae=enable_cae
    )
    print(f"Async credential strategy: {selection.strategy}")
    try:
        return await test_credential_async(
            selection.credential,
            scope,
            enable_cae=selection.enable_cae,
        )
    finally:
        await selection.credential.close()


def main() -> int:
    args = _arguments()
    deployment_environment = detect_environment()
    selection = build_credential(
        deployment_environment, enable_cae=args.enable_cae
    )

    print(f"Detected environment: {deployment_environment.value}")
    print(f"Sync credential strategy: {selection.strategy}")
    try:
        sync_succeeded = test_credential(
            selection.credential,
            args.scope,
            enable_cae=selection.enable_cae,
        )
    finally:
        selection.credential.close()

    async_succeeded = asyncio.run(
        _run_async(args.scope, args.enable_cae, deployment_environment)
    )
    return 0 if sync_succeeded and async_succeeded else 1


if __name__ == "__main__":
    raise SystemExit(main())
