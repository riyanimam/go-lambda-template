# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Security

- **CRITICAL**: Upgraded Go to v1.24 to fix 9 security vulnerabilities in standard library
  - GO-2025-4175: Improper DNS name constraint validation in crypto/x509
  - GO-2025-4155: Excessive resource consumption in crypto/x509 error string printing
  - GO-2025-4013: Panic when validating DSA certificates in crypto/x509
  - GO-2025-4012: Memory exhaustion from unlimited cookie parsing in net/http
  - GO-2025-4011: Memory exhaustion from DER payload parsing in encoding/asn1
  - GO-2025-4010: IPv6 hostname validation bypass in net/url
  - GO-2025-4009: Quadratic complexity in PEM parsing in encoding/pem
  - GO-2025-4008: ALPN negotiation information leakage in crypto/tls
  - GO-2025-4007: Quadratic complexity in name constraint checking in crypto/x509

### Changed

- Updated Go to version 1.24 (latest stable) - fixes 9 critical vulnerabilities
- Updated CI testing matrix to Go 1.23.x and 1.24.x
- Updated golangci-lint target version to 1.24
- Updated `github.com/stretchr/testify` to v1.11.1 (from v1.7.2)
- Updated all documentation to reflect Go 1.23 requirement

### Added

- Initial project setup with Go 1.23 support
- AWS Lambda handler implementation
- Comprehensive CI/CD workflows
  - Multi-version testing (Go 1.22.x, 1.23.x)
  - Code quality checks (golangci-lint, gofmt)
  - Security scanning (gosec, govulncheck, CodeQL)
  - PR title validation
- Development tooling
  - golangci-lint configuration with 20+ linters
  - Lefthook for Git hooks (replaced Python pre-commit)
  - Pre-commit hooks (formatting, linting, YAML/Markdown)
  - Pre-push hooks (tests, build verification)
- Comprehensive documentation
  - README with badges and quick start guide
  - DEVELOPMENT.md with detailed workflow
  - CHANGELOG.md following Keep a Changelog format
- Testing infrastructure with coverage reporting
- Terraform deployment configuration

### Changed

- Migrated from Python pre-commit to Lefthook
- Updated CI workflow to use latest GitHub Actions (v4, v5)
- Improved .gitignore for Go projects

### Removed

- Python pre-commit configuration (.pre-commit-config.yaml)
- Legacy linting tools (golint)

## [0.1.0] - 2024-01-XX

### Added

- Initial release of go-lambda-template
- Basic AWS Lambda handler
- Terraform infrastructure code
- MIT License

[0.1.0]: https://github.com/riyanimam/go-lambda-template/releases/tag/v0.1.0
[unreleased]: https://github.com/riyanimam/go-lambda-template/compare/v0.1.0...HEAD
