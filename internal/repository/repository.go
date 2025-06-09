package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Repository struct {
	rdb  *redis.Client
	pgDb *pgxpool.Pool
}

func NewRepository(rdb *redis.Client, pgDb *pgxpool.Pool) *Repository {
	return &Repository{rdb: rdb, pgDb: pgDb}
}
