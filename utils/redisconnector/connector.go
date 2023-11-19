package redisconnector

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Connector struct {
	*redis.Client
	cfg Config
}

func New(cfg Config) *Connector {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err)
	}

	return &Connector{
		Client: client,
		cfg:    cfg,
	}
}
