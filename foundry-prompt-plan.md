# Foundry prompt plan

This plan covers Azure AI client libraries that have published packages for .NET,
Java, JavaScript/TypeScript, and Python. Each case expands into four parity prompts
and four persisted reference applications.

Use the current SDK source, README, and samples to verify each language-specific
implementation. Dedicated skill coverage is useful evidence but is not an inclusion
requirement.

## Included libraries and cases

### Azure AI Agents

Package family: `Azure.AI.Agents.Persistent`, `azure-ai-agents-persistent`,
`@azure/ai-agents`, and `azure-ai-agents`.

1. **Basic persistent conversation** — Create an agent, thread, message, and run;
   poll to completion; read assistant text; and delete created resources.
   - Evaluation focus: correct identifiers across operations, run polling, typed
     message-content traversal, and cleanup.
   - Status: the first four-language prompt and reference set is complete.
2. **Function tool execution** — Define a local weather function, create an agent
   with the function tool, process required tool calls, submit outputs, and print the
   final response.
   - Evaluation focus: tool schema, required-action detection, argument decoding,
     tool-call correlation, and output submission.
3. **File search** — Upload a document, create a vector store, attach file search to
   an agent, ask a grounded question, and clean up files and vector stores.
   - Evaluation focus: file purpose, vector-store readiness, tool resources,
     grounded run workflow, and dependent-resource cleanup.

### Azure AI Projects

Package family: `Azure.AI.Projects`, `azure-ai-projects`,
`@azure/ai-projects`, and `azure-ai-projects`.

1. **Project resource inventory** — List project connections and deployments, select
   entries by name, and print their SDK-defined metadata.
   - Evaluation focus: Projects client operation groups, paging, name-based
     retrieval, and typed connection/deployment results.
2. **Dataset lifecycle** — Upload a local data file as a project dataset, retrieve
   its version and metadata, download it, and delete the created dataset version.
   - Evaluation focus: file upload, dataset and version identifiers, download
     handling, and resource lifecycle.
   - Status: the four-language prompt and reference set is complete.
3. **Evaluation run** — Create an evaluation from a small JSONL or CSV dataset,
   select a built-in evaluator, poll the evaluation, and print metric results.
   - Evaluation focus: evaluator configuration, dataset reference, long-running
     operation handling, and metric-result traversal.
   - Status: the four-language prompt and reference set is complete.

### Azure AI Voice Live

Package family: `Azure.AI.VoiceLive`, `azure-ai-voicelive`,
`@azure/ai-voicelive`, and `azure-ai-voicelive`.

1. **Audio-file conversation** — Configure a manual-turn Voice Live session, send a
   PCM16 input file in chunks, commit the input, request a response, and write
   returned audio to a file.
   - Evaluation focus: session configuration, audio format, chunk encoding, input
     commit, response events, and audio-delta assembly.
2. **Voice function calling** — Register a local function, start a voice session,
   handle a function-call event, return the correlated tool output, and continue the
   response.
   - Evaluation focus: realtime event loop, function schema, call identifiers,
     argument decoding, output submission, and response continuation.

### Azure AI Content Safety

Package family: `Azure.AI.ContentSafety`, `azure-ai-contentsafety`,
`@azure-rest/ai-content-safety`, and `azure-ai-contentsafety`.

Wave 2 status: skipped. The official JavaScript source tag
`@azure-rest/ai-content-safety_1.0.2` contains the required package and APIs, but
version 1.0.2 isn't restorable from the required Azure SDK private npm registry.
On 2026-08-26, registry metadata listed only 1.0.3 alpha builds (from
`1.0.3-alpha.20260302.1` through `1.0.3-alpha.20260825.1`);
`npm install` returned `ETARGET` for 1.0.2, and the registry's direct 1.0.2
tarball URL returned 404. The .NET 1.0.0, Java 1.0.20, and Python 1.0.0
references restored and compiled, and the JavaScript APIs were source-verified,
but the four-language contract requires every package to restore from the
mandated registry.

