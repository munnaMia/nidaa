package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Service  *Service
	Database *DB
}

type Service struct {
	Name      string
	Version   string
	HttpPort  int
	SecretKey string
}

type DB struct {
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
	SSL_Mode    string
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

	// fetching service configuration
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

	scrkey := os.Getenv("SECRET_KEY")
	if scrkey == "" {
		slog.Error("secret key string not found")
		os.Exit(1)
	}

	// fetching db configurations
	dbHost := os.Getenv("DB_HOST")
	if svrName == "" {
		slog.Error("db host string not found")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if svrName == "" {
		slog.Error("db port  not found")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if svrName == "" {
		slog.Error("db user not found")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if svrName == "" {
		slog.Error("db password not found")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if svrName == "" {
		slog.Error("db name not found")
		os.Exit(1)
	}

	sslMode := os.Getenv("SSL_MODE")
	if svrName == "" {
		slog.Error("sslmode not found")
		os.Exit(1)
	}

	configs = &Configuration{
		Service: &Service{
			Name:      svrName,
			Version:   version,
			HttpPort:  int(httpPort),
			SecretKey: scrkey,
		},
		Database: &DB{
			DB_Host:     dbHost,
			DB_Port:     dbPort,
			DB_User:     dbUser,
			DB_Password: dbPassword,
			DB_Name:     dbName,
			SSL_Mode:    sslMode,
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
