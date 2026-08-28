"""Offline unit tests for environment selection and authentication diagnostics."""

from __future__ import annotations

import unittest

from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError

from connectivity_tester import _failure_result
from credential_factory import build_sync_credential
from environment_detector import RuntimeEnvironment, detect_environment


class EnvironmentDetectorTests(unittest.TestCase):
    def test_defaults_to_development(self) -> None:
        self.assertEqual(detect_environment({}), RuntimeEnvironment.DEV)

    def test_ci_takes_precedence_over_production(self) -> None:
        values = {"PIPELINE_WORKSPACE": "work", "IDENTITY_ENDPOINT": "endpoint"}
        self.assertEqual(detect_environment(values), RuntimeEnvironment.CI)

    def test_detects_workload_identity_as_production(self) -> None:
        values = {"AZURE_FEDERATED_TOKEN_FILE": "token"}
        self.assertEqual(detect_environment(values), RuntimeEnvironment.PRODUCTION)


class CredentialFactoryTests(unittest.TestCase):
    def test_production_selects_user_assigned_identity(self) -> None:
        selection = build_sync_credential(
            RuntimeEnvironment.PRODUCTION,
            environ={"AZURE_MANAGED_IDENTITY_CLIENT_ID": "client-id"},
        )
        self.assertIn("user-assigned managed identity", selection.strategy)
        selection.credential.close()

    def test_rejects_incomplete_pipeline_federation_settings(self) -> None:
        with self.assertRaisesRegex(ValueError, "AZURESUBSCRIPTION_TENANT_ID"):
            build_sync_credential(
                RuntimeEnvironment.CI,
                environ={"SYSTEM_ACCESSTOKEN": "token"},
            )

    def test_accepts_standard_azure_pipelines_variables(self) -> None:
        selection = build_sync_credential(
            RuntimeEnvironment.CI,
            environ={
                "AZURESUBSCRIPTION_TENANT_ID": "tenant-id",
                "AZURESUBSCRIPTION_CLIENT_ID": "client-id",
                "AZURESUBSCRIPTION_SERVICE_CONNECTION_ID": "connection-id",
                "SYSTEM_ACCESSTOKEN": "token",
            },
        )
        self.assertIn("Azure Pipelines workload identity", selection.strategy)
        selection.credential.close()


class AuthenticationDiagnosticTests(unittest.TestCase):
    def test_reports_expired_certificate(self) -> None:
        error = ClientAuthenticationError(
            message="AADSTS700027: The certificate has expired."
        )
        result = _failure_result(error, enable_cae=True)
        self.assertIn("certificate is expired", result.failure_reason or "")
        self.assertTrue(result.cae_requested)

    def test_reports_unavailable_identity(self) -> None:
        result = _failure_result(
            CredentialUnavailableError("IMDS endpoint unavailable"),
            enable_cae=False,
        )
        self.assertIn("No identity is available", result.failure_reason or "")


if __name__ == "__main__":
    unittest.main()
