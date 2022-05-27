package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Li-Khan/my-forum/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

// * NewPostgresRepository - Connect to the database
func NewPostgresRepository(c *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s%s/%s", c.UserDB, c.PasswordDB, c.HostDB, c.PortDB, c.NameDB)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	cfg.MaxConns = 25
	cfg.MaxConnLifetime = 5 * time.Minute

	db, err := pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	return db, nil
}
