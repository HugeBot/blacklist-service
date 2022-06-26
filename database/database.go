package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ctx = context.Background()
)

type Database struct {
	pool *pgxpool.Pool
}

type Options struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func New(opts *Options) (*Database, error) {
	pool, err := pgxpool.Connect(ctx, fmt.Sprintf("postgresl://%s:%s@%s:%d/%s", opts.Username, opts.Password, opts.Host, opts.Port, opts.Name))
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	database := &Database{
		pool: pool,
	}

	return database, nil
}
