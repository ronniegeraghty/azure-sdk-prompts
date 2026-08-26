package com.example;

import com.azure.ai.agents.persistent.*;
import com.azure.ai.agents.persistent.models.*;
import com.azure.core.util.BinaryData;
import com.azure.identity.DefaultAzureCredentialBuilder;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.json.JsonMapper;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

public final class AgentFunctionTool {
    private static final String USER_MESSAGE = "What is the weather in Seattle in celsius?";

    private AgentFunctionTool() {
    }

    public static void main(String[] args) throws Exception {
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

        FunctionDefinition definition = new FunctionDefinition(
            "get_weather",
            BinaryData.fromObject(Map.of(
                "type", "object",
                "properties", Map.of(
                    "location", Map.of("type", "string"),
                    "unit", Map.of("type", "string", "enum", List.of("c", "f"))),
                "required", List.of("location", "unit"))))
            .setDescription("Gets the current weather for a location.");
        FunctionToolDefinition tool = new FunctionToolDefinition(definition);

        PersistentAgent agent = null;
        PersistentAgentThread thread = null;
        try {
            agent = administration.createAgent(new CreateAgentOptions(model)
                .setName("hyoka-weather-agent")
                .setInstructions("Use get_weather for every weather question.")
                .setTools(List.of(tool)));
            thread = threads.createThread();
            messages.createMessage(thread.getId(), MessageRole.USER, USER_MESSAGE);
            ThreadRun run = runs.createRun(new CreateRunOptions(thread.getId(), agent.getId()));

            do {
                Thread.sleep(500);
                run = runs.getRun(thread.getId(), run.getId());
                if (RunStatus.REQUIRES_ACTION.equals(run.getStatus())
                    && run.getRequiredAction() instanceof SubmitToolOutputsAction action) {
                    List<ToolOutput> outputs = new ArrayList<>();
                    for (RequiredToolCall call : action.getSubmitToolOutputs().getToolCalls()) {
                        if (!(call instanceof RequiredFunctionToolCall functionCall)
                            || !"get_weather".equals(functionCall.getFunction().getName())) {
                            throw new IllegalStateException("Unexpected tool call.");
                        }
                        JsonNode arguments = new JsonMapper()
                            .readTree(functionCall.getFunction().getArguments());
                        String output = getWeather(
                            arguments.get("location").asText(),
                            arguments.get("unit").asText());
                        outputs.add(new ToolOutput()
                            .setToolCallId(functionCall.getId())
                            .setOutput(output));
                    }
                    run = runs.submitToolOutputsToRun(thread.getId(), run.getId(), outputs);
                }
            } while (RunStatus.QUEUED.equals(run.getStatus())
                || RunStatus.IN_PROGRESS.equals(run.getStatus())
                || RunStatus.REQUIRES_ACTION.equals(run.getStatus()));

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
        }
    }

    private static String getWeather(String location, String unit) throws Exception {
        if (!location.toLowerCase().contains("seattle")
            || !List.of("c", "f").contains(unit)) {
            throw new IllegalArgumentException("Unsupported weather request.");
        }
        return new JsonMapper().writeValueAsString(Map.of(
            "location", "Seattle",
            "temperature", unit.equals("c") ? 21 : 70,
            "unit", unit));
    }

    private static String requireEnvironmentVariable(String name) {
        String value = System.getenv(name);
        if (value == null || value.isBlank()) {
            throw new IllegalStateException(name + " is required.");
        }
        return value;
    }
}
