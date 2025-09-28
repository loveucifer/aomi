# Contributing to Aomi

Thank you for your interest in contributing to Aomi! This document outlines the process for contributing to this project.

## Table of Contents
- [Getting Started](#getting-started)
- [Development Workflow](#development-workflow)
- [Code Standards](#code-standards)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Issues](#issues)

## Getting Started

1. Fork the repository
2. Clone your fork:
   ```bash
   git clone https://github.com/YOUR_USERNAME/aomi.git
   cd aomi
   ```
3. Create a new branch for your feature or bug fix:
   ```bash
   git checkout -b feature/my-feature
   ```

## Development Workflow

1. Make your changes
2. Test your changes locally:
   ```bash
   go build -o aomi ./cmd/aomi/
   ./aomi sample.json test.csv
   ```
3. Add or update tests as necessary
4. Update documentation if needed
5. Run tests to ensure everything works:
   ```bash
   go test ./...
   ```

## Code Standards

- Follow Go formatting standards (`go fmt`)
- Add comments to exported functions and types
- Keep functions focused and readable
- Use meaningful variable and function names
- Follow the existing project structure and patterns

## Testing

- Add unit tests for new functionality
- Ensure all tests pass before submitting a pull request
- Test end-to-end functionality manually when possible

## Pull Request Process

1. Ensure your branch is up-to-date with the main branch:
   ```bash
   git fetch upstream
   git merge upstream/main
   ```
2. Run all tests and ensure they pass
3. Update the README or documentation if necessary
4. Submit your pull request with a clear description of your changes
5. Address any review comments

## Issues

When creating issues, please include:
- A clear, descriptive title
- Steps to reproduce the issue
- Expected vs actual behavior
- Version information
- Any relevant logs or error messages

## Questions?

If you have questions about contributing, feel free to open an issue with the "question" label.