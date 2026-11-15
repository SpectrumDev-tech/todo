package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	db, err := storage.Open(ctx, cfg.Database, log)
	if err != nil {
		return fmt.Errorf("init storage: %w", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Error("close db connection", slog.Any("error", err))
		}
	}()

	router := http.NewServeMux()
	router.HandleFunc("/status", helloHandler)
	srv := &http.Server{
		Addr:    cfg.Server.Addr(),
		Handler: router,
	}

	serverError := make(chan error, 1)

	go func() {
		log.Info("starting server",
			slog.String("addr", cfg.Server.Addr()),
		)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverError <- err
			log.Error("HTTP server error", slog.Any("error", err))
		}
	}()

	select {
	case err := <-serverError:
		return fmt.Errorf("server failed: %w", err)
	case <-ctx.Done():
		log.Info("shutdown signal")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown: %w", err)
	}
	log.Info("server shutdown complete")
	return nil
}
