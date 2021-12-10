package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// Config ...
type Config struct {
	Smtp Smtp
	To []string `json:"tru" required:"true"`
}

type Smtp struct {
	From     string `json:"from" required:"true"`
	Password string `json:"password" required:"true"`
	Host     string `json:"host" required:"true"`
	Port string `json:"port" default:"587"`
}

func Parse() (Config, error) {
	var config Config

	if err := godotenv.Load(); err != nil {
		if !os.IsNotExist(err) {
			return config, err
		}
	}

	if err := envconfig.Process("", &config); err != nil {
		return config, err
	}

	return config, nil
}
