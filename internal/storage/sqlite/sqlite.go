package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/SpectrumDev-tech/todo/internal/storage/migrate"
	"github.com/pressly/goose/v3"

	_ "modernc.org/sqlite"
)

var (
	ErrOpenDatabase = errors.New("failed opening database")
	ErrPingDatabase = errors.New("failed pinging database")
	ErrSetDialect   = errors.New("failed setting dialect")
)

func Init(ctx context.Context, path string, log *slog.Logger) (*sql.DB, error) {
	// Создаём папку (если ещё не создана) под файл db
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("failed creating database directory: %w", err)
	}
	// Init db
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOpenDatabase, err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrPingDatabase, err)
	}

	log.Debug("connected to database", slog.String("path", path))
	return db, nil
}

func Migrate(ctx context.Context, db *sql.DB, log *slog.Logger) error {
	fsys, err := migrationsFS()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrSetDialect, err)
	}
	return migrate.Up(ctx, db, goose.DialectSQLite3, fsys, log)
}
