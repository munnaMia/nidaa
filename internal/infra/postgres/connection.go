package postgres

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/munnaMia/nidaa/internal/config"
)

// construct a dsn string for postgres
func getDSN(cnf *config.Configuration) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s&application_name=%s",
		cnf.Database.DB_User,
		cnf.Database.DB_Password,
		cnf.Database.DB_Host,
		cnf.Database.DB_Port,
		cnf.Database.DB_Name,
		cnf.Database.SSL_Mode,
		cnf.Service.Name,
	)
}

func NewConnection(ctx context.Context, cnf *config.Configuration) (*sql.DB, error) {
	dsn := getDSN(cnf)

	pool, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.PingContext(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping the database, %w", err)
	}

	pool.SetMaxIdleConns(6)
	pool.SetMaxOpenConns(10)

	return pool, err
}
