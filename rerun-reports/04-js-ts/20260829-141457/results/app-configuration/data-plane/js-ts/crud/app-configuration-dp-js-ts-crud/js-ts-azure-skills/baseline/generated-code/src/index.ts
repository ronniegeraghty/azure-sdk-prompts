import {
  AppConfigurationClient,
  featureFlagContentType,
} from "@azure/app-configuration";
import { RestError } from "@azure/core-rest-pipeline";

const settingKey = "app:Settings:FontSize";
const featureFlagKey = ".appconfig.featureflag/BetaFeature";

async function main(): Promise<void> {
  const connectionString = process.env.AZURE_APP_CONFIG_CONNECTION_STRING;
  if (!connectionString) {
    throw new Error(
      "Set AZURE_APP_CONFIG_CONNECTION_STRING before running this program.",
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
  console.log(`Value for ${settingKey}: ${setting.value}`);

  for await (const matchingSetting of client.listConfigurationSettings({
    keyFilter: "app:Settings:*",
  })) {
    console.log(
      `${matchingSetting.key} [${matchingSetting.label ?? "no label"}] = ${
        matchingSetting.value ?? ""
      }`,
    );
  }

  await client.setConfigurationSetting({
    key: featureFlagKey,
    contentType: featureFlagContentType,
    value: JSON.stringify({
      id: "BetaFeature",
      description: "Enables the beta feature.",
      enabled: false,
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
    const requestId = error.response?.headers.get("x-ms-request-id");
    if (requestId) {
      console.error(`Request ID: ${requestId}`);
    }
  } else if (error instanceof Error) {
    console.error(`Error: ${error.message}`);
  } else {
    console.error("An unknown error occurred.", error);
  }

  process.exitCode = 1;
});
