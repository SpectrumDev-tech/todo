package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestValidatePathConfig(t *testing.T) {
	tmpDir := t.TempDir()

	// Создаём валидный конфиг файл
	validYAML := filepath.Join(tmpDir, "config.yaml")
	if err := os.WriteFile(validYAML, []byte("test: true"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Файл с инвалидным расширением
	invalidYAML := filepath.Join(tmpDir, "config.txt")
	if err := os.WriteFile(invalidYAML, []byte("test: true"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Тестируем вложенность
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Valid YAML file",
			path:    validYAML,
			wantErr: false,
		},
		{
			name:    "Invalid YAML file",
			path:    invalidYAML,
			wantErr: true,
		},
		{
			name:    "File does not exist",
			path:    filepath.Join(tmpDir, "nonextconfig.yaml"),
			wantErr: true,
		},
		{
			name:    "Path is a directory",
			path:    subDir,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePathConfig(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePathConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
