package rest

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/munnaMia/nidaa/rest/middleware"
)

// Start the server
func Start() {
	mux := http.NewServeMux()

	// prepare the middleware & manager
	mdlw := middleware.NewMiddleware()
	mdlwMngr := middleware.NewManager()

	// register global middlewares
	mdlwMngr.GlobalMiddleware(
		mdlw.Logger,
	)

	// register routes...?

	server := http.Server{
		Handler: mdlwMngr.Wrap(mux),
		Addr:    "8080",
	}

	// initializing the go server
	slog.Info("Starting the server", "PORT", "8080")
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
