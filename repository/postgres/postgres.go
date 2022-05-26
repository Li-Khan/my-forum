package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/Li-Khan/my-forum/config"
	"github.com/Li-Khan/my-forum/lib/mistake"
	"github.com/jackc/pgx/v4/pgxpool"
)

// * NewPostgresRepository - Connect to the database
func NewPostgresRepository(c *config.Config) (db *pgxpool.Pool, err error) {
	defer func() {
		err = mistake.WrapIfErr("postgres", err)
	}()

	dsn := fmt.Sprintf("postgres://%s:%s@%s%s/%s", c.UserDB, c.PasswordDB, c.HostDB, c.PortDB, c.NameDB)

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 25
	cfg.MaxConnLifetime = 5 * time.Minute

	db, err = pgxpool.ConnectConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}
