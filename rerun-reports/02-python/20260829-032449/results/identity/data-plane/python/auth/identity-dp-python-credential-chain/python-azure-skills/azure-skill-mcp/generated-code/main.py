"""Run synchronous and asynchronous Azure credential connectivity tests."""

from __future__ import annotations

import asyncio
import os

from connectivity_tester import test_credential_async, test_credential_sync
from credential_factory import build_async_credential, build_sync_credential
from environment_detector import detect_environment


ARM_SCOPE = "https://management.azure.com/.default"


def _cae_enabled() -> bool:
    value = os.getenv("AZURE_ENABLE_CAE", "false").strip().lower()
    if value in {"1", "true", "yes", "on"}:
        return True
    if value in {"0", "false", "no", "off"}:
        return False
    raise ValueError(
        "AZURE_ENABLE_CAE must be one of: true, false, 1, 0, yes, no, on, off"
    )


async def _run_async(environment, enable_cae: bool) -> None:
    selection = build_async_credential(environment, enable_cae=enable_cae)
    print(f"\nAsync credential strategy: {selection.strategy}")
    try:
        await test_credential_async(
            selection.credential,
            ARM_SCOPE,
            enable_cae=selection.enable_cae,
        )
    finally:
        await selection.credential.close()


def main() -> None:
    environment = detect_environment()
    enable_cae = _cae_enabled()
    print(f"Detected environment: {environment.value}")
    print(f"CAE token requests enabled: {'yes' if enable_cae else 'no'}")

    selection = build_sync_credential(environment, enable_cae=enable_cae)
    print(f"\nSync credential strategy: {selection.strategy}")
    try:
        test_credential_sync(
            selection.credential,
            ARM_SCOPE,
            enable_cae=selection.enable_cae,
        )
    finally:
        selection.credential.close()

    asyncio.run(_run_async(environment, enable_cae))


if __name__ == "__main__":
    main()
