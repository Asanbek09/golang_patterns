package main

import (
	"example/models"
	"log"
	"os"
	"testing"
)

var testApp application

func TestMain(m *testing.M) {
	dsn := "postgres://postgres:secretPassword@localhost:5432/breeders?sslmode=disable"
	db, err := initPostgresDB(dsn)

	if err != nil {
		log.Panic(err)
	}

	testApp = application{
		DB: db,
		Models: *models.New(db),
	}

	os.Exit(m.Run())
}