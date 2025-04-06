package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/kongzyeons/go-bank/cmd/api/message"
	"github.com/kongzyeons/go-bank/cmd/api/router"
	"github.com/kongzyeons/go-bank/config"
	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/kongzyeons/go-bank/pkg/kafka"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/kongzyeons/go-bank/pkg/redis"
)

// @title API-BANK
// @description This is a sample server for a github.com/kongzyeons/go-bank
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// init config
	cfg := config.InitConfig()

	// inti redis
	redisClient := redis.InitRedis()
	defer redisClient.Close()

	// init db
	db := postgresql.InitPostgresql()
	defer db.Close()

	// init kafka producer
	producer := kafka.InitKafkaProducer()
	defer producer.Close()

	// init kafka consumer
	consumer := kafka.InitKafkaConsumer()
	defer consumer.Close()

	// init message queue
	consumerHandler := message.InitMessage(db)
	log.Println("Account consumer started...")
	go func() {
		for {
			consumer.Consume(context.Background(), events.Topics, consumerHandler)
		}
	}()

	app := router.InitServer()
	router.InitRouter(
		app, redisClient,
		db, producer,
	)
	log.Println("server started...")
	app.Listen(cfg.Port)
}
