Sure! Here's a clear, professional English documentation for your proto file that any programmer can understand:

---

# Agent Service Protocol Buffer Definition Documentation

This `.proto` file defines the gRPC service and messages for managing **Agent** registration, command distribution, and command result reporting in a backend system.

---

## Package and Options

* **Package:** `agent`
* **Go Package:** `central-agent-controller/proto`
* **Syntax:** `proto3`

---

## Messages

### `AgentConfigRequest`

Sent by the agent to register or update its configuration on the server.

| Field           | Type              | Description                                                           |
| --------------- | ----------------- | --------------------------------------------------------------------- |
| `agent_id`      | `string`          | Unique identifier of the agent                                        |
| `hostname`      | `string`          | Agent's hostname                                                      |
| `ip_address`    | `string`          | Agent's current IP address                                            |
| `os_version`    | `string`          | Operating system version                                              |
| `agent_version` | `string`          | Version of the agent software                                         |
| `capabilities`  | `repeated string` | List of agent capabilities (e.g. scan, update, firewall, application) |
| `timestamp`     | `int64`           | Unix timestamp of this configuration update                           |
| `api_key`       | `string`          | API key for authenticating the agent                                  |

---

### `AgentCommandResponse`

Server response to an agent's config request, containing commands for the agent to execute.

| Field      | Type              | Description                                                                               |
| ---------- | ----------------- | ----------------------------------------------------------------------------------------- |
| `commands` | `repeated string` | List of commands to execute (e.g., "START\_SCAN", "WAIT", "UPDATE\_SOFTWARE", "FIND\_ME") |
| `reason`   | `string`          | Optional explanation or reason for commands or status                                     |

---

### `Status` (enum)

Defines possible statuses for command execution results:

| Value                        | Meaning                            |
| ---------------------------- | ---------------------------------- |
| `STATUS_UNSPECIFIED`         | Status not specified               |
| `STATUS_SUCCESS`             | Command executed successfully      |
| `STATUS_IN_PROGRESS`         | Command is still running           |
| `STATUS_FAILED_AGENT_ERROR`  | Command failed due to agent error  |
| `STATUS_FAILED_SERVER_ERROR` | Command failed due to server error |
| `STATUS_REJECTED`            | Command was rejected by the agent  |

---

### `AgentCommandResultRequest`

Message sent from the agent to the server reporting the result of a command execution.

| Field              | Type                 | Description                                        |
| ------------------ | -------------------- | -------------------------------------------------- |
| `agent_id`         | `string`             | Agent identifier                                   |
| `executed_command` | `string`             | Name of the command executed (e.g., "START\_SCAN") |
| `status`           | `Status` enum        | Execution status of the command                    |
| `result`           | `string`             | Result details or summary                          |
| `current_config`   | `AgentConfigRequest` | Agent's current configuration after execution      |
| `timestamp`        | `int64`              | Unix timestamp of the result report                |

---

### `AcknowledgementResponse`

Server response acknowledging receipt of a command result and optionally sending the next command.

| Field          | Type     | Description                            |
| -------------- | -------- | -------------------------------------- |
| `message`      | `string` | Acknowledgement or instruction message |
| `next_command` | `string` | Next command for the agent, if any     |

---

## Service: `AgentService`

Defines the gRPC service interface with two RPC methods:

| Method       | Request Type                | Response Type             | Description                                                                                    |
| ------------ | --------------------------- | ------------------------- | ---------------------------------------------------------------------------------------------- |
| `SendConfig` | `AgentConfigRequest`        | `AgentCommandResponse`    | Agent sends its configuration; server responds with commands to execute                        |
| `SendResult` | `AgentCommandResultRequest` | `AcknowledgementResponse` | Agent sends result of command execution; server acknowledges and optionally sends next command |

---

## Summary

This protocol enables a **bidirectional workflow** between a central server and distributed agents:

1. **Agent registration/heartbeat** through `SendConfig`, receiving commands to execute.
2. **Agent reports command results** via `SendResult`, getting acknowledgements and new commands.

The design supports flexible commands and tracks the agent’s capabilities and status.

---

If you want, I can help you write example usage snippets or detailed implementation guidance next.
