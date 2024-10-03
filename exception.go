// Package exception provides functions for handling errors.
package exception

// PanicIfNeeded panics if the given error is not nil.
func PanicIfNeeded(err interface{}) {
	// Check if the given error is not nil.
	if err != nil {
		panic(err)
	}
}
