package main

import (
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	user_task "github.com/kongzyeons/go-bank/internal/services/task"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func main() {
	db := postgresql.InitPostgresql()
	defer db.Close()
	userRepo := user_repo.NewUserRepo(db)
	userGreetingRepo := usergreeting_repo.NewUserGreetingRepo(db)
	bannerRepo := banner_repo.NewBannerRepo(db)
	taskSvc := user_task.NewTaskSvc(
		db, userRepo, userGreetingRepo, bannerRepo,
	)
	taskSvc.CreateTable()
}
