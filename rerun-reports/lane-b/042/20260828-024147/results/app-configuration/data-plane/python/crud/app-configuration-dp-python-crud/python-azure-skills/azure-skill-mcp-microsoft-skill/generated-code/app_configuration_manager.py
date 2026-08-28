import os
import sys

from azure.appconfiguration import (
    AzureAppConfigurationClient,
    ConfigurationSetting,
    FeatureFlagConfigurationSetting,
)
from azure.core.exceptions import HttpResponseError
from azure.identity import DefaultAzureCredential


SETTING_KEY = "app:Settings:FontSize"
SETTING_VALUE = "24"
PRODUCTION_LABEL = "Production"
FEATURE_ID = "BetaFeature"


def main() -> int:
    endpoint = os.getenv("AZURE_APPCONFIGURATION_ENDPOINT")
    if not endpoint:
        print(
            "Set AZURE_APPCONFIGURATION_ENDPOINT to an App Configuration "
            "endpoint, such as https://<name>.azconfig.io.",
            file=sys.stderr,
        )
        return 2

    credential = DefaultAzureCredential()

    try:
        with AzureAppConfigurationClient(
            base_url=endpoint,
            credential=credential,
        ) as client:
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

            client.delete_configuration_setting(key=SETTING_KEY)
    except HttpResponseError as error:
        status = (
            f" (HTTP {error.status_code})"
            if error.status_code is not None
            else ""
        )
        print(
            f"Azure App Configuration request failed{status}: {error.message}",
            file=sys.stderr,
        )
        return 1
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
