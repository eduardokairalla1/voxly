// * Implementation of custom error. *

package errors

// --- IMPORTS ---
import (
	"fmt"
)

// --- CODE ---

// Error implements the error interface.
func (e *ReclamisError) Error() string {
	return fmt.Sprintf("[%d] %s (%s)", e.Code, e.Message, e.Kind)
}
