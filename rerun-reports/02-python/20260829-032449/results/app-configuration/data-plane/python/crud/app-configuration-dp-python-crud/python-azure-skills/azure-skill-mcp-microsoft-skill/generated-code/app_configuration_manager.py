"""Manage key-values and a feature flag in Azure App Configuration."""

import logging
import os

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    ConfigurationSetting,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential


KEY = "app:Settings:FontSize"
KEY_FILTER = "app:Settings:*"
PRODUCTION_LABEL = "Production"
FEATURE_ID = "BetaFeature"

logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
logger = logging.getLogger(__name__)


def log_http_error(error: HttpResponseError) -> None:
    """Log Azure response details that are useful for troubleshooting."""
    error_code = error.error.code if error.error else "N/A"
    request_id = (
        error.response.headers.get("x-ms-request-id") if error.response else "N/A"
    )
    logger.error(
        "Azure App Configuration request failed: status=%s code=%s "
        "request_id=%s message=%s",
        error.status_code,
        error_code,
        request_id,
        error.message,
    )


def main() -> int:
    endpoint = os.getenv("AZURE_APPCONFIGURATION_ENDPOINT")
    if not endpoint:
        logger.error("Set the AZURE_APPCONFIGURATION_ENDPOINT environment variable.")
        return 2

    try:
        with DefaultAzureCredential() as credential:
            with AzureAppConfigurationClient(
                base_url=endpoint,
                credential=credential,
            ) as client:
                client.set_configuration_setting(
                    ConfigurationSetting(key=KEY, value="24")
                )
                client.set_configuration_setting(
                    ConfigurationSetting(
                        key=KEY,
                        value="24",
                        label=PRODUCTION_LABEL,
                    )
                )

                setting = client.get_configuration_setting(key=KEY)
                print(f"{setting.key} = {setting.value}")

                print(f"Settings matching {KEY_FILTER}:")
                for matching_setting in client.list_configuration_settings(
                    key_filter=KEY_FILTER
                ):
                    label = matching_setting.label or "(no label)"
                    print(
                        f"{matching_setting.key} [{label}] = "
                        f"{matching_setting.value}"
                    )

                client.set_configuration_setting(
                    FeatureFlagConfigurationSetting(
                        feature_id=FEATURE_ID,
                        enabled=True,
                    )
                )

                client.delete_configuration_setting(key=KEY)
                print(f"Deleted {KEY} with no label.")
    except ClientAuthenticationError as error:
        logger.error("Azure authentication failed: %s", error.message)
        return 1
    except HttpResponseError as error:
        log_http_error(error)
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
