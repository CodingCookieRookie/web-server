# WebSocket Server

This project implements a WebSocket server designed to accept a websocket connection and respond with a random and unique big integer number.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go 1.20 or later

### Installation

First, clone the repository to your local machine:

```bash
git clone https://yourrepository.com/websocket-server.git
cd websocket-server
go run main.go
```

## Benchmark Feature (Optional, make sure websocket server is running)
This project supports a benchmark feature to test the server's throughput and number of concurrent connections it can handle. The number of concurrent connections and messages sent can be varied to test the server's concurrent load handling capability and measure throughput.

```bash
cd benchmark
go run benchmark.go
```

## Configuration (Optional)
This project supports optional configuration through a `.env` file to customise certain aspects of its operation. Using a `.env` file is optional but recommended for customising your development environment without altering the application's default settings.

### Creating a .env File
1. Create a file named `.env` in the root directory of the project.
2. Add configuration variables to the `.env` file as needed. Below are the available variables you can configure:

```plaintext
HOST = "localhost"
PORT = "8080"
LOG_FILE = "info.log"
```

#### Server configuration
HOST, PORT

#### Log configuration (for saving logs in log file)
LOG_FILE

