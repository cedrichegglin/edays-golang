# edays-golang

A Go application demonstrating best practices for project structure, testing, and deployment.

## 🚀 Features

- **Clean Architecture**: Well-organized project structure following Go best practices
- **HTTP Server**: RESTful API with middleware support
- **Configuration Management**: Environment-based configuration
- **Logging**: Structured JSON logging with configurable levels
- **Testing**: Comprehensive test coverage with benchmarks
- **Docker Support**: Containerized deployment
- **CI/CD**: GitHub Actions workflow for testing and building
- **API Documentation**: OpenAPI 3.0 specification

## 📁 Project Structure

```
edays-golang/
├── cmd/                    # Application entrypoints
│   └── edays/             # Main application
├── pkg/                   # Public packages
│   └── greetings/         # Greeting functionality
├── internal/              # Private application code
│   ├── config/           # Configuration management
│   ├── logger/           # Logging utilities
│   └── middleware/       # HTTP middleware
├── api/                   # API specifications
│   └── v1/               # API version 1
├── examples/              # Usage examples
├── docs/                  # Documentation
├── scripts/               # Build and deployment scripts
├── .github/               # GitHub workflows
├── go.mod                 # Go module definition
├── go.sum                 # Go module checksums
├── Makefile              # Build automation
├── Dockerfile            # Container definition
└── README.md             # This file
```

## 🛠️ Getting Started

### Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/cedrichegglin/edays-golang.git
cd edays-golang
```

2. Download dependencies:
```bash
go mod download
```

3. Run the application:
```bash
go run ./cmd/edays
```

Or using Make:
```bash
make run
```

### Configuration

The application can be configured using environment variables:

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | 8080 | Server port |
| `LOG_LEVEL` | info | Log level (debug, info, warn, error) |
| `READ_TIMEOUT` | 10 | HTTP read timeout in seconds |
| `WRITE_TIMEOUT` | 10 | HTTP write timeout in seconds |
| `ENVIRONMENT` | development | Environment (development, production) |

## 🧪 Testing

Run tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
make test-coverage
```

Run benchmarks:
```bash
go test -bench=. ./...
```

## 🏗️ Building

Build the application:
```bash
make build
```

Build for Linux:
```bash
make build-linux
```

## 🐳 Docker

Build Docker image:
```bash
make docker-build
```

Run with Docker:
```bash
docker run -p 8080:8080 edays
```

## 📚 API Documentation

### Endpoints

#### Health Check
```http
GET /health
```

Returns the health status of the application.

**Response:**
```
200 OK
OK
```

#### Get Greeting
```http
GET /api/v1/greeting
```

Returns a random greeting message.

**Response:**
```json
{
  "message": "Hello, World!"
}
```

### OpenAPI Specification

The complete API specification is available in `api/v1/openapi.yaml`.

## 🛠️ Development

### Code Style

- Use `gofmt` for formatting
- Follow Go naming conventions
- Write comprehensive tests
- Document public APIs

### Available Make Commands

```bash
make help          # Show all available commands
make build         # Build the application
make test          # Run tests
make test-coverage # Run tests with coverage
make fmt           # Format code
make lint          # Lint code
make clean         # Clean build artifacts
make deps          # Download and tidy dependencies
make docs          # Generate documentation
```

### Installing Development Tools

```bash
make install-tools
```

This will install:
- golangci-lint for code linting
- goimports for import formatting

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go team for the excellent language and tooling
- The Go community for best practices and conventions