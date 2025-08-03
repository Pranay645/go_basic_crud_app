# Go Basic CRUD Application

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16+-4169E1?logo=postgresql)

A lightweight REST API implementation using Go's standard libraries for HTTP handling and PostgreSQL database operations.

## Features
- Full CRUD operations implementation
- RESTful API endpoints
- Middleware support for:
  - Request logging
  - Security headers
- PostgreSQL database integration
- Environment configuration
- Structured error handling

## Technology Stack
- **Language**: Go 1.21+
- **Database**: PostgreSQL 16+
- **HTTP Library**: net/http
- **SQL Driver**: github.com/lib/pq
- **Configuration**: Environment variables

## Installation
1. Clone repository:
```bash
git clone https://github.com/Pranay645/go_basic_crud_app.git
cd go_basic_crud_app
```

2. Database setup:
```sql
CREATE DATABASE go_crud;
```

3. Install dependencies:
```bash
go mod download
```

## Configuration
Create `.env` file:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=go_crud
```

## API Endpoints
| Method | Path | Description |
|--------|------|-------------|
| POST   | /cars | Create new item |
| GET    | /cars/{id} | Get single item |
| PUT    | /cars | Update item |
| DELETE | /cars/{id} | Delete item |

**Example Request:**
```bash
curl -X POST -H "Content-Type: application/json" \
-d '{"name":"New Item","description":"Sample description"}' \
http://localhost:8080/cars
```

## Project Structure
```
.
├── handlers/      # HTTP request handlers
├── middleware/    # Custom middleware
├── models/        # Data models
├── repository/    # Database operations
├── main.go        # Application entry point
├── go.mod         # Dependency management
└── .env.example   # Environment template
```

## Running the Application
```bash
# Start development server
go run main.go

# Build binary
go build -o app
```

## License
MIT License - See [LICENSE](LICENSE) for details

## Contributing
Contributions welcome! Please open an issue first to discuss proposed changes.