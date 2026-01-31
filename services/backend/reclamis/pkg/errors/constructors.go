// * Constructors for custom error types. *

package errors

// --- CODE ---
func NewBaseReclamisError(message string) *ReclamisError {
	return &ReclamisError{
		Kind:    KindBaseReclamisError,
		Message: message,
		Code:    500,
	}
}
