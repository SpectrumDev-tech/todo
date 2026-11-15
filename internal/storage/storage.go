package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/SpectrumDev-tech/todo/internal/config"
	"github.com/SpectrumDev-tech/todo/internal/storage/sqlite"
)

// Open database with driver from cfg
// TODO: add support choose driver
func Open(ctx context.Context, cfg config.DatabaseCfg, log *slog.Logger) (*sql.DB, error) {
	var (
		db              *sql.DB
		err             error
		applyMigrations func(context.Context, *sql.DB, *slog.Logger) error
	)

	switch cfg.Driver {
	case "sqlite":
		db, err = sqlite.Init(ctx, cfg.Path, log)
		applyMigrations = sqlite.Migrate

	// TODO: case for PostgreSQL
	default:
		return nil, fmt.Errorf("unknown driver: %s", cfg.Driver)
	}
	if err != nil {
		return nil, err
	}

	if err = applyMigrations(ctx, db, log); err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}
