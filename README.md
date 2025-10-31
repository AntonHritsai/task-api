# Go Task Server API

A RESTful API for a simple task management system, built with Go. This project is designed to demonstrate a clean, layered application architecture and best practices for building backend services.

The API provides full CRUD (Create, Read, Update, Delete) functionality for tasks, along with several endpoints for filtering and searching.

---

## Architecture

This project follows a classic 3-tier (or layered) architecture to ensure separation of concerns, testability, and maintainability:

-   **Handlers (Controller Layer):** Responsible for handling HTTP requests and responses. It parses incoming data, validates it, and calls the appropriate service. It knows nothing about the database.
-   **Services (Business Logic Layer):** Contains all the business logic of the application. It orchestrates operations and coordinates the repository layer to access data. It knows nothing about HTTP.
-   **Repositories (Data Access Layer):** Responsible for all communication with the database. It encapsulates all SQL/GORM queries and provides a clean API for the service layer to work with data.

---

## Technologies Used

-   **Go**: The core programming language.
-   **Gin**: A high-performance HTTP web framework.
-   **GORM**: A developer-friendly ORM library for Go.
-   **MySQL**: The relational database used for data storage.
-   **GoDotEnv**: For managing environment variables.

---

## API Endpoints

### Tasks

| Method | Endpoint              | Description                                        |
| :----- | :-------------------- | :------------------------------------------------- |
| `POST` | `/tasks`              | Creates a new task.                                |
| `GET`  | `/tasks`              | Retrieves a list of all tasks.                     |
| `GET`  | `/tasks/:id`          | Retrieves a single task by its ID.                 |
| `PUT`  | `/tasks/:id`          | Updates an existing task by its ID.                |
| `DELETE`| `/tasks/:id`         | Deletes a task by its ID.                          |

### Filter & Search

| Method | Endpoint              | Description                                        |
| :----- | :-------------------- | :------------------------------------------------- |
| `GET`  | `/tasks/today`        | Retrieves all tasks with a deadline of today.      |
| `GET`  | `/tasks/overdue`      | Retrieves all unfinished tasks that are past their deadline. |
| `GET`  | `/tasks/search`       | Searches for tasks by title. (e.g., `/tasks/search?title=refactor`) |

---

## Getting Started

### Prerequisites

-   Go (version 1.21 or newer recommended)
-   A running MySQL instance

### Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/AntonHritsai/task-api.git
    cd task-api
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the root of the project and add your database connection details:
    ```env
    DB_USER="your_db_user"
    DB_PASSWORD="your_db_password"
    DB_HOST="127.0.0.1"
    DB_PORT="3306"
    DB_NAME="your_db_name"
    ```

3.  **Install dependencies:**
    Go will automatically download the necessary dependencies when you build or run the project.

    ```bash
    go mod tidy
    ```

4.  **Run the application:**
    ```bash
    go run ./cmd/api/main.go
    ```

The server will start on `http://localhost:8080`.
