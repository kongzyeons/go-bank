package router

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/kongzyeons/go-bank/cmd/api/docs"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kongzyeons/go-bank/internal/handlers"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
	banner_svc "github.com/kongzyeons/go-bank/internal/services/api/banner"
	user_svc "github.com/kongzyeons/go-bank/internal/services/api/user"
)

func InitRouter(
	app *fiber.App,
	redisClient *redis.Client,
	db *sqlx.DB,
) {

	// setup repository
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)
	bannerRepo := banner_repo.NewBannerRepo(db)

	// setup services
	authSvc := auth_service.NewAuthSvc(
		db, redisClient,
		userRepo, userGreetingRepo,
	)
	userSvc := user_svc.NewUserSvc(userGreetingRepo)
	bannerSvc := banner_svc.NewBannerSvc(bannerRepo)

	// setup handler
	authHandler := handlers.NewAuthHandler(authSvc)
	userHandler := handlers.NewUserHandler(userSvc)
	bannerHandler := handlers.NewBannerHandler(bannerSvc)

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

	// user
	routeUser := route.Group("/user", middlewareAuth.AuthRequired)
	routeUser.Get("/greeting", userHandler.GetGeeting)

	// banner
	routeBanner := route.Group("/banner", middlewareAuth.AuthRequired)
	routeBanner.Post("/getlist", bannerHandler.GetList)

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
