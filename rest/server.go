package rest

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

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

	// register all the routes
	svr.userHanlder.RegisterRoute(mux, mdlwMngr)

	// server configuration prepare
	addr := ":" + strconv.Itoa(svr.config.Service.HttpPort)
	wrapedMux := mdlwMngr.Wrap(mux)

	server := http.Server{
		Handler: wrapedMux,
		Addr:    addr,
	}

	// initializing the go server
	slog.Info("Starting the server", "PORT", addr)
	if err := server.ListenAndServe(); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
