package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type RepositoryConfig struct {
	db *sqlx.DB
}

func (repository *RepositoryConfig) Repository() *RepositoryConfig {
	db, err := sqlx.Connect("postgres", "host=localhost user=root password=root dbname=adopler sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	repo := &RepositoryConfig{db}
	return repo
}
