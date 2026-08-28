# Azure Resource Group Manager

This sample uses the Azure management-plane SDK and `DefaultAzureCredential` to:

1. Create or update a resource group.
2. List every resource group in the subscription.
3. Retrieve the created resource group's details.
4. Merge one or more tags without discarding existing tags.
5. Optionally delete the resource group and wait for deletion to finish.

## Requirements

- Python 3.9 or newer
- An Azure identity with Resource Group Contributor permissions at the
  subscription scope
- `AZURE_SUBSCRIPTION_ID` set to the target subscription ID

Install the required packages:

```powershell
python -m pip install -r requirements.txt
```

`DefaultAzureCredential` supports local developer-tool credentials, service
principal environment variables, workload identity, and managed identity. No
credentials are stored in this project.

## Usage

Create, list, inspect, and tag a group while retaining it:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
python .\resource_group_manager.py --name "example-rg" --location "eastus" --tag "environment=dev"
```

Run the same workflow and delete the group at the end:

```powershell
python .\resource_group_manager.py --name "example-rg" --location "eastus" --tag "environment=dev" --delete
```

Deleting a resource group also deletes every resource it contains. The script
therefore requires the explicit `--delete` option.

## References

- [Manage Azure resource groups by using Python](https://learn.microsoft.com/azure/azure-resource-manager/management/manage-resource-groups-python)
- [ResourceGroupsOperations API reference](https://learn.microsoft.com/python/api/azure-mgmt-resource/azure.mgmt.resource.resources.operations.resourcegroupsoperations?view=azure-python)
- [`azure-identity` on PyPI](https://pypi.org/project/azure-identity/)
- [`azure-mgmt-resource` on PyPI](https://pypi.org/project/azure-mgmt-resource/)
