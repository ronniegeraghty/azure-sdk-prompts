---
name: azure-sdk-for-rust-bestpractices
description: "Best practices for generating Rust code that uses the Azure SDK for Rust. Use when generating Rust code targeting Azure services to ensure correct crate selection, authentication patterns, error handling, pagination, and idiomatic SDK usage."
applyTo: "**/*.rs,**/Cargo.toml"
---

# Azure SDK for Rust Best Practices

You are generating Rust code that uses the **Azure SDK for Rust**. In addition to the general Azure best practices, these best practices demonstrate how to produce correct code which uses the Azure SDK for Rust effectively.

## References

- [Azure SDK for Rust repository README](https://github.com/Azure/azure-sdk-for-rust/blob/main/README.md)
- [azure_core package README](https://github.com/Azure/azure-sdk-for-rust/blob/main/sdk/core/azure_core/README.md)
- [Azure SDK Design Guidelines for Rust](https://azure.github.io/azure-sdk/rust_introduction.html)

## Crate Selection

Use only the **official** `azure_*` crates published by the Azure SDK team. Do NOT use legacy or unofficial crates.

### Official Crates

| Service           | Crate                     |
| ----------------- | ------------------------- |
| Core              | `azure_core`              |
| Identity          | `azure_identity`          |
| Storage Blobs     | `azure_storage_blob`      |
| Storage Common    | `azure_storage`           |
| Key Vault Secrets | `azure_key_vault_secrets` |
| Cosmos            | `azure_cosmos`            |

### Banned Legacy Crates (NEVER use)

- `azure_storage_blobs` (plural — unofficial)
- `azure_storage_queues`
- `azure_storage_datalake`
- `azure_data_cosmos`
- `azure_security_keyvault`
- `azure_messaging_servicebus`

## Installing Crates

Use `cargo add` to install crates:

```sh
cargo add azure_identity azure_security_keyvault_secrets tokio
```

## Authentication

Always use token-based credentials from `azure_identity`. Never use connection strings, shared access signatures, or hardcoded keys.

### Preferred Credential Types

- `DeveloperToolsCredential` — recommended for development, NOT appropriate for production.
- `AzureCliCredential`
- `EnvironmentCredential`
- `ManagedIdentityCredential`
- `ClientSecretCredential`
- `WorkloadIdentityCredential`

### Example: Creating a Client with Credentials

```rust
use azure_identity::DeveloperToolsCredential;
use azure_security_keyvault_secrets::{SecretClient, SecretClientOptions};

let credential = DeveloperToolsCredential::new(None)?;

let options = SecretClientOptions {
    api_version: "7.5".to_string(),
    ..Default::default()
};

let client = SecretClient::new(
    "https://<your-key-vault-name>.vault.azure.net/",
    credential.clone(),
    Some(options),
)?;
```

## Service Client Configuration

Client types have names ending in `Client` (e.g., `SecretClient`, `BlobClient`). Instantiate them with `new()`, passing the endpoint, credential, and optional `ClientOptions`.

- Always use `..Default::default()` in struct initialization to protect against future breaking changes when fields are added.
- Annotate with `#[allow(clippy::needless_update)]` if you assign all fields.

## Error Handling

Use `Result<T, E>` and the `?` operator. Avoid `.unwrap()` except in trivial examples with clear context.

### Pattern: Matching on Error Kind

```rust
use azure_core::{error::ErrorKind, http::StatusCode};

match client.get_secret("secret-name", None).await {
    Ok(secret) => println!("Secret: {:?}", secret.into_model()?.value),
    Err(e) => match e.kind() {
        ErrorKind::HttpResponse { status, error_code, .. } if *status == StatusCode::NotFound => {
            if let Some(code) = error_code {
                println!("ErrorCode: {}", code);
            } else {
                println!("Secret not found, but no error code provided.");
            }
        },
        _ => println!("An error occurred: {e:?}"),
    },
}
```

## Accessing HTTP Response Details

Service methods return `Response<T>`. Use `into_model()` to deserialize, or `deconstruct()` to access status, headers, and body separately.

```rust
let response = client.get_secret("secret-name", None).await?;
let secret = response.into_model()?;

// Or access raw HTTP details:
let response = client.get_secret("secret-name", None).await?;
let (status, headers, body) = response.deconstruct();
```

## Pagination with `Pager<T>`

When a service function returns a `Pager<T>` type collection, iterate over the pager directly (using `futures::TryStreamExt`). `Pager<T>` automatically flattens pages into individual items via the `Page` trait — even when `T` is a response envelope like `ListBlobsResponse`, calling `try_next()` yields individual items (e.g., `BlobItem`), **not** pages. Do **not** use `into_pages()` unless you have an explicit requirement for page-level access.

### Example: Listing Key Vault secrets

```rust
use azure_security_keyvault_secrets::ResourceExt;
use futures::TryStreamExt as _;

let mut pager = client.list_secret_properties(None)?;
while let Some(secret) = pager.try_next().await? {
    let name = secret.resource_id()?.name;
    println!("Found secret with name: {}", name);
}
```

### Example: Listing blobs in a container

```rust
use futures::TryStreamExt as _;

let mut pager = container_client.list_blobs(None)?;
while let Some(blob) = pager.try_next().await? {
    let name = blob.name.as_deref().unwrap_or("<unknown>");
    let size = blob
        .properties
        .as_ref()
        .and_then(|p| p.content_length)
        .unwrap_or(0);
    println!("Blob: {name} ({size} bytes)");
}
```

Page-by-page iteration via `into_pages()` is only necessary when there is an explicit requirement for page-level access.

If you MUST iterate over the pages, using `into_pages()`, make sure you handle the nested structure correctly:

```rust
let mut pager = client.list_secret_properties(None)?.into_pages();
while let Some(secrets) = pager.try_next().await? {
    let secrets = secrets.into_model()?.value;
    for secret in secrets {
        let name = secret.resource_id()?.name;
        println!("Found secret with name: {}", name);
    }
}
```

## Long-Running Operations with `Poller<T>`

`Poller<T>` implements `IntoFuture`, so you can simply `.await` it:

```rust
let certificate = client
    .create_certificate("certificate-name", body.try_into()?, None)?
    .await?
    .into_model()?;
```

Or iterate over status updates using `futures::Stream`:

```rust
use futures::stream::TryStreamExt as _;

let mut poller = client.create_certificate("certificate-name", body.try_into()?, None)?;
while let Some(operation) = poller.try_next().await? {
    let operation = operation.into_model()?;
    match operation.status.as_deref().unwrap_or("unknown") {
        "inProgress" => continue,
        "completed" => {
            let target = operation.target.ok_or("expected target")?;
            println!("Created certificate {}", target);
            break;
        },
        status => Err(format!("operation terminated with status {status}"))?,
    }
}
```

## Async Runtime

Use `#[tokio::main]` as the default async runtime. If using a different runtime, disable default features and provide a custom `AsyncRuntime` implementation via `azure_core::async_runtime::set_async_runtime()`.

## Logging and Debugging

- Models implement `core::fmt::Debug` with PII-safe formatting by default (prints `Secret { .. }`).
- Enable the `debug` feature on `azure_core` for full field output (NOT recommended for production).

## Feature Flags

Key `azure_core` features:

- `reqwest` (default): sets `reqwest` as the HTTP client
- `tokio` (default): sets `tokio` as the async runtime
- `reqwest_gzip` / `reqwest_deflate` (default): compression support
- `reqwest_native_tls` (default): native TLS
- `xml`: XML support
- `hmac_rust` / `hmac_openssl`: HMAC signing
- `derive`: derive macros like `SafeDebug`

## Custom HTTP Policies

Add per-call or per-try policies by implementing `Policy` and adding to `ClientOptions`:

```rust
use azure_core::http::{
    policies::{Policy, PolicyResult},
    ClientOptions, Context, Request,
};
use std::sync::Arc;

#[derive(Debug)]
struct MyPolicy;

#[async_trait::async_trait]
impl Policy for MyPolicy {
    async fn send(
        &self,
        ctx: &Context,
        request: &mut Request,
        next: &[Arc<dyn Policy>],
    ) -> PolicyResult {
        // Modify request here
        next[0].send(ctx, request, &next[1..]).await
    }
}

let mut options = SecretClientOptions::default();
options.client_options.per_call_policies.push(Arc::new(MyPolicy));
```

## Known Issues

If you encounter hangs when making multiple HTTP operations, disable connection pooling:

```rust
let client = Arc::new(
    reqwest::ClientBuilder::new()
        .pool_max_idle_per_host(0)
        .build()?,
);
```

## Summary Checklist

- [ ] Use official `azure_*` crates only
- [ ] Authenticate with `azure_identity` token credentials
- [ ] Handle errors with `Result` and `?` operator
- [ ] Use `..Default::default()` in struct initialization
- [ ] Use `Pager<T>` with `TryStreamExt` for pagination
- [ ] Use `Poller<T>` with `.await` or stream for LROs
- [ ] Use `#[tokio::main]` async runtime
- [ ] Never hardcode secrets, keys, or connection strings
