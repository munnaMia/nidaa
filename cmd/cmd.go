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
)

// Start the nidaa application
func Run() {
	ctx := context.Background()

	// setup a default logger.
	lg := logger.NewLogger(false, false)
	slog.SetDefault(lg)

	// geting configuration from the env
	cnf := config.GetConfig()

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
	userHandler := user.NewHandler(userUsecase)

	// create a new rest server
	svr := rest.NewServer(
		cnf,
		userHandler,
	)

	// starting the server
	svr.Start()
}
