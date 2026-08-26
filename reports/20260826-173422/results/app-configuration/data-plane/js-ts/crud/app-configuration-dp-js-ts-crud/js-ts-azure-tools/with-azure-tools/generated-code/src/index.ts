import {
  AppConfigurationClient,
  type ConfigurationSettingParam,
  type FeatureFlagValue,
  featureFlagContentType,
  featureFlagPrefix,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "AZURE_APPCONFIG_CONNECTION_STRING must be set before running the program.",
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

  console.log("Matching configuration settings:");
  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    const label = matchingSetting.label ?? "(no label)";
    console.log(
      `- ${matchingSetting.key} [${label}] = ${matchingSetting.value}`,
    );
  }

  const betaFeatureFlag: ConfigurationSettingParam<FeatureFlagValue> = {
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

  await client.setConfigurationSetting(betaFeatureFlag);
  console.log("Created feature flag BetaFeature.");

  await client.deleteConfigurationSetting({ key: settingKey });
  console.log(`Deleted the unlabeled setting ${settingKey}.`);
}

main().catch((error: unknown) => {
  if (error instanceof RestError) {
    console.error("Azure App Configuration request failed.", {
      code: error.code,
      statusCode: error.statusCode,
      message: error.message,
      requestId: error.request?.headers.get("x-ms-request-id"),
    });
  } else if (error instanceof Error) {
    console.error(error.message);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
});