1. **Text harm analysis** — Analyze supplied text and print the severity returned
   for every harm category.
   - Evaluation focus: text request model, category result traversal, and severity
     mapping.
   - Status: skipped because the library-level JavaScript package restore failed.
2. **Image harm analysis** — Load an image, submit its bytes, and print category
   severities from the image result.
   - Evaluation focus: binary image encoding, image request model, and image-specific
     result traversal.
   - Status: skipped because the library-level JavaScript package restore failed.
3. **Blocklist lifecycle** — Create a blocklist, add terms, analyze text with that
   blocklist, report matched terms, and delete the blocklist.
   - Evaluation focus: blocklist and term identifiers, blocklist-enabled analysis,
     matched-term results, and cleanup.
   - Status: skipped because the library-level JavaScript package restore failed.

### Azure AI Content Understanding

Package family: `Azure.AI.ContentUnderstanding`,
`azure-ai-contentunderstanding`, `@azure/ai-content-understanding`, and
`azure-ai-contentunderstanding`.

1. **Invoice analysis** — Analyze an invoice with a prebuilt or selected analyzer,
   poll to completion, and print extracted fields with confidence values.
   - Evaluation focus: analyzer selection, document submission, poller result,
     field typing, and confidence extraction.
2. **Custom analyzer lifecycle** — Create an analyzer with a field schema, analyze a
   sample document, retrieve its result, and delete the analyzer.
   - Evaluation focus: analyzer schema, create/update operation, analyzer ID reuse,
     structured fields, and cleanup.
3. **Semantic document chunking** — Analyze a document with semantic chunking and
   print each chunk's content, span, and associated metadata.
   - Evaluation focus: chunking configuration and traversal of document chunk
     results rather than flattened text.

### Azure AI Document Intelligence

Package family: `Azure.AI.DocumentIntelligence`,
`azure-ai-documentintelligence`, `@azure-rest/ai-document-intelligence`, and
`azure-ai-documentintelligence`.

Wave 3 status: skipped. The official JavaScript source tag
`@azure-rest/ai-document-intelligence_1.1.0` identifies the released 1.1.0
package, but that version isn't restorable from the required Azure SDK private
npm registry. On 2026-08-27, registry metadata listed only 1.1.0 and 1.1.1
alpha builds; `npm view` returned 404 for 1.1.0, and the registry's `latest`
tag pointed to `1.1.1-alpha.20260826.1`. The four-language contract requires
every package to restore from the mandated registry.

1. **Prebuilt invoice extraction** — Analyze an invoice with `prebuilt-invoice` and
   print vendor, invoice ID, date, and total from typed document fields.
   - Evaluation focus: prebuilt model selection, poller use, document fields,
     typed values, and confidence.
   - Status: skipped because the library-level JavaScript package restore failed.
2. **Layout and Markdown extraction** — Analyze a document with the layout model and
   Markdown output, then print paragraphs, tables, and page information.
   - Evaluation focus: output-content format, layout result hierarchy, table cells,
     and page spans.
   - Status: skipped because the library-level JavaScript package restore failed.
3. **Custom classification** — Submit a document to a custom classifier, poll the
   operation, and print each classified document's type and confidence.
   - Evaluation focus: classifier ID, classification poller, per-document results,
     and confidence values.
   - Status: skipped because the library-level JavaScript package restore failed.

### Azure AI Vision Image Analysis

Package family: `Azure.AI.Vision.ImageAnalysis`,
`azure-ai-vision-imageanalysis`, `@azure-rest/ai-vision-image-analysis`, and
`azure-ai-vision-imageanalysis`.

Wave 3 status: skipped. Official JavaScript source has no 1.0.0 release tag;
its latest release is `1.0.0-beta.3` from 2024-07-18. On 2026-08-27, the
required Azure SDK private npm registry likewise listed only 1.0.0 alpha and
beta builds, and `npm view` returned 404 for 1.0.0. Although the registry's
`latest` tag points to `1.0.0-beta.3`, both `npm install` and an isolated
`npm pack` for that version returned `E401`, so the package isn't restorable
from the mandated registry. The four-language contract requires every package
to restore from that registry.

