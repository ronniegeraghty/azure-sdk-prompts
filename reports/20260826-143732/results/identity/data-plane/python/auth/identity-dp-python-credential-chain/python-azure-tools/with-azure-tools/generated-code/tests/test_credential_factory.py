from __future__ import annotations

import asyncio
import unittest

from credential_factory import build_async_credential, build_credential
from environment_detector import DeploymentEnvironment


class CredentialFactoryTests(unittest.TestCase):
    def test_development_strategy_uses_developer_tools(self) -> None:
        bundle = build_credential(DeploymentEnvironment.DEV, {})
        try:
            self.assertIn("Azure CLI", bundle.strategy)
            self.assertIn("VS Code", bundle.strategy)
        finally:
            bundle.credential.close()

    def test_ci_strategy_supports_azure_pipelines(self) -> None:
        bundle = build_credential(
            DeploymentEnvironment.CI,
            {
                "AZURE_TENANT_ID": "tenant",
                "AZURE_CLIENT_ID": "client",
                "AZURE_SERVICE_CONNECTION_ID": "connection",
                "SYSTEM_ACCESSTOKEN": "fake-token",
            },
        )
        try:
            self.assertIn("Azure Pipelines", bundle.strategy)
            self.assertIn("environment credential", bundle.strategy)
        finally:
            bundle.credential.close()

    def test_production_prefers_user_assigned_managed_identity(self) -> None:
        bundle = build_credential(
            DeploymentEnvironment.PRODUCTION,
            {"AZURE_MANAGED_IDENTITY_CLIENT_ID": "managed-client"},
        )
        try:
            self.assertEqual("user-assigned managed identity", bundle.strategy)
        finally:
            bundle.credential.close()

    def test_production_adds_workload_identity_fallback(self) -> None:
        bundle = build_credential(
            DeploymentEnvironment.PRODUCTION,
            {
                "AZURE_TENANT_ID": "tenant",
                "AZURE_CLIENT_ID": "client",
                "AZURE_FEDERATED_TOKEN_FILE": "token-file",
            },
        )
        try:
            self.assertIn("system-assigned managed identity", bundle.strategy)
            self.assertIn("workload identity fallback", bundle.strategy)
        finally:
            bundle.credential.close()

    def test_async_factory_builds_each_strategy(self) -> None:
        async def build_and_close() -> None:
            for environment in DeploymentEnvironment:
                bundle = build_async_credential(environment, {})
                self.assertTrue(bundle.strategy)
                await bundle.credential.close()

        asyncio.run(build_and_close())


if __name__ == "__main__":
    unittest.main()
