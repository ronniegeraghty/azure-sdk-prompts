import unittest

from azure.core.exceptions import ClientAuthenticationError
from azure.identity import CredentialUnavailableError

from connectivity_tester import classify_authentication_failure


class AuthenticationFailureTests(unittest.TestCase):
    def test_reports_expired_certificate(self):
        error = ClientAuthenticationError(
            message="The configured X509 certificate has expired."
        )
        self.assertEqual(
            "the service principal certificate has expired",
            classify_authentication_failure(error),
        )

    def test_reports_wrong_tenant(self):
        error = ClientAuthenticationError(
            message="AADSTS90002: Tenant 'example' not found."
        )
        self.assertEqual(
            "the tenant ID or authority is incorrect",
            classify_authentication_failure(error),
        )

    def test_reports_missing_credential(self):
        error = CredentialUnavailableError(
            message="No credential in this chain is available."
        )
        self.assertEqual(
            "no credential in the configured chain could attempt authentication",
            classify_authentication_failure(error),
        )


if __name__ == "__main__":
    unittest.main()
