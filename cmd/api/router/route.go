package router

import (
	"github.com/IBM/sarama"
	"github.com/jmoiron/sqlx"
	_ "github.com/kongzyeons/go-bank/cmd/api/docs"
	"github.com/redis/go-redis/v9"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/kongzyeons/go-bank/internal/handlers"
	"github.com/kongzyeons/go-bank/internal/queues"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	debitcard_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	account_svc "github.com/kongzyeons/go-bank/internal/services/api/account"
	auth_service "github.com/kongzyeons/go-bank/internal/services/api/auth"
	banner_svc "github.com/kongzyeons/go-bank/internal/services/api/banner"
	debitcard_svc "github.com/kongzyeons/go-bank/internal/services/api/debit-card"
	user_svc "github.com/kongzyeons/go-bank/internal/services/api/user"
)

func InitRouter(
	app *fiber.App,
	redisClient *redis.Client,
	db *sqlx.DB,
	producer sarama.SyncProducer,
) {
	// setup noti
	// lineNoti := line.NewLineAPI()

	// setup queue
	eventProducer := queues.NewEventProducer(producer)

	// setup repository
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)
	bannerRepo := banner_repo.NewBannerRepo(db)
	accountRepo := account_repo.NewAccountRepo(db)
	accountDetailRepo := accountdetail_repo.NewAccountDetailRepo(db)
	accountBalanceRepo := accountbalance_repo.NewaccountBalanceRepo(db)
	debitCardRepo := debitcard_repo.NewDebitCardRepo(db)
	transectionRepo := transaction_repo.NewTransactionRepo(db)

	// setup services
	authSvc := auth_service.NewAuthSvc(
		db, redisClient,
		userRepo, userGreetingRepo,
		transectionRepo,
	)
	userSvc := user_svc.NewUserSvc(userGreetingRepo)
	bannerSvc := banner_svc.NewBannerSvc(bannerRepo)
	accountSvc := account_svc.NewAccountSvc(
		db, eventProducer,
		accountRepo, accountDetailRepo, accountBalanceRepo,
		transectionRepo,
	)
	debitCardSvc := debitcard_svc.NewDebitCardSvc(debitCardRepo)

	// setup handler
	authHandler := handlers.NewAuthHandler(authSvc)
	userHandler := handlers.NewUserHandler(userSvc)
	bannerHandler := handlers.NewBannerHandler(bannerSvc)
	accountHandler := handlers.NewAccountHandler(accountSvc)
	debitCardHandler := handlers.NewDebitCardHandler(debitCardSvc)

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

	// account
	routeAccount := route.Group("/account", middlewareAuth.AuthRequired)
	routeAccount.Post("/getlist", accountHandler.GetList)
	routeAccount.Put("/edit/:accountID", accountHandler.Edit)
	routeAccount.Put("/setIsmain", accountHandler.SetIsmain)
	routeAccount.Get("/getQrcode/:accountID", accountHandler.GetQrcode)
	routeAccount.Put("/addMoney/:accountID", accountHandler.AddMoney)

	// debitCard
	routeDebitCard := route.Group("/debitCard", middlewareAuth.AuthRequired)
	routeDebitCard.Post("/getlist", debitCardHandler.GetList)

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
