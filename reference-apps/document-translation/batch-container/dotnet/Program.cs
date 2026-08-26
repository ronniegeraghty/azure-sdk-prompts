using Azure.AI.Translation.Document;
using Azure.Identity;

if (args.Length != 3)
{
    Console.Error.WriteLine(
        "Usage: batch-container "
        + "<source-container-uri> <target-container-uri> <target-language>");
    return 2;
}

string endpoint = Environment.GetEnvironmentVariable(
    "AZURE_DOCUMENT_TRANSLATION_ENDPOINT")
    ?? throw new InvalidOperationException(
        "AZURE_DOCUMENT_TRANSLATION_ENDPOINT is required.");
DocumentTranslationClient client = new(
    new Uri(endpoint),
    new DefaultAzureCredential());
DocumentTranslationInput input = new(
    new Uri(args[0]),
    new Uri(args[1]),
    args[2]);

DocumentTranslationOperation operation =
    await client.StartTranslationAsync(input);
await operation.WaitForCompletionAsync();
Console.WriteLine($"Overall status: {operation.Status}");

await foreach (
    DocumentStatusResult document in operation.GetDocumentStatusesAsync())
{
    Console.WriteLine(
        $"Document: id={document.Id} status={document.Status}");
    if (document.Error is not null)
    {
        Console.WriteLine(
            $"Failure: code={document.Error.Code} "
            + $"message={document.Error.Message}");
    }
}

return 0;
