package com.example;

import com.azure.ai.agents.persistent.MessagesClient;
import com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;
import com.azure.ai.agents.persistent.PersistentAgentsClient;
import com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;
import com.azure.ai.agents.persistent.RunsClient;
import com.azure.ai.agents.persistent.ThreadsClient;
import com.azure.ai.agents.persistent.models.CreateAgentOptions;
import com.azure.ai.agents.persistent.models.CreateRunOptions;
import com.azure.ai.agents.persistent.models.MessageContent;
import com.azure.ai.agents.persistent.models.MessageRole;
import com.azure.ai.agents.persistent.models.MessageTextContent;
import com.azure.ai.agents.persistent.models.PersistentAgent;
import com.azure.ai.agents.persistent.models.PersistentAgentThread;
import com.azure.ai.agents.persistent.models.RunStatus;
import com.azure.ai.agents.persistent.models.ThreadMessage;
import com.azure.ai.agents.persistent.models.ThreadRun;
import com.azure.identity.DefaultAzureCredentialBuilder;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.List;

public final class BasicAgentConversation {
    private static final String AGENT_NAME = "hyoka-basic-agent";
    private static final String AGENT_INSTRUCTIONS =
        "Answer the user's question clearly and concisely.";
    private static final String USER_MESSAGE = "What is the capital of France?";
    private static final long POLL_INTERVAL_MILLIS = 500;

    private BasicAgentConversation() {
    }

    public static void main(String[] args) throws InterruptedException {
        String projectEndpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
        String modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

        PersistentAgentsClient client = new PersistentAgentsClientBuilder()
            .endpoint(projectEndpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
        PersistentAgentsAdministrationClient agentsClient =
            client.getPersistentAgentsAdministrationClient();
        ThreadsClient threadsClient = client.getThreadsClient();
        MessagesClient messagesClient = client.getMessagesClient();
        RunsClient runsClient = client.getRunsClient();

        PersistentAgent agent = null;
        PersistentAgentThread thread = null;
        try {
            CreateAgentOptions agentOptions = new CreateAgentOptions(modelDeploymentName)
                .setName(AGENT_NAME)
                .setInstructions(AGENT_INSTRUCTIONS);
            agent = agentsClient.createAgent(agentOptions);
            thread = threadsClient.createThread();
            messagesClient.createMessage(thread.getId(), MessageRole.USER, USER_MESSAGE);

            ThreadRun run = runsClient.createRun(
                new CreateRunOptions(thread.getId(), agent.getId()));
            while (!isTerminal(run.getStatus())) {
                Thread.sleep(POLL_INTERVAL_MILLIS);
                run = runsClient.getRun(thread.getId(), run.getId());
            }

            if (run.getStatus() != RunStatus.COMPLETED) {
                throw new IllegalStateException(
                    "Agent run ended with status " + run.getStatus()
                        + formatRunError(run));
            }

            printAssistantMessagesChronologically(messagesClient, thread.getId());
        } finally {
            try {
                if (thread != null) {
                    threadsClient.deleteThread(thread.getId());
                }
            } finally {
                if (agent != null) {
                    agentsClient.deleteAgent(agent.getId());
                }
            }
        }
    }

    private static boolean isTerminal(RunStatus status) {
        return status != RunStatus.QUEUED
            && status != RunStatus.IN_PROGRESS
            && status != RunStatus.CANCELLING;
    }

    private static String formatRunError(ThreadRun run) {
        if (run.getLastError() == null) {
            return ".";
        }
        return ": " + run.getLastError().getMessage();
    }

    private static void printAssistantMessagesChronologically(
        MessagesClient messagesClient, String threadId) {

        List<ThreadMessage> messages = new ArrayList<>();
        messagesClient.listMessages(threadId).forEach(messages::add);
        messages.sort(Comparator.comparing(ThreadMessage::getCreatedAt));

        for (ThreadMessage message : messages) {
            if (message.getRole() != MessageRole.AGENT) {
                continue;
            }
            for (MessageContent content : message.getContent()) {
                if (content instanceof MessageTextContent textContent) {
                    System.out.println(textContent.getText().getValue());
                }
            }
        }
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(
                "Required environment variable " + name + " is not set.");
        }
        return value;
    }
}
