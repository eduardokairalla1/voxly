// * Utility functions for router setup. *

package router

// --- IMPORTS ---
import (
	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/config"
	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/handlers"
)

// --- CODE ---
func newHandlers(configs *config.Configs) *handlersContainer {
	return &handlersContainer{
		system: handlers.NewSystemHandlers(configs),
	}
}
