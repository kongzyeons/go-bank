package router

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/kongzyeons/go-bank/cmd/api/docs"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kongzyeons/go-bank/internal/handlers"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
	homepage_svc "github.com/kongzyeons/go-bank/internal/services/api/homepage"
)

func InitRouter(
	app *fiber.App,
	redisClient *redis.Client,
	db *sqlx.DB,
) {

	// setup repository
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)

	// setup services
	authSvc := auth_service.NewAuthSvc(
		db, redisClient,
		userRepo, userGreetingRepo,
	)
	homePage := homepage_svc.NewHomePageSvc(userGreetingRepo)

	// setup handler
	authHandler := handlers.NewAuthHandler(authSvc)
	homePageHandler := handlers.NewHomePageHandler(homePage)

	// setup route
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	route := app.Group("/api/v1")
	route.Post("/register", authHandler.Register)
	route.Post("/login", authHandler.Login)

	middlewareAuth := NewMiddlewareAuth(redisClient)
	// auth
	routeAuth := route.Group("/auth", middlewareAuth.AuthRequired)
	routeAuth.Get("/ping", authHandler.Ping)
	routeAuth.Post("/refresh", authHandler.Refresh)
	routeAuth.Post("/logout", authHandler.Logout)

	// homepage
	routeHomePage := route.Group("/homepage", middlewareAuth.AuthRequired)
	routeHomePage.Get("/greeting", homePageHandler.GetUserGreetings)

}

func InitServer() *fiber.App {
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
