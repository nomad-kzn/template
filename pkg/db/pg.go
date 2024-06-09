package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq" // postgres
)

const driver = "postgres"

type PostgresClient interface {
	Connect(ctx context.Context) error
	Database() *sql.DB
	IsConnected() bool
	Close() error
}

type client struct {
	cfg       *DatabaseConfig
	db        *sql.DB
	connected bool
}

func NewPostgresClient(cfg *DatabaseConfig) PostgresClient {
	c := &client{
		cfg: cfg,
	}
	return c
}

func (c *client) IsConnected() bool {
	return c.connected
}

func (c *client) Connect(_ context.Context) error {
	if c.cfg == nil {
		return errors.New("db Connect err: config required")
	}

	db, err := sql.Open(driver, c.cfg.URI())
	if err != nil {
		return fmt.Errorf("database connect err: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("database ping err: %w", err)
	}

	c.db = db
	c.connected = true

	return nil
}

func (c *client) Database() *sql.DB {
	return c.db
}

func (c *client) Close() error {
	return c.db.Close()
}
