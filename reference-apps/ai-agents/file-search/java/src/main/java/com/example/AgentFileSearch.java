package com.example;

import com.azure.ai.agents.persistent.*;
import com.azure.ai.agents.persistent.models.*;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;
import java.util.List;

public final class AgentFileSearch {
    private static final String GUIDE
        = "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long "
        + "and hikers should bring a rain jacket.";
    private static final String QUESTION
        = "According to the uploaded guide, how long is the Cascade Loop "
        + "and what should hikers bring?";

    private AgentFileSearch() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
        String model = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");
        PersistentAgentsClient client = new PersistentAgentsClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
        PersistentAgentsAdministrationClient administration
            = client.getPersistentAgentsAdministrationClient();
        ThreadsClient threads = client.getThreadsClient();
        MessagesClient messages = client.getMessagesClient();
        RunsClient runs = client.getRunsClient();
        FilesClient files = client.getFilesClient();
        VectorStoresClient vectorStores = client.getVectorStoresClient();

        FileInfo uploadedFile = null;
        VectorStore vectorStore = null;
        PersistentAgent agent = null;
        PersistentAgentThread thread = null;
        try {
            uploadedFile = files.uploadFile(new UploadFileRequest(
                new FileDetails(BinaryData.fromString(GUIDE)).setFilename("trail-guide.txt"),
                FilePurpose.AGENTS));
            vectorStore = vectorStores.createVectorStore(
                List.of(uploadedFile.getId()),
                "hyoka-trail-guide",
                null,
                null,
                null,
                null);

            while (VectorStoreStatus.IN_PROGRESS.equals(vectorStore.getStatus())) {
                Thread.sleep(500);
                vectorStore = vectorStores.getVectorStore(vectorStore.getId());
            }
            if (!VectorStoreStatus.COMPLETED.equals(vectorStore.getStatus())) {
                throw new IllegalStateException(
                    "Vector store ended with status " + vectorStore.getStatus());
            }

            FileSearchToolResource searchResource = new FileSearchToolResource()
                .setVectorStoreIds(List.of(vectorStore.getId()));
            agent = administration.createAgent(new CreateAgentOptions(model)
                .setName("hyoka-trail-guide-agent")
                .setInstructions("Use file search to answer questions about the uploaded guide.")
                .setTools(List.of(new FileSearchToolDefinition()))
                .setToolResources(new ToolResources().setFileSearch(searchResource)));

            thread = threads.createThread();
            messages.createMessage(thread.getId(), MessageRole.USER, QUESTION);
            ThreadRun run = runs.createRun(new CreateRunOptions(thread.getId(), agent.getId()));
            do {
                Thread.sleep(500);
                run = runs.getRun(thread.getId(), run.getId());
            } while (RunStatus.QUEUED.equals(run.getStatus())
                || RunStatus.IN_PROGRESS.equals(run.getStatus()));

            if (!RunStatus.COMPLETED.equals(run.getStatus())) {
                throw new IllegalStateException("Run ended with status " + run.getStatus());
            }

            for (ThreadMessage message : messages.listMessages(
                thread.getId(), null, null, ListSortOrder.ASCENDING, null, null)) {
                if (!MessageRole.AGENT.equals(message.getRole())) {
                    continue;
                }
                for (MessageContent content : message.getContent()) {
                    if (content instanceof MessageTextContent text) {
                        System.out.println(text.getText().getValue());
                    }
                }
            }
        } finally {
            if (thread != null) {
                threads.deleteThread(thread.getId());
            }
            if (agent != null) {
                administration.deleteAgent(agent.getId());
            }
            if (vectorStore != null) {
                vectorStores.deleteVectorStore(vectorStore.getId());
            }
            if (uploadedFile != null) {
                files.deleteFile(uploadedFile.getId());
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
