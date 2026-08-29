import { AppConfigurationClient } from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";
const productionLabel = "Production";
const featureFlagId = "BetaFeature";
const featureFlagKey = `.appconfig.featureflag/${featureFlagId}`;
const featureFlagContentType =
  "application/vnd.microsoft.appconfig.ff+json;charset=utf-8";

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
    label: productionLabel,
  });

  const setting = await client.getConfigurationSetting({ key: settingKey });
  console.log(`${setting.key} = ${setting.value}`);

  console.log('Settings matching "app:Settings:*":');
  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    console.log(
      `${matchingSetting.key}${
        matchingSetting.label ? ` [${matchingSetting.label}]` : ""
      } = ${matchingSetting.value}`,
    );
  }

  await client.setConfigurationSetting({
    key: featureFlagKey,
    contentType: featureFlagContentType,
    value: JSON.stringify({
      id: featureFlagId,
      description: "Enables the beta feature.",
      enabled: true,
      conditions: {
        client_filters: [],
      },
    }),
  });

  await client.deleteConfigurationSetting({ key: settingKey });
  console.log(`Deleted ${settingKey}.`);
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
