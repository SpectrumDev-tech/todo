package main

import (
	"fmt"
	"net/http"

	"github.com/SpectrumDev-tech/todo/internal/config"
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
	fmt.Printf("Server Port: %d\n", cfg.Server.Port)

	http.HandleFunc("/status", helloHandler)

	if err := http.ListenAndServe(cfg.Server.Addr(), nil); err != nil {
		return err
	}
	return nil
}
