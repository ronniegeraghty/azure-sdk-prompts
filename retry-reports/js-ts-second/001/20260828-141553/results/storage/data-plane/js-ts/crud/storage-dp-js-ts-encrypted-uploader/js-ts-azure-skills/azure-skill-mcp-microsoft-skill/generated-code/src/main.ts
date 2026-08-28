import { createAzureClients } from "./config.js";
import { EncryptedBlobStore } from "./encryptedBlobStore.js";
import { describeServiceError } from "./errors.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const clients = createAzureClients();
  const keyManagement = new KeyManagement(
    clients.keyClient,
    clients.credential,
    clients.keyVaultUrl,
    clients.keyName,
  );
  const encryptedBlobStore = new EncryptedBlobStore(
    clients.containerClient,
    keyManagement,
  );

  const blobName = `round-trip-${Date.now()}.txt`;
  const sample = Buffer.from(
    "Client-side encryption keeps this plaintext out of Azure services.",
    "utf8",
  );

  const upload = await encryptedBlobStore.upload(blobName, sample);
  const decrypted = await encryptedBlobStore.download(blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  console.error(`Round-trip failed: ${describeServiceError(error)}`);
  process.exitCode = 1;
});
