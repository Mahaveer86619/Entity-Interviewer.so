# Entity Interviewer.so

Entity Interviewer.so is a streamlined web application that generates personalized interview questions from job descriptions and provides real-time Text-to-Speech (TTS) conversion. It utilizes minimal Go packages and HTMX for dynamic UI interactions, along with a reusable admin microservice for monitoring and dashboarding.

## Key Technologies

-   **Backend Logic:** Go (Golang) with the standard library (net/http) and HTMX for server-side rendering.
-   **Dynamic UI:** HTMX for AJAX-driven UI interactions.
-   **Database:** PostgreSQL for robust relational data storage.
-   **Microservices:** Python (FastAPI) for specialized NLP and TTS tasks.
-   **NLP Pipelines:** Hugging Face Transformers for text processing.
-   **Question Generation:** Language Learning Models (LLMs) for contextual question generation.
-   **TTS Conversion:** F5 TTS for high-quality audio output.
-   **Containerization:** Docker for consistent deployments.
-   **Database Management:** Adminer for PostgreSQL.

## Key Features

-   **Interview Question Generation:** Generate job-specific interview questions from user input.
-   **Real-time TTS Conversion:** Convert generated questions to speech in real-time.
-   **Secure Authentication:** Implement secure user authentication and session management.
-   **Dynamic UI:** Provide a dynamic, interactive, and responsive user interface using HTMX.
-   **Scalable Architecture:** Design a scalable and efficient microservices architecture.
-   **Database Management:** Use Adminer for database management and inspection.
-   **Reusable Admin Service:** Develop a reusable admin microservice for health checks and dashboarding.

## Development Steps

1.  **PostgreSQL Setup:** Set up PostgreSQL for data storage and define database tables.
2.  **Python Microservices:** Build FastAPI microservices for NLP and TTS tasks.
    -   Integrate Hugging Face Transformers for text processing.
    -   Implement LLM-based question generation.
    -   Integrate F5 TTS for audio output.
3.  **Go Backend Development:** Develop the Go backend using `net/http` and HTMX.
    -   Implement database connections and API endpoints.
4.  **Frontend Development:** Design the UI with HTMX for dynamic interactions.
    -   Implement input forms and audio playback for TTS output.
5.  **Dockerization:** Containerize each microservice using Docker.
    -   Use Docker Compose to orchestrate services.
6.  **Testing and Deployment:** Write unit and integration tests, then deploy on a cloud platform.
7.  **Adminer Integration:** Integrate Adminer for database management.
8.  **Admin Service Development:** Build a reusable Go admin service.

## Security Considerations

-   **Input Validation:** Implement rigorous input validation to prevent injection attacks.
-   **JWT Tokens:** Use JWT tokens for secure authentication and authorization.

## Microservice Descriptions

-   `go-server`: Go backend application.
-   `python-llm`: Python microservice for LLM-based question generation.
-   `python-tts`: Python microservice for TTS conversion.
-   `postgres`: PostgreSQL database service.
-   `adminer`: Database management tool.
-   `go-admin`: Reusable Go admin service.

## Docker Compose Services

-   `go-server`
-   `python-llm`
-   `python-tts`
-   `postgres`
-   `adminer`
-   `go-admin`

## Project File Structure

```
Interviewer.so/
├── go-server/             # Go backend application
│   ├── cmd/
│   │   └── main.go       # Entry point for the Go application
│   ├── internal/
│   │   ├── handlers/     # HTTP handlers
│   │   │   ├── api.go
│   │   │   └── web.go
│   │   ├── services/     # Business logic services
│   │   │   ├── question_service.go
│   │   │   └── tts_service.go
│   │   ├── models/       # Data models
│   │   │   └── models.go
│   │   ├── config/       # Configuration files
│   │   │   └── config.go
│   │   └── middleware/   # Middleware functions
│   │       └── auth.go
│   ├── static/           # Static files (HTML, CSS, JS)
│   │   ├── index.html    # Main HTML page
│   │   └── ...
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── python-llm/             # Python LLM microservice
│   ├── app/
│   │   ├── main.py       # FastAPI application
│   │   └── models.py     # Pydantic models
│   ├── requirements.txt
│   └── Dockerfile
├── python-tts/             # Python TTS microservice
│   ├── app/
│   │   ├── main.py       # FastAPI application
│   │   └── models.py     # Pydantic models
│   ├── requirements.txt
│   └── Dockerfile
├── go-admin/              # Reusable Go Admin Service
│   ├── cmd/
│   │   └── main.go       # Entry point for the Admin Service
│   ├── internal/
│   │   ├── handlers/     # HTTP handlers for registration and dashboard
│   │   │   ├── admin.go
│   │   ├── services/     # Health check and data management services
│   │   │   ├── health_check.go
│   │   │   ├── data_service.go
│   │   ├── models/       # Data models
│   │   │   └── models.go
│   ├── static/           # Static files for the admin dashboard
│   │   ├── dashboard.html
│   │   └── ...
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
├── docker-compose.yml     # Docker Compose configuration
└── .env                  # Environment variables
```

## How to Contribute and Use It Yourself

**Prerequisites:**

-   Docker and Docker Compose
-   Go (Golang)
-   Python 3.x
-   PostgreSQL

**Setup:**

1.  **Clone the repository:**

    ```bash
    git clone [repository URL]
    cd Interviewer.so
    ```

2.  **Set up environment variables:**

    -   Create a `.env` file in the root directory.
    -   Add necessary environment variables (e.g., database connection strings, API keys).

3.  **Start the services using Docker Compose:**

    ```bash
    docker-compose up --build
    ```

4.  **Access the application:**

    -   Open your web browser and navigate to `http://localhost:[port]` (depending on your Go server configuration).
    -   Access Adminer via `http://localhost:[adminer port]`.
    -   Access the admin panel via `http://localhost:[go-admin port]`

**Contribution Guidelines:**

1.  **Fork the repository:** Create your own fork of the project.
2.  **Create a branch:** Create a new branch for your feature or bug fix.
3.  **Make your changes:** Implement your changes and ensure they are well-documented.
4.  **Test your changes:** Write unit and integration tests to verify your changes.
5.  **Commit your changes:** Commit your changes with clear and concise commit messages.
6.  **Push your changes:** Push your changes to your fork.
7.  **Create a pull request:** Submit a pull request to the main repository.

**Running the Go Services Locally (Without Docker):**

1.  **Navigate to the Go server directory:**

    ```bash
    cd go-server/cmd
    ```

2.  **Run the Go server:**

    ```bash
    go run main.go
    ```

3.  Repeat steps for the go-admin service.
**Running the python services locally(without docker):**

1. navigate to the python services directory
2. create a virtual enviroment.
3. install the requirements.txt with pip install -r requirements.txt
4. run the main.py file.