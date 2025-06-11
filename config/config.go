package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Type string `yaml:"type"`
	}
	Sqlite struct {
		Path string `yaml:"path"`
	}
	Mysql struct {
		Dsn string `yaml:"dsn"`
	}
	Session struct {
		SecretKey string `yaml:"secret_key"`
	}
	Redis struct {
		Addr string `yaml:"addr"`
	}
	Server struct {
		Port string `yaml:"port"`
	}
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
