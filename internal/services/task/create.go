package task_svc

import (
	"log"

	"github.com/jmoiron/sqlx"
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
)

type TaskCreateSvc interface {
	CreateTable() error
}

type taskCreateSvc struct {
	db                   *sqlx.DB
	userRepo             user_repo.UserRepo
	userGreetingRepo     usergreeting_repo.UserGreetingRepo
	bannerRepo           banner_repo.BannerRepo
	accountRepo          account_repo.AccountRepo
	accountBalanceRepo   accountbalance_repo.AccountBalanceRepo
	accountDetailRepo    accountdetail_repo.AccountDetailRepo
	accountFalgRepo      accountflag_repo.AccountFlagRepo
	debitcardRepo        debitcard_repo.DebitCardRepo
	debitcardstatuRepo   debitcardstatus_repo.DebitCardStatusRepo
	debitCardSDetailRepo debitcarddetails_repo.DebitCardSDetailRepo
	debitCarddesignRepo  debitcarddesign_repo.DebitCarddesignRepo
	transectionRepo      transaction_repo.TransactionRepo
}

func NewTaskCreateSvc(
	db *sqlx.DB,
	userRepo user_repo.UserRepo,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
	bannerRepo banner_repo.BannerRepo,
	accountRepo account_repo.AccountRepo,
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo,
	accountDetail accountdetail_repo.AccountDetailRepo,
	accountFalgRepo accountflag_repo.AccountFlagRepo,
	debitcardRepo debitcard_repo.DebitCardRepo,
	debitcardstatuRepo debitcardstatus_repo.DebitCardStatusRepo,
	debitCardSDetailRepo debitcarddetails_repo.DebitCardSDetailRepo,
	debitCarddesignRepo debitcarddesign_repo.DebitCarddesignRepo,
	transectionRepo transaction_repo.TransactionRepo,
) TaskCreateSvc {
	return &taskCreateSvc{
		db:                   db,
		userRepo:             userRepo,
		userGreetingRepo:     userGreetingRepo,
		bannerRepo:           bannerRepo,
		accountRepo:          accountRepo,
		accountBalanceRepo:   accountBalanceRepo,
		accountDetailRepo:    accountDetail,
		accountFalgRepo:      accountFalgRepo,
		debitcardRepo:        debitcardRepo,
		debitcardstatuRepo:   debitcardstatuRepo,
		debitCardSDetailRepo: debitCardSDetailRepo,
		debitCarddesignRepo:  debitCarddesignRepo,
		transectionRepo:      transectionRepo,
	}
}

func (svc *taskCreateSvc) CreateTable() error {
	// create user table
	err := svc.createTableUser()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'users' created successfully!")

	// create banner table
	err = svc.createTableBanner()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'banners' created successfully!")

	// create account table
	err = svc.createTableAccount()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'account' created successfully!")

	// create debit card table
	err = svc.createTableDebitCard()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'debit card' created successfully!")

	// create transection table
	err = svc.createTableTransection()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Table 'transection' created successfully!")
	return nil
}

func (svc *taskCreateSvc) createTableUser() error {
	err := svc.userRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.userGreetingRepo.CreateTable()
	return err
}

func (svc *taskCreateSvc) createTableBanner() error {
	err := svc.bannerRepo.CreateTable()
	return err
}

func (svc *taskCreateSvc) createTableAccount() error {
	err := svc.accountRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.accountBalanceRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.accountDetailRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.accountFalgRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.accountRepo.CreateTableView()
	return err
}

func (svc *taskCreateSvc) createTableDebitCard() error {
	err := svc.debitcardRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.debitcardstatuRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.debitCardSDetailRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.debitCarddesignRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.debitcardRepo.CreateTableView()
	if err != nil {
		return err
	}
	return nil
}

func (svc *taskCreateSvc) createTableTransection() error {
	err := svc.transectionRepo.CreateTable()
	if err != nil {
		return err
	}
	return nil
}
