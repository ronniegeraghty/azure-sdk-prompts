# Azure Resource Group Manager

This local Python CLI uses the Azure management plane SDK to create, list, get,
tag, and delete resource groups. It does not contain credentials or execute any
operation until you run it.

## Install

Python 3.9 or newer is required.

```powershell
python -m pip install -r requirements.txt
```

Required packages:

- `azure-identity` for Microsoft Entra authentication
- `azure-mgmt-resource` for Resource Group management

Set the subscription ID and configure any authentication method supported by
`DefaultAzureCredential`, such as service principal environment variables or a
managed identity:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
$env:AZURE_TENANT_ID = "<tenant-id>"
$env:AZURE_CLIENT_ID = "<client-id>"
$env:AZURE_CLIENT_SECRET = "<client-secret>"
```

Do not store those values in source control.

## Usage

Run individual operations:

```powershell
python resource_group_manager.py create --name example-rg --location eastus
python resource_group_manager.py list
python resource_group_manager.py get --name example-rg
python resource_group_manager.py tag --name example-rg --key environment --value dev
python resource_group_manager.py delete --name example-rg --yes
```

Run the complete create, list, get, tag, and delete lifecycle:

```powershell
python resource_group_manager.py workflow --name example-rg --location eastus --yes
```

Deletion removes the resource group and every resource it contains. Both
deletion paths require `--yes`.
