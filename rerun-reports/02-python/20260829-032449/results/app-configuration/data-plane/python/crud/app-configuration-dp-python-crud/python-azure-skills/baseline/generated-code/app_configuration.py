import os

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import HttpResponseError


CONNECTION_STRING_ENV = "AZURE_APP_CONFIG_CONNECTION_STRING"
SETTING_KEY = "app:Settings:FontSize"


def main() -> None:
    connection_string = os.environ.get(CONNECTION_STRING_ENV)
    if not connection_string:
        raise RuntimeError(
            f"Set the {CONNECTION_STRING_ENV} environment variable before running."
        )

    client = AzureAppConfigurationClient.from_connection_string(connection_string)

    try:
        client.set_configuration_setting(key=SETTING_KEY, value="24")
        client.set_configuration_setting(
            key=SETTING_KEY,
            label="Production",
            value="24",
        )

        setting = client.get_configuration_setting(key=SETTING_KEY)
        print(setting.value)

        for matching_setting in client.list_configuration_settings(
            key_filter="app:Settings:*"
        ):
            print(
                f"{matching_setting.key}={matching_setting.value} "
                f"(label={matching_setting.label!r})"
            )

        client.set_configuration_setting(
            FeatureFlagConfigurationSetting(
                feature_id="BetaFeature",
                enabled=True,
            )
        )

        client.delete_configuration_setting(key=SETTING_KEY)
    except HttpResponseError as error:
        print(
            f"Azure App Configuration request failed "
            f"(status={error.status_code}): {error.message}"
        )
        raise


if __name__ == "__main__":
    main()
