using Azure.AI.Projects;
using Azure.Identity;
using Azure.Storage.Blobs;

string endpoint = RequireEnvironmentVariable("FOUNDRY_PROJECT_ENDPOINT");
string datasetName = RequireEnvironmentVariable("DATASET_NAME");
string datasetVersion = RequireEnvironmentVariable("DATASET_VERSION");
string inputPath = RequireEnvironmentVariable("DATA_FILE_PATH");
string downloadPath = RequireEnvironmentVariable("DOWNLOAD_FILE_PATH");
byte[] expectedBytes = await File.ReadAllBytesAsync(inputPath);

AIProjectClient client = new(new Uri(endpoint), new DefaultAzureCredential());
bool uploaded = false;

try
{
    await client.Datasets.UploadFileAsync(
        name: datasetName,
        version: datasetVersion,
        filePath: inputPath);
    uploaded = true;

    AIProjectDataset dataset =
        await client.Datasets.GetDatasetAsync(datasetName, datasetVersion);
    Console.WriteLine(
        $"Dataset: name={dataset.Name} version={dataset.Version} id={dataset.Id} "
        + $"type={dataset.GetType().Name} dataUri={dataset.DataUri}");

    DatasetCredential credential =
        await client.Datasets.GetCredentialsAsync(datasetName, datasetVersion);
    Uri blobUri = credential.BlobReference.BlobUri;
    UriBuilder authorizedBlobUri =
        new(credential.BlobReference.Credential.SasUri)
        {
            Path = blobUri.AbsolutePath,
        };
    BlobClient blobClient = new(authorizedBlobUri.Uri);

    string? parent = Path.GetDirectoryName(Path.GetFullPath(downloadPath));
    if (parent is not null)
    {
        Directory.CreateDirectory(parent);
    }
    await blobClient.DownloadToAsync(downloadPath);

    byte[] downloadedBytes = await File.ReadAllBytesAsync(downloadPath);
    if (!expectedBytes.AsSpan().SequenceEqual(downloadedBytes))
    {
        throw new InvalidOperationException("Downloaded bytes don't match the source file.");
    }
    Console.WriteLine($"Downloaded bytes verified: {downloadedBytes.Length}");
}
finally
{
    if (uploaded)
    {
        await client.Datasets.DeleteAsync(datasetName, datasetVersion);
    }
}

static string RequireEnvironmentVariable(string name) =>
    Environment.GetEnvironmentVariable(name)
    ?? throw new InvalidOperationException($"{name} is required.");
