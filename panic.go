// Package exception provides functions for handling errors.
package exception

import (
	"fmt"

	"github.com/go-logr/logr"
)

// PanicIfNeeded panics if the given error is not nil.
func DoPanic(err interface{}) {
	// Check if the given error is not nil.
	if err != nil {
		// Check if the error is of type TradingError.
		if te, ok := err.(TradingError); ok {
			// Print the message and code.
			fmt.Printf("Custom Error: %s, Code: %s\n", te.Error(), te.ErrorCode())
			fmt.Printf("Details: %s\n", te.Details())
			fmt.Printf("Original Error: %s\n", te.OriginalError().Error())
		}
		// Panic with the error.
		panic(err)
	}
}

// PanicIfNeeded panics if the given error is not nil.
func DontPanic(err interface{}) {
	// Check if the given error is not nil.
	if err != nil {
		// Check if the error is of type TradingError.
		if te, ok := err.(TradingError); ok {
			// Print the message and code.
			fmt.Printf("Custom Error: %s, Code: %s\n", te.Error(), te.ErrorCode())
			fmt.Printf("Details: %s\n", te.Details())
			fmt.Printf("Original Error: %s\n", te.OriginalError().Error())
		}
	}
}

// PanicIfNeeded panics if the given error is not nil.
func DoPanicWithLog(err interface{}, logger logr.Logger) {
	// Check if the given error is not nil.
	error_message := ""
	if err != nil {
		// Check if the error is of type TradingError.
		if te, ok := err.(TradingError); ok {
			// Print the message and code.
			if logger.GetSink() != nil {
				if te.Error() != "" {
					error_message = fmt.Sprintf("Custom Error: %s, ", te.Error())
					if te.ErrorCode() != "" {
						error_message = fmt.Sprintf("%sCode: %s, ", error_message, te.ErrorCode())
					}
					if te.Details() != "" {
						error_message = fmt.Sprintf("%sDetails: %s, ", error_message, te.Details())
					}
					if te.OriginalError().Error() != "" {
						error_message = fmt.Sprintf("%sOriginal Error: %s, ", error_message, te.OriginalError().Error())
					}
				}
				logger.Error(err.(error), error_message)
			}
		}
		// Panic with the error.
		panic(err)
	}
}

// PanicIfNeeded panics if the given error is not nil.
func DontPanicWithLog(err interface{}, logger logr.Logger) {
	error_message := ""
	if err != nil {
		// Check if the error is of type TradingError.
		if te, ok := err.(TradingError); ok {
			// Print the message and code.
			if logger.GetSink() != nil {
				if te.Error() != "" {
					error_message = fmt.Sprintf("Custom Error: %s, ", te.Error())
					if te.ErrorCode() != "" {
						error_message = fmt.Sprintf("%sCode: %s, ", error_message, te.ErrorCode())
					}
					if te.Details() != "" {
						error_message = fmt.Sprintf("%sDetails: %s, ", error_message, te.Details())
					}
					if te.OriginalError().Error() != "" {
						error_message = fmt.Sprintf("%sOriginal Error: %s, ", error_message, te.OriginalError().Error())
					}
				}
				logger.Error(err.(error), error_message)
			}
		}
	}
}
