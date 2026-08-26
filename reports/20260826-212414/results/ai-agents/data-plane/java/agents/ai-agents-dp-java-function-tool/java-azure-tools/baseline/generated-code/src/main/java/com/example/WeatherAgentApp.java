package com.example;

import com.azure.ai.agents.persistent.MessagesClient;
import com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;
import com.azure.ai.agents.persistent.PersistentAgentsClient;
import com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;
import com.azure.ai.agents.persistent.RunsClient;
import com.azure.ai.agents.persistent.ThreadsClient;
import com.azure.ai.agents.persistent.models.CreateAgentOptions;
import com.azure.ai.agents.persistent.models.CreateRunOptions;
import com.azure.ai.agents.persistent.models.FunctionDefinition;
import com.azure.ai.agents.persistent.models.FunctionToolDefinition;
import com.azure.ai.agents.persistent.models.MessageContent;
import com.azure.ai.agents.persistent.models.MessageRole;
import com.azure.ai.agents.persistent.models.MessageTextContent;
import com.azure.ai.agents.persistent.models.PersistentAgent;
import com.azure.ai.agents.persistent.models.PersistentAgentThread;
import com.azure.ai.agents.persistent.models.RequiredFunctionToolCall;
import com.azure.ai.agents.persistent.models.RequiredToolCall;
import com.azure.ai.agents.persistent.models.RunStatus;
import com.azure.ai.agents.persistent.models.SubmitToolOutputsAction;
import com.azure.ai.agents.persistent.models.ThreadMessage;
import com.azure.ai.agents.persistent.models.ThreadRun;
import com.azure.ai.agents.persistent.models.ToolOutput;
import com.azure.core.http.rest.PagedIterable;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;

import java.util.ArrayList;
import java.util.Comparator;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;

public final class WeatherAgentApp {
    private static final String USER_MESSAGE = "What is the weather in Seattle in celsius?";
    private static final ObjectMapper JSON = new ObjectMapper();

    private WeatherAgentApp() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requiredEnvironmentVariable("PROJECT_ENDPOINT");
        String modelDeploymentName = requiredEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

        PersistentAgentsClient agentsClient = new PersistentAgentsClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();
        PersistentAgentsAdministrationClient administrationClient
            = agentsClient.getPersistentAgentsAdministrationClient();
        ThreadsClient threadsClient = agentsClient.getThreadsClient();
        MessagesClient messagesClient = agentsClient.getMessagesClient();
        RunsClient runsClient = agentsClient.getRunsClient();

