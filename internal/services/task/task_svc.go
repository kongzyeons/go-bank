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
	accountflag_repo "github.com/kongzyeons/go-bank/internal/repositories/account-flags"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	debitcard_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card"
	debitcarddesign_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-design"
	debitcarddetails_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-details"
	debitcardstatus_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card-status"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type TaskSvc interface {
	CreateTable()
	MockData()
}

type taskSvc struct {
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

func NewTaskSvc(
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
) TaskSvc {
	return &taskSvc{
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

	// create debit card table
	err = svc.createTableDebitCard()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'debit card' created successfully!")

	// create transection table
	err = svc.createTableTransection()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'transection' created successfully!")

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

	// insert debit card
	err = svc.insertDebitCard(tx, userID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'debit card' successfully!")

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
	err = svc.accountFalgRepo.CreateTable()
	if err != nil {
		return err
	}
	err = svc.accountRepo.CreateTableView()
	return err
}

func (svc *taskSvc) createTableDebitCard() error {
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

func (svc *taskSvc) createTableTransection() error {
	err := svc.transectionRepo.CreateTable()
	if err != nil {
		return err
	}
	return nil
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
			Name:           types.NewNullString(names[i]),
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

func (svc *taskSvc) insertDebitCard(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	names := []string{"My Salary", "For My Dream", "For My Dream", "For My Dream"}
	status := []string{"In progress", "In progress", "", ""}
	numbers := []string{"", "", "9440-78xx-xxxx-3115", "9440-78xx-xxxx-3115"}
	for i := 0; i < 4; i++ {
		cardID, err := svc.debitcardRepo.Insert(tx, orm.DebitCard{
			UserID:      userID,
			Name:        types.NewNullString(names[i]),
			CreatedBy:   "admin",
			CreatedDate: timeNow,
			UpdatedBy:   types.NewNullString("admin"),
			UpdatedDate: types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
		err = svc.debitcardstatuRepo.Insert(tx, orm.DebitCardStatus{
			CardID:      cardID,
			UserID:      userID,
			Status:      types.NewNullString(status[i]),
			CreatedBy:   "admin",
			CreatedDate: timeNow,
			UpdatedBy:   types.NewNullString("admin"),
			UpdatedDate: types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
		err = svc.debitCardSDetailRepo.Insert(tx, orm.DebitCardDetail{
			CardID:      cardID,
			UserID:      userID,
			Issuer:      types.NewNullString("TestLab"),
			Number:      types.NewNullString(numbers[i]),
			CreatedBy:   "admin",
			CreatedDate: timeNow,
			UpdatedBy:   types.NewNullString("admin"),
			UpdatedDate: types.NewNullTime(timeNow),
		})
		if err != nil {
			return err
		}
		err = svc.debitCarddesignRepo.Insert(tx, orm.DebitCardDesign{
			CardID:      cardID,
			UserID:      userID,
			Color:       types.NewNullString("red"),
			BorderColor: types.NewNullString("blue"),
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

func (svc *taskSvc) MockData() {
}
