package cmd

import (
	"context"
	"log/slog"

	"github.com/munnaMia/nidaa/internal/config"
	"github.com/munnaMia/nidaa/internal/infra/postgres"
	"github.com/munnaMia/nidaa/internal/usecase"
	"github.com/munnaMia/nidaa/rest"
	"github.com/munnaMia/nidaa/rest/handler/user"
	"github.com/munnaMia/nidaa/util/logger"
	"github.com/munnaMia/nidaa/util/responder"
)

// Start the nidaa application
func Run() {
	ctx := context.Background()

	// setup a default logger.
	lg := logger.NewLogger(false, false)
	slog.SetDefault(lg)

	// geting configuration from the env
	cnf := config.GetConfig()

	// initializing a http responder
	httpResponder := responder.NewHttpResponder()

	// initializing db connection
	pool, err := postgres.NewConnection(ctx, cnf)
	if err != nil {
		slog.Error("Error while initializing db connection", "err", err)
		return
	}
	defer pool.Close()

	// initialized repositories
	userRepo := postgres.NewUserRepository(pool)

	// initialized usecases
	userUsecase := usecase.NewUserUseCase(userRepo)

	// creating handlers
	userHandler := user.NewHandler(userUsecase, httpResponder)

	// create a new rest server
	svr := rest.NewServer(
		cnf,
		userHandler,
	)

	// starting the server
	svr.Start()
}
