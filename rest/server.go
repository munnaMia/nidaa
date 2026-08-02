package rest

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/munnaMia/nidaa/internal/config"
	"github.com/munnaMia/nidaa/rest/handler/user"
	"github.com/munnaMia/nidaa/rest/middleware"
)

type Server struct {
	config      *config.Configuration
	userHanlder *user.Handler
}

// create a new rest server
func NewServer(
	cnf *config.Configuration,
	usrHndlr *user.Handler,
) *Server {
	return &Server{
		config:      cnf,
		userHanlder: usrHndlr,
	}
}

// Start the server
func (svr *Server) Start() {
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
