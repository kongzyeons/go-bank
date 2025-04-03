package user_task

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
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
	db               *sqlx.DB
	userRepo         user_repo.UserRepo
	userGreetingRepo usergreeting_repo.UserGreetingRepo
	bannerRepo       banner_repo.BannerRepo
}

func NewTaskSvc(
	db *sqlx.DB,
	userRepo user_repo.UserRepo,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
	bannerRepo banner_repo.BannerRepo,
) TaskSvc {
	return &taskSvc{
		db:               db,
		userRepo:         userRepo,
		userGreetingRepo: userGreetingRepo,
		bannerRepo:       bannerRepo,
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

	// create user_greeting table
	err = svc.createTableUserGreeting()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'user_greetings' created successfully!")

	// create banner table
	err = svc.createTableBanner()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'banners' created successfully!")

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

	// insert user_greeting
	err = svc.insertUserGreeting(tx, userID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'user_greetings' successfully!")

	// insert banner
	err = svc.insertBanner(tx, userID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Insert 'banners' successfully!")

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
	return err
}

func (svc *taskSvc) createTableUserGreeting() error {
	err := svc.userGreetingRepo.CreateTable()
	return err
}

func (svc *taskSvc) createTableBanner() error {
	err := svc.bannerRepo.CreateTable()
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
	return userID, err
}

func (svc *taskSvc) insertUserGreeting(tx *sql.Tx, userID string) (err error) {
	timeNow := time.Now().UTC()
	err = svc.userGreetingRepo.Insert(tx, orm.UserGreeting{
		UserID:      userID,
		Greeting:    "Have a nice day Clare",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	return err
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

func (svc *taskSvc) MockData() {
}
