package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	ctx  = context.Background()
	pool *pgxpool.Pool
)

type Options struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func (o *Options) init() {
	if o.Host == "" {
		o.Host = "localhost"
	}

	if o.Port == 0 {
		o.Port = 5432
	}

	if o.Name == "" {
		log.Fatal("database > name | must be defined")
	}

	if o.Username == "" {
		log.Fatal("database > username | must be defined")
	}

	if o.Password == "" {
		log.Fatal("database > password | must be defined")
	}
}

func Init(opts *Options) error {

	opts.init()

	newPool, err := pgxpool.Connect(ctx, fmt.Sprintf("postgresl://%s:%s@%s:%d/%s", opts.Username, opts.Password, opts.Host, opts.Port, opts.Name))
	if err != nil {
		return err
	}

	if err := pool.Ping(ctx); err != nil {
		return err
	}

	pool = newPool

	return nil
}

func Query(sql string, args ...interface{}) (pgx.Rows, error) {
	return pool.Query(ctx, sql, args...)
}

func QueryRow(sql string, args ...interface{}) pgx.Row {
	return pool.QueryRow(ctx, sql, args...)
}
