package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ExpectedHash string `envconfig:"EXPECTED_HASH"`
}

func Load() Config {
	err := godotenv.Load("./env")
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	var config Config
	envconfig.MustProcess("RAJDS", &config)
	return config
}
