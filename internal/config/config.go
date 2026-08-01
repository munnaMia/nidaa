package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Service *Service
}

type Service struct {
	Name     string
	Version  string
	HttpPort int
}

var configs *Configuration

// load the env's to the configuration
func loadConfig(path ...string) {
	err := godotenv.Load(path...)
	if err != nil {
		slog.Error(
			"No .env file found falling back to system environment variables",
			"error", err,
		)
		return
	}

	svrName := os.Getenv("SERVICE_NAME")
	if svrName == "" {
		slog.Error("Service name not found")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		slog.Error("Service name not found")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(os.Getenv("HTTP_PORT"), 10, 64)
	if err != nil {
		slog.Error("Could not able to parsing HTTP port address")
		os.Exit(1)
	}

	configs = &Configuration{
		Service: &Service{
			Name:     svrName,
			Version:  version,
			HttpPort: int(httpPort),
		},
	}
}

// take multiple path's as an parameter if the path string is empty then it search
// for the envs on the root .env file return a pointer to configuration env's location
func GetConfig(path ...string) *Configuration {
	if configs == nil {
		loadConfig(path...)
	}
	return configs
}
