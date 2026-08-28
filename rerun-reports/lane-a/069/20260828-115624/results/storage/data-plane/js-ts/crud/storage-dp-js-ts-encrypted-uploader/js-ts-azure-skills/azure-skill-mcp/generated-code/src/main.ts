import { createAzureConnections } from "./config.js";
import { EncryptedBlobClient } from "./encryptedBlobClient.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const connections = createAzureConnections();
  const keyManagement = await KeyManagement.create(
    connections.keyClient,
    connections.credential,
    connections.keyName,
    connections.keyVersion,
  );
  const encryptedBlobs = new EncryptedBlobClient(
    connections.containerClient,
    keyManagement,
  );

  const blobName = `encrypted-demo-${Date.now()}.bin`;
  const sample = "Client-side encryption with Azure Key Vault envelope encryption.";

  const upload = await encryptedBlobs.upload(blobName, sample);
  const decrypted = await encryptedBlobs.download(blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Round-trip failed: ${message}`);
  process.exitCode = 1;
});
