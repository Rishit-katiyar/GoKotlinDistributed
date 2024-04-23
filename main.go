package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"./utils"
)

var (
	logger *utils.Logger
)

func main() {
	// Initialize logger
	logger = utils.NewLogger("main.log")

	// Create a coordinator with 5 workers
	coordinator := NewCoordinator(5)

	// Start HTTP server
	logger.Info("Starting HTTP server...")
	if err := startHTTPServer(coordinator); err != nil {
		logger.Error("Failed to start HTTP server: ", err)
		os.Exit(1)
	}
}

func startHTTPServer(coordinator *Coordinator) error {
	// Define HTTP server configurations
	server := &http.Server{
		Addr:         ":8080",
		Handler:      createHTTPHandler(coordinator),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start HTTP server
	go func() {
		logger.Info("Server listening on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Error("Failed to start HTTP server: ", err)
		}
	}()

	// Gracefully handle server shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Info("Shutting down server...")

	// Create a deadline for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown error: ", err)
		return err
	}
	logger.Info("Server stopped gracefully")
	return nil
}

func createHTTPHandler(coordinator *Coordinator) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/task", TaskHandler(coordinator))
	mux.HandleFunc("/results", ResultsHandler())
	mux.HandleFunc("/stop", StopHandler(coordinator))
	return mux
}
