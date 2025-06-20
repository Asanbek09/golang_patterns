package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepository(conn *sql.DB) Repository {
	return &postgresRepository {
		DB: conn,
	}
}