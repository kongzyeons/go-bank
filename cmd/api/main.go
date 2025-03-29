package main

import (
	"github.com/kongzyeons/go-bank/cmd/api/router"
	"github.com/kongzyeons/go-bank/config"
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

	cfg := config.InitConfig()

	app := router.NewServer()
	router.NewRouter(app)
	app.Listen(cfg.Port)
}
