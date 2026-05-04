# Key Vault Secrets Client

A minimal Azure Key Vault client for reading and writing secrets.

## Setup

```bash
pip install -r requirements.txt
```

## Configuration

Set the following environment variable before running:

```bash
export AZURE_KEYVAULT_URL="https://<your-vault-name>.vault.azure.net/"
```

Authentication uses `DefaultAzureCredential` — make sure you are logged
in via `az login` or have a service principal configured.

## Usage

```bash
python main.py
```
