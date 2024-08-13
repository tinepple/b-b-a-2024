package main

import (
	"log"
	"net/http"

	"backend-bootcamp-assignment-2024/internal/handler"
	"backend-bootcamp-assignment-2024/internal/storage"

	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = db.Close()
	}()

	storageRepo, err := storage.New(db)
	if err != nil {
		log.Fatal(err)
	}

	apiHandler := handler.New(storageRepo)

	if err := http.ListenAndServe(":3333", apiHandler); err != nil {
		log.Fatal(err)
	}
}
