package main

import (
	"backend-bootcamp-assignment-2024/internal/services/kafka_consume_service"
	"backend-bootcamp-assignment-2024/internal/services/sender_service"
	"backend-bootcamp-assignment-2024/internal/storage"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", os.Getenv("KAFKA_HOST"), os.Getenv("KAFKA_PORT"))}, nil)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	partConsumer, err := consumer.ConsumePartition("ping", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer partConsumer.Close()

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

	senderService := sender_service.New()
	logger := logrus.New()
	kafkaService := kafka_consume_service.New(partConsumer, senderService, storageRepo, logger)

	if err = kafkaService.Consume(ctx); err != nil {
		log.Fatal(err)
	}
}
