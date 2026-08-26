import os
import sys

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import HttpResponseError


CONNECTION_STRING_ENV = "AZURE_APP_CONFIG_CONNECTION_STRING"
SETTING_KEY = "app:Settings:FontSize"
SETTING_VALUE = "24"
PRODUCTION_LABEL = "Production"


def create_client() -> AzureAppConfigurationClient:
    connection_string = os.environ.get(CONNECTION_STRING_ENV)
    if not connection_string:
        raise RuntimeError(
            f"Set the {CONNECTION_STRING_ENV} environment variable before running."
        )

    return AzureAppConfigurationClient.from_connection_string(connection_string)


def main() -> int:
    try:
        client = create_client()
    except RuntimeError as error:
        print(f"Configuration error: {error}", file=sys.stderr)
        return 2

    try:
        with client:
            client.set_configuration_setting(key=SETTING_KEY, value=SETTING_VALUE)
            client.set_configuration_setting(
                key=SETTING_KEY,
                value=SETTING_VALUE,
                label=PRODUCTION_LABEL,
            )

            setting = client.get_configuration_setting(key=SETTING_KEY)
            print(f"{setting.key} = {setting.value}")

            print(f'Settings matching "app:Settings:*":')
            for matching_setting in client.list_configuration_settings(
                key_filter="app:Settings:*"
            ):
                label = matching_setting.label or "(no label)"
                print(
                    f"  {matching_setting.key} = {matching_setting.value} "
                    f"[label: {label}]"
                )

            beta_feature = FeatureFlagConfigurationSetting(
                feature_id="BetaFeature",
                enabled=True,
            )
            client.set_configuration_setting(beta_feature)

            client.delete_configuration_setting(key=SETTING_KEY)
    except HttpResponseError as error:
        status_code = (
            error.response.status_code if error.response is not None else "unknown"
        )
        print(
            f"Azure App Configuration request failed "
            f"(status {status_code}): {error.message}",
            file=sys.stderr,
        )
        return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
