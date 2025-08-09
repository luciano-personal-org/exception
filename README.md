# Exception Package

A comprehensive error handling package for Go applications, specifically designed for algorithmic trading platforms. This package provides structured error types and flexible error handling functions with optional logging and panic control.

## Features

- 🎯 **Structured Error Types**: Rich error contexts with codes, details, and original error chaining
- 🔍 **Flexible Error Handling**: Functions with panic/no-panic variants
- 📝 **Integrated Logging**: Optional structured logging support using `logr`
- 💼 **Trading-Specific**: Purpose-built for financial and trading applications
- 🔧 **Simple API**: Easy-to-use interface with clear naming conventions

## Installation

```bash
go get github.com/luciano-personal-org/exception
```

## Quick Start

```go
package main

import (
    "errors"
    "fmt"
    "github.com/luciano-personal-org/exception"
)

func main() {
    // Create a trading error
    tradingErr := exception.NewTradingError("INVALID_SYMBOL", "Stock symbol not found")
    tradingErr.SetDetails("Symbol 'XYZ' is not listed on any exchange")
    tradingErr.SetOriginalError(errors.New("database lookup failed"))
    
    // Handle error without panicking
    exception.DontPanic(tradingErr)
    
    // Or handle with panic (use carefully)
    // exception.DoPanic(tradingErr)
}
```

## Error Types

### TradingError Interface

The primary error type for trading-related operations:

```go
type TradingError interface {
    Details() string
    ErrorCode() string
    Error() string
    OriginalError() error
    SetOriginalError(err error)
    SetDetails(details string)
}
```

**Example Usage:**

```go
// Create a new trading error
err := exception.NewTradingError("ORDER_REJECTED", "Insufficient funds")
err.SetDetails("Account balance: $100.00, Order amount: $150.00")
err.SetOriginalError(errors.New("account validation failed"))

fmt.Println(err.Error())         // "Insufficient funds"
fmt.Println(err.ErrorCode())     // "ORDER_REJECTED"
fmt.Println(err.Details())       // "Account balance: $100.00, Order amount: $150.00"
```

### ValidationError Struct

Simple validation error for input validation:

```go
type ValidationError struct {
    Message string
}
```

**Example Usage:**

```go
validationErr := exception.ValidationError{
    Message: "Price must be greater than zero",
}
fmt.Println(validationErr.Error()) // "Price must be greater than zero"
```

## Error Handling Functions

### Basic Error Handling

#### `DoPanic(err interface{})`
Panics if the error is not nil. Provides detailed output for TradingError types.

```go
err := exception.NewTradingError("CONNECTION_FAILED", "Cannot connect to exchange")
exception.DoPanic(err) // Will panic with detailed error information
```

#### `DontPanic(err interface{})`
Logs error details without panicking. Safe for production use.

```go
err := exception.NewTradingError("RATE_LIMITED", "API rate limit exceeded")
exception.DontPanic(err) // Logs error details, continues execution
```

### Logging-Enhanced Error Handling

#### `DoPanicWithLog(err interface{}, logger logr.Logger)`
Panics with structured logging support.

```go
import "github.com/go-logr/logr"

logger := // ... initialize your logger
err := exception.NewTradingError("MARKET_CLOSED", "Market is currently closed")
exception.DoPanicWithLog(err, logger) // Logs then panics
```

#### `DontPanicWithLog(err interface{}, logger logr.Logger)`
Logs errors using structured logging without panicking.

```go
logger := // ... initialize your logger
err := exception.NewTradingError("STALE_PRICE", "Price data is outdated")
exception.DontPanicWithLog(err, logger) // Structured logging, no panic
```

## Common Use Cases

### Trading Application Error Handling

```go
func ProcessOrder(symbol string, quantity int, price float64) error {
    // Validate inputs
    if quantity <= 0 {
        return exception.ValidationError{
            Message: "Quantity must be positive",
        }
    }
    
    // Business logic that might fail
    if !isMarketOpen() {
        tradingErr := exception.NewTradingError("MARKET_CLOSED", "Cannot process order")
        tradingErr.SetDetails(fmt.Sprintf("Market hours: 9:30 AM - 4:00 PM EST"))
        return tradingErr
    }
    
    // Process order...
    return nil
}

func main() {
    err := ProcessOrder("AAPL", -10, 150.0)
    if err != nil {
        // Log error without crashing the application
        exception.DontPanic(err)
    }
}
```

### Error Chain Tracking

```go
func connectToDatabase() error {
    return errors.New("connection timeout")
}

func fetchMarketData() error {
    originalErr := connectToDatabase()
    if originalErr != nil {
        tradingErr := exception.NewTradingError("DATA_UNAVAILABLE", "Cannot fetch market data")
        tradingErr.SetOriginalError(originalErr)
        tradingErr.SetDetails("Database connection failed after 3 retries")
        return tradingErr
    }
    return nil
}
```

## Best Practices

1. **Use TradingError for Business Logic**: Reserve TradingError for domain-specific errors in your trading application.

2. **Chain Errors Properly**: Always set the original error when wrapping errors to maintain the error chain.

3. **Choose Panic vs No-Panic Carefully**: 
   - Use `DoPanic*` functions only for truly unrecoverable errors
   - Use `DontPanic*` functions for errors that can be handled gracefully

4. **Provide Meaningful Error Codes**: Use consistent, descriptive error codes that can be used for error handling logic.

5. **Include Contextual Details**: Add relevant context in the Details field to help with debugging.

## Error Output Format

When using TradingError with the panic functions, you'll get structured output like:

```
Custom Error: Insufficient funds, Code: ORDER_REJECTED
Details: Account balance: $100.00, Order amount: $150.00
Original Error: account validation failed
```

## Dependencies

- [github.com/go-logr/logr](https://github.com/go-logr/logr) v1.4.2 - Structured logging interface

## Contributing

This package is part of a larger algorithmic trading platform. When contributing:

1. Maintain backward compatibility
2. Add tests for new functionality
3. Follow existing naming conventions
4. Update documentation for new features

## License

Part of the Algorithmic Trading Platform - see project root for license details.