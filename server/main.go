package server

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"nexus/internal/router"
	"os"
	"os/signal"
	"time"
)

func StartServer() error {
	r := router.SetupRoutes()
	// Server configuration
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server
	go func() {
		log.Printf("Server is starting on port %s\n", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
	return nil
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello World"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Custom middleware example
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request received: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
