// * Routers setup for the Reclamis service. *
package router

// --- IMPORTS ---
import (
	"github.com/gin-gonic/gin"

	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/config"
)

// --- CODE ---

// Setup initializes and configures the Gin router with routes and handlers.
//
// param configs: Configuration settings for the service.
//
// return: Configured Gin engine.
func Setup(configs *config.Configs) *gin.Engine {

	// Create Gin router
	router := gin.Default()

	// Get handlers
	handlers := newHandlers(configs)

	// System routes
	system := router.Group("/system")
	{
		system.GET("/health", handlers.system.GetHealth)
		system.GET("/info", handlers.system.GetInfo)
	}

	// Return router
	return router
}
