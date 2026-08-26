from __future__ import annotations

import asyncio
import io
import unittest
from contextlib import redirect_stdout

from azure.core.credentials import AccessToken
from azure.core.exceptions import ClientAuthenticationError

from connectivity_tester import test_credential, test_credential_async


class FakeCredential:
    def __init__(self) -> None:
        self.enable_cae = False

    def get_token(self, *scopes, **kwargs) -> AccessToken:
        self.enable_cae = kwargs["enable_cae"]
        return AccessToken("not-a-real-token", 2_000_000_000)


class FailingCredential:
    def get_token(self, *scopes, **kwargs) -> AccessToken:
        raise ClientAuthenticationError(
            message="AADSTS90002: Tenant 'wrong' not found."
        )


class FakeAsyncCredential:
    def __init__(self) -> None:
        self.enable_cae = False

    async def get_token(self, *scopes, **kwargs) -> AccessToken:
        self.enable_cae = kwargs["enable_cae"]
        return AccessToken("not-a-real-token", 2_000_000_000)


class ConnectivityTesterTests(unittest.TestCase):
    def test_sync_success_forwards_cae(self) -> None:
        credential = FakeCredential()
        output = io.StringIO()
        with redirect_stdout(output):
            successful = test_credential(
                credential, "scope", enable_cae=True  # type: ignore[arg-type]
            )
        self.assertTrue(successful)
        self.assertTrue(credential.enable_cae)
        self.assertIn("SUCCESS", output.getvalue())

    def test_sync_failure_explains_wrong_tenant(self) -> None:
        output = io.StringIO()
        with redirect_stdout(output):
            successful = test_credential(
                FailingCredential(), "scope"  # type: ignore[arg-type]
            )
        self.assertFalse(successful)
        self.assertIn("tenant ID is incorrect", output.getvalue())

    def test_async_success_forwards_cae(self) -> None:
        credential = FakeAsyncCredential()
        output = io.StringIO()
        with redirect_stdout(output):
            successful = asyncio.run(
                test_credential_async(
                    credential,  # type: ignore[arg-type]
                    "scope",
                    enable_cae=True,
                )
            )
        self.assertTrue(successful)
        self.assertTrue(credential.enable_cae)
        self.assertIn("SUCCESS", output.getvalue())


if __name__ == "__main__":
    unittest.main()
