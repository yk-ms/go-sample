# Go Sample Project

This is a sample Go project configured to run in a Docker container using Docker Compose. The project includes a simple Go application and provides a Makefile for easy build and run commands.

## Project Structure

- `Dockerfile`: Defines the Docker image for the Go application.
- `docker-compose.yml`: Manages multi-container Docker applications.
- `Makefile`: Provides commands to build and run the application.

## Prerequisites

- Docker
- Docker Compose
- Make

## Setup and Usage

1. **Build the Docker Image**

   ```bash
   make build
   ```

2. **Run the Application**

   ```bash
   make up
   ```

   The application will be accessible at `http://localhost:8080`.

## Sample

The application provides a simple API to get item.

```bash
curl http://localhost:8080/api/todo
```

The response will be:

```json
[
  {
    "id": 1,
    "title": "Sample Item",
    "description": "This is a sample item.",
    "completed": false
  }
]
```

create item

```bash
curl -X POST http://localhost:8080/api/todo \
-H "Content-Type: application/json" \
-d '{"id":"1", "title":"sample title", "description":"sample description", "completed":false}'
```

## License

This project is licensed under the MIT License.
