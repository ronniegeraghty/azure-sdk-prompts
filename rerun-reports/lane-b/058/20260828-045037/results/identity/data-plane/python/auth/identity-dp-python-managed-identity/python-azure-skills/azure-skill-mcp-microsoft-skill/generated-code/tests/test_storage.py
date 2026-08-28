import unittest
from unittest.mock import MagicMock, patch

from azure.core.exceptions import HttpResponseError

from managed_identity_demo.storage import AzureOperationError, list_container_names


class StorageTests(unittest.TestCase):
    @patch("managed_identity_demo.storage.BlobServiceClient")
    def test_lists_container_names(self, client_class):
        service = client_class.return_value.__enter__.return_value
        first = MagicMock()
        first.name = "documents"
        second = MagicMock()
        second.name = "images"
        service.list_containers.return_value = [first, second]

        result = list_container_names(
            "https://example.blob.core.windows.net", MagicMock()
        )

        self.assertEqual(result, ["documents", "images"])

    @patch("managed_identity_demo.storage.BlobServiceClient")
    def test_translates_forbidden_response(self, client_class):
        service = client_class.return_value.__enter__.return_value
        response = MagicMock()
        response.status_code = 403
        service.list_containers.side_effect = HttpResponseError(response=response)

        with self.assertRaisesRegex(AzureOperationError, "Storage Blob Data Reader"):
            list_container_names(
                "https://example.blob.core.windows.net", MagicMock()
            )


if __name__ == "__main__":
    unittest.main()
