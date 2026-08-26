import {
  AppConfigurationClient,
  type ConfigurationSetting,
  type FeatureFlagValue,
  featureFlagContentType,
  featureFlagPrefix,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";
const settingValue = "24";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APP_CONFIGURATION_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "Set AZURE_APP_CONFIGURATION_CONNECTION_STRING before running this program.",
    );
  }

  const client = new AppConfigurationClient(connectionString);

  await client.setConfigurationSetting({
    key: settingKey,
    value: settingValue,
  });

  await client.setConfigurationSetting({
    key: settingKey,
    value: settingValue,
    label: "Production",
  });

  const setting = await client.getConfigurationSetting({ key: settingKey });
  console.log(`${setting.key} = ${setting.value}`);

  console.log("Matching configuration settings:");
  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    console.log(
      `${matchingSetting.key} [${matchingSetting.label ?? "no label"}] = ${matchingSetting.value}`,
    );
  }

  const betaFeature: ConfigurationSetting<FeatureFlagValue> = {
    key: `${featureFlagPrefix}BetaFeature`,
    isReadOnly: false,
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
  await client.setConfigurationSetting(betaFeature);

  await client.deleteConfigurationSetting({ key: settingKey });
  console.log(`Deleted the unlabeled setting "${settingKey}".`);
}

main().catch((error: unknown) => {
  if (error instanceof RestError) {
    console.error(
      `Azure App Configuration request failed (${error.statusCode ?? "no status"}): ${error.message}`,
    );
  } else if (error instanceof Error) {
    console.error(`Program failed: ${error.message}`);
  } else {
    console.error("Program failed with an unknown error.");
  }

  process.exitCode = 1;
});
