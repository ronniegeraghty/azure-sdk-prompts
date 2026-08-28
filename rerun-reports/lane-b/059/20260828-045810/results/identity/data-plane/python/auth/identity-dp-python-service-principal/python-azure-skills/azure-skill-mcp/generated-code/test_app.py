"""Local-only tests for configuration and Azure error handling."""

from __future__ import annotations

import os
import unittest
from unittest.mock import MagicMock, patch

from azure.core.exceptions import ClientAuthenticationError

import app


class AzureSettingsTests(unittest.TestCase):
    @patch.dict(os.environ, {}, clear=True)
    def test_missing_environment_variables_are_reported(self) -> None:
        with self.assertRaisesRegex(
            app.ConfigurationError, "AZURE_CLIENT_SECRET"
        ):
            app.AzureSettings.from_environment()


class RunTests(unittest.TestCase):
    SETTINGS = {
        "AZURE_TENANT_ID": "tenant",
        "AZURE_CLIENT_ID": "client",
        "AZURE_CLIENT_SECRET": "secret",
        "AZURE_SUBSCRIPTION_ID": "subscription",
    }

    @patch.dict(os.environ, SETTINGS, clear=True)
    @patch("app.list_resource_group_names", return_value=["example-rg"])
    def test_success(self, list_names: MagicMock) -> None:
        with patch("app.create_credential", return_value=MagicMock()):
            self.assertEqual(app.run(), 0)
        list_names.assert_called_once()

    @patch.dict(os.environ, SETTINGS, clear=True)
    @patch(
        "app.list_resource_group_names",
        side_effect=ClientAuthenticationError("invalid credential"),
    )
    def test_authentication_failure_has_distinct_exit_code(
        self, list_names: MagicMock
    ) -> None:
        with patch("app.create_credential", return_value=MagicMock()):
            self.assertEqual(app.run(), 3)
        list_names.assert_called_once()


if __name__ == "__main__":
    unittest.main()
