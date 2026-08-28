# Azure Resource Group Management

Install the required packages:

```powershell
python -m pip install -r requirements.txt
```

Authenticate locally with a credential supported by `DefaultAzureCredential`, then
set the subscription ID:

```powershell
$env:AZURE_SUBSCRIPTION_ID = "<subscription-id>"
```

List all resource groups, retrieve the selected group, and add or replace a tag:

```powershell
python .\manage_resource_groups.py `
  --resource-group "<resource-group-name>" `
  --tag-name "environment" `
  --tag-value "development"
```

To also delete the group, explicitly enable and confirm deletion:

```powershell
python .\manage_resource_groups.py `
  --resource-group "<resource-group-name>" `
  --tag-name "environment" `
  --tag-value "development" `
  --delete `
  --confirm-delete "<resource-group-name>"
```

The identity must have permissions to read, update, and delete resource groups in
the selected subscription. The script does not create or deploy Azure resources.
