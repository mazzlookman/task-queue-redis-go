# Task Queue System with Go and Redis

A **Task Queue System** implemented in Go using **Redis** for managing task queues. This project demonstrates how to use a task queue to manage background jobs effectively with Go's concurrency and Redis as a lightweight message broker.

## Features

- **Task Management**:
  - Add tasks to a Redis-backed queue.
  - Monitor task status (e.g., `pending`, `in-progress`, `done`).
- **Worker System**:
  - Background workers process tasks from the queue.
  - Tasks are processed asynchronously.
- **API**:
  - RESTful API for managing tasks using [Fiber](https://gofiber.io).
- **Scalability**:
  - Add more workers to scale task processing.

## Tech Stack

- **Programming Language**: Go
- **Message Broker**: Redis
- **Web Framework**: Fiber
- **Package Manager**: Go Modules

## Prerequisites

1. **Go**: Install [Go](https://golang.org/doc/install).
2. **Redis**: Install [Redis](https://redis.io/docs/latest/operate/oss_and_stack/install/install-redis/) server and ensure it's running locally or accessible remotely.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-username/task-queue-system.git
   cd task-queue-system
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run Redis (if not running already):
   ```bash
   redis-server
   ```

## Usage

### 1. Start the Application

Run the server:
```bash
go run main.go
```

The server will start at `http://localhost:3000`.

### 2. API Endpoints

| Method | Endpoint               | Description                     |
|--------|------------------------|---------------------------------|
| POST   | `/tasks`               | Add a new task to the queue.    |
| GET    | `/tasks/:id/status`    | Get the status of a specific task. |

#### Example Requests

1. **Add a Task**
   ```bash
   curl -X POST http://localhost:3000/tasks
   ```

   Response:
   ```json
   {
       "taskID": "123e4567-e89b-12d3-a456-426614174000"
   }
   ```

2. **Check Task Status**
   ```bash
   curl http://localhost:3000/tasks/123e4567-e89b-12d3-a456-426614174000/status
   ```

   Response:
   ```json
   {
       "status": "pending"
   }
   ```

### 3. Worker

The worker runs in the background to process tasks. By default, the worker starts alongside the application. It listens to the `task_queue` and processes tasks sequentially.

To add more workers for parallel processing, run additional instances of the worker.

---

## Peek Task Queue in Redis

Use redis-cli to monitor Redis data:
```bash
redis-cli -h localhost -p 6379
```

1. View all tasks in the queue:
   ```bash
   LRANGE task_queue 0 -1
   ```
2. View the status of a task:
   ```bash
   GET task:<taskID>:status
   ```

## How It Works?

1. **Adding a Task**: The task is added to a Redis-backed queue (`task_queue`) with a status of `pending`.
2. **Worker Processing**:
   - A worker continuously monitors the queue for new tasks.
   - When a task is picked, its status is updated to `in-progress`.
   - After processing, the status is updated to `done`.
3. **Monitoring**: The API provides endpoints to check task status by task ID.