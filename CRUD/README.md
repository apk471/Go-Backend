# Go CRUD API

A RESTful CRUD API built with Go for managing users. This project demonstrates a clean architecture with SQLite storage, configuration management, and structured logging.

## Features

- Create users via POST endpoint
- Get user by ID
- Get all users
- Input validation using `go-playground/validator`
- SQLite database storage
- YAML-based configuration
- Structured logging with `slog`
- Graceful server shutdown
- Clean architecture with separation of concerns

## Project Structure

```
go-crud-api/
├── cmd/
│   └── api/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── http/
│   │   └── handlers/
│   │       └── api/
│   │           └── handle.go    # HTTP request handlers
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── types/
│   │   └── types.go            # Data models
│   └── utils/
│       └── response/
│           └── response.go      # HTTP response utilities
├── storage/
│   └── storage.db               # SQLite database (auto-generated)
├── go.mod                       # Go module dependencies
├── go.sum                       # Dependency checksums
└── README.md                    # This file
```

## Prerequisites

- Go 1.25.5 or higher
- SQLite3 (usually comes with Go SQLite driver)

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd go-crud-api
```

2. Install dependencies:

```bash
go mod download
```

3. Create a configuration file. You can use the example in `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

## Configuration

The application supports configuration via:

1. **Environment variable**: `CONFIG_PATH` pointing to the config file
2. **Command-line flag**: `-config` flag with the path to config file

### Configuration Options

- `env`: Environment name (e.g., "dev", "prod")
- `storage_path`: Path to SQLite database file
- `http_server.address`: Server address and port

### Example Usage

```bash
# Using command-line flag
go run cmd/api/main.go -config config/local.yaml

# Using environment variable
CONFIG_PATH=config/local.yaml go run cmd/api/main.go
```

## Running the Application

### Development Mode

```bash
go run cmd/api/main.go -config config/local.yaml
```

### Build and Run

```bash
# Build the binary
go build -o bin/api cmd/api/main.go

# Run the binary
./bin/api -config config/local.yaml
```

The server will start on the address specified in your config file (default: `localhost:8082`).

## API Endpoints

### Create User

**POST** `/api/users`

Creates a new user in the database.

**Request Body:**

```json
{
  "id": 0,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25
}
```

**Validation Rules:**

- `name`: Required, minimum 2 characters, maximum 100 characters
- `email`: Required, must be a valid email format
- `age`: Required, minimum 18, maximum 100

**Response (201 Created):**

```json
{
  "id": 1
}
```

**Response (400 Bad Request):**

```json
{
  "status_code": "error",
  "error": "name is required, email is not a valid email"
}
```

### Get User by ID

**GET** `/api/students/{id}`

Retrieves a user by their ID.

**Response (200 OK):**

```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john@example.com",
  "age": 25
}
```

**Response (500 Internal Server Error):**

```json
{
  "status_code": "error",
  "error": "no user found with id 999"
}
```

### Get All Users

**GET** `/api/students`

Retrieves all users from the database.

**Response (200 OK):**

```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "age": 25
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane@example.com",
    "age": 30
  }
]
```

## Database Schema

The application automatically creates the following table on startup:

```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    email TEXT,
    age INTEGER
)
```

## Dependencies

- [cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Configuration management
- [go-playground/validator](https://github.com/go-playground/validator) - Input validation
- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver
- Standard library packages: `net/http`, `log/slog`, `database/sql`

## Development

### Project Architecture

The project follows a clean architecture pattern:

- **cmd/**: Application entry points
- **internal/**: Private application code
  - **config/**: Configuration loading and management
  - **http/handlers/**: HTTP request handlers
  - **storage/**: Data persistence layer (interface and implementations)
  - **types/**: Domain models
  - **utils/**: Utility functions

### Adding New Features

1. **New Endpoint**: Add handler in `internal/http/handlers/api/`
2. **New Storage Method**: Add to `Storage` interface in `internal/storage/storage.go` and implement in `internal/storage/sqlite/sqlite.go`
3. **New Type**: Add to `internal/types/types.go`

## Graceful Shutdown

The application supports graceful shutdown. When you send an interrupt signal (Ctrl+C), the server will:

1. Stop accepting new requests
2. Wait up to 10 seconds for existing requests to complete
3. Shut down gracefully

## Logging

The application uses structured logging with Go's `slog` package. Logs include:

- Server startup and shutdown events
- Request handling information
- Error messages with context
- Storage operations
