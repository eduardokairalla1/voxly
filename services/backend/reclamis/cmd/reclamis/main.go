// * Service entry point *

package main

// --- IMPORTS ---
import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/config"
	"github.com/eduardokairalla1/voxly/backend/reclamis/internal/router"
)

// --- CODE ---
func main() {

	// Load configuration
	configs := config.Load()

	// Setup router
	routers := router.Setup(configs)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + configs.Port,
		Handler: routers,
	}

	// Start server in a goroutine
	go func() {

		// log server start
		log.Printf("Server starting on port %s", configs.Port)

		// Start listening and serving HTTP requests
		err := srv.ListenAndServe()

		// error occurred while starting or running the server: log and exit
		if err != nil && err != http.ErrServerClosed {
			log.Printf("Server failed: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)
	<-quit

	// log server shutdown
	log.Println("Shutting down server...")

	// Create context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt graceful server shutdown
	_ = srv.Shutdown(ctx)

	// log server exit
	log.Println("Server exited")
}
