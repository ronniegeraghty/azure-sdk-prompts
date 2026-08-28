from unittest.mock import Mock, patch

from azure.core.exceptions import ClientAuthenticationError, HttpResponseError

from managed_identity_demo.cli import run


@patch("managed_identity_demo.cli.list_container_names")
@patch("managed_identity_demo.cli.create_blob_service_client")
@patch("managed_identity_demo.cli.create_credential")
def test_run_lists_containers(create_credential, create_client, list_names, capsys):
    list_names.return_value = iter(["alpha", "beta"])

    result = run(
        [
            "--identity",
            "system",
            "--account-url",
            "https://example.blob.core.windows.net",
        ]
    )

    assert result == 0
    assert capsys.readouterr().out == "alpha\nbeta\n"


@patch("managed_identity_demo.cli.create_credential")
def test_run_reports_authentication_failure(create_credential, capsys):
    create_credential.side_effect = ClientAuthenticationError("token unavailable")

    result = run(
        [
            "--identity",
            "system",
            "--account-url",
            "https://example.blob.core.windows.net",
        ]
    )

    assert result == 3
    assert "Authentication failed" in capsys.readouterr().err


@patch("managed_identity_demo.cli.list_container_names")
@patch("managed_identity_demo.cli.create_blob_service_client")
@patch("managed_identity_demo.cli.create_credential")
def test_run_distinguishes_authorization_failure(
    create_credential, create_client, list_names, capsys
):
    list_names.side_effect = HttpResponseError(message="forbidden", response=Mock())

    result = run(
        [
            "--identity",
            "user",
            "--client-id",
            "identity-client-id",
            "--account-url",
            "https://example.blob.core.windows.net",
        ]
    )

    assert result == 4
    assert "data-plane role" in capsys.readouterr().err
