package user_task

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type TaskSvc interface {
	CreateTable()
	MockData()
}

type taskSvc struct {
	db                 *sqlx.DB
	userRepo           user_repo.UserRepo
	userGreetingRepo   usergreeting_repo.UserGreetingRepo
	bannerRepo         banner_repo.BannerRepo
	accountRepo        account_repo.AccountRepo
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo
	accountDetailRepo  accountdetail_repo.AccountDetailRepo
}

func NewTaskSvc(
	db *sqlx.DB,
	userRepo user_repo.UserRepo,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
	bannerRepo banner_repo.BannerRepo,
	accountRepo account_repo.AccountRepo,
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo,
	accountDetail accountdetail_repo.AccountDetailRepo,
) TaskSvc {
	return &taskSvc{
		db:                 db,
		userRepo:           userRepo,
		userGreetingRepo:   userGreetingRepo,
		bannerRepo:         bannerRepo,
		accountRepo:        accountRepo,
		accountBalanceRepo: accountBalanceRepo,
		accountDetailRepo:  accountDetail,
	}
}

func (svc *taskSvc) CreateTable() {
	// create user table
	err := svc.createTableUser()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'users' created successfully!")

	// create banner table
	err = svc.createTableBanner()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'banners' created successfully!")

	// create account table
	err = svc.createTableAccount()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'account' created successfully!")

	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	// insert user
	userID, err := svc.insertUser(tx)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'users' successfully!")

	// insert banner
	err = svc.insertBanner(tx, userID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'banners' successfully!")

	// insert acocunts
	err = svc.insertAccount(tx, userID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'acocunts' successfully!")

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	return
}

func (svc *taskSvc) createTableUser() error {
	err := svc.userRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.userGreetingRepo.CreateTable()
	return err
}

func (svc *taskSvc) createTableBanner() error {
	err := svc.bannerRepo.CreateTable()
	return err
}

func (svc *taskSvc) createTableAccount() error {
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
	err = svc.accountRepo.CreateTableView()
	return err
}

func (svc *taskSvc) insertUser(tx *sql.Tx) (userID string, err error) {
	timeNow := time.Now().UTC()
	userID, err = svc.userRepo.Insert(tx, orm.User{
		Name:        "admin",
		Password:    "123456",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	if err != nil {
		return userID, err
	}
	err = svc.userGreetingRepo.Insert(tx, orm.UserGreeting{
		UserID:      userID,
		Greeting:    "Have a nice day Clare",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	return userID, err
}

func (svc *taskSvc) insertBanner(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	titles := []string{"Emily", "AbcdEfghiJKlmN", "Jone Kiersten", "Emily", "Emily", "MarkYu Gonzales"}
	for i := range titles {
		err := svc.bannerRepo.Insert(tx, orm.Banner{
			UserID:      userID,
			Title:       types.NewNullString(titles[i]),
			Image:       types.NewNullString("https://dummyimage.com/54x54/999/fff"),
			CreatedBy:   "admin",
			CreatedDate: timeNow,
			UpdatedBy:   types.NewNullString("admin"),
			UpdatedDate: types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *taskSvc) insertAccount(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	names := []string{
		"Saving Account", "Saving Account",
		"Credit Loan", "Travel New York", "Need To Repay",
	}
	typeAccounts := []string{
		"Smart account", "Smart account", "Credit Loan",
		"Goal driven savings", "Credit Loan",
	}
	issuers := []string{"TestLab", "TestLab", "", "TestLab", ""}
	amounts := []float64{
		62000, 8837999, 300.1, 30000, 30000,
	}
	isManinAccount := []bool{true, false, false, false, false}
	progress := []int64{0, 0, 0, 24, 0}
	for i := 0; i < 5; i++ {
		accountID, err := svc.accountRepo.Insert(tx, orm.Account{
			UserID:        userID,
			Name:          types.NewNullString(names[i]),
			Type:          types.NewNullString(typeAccounts[i]),
			Currency:      types.NewNullString("THB"),
			AccountNumber: types.NewNullString("568-2-81740-9"),
			Issuer:        types.NewNullString(issuers[i]),
			CreatedBy:     "admin",
			CreatedDate:   timeNow,
			UpdatedBy:     types.NewNullString("admin"),
			UpdatedDate:   types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
		err = svc.accountBalanceRepo.Insert(tx, orm.AccountBalance{
			AccountID:   accountID,
			UserID:      userID,
			Amount:      types.NewNullFloat64(amounts[i]),
			CreatedBy:   "admin",
			CreatedDate: timeNow,
			UpdatedBy:   types.NewNullString("admin"),
			UpdatedDate: types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
		err = svc.accountDetailRepo.Insert(tx, orm.AccountDetail{
			AccountID:      accountID,
			UserID:         userID,
			Color:          types.NewNullString("red"),
			IsManinAccount: isManinAccount[i],
			Progress:       types.NewNullInt64(progress[i]),
			CreatedBy:      "admin",
			CreatedDate:    timeNow,
			UpdatedBy:      types.NewNullString("admin"),
			UpdatedDate:    types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *taskSvc) MockData() {
}
