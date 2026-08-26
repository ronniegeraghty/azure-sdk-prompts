# Multilingual-translation reference applications

These applications implement the four `multilingual-translation.prompt.md` files.
Restoring or building them doesn't call Azure. Running them requires
`AZURE_TEXT_TRANSLATION_ENDPOINT`, Azure credentials, and one text argument.

Each application sends one request with French and Japanese targets, leaves source
language detection to the service, and prints returned language codes and text.
