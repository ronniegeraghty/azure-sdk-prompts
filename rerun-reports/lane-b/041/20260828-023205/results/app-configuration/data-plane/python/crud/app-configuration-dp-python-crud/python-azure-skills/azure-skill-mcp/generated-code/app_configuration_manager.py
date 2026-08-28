import os
import sys

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    ConfigurationSetting,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import HttpResponseError


CONNECTION_STRING_ENV = "AZURE_APPCONFIG_CONNECTION_STRING"
SETTING_KEY = "app:Settings:FontSize"
SETTING_VALUE = "24"
PRODUCTION_LABEL = "Production"
KEY_FILTER = "app:Settings:*"


def main() -> int:
    connection_string = os.environ.get(CONNECTION_STRING_ENV)
    if not connection_string:
        print(
            f"Set the {CONNECTION_STRING_ENV} environment variable before running.",
            file=sys.stderr,
        )
        return 2

    client = AzureAppConfigurationClient.from_connection_string(connection_string)

    try:
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
        print(f"{setting.key} = {setting.value}")

        for matching_setting in client.list_configuration_settings(
            key_filter=KEY_FILTER
        ):
            label = matching_setting.label or "<no label>"
            print(
                f"{matching_setting.key} [{label}] = {matching_setting.value}"
            )

        client.set_configuration_setting(
            FeatureFlagConfigurationSetting(
                feature_id="BetaFeature",
                enabled=True,
            )
        )

        client.delete_configuration_setting(key=SETTING_KEY)
    except HttpResponseError as error:
        status_code = error.status_code or "unknown"
        print(
            f"Azure App Configuration request failed "
            f"(status {status_code}): {error.message}",
            file=sys.stderr,
        )
        return 1
    finally:
        client.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
