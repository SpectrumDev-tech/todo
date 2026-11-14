package config

import (
	"errors"
	"flag"
	"fmt"
)

var (
	ErrValidatePath     = errors.New("invalid path")
	ErrPathNotSpecified = errors.New("path not specified")
)

// Получаем путь к конфигу из аргумента
func getPathFromArgument() (string, error) {
	var path string
	// Парсим аргументы по --config
	flag.StringVar(&path, "config", "", "path to config.yaml")
	flag.Parse()

	// Если удалось что-то получить по аргументу --config - валидируем и возвращаем
	// При возврате из валидатора ошибки - возвращаем ошибку, config будет паниковать сам
	if path != "" {
		if err := validatePathConfig(path); err != nil {
			return "", fmt.Errorf("%w: %s", ErrValidatePath, err)
		}
		return path, nil
	}

	// В остальных случаях возвращаем zero value для path и ошибку
	return "", ErrPathNotSpecified
}
