using Azure.AI.Projects;
using Azure.Identity;

string endpoint = RequireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
string connectionName = RequireEnvironmentVariable("CONNECTION_NAME");
string deploymentName = RequireEnvironmentVariable("DEPLOYMENT_NAME");
AIProjectClient client = new(new Uri(endpoint), new DefaultAzureCredential());

Console.WriteLine("Connections:");
await foreach (AIProjectConnection connection in client.Connections.GetConnectionsAsync())
{
    PrintConnection(connection);
}

Console.WriteLine("Selected connection:");
AIProjectConnection selectedConnection = await client.Connections.GetConnectionAsync(
    connectionName,
    includeCredentials: false);
PrintConnection(selectedConnection);

Console.WriteLine("Model deployments:");
await foreach (AIProjectDeployment deployment in client.Deployments.GetDeploymentsAsync())
{
    if (deployment is ModelDeployment modelDeployment)
    {
        PrintDeployment(modelDeployment);
    }
}

Console.WriteLine("Selected model deployment:");
AIProjectDeployment selectedDeployment =
    await client.Deployments.GetDeploymentAsync(deploymentName);
if (selectedDeployment is not ModelDeployment selectedModelDeployment)
{
    throw new InvalidOperationException(
        $"{deploymentName} is not a model deployment.");
}
PrintDeployment(selectedModelDeployment);

static void PrintConnection(AIProjectConnection connection) =>
    Console.WriteLine(
        $"{connection.Name} | {connection.Type} | {connection.Target} | default={connection.IsDefault}");

static void PrintDeployment(ModelDeployment deployment) =>
    Console.WriteLine(
        $"{deployment.Name} | {deployment.ModelPublisher} | {deployment.ModelName} | {deployment.ModelVersion}");

static string RequireEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException($"{name} is required.");
