# 🚀 GoKotlinDistributed

GoKotlinDistributed is a distributed computing project demonstrating seamless interoperation between Go and Kotlin.

## Overview

GoKotlinDistributed facilitates efficient communication and coordination between Go and Kotlin components, enabling the distribution and processing of computational tasks across multiple worker nodes. Leveraging the strengths of both languages, this project demonstrates the integration and collaboration between different technologies in a distributed computing environment.

## Features

- **Distributed Task Processing:** Efficiently distribute computational tasks across multiple worker nodes for parallel processing.
- **Interoperation between Go and Kotlin:** Seamlessly integrate and communicate between Go and Kotlin components for enhanced functionality.
- **Worker Status Monitoring:** Monitor the status and activity of worker nodes in real-time for effective task management.
- **Task Prioritization:** Prioritize tasks based on predefined criteria to optimize resource utilization and task execution.
- **Worker Statistics Tracking:** Track and analyze worker performance metrics to identify bottlenecks and optimize system performance.
- **Log Rotation and Management:** Implement log rotation and management to ensure efficient log storage and maintenance.

## Installation

### Prerequisites

Ensure that the following prerequisites are met before installing GoKotlinDistributed:

- Go (version >= 1.15)
- Kotlin (version >= 1.4)
- Git

### Clone the Repository

Clone the GoKotlinDistributed repository to your local machine using the following command:

```bash
git clone https://github.com/Rishit-katiyar/GoKotlinDistributed.git
```

### Build and Run

#### Go Server

1. Navigate to the `GoServer` directory:

```bash
cd GoServer
```

2. Build the Go server:

```bash
go build -o server main.go coordinator.go worker.go http_handlers.go
```

3. Run the Go server:

```bash
./server
```

#### Kotlin Client

1. Navigate to the `KotlinClient` directory:

```bash
cd ../KotlinClient
```

2. Build the Kotlin client:

```bash
./gradlew build
```

3. Run the Kotlin client:

```bash
./gradlew run
```

## Troubleshooting

Encountering issues during installation or execution? Try the following troubleshooting steps:

### Go Server

- If the Go server fails to build, ensure that Go is installed and the required dependencies are available.
- If the Go server fails to run, check for any errors reported in the console output and ensure that the necessary ports (e.g., 8080) are not already in use.

### Kotlin Client

- If the Kotlin client fails to build, ensure that Kotlin is installed and the required dependencies are available.
- If the Kotlin client fails to run, check for any errors reported in the console output and ensure that the server is running and accessible.

## Usage

Utilize the GoKotlinDistributed project for efficient distributed task processing:

- Send tasks to the server using HTTP POST requests to `http://localhost:8080/task`.
- Retrieve task results from the server using HTTP GET requests to `http://localhost:8080/results`.
- Stop the workers using an HTTP GET request to `http://localhost:8080/stop`.
- View worker statistics using an HTTP GET request to `http://localhost:8080/stats`.

## License

This project is licensed under the [GNU General Public License v3.0](LICENSE).
