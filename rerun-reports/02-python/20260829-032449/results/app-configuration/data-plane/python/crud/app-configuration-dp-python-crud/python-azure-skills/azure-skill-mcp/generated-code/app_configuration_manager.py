import logging
import os
import sys

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    ConfigurationSetting,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import ClientAuthenticationError, HttpResponseError
from azure.identity import DefaultAzureCredential


SETTING_KEY = "app:Settings:FontSize"
SETTING_VALUE = "24"
PRODUCTION_LABEL = "Production"
FEATURE_ID = "BetaFeature"


def create_client(
    endpoint: str, credential: DefaultAzureCredential
) -> AzureAppConfigurationClient:
    return AzureAppConfigurationClient(base_url=endpoint, credential=credential)


def manage_settings(client: AzureAppConfigurationClient) -> None:
    client.set_configuration_setting(
        ConfigurationSetting(key=SETTING_KEY, value=SETTING_VALUE)
    )
    client.set_configuration_setting(
        ConfigurationSetting(
            key=SETTING_KEY,
            value=SETTING_VALUE,
            label=PRODUCTION_LABEL,
        )
    )

    setting = client.get_configuration_setting(key=SETTING_KEY)
    print(setting.value)

    for matching_setting in client.list_configuration_settings(
        key_filter="app:Settings:*"
    ):
        print(
            f"{matching_setting.key} "
            f"(label={matching_setting.label!r}): {matching_setting.value}"
        )

    client.set_configuration_setting(
        FeatureFlagConfigurationSetting(feature_id=FEATURE_ID, enabled=True)
    )

    client.delete_configuration_setting(key=SETTING_KEY)


def main() -> int:
    endpoint = os.environ.get("AZURE_APPCONFIG_ENDPOINT")
    if not endpoint:
        logging.error(
            "Set AZURE_APPCONFIG_ENDPOINT to the App Configuration endpoint."
        )
        return 2

    try:
        with DefaultAzureCredential() as credential:
            with create_client(endpoint, credential) as client:
                manage_settings(client)
    except ClientAuthenticationError as error:
        logging.error("Azure authentication failed: %s", error.message)
        return 1
    except HttpResponseError as error:
        request_id = (
            error.response.headers.get("x-ms-request-id")
            if error.response is not None
            else None
        )
        logging.error(
            "App Configuration request failed (status=%s, request_id=%s): %s",
            error.status_code,
            request_id or "unknown",
            error.message,
        )
        return 1

    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    sys.exit(main())
