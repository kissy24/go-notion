package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Secret SecretConfig
}

type SecretConfig struct {
	ApiKey     string `toml:"key"`
	DatabaseId string `toml:"id"`
}

func (config Config) ParseToml() SecretConfig {
	_, err := toml.DecodeFile("../config/config.toml", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config.Secret
}
