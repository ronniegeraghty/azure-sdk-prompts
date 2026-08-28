"""Local-only tests that never contact Azure."""

from __future__ import annotations

import asyncio
import time
import unittest

from azure.core.credentials import AccessToken
from azure.core.exceptions import ClientAuthenticationError

from connectivity_tester import (
    explain_authentication_failure,
    test_credential as run_sync_test,
    test_credential_async as run_async_test,
)
from credential_factory import build_credential
from environment_detector import DeploymentEnvironment, detect_environment


class FakeCredential:
    def __init__(self) -> None:
        self.enable_cae: bool | None = None

    def get_token(self, *scopes: str, **kwargs: object) -> AccessToken:
        self.enable_cae = bool(kwargs.get("enable_cae"))
        return AccessToken("not-a-real-token", int(time.time()) + 3600)


class AsyncFakeCredential:
    def __init__(self) -> None:
        self.enable_cae: bool | None = None

    async def get_token(
        self, *scopes: str, **kwargs: object
    ) -> AccessToken:
        self.enable_cae = bool(kwargs.get("enable_cae"))
        return AccessToken("not-a-real-token", int(time.time()) + 3600)

    async def close(self) -> None:
        return None


class EnvironmentDetectorTests(unittest.TestCase):
    def test_defaults_to_dev(self) -> None:
        self.assertEqual(
            detect_environment({}).environment,
            DeploymentEnvironment.DEV,
        )

    def test_ci_takes_precedence_over_managed_identity(self) -> None:
        result = detect_environment(
            {"CI": "true", "IDENTITY_ENDPOINT": "http://localhost"}
        )
        self.assertEqual(result.environment, DeploymentEnvironment.CI)

    def test_workload_identity_is_production(self) -> None:
        result = detect_environment(
            {"AZURE_FEDERATED_TOKEN_FILE": "fake-token-path"}
        )
        self.assertEqual(result.environment, DeploymentEnvironment.PRODUCTION)


class CredentialFactoryTests(unittest.TestCase):
    def test_user_assigned_identity_is_described(self) -> None:
        selection = build_credential(
            DeploymentEnvironment.PRODUCTION,
            {"AZURE_MANAGED_IDENTITY_CLIENT_ID": "fake-client-id"},
        )
        self.assertIn("user-assigned managed identity", selection.strategy)
        selection.credential.close()

    def test_ci_service_connection_is_added(self) -> None:
        selection = build_credential(
            DeploymentEnvironment.CI,
            {
                "AZURE_TENANT_ID": "fake-tenant",
                "AZURE_CLIENT_ID": "fake-client",
                "AZURE_SERVICE_CONNECTION_ID": "fake-connection",
                "SYSTEM_ACCESSTOKEN": "fake-access-token",
            },
        )
        self.assertIn("Azure Pipelines service connection", selection.strategy)
        selection.credential.close()


class ConnectivityTests(unittest.TestCase):
    def test_sync_tester_forwards_cae(self) -> None:
        credential = FakeCredential()
        result = run_sync_test(
            credential,
            "https://management.azure.com/.default",
            enable_cae=True,
        )
        self.assertTrue(result.succeeded)
        self.assertTrue(credential.enable_cae)

    def test_async_tester_forwards_cae(self) -> None:
        credential = AsyncFakeCredential()
        result = asyncio.run(
            run_async_test(
                credential,
                "https://management.azure.com/.default",
                enable_cae=True,
            )
        )
        self.assertTrue(result.succeeded)
        self.assertTrue(credential.enable_cae)

    def test_expired_secret_is_specific(self) -> None:
        detail = explain_authentication_failure(
            ClientAuthenticationError(
                message="AADSTS7000222: The provided client secret has expired."
            )
        )
        self.assertIn("client secret has expired", detail)

    def test_chained_unavailable_credentials_are_specific(self) -> None:
        detail = explain_authentication_failure(
            ClientAuthenticationError(
                message=(
                    "AzureCliCredential: Azure CLI not found on path\n"
                    "ManagedIdentityCredential: authentication unavailable"
                )
            )
        )
        self.assertIn("No configured credential source is available", detail)


if __name__ == "__main__":
    unittest.main()
