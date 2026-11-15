package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestServerCfg_Addr(t *testing.T) {
	tests := []struct {
		name   string
		server ServerCfg
		want   string
	}{
		{
			name: "Default port and address",
			server: ServerCfg{
				Port: 8080,
				Host: "localhost",
			},
			want: "localhost:8080",
		},
		{
			name: "Custom addr",
			server: ServerCfg{
				Port: 7598,
				Host: "255.0.4.6",
			},
			want: "255.0.4.6:7598",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.server.Addr(); got != tt.want {
				t.Errorf("ServerCfg.Addr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMustLoadConfig_Success(t *testing.T) {
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "local.yaml")
	yamlContent := []byte(`server:
  port: 8080
  host: "localhost"
database:
  driver: "sqlite"
  path: "./data/test.db"
`)

	if err := os.WriteFile(configPath, yamlContent, 0644); err != nil {
		t.Fatalf("failed to write temp config file: %v", err)
	}
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--config", configPath}

	var cfg *Config = MustLoadConfig()

	if cfg == nil {
		t.Fatalf("expected config to be loaded, got nil")
	}
	if cfg.Server.Port != 8080 {
		t.Errorf("expected port 8080, got %d", cfg.Server.Port)
	}
	if cfg.Server.Host != "localhost" {
		t.Errorf("expected host 'localhost', got %s", cfg.Server.Host)
	}
}

func TestMustLoadConfig_PanicOnMissionFile(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	os.Args = []string{"cmd", "--config", "nonexistent.yaml"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic")
		}
	}()

	MustLoadConfig()
}
