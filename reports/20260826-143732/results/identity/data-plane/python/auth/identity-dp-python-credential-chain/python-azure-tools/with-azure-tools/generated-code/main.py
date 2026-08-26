"""Run sync and async Azure Resource Manager authentication checks."""

from __future__ import annotations

import argparse
import asyncio

from connectivity_tester import test_credential, test_credential_async
from credential_factory import build_async_credential, build_credential
from environment_detector import DeploymentEnvironment, detect_environment


ARM_SCOPE = "https://management.azure.com/.default"


def _parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test the environment-specific Azure credential chain."
    )
    parser.add_argument(
        "--enable-cae",
        action="store_true",
        help="request a Continuous Access Evaluation capable token",
    )
    return parser.parse_args()


async def _run_async(
    environment: DeploymentEnvironment,
    enable_cae: bool,
) -> bool:
    bundle = build_async_credential(environment)
    print(f"\nAsync credential strategy: {bundle.strategy}")
    try:
        return await test_credential_async(
            bundle.credential,
            ARM_SCOPE,
            enable_cae=enable_cae,
        )
    finally:
        await bundle.credential.close()


def main() -> int:
    args = _parse_args()
    detection = detect_environment()
    print(f"Detected environment: {detection.environment.value}")
    print(f"Detection reason: {detection.reason}")

    sync_bundle = build_credential(detection.environment)
    print(f"\nSync credential strategy: {sync_bundle.strategy}")
    try:
        sync_success = test_credential(
            sync_bundle.credential,
            ARM_SCOPE,
            enable_cae=args.enable_cae,
        )
    finally:
        sync_bundle.credential.close()

    async_success = asyncio.run(
        _run_async(detection.environment, args.enable_cae)
    )
    return 0 if sync_success and async_success else 1


if __name__ == "__main__":
    raise SystemExit(main())
