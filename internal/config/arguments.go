package config

import (
	"errors"
	"flag"
	"fmt"
)

var (
	ErrValidatePath     = errors.New("invalid path")
	ErrPathNotSpecified = errors.New("path not specified")
	ErrFailedParse      = errors.New("failed to parse arguments")
)

// Получаем путь к конфигу из аргумента
func getPathFromArgument(args []string) (string, error) {
	//Создаём изолированный флаг, чтобы не трогать глобальный
	fs := flag.NewFlagSet("config", flag.ContinueOnError)

	var path string
	// Парсим аргументы по --config
	fs.StringVar(&path, "config", "", "path to config.yaml")
	if err := fs.Parse(args); err != nil {
		return "", fmt.Errorf("%w: %v", ErrFailedParse, err)
	}
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
