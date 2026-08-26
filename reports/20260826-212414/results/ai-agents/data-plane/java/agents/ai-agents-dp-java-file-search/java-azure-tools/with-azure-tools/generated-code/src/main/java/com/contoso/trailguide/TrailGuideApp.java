package com.contoso.trailguide;

import com.azure.ai.agents.persistent.FilesClient;
import com.azure.ai.agents.persistent.MessagesClient;
import com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;
import com.azure.ai.agents.persistent.PersistentAgentsClient;
import com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;
import com.azure.ai.agents.persistent.RunsClient;
import com.azure.ai.agents.persistent.ThreadsClient;
import com.azure.ai.agents.persistent.VectorStoresClient;
import com.azure.ai.agents.persistent.models.CreateAgentOptions;
import com.azure.ai.agents.persistent.models.CreateRunOptions;
import com.azure.ai.agents.persistent.models.FileDetails;
import com.azure.ai.agents.persistent.models.FileInfo;
import com.azure.ai.agents.persistent.models.FilePurpose;
import com.azure.ai.agents.persistent.models.FileSearchToolDefinition;
import com.azure.ai.agents.persistent.models.FileSearchToolResource;
import com.azure.ai.agents.persistent.models.ListSortOrder;
import com.azure.ai.agents.persistent.models.MessageContent;
import com.azure.ai.agents.persistent.models.MessageRole;
import com.azure.ai.agents.persistent.models.MessageTextContent;
import com.azure.ai.agents.persistent.models.PersistentAgent;
import com.azure.ai.agents.persistent.models.PersistentAgentThread;
import com.azure.ai.agents.persistent.models.RunStatus;
import com.azure.ai.agents.persistent.models.ThreadMessage;
import com.azure.ai.agents.persistent.models.ThreadRun;
import com.azure.ai.agents.persistent.models.ToolResources;
import com.azure.ai.agents.persistent.models.UploadFileRequest;
import com.azure.ai.agents.persistent.models.VectorStore;
import com.azure.ai.agents.persistent.models.VectorStoreFileCount;
import com.azure.ai.agents.persistent.models.VectorStoreStatus;
import com.azure.core.http.rest.PagedIterable;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;

import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Path;
import java.time.Duration;
import java.time.Instant;
import java.util.List;

public final class TrailGuideApp {
    private static final String DOCUMENT_TEXT =
        "The Contoso Trail Guide says the Cascade Loop is 42 kilometers long and hikers should bring a rain jacket.";
    private static final String QUESTION =
        "According to the uploaded guide, how long is the Cascade Loop and what should hikers bring?";
    private static final Duration POLL_INTERVAL = Duration.ofMillis(500);
    private static final Duration OPERATION_TIMEOUT = Duration.ofMinutes(5);

    private TrailGuideApp() {
    }

    public static void main(String[] args) throws Throwable {
        String endpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
        String modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

        PersistentAgentsClient client = new PersistentAgentsClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        PersistentAgentsAdministrationClient administrationClient =
            client.getPersistentAgentsAdministrationClient();
        ThreadsClient threadsClient = client.getThreadsClient();
        MessagesClient messagesClient = client.getMessagesClient();
        RunsClient runsClient = client.getRunsClient();
        FilesClient filesClient = client.getFilesClient();
        VectorStoresClient vectorStoresClient = client.getVectorStoresClient();

        String fileId = null;
        String vectorStoreId = null;
        String agentId = null;
        String threadId = null;
        Throwable primaryFailure = null;

        try {
            Path documentPath = Path.of("contoso-trail-guide.txt");
            Files.writeString(documentPath, DOCUMENT_TEXT, StandardCharsets.UTF_8);

            FileInfo uploadedFile = filesClient.uploadFile(new UploadFileRequest(
                new FileDetails(BinaryData.fromFile(documentPath)).setFilename(documentPath.getFileName().toString()),
                FilePurpose.AGENTS));
            fileId = uploadedFile.getId();

            VectorStore vectorStore = vectorStoresClient.createVectorStore(
                List.of(fileId), "hyoka-trail-guide-store", null, null, null, null);
            vectorStoreId = vectorStore.getId();
            vectorStore = waitForVectorStore(vectorStoresClient, vectorStore);
            requireSuccessfulIndexing(vectorStore);

            FileSearchToolResource fileSearchResource =
                new FileSearchToolResource().setVectorStoreIds(List.of(vectorStoreId));
            CreateAgentOptions agentOptions = new CreateAgentOptions(modelDeploymentName)
                .setName("hyoka-trail-guide-agent")
                .setInstructions("Answer questions using the uploaded trail guide. Use file search before answering.")
                .setTools(List.of(new FileSearchToolDefinition()))
                .setToolResources(new ToolResources().setFileSearch(fileSearchResource));
            PersistentAgent agent = administrationClient.createAgent(agentOptions);
            agentId = agent.getId();

            PersistentAgentThread thread = threadsClient.createThread();
            threadId = thread.getId();
            messagesClient.createMessage(threadId, MessageRole.USER, QUESTION);

            ThreadRun run = runsClient.createRun(new CreateRunOptions(threadId, agentId));
            run = waitForRun(runsClient, threadId, run);
            if (run.getStatus() != RunStatus.COMPLETED) {
                String detail = run.getLastError() == null ? "" : ": " + run.getLastError().getMessage();
                throw new IllegalStateException("Agent run ended with status " + run.getStatus() + detail);
            }

            printAgentTextChronologically(messagesClient, threadId);
        } catch (Throwable failure) {
            primaryFailure = failure;
            throw failure;
        } finally {
            RuntimeException cleanupFailure = cleanUp(
                threadsClient,
                administrationClient,
                vectorStoresClient,
                filesClient,
                threadId,
                agentId,
                vectorStoreId,
                fileId);
            if (cleanupFailure != null) {
                if (primaryFailure != null) {
                    primaryFailure.addSuppressed(cleanupFailure);
                } else {
                    throw cleanupFailure;
                }
            }
        }
    }

