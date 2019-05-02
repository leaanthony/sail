package errors

import (
	"fmt"
)

// SError is like an error but just holds error types
type SError struct {
	Type int
	Message string
}

// Error returns a string representing the error
func (s *SError) Error() string {
	message := "Error[%d]"
	if s.Message != "" {
		message += ": %s"
	}
	return fmt.Sprintf(message, s.Type, s.Message)
}

// New returns a new Error of the given type
func New(Type int, message ...string) *SError {
	result := &SError{
		Type: Type,
	}
	if len(message) > 0 {
		result.Message = message[0]
	}
	return result
}

