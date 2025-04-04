package main

import (
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	accountflag_repo "github.com/kongzyeons/go-bank/internal/repositories/account-flags"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	debitcard_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card"
	debitcarddesign_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-design"
	debitcarddetails_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-details"
	debitcardstatus_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-status"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
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
	accountFalgRepo := accountflag_repo.NewAccountFlagRepo(db)
	debitcardRepo := debitcard_repo.NewDebitCardRepo(db)
	debitcardstatuRepo := debitcardstatus_repo.NewDebitCardStatusRepo(db)
	debitCardSDetailRepo := debitcarddetails_repo.NewDebitCardSDetailRepo(db)
	debitCarddesignRepo := debitcarddesign_repo.NewDebitCarddesignRepo(db)
	transectionRepo := transaction_repo.NewTransactionRepo(db)

	taskSvc := user_task.NewTaskSvc(
		db,
		userRepo, userGreetingRepo,
		bannerRepo,
		accountRepo, accountbalanceRepo, accountdetailRepo, accountFalgRepo,
		debitcardRepo, debitcardstatuRepo, debitCardSDetailRepo, debitCarddesignRepo,
		transectionRepo,
	)
	taskSvc.CreateTable()
}
