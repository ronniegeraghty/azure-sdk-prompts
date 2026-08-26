using Azure.AI.Translation.Text;
using Azure.Identity;

if (args.Length != 4)
{
    Console.Error.WriteLine(
        "Usage: transliteration <language> <from-script> <to-script> <text>");
    return 2;
}

string endpoint = Environment.GetEnvironmentVariable("AZURE_TEXT_TRANSLATION_ENDPOINT")
    ?? throw new InvalidOperationException(
        "AZURE_TEXT_TRANSLATION_ENDPOINT is required.");
TextTranslationClient client = new(
    new DefaultAzureCredential(),
    new Uri(endpoint));

IReadOnlyList<TransliteratedText> results = (
    await client.TransliterateAsync(args[0], args[1], args[2], args[3])).Value;
TransliteratedText result = results.FirstOrDefault()
    ?? throw new InvalidOperationException(
        "The service returned no transliteration.");

Console.WriteLine($"Transliterated text: {result.Text}");
Console.WriteLine($"Returned script: {result.Script}");
return 0;