        PersistentAgent agent = null;
        PersistentAgentThread thread = null;
        try {
            FunctionToolDefinition weatherTool = new FunctionToolDefinition(
                new FunctionDefinition("get_weather", BinaryData.fromObject(weatherParameters()))
                    .setDescription("Get the current weather for a location in celsius or fahrenheit."));

            agent = administrationClient.createAgent(
                new CreateAgentOptions(modelDeploymentName)
                    .setName("hyoka-weather-agent")
                    .setInstructions(
                        "Answer weather questions by calling the get_weather function. "
                            + "Do not answer a weather question without using that function.")
                    .setTools(List.of(weatherTool)));

            thread = threadsClient.createThread();
            messagesClient.createMessage(thread.getId(), MessageRole.USER, USER_MESSAGE);

            ThreadRun run = runsClient.createRun(new CreateRunOptions(thread.getId(), agent.getId()));
            run = pollRun(thread.getId(), run, runsClient);

            if (run.getStatus() != RunStatus.COMPLETED) {
                String detail = run.getLastError() == null ? "" : ": " + run.getLastError().getMessage();
                throw new IllegalStateException("Run ended with status " + run.getStatus() + detail);
            }

            printAssistantMessagesChronologically(messagesClient, thread.getId());
        } finally {
            try {
                if (thread != null) {
                    threadsClient.deleteThread(thread.getId());
                }
            } finally {
                if (agent != null) {
                    administrationClient.deleteAgent(agent.getId());
                }
            }
        }
    }

    private static ThreadRun pollRun(String threadId, ThreadRun run, RunsClient runsClient)
        throws InterruptedException {
        while (isNonTerminal(run.getStatus())) {
            if (run.getStatus() == RunStatus.REQUIRES_ACTION) {
                run = submitRequiredToolOutputs(threadId, run, runsClient);
            } else {
                Thread.sleep(500);
                run = runsClient.getRun(threadId, run.getId());
            }
        }
        return run;
    }

    private static ThreadRun submitRequiredToolOutputs(
        String threadId, ThreadRun run, RunsClient runsClient) {
        if (!(run.getRequiredAction() instanceof SubmitToolOutputsAction)) {
            throw new IllegalStateException("Run requested an unsupported action.");
        }

        SubmitToolOutputsAction action = (SubmitToolOutputsAction) run.getRequiredAction();
        List<ToolOutput> outputs = new ArrayList<>();
        for (RequiredToolCall toolCall : action.getSubmitToolOutputs().getToolCalls()) {
            if (!(toolCall instanceof RequiredFunctionToolCall)) {
                throw new IllegalStateException("Run requested an unsupported tool call.");
            }

            RequiredFunctionToolCall functionCall = (RequiredFunctionToolCall) toolCall;
            if (!"get_weather".equals(functionCall.getFunction().getName())) {
                throw new IllegalStateException(
                    "Run requested unknown function: " + functionCall.getFunction().getName());
            }

            WeatherArguments arguments = decodeArguments(functionCall.getFunction().getArguments());
            String output = getWeather(arguments.location, arguments.unit);
            outputs.add(new ToolOutput()
                .setToolCallId(functionCall.getId())
                .setOutput(output));
        }

        return runsClient.submitToolOutputsToRun(threadId, run.getId(), outputs);
    }

    private static WeatherArguments decodeArguments(String arguments) {
        try {
            WeatherArguments decoded = JSON.readValue(arguments, WeatherArguments.class);
            if (decoded.location == null || decoded.location.isBlank()) {
                throw new IllegalArgumentException("location must be a non-empty string");
            }
            if (!Set.of("c", "f").contains(decoded.unit)) {
                throw new IllegalArgumentException("unit must be either 'c' or 'f'");
            }
            return decoded;
        } catch (JsonProcessingException e) {
            throw new IllegalArgumentException("Invalid get_weather arguments: " + arguments, e);
        }
    }

    private static String getWeather(String location, String unit) {
        if (!"Seattle".equalsIgnoreCase(location.trim())) {
            throw new IllegalArgumentException("Unsupported location: " + location);
        }

        int temperature = "c".equals(unit) ? 21 : 70;
        Map<String, Object> result = new LinkedHashMap<>();
        result.put("location", location);
        result.put("temperature", temperature);
        result.put("unit", unit);
        try {
            return JSON.writeValueAsString(result);
        } catch (JsonProcessingException e) {
            throw new IllegalStateException("Could not encode weather result.", e);
        }
    }

    private static Map<String, Object> weatherParameters() {
        Map<String, Object> location = Map.of(
            "type", "string",
            "description", "The city whose weather is requested.");
        Map<String, Object> unit = Map.of(
            "type", "string",
            "enum", List.of("c", "f"),
            "description", "Temperature unit: c for celsius or f for fahrenheit.");

        Map<String, Object> schema = new LinkedHashMap<>();
        schema.put("type", "object");
        schema.put("properties", Map.of("location", location, "unit", unit));
        schema.put("required", List.of("location", "unit"));
        schema.put("additionalProperties", false);
        return schema;
    }

    private static void printAssistantMessagesChronologically(
        MessagesClient messagesClient, String threadId) {
        PagedIterable<ThreadMessage> listedMessages = messagesClient.listMessages(threadId);
        List<ThreadMessage> messages = new ArrayList<>();
        listedMessages.forEach(messages::add);
        messages.sort(Comparator.comparing(ThreadMessage::getCreatedAt));

        for (ThreadMessage message : messages) {
            if (MessageRole.AGENT.equals(message.getRole())) {
                for (MessageContent content : message.getContent()) {
                    if (content instanceof MessageTextContent) {
                        System.out.println(((MessageTextContent) content).getText().getValue());
                    }
                }
            }
        }
    }

    private static boolean isNonTerminal(RunStatus status) {
        return status == RunStatus.QUEUED
            || status == RunStatus.IN_PROGRESS
            || status == RunStatus.REQUIRES_ACTION
            || status == RunStatus.CANCELLING;
    }

    private static String requiredEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException("Environment variable " + name + " is required.");
        }
        return value;
    }

    public static final class WeatherArguments {
        public String location;
        public String unit;
    }
}
