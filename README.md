
# Redis-Go: A Lightweight In-Memory Database

Redis-Go is a lightweight, in-memory, non-persistent database built in Go, inspired by the widely used Redis. It is designed for caching and supports multiple concurrent connections using Go's goroutines, making it fast and efficient. The application is also Dockerized for easy deployment and usage.

---

## Features

- **In-Memory Database:** Stores key-value pairs in memory for quick access.
- **Concurrent Connections:** Handles multiple simultaneous connections using Go's multi-threading capabilities.
- **Basic Commands Supported:**
  - `PING`: Responds with `PONG`.
  - `ECHO <message>`: Responds with the provided `<message>`.
  - `SET <key> <value>`: Stores a value for a given key.
  - `SET <key> <value> PX <milliseconds>`: Stores a value for a given key with an expiry time in milliseconds.
  - `GET <key>`: Retrieves the value of a key.
- **Data Integrity:** Implements mutexes to ensure mutual exclusivity and prevent data corruption.
- **Dockerized:** Easily deployable as a Docker container.

---

## Getting Started

### Prerequisites

- Docker installed on your system.
- (Optional) Redis CLI or netcat to interact with the application.

---

### Building and Running the Application

1. **Build the Docker Image:**

   ```bash
   docker build -t redis-go .
   ```

2. **Run the Docker Container:**

   ```bash
   docker run -p 6379:6379 redis-go
   ```

---

### Interacting with the Application

The application listens on port `6379`. You can interact with it using:

- **Redis CLI:** Connect using the Redis CLI.
- **Netcat (Windows):** Establish a TCP connection using netcat.

#### Supported Commands

- **PING:**  
  Sends `PING`, and the server responds with `PONG`.

  ```bash
  > PING
  < PONG
  ```

- **ECHO `<message>`:**  
  Sends a message to the server, and it echoes back the same message.

  ```bash
  > ECHO Hello
  < Hello
  ```

- **SET `<key>` `<value>`:**  
  Sets a key-value pair in memory.

  ```bash
  > SET foo bar
  < OK
  ```

- **SET `<key>` `<value>` PX `<milliseconds>`:**  
  Sets a key-value pair with an expiration time in milliseconds.

  ```bash
  > SET foo bar PX 100
  < OK
  ```

- **GET `<key>`:**  
  Retrieves the value of a given key.

  ```bash
  > GET foo
  < bar
  ```

---

## Concurrency and Data Safety

Redis-Go is built with Go's goroutines for handling multiple concurrent connections, ensuring high performance and responsiveness. Mutexes are used internally to maintain data integrity, preventing race conditions and ensuring thread-safe operations.



---

## Acknowledgements

Redis-Go is inspired by [Redis](https://redis.io/) and aims to provide a lightweight alternative for basic caching and in-memory operations.
