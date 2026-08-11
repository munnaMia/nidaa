package cmd

import (
	"context"
	"log/slog"

	"github.com/munnaMia/nidaa/internal/config"
	"github.com/munnaMia/nidaa/internal/infra/auth"
	"github.com/munnaMia/nidaa/internal/infra/postgres"
	"github.com/munnaMia/nidaa/internal/usecase"
	"github.com/munnaMia/nidaa/rest"
	"github.com/munnaMia/nidaa/rest/handler/user"
	jsonhelper "github.com/munnaMia/nidaa/util/jsonHelper"
	"github.com/munnaMia/nidaa/util/logger"
	"github.com/munnaMia/nidaa/util/responder"
	"github.com/munnaMia/nidaa/util/validate"
)

// Start the nidaa application
func Run() {
	ctx := context.Background()

	// setup a default logger.
	lg := logger.NewLogger(false, false)
	slog.SetDefault(lg)

	// geting configuration from the env
	cnf := config.GetConfig()

	// initializing a sevices
	httpResponder := responder.NewHttpResponder()
	jwtService := auth.NewJWTService(cnf.Service.SecretKey)
	jsonHelper := jsonhelper.NewJsonHelper()
	validate := validate.NewValidate()

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
	userUsecase := usecase.NewUserUseCase(
		userRepo,
		jwtService,
	)

	// creating handlers
	userHandler := user.NewHandler(
		userUsecase,
		jsonHelper,
		httpResponder,
		validate,
	)

	// create a new rest server
	svr := rest.NewServer(
		cnf,
		jwtService,
		userHandler,
	)

	// starting the server
	svr.Start()
}
