package database

import (
	"context"
	"database/sql"
	"fmt"
	"logger-test/logger" 
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() Service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, database)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err)
		os.Exit(1)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		logger.Error("Database is down", "error", err)
		os.Exit(1)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
