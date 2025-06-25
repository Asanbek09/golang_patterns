package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthByID(id int) (*DogOfMonth, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepository(conn *sql.DB) Repository {
	return &postgresRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{
		DB: nil,
	}
}