1. **Caption, tags, and objects** — Analyze an image for caption, tags, and detected
   objects and print text, confidence, and object rectangles.
   - Evaluation focus: feature selection and traversal of three distinct typed
     result collections.
   - Status: skipped because the library-level JavaScript package restore failed.
2. **OCR** — Analyze an image with the read feature and print text grouped by block
   and line with bounding polygons.
   - Evaluation focus: read feature selection, block/line hierarchy, and geometry.
   - Status: skipped because the library-level JavaScript package restore failed.

### Azure AI Vision Face

Package family: `Azure.AI.Vision.Face`, `azure-ai-vision-face`,
`@azure-rest/ai-vision-face`, and `azure-ai-vision-face`.

1. **Face detection** — Detect faces in an image and print each face ID, rectangle,
   and requested recognition attributes.
   - Evaluation focus: detection and recognition model selection, requested
     attributes, and typed face results.
2. **Stateless face verification** — Detect one face in each of two images, verify
   the two face IDs, and print the match decision and confidence.
   - Evaluation focus: carrying detected IDs into verification and reading the
     verification result.
3. **Person-group identification** — Create a person group, add people and face
   images, train it, identify faces from a query image, and remove created resources.
   - Evaluation focus: administration/client separation, training poller, detected
     face IDs, ranked candidates, and cleanup.

### Azure AI Text Translation

Package family: `Azure.AI.Translation.Text`, `azure-ai-translation-text`,
`@azure-rest/ai-translation-text`, and `azure-ai-translation-text`.

1. **Multilingual translation** — Translate supplied text into French and Japanese
   and print detected source language plus each target translation.
   - Evaluation focus: multiple target languages, detected-language result, and
     per-target translation traversal.
   - Status: the four-language prompt and reference set is complete.
2. **Transliteration** — Transliterate text between specified scripts and print the
   transliterated output with the script metadata.
   - Evaluation focus: language/from-script/to-script parameters and transliteration
     result types.
   - Status: the four-language prompt and reference set is complete.

### Azure AI Document Translation

Package family: `Azure.AI.Translation.Document`,
`azure-ai-translation-document`, `@azure-rest/ai-translation-document`, and
`azure-ai-translation-document`.

1. **Single-document translation** — Upload or read one document, translate it to a
   target language, and persist the returned translated document.
   - Evaluation focus: document content and media type, target language, binary
     response handling, and output-file integrity.
   - Status: the four-language prompt and reference set is complete.
2. **Batch container translation** — Start a source-to-target container translation,
   poll the batch operation, and enumerate each document's status and failure details.
   - Evaluation focus: source/target inputs, long-running operation, per-document
     paging, status, and failure result types.
   - Status: the four-language prompt and reference set is complete.

## Scope summary

- 10 included library families.
- 26 language-independent cases.
- 104 prompts and 104 reference applications after four-language expansion.
- Prompt-specific criteria cover only the SDK workflow. Shared criteria continue to
  cover authentication, secrets, general error handling, and language conventions.
- Live Azure execution is optional. Every reference application must restore and
  compile without contacting Azure.

## Implementation order

1. Finish Azure AI Agents.
2. Add Azure AI Projects and Content Safety to establish paging, resource lifecycle,
   text, and binary request patterns.
3. Add Document Intelligence, Image Analysis, and both Translation libraries.
4. Add Content Understanding and Face, which require more involved result models and
   resource administration.
5. Add Voice Live last because realtime event validation needs the most specialized
   test harness.

## Excluded or deferred libraries

- Azure AI OpenAI: no longer the supported Azure SDK direction.
- Azure AI Inference: a currently published .NET package wasn't confirmed.
- Anomaly Detector and Personalizer: retire on October 1, 2026.
- Metrics Advisor: retired.
- Form Recognizer: superseded by Document Intelligence for current service APIs.
- Text Analytics/Language Text, Language Conversations, Question Answering, Azure AI
  Discovery, Agent Server, Evaluation, ML, and Transcription: defer until supported
  package parity is confirmed across all four languages.
