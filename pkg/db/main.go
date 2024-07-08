package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var Connection *pgxpool.Pool

func InitDB() (*pgxpool.Pool, error) {
	fmt.Println("Connecting to database...")
	pool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		if err != nil {
		}
		os.Exit(1)
	}
	fmt.Println("Connected to database!")
	fmt.Println("Checking database health...")

	pingErr := pool.Ping(context.Background())
	if pingErr != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to ping the database: %v\n", pingErr)
		if err != nil {
		}
		os.Exit(1)
	}
	fmt.Println("Database healthy!")
	return pool, nil
}
