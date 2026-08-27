# Cosmos DB Python three-arm analysis

## Result

| Arm | Prompt checks | Python checks | Workspace | Azure MCP check | Total |
|---|---:|---:|---:|---:|---:|
| Baseline | 5/6 | 3/5 | 0/1 | 0/1 | 8/13 |
| Azure MCP + general Azure skills | 5/6 | 4/5 | 0/1 | 0/1 | 9/13 |
| Azure MCP + general Azure skills + Python SDK skills | 5/6 | 5/5 | 0/1 | 0/1 | 10/13 |

The prompt-specific result was unchanged. All arms missed the
`enable_cross_partition_query` criterion. The gains came from generic Python
criteria:

- The Azure arm added correct client lifecycle management.
- The SDK arm also added `DefaultAzureCredential`, reaching 5/5 Python checks.

The workspace check reported failure even though each arm generated
`cosmos_crud.py` and `requirements.txt`; treat that as a grader/reporting issue,
not missing output.

## Skill and MCP evidence

| Arm | General Azure skills | Python SDK skill | Azure MCP |
|---|---|---|---|
| Baseline | Not configured | Not configured | Not configured |
| Azure | Loaded, not explicitly invoked | Not configured | Used: 5 successful calls |
| Azure + SDK | Loaded, not explicitly invoked | `azure-cosmos-py` loaded and invoked | Configured, no calls |

The Azure arm used `azure-get_azure_bestpractices` twice and
`azure-documentation` three times. Despite these calls, the tool grader recorded
0/1 because it looked for an `azure` tool call rather than tools served by the
Azure MCP server.

The SDK arm explicitly invoked `azure-cosmos-py`, then read its partitioning
references. This is direct evidence that the language-specific plugin was used.
There is no equivalent invocation evidence for the general Azure skills in
either enhanced arm.

## SDK v1 duration validation

The generated report contains nonzero elapsed durations for correlated tool
events, including the SDK skill invocation (3,560 ms) and all five Azure MCP
calls (4,437-10,234 ms). This confirms the SDK v1 tool-duration fix works in a
live evaluation report.

## Conclusion

The three-arm setup runs successfully and the language-specific skill provides
an incremental improvement. However, this single run does not demonstrate
general Azure skill invocation: arm 2 improved through Azure MCP usage, while
arm 3 improved through the Python SDK skill and did not call Azure MCP.
