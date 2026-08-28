import { createAzureConfiguration } from "./config.js";
import { EncryptedBlobStorage } from "./encryptedBlobStorage.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const configuration = createAzureConfiguration();
  const keyManagement = new KeyManagement(
    configuration.keyClient,
    configuration.credential,
    configuration.keyVaultUrl,
    configuration.keyName,
  );
  const encryptedBlobStorage = new EncryptedBlobStorage(
    configuration.containerClient,
    keyManagement,
  );

  const blobName =
    process.env.AZURE_STORAGE_BLOB_NAME?.trim() ||
    "sample.txt.encrypted";
  const sample = "Client-side encryption with Azure Key Vault!";

  const upload = await encryptedBlobStorage.upload(
    blobName,
    Buffer.from(sample, "utf8"),
    "application/octet-stream",
  );
  const decrypted = await encryptedBlobStorage.download(blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Round-trip failed: ${message}`);
  process.exitCode = 1;
});
