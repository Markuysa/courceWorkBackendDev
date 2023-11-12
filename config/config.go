package config

import "github.com/Markuysa/courceWorkBackendDev/pkg/pgconnector"

type Config struct {
	Postgres pgconnector.Config
}
