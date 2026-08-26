# Transliteration reference applications

These applications implement the four `transliteration.prompt.md` files. Restoring
or building them doesn't call Azure. Running them requires
`AZURE_TEXT_TRANSLATION_ENDPOINT`, Azure credentials, and language, source-script,
target-script, and text arguments.

Each application invokes transliteration and prints the service-returned text and
script metadata.
