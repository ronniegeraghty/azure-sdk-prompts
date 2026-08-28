"""Run synchronous and asynchronous Azure credential connectivity tests."""

from __future__ import annotations

import argparse
import asyncio

from connectivity_tester import test_credential, test_credential_async
from credential_factory import build_async_credential, build_credential
from environment_detector import detect_environment


ARM_SCOPE = "https://management.azure.com/.default"


def _parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Test an environment-specific Azure credential chain."
    )
    parser.add_argument(
        "--enable-cae",
        action="store_true",
        help="Request Continuous Access Evaluation capable tokens.",
    )
    parser.add_argument(
        "--scope",
        default=ARM_SCOPE,
        help=f"Azure token scope (default: {ARM_SCOPE}).",
    )
    return parser.parse_args()


async def _run_async(scope: str, enable_cae: bool) -> bool:
    detection = detect_environment()
    selection = build_async_credential(detection.environment)
    print(f"Async credential strategy: {selection.strategy}")
    try:
        result = await test_credential_async(
            selection.credential,
            scope,
            enable_cae=enable_cae,
        )
        return result.succeeded
    finally:
        await selection.credential.close()


def main() -> int:
    args = _parse_args()
    detection = detect_environment()
    print(f"Detected environment: {detection.environment.value}")
    print(f"Detection evidence: {', '.join(detection.evidence)}")

    selection = build_credential(detection.environment)
    print(f"Sync credential strategy: {selection.strategy}")
    try:
        sync_result = test_credential(
            selection.credential,
            args.scope,
            enable_cae=args.enable_cae,
        )
    finally:
        close = getattr(selection.credential, "close", None)
        if close:
            close()

    async_succeeded = asyncio.run(_run_async(args.scope, args.enable_cae))
    return 0 if sync_result.succeeded and async_succeeded else 1


if __name__ == "__main__":
    raise SystemExit(main())
