package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func DbConnector() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}
	defer dbpool.Close()
}
