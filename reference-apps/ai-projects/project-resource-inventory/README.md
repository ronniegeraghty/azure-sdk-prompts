# Project resource-inventory reference applications

These applications implement the four `project-resource-inventory.prompt.md` files.
Restoring or building them doesn't call Azure. Running them requires a Foundry
project with the selected connection and deployment.

Each application traverses both pageable inventories, performs both name-based
retrievals, and prints only typed resource metadata.
