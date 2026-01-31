// * Custom error types for the application. *

package errors

// --- TYPES & CONSTANTS ---
type Kind string

const (
	KindBaseReclamisError Kind = "BASE_RECLAMIS_ERROR"
)

type ReclamisError struct {
	Kind    Kind
	Message string
	Code    int
}
