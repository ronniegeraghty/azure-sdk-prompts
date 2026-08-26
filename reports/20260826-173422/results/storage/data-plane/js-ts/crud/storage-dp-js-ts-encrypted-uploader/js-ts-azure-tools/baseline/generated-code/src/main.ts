import { createAzureConnections } from "./config.js";
import { EncryptedBlobClient } from "./encryptedBlob.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const connections = createAzureConnections();
  const keyManagement = new KeyManagement(
    connections.keyClient,
    connections.credential,
    connections.keyName,
  );
  const encryptedBlobs = new EncryptedBlobClient(
    connections.containerClient,
    keyManagement,
  );

  const blobName = `encrypted-demo-${Date.now()}.bin`;
  const sample = Buffer.from(
    "Client-side encryption with an Azure Key Vault protected data key.",
    "utf8",
  );

  const upload = await encryptedBlobs.upload(blobName, sample);
  const decrypted = await encryptedBlobs.download(blobName);

  console.log(`Vault key ID: ${upload.keyId}`);
  console.log(`Wrapped DEK (base64): ${upload.wrappedKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  console.error("Encrypted blob round-trip failed:", error);
  process.exitCode = 1;
});
