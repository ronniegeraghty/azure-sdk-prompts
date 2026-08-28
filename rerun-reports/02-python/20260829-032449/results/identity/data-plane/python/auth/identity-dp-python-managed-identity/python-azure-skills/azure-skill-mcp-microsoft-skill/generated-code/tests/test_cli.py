from unittest import TestCase
from unittest.mock import MagicMock, patch

from azure.core.exceptions import (
    ClientAuthenticationError,
    HttpResponseError,
    ServiceRequestError,
)
from azure.identity import CredentialUnavailableError

from managed_identity_demo.__main__ import run


class RunTests(TestCase):
    def test_account_url_is_required(self):
        with patch.dict("os.environ", {}, clear=True):
            self.assertEqual(run([]), 2)

    @patch("managed_identity_demo.__main__.list_blob_containers")
    @patch("managed_identity_demo.__main__.create_credential")
    def test_lists_containers(self, create_credential, list_containers):
        credential = MagicMock()
        create_credential.return_value = credential
        list_containers.return_value = ["one", "two"]

        result = run(
            [
                "--identity",
                "system",
                "--account-url",
                "https://example.blob.core.windows.net",
            ]
        )

        self.assertEqual(result, 0)
        list_containers.assert_called_once_with(
            "https://example.blob.core.windows.net",
            credential,
        )

    def _run_with_error(self, error: Exception) -> int:
        with patch(
            "managed_identity_demo.__main__.create_credential"
        ) as factory, patch(
            "managed_identity_demo.__main__.list_blob_containers"
        ) as operation:
            factory.return_value = MagicMock()
            operation.side_effect = error
            return run(
                [
                    "--account-url",
                    "https://example.blob.core.windows.net",
                ]
            )

    def test_credential_unavailable_exit_code(self):
        self.assertEqual(
            self._run_with_error(CredentialUnavailableError("unavailable")),
            3,
        )

    def test_authentication_failure_exit_code(self):
        self.assertEqual(
            self._run_with_error(ClientAuthenticationError("failed")),
            4,
        )

    def test_authorization_failure_exit_code(self):
        response = MagicMock()
        response.status_code = 403
        self.assertEqual(
            self._run_with_error(HttpResponseError("forbidden", response=response)),
            5,
        )

    def test_network_failure_exit_code(self):
        self.assertEqual(
            self._run_with_error(ServiceRequestError("network unavailable")),
            6,
        )
