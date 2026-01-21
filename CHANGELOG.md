# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Changed

- Updated Go to version 1.23 (latest stable)
- Updated golangci-lint to v1.63.4 (latest)
- Updated Terraform to v1.10.3 (latest)
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
