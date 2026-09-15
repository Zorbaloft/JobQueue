# JobQueue

Project under construction, with the goal of expanding knowledge in Go.

The aim is to build a job queue: a server receives tasks, enqueues them, and processes them. At this stage the queue itself does not exist yet — what already works is TCP communication between a producer and a server.

## Current state

- **TCP server** (`tcp-server.go`) — listens on port `8080`, accepts connections, and logs received messages.
- **Producer** (`producer.go`) — connects to the server and sends a message (currently `"Hello World"`).
- **Entry point** (`main.go`) — starts the server, sends the message, and shuts the server down cleanly.

The Go module is `JobQueue.com/jobqueue` and uses Go 1.27. 

## How to run

From the project root:

```bash
go run .
```

You should see the received message in the server log, something like:

```text
received from 127.0.0.1:...: Hello World
```

## Next steps

- Introduce an in-memory job queue
- Separate the server lifecycle from message sending (today the producer runs in the same process)
- Define a simple protocol to enqueue, process, and report the status of each job
