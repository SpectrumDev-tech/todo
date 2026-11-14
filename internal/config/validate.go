package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// Проверяем есть ли вообще такой путь.
// Является ли конечный файл - yaml файлом.
func validatePathConfig(path string) error {
	cleanPath := filepath.Clean(path)

	info, err := os.Stat(cleanPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file does not exists %s: %q", cleanPath, err)
		}
		return fmt.Errorf("any error: %q", err)
	}

	if info.IsDir() {
		return fmt.Errorf("%s is a directory", cleanPath)
	}

	ext := filepath.Ext(cleanPath)
	if ext != ".yaml" {
		return fmt.Errorf("%s is not a YAML file", cleanPath)
	}
	return nil
}
