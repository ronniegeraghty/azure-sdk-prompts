import unittest

from environment_detector import RuntimeEnvironment, detect_environment


class DetectEnvironmentTests(unittest.TestCase):
    def test_defaults_to_development(self):
        result = detect_environment({})
        self.assertEqual(RuntimeEnvironment.DEV, result.environment)

    def test_ci_marker_takes_precedence(self):
        result = detect_environment(
            {"CI": "true", "IDENTITY_ENDPOINT": "http://localhost/identity"}
        )
        self.assertEqual(RuntimeEnvironment.CI, result.environment)

    def test_managed_identity_endpoint_means_production(self):
        result = detect_environment({"IDENTITY_ENDPOINT": "http://localhost/identity"})
        self.assertEqual(RuntimeEnvironment.PRODUCTION, result.environment)

    def test_workload_identity_means_production(self):
        result = detect_environment(
            {
                "AZURE_TENANT_ID": "tenant",
                "AZURE_CLIENT_ID": "client",
                "AZURE_FEDERATED_TOKEN_FILE": "token-file",
            }
        )
        self.assertEqual(RuntimeEnvironment.PRODUCTION, result.environment)

    def test_override_is_honored(self):
        result = detect_environment({"AZURE_AUTH_ENVIRONMENT": "production"})
        self.assertEqual(RuntimeEnvironment.PRODUCTION, result.environment)

    def test_invalid_override_is_rejected(self):
        with self.assertRaisesRegex(ValueError, "must be one of"):
            detect_environment({"AZURE_AUTH_ENVIRONMENT": "staging"})


if __name__ == "__main__":
    unittest.main()
