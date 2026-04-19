# Go Challenge: Concurrent Observer System

## Goal

Implement the **Observer (Publish–Subscribe) pattern** in Go using concurrency

## Scenario

You are building a system where clients can **subscribe to events**, and when an event occurs, all subscribers must be **notified concurrently**.

## Requirements

### Core Concepts

Your implementation must use:

* **goroutines**
* **channels**
* **`select` statement**
* **`sync.WaitGroup`**

### HTTP Layer

Use an HTTP framework (e.g., `echo`) to expose endpoints:

### (Optional) Graceful Shutdown

* On `SIGINT`:

    * Stop accepting new requests
    * Ensure all pending notifications are delivered

### curl POST examples:
```shell
curl -X POST http://localhost:1323/publish \
  -H "Content-Type: application/json" \
  -d '{
    "event_type": "deployment",
    "service": "payments-api",
    "status": "success"
  }'
```

```shell
curl -X POST http://localhost:1323/publish \
  -H "Content-Type: application/json" \
  -d '{
    "event_type": "deployment",
    "service": "auth-service",
    "status": "failure"
  }'
```

```shell
curl -X POST http://localhost:1323/publish \
  -H "Content-Type: application/json" \
  -d '{
    "event_type": "deployment",
    "service": "catalog-service",
    "status": "success"
  }'
```

