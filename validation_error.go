// Package exception contains the validation error exception.
package exception

// ValidationError struct for validation error.
type ValidationError struct {
	// Message is the validation error message.
	Message string
}

// Error returns the validation error message.
func (validationError ValidationError) Error() string {
	// Return the validation error message.
	return validationError.Message
}
