package cmd

import (
	"log/slog"

	"github.com/munnaMia/nidaa/rest"
	"github.com/munnaMia/nidaa/util/logger"
)

// Start the nidaa application
func Run() {

	// prepare everything then run the server

	// setup a default logger.
	lg := logger.NewLogger(false, false)
	slog.SetDefault(lg)

	// starting the server
	rest.Start()
}
