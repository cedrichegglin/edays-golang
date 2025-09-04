package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cedrichegglin/edays-golang/internal/config"
	"github.com/cedrichegglin/edays-golang/internal/logger"
	"github.com/cedrichegglin/edays-golang/internal/middleware"
	"github.com/cedrichegglin/edays-golang/pkg/greetings"
)

func main() {
	// Initialize configuration
	cfg := config.Load()

	// Initialize logger
	logger := logger.New(cfg.LogLevel)

	// Create HTTP server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      setupRoutes(logger),
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
	}

	// Start server in a goroutine
	go func() {
		logger.Info("Starting server", "port", cfg.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Server failed to start", "error", err)
			os.Exit(1)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Create a deadline for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server forced to shutdown", "error", err)
		os.Exit(1)
	}

	logger.Info("Server exited")
}

func setupRoutes(logger *logger.Logger) http.Handler {
	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// API routes
	mux.HandleFunc("/api/v1/greeting", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		greeting := greetings.Greet()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "%s"}`, greeting)
	})

	// Apply middleware
	handler := middleware.Logging(logger)(mux)
	handler = middleware.Recovery(logger)(handler)
	handler = middleware.CORS()(handler)

	return handler
}
