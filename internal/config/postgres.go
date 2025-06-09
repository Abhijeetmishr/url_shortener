package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// make sure your function start with uppercase to call outside of the directory.
func ConnectDatabase() *pgxpool.Pool {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	fmt.Println("envs: ", host, port, user, password, dbName)
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=5&pool_min_conns=1", user, password, host, port, dbName)

	dbpool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Panicf("unable to connect to database: %v\n", err)
	}

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Panicf("Postgres QueryRow failed:: %v\n", err)
	}

	fmt.Println("Successfully connected to postgres database")
	return dbpool
}
