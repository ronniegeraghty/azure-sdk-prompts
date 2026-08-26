import logging
import os

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
KEY_FILTER = "app:Settings:*"


def main() -> int:
    endpoint = os.environ.get("AZURE_APPCONFIGURATION_ENDPOINT")
    if not endpoint:
        raise RuntimeError(
            "Set AZURE_APPCONFIGURATION_ENDPOINT to your App Configuration endpoint."
        )

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
            print(f"{setting.key} = {setting.value}")

            print(f"Settings matching {KEY_FILTER!r}:")
            for matching_setting in client.list_configuration_settings(
                key_filter=KEY_FILTER
            ):
                label = matching_setting.label or "(no label)"
                print(
                    f"  {matching_setting.key} [{label}] = "
                    f"{matching_setting.value}"
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
        logging.error(
            "Azure App Configuration request failed (HTTP %s): %s",
            status_code,
            error.message,
        )
        return 1
    finally:
        credential.close()

    return 0


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    raise SystemExit(main())
