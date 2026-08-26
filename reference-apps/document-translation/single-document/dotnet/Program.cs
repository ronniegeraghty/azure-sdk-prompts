using Azure.AI.Translation.Document;
using Azure.Identity;

if (args.Length != 3)
{
    Console.Error.WriteLine(
        "Usage: single-document <input-path> <target-language> <output-path>");
    return 2;
}

string endpoint = Environment.GetEnvironmentVariable(
    "AZURE_DOCUMENT_TRANSLATION_ENDPOINT")
    ?? throw new InvalidOperationException(
        "AZURE_DOCUMENT_TRANSLATION_ENDPOINT is required.");
string mediaType = GetMediaType(args[0]);
SingleDocumentTranslationClient client = new(
    new Uri(endpoint),
    new DefaultAzureCredential());

await using FileStream input = File.OpenRead(args[0]);
MultipartFormFileData document = new(
    Path.GetFileName(args[0]),
    input,
    mediaType);
DocumentTranslateContent content = new(document);
BinaryData translated = (
    await client.TranslateAsync(args[1], content)).Value;

await using FileStream output = File.Create(args[2]);
await translated.ToStream().CopyToAsync(output);
return 0;

static string GetMediaType(string path) =>
    Path.GetExtension(path).ToLowerInvariant() switch
    {
        ".txt" => "text/plain",
        ".html" or ".htm" => "text/html",
        ".docx" => "application/vnd.openxmlformats-officedocument"
            + ".wordprocessingml.document",
        ".pdf" => "application/pdf",
        _ => throw new ArgumentException(
            $"Unsupported document extension: {Path.GetExtension(path)}"),
    };
