using Azure.AI.Translation.Text;
using Azure.Identity;

if (args.Length != 1)
{
    Console.Error.WriteLine("Usage: multilingual-translation <text>");
    return 2;
}

string endpoint = Environment.GetEnvironmentVariable("AZURE_TEXT_TRANSLATION_ENDPOINT")
    ?? throw new InvalidOperationException(
        "AZURE_TEXT_TRANSLATION_ENDPOINT is required.");

TextTranslationClient client = new(
    new DefaultAzureCredential(),
    new Uri(endpoint));
TranslateInputItem input = new(
    args[0],
    new[]
    {
        new TranslationTarget("fr"),
        new TranslationTarget("ja"),
    });

TranslatedTextItem result = (
    await client.TranslateAsync(input)).Value;

Console.WriteLine(
    $"Detected source language: {result.DetectedLanguage?.Language}");
foreach (TranslationText translation in result.Translations)
{
    Console.WriteLine(
        $"Translation ({translation.Language}): {translation.Text}");
}

return 0;
