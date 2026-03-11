# Go CRUD API With Gin

A small REST API built with Go, Gin, and MongoDB. The project currently exposes health checks and basic note creation and listing endpoints.

## Stack

- Go 1.25
- Gin
- MongoDB Go Driver v2
- godotenv
- Air for local live reload

## Project Structure

```text
.
|-- cmd/
|   `-- api/
|       `-- main.go
|-- internal/
|   |-- config/
|   |   `-- config.go
|   |-- db/
|   |   `-- mongo.go
|   |-- handler/
|   |   `-- notes_handler.go
|   |-- models/
|   |   `-- notes.go
|   |-- repository/
|   |   `-- note_repository.go
|   `-- server/
|       `-- router.go
|-- tmp/
|-- .air.toml
|-- .gitignore
|-- go.mod
`-- README.md
```

## Features

- Loads configuration from environment variables and `.env`
- Connects to MongoDB on startup
- Exposes a root route and health check route
- Supports creating notes
- Supports listing all notes

## Environment Variables

Create a `.env` file in the project root:

```env
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=notesdb
PORT=8080
```

Required variables:

- `MONGO_URI`: MongoDB connection string
- `MONGO_DB_NAME`: Database name used by the API
- `PORT`: HTTP server port

## Getting Started

### 1. Install dependencies

```bash
go mod download
```

### 2. Run the API

```bash
go run ./cmd/api
```

The server starts at:

```text
http://localhost:8080
```

If you set a different `PORT`, use that instead.

## Run With Air

Install Air once:

```bash
go install github.com/air-verse/air@latest
```

Then start live reload from the project root:

```bash
air
```

This repository already includes `.air.toml` configured to build `./cmd/api` into `tmp/api.exe`.

## API Endpoints

### `GET /`

Returns a simple welcome response.

Example response:

```json
{
  "success": true,
  "message": "Welcome to Go CRUD API"
}
```

### `GET /health`

Returns application health status.

Example response:

```json
{
  "status": "ok"
}
```

### `GET /notes`

Returns all notes from the `notes` collection.

Example response:

```json
{
  "success": true,
  "data": [
    {
      "id": "...",
      "title": "First note",
      "content": "Example content",
      "pinned": false,
      "created_at": "2026-03-11T10:00:00Z",
      "updated_at": "2026-03-11T10:00:00Z"
    }
  ]
}
```

### `POST /notes`

Creates a new note.

Request body:

```json
{
  "title": "My note",
  "content": "Write something useful here",
  "pinned": true
}
```

Successful response:

```json
{
  "success": true,
  "data": {
    "id": "...",
    "title": "My note",
    "content": "Write something useful here",
    "pinned": true,
    "created_at": "2026-03-11T10:00:00Z",
    "updated_at": "2026-03-11T10:00:00Z"
  }
}
```

Validation rules:

- `title` is required
- `content` is required
- `pinned` is optional

## Example cURL Requests

Create a note:

```bash
curl -X POST http://localhost:8080/notes \
	-H "Content-Type: application/json" \
	-d '{"title":"My note","content":"Hello from Gin","pinned":false}'
```

List notes:

```bash
curl http://localhost:8080/notes
```

Health check:

```bash
curl http://localhost:8080/health
```

## Notes

- The API uses the `notes` collection in MongoDB.
- Configuration loading fails fast if any required environment variable is missing.
- The current implementation includes create and list operations only.
