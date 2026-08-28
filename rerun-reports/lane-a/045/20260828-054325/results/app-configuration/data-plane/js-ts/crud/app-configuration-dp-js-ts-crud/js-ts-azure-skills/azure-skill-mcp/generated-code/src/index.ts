import {
  AppConfigurationClient,
  ConfigurationSetting,
  FeatureFlagValue,
  featureFlagContentType,
  featureFlagPrefix,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APPCONFIG_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "Set the AZURE_APPCONFIG_CONNECTION_STRING environment variable before running this program.",
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

  console.log("Matching settings:");
  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    const label = matchingSetting.label ?? "(no label)";
    console.log(
      `- ${matchingSetting.key} [${label}] = ${matchingSetting.value}`,
    );
  }

  const betaFeature: ConfigurationSetting<FeatureFlagValue> = {
    key: `${featureFlagPrefix}BetaFeature`,
    isReadOnly: false,
    contentType: featureFlagContentType,
    value: {
      id: "BetaFeature",
      enabled: true,
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
    console.error("Azure App Configuration request failed:", {
      statusCode: error.statusCode,
      code: error.code,
      message: error.message,
      requestId: error.request?.headers.get("x-ms-client-request-id"),
    });
  } else if (error instanceof Error) {
    console.error("Application error:", error.message);
  } else {
    console.error("Unexpected error:", error);
  }

  process.exitCode = 1;
});
