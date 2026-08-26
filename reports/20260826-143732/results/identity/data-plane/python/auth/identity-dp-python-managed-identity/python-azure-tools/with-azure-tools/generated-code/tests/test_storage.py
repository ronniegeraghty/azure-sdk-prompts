from unittest.mock import MagicMock, patch

import pytest

from managed_identity_demo.storage import (
    StorageConfigurationError,
    list_container_names,
    validate_account_url,
)


@pytest.mark.parametrize(
    "url",
    [
        "http://account.blob.core.windows.net",
        "https://user:password@account.blob.core.windows.net",
        "https://account.blob.core.windows.net?sig=secret",
        "not-a-url",
    ],
)
def test_validate_account_url_rejects_unsafe_values(url):
    with pytest.raises(StorageConfigurationError):
        validate_account_url(url)


@patch("managed_identity_demo.storage.BlobServiceClient")
def test_list_container_names_uses_token_credential(mock_client_type):
    credential = MagicMock()
    client = mock_client_type.return_value.__enter__.return_value
    client.list_containers.return_value = [{"name": "one"}, {"name": "two"}]

    names = list_container_names(
        "https://account.blob.core.windows.net/",
        credential,
    )

    assert names == ["one", "two"]
    mock_client_type.assert_called_once_with(
        account_url="https://account.blob.core.windows.net",
        credential=credential,
        retry_total=3,
        retry_backoff_factor=0.8,
        connection_timeout=10,
        read_timeout=30,
    )
