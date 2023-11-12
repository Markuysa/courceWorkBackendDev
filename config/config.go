package config

import "github.com/Markuysa/courceWorkBackendDev/pkg/pgconnector"

type Config struct {
	Postgres pgconnector.Config
}

func LoadConfig() (cfg Config, err error) {

	return cfg, err
}
