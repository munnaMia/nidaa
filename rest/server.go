package rest

import (
	"log/slog"
	"net/http"
	"os"
)

// Start the server
func Start() {
	mux := http.NewServeMux()

	// prepare the middleware & manager
	// register the middleware and manager
	// register routes
	// run the server

	server := http.Server{
		Handler: mux,
		Addr:    "8080",
	}

	// initializing the go server
	slog.Info("Starting the server", "PORT", "8080")
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
