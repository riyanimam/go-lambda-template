# go-lambda-template

[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![AWS Lambda](https://img.shields.io/badge/AWS-Lambda-FF9900?style=flat&logo=amazon-aws)](https://aws.amazon.com/lambda/)
[![golangci-lint](https://img.shields.io/badge/golangci--lint-1.63-blue?style=flat)](https://golangci-lint.run/)
[![Code Style](https://img.shields.io/badge/code%20style-gofmt-blue.svg)](https://pkg.go.dev/cmd/gofmt)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](./DEVELOPMENT.md)

A production-ready Go template for AWS Lambda functions with comprehensive tooling, testing, and CI/CD pipelines.

## Features

- 🚀 **AWS Lambda Ready**: Pre-configured for AWS Lambda deployment with custom runtime
- 🛠️ **Modern Tooling**: golangci-lint, gofmt, lefthook for code quality
- 🧪 **Testing**: Built-in test infrastructure with coverage reporting
- 📦 **CI/CD**: GitHub Actions workflows for testing, security scanning, and deployment
- 🔒 **Security**: Integrated gosec, govulncheck, and CodeQL scanning
- 📝 **Documentation**: Comprehensive development and deployment guides

## Prerequisites

- **Go**: 1.23 or higher
- **Git**: 2.30 or higher
- **AWS CLI**: Configured with appropriate credentials
- **Terraform**: 1.10 or higher (for infrastructure deployment)

## Quick Start

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/riyanimam/go-lambda-template.git
   cd go-lambda-template
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up Git hooks (optional but recommended):

   ```bash
   # Install lefthook
   go install github.com/evilmartians/lefthook@latest

   # Install hooks
   lefthook install
   ```

### Development

1. **Build the function**:

   ```bash
   go build -o bootstrap ./src/main.go
   ```

2. **Run tests**:

   ```bash
   go test -v ./...
   ```

3. **Run tests with coverage**:

   ```bash
   go test -v -race -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out
   ```

4. **Lint code**:

   ```bash
   # Install golangci-lint
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

   # Run linter
   golangci-lint run ./...
   ```

5. **Format code**:

   ```bash
   gofmt -l -s -w .
   ```

### Building for AWS Lambda

Build the Lambda function for deployment:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w" -o bootstrap ./src/main.go
zip lambda-function.zip bootstrap
```

The `-ldflags="-s -w"` flags reduce the binary size by removing debug information and symbol tables.

### Deployment

Deploy using Terraform:

```bash
cd terraform
terraform init
terraform plan
terraform apply
```

See [terraform/README.md](terraform/README.md) for detailed deployment instructions.

## Project Structure

```
.
├── .github/
│   └── workflows/          # GitHub Actions CI/CD workflows
├── src/
│   ├── main.go            # Lambda handler implementation
│   └── main_test.go       # Handler tests
├── terraform/             # Infrastructure as Code
├── .golangci.yml          # golangci-lint configuration
├── lefthook.yml           # Git hooks configuration
├── go.mod                 # Go module dependencies
└── README.md              # This file
```

## Available Commands

| Command                                      | Description                      |
| -------------------------------------------- | -------------------------------- |
| `go build -o bootstrap ./src/main.go`        | Build the Lambda function        |
| `go test -v ./...`                           | Run all tests                    |
| `go test -v -race -coverprofile=coverage.out ./...` | Run tests with race detection and coverage |
| `golangci-lint run ./...`                    | Run comprehensive linting        |
| `gofmt -l -s -w .`                           | Format all Go code               |
| `go mod tidy`                                | Update dependencies              |
| `lefthook run pre-commit`                    | Run pre-commit hooks manually    |

## CI/CD Workflows

This template includes comprehensive GitHub Actions workflows:

- **CI**: Multi-version testing (Go 1.22.x, 1.23.x) and artifact building
- **Code Quality**: Linting, formatting checks, and test coverage
- **Security**: gosec, govulncheck, and CodeQL scanning
- **PR Validation**: Conventional commit message validation

## Contributing

Contributions are welcome! Please see [DEVELOPMENT.md](./DEVELOPMENT.md) for development guidelines.

1. Create a feature branch
2. Make your changes
3. Run tests and linting
4. Submit a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Support

For questions or issues:

- Open an issue on GitHub
- Review the [DEVELOPMENT.md](./DEVELOPMENT.md) guide
- Check existing issues and discussions