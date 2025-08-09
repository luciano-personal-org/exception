# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go package (`github.com/luciano-personal-org/exception`) that provides error handling utilities for an algorithmic trading platform. The package is part of a larger mono-repository structure (`00-algo-platform/03-packages/exception`).

## Architecture

The package implements a structured error handling system with three main components:

### Core Error Types

1. **TradingError Interface** (`trading_error.go`): 
   - Primary error interface with methods: `Details()`, `ErrorCode()`, `Error()`, `OriginalError()`, `SetOriginalError()`, `SetDetails()`
   - Implemented by `tradingErrorImpl` struct with fields for error code, message, details, and original error
   - Constructor: `NewTradingError(errorCode string, message string)`

2. **ValidationError Struct** (`validation_error.go`):
   - Simple validation error with a `Message` field
   - Implements standard `error` interface with `Error()` method

### Error Handling Functions

**Panic Functions** (`panic.go`):
- `DoPanic(err interface{})`: Panics on any error, with special handling for TradingError
- `DontPanic(err interface{})`: Logs error details without panicking
- `DoPanicWithLog(err interface{}, logger logr.Logger)`: Panic with structured logging
- `DontPanicWithLog(err interface{}, logger logr.Logger)`: Log without panicking

## Dependencies

- `github.com/go-logr/logr v1.4.2`: For structured logging interface

## Common Commands

### Build and Test
```bash
# Build the module
go build

# Run tests (no test files currently exist)
go test

# Format code
go fmt ./...

# Lint and vet code
go vet ./...

# Tidy dependencies
go mod tidy

# Download dependencies
go mod download
```

### Development Commands
```bash
# Check module dependencies
go list -m all

# View module info
go mod graph

# Verify dependencies
go mod verify
```

## Key Patterns

- All error handling functions accept `interface{}` to handle any error type
- TradingError provides rich error context with codes, details, and original errors
- Functions come in paired variants (panic/no-panic, with/without logging)
- Consistent error message formatting using `fmt.Sprintf`
- Conditional logging based on logger sink availability