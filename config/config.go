package config

import (
	"encoding/json"
	"os"

	"github.com/Markuysa/courceWorkBackendDev/utils/duration"
	"github.com/Markuysa/courceWorkBackendDev/utils/oteltrace"
	"github.com/Markuysa/courceWorkBackendDev/utils/pgconn"
	"github.com/Markuysa/courceWorkBackendDev/utils/redisconnector"
)

type Config struct {
	Trace    oteltrace.Config
	Postgres pgconn.Config
	Redis    redisconnector.Config
	HTTP     struct {
		URI string
	}
	ServiceName      string
	AdminServiceName string
	Auth             struct {
		AccessTTL  duration.Duration
		SessionTTL duration.Duration
	}
	Telegram struct {
		Token string
	}
}

func LoadConfig() (cfg Config, err error) {
	jsonFile, err := os.Open(os.Getenv("CONFIG"))
	if err != nil {
		return cfg, err
	}

	err = json.NewDecoder(jsonFile).Decode(&cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, err
}
