package router

import (
	_ "github.com/kongzyeons/go-bank/cmd/api/docs"
	"github.com/kongzyeons/go-bank/cmd/api/router/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kongzyeons/go-bank/internal/handlers"
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
)

func NewRouter(app *fiber.App) {
	// setup repository

	// setup services
	authSvc := auth_service.NewAuthSvc()

	// setup handler
	authHandler := handlers.NewAuthHandler(authSvc)

	// setup route
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	route := app.Group("/api/v1")
	route.Post("/register", authHandler.Register)
	route.Post("/login", authHandler.Register)

	routeAuth := route.Group("/auth", middleware.GetProtected)
	routeAuth.Get("/ping", authHandler.Ping)
	routeAuth.Post("/logout", authHandler.Logout)

}

func NewServer() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	// rateLimiter := limiter.New(limiter.Config{
	// 	Max:        20, // default 20
	// 	Expiration: 1,  //default 1 minute
	// 	LimitReached: func(c *fiber.Ctx) error {
	// 		return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests, please try again later.")
	// 	},
	// })
	// // Apply the rate limiter middleware
	// app.Use(rateLimiter)
	return app

}
