import {
  AppConfigurationClient,
  type FeatureFlagValue,
  type SetConfigurationSettingParam,
  featureFlagContentType,
  featureFlagPrefix,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "Set the AZURE_APPCONFIG_CONNECTION_STRING environment variable.",
    );
  }

  const client = new AppConfigurationClient(connectionString);

  await client.setConfigurationSetting({
    key: settingKey,
    value: "24",
  });

  await client.setConfigurationSetting({
    key: settingKey,
    value: "24",
    label: "Production",
  });

  const setting = await client.getConfigurationSetting({ key: settingKey });
  console.log(`${setting.key} = ${setting.value}`);

  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    const label = matchingSetting.label ?? "(no label)";
    console.log(
      `${matchingSetting.key} [${label}] = ${matchingSetting.value ?? ""}`,
    );
  }

  const betaFeature: SetConfigurationSettingParam<FeatureFlagValue> = {
    key: `${featureFlagPrefix}BetaFeature`,
    contentType: featureFlagContentType,
    value: {
      id: "BetaFeature",
      enabled: true,
      description: "Enables the beta experience.",
      conditions: {
        clientFilters: [],
      },
    },
  };

  await client.setConfigurationSetting(betaFeature);
  await client.deleteConfigurationSetting({ key: settingKey });
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
