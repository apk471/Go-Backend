# Go Backend Learning Repository

A collection of Go backend development examples and implementations demonstrating various concepts including CRUD operations, REST APIs, authentication patterns, caching strategies, and request handling.

## 📋 Table of Contents

- [Overview](#overview)
- [Projects](#projects)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Descriptions](#project-descriptions)
- [Technologies Used](#technologies-used)
- [Contributing](#contributing)

## 🎯 Overview

This repository serves as a comprehensive learning resource and reference implementation for building backend applications in Go. It covers fundamental to advanced concepts including database operations, authentication, authorization, caching, and various HTTP request handling patterns.

## 🚀 Projects

### 1. **CRUD**
A complete RESTful CRUD API with clean architecture, demonstrating user management with SQLite storage, validation, and structured logging.

**Key Features:**
- Clean architecture with separation of concerns
- SQLite database integration
- Input validation using `go-playground/validator`
- YAML-based configuration management
- Structured logging with `slog`
- Graceful server shutdown

**Tech Stack:** Go 1.25.5, SQLite3, go-playground/validator

[📖 Detailed Documentation](CRUD/README.md)

### 2. **REST**
A task/organization manager REST API built with MongoDB, demonstrating document database operations and RESTful principles.

**Key Features:**
- MongoDB integration for data persistence
- CRUD operations for organizations
- Index creation for optimized queries
- Repository pattern implementation

**Tech Stack:** Go 1.25.5, MongoDB (mongo-driver)

**Endpoints:**
- `GET /organizations` - List all organizations
- `POST /organizations` - Create new organization
- `GET /organizations/{id}` - Get organization by ID
- `PUT /organizations/{id}` - Update organization
- `DELETE /organizations/{id}` - Delete organization

### 3. **REST_Cache**
Enhanced REST API with Redis caching layer for improved performance and reduced database load.

**Key Features:**
- MongoDB for persistent storage
- Redis for caching frequently accessed data
- Cache invalidation strategies
- Reduced database queries through intelligent caching

**Tech Stack:** Go 1.25.5, MongoDB, Redis

### 4. **OAuth**
Implementation of OAuth 2.0 authentication flow using Google OAuth provider.

**Key Features:**
- OAuth 2.0 authorization code flow
- Google OAuth integration
- Session management after OAuth authentication
- Protected route middleware
- User info retrieval from OAuth provider

**Tech Stack:** Go 1.25.5, golang.org/x/oauth2

**Endpoints:**
- `/login` - Initiate OAuth flow
- `/oauth/callback` - OAuth callback handler
- `/protected` - Protected endpoint requiring authentication

### 5. **RBAC Auth (Role-Based Access Control)**
Demonstrates role-based access control with JWT authentication.

**Key Features:**
- JWT token generation and validation
- Role-based authorization middleware
- Multiple role support (Admin, User)
- Protected routes based on user roles

**Tech Stack:** Go 1.25.5, JWT (golang-jwt/jwt)

**Endpoints:**
- `/admin` - Admin-only access
- `/user` - User and Admin access

### 6. **Stateful Auth**
Session-based authentication using server-side session storage.

**Key Features:**
- Server-side session management
- Cookie-based session tracking
- Session creation and deletion
- Session middleware for protected routes

**Tech Stack:** Go 1.25.5

**Endpoints:**
- `/login` - User login with session creation
- `/protected` - Protected endpoint
- `/logout` - Session termination

### 7. **Stateless Auth**
JWT-based stateless authentication implementation.

**Key Features:**
- JWT token generation
- Token-based authentication
- No server-side session storage
- Middleware for token validation

**Tech Stack:** Go 1.25.5, JWT

**Endpoints:**
- `/login` - Login and receive JWT token
- `/protected` - Protected endpoint requiring valid JWT

### 8. **Transformation-Validations**
Demonstrates various validation and transformation patterns for incoming data.

**Key Features:**
- Syntactic validation (format checking)
- Semantic validation (business logic validation)
- Complex validation (cross-field validation)
- Simple transformations (normalization)
- Semantic transformations (model conversion)
- Complex transformations (calculation logic)

**Tech Stack:** Go 1.25.5

**Examples:**
- Email normalization and validation
- Password hashing
- Age validation
- Date range validation
- Price calculation with discounts

### 9. **Requests**
Basic HTTP request handling examples with frontend and backend separation.

**Key Features:**
- HTTP request/response handling
- Frontend-backend integration examples
- Basic routing patterns

### 10. **MultipartRequests**
Demonstrates handling of multipart form data and file uploads.

**Key Features:**
- File upload handling
- Multipart form data parsing
- Frontend form implementation
- File storage management

## 📦 Prerequisites

- **Go**: Version 1.25.5 or higher
- **MongoDB**: For REST and REST_Cache projects (optional, for those projects only)
- **Redis**: For REST_Cache project (optional)
- **SQLite3**: For CRUD project (usually comes with Go SQLite driver)

## 🏁 Getting Started

1. **Clone the repository:**
   ```bash
   git clone <repository-url>
   cd Go-Backend
   ```

2. **Choose a project:**
   Navigate to any project directory:
   ```bash
   cd CRUD
   # or
   cd REST
   # or any other project
   ```

3. **Install dependencies:**
   ```bash
   go mod download
   ```

4. **Run the project:**
   ```bash
   go run main.go
   ```
   
   Or for projects with custom configuration:
   ```bash
   go run main.go -config config/local.yaml
   ```

5. **Test the endpoints:**
   Use curl, Postman, or any HTTP client to test the API endpoints.

## 📖 Project Descriptions

### Authentication Patterns

This repository demonstrates four different authentication approaches:

1. **Stateful Auth**: Traditional session-based authentication with server-side storage
2. **Stateless Auth**: Modern JWT-based authentication without server state
3. **OAuth**: Third-party authentication using OAuth 2.0 (Google)
4. **RBAC Auth**: Role-based access control with JWT tokens

Choose the appropriate pattern based on your application requirements:
- Use **Stateful** for traditional web applications with server sessions
- Use **Stateless** for microservices and APIs
- Use **OAuth** for social login integration
- Use **RBAC** when you need fine-grained permission control

### Database Operations

- **CRUD**: Demonstrates SQLite operations with clean architecture
- **REST**: Shows MongoDB operations with the official Go driver
- **REST_Cache**: Adds Redis caching layer to MongoDB operations

### Request Handling

- **Requests**: Basic HTTP request handling
- **MultipartRequests**: File upload and multipart form handling
- **Transformation-Validations**: Input validation and transformation patterns

## 🛠 Technologies Used

- **Language**: Go 1.25.5
- **Databases**: 
  - MongoDB (mongo-driver)
  - SQLite3
  - Redis
- **Authentication**: 
  - JWT (golang-jwt/jwt)
  - OAuth 2.0 (golang.org/x/oauth2)
- **Validation**: go-playground/validator
- **Configuration**: cleanenv, godotenv
- **HTTP**: net/http (standard library)

## 📝 Common Patterns

### Running a Server
Most projects run on port `:8080` by default:
```bash
go run main.go
```

### Authentication Testing
For projects with authentication, you'll typically:
1. Call the `/login` endpoint to get credentials/token
2. Use the token/session in subsequent requests to `/protected` endpoints

### Configuration
Projects with configuration files (like CRUD) support:
- Environment variable: `CONFIG_PATH=config/local.yaml`
- Command-line flag: `-config config/local.yaml`

## 🤝 Contributing

This is a learning repository. Feel free to:
- Add new examples
- Improve existing implementations
- Fix bugs
- Enhance documentation

## 📄 Notes

- Each project is self-contained with its own `go.mod` file
- Most projects use the standard port `:8080` (except CRUD which uses `:8082`)
- The `REST_Cache/data/` directory contains MongoDB data files and is excluded from version control
- Projects demonstrate different architectural patterns - from simple handlers to clean architecture

## 🔒 Security Notes

These examples are for learning purposes. For production use:
- Use proper secret management (environment variables, secret managers)
- Implement rate limiting
- Add proper CORS configuration
- Use HTTPS/TLS
- Implement proper password hashing (bcrypt, argon2)
- Validate and sanitize all inputs
- Add proper logging and monitoring
- Use secure session storage
- Implement CSRF protection for stateful authentication

## 📚 Further Learning

Each project demonstrates specific concepts. To deepen your understanding:
- Read the code comments
- Modify the examples
- Combine patterns from different projects
- Add tests
- Implement additional features

---

**Happy Learning! 🚀**
