package server

import (
	"fmt"
	"github/Ndraaa15/fitness-tracker-api/cmd/config"
	"log"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
)

type server struct {
	server *http.Server
}

func newServer() (*server, error) {
	srv := &server{
		server: &http.Server{
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	_, err := sqlx.Connect("postgres", config.DBConfig())

	if err != nil {
		log.Printf("Error connecting to database: %v", err)
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	return srv, nil
}
