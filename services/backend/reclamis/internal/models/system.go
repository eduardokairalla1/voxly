// * Data models for system informations. *

package models

// --- TYPES ---
type Health struct {
	Status string `json:"status"`
}

type Info struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
}
