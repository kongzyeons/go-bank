package main

import (
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
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
	accountRepo := account_repo.NewAccountRepo(db)
	accountbalanceRepo := accountbalance_repo.NewaccountBalanceRepo(db)
	accountdetailRepo := accountdetail_repo.NewAccountDetailRepo(db)
	taskSvc := user_task.NewTaskSvc(
		db,
		userRepo, userGreetingRepo,
		bannerRepo,
		accountRepo, accountbalanceRepo, accountdetailRepo,
	)
	taskSvc.CreateTable()
}
