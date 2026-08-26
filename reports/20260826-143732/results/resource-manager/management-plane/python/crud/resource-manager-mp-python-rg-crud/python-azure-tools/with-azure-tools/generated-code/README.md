# Azure Resource Group Manager

This Python 3.10+ command-line script uses the Azure management-plane SDK and
`DefaultAzureCredential`. It does not contain credentials or a subscription ID.

## Install

```powershell
python -m venv .venv
.\.venv\Scripts\Activate.ps1
python -m pip install -r requirements.txt
```

Required packages:

- `azure-identity` for Microsoft Entra authentication
- `azure-mgmt-resource` for Azure Resource Manager operations

Set the target subscription:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
```

`DefaultAzureCredential` can use managed identity in Azure or supported local
developer credentials. Grant the identity only the Resource Group permissions
needed for the selected operations.

## Use

List all Resource Groups:

```powershell
python .\resource_groups.py list
```

Get details of a Resource Group:

```powershell
python .\resource_groups.py get "<resource-group-name>"
```

Add a tag while preserving existing tags:

```powershell
python .\resource_groups.py tag "<resource-group-name>" --tag Environment=Test
```

Delete a Resource Group and all resources it contains:

```powershell
python .\resource_groups.py delete "<resource-group-name>" --yes
```

The `--subscription-id` global option can be used instead of the environment
variable. Add `--verbose` before the command for informational logs.

## References

- [Manage Azure resource groups by using Python](https://learn.microsoft.com/azure/azure-resource-manager/management/manage-resource-groups-python)
- [Azure Identity client library for Python](https://learn.microsoft.com/python/api/overview/azure/identity-readme)
