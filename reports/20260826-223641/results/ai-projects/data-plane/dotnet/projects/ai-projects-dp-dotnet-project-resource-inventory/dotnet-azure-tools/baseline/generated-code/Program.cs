using Azure.AI.Projects;
using Azure.Identity;

string endpoint = GetRequiredEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
string connectionName = GetRequiredEnvironmentVariable("CONNECTION_NAME");
string deploymentName = GetRequiredEnvironmentVariable("DEPLOYMENT_NAME");

var projectClient = new AIProjectClient(
    new Uri(endpoint, UriKind.Absolute),
    new DefaultAzureCredential());

Console.WriteLine("Project connections");
Console.WriteLine("-------------------");

await foreach (AIProjectConnection connection in projectClient.Connections.GetConnectionsAsync())
{
    PrintConnection(connection);
}

Console.WriteLine();
Console.WriteLine($"Connection '{connectionName}'");
Console.WriteLine(new string('-', 13 + connectionName.Length));

AIProjectConnection selectedConnection =
    await projectClient.Connections.GetConnectionAsync(
        connectionName,
        includeCredentials: false);
PrintConnection(selectedConnection);

Console.WriteLine();
Console.WriteLine("Model deployments");
Console.WriteLine("-----------------");

await foreach (AIProjectDeployment deployment in projectClient.Deployments.GetDeploymentsAsync())
{
    if (deployment is ModelDeployment modelDeployment)
    {
        PrintModelDeployment(modelDeployment);
    }
}

Console.WriteLine();
Console.WriteLine($"Model deployment '{deploymentName}'");
Console.WriteLine(new string('-', 19 + deploymentName.Length));

AIProjectDeployment selectedDeployment =
    await projectClient.Deployments.GetDeploymentAsync(deploymentName);

if (selectedDeployment is not ModelDeployment selectedModelDeployment)
{
    throw new InvalidOperationException(
        $"Deployment '{deploymentName}' is not a model deployment.");
}

PrintModelDeployment(selectedModelDeployment);

static string GetRequiredEnvironmentVariable(string name)
{
    string? value = Environment.GetEnvironmentVariable(name);

    if (string.IsNullOrWhiteSpace(value))
    {
        throw new InvalidOperationException(
            $"Required environment variable '{name}' is not set.");
    }

    return value;
}

static void PrintConnection(AIProjectConnection connection)
{
    Console.WriteLine(
        $"Name: {connection.Name}; Type: {connection.Type}; " +
        $"Target: {connection.Target}; Default: {connection.IsDefault}");
}

static void PrintModelDeployment(ModelDeployment deployment)
{
    Console.WriteLine(
        $"Name: {deployment.Name}; Publisher: {deployment.ModelPublisher}; " +
        $"Model: {deployment.ModelName}; Version: {deployment.ModelVersion}");
}
