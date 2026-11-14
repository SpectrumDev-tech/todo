package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/SpectrumDev-tech/todo/internal/config"
	"github.com/SpectrumDev-tech/todo/internal/logger"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	cfg := config.MustLoadConfig()
	log := logger.SetupLogger(slog.LevelDebug, os.Stdout)

	log.Info("starting server",
		slog.String("addr", cfg.Server.Addr()),
		slog.Any("database", cfg.Database),
	)

	http.HandleFunc("/status", helloHandler)

	if err := http.ListenAndServe(cfg.Server.Addr(), nil); err != nil {
		return err
	}
	return nil
}
