from __future__ import annotations

import unittest
from unittest.mock import Mock, patch

from managed_identity_demo.storage import create_blob_service_client


class CreateBlobServiceClientTests(unittest.TestCase):
    @patch("managed_identity_demo.storage.BlobServiceClient")
    def test_passes_token_credential_to_sdk_client(self, client_type) -> None:
        credential = Mock()

        create_blob_service_client(
            "https://account.blob.core.windows.net/",
            credential,
        )

        client_type.assert_called_once_with(
            account_url="https://account.blob.core.windows.net",
            credential=credential,
            retry_total=4,
            retry_connect=4,
            retry_read=4,
            retry_status=4,
        )

    def test_rejects_insecure_endpoint(self) -> None:
        with self.assertRaisesRegex(ValueError, "HTTPS"):
            create_blob_service_client("http://account.example", Mock())


if __name__ == "__main__":
    unittest.main()

