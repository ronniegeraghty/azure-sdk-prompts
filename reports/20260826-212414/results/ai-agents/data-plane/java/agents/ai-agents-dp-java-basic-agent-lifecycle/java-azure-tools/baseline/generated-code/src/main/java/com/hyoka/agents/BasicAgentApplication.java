package com.hyoka.agents;

import com.azure.ai.agents.persistent.MessagesClient;
import com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;
import com.azure.ai.agents.persistent.PersistentAgentsClient;
import com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;
import com.azure.ai.agents.persistent.RunsClient;
import com.azure.ai.agents.persistent.ThreadsClient;
import com.azure.ai.agents.persistent.models.CreateAgentOptions;
import com.azure.ai.agents.persistent.models.CreateRunOptions;
import com.azure.ai.agents.persistent.models.ListSortOrder;
import com.azure.ai.agents.persistent.models.MessageContent;
import com.azure.ai.agents.persistent.models.MessageRole;
import com.azure.ai.agents.persistent.models.MessageTextContent;
import com.azure.ai.agents.persistent.models.PersistentAgent;
import com.azure.ai.agents.persistent.models.PersistentAgentThread;
import com.azure.ai.agents.persistent.models.RunStatus;
import com.azure.ai.agents.persistent.models.ThreadMessage;
import com.azure.ai.agents.persistent.models.ThreadRun;
import com.azure.identity.DefaultAzureCredentialBuilder;

public final class BasicAgentApplication {
    private static final String AGENT_NAME = "hyoka-basic-agent";
    private static final String AGENT_INSTRUCTIONS = "Answer the user's question clearly and concisely.";
    private static final String USER_MESSAGE = "What is the capital of France?";
    private static final long POLL_INTERVAL_MILLIS = 500L;

    private BasicAgentApplication() {
    }

    public static void main(String[] args) throws InterruptedException {
        String projectEndpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
        String modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

        PersistentAgentsClient agentsClient = new PersistentAgentsClientBuilder()
            .endpoint(projectEndpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        PersistentAgentsAdministrationClient administrationClient =
            agentsClient.getPersistentAgentsAdministrationClient();
        ThreadsClient threadsClient = agentsClient.getThreadsClient();
        MessagesClient messagesClient = agentsClient.getMessagesClient();
        RunsClient runsClient = agentsClient.getRunsClient();

        PersistentAgent agent = administrationClient.createAgent(
            new CreateAgentOptions(modelDeploymentName)
                .setName(AGENT_NAME)
                .setInstructions(AGENT_INSTRUCTIONS));

        PersistentAgentThread thread = null;
        try {
            thread = threadsClient.createThread();
            messagesClient.createMessage(thread.getId(), MessageRole.USER, USER_MESSAGE);

            ThreadRun run = runsClient.createRun(new CreateRunOptions(thread.getId(), agent.getId()));
            run = waitForTerminalStatus(runsClient, thread.getId(), run);

            if (!RunStatus.COMPLETED.equals(run.getStatus())) {
                throw new IllegalStateException(describeUnsuccessfulRun(run));
            }

            printAssistantMessages(messagesClient, thread.getId());
        } finally {
            try {
                if (thread != null) {
                    threadsClient.deleteThread(thread.getId());
                }
            } finally {
                administrationClient.deleteAgent(agent.getId());
            }
        }
    }

    private static ThreadRun waitForTerminalStatus(
        RunsClient runsClient,
        String threadId,
        ThreadRun run
    ) throws InterruptedException {
        ThreadRun currentRun = run;
        while (!isTerminal(currentRun.getStatus())) {
            Thread.sleep(POLL_INTERVAL_MILLIS);
            currentRun = runsClient.getRun(threadId, currentRun.getId());
        }
        return currentRun;
    }

    private static boolean isTerminal(RunStatus status) {
        return RunStatus.COMPLETED.equals(status)
            || RunStatus.FAILED.equals(status)
            || RunStatus.CANCELLED.equals(status)
            || RunStatus.EXPIRED.equals(status);
    }

    private static void printAssistantMessages(MessagesClient messagesClient, String threadId) {
        for (ThreadMessage message : messagesClient.listMessages(
            threadId,
            null,
            null,
            ListSortOrder.ASCENDING,
            null,
            null
        )) {
            if (!MessageRole.AGENT.equals(message.getRole())) {
                continue;
            }

            for (MessageContent content : message.getContent()) {
                if (content instanceof MessageTextContent textContent) {
                    System.out.println(textContent.getText().getValue());
                }
            }
        }
    }

    private static String describeUnsuccessfulRun(ThreadRun run) {
        String message = "Agent run ended with status " + run.getStatus();
        if (run.getLastError() != null) {
            message += ": " + run.getLastError().getMessage();
        }
        return message;
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Required environment variable is not set: " + name);
        }
        return value;
    }
}
