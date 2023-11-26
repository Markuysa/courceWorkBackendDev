package pgconn

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type Connector struct {
	*sqlx.DB
	cfg Config
}

func New(cfg Config) *Connector {
	uri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DB,
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
