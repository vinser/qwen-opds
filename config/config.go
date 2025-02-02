package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Library struct {
		Stock string `yaml:"STOCK"`
	}
	Database struct {
		DSN string `yaml:"DSN"`
	}
	OPDS struct {
		Port int `yaml:"PORT"`
	}
}

func LoadConfig(filepath string) *Config {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	var cfg Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	return &cfg
}

func CreateDefaultConfig() {
	// Создание файла конфигурации
}
