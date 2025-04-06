package account_queue

import (
	"context"
	"encoding/json"
	"log"
	"reflect"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/queues"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/pkg/line"
)

type accountEventHandler struct {
	db                 *sqlx.DB
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo
	transectionRepo    transaction_repo.TransactionRepo
	lineNoti           line.LineAPI
}

func NewAccountEventHandler(
	db *sqlx.DB,
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo,
	transectionRepo transaction_repo.TransactionRepo,
	lineNoti line.LineAPI,
) queues.EventHandler {
	return &accountEventHandler{
		db:                 db,
		accountBalanceRepo: accountBalanceRepo,
		transectionRepo:    transectionRepo,
		lineNoti:           lineNoti,
	}
}

func (svc *accountEventHandler) Handle(topic string, eventBytes []byte) {

	switch topic {

	case reflect.TypeOf(events.AccountAddMoneyEvent{}).Name():
		event := &events.AccountAddMoneyEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		// TODO : repositoy
		dataDB, err := svc.accountBalanceRepo.GetByID(event.AccountID)
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		if dataDB == nil {
			res := response.Notfound[any]("not found account id")
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		// begin transection
		tx, err := svc.db.BeginTx(context.Background(), nil)
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		err = svc.accountBalanceRepo.Update(tx, addMoneyToUpdate(*event, *dataDB))
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		err = svc.transectionRepo.Insert(tx, orm.Transaction{
			UserID:      event.UserID,
			Name:        types.NewNullString("account:addMoney"),
			IsBank:      true,
			CreatedBy:   event.Username,
			CreatedDate: time.Now().UTC(),
		})
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		//commit transaction
		err = tx.Commit()
		if err != nil {
			res := response.InternalServerError[any](err, err.Error())
			jsonData, _ := json.MarshalIndent(res, "", "   ")
			time.Sleep(3 * time.Second)
			svc.lineNoti.SendMessage(string(jsonData))
			return
		}

		result := &models.AccountAddMoneyRes{
			AccountID:     dataDB.AccountID,
			AmmountAdd:    event.Ammount,
			AmmountResult: dataDB.Amount.Float64 + event.Ammount,
		}
		res := response.Ok(&result)
		jsonData, _ := json.MarshalIndent(res, "", "   ")
		time.Sleep(3 * time.Second)
		svc.lineNoti.SendMessage(string(jsonData))

	default:
		log.Println("no event handler")
	}
}
