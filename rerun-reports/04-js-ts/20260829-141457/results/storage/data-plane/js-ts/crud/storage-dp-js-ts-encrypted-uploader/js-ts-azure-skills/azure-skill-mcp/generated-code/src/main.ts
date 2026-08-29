import { createAppConfig } from "./config.js";
import { EncryptedBlobClient } from "./encrypted-blob-client.js";
import { getErrorMessage } from "./errors.js";
import { KeyManagementClient } from "./key-management.js";

async function main(): Promise<void> {
  const config = createAppConfig();
  const keyManagementClient = new KeyManagementClient(
    config.keyClient,
    config.credential,
    config.keyVaultUrl,
    config.keyName,
  );
  const encryptedBlobClient = new EncryptedBlobClient(
    config.containerClient,
    keyManagementClient,
  );

  const sample = "Client-side envelope encryption with Azure Key Vault";
  const uploadResult = await encryptedBlobClient.upload(config.blobName, sample);
  const decrypted = await encryptedBlobClient.download(config.blobName);

  console.log(`Vault key ID: ${uploadResult.keyId}`);
  console.log(`Wrapped DEK (base64): ${uploadResult.wrappedKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  console.error(`Round-trip failed: ${getErrorMessage(error)}`);
  process.exitCode = 1;
});
