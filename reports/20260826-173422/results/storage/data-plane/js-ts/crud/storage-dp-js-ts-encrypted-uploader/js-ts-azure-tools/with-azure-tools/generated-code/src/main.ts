import { createAppConfiguration } from "./config.js";
import { EncryptedBlobStorage } from "./encrypted-blob-storage.js";
import { KeyManagement } from "./key-management.js";

async function main(): Promise<void> {
  const configuration = createAppConfiguration();
  const keyManagement = new KeyManagement(
    configuration.keyClient,
    configuration.credential,
    configuration.keyName,
  );
  const encryptedBlobStorage = new EncryptedBlobStorage(
    configuration.containerClient,
    keyManagement,
  );

  const blobName = `envelope-encryption-demo-${Date.now()}.txt`;
  const sample = Buffer.from(
    "Client-side encryption with Azure Key Vault and Blob Storage.",
    "utf8",
  );

  try {
    const upload = await encryptedBlobStorage.uploadBuffer(
      blobName,
      sample,
      "text/plain; charset=utf-8",
    );
    const decrypted = await encryptedBlobStorage.downloadBuffer(blobName);

    console.log(`Vault key ID: ${upload.keyId}`);
    console.log(`Wrapped DEK (base64): ${upload.wrappedKeyBase64}`);
    console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
  } finally {
    sample.fill(0);
  }
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Demo failed: ${message}`);
  process.exitCode = 1;
});
