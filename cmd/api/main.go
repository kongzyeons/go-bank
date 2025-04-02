package main

import (
	"github.com/kongzyeons/go-bank/cmd/api/router"
	"github.com/kongzyeons/go-bank/config"
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
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// inti redis
	redisClient := redis.InitRedis()
	defer redisClient.Close()

	// init db
	db := postgresql.InitPostgresql()
	defer db.Close()

	cfg := config.InitConfig()

	app := router.InitServer()
	router.InitRouter(app, redisClient, db)
	app.Listen(cfg.Port)
}
