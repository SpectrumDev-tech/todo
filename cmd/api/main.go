package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/SpectrumDev-tech/todo/internal/config"
	"github.com/SpectrumDev-tech/todo/internal/logger"
	"github.com/SpectrumDev-tech/todo/internal/storage"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
}

func main() {
	cfg := config.MustLoadConfig()
	log := logger.SetupLogger(slog.LevelDebug, os.Stdout)

	if err := run(cfg, log); err != nil {
		log.Error("applications stopped with error", slog.Any("error", err))
		os.Exit(1)
	}
}

func run(cfg *config.Config, log *slog.Logger) error {
	db, err := storage.Open(cfg.Database, log)
	if err != nil {
		return fmt.Errorf("init storage: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error("close db connection", slog.Any("error", err))
		}
	}()

	log.Info("starting server",
		slog.String("addr", cfg.Server.Addr()),
	)

	http.HandleFunc("/status", helloHandler)

	if err := http.ListenAndServe(cfg.Server.Addr(), nil); err != nil {
		return err
	}
	return nil
}
