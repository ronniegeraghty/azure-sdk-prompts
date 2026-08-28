from __future__ import annotations

import asyncio

from credential_factory import build_async_credential, build_sync_credential
from environment_detector import RuntimeEnvironment


def test_builds_sync_developer_chain() -> None:
    selection = build_sync_credential(RuntimeEnvironment.DEV, enable_cae=True)
    try:
        assert selection.enable_cae is True
        assert "Azure CLI" in selection.strategy
    finally:
        selection.credential.close()


def test_builds_async_developer_chain() -> None:
    async def exercise() -> None:
        selection = build_async_credential(RuntimeEnvironment.DEV)
        try:
            assert selection.enable_cae is False
            assert "VS Code" in selection.strategy
        finally:
            await selection.credential.close()

    asyncio.run(exercise())
