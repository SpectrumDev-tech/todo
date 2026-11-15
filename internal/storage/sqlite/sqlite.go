package sqlite

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pressly/goose/v3"

	_ "modernc.org/sqlite"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var (
	ErrOpenDatabase = errors.New("failed opening database")
	ErrPingDatabase = errors.New("failed pinging database")
	ErrSetDialect   = errors.New("failed setting dialect")
)

func Init(path string, log *slog.Logger) (*sql.DB, error) {
	// Создаём папку (если ещё не создана) под файл db
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("failed creating database directory: %w", err)
	}
	// Init db
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrOpenDatabase, err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrPingDatabase, err)
	}

	log.Debug("connected to database", slog.String("path", path))

	// Complete migrations
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrSetDialect, err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %v", err)
	}

	return db, nil
}
