# Azure Service Bus TypeScript demo

Install the required Service Bus SDK and development dependencies:

```powershell
npm install
```

The Azure messaging package used by the program is:

```powershell
npm install @azure/service-bus
```

Create the queue, topic, and subscription in an Azure Service Bus namespace
before running the example. Copy `.env.example` to `.env`, replace the
placeholders, load those variables into the shell, and run:

```powershell
npm start
```

The connection string must allow sending to the queue and topic and receiving
from the queue and topic subscription. Do not commit a real connection string.
