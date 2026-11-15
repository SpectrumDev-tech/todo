package storage

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/SpectrumDev-tech/todo/internal/config"
	"github.com/SpectrumDev-tech/todo/internal/storage/sqlite"
)

// Open database with driver from cfg
// TODO: add support choose driver
func Open(cfg config.DatabaseCfg, log *slog.Logger) (*sql.DB, error) {
	switch cfg.Driver {
	case "sqlite":
		db, err := sqlite.Init(cfg.Path, log)
		if err != nil {
			return nil, fmt.Errorf("failed to init sqlite db: %w", err)
		}
		return db, nil

	// TODO: case for PostgreSQL

	default:
		return nil, fmt.Errorf("unknown driver: %s", cfg.Driver)
	}
}