    private static VectorStore waitForVectorStore(VectorStoresClient client, VectorStore vectorStore)
        throws InterruptedException {
        Instant deadline = Instant.now().plus(OPERATION_TIMEOUT);
        while (vectorStore.getStatus() == VectorStoreStatus.IN_PROGRESS) {
            if (Instant.now().isAfter(deadline)) {
                throw new IllegalStateException("Timed out waiting for vector store indexing");
            }
            Thread.sleep(POLL_INTERVAL.toMillis());
            vectorStore = client.getVectorStore(vectorStore.getId());
        }
        return vectorStore;
    }

    private static void requireSuccessfulIndexing(VectorStore vectorStore) {
        VectorStoreFileCount counts = vectorStore.getFileCounts();
        boolean allFilesIndexed = counts != null
            && counts.getTotal() > 0
            && counts.getCompleted() == counts.getTotal()
            && counts.getFailed() == 0
            && counts.getCancelled() == 0;
        if (vectorStore.getStatus() != VectorStoreStatus.COMPLETED || !allFilesIndexed) {
            String countsDescription = counts == null
                ? "unavailable"
                : String.format(
                    "total=%d, completed=%d, failed=%d, cancelled=%d",
                    counts.getTotal(),
                    counts.getCompleted(),
                    counts.getFailed(),
                    counts.getCancelled());
            throw new IllegalStateException(
                "Vector store indexing was not successful. Status=" + vectorStore.getStatus()
                    + ", file counts: " + countsDescription);
        }
    }

    private static ThreadRun waitForRun(RunsClient client, String threadId, ThreadRun run)
        throws InterruptedException {
        Instant deadline = Instant.now().plus(OPERATION_TIMEOUT);
        while (run.getStatus() == RunStatus.QUEUED
            || run.getStatus() == RunStatus.IN_PROGRESS
            || run.getStatus() == RunStatus.CANCELLING) {
            if (Instant.now().isAfter(deadline)) {
                throw new IllegalStateException("Timed out waiting for the agent run to reach a terminal status");
            }
            Thread.sleep(POLL_INTERVAL.toMillis());
            run = client.getRun(threadId, run.getId());
        }
        if (run.getStatus() == RunStatus.REQUIRES_ACTION) {
            throw new IllegalStateException("Agent run requires unsupported external tool output");
        }
        return run;
    }

    private static void printAgentTextChronologically(MessagesClient client, String threadId) {
        PagedIterable<ThreadMessage> messages =
            client.listMessages(threadId, null, null, ListSortOrder.ASCENDING, null, null);
        for (ThreadMessage message : messages) {
            if (message.getRole() != MessageRole.ASSISTANT) {
                continue;
            }
            for (MessageContent content : message.getContent()) {
                if (content instanceof MessageTextContent textContent) {
                    System.out.println(textContent.getText().getValue());
                }
            }
        }
    }

    private static RuntimeException cleanUp(
        ThreadsClient threadsClient,
        PersistentAgentsAdministrationClient administrationClient,
        VectorStoresClient vectorStoresClient,
        FilesClient filesClient,
        String threadId,
        String agentId,
        String vectorStoreId,
        String fileId) {

        RuntimeException failure = null;
        failure = delete(failure, "thread", threadId, threadsClient::deleteThread);
        failure = delete(failure, "agent", agentId, administrationClient::deleteAgent);
        failure = delete(failure, "vector store", vectorStoreId, vectorStoresClient::deleteVectorStore);
        failure = delete(failure, "uploaded file", fileId, filesClient::deleteFile);
        return failure;
    }

    private static RuntimeException delete(
        RuntimeException accumulatedFailure,
        String resourceType,
        String resourceId,
        ResourceDeleter deleter) {

        if (resourceId == null) {
            return accumulatedFailure;
        }
        try {
            deleter.delete(resourceId);
        } catch (RuntimeException exception) {
            RuntimeException wrapped =
                new RuntimeException("Failed to delete " + resourceType + " " + resourceId, exception);
            if (accumulatedFailure == null) {
                return wrapped;
            }
            accumulatedFailure.addSuppressed(wrapped);
        }
        return accumulatedFailure;
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }

    @FunctionalInterface
    private interface ResourceDeleter {
        void delete(String resourceId);
    }
}
