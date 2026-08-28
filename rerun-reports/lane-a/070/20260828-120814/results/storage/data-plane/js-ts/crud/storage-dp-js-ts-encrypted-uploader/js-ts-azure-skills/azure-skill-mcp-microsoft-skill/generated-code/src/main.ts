import { buildAzureConnections } from "./config.js";
import { EncryptedBlobClient } from "./encryptedBlobClient.js";
import { KeyManagement } from "./keyManagement.js";

async function main(): Promise<void> {
  const connections = buildAzureConnections();
  const keyManagement = new KeyManagement(
    connections.keyClient,
    connections.credential,
    connections.vaultUrl,
    connections.keyName,
  );
  const encryptedBlobClient = new EncryptedBlobClient(
    connections.containerClient,
    keyManagement,
  );

  const blobName = process.env.DEMO_BLOB_NAME?.trim() || "encrypted-demo.txt";
  const sampleText =
    process.env.DEMO_TEXT ??
    "Hello from client-side envelope encryption.";

  const uploadResult = await encryptedBlobClient.upload(
    blobName,
    Buffer.from(sampleText, "utf8"),
    "text/plain; charset=utf-8",
  );
  const decrypted = await encryptedBlobClient.download(blobName);

  console.log(`Vault key ID: ${uploadResult.keyId}`);
  console.log(`Wrapped DEK (base64): ${uploadResult.wrappedDataKeyBase64}`);
  console.log(`Decrypted output: ${decrypted.toString("utf8")}`);
}

main().catch((error: unknown) => {
  const message = error instanceof Error ? error.message : String(error);
  console.error(`Round-trip failed: ${message}`);
  process.exitCode = 1;
});
