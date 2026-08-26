# Basic agent lifecycle reference applications

These applications implement the scenario in the four
`basic-agent-lifecycle.prompt.md` files. They provide reviewable examples and
compile-time regression checks for the Azure AI Agents SDKs.

Restoring or building the projects doesn't call Azure. Running an application
requires `PROJECT_ENDPOINT`, `MODEL_DEPLOYMENT_NAME`, Azure credentials, and a
compatible Foundry project.

| Language | SDK package |
| --- | --- |
| .NET | `Azure.AI.Agents.Persistent` |
| Java | `com.azure:azure-ai-agents-persistent` |
| JavaScript/TypeScript | `@azure/ai-agents` |
| Python | `azure-ai-agents` |

