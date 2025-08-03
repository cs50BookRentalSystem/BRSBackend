# CS50 Book Rental System Backend

#### Video Demo: [Link to Video Demo]

#### Description:
A Go-based backend for a book rental system, providing APIs for managing books, students, and rentals.

## Introduction

This repository contains the Go backend for the CS50 Book Rental System (BRS). It provides a complete and robust API for managing book inventories, student rentals, and librarian administration. Designed for efficiency and maintainability, this backend serves as the server-side foundation for the BRS full-stack application.

---

## Project Overview

### Tech Stack

The BRS Backend is built upon a modern technology stack, chosen for performance and developer productivity.

*   **Backend:**
    *   **Go (Golang):** The high-performance language used for the core application logic.
    *   **Chi:** A lightweight and composable router for building structured HTTP services.
    *   **GORM:** A developer-friendly ORM library to simplify database interactions.
    *   **SQLite:** A self-contained, serverless SQL database engine, ideal for local development.
*   **API and Documentation:**
    *   **OpenAPI (Swagger):** The standard for defining and documenting the project's RESTful API.
    *   **oapi-codegen:** A tool for generating Go code directly from the OpenAPI specification.
*   **Frontend (Separate Repository):**
    *   **JavaScript (React.js):** The library used for the separate frontend application.

### Team

This project was brought to life through the collaborative efforts of a dedicated two-person team:

*   **Danny**: Frontend Development ([https://github.com/dannykhant](https://github.com/dannykhant))
*   **Dther**: Backend Development ([https://github.com/dtherhtun](https://github.com/dtherhtun))

---

## Core Features

The BRS Backend is equipped with a wide range of features to support a fully functional book rental system.

*   **Librarian Authentication:** Secure and reliable authentication for librarians, with session management to protect administrative endpoints.
*   **Book Management:** Comprehensive CRUD (Create, Read, Update, Delete) functionality for managing the book inventory. Librarians can add new titles, update book details, and adjust stock levels.
*   **Student Management:** A complete set of tools for managing student records, including the ability to add new students, view their rental history, and manage their accounts.
*   **Rental and Return Processing:** A streamlined workflow for processing book rentals and returns. The system tracks the status of each rental, from the moment a book is checked out to when it is returned.
*   **Overdue Rental Tracking:** An automated system for identifying and reporting overdue rentals, with a configurable rental period to suit the library's policies.
*   **Comprehensive Reporting:** Detailed reports on rental activities, including the most popular books, the number of active rentals, and a list of overdue items.
*   **Interactive API Documentation:** A user-friendly Swagger UI for exploring and interacting with the API, providing clear documentation for all endpoints, request payloads, and response formats.

---

## System Architecture

The BRS Backend is designed with a clean and layered architecture, promoting separation of concerns and making the codebase easy to navigate and maintain.

*   **`cmd/`:** The entry point of the application, responsible for initializing the command-line interface and starting the server.
*   **`openapi/`:** Contains the OpenAPI specification (`api.yaml`) and the generated Go code for the API, ensuring a contract-first approach to development.
*   **`pkg/`:** The heart of the application, containing all the core business logic and functionality.
    *   **`api/`:** The generated API handlers and Swagger UI assets.
    *   **`config/`:** Handles the loading and parsing of application configuration from YAML files.
    *   **`dto/`:** Data Transfer Objects (DTOs) that define the structure of data exchanged between the client and the server.
    *   **`handlers/`:** The HTTP request handlers that bridge the gap between the API and the underlying business logic.
    *   **`middleware/`:** A collection of HTTP middleware for handling cross-cutting concerns such as authentication, CORS, and request logging.
    *   **`models/`:** The database models that represent the core entities of the system, such as books, students, and rentals.
    *   **`repository/`:** The data access layer, responsible for all interactions with the database.
    *   **`services/`:** The business logic layer, where the core application services and use cases are implemented.
    *   **`validation/`:** Provides utilities for validating incoming data and ensuring data integrity.

---

## Getting Started

Follow these instructions to get a local instance of the BRS Backend up and running for development and testing.

### Prerequisites

*   **Go:** Version 1.18 or higher.
*   **Git:** For cloning the repository.
*   **Docker (Optional):** For running the application in a containerized environment.

### Configuration

The application is configured using a YAML file. For local development, you can create a `config.dev.yaml` file in the root of the project.

#### Example `config.dev.yaml`

```yaml
server:
  port: 8080
  env: dev
database:
  dsn: "./brs.sqlite"
librarian:
  user: "admin"
  pass: "securePasswd"
rent:
  rental_days: 7 # Number of days a book can be rented before it's considered overdue
```

#### Configuration Details

*   `server.port`: The port on which the application will run.
*   `server.env`: The application environment (e.g., `dev`, `prod`). When set to `prod`, the automatic seeding of sample data at startup is disabled.
*   `database.dsn`: The data source name for the database connection.
*   `librarian.user`: The username for the default librarian account.
*   `librarian.pass`: The password for the default librarian account.
*   `rent.rental_days`: The maximum number of days a book can be rented before it is considered overdue.

### Installation and Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/cs50BookRentalSystem/BRSBackend.git
    cd BRSBackend
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

### Database Initialization and Seeding

The application uses an SQLite database, and the database file (e.g., `brs.sqlite`) is automatically created during the initial startup if it doesn't already exist. The project offers two methods for populating the database with sample data:

**1. Automatic Seeding (Default)**

By default, the application will automatically seed the database with a predefined set of books and students when it starts. This feature is enabled as long as the `server.env` in your configuration file is not set to `prod`. This is the recommended approach for a quick and easy setup for local development.

**2. Manual Seeding (Optional)**

For more control over the seeding process, you can use the provided shell script to manually populate the database. This script makes a series of API calls to the running application to add a comprehensive list of books and students.

**To use the manual seeding script, you must first have the application running.** Once the server is active, you can execute the following command in a separate terminal:

```bash
./import_data.sh
```

This script will first authenticate with the API and then proceed to add the sample data. This method is ideal for testing the API endpoints or for situations where you want to re-seed the database without restarting the application.

### Running the Applicationpro

*   **With default configuration:**
    ```bash
    go run main.go
    ```

*   **With a custom configuration file:**
    ```bash
    go run main.go --config config.prod.yaml
    ```

The server will be accessible at `http://localhost:8080` (or the port specified in your configuration).

### Running with Docker

1.  **Build the Docker image:**
    ```bash
    docker build -t brs-backend .
    ```

2.  **Run the Docker container:**
    ```bash
    docker run -p 8080:8080 --name brs-backend-container brs-backend
    ```

You can also pull the latest image from the GitHub Container Registry:
```bash
docker run -p 8080:8080 --name brs-backend-container ghcr.io/cs50bookrentalsystem/brsbackend:latest
```

---

## API Documentation

The BRS Backend provides a fully documented API using the OpenAPI specification. You can access the interactive Swagger UI in your browser to explore the available endpoints, view request and response schemas, and test the API in real-time.

*   **Swagger UI:** `http://localhost:8080/swagger/index.html`

---

## Testing

To ensure the quality and reliability of the codebase, the project includes a suite of unit and integration tests. You can run these tests using the following command:

```bash
go test -v ./...
```

This will execute all test files in the project and provide a summary of the results.