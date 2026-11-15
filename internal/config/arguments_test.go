package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPathFromArgs(t *testing.T) {
	tmpDir := t.TempDir()

	//Создаём валидный конфиг
	validYAML := filepath.Join(tmpDir, "config.yaml")
	if err := os.WriteFile(validYAML, []byte("test: true"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	// Создаём инвалид файл
	invalidYAML := filepath.Join(tmpDir, "invalidext.txt")
	if err := os.WriteFile(invalidYAML, []byte("test: true"), 0644); err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{
			name:    "Valid config path provided",
			args:    []string{"--config", validYAML},
			want:    validYAML,
			wantErr: false,
		},
		{
			name:    "Config flag not specified",
			args:    []string{},
			want:    "",
			wantErr: true,
		},
		{
			name:    "File does not exist",
			args:    []string{"--config", filepath.Join(tmpDir, "nonexistant.yaml")},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Invalid file extension",
			args:    []string{"--config", invalidYAML},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getPathFromArgument(tt.args)

			if (err != nil) != tt.wantErr {
				t.Errorf("getPathFromArgument() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("getPathFromArgument() got = %v, want %v", got, tt.want)
			}
		})
	}
}
