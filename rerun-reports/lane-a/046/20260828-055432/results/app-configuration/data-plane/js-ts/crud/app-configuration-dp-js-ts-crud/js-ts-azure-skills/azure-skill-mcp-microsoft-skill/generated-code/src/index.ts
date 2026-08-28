import {
  AppConfigurationClient,
  type FeatureFlagValue,
  type SetConfigurationSettingParam,
  featureFlagContentType,
  featureFlagPrefix,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const key = "app:Settings:FontSize";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "AZURE_APPCONFIG_CONNECTION_STRING must contain an Azure App Configuration connection string.",
    );
  }

  const client = new AppConfigurationClient(connectionString);

  await client.setConfigurationSetting({ key, value: "24" });

  await client.setConfigurationSetting({
    key,
    value: "24",
    label: "Production",
  });

  const setting = await client.getConfigurationSetting({ key });
  console.log(`${setting.key} = ${setting.value}`);

  console.log('Settings matching "app:Settings:*":');
  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    const label = matchingSetting.label ?? "(no label)";
    console.log(
      `${matchingSetting.key} [${label}] = ${matchingSetting.value ?? ""}`,
    );
  }

  const featureFlag: SetConfigurationSettingParam<FeatureFlagValue> = {
    key: `${featureFlagPrefix}BetaFeature`,
    contentType: featureFlagContentType,
    value: {
      id: "BetaFeature",
      enabled: true,
      description: "Enables the beta feature.",
      conditions: {
        clientFilters: [],
      },
    },
  };

  await client.setConfigurationSetting(featureFlag);
  await client.deleteConfigurationSetting({ key });
}

main().catch((error: unknown) => {
  if (error instanceof RestError) {
    console.error(
      `Azure App Configuration request failed (${error.statusCode ?? "unknown status"}): ${error.message}`,
    );
    if (error.code) {
      console.error(`Error code: ${error.code}`);
    }
  } else if (error instanceof Error) {
    console.error(`Application error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
});
