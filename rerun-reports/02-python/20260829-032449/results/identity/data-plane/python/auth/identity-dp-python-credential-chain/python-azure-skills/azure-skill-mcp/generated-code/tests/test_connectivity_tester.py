import asyncio
import time
import unittest

from azure.core.credentials import AccessToken
from azure.core.exceptions import ClientAuthenticationError

from connectivity_tester import test_credential_async, test_credential_sync


class SuccessfulSyncCredential:
    def get_token(self, *scopes, **kwargs):
        return AccessToken("not-logged", int(time.time()) + 3600)


class FailedSyncCredential:
    def get_token(self, *scopes, **kwargs):
        raise ClientAuthenticationError(
            message="AADSTS7000222: The provided client secret keys are expired."
        )


class SuccessfulAsyncCredential:
    async def get_token(self, *scopes, **kwargs):
        return AccessToken("not-logged", int(time.time()) + 3600)


class ConnectivityTesterTests(unittest.TestCase):
    def test_sync_success_records_cae_and_expiry(self):
        result = test_credential_sync(
            SuccessfulSyncCredential(), "scope", enable_cae=True
        )
        self.assertTrue(result.succeeded)
        self.assertTrue(result.cae_requested)
        self.assertIsNotNone(result.expires_on)

    def test_sync_failure_reports_specific_reason(self):
        result = test_credential_sync(FailedSyncCredential(), "scope")
        self.assertFalse(result.succeeded)
        self.assertEqual(
            result.failure_category, "expired client certificate or secret"
        )

    def test_async_success_records_cae_and_expiry(self):
        result = asyncio.run(
            test_credential_async(
                SuccessfulAsyncCredential(), "scope", enable_cae=True
            )
        )
        self.assertTrue(result.succeeded)
        self.assertTrue(result.cae_requested)
        self.assertIsNotNone(result.expires_on)


if __name__ == "__main__":
    unittest.main()
