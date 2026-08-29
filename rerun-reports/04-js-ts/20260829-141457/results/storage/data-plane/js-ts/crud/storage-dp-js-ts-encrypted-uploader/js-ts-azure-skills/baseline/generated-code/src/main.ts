import { createAzureConfiguration } from "./config.js";
import { EncryptedBlobClient } from "./encryptedBlobClient.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const configuration = createAzureConfiguration();
  const keyManagement = new KeyManagement(
    configuration.keyClient,
    configuration.credential,
    configuration.keyName,
  );
  const encryptedBlobClient = new EncryptedBlobClient(
    configuration.containerClient,
    keyManagement,
  );

  const blobName = `encrypted-demo-${Date.now()}.bin`;
  const sample = "Hello from client-side encrypted Azure Blob Storage!";

  const upload = await encryptedBlobClient.upload(
    blobName,
    Buffer.from(sample, "utf8"),
  );
  const decrypted = await encryptedBlobClient.download(blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedDataKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  console.error("Encrypted blob round-trip failed:", error);
  process.exitCode = 1;
});
