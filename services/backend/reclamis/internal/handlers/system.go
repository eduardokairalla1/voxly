// * Handlers for system-related endpoints. *

package handlers

// --- IMPORTS ---
import (
	"github.com/gin-gonic/gin"

	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/config"
	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/models"
)

// --- CONSTRUCTORS ---
func NewSystemHandlers(configs *config.Configs) *SystemHandlers {
	return &SystemHandlers{
		configs: configs,
	}
}

// --- CODE ---

// GetHealth handles the /health endpoint.
func (h *SystemHandlers) GetHealth(context *gin.Context) {
	context.JSON(200, models.Health{
		Status: "OK",
	})
}

// GetInfo handles the /info endpoint.
func (h *SystemHandlers) GetInfo(context *gin.Context) {
	context.JSON(200, models.Info{
		Name:        "Reclamis",
		Description: "This is Voxly's REST API service.",
		Version:     h.configs.Version,
	})
}
