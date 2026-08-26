from __future__ import annotations

import unittest

from environment_detector import DeploymentEnvironment, detect_environment


class EnvironmentDetectorTests(unittest.TestCase):
    def test_defaults_to_development(self) -> None:
        result = detect_environment({})
        self.assertEqual(DeploymentEnvironment.DEV, result.environment)

    def test_ci_takes_precedence_over_azure_host_markers(self) -> None:
        result = detect_environment(
            {"TF_BUILD": "True", "IDENTITY_ENDPOINT": "http://localhost"}
        )
        self.assertEqual(DeploymentEnvironment.CI, result.environment)

    def test_detects_workload_identity_as_production(self) -> None:
        result = detect_environment(
            {
                "AZURE_TENANT_ID": "tenant",
                "AZURE_CLIENT_ID": "client",
                "AZURE_FEDERATED_TOKEN_FILE": "token",
            }
        )
        self.assertEqual(DeploymentEnvironment.PRODUCTION, result.environment)

    def test_override_is_validated(self) -> None:
        with self.assertRaises(ValueError):
            detect_environment({"APP_ENVIRONMENT": "staging"})


if __name__ == "__main__":
    unittest.main()
