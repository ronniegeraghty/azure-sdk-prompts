import { createAppConfig } from "./config.js";
import { EncryptedBlobStorage } from "./encryptedBlobStorage.js";
import { EncryptedBlobError } from "./errors.js";
import { KeyVaultKeyManager } from "./keyManagement.js";

async function main(): Promise<void> {
  const config = createAppConfig();
  const keyManager = new KeyVaultKeyManager(
    config.keyClient,
    config.credential,
    config.keyName,
  );
  const encryptedStorage = new EncryptedBlobStorage(
    config.containerClient,
    keyManager,
  );
  const sample = "Client-side encryption with Azure Key Vault envelope keys.";

  const upload = await encryptedStorage.upload(config.blobName, sample);
  const decrypted = await encryptedStorage.download(config.blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedDataKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  if (error instanceof EncryptedBlobError) {
    console.error(
      `[${error.category}] ${error.operation}: ${error.message}`,
    );
  } else {
    console.error(error instanceof Error ? error.message : error);
  }

  process.exitCode = 1;
});
