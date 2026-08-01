package cmd

import (
	"log/slog"

	"github.com/munnaMia/nidaa/internal/config"
	"github.com/munnaMia/nidaa/rest"
	"github.com/munnaMia/nidaa/util/logger"
)

// Start the nidaa application
func Run() {

	// prepare everything then run the server
	cnf := config.GetConfig()

	// setup a default logger.
	lg := logger.NewLogger(false, false)
	slog.SetDefault(lg)

	// create a new rest sever and pass the dependency ok .

	// starting the server
	rest.Start()
}
