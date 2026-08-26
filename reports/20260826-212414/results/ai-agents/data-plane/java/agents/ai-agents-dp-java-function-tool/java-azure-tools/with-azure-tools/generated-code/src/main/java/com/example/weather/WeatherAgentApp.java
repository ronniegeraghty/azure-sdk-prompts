package com.example.weather;

import com.azure.ai.agents.persistent.MessagesClient;
import com.azure.ai.agents.persistent.PersistentAgentsAdministrationClient;
import com.azure.ai.agents.persistent.PersistentAgentsClient;
import com.azure.ai.agents.persistent.PersistentAgentsClientBuilder;
import com.azure.ai.agents.persistent.RunsClient;
import com.azure.ai.agents.persistent.ThreadsClient;
import com.azure.ai.agents.persistent.models.CreateAgentOptions;
import com.azure.ai.agents.persistent.models.FunctionDefinition;
import com.azure.ai.agents.persistent.models.FunctionToolDefinition;
import com.azure.ai.agents.persistent.models.ListSortOrder;
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
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;
import java.util.Map;

public final class WeatherAgentApp {
    private static final String FUNCTION_NAME = "get_weather";
    private static final String USER_MESSAGE = "What is the weather in Seattle in celsius?";
    private static final long POLL_INTERVAL_MILLIS = 500;

    private WeatherAgentApp() {
    }

