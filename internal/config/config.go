package config

import (
	"net"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Server   ServerCfg   `yaml:"server"`
	Database DatabaseCfg `yaml:"database"`
}

type ServerCfg struct {
	Port int    `yaml:"port" env-required:"true"`
	Host string `yaml:"host" env-required:"true"`
}

func (s *ServerCfg) Addr() string {
	return net.JoinHostPort(s.Host, strconv.Itoa(s.Port))
}

type DatabaseCfg struct {
	Driver string `yaml:"driver" env-default:"sqlite"`
	Path   string `yaml:"path" env-required:"true"`
}

func MustLoadConfig() *Config {
	var cfg Config

	path, err := getPathFromArgument()
	if err != nil {
		panic(err)
	}

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(err)
	}
	return &cfg
}
