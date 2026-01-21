# Development Guide

This guide covers the development workflow for the go-lambda-template project.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Testing](#testing)
- [Code Quality](#code-quality)
- [Git Workflow](#git-workflow)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)

## Prerequisites

### Required Software

- **Go**: 1.24 or higher

  - Download from [go.dev](https://go.dev/dl/)
  - Verify installation: `go version`

- **Git**: 2.30 or higher

  - Download from [git-scm.com](https://git-scm.com/)
  - Verify installation: `git --version`

- **AWS CLI**: Latest version

  - Install from [AWS CLI Installation Guide](https://aws.amazon.com/cli/)
  - Configure with `aws configure`

- **Terraform**: 1.10 or higher

  - Download from [terraform.io](https://www.terraform.io/downloads)
  - Verify installation: `terraform version`

### Optional but Recommended

- **golangci-lint**: Comprehensive linter

  ```bash
  go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
  ```

- **lefthook**: Git hooks manager

  ```bash
  go install github.com/evilmartians/lefthook@latest
  ```

- **govulncheck**: Vulnerability scanner

  ```bash
  go install golang.org/x/vuln/cmd/govulncheck@latest
  ```

## Getting Started

### Initial Setup

1. **Clone the repository**:

   ```bash
   git clone https://github.com/riyanimam/go-lambda-template.git
   cd go-lambda-template
   ```

1. **Install dependencies**:

   ```bash
   go mod download
   go mod verify
   ```

1. **Install development tools**:

   ```bash
   # Install golangci-lint
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

   # Install lefthook
   go install github.com/evilmartians/lefthook@latest

   # Install govulncheck
   go install golang.org/x/vuln/cmd/govulncheck@latest
   ```

1. **Set up Git hooks** (recommended):

   ```bash
   lefthook install
   ```

   This installs pre-commit and pre-push hooks that will:

   - Format code with `gofmt`
   - Run `go vet`
   - Check Terraform formatting
   - Lint YAML and Markdown files
   - Run tests before pushing

### Environment Setup

Create a `.env` file in the project root for local development (already in `.gitignore`):

```bash
AWS_REGION=us-east-1
AWS_PROFILE=your-profile-name
# Add other environment-specific variables here
```

## Development Workflow

### Code Organization

```
src/
├── main.go        # Lambda handler entry point
└── main_test.go   # Handler tests
```

### Building

#### Local Build

```bash
# Standard build
go build -o bootstrap ./src/main.go

# Build with optimizations (smaller binary)
go build -ldflags="-s -w" -o bootstrap ./src/main.go
```

#### Build for AWS Lambda

```bash
# Build for Linux AMD64 (AWS Lambda runtime)
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -tags lambda.norpc -ldflags="-s -w" -o bootstrap ./src/main.go

# Create deployment package
zip lambda-function.zip bootstrap
```

**Build flags explained**:

- `GOOS=linux GOARCH=amd64`: Target Linux AMD64 architecture
- `CGO_ENABLED=0`: Disable CGO for static binary
- `-tags lambda.norpc`: Use optimized Lambda runtime
- `-ldflags="-s -w"`: Strip debug info and symbol table (reduces size)

### Running Locally

You can test the Lambda function locally using the AWS SAM CLI:

```bash
# Install AWS SAM CLI
# https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html

# Invoke function locally
sam local invoke -e event.json
```

## Testing

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run tests with coverage
go test -v -race -coverprofile=coverage.out ./...

# View coverage report in terminal
go tool cover -func=coverage.out

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage.html
```

### Test Coverage Goals

- **Minimum**: 80% coverage
- **Target**: 90%+ coverage
- Coverage reports are generated in CI/CD pipelines

### Writing Tests

Follow Go testing best practices:

```go
package main

import (
    "context"
    "testing"
)

func TestHandler(t *testing.T) {
    ctx := context.Background()
    
    result, err := handler(ctx)
    
    if err != nil {
        t.Fatalf("handler returned error: %v", err)
    }
    
    expected := "Hello, world!"
    if result != expected {
        t.Errorf("handler returned %q, want %q", result, expected)
    }
}
```

## Code Quality

### Formatting

Always format code before committing:

```bash
# Format all Go files
gofmt -l -s -w .

# Check formatting (returns files that need formatting)
gofmt -l -s .
```

### Linting

Run the comprehensive linter:

```bash
# Run golangci-lint with all configured linters
golangci-lint run ./...

# Run with auto-fix where possible
golangci-lint run --fix ./...

# Run specific linters
golangci-lint run --disable-all --enable=errcheck,gosimple,govet ./...
```

The project uses a comprehensive `.golangci.yml` configuration with 20+ linters.

### Static Analysis

```bash
# Run go vet
go vet ./...

# Check for vulnerabilities
govulncheck ./...
```

### Pre-commit Checks

If you installed lefthook, these checks run automatically before each commit:

- Code formatting (`gofmt`)
- Static analysis (`go vet`)
- Terraform formatting
- YAML linting
- Markdown formatting
- Trailing whitespace check

Run checks manually:

```bash
lefthook run pre-commit
```

### Pre-push Checks

These checks run automatically before each push:

- Full test suite with race detection
- Build verification

Run checks manually:

```bash
lefthook run pre-push
```

## Git Workflow

### Branch Naming

Follow this convention:

- `feature/description` - New features
- `fix/description` - Bug fixes
- `docs/description` - Documentation updates
- `refactor/description` - Code refactoring
- `test/description` - Test additions/updates

### Commit Messages

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Types**: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert

**Examples**:

```
feat: Add structured logging to Lambda handler
fix: Resolve race condition in concurrent requests
docs: Update deployment guide with new IAM permissions
```

### Pull Request Process

1. **Create a feature branch**:

   ```bash
   git checkout -b feature/your-feature-name
   ```

1. **Make changes and commit**:

   ```bash
   git add .
   git commit -m "feat: Add your feature"
   ```

1. **Push to GitHub**:

   ```bash
   git push origin feature/your-feature-name
   ```

1. **Create Pull Request**:

   - Ensure all CI checks pass
   - Request review from team members
   - Address review comments
   - Merge when approved

## Deployment

### Local Deployment with Terraform

1. **Navigate to terraform directory**:

   ```bash
   cd terraform
   ```

1. **Initialize Terraform**:

   ```bash
   terraform init
   ```

1. **Review planned changes**:

   ```bash
   terraform plan
   ```

1. **Apply changes**:

   ```bash
   terraform apply
   ```

### CI/CD Deployment

The project uses GitHub Actions for automated deployment:

- **Pull Requests**: Run tests, linting, and security scans
- **Main Branch**: Run full test suite and create deployment artifacts

Deployment artifacts are available in the GitHub Actions runs.

## Troubleshooting

### Common Issues

#### Module Download Failures

```bash
# Clear module cache
go clean -modcache

# Re-download dependencies
go mod download
```

#### Build Failures on Lambda

Ensure you're building for the correct architecture:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./src/main.go
```

#### golangci-lint Issues

```bash
# Update to latest version
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Clear cache
golangci-lint cache clean
```

#### Lefthook Not Running

```bash
# Reinstall hooks
lefthook uninstall
lefthook install
```

### Getting Help

- Check existing [GitHub Issues](https://github.com/riyanimam/go-lambda-template/issues)
- Review [AWS Lambda Go documentation](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- Consult [Go documentation](https://go.dev/doc/)

## Best Practices

1. **Always run tests before pushing**
1. **Keep functions small and focused**
1. **Use meaningful variable and function names**
1. **Document exported functions and types**
1. **Handle errors explicitly**
1. **Use contexts for cancellation and timeouts**
1. **Keep dependencies minimal**
1. **Review security scan results**

## Additional Resources

- [Effective Go](https://go.dev/doc/effective_go)
- [AWS Lambda Go Best Practices](https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html)
- [golangci-lint Documentation](https://golangci-lint.run/)
- [Conventional Commits](https://www.conventionalcommits.org/)
