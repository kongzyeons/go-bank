package task_svc

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
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

type TaskInsertSvc interface {
	InsertAdminData() error
	InsertSimpleData() error
	MockDataBanner()
}

type taskInsertSvc struct {
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

func NewTaskInsertSvc(
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
) TaskInsertSvc {
	return &taskInsertSvc{
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

func (svc *taskInsertSvc) InsertAdminData() error {
	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	// insert user
	userID, err := svc.insertUserAdmin(tx)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'users' successfully!")

	// insert banner
	err = svc.insertBanner(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'banners' successfully!")

	// insert acocunts
	err = svc.insertAccount(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'acocunts' successfully!")

	// insert debit card
	err = svc.insertDebitCard(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'debit card' successfully!")

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	return nil
}

func (svc *taskInsertSvc) InsertSimpleData() error {
	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}

	// insert user
	userID, err := svc.insertUser(tx)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'users' successfully!")

	// insert banner
	err = svc.insertBanner(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'banners' successfully!")

	// insert acocunts
	err = svc.insertAccount(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'acocunts' successfully!")

	// insert debit card
	err = svc.insertDebitCard(tx, userID)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Insert 'debit card' successfully!")

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return err
	}
	return nil
}

func (svc *taskInsertSvc) insertUser(tx *sql.Tx) (userID string, err error) {
	timeNow := time.Now().UTC()
	name := uuid.New()
	userID, err = svc.userRepo.Insert(tx, orm.User{
		Name:        name.String(),
		Password:    types.NewNullString("123456"),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	if err != nil {
		return userID, err
	}
	err = svc.userGreetingRepo.Insert(tx, orm.UserGreeting{
		UserID:      userID,
		Greeting:    types.NewNullString("Have a nice day Clare"),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	return userID, err
}

func (svc *taskInsertSvc) insertUserAdmin(tx *sql.Tx) (userID string, err error) {
	timeNow := time.Now().UTC()
	userID, err = svc.userRepo.InsertMock(tx, orm.User{
		Name:        "admin",
		Password:    types.NewNullString("123456"),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	if err != nil {
		return userID, err
	}
	err = svc.userGreetingRepo.Insert(tx, orm.UserGreeting{
		UserID:      userID,
		Greeting:    types.NewNullString("Have a nice day Clare"),
		CreatedBy:   types.NewNullString("admin"),
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	return userID, err
}

func (svc *taskInsertSvc) insertBanner(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	titles := []string{"Emily", "AbcdEfghiJKlmN", "Jone Kiersten", "Emily", "Emily", "MarkYu Gonzales"}
	for i := range titles {
		err := svc.bannerRepo.Insert(tx, orm.Banner{
			UserID:      userID,
			Title:       types.NewNullString(titles[i]),
			Image:       types.NewNullString("https://dummyimage.com/54x54/999/fff"),
			CreatedBy:   types.NewNullString("admin"),
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

func (svc *taskInsertSvc) insertAccount(tx *sql.Tx, userID string) (err error) {
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
			CreatedBy:     types.NewNullString("admin"),
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
			CreatedBy:   types.NewNullString("admin"),
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
			CreatedBy:      types.NewNullString("admin"),
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

func (svc *taskInsertSvc) insertDebitCard(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	names := []string{"My Salary", "For My Dream", "For My Dream", "For My Dream"}
	status := []string{"In progress", "In progress", "", ""}
	numbers := []string{"", "", "9440-78xx-xxxx-3115", "9440-78xx-xxxx-3115"}
	for i := 0; i < 4; i++ {
		cardID, err := svc.debitcardRepo.Insert(tx, orm.DebitCard{
			UserID:      userID,
			Name:        types.NewNullString(names[i]),
			CreatedBy:   types.NewNullString("admin"),
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
			CreatedBy:   types.NewNullString("admin"),
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
			CreatedBy:   types.NewNullString("admin"),
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
			CreatedBy:   types.NewNullString("admin"),
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

func (svc *taskInsertSvc) MockDataBanner() {

	var defaultUUID uuid.UUID // zero value

	var banners []orm.Banner
	for i := 0; i < 50000; i++ {
		banners = append(banners, orm.Banner{
			UserID:      defaultUUID.String(),
			Title:       types.NewNullString("test"),
			Description: types.NewNullString("test"),
			Image:       types.NewNullString("test"),
			DummyCol11:  types.NewNullString("test"),
			CreatedBy:   types.NewNullString("test"),
			CreatedDate: time.Now(),
			UpdatedBy:   types.NewNullString("test"),
			UpdatedDate: types.NewNullTime(time.Now()),
		})
	}

	err := InsertBannersConcurrently(svc.db, banners)
	if err != nil {
		fmt.Println("Insert failed:", err)
	} else {
		fmt.Println("Insert successful")
	}

}

func insertBannerChunk(db *sqlx.DB, banners []orm.Banner) error {
	query := `INSERT INTO banners (
		user_id, title, description, image, 
		dummy_col_11, created_by, created_date, updated_by, updated_date
	) VALUES `

	valueStrings := make([]string, 0, len(banners))
	valueArgs := make([]interface{}, 0, len(banners)*9)

	for i, b := range banners {
		valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			i*9+1, i*9+2, i*9+3, i*9+4, i*9+5, i*9+6, i*9+7, i*9+8, i*9+9))

		valueArgs = append(valueArgs,
			b.UserID,
			b.Title,
			b.Description,
			b.Image,
			b.DummyCol11,
			b.CreatedBy,
			b.CreatedDate,
			b.UpdatedBy,
			b.UpdatedDate,
		)
	}

	query += strings.Join(valueStrings, ",")
	_, err := db.Exec(query, valueArgs...)
	return err
}

func InsertBannersConcurrently(db *sqlx.DB, banners []orm.Banner) error {
	const chunkSize = 1000
	var wg sync.WaitGroup
	errChan := make(chan error, len(banners)/chunkSize+1)

	for i := 0; i < len(banners); i += chunkSize {
		end := i + chunkSize
		if end > len(banners) {
			end = len(banners)
		}
		chunk := banners[i:end]

		wg.Add(1)
		go func(chunkData []orm.Banner) {
			defer wg.Done()
			if err := insertBannerChunk(db, chunkData); err != nil {
				errChan <- err
			}
		}(chunk)
	}

	wg.Wait()
	close(errChan)

	if len(errChan) > 0 {
		return <-errChan
	}
	return nil
}