    public static void main(String[] args) throws InterruptedException {
        String endpoint = requireEnvironmentVariable("PROJECT_ENDPOINT");
        String modelDeploymentName = requireEnvironmentVariable("MODEL_DEPLOYMENT_NAME");

        PersistentAgentsClient client = new PersistentAgentsClientBuilder()
            .endpoint(endpoint)
            .credential(new DefaultAzureCredentialBuilder().build())
            .buildClient();

        PersistentAgentsAdministrationClient administrationClient
            = client.getPersistentAgentsAdministrationClient();
        ThreadsClient threadsClient = client.getThreadsClient();
        MessagesClient messagesClient = client.getMessagesClient();
        RunsClient runsClient = client.getRunsClient();

        PersistentAgent agent = null;
        PersistentAgentThread thread = null;
        try {
            agent = administrationClient.createAgent(
                new CreateAgentOptions(modelDeploymentName)
                    .setName("hyoka-weather-agent")
                    .setInstructions(
                        "Answer weather questions by calling get_weather. "
                            + "You must use get_weather for every weather question.")
                    .setTools(List.of(createWeatherTool())));

            thread = threadsClient.createThread();
            messagesClient.createMessage(thread.getId(), MessageRole.USER, USER_MESSAGE);

            ThreadRun run = runsClient.createRun(thread.getId(), agent.getId());
            run = pollRun(runsClient, thread.getId(), run);
            if (run.getStatus() != RunStatus.COMPLETED) {
                throw new IllegalStateException(describeRunFailure(run));
            }

            printAssistantMessages(messagesClient, thread.getId());
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

    private static FunctionToolDefinition createWeatherTool() {
        Map<String, Object> parameters = Map.of(
            "type", "object",
            "properties", Map.of(
                "location", Map.of(
                    "type", "string",
                    "description", "The city whose weather is requested."),
                "unit", Map.of(
                    "type", "string",
                    "description", "The temperature unit.",
                    "enum", List.of("c", "f"))),
            "required", List.of("location", "unit"),
            "additionalProperties", false);

        FunctionDefinition definition = new FunctionDefinition(
            FUNCTION_NAME,
            BinaryData.fromObject(parameters))
            .setDescription("Get deterministic weather for a location.");
        return new FunctionToolDefinition(definition);
    }

    private static ThreadRun pollRun(RunsClient runsClient, String threadId, ThreadRun run)
        throws InterruptedException {
        while (!isTerminal(run.getStatus())) {
            if (run.getStatus() == RunStatus.REQUIRES_ACTION) {
                run = submitRequiredToolOutputs(runsClient, threadId, run);
            } else if (run.getStatus() == RunStatus.QUEUED
                || run.getStatus() == RunStatus.IN_PROGRESS
                || run.getStatus() == RunStatus.CANCELLING) {
                Thread.sleep(POLL_INTERVAL_MILLIS);
                run = runsClient.getRun(threadId, run.getId());
            } else {
                throw new IllegalStateException("Unexpected run status: " + run.getStatus());
            }
        }
        return run;
    }

    private static ThreadRun submitRequiredToolOutputs(
        RunsClient runsClient, String threadId, ThreadRun run) {
        if (!(run.getRequiredAction() instanceof SubmitToolOutputsAction action)) {
            throw new IllegalStateException("Run requires an unsupported action.");
        }

        List<ToolOutput> outputs = new ArrayList<>();
        for (RequiredToolCall toolCall : action.getSubmitToolOutputs().getToolCalls()) {
            if (!(toolCall instanceof RequiredFunctionToolCall functionCall)) {
                throw new IllegalStateException("Run requested a non-function tool.");
            }
            if (!FUNCTION_NAME.equals(functionCall.getFunction().getName())) {
                throw new IllegalStateException(
                    "Run requested unknown function: " + functionCall.getFunction().getName());
            }

            WeatherArguments arguments = BinaryData
                .fromString(functionCall.getFunction().getArguments())
                .toObject(WeatherArguments.class);
            String result = getWeather(arguments);
            outputs.add(new ToolOutput()
                .setToolCallId(functionCall.getId())
                .setOutput(result));
        }

        if (outputs.isEmpty()) {
            throw new IllegalStateException("Run required action but supplied no function calls.");
        }
        return runsClient.submitToolOutputsToRun(threadId, run.getId(), outputs);
    }

    private static String getWeather(WeatherArguments arguments) {
        if (arguments == null || arguments.location == null || arguments.location.isBlank()) {
            throw new IllegalArgumentException("get_weather requires a non-empty location.");
        }
        if (arguments.unit == null) {
            throw new IllegalArgumentException("get_weather requires unit c or f.");
        }

        String location = arguments.location.trim();
        String unit = arguments.unit.toLowerCase(Locale.ROOT);
        if (!"Seattle".equalsIgnoreCase(location)) {
            throw new IllegalArgumentException("get_weather supports only Seattle.");
        }
        if (!"c".equals(unit) && !"f".equals(unit)) {
            throw new IllegalArgumentException("get_weather unit must be c or f.");
        }

        int temperature = "c".equals(unit) ? 21 : 70;
        return BinaryData.fromObject(new WeatherResult("Seattle", unit, temperature)).toString();
    }

    private static void printAssistantMessages(MessagesClient messagesClient, String threadId) {
        for (ThreadMessage message : messagesClient.listMessages(
            threadId, null, null, ListSortOrder.ASCENDING, null, null)) {
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

    private static boolean isTerminal(RunStatus status) {
        return status == RunStatus.COMPLETED
            || status == RunStatus.FAILED
            || status == RunStatus.CANCELLED
            || status == RunStatus.EXPIRED;
    }

    private static String describeRunFailure(ThreadRun run) {
        if (run.getLastError() == null) {
            return "Run ended with status " + run.getStatus() + ".";
        }
        return "Run ended with status " + run.getStatus() + ": "
            + run.getLastError().getCode() + " - " + run.getLastError().getMessage();
    }

    private static String requireEnvironmentVariable(String name) {
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

    public static final class WeatherResult {
        private final String location;
        private final String unit;
        private final int temperature;

        public WeatherResult(String location, String unit, int temperature) {
            this.location = location;
            this.unit = unit;
            this.temperature = temperature;
        }

        public String getLocation() {
            return location;
        }

        public String getUnit() {
            return unit;
        }

        public int getTemperature() {
            return temperature;
        }
    }
}
