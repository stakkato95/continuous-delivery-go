package config

import "github.com/stakkato95/service-engineering-go-lib/config"

type Config struct {
	PostgresService string `mapstructure:"POSTGRES_SERVICE"`
}

var AppConfig Config

func init() {
	config.Init(&AppConfig, Config{})
}
