import unittest

from environment_detector import RuntimeEnvironment, detect_environment


class EnvironmentDetectorTests(unittest.TestCase):
    def test_defaults_to_dev(self):
        self.assertEqual(detect_environment({}), RuntimeEnvironment.DEV)

    def test_detects_ci_before_azure_host(self):
        environment = detect_environment(
            {"TF_BUILD": "True", "IDENTITY_ENDPOINT": "http://localhost"}
        )
        self.assertEqual(environment, RuntimeEnvironment.CI)

    def test_detects_managed_identity_host_as_production(self):
        self.assertEqual(
            detect_environment({"IDENTITY_ENDPOINT": "http://localhost"}),
            RuntimeEnvironment.PRODUCTION,
        )

    def test_detects_workload_identity_as_production(self):
        environment = detect_environment(
            {
                "AZURE_TENANT_ID": "tenant",
                "AZURE_CLIENT_ID": "client",
                "AZURE_FEDERATED_TOKEN_FILE": "token-file",
            }
        )
        self.assertEqual(environment, RuntimeEnvironment.PRODUCTION)

    def test_explicit_override_wins(self):
        self.assertEqual(
            detect_environment({"APP_ENV": "local", "CI": "true"}),
            RuntimeEnvironment.DEV,
        )

    def test_invalid_override_is_rejected(self):
        with self.assertRaisesRegex(ValueError, "APP_ENV must be one of"):
            detect_environment({"APP_ENV": "staging"})


if __name__ == "__main__":
    unittest.main()
