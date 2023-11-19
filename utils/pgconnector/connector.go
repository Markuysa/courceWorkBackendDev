package pgconnector

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type Connector struct {
	*sqlx.DB
	cfg Config
}

func New(cfg Config) *Connector {
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
	)

	conn, err := sqlx.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}

	if err = conn.Ping(); err != nil {
		log.Fatal(err)
	}

	return &Connector{
		DB:  conn,
		cfg: cfg,
	}
}
