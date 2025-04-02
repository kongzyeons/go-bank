package user_task

import (
	"context"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/types"
)

type UserTask interface {
	CreateTable()
	MockData()
}

type userTask struct {
	db               *sqlx.DB
	userRepo         user_repo.UserRepo
	userGreetingRepo usergreeting_repo.UserGreetingRepo
}

func NewUserTask(
	db *sqlx.DB,
	userRepo user_repo.UserRepo,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
) UserTask {
	return &userTask{
		db:               db,
		userRepo:         userRepo,
		userGreetingRepo: userGreetingRepo,
	}
}

func (svc *userTask) CreateTable() {
	// create user table
	err := svc.userRepo.CreateTable()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'users' created successfully!")

	// create user_greeting table
	err = svc.userGreetingRepo.CreateTable()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Table 'user_greetings' created successfully!")

	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	timeNow := time.Now().UTC()

	// insert user
	userID, err := svc.userRepo.Insert(tx, orm.User{
		Name:        "admin",
		Password:    "123456",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})
	if err != nil {
		log.Println(err)
		return
	}

	// insert user_greeting
	err = svc.userGreetingRepo.Insert(tx, orm.UserGreeting{
		UserID:      userID,
		Greeting:    "Have a nice day Clare",
		CreatedBy:   "admin",
		CreatedDate: timeNow,
		UpdatedBy:   types.NewNullString("admin"),
		UpdatedDate: types.NewNullTime(timeNow),
	})

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println(err)
		return
	}

	return
}
func (svc *userTask) MockData() {
}
