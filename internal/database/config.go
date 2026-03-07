package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() *pgxpool.Pool {

	databaseUrl := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.New(context.Background(), databaseUrl)

	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	return pool
}