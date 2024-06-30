# HelloFresh

HelloFresh is a global meal-kit company that delivers fresh ingredients and easy-to-follow recipes to customers' doors.
Our mission is to make cooking fun, easy, and convenient by providing all the necessary ingredients in just the right
quantities, along with detailed recipe instructions.

## Overview

This repository contains the source code for the HelloFresh backend services. The application is composed of several
microservices that handle different aspects of the system, including user management, product management, and order
processing.

## Getting Started

Follow the instructions below to set up and run the application.

### Prerequisites

- [Docker](https://www.docker.com/get-started) installed on your machine.
- [Go](https://golang.org/doc/install) programming language installed.

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/iam-benyamin/hellofresh.git
   cd hellofresh
   ```
2. Build and start the Docker containers:
   ```docker compose up --build -d```

3. Run the Go services:

```bash
go run ./cmd/user/main.go
go run ./cmd/product/main.go
go run ./cmd/order/main.go
```

### API Documentation

To interact with the endpoints, you can import the Postman documentation provided in the repository:

### Project Structure

The project is organized into several directories:

`adapter/`: Contains third party adapter-related code. like rabbitmq 
`cmd/`: Contains the main applications for each microservice.
`user/`: Handles user management (registration, authentication, etc.).
`config/`: Contains code that responsible for read environment variables form file or os env
`contract/`: Contains protobuf contracts.
`delivery/`: Delivery logic like grpc for product and user and http for order.
`docs/`: Contains documentation files, including Postman collections.
`entity/`: Defines entities and data structures.
`logger/`: Logging utilities and configurations.
`logs/`: Directory for log files.
`param/`: Parameter definitions.
`pkg/`: Contains utility packages.
`repository/`: Data access layer and repository implementations.
`serve/`: Server configurations and initialization.
`service/`: Business logic and service layer.
`validator/`: Validation logic and utilities.
`proto_compile_code/`: list of command for compile protobuf code.
`.golangci.yaml`: Configuration file for GolangCI-Lint.
`config.yaml`: Application configuration file.


have fine ;)