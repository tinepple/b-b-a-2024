package main

import (
	"backend-bootcamp-assignment-2024/internal/handler"
	"backend-bootcamp-assignment-2024/internal/services/auth_service"
	"backend-bootcamp-assignment-2024/internal/services/kafka_produce_service"
	"backend-bootcamp-assignment-2024/internal/storage"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		_ = db.Close()
	}()

	storageRepo, err := storage.New(db)
	if err != nil {
		log.Fatalf("Failed to create storage: %v", err)
	}

	authService := auth_service.New()
	if err != nil {
		log.Fatalf("Failed to create auth service: %v", err)
	}

	producer, err := sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()
	kafkaService := kafka_produce_service.New(producer)

	logger := logrus.New()
	apiHandler := handler.New(storageRepo, authService, kafkaService, logger)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APISERVER_PORT")), apiHandler); err != nil {
		log.Fatal(err)
	}
}
