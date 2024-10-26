// Package exception provides functions for handling errors.
package exception

import (
	"fmt"
)

// PanicIfNeeded panics if the given error is not nil.
func PanicIfNeeded(err interface{}) {
	// Check if the given error is not nil.
	if err != nil {
		// Check if the error is of type TradingError.
		if te, ok := err.(TradingError); ok {
			// Print the message and code.
			fmt.Printf("Custom Error: %s, Code: %s\n", te.Error(), te.ErrorCode())
			fmt.Printf("Original Error: %s\n", te.OriginalError().Error())
		}
		// Panic with the error.
		panic(err)
	}
}
