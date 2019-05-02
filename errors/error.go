package errors

import (
	"fmt"
)

// SErrorStruct is like an error but just holds error types
type SErrorStruct struct {
	Type int
	Message string
}

// Error returns a string representing the error
func (s *SErrorStruct) Error() string {
	message := "Error[%d]"
	if s.Message != "" {
		message += ": %s"
	}
	return fmt.Sprintf(message, s.Type, s.Message)
}

// SError returns a new Error of the given type
func SError(Type int, message ...string) *SErrorStruct {
	result := &SErrorStruct{
		Type: Type,
	}
	if len(message) > 0 {
		result.Message = message[0]
	}
	return result
}

