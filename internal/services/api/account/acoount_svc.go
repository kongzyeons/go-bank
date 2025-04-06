package account_svc

import (
	"bytes"
	"context"
	"encoding/base64"
	"image/png"
	"math"
	"strings"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/queues"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type AccountSvc interface {
	GetList(req models.AccountGetListReq) response.Response[*models.AccountGetListRes]
	Edit(req models.AccountEditReq) response.Response[*models.AccountEditRes]
	GetQrcode(req models.AccountGetQrcodeReq) response.Response[*models.AccountGetQrcodeRes]
	SetIsmain(req models.AccountSetIsmainReq) response.Response[*models.AccountSetIsmainRes]
	AddMoney(req models.AccountAddMoneyReq) response.Response[any]
	Withdrawl(req models.AccountWithdrawlReq) response.Response[any]
}

type accountSvc struct {
	db                 *sqlx.DB
	eventProducer      queues.EventProducer
	accountRepo        account_repo.AccountRepo
	accountDetailRepo  accountdetail_repo.AccountDetailRepo
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo
	transectionRepo    transaction_repo.TransactionRepo
}

func NewAccountSvc(
	db *sqlx.DB,
	eventProducer queues.EventProducer,
	accountRepo account_repo.AccountRepo,
	accountDetailRepo accountdetail_repo.AccountDetailRepo,
	accountBalanceRepo accountbalance_repo.AccountBalanceRepo,
	transectionRepo transaction_repo.TransactionRepo,
) AccountSvc {
	return &accountSvc{
		db:                 db,
		eventProducer:      eventProducer,
		accountRepo:        accountRepo,
		accountDetailRepo:  accountDetailRepo,
		accountBalanceRepo: accountBalanceRepo,
		transectionRepo:    transectionRepo,
	}
}

func (svc *accountSvc) GetList(req models.AccountGetListReq) response.Response[*models.AccountGetListRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.AccountGetListRes](valMap)
	}

	// dafault
	if req.SortBy.Field == "" {
		req.SortBy.Field = "updatedDate"
		req.SortBy.Mode = "desc"
	}

	mapSortMode := map[string]bool{
		"asc": true, "ASC": true,
		"desc": true, "DESC": true,
	}
	if _, ok := mapSortMode[req.SortBy.Mode]; !ok {
		valMap := map[string]string{
			"mode": "mode must be asc desc",
		}
		return response.ValidationFailed[*models.AccountGetListRes](valMap)
	}
	fieldSort, err := validation.GetField(req.SortBy.Field, "json", models.AccountGetListResult{})
	if err != nil {
		return response.BadRequest[*models.AccountGetListRes]("field sort not found")
	}
	if fieldSort.Tag.Get("db") == "" {
		return response.BadRequest[*models.AccountGetListRes]("field sort not found")
	}
	req.SortBy.Field = fieldSort.Tag.Get("db")
	req.SortBy.FieldType = fieldSort.Type.Kind()

	dataDB, total, err := svc.accountRepo.GetList(req)
	if err != nil {
		return response.InternalServerError[*models.AccountGetListRes](err, err.Error())
	}

	var totalAnount float64
	results := make([]models.AccountGetListResult, len(dataDB))
	if len(dataDB) <= 0 {
		res := &models.AccountGetListRes{
			Results:      results,
			TotalAnount:  totalAnount,
			TotalPages:   0,
			TotalResults: total,
			Page:         req.Page,
			PerPage:      req.PerPage,
		}
		return response.Ok(&res)
	}
	for i := range dataDB {
		results[i] = toBannerGetListResult(dataDB[i])
		totalAnount += results[i].Amount
	}
	totalPages := int64(math.Ceil(float64(total) / float64(req.PerPage)))
	res := &models.AccountGetListRes{
		Results:      results,
		TotalAnount:  totalAnount,
		TotalPages:   totalPages,
		TotalResults: total,
		Page:         req.Page,
		PerPage:      req.PerPage,
	}
	return response.Ok(&res)
}

func (svc *accountSvc) Edit(req models.AccountEditReq) response.Response[*models.AccountEditRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.AccountEditRes](valMap)
	}

	dataDB, err := svc.accountDetailRepo.GetByID(strings.TrimSpace(req.AccountID))
	if err != nil {
		return response.InternalServerError[*models.AccountEditRes](err, err.Error())
	}
	if dataDB == nil {
		return response.BadRequest[*models.AccountEditRes]("account id not found")
	}
	if dataDB.UserID != req.UserID {
		return response.Unauthorized[*models.AccountEditRes]("user id unauthorized")
	}

	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AccountEditRes](err, err.Error())
	}

	err = svc.accountDetailRepo.Update(tx, editToUpdate(req, *dataDB))
	if err != nil {
		return response.InternalServerError[*models.AccountEditRes](err, err.Error())
	}
	err = svc.transectionRepo.Insert(tx, orm.Transaction{
		UserID:      dataDB.UserID,
		Name:        types.NewNullString("account:edit"),
		CreatedBy:   req.Name,
		CreatedDate: time.Now().UTC(),
	})
	if err != nil {
		return response.InternalServerError[*models.AccountEditRes](err, err.Error())
	}
	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AccountEditRes](err, err.Error())
	}

	res := &models.AccountEditRes{
		UpdatedBy:   req.Username,
		UpdatedDate: time.Now().UTC(),
	}
	return response.Ok(&res)

}

func (svc *accountSvc) GetQrcode(req models.AccountGetQrcodeReq) response.Response[*models.AccountGetQrcodeRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.AccountGetQrcodeRes](valMap)
	}

	dataDB, err := svc.accountRepo.GetByID(strings.TrimSpace(req.AccountID))
	if err != nil {
		return response.InternalServerError[*models.AccountGetQrcodeRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.AccountGetQrcodeRes]("not found account id")
	}
	if dataDB.UserID != req.UserID {
		return response.Unauthorized[*models.AccountGetQrcodeRes]("user id unauthorized")
	}

	qrCode, err := qr.Encode(dataDB.AccountID, qr.M, qr.Auto)
	if err != nil {
		return response.InternalServerError[*models.AccountGetQrcodeRes](err, err.Error())
	}
	qrCode, _ = barcode.Scale(qrCode, 256, 256)

	// Encode the image to PNG in a buffer
	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		return response.InternalServerError[*models.AccountGetQrcodeRes](err, err.Error())
	}

	res := &models.AccountGetQrcodeRes{
		QrcodeBase64: base64.StdEncoding.EncodeToString(buf.Bytes()),
	}
	return response.Ok(&res)
}

func (svc *accountSvc) SetIsmain(req models.AccountSetIsmainReq) response.Response[*models.AccountSetIsmainRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.AccountSetIsmainRes](valMap)
	}

	dataDB, err := svc.accountDetailRepo.GetByID(strings.TrimSpace(req.AccountID))
	if err != nil {
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.AccountSetIsmainRes]("not found account id")
	}
	if dataDB.IsManinAccount {
		return response.BadRequest[*models.AccountSetIsmainRes]("this account id is main")
	}
	if dataDB.UserID != req.UserID {
		return response.Unauthorized[*models.AccountSetIsmainRes]("user id unauthorized")
	}

	dataDBIsmain, err := svc.accountDetailRepo.GetByID(strings.TrimSpace(req.AccountIDIsmain))
	if err != nil {
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}
	if dataDBIsmain == nil {
		return response.Notfound[*models.AccountSetIsmainRes]("not found account id")
	}
	if !dataDBIsmain.IsManinAccount {
		return response.BadRequest[*models.AccountSetIsmainRes]("this account id is not main")
	}
	if dataDBIsmain.UserID != req.UserID {
		return response.Unauthorized[*models.AccountSetIsmainRes]("user id unauthorized")
	}

	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}

	err = svc.accountDetailRepo.Update(tx, setIsmainToUpdate(req, *dataDB, true))
	if err != nil {
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}

	err = svc.accountDetailRepo.Update(tx, setIsmainToUpdate(req, *dataDBIsmain, false))
	if err != nil {
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}

	err = svc.transectionRepo.Insert(tx, orm.Transaction{
		UserID:      dataDB.UserID,
		Name:        types.NewNullString("account:setIsmain"),
		CreatedBy:   req.Username,
		CreatedDate: time.Now().UTC(),
	})
	if err != nil {
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}
	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AccountSetIsmainRes](err, err.Error())
	}

	res := &models.AccountSetIsmainRes{
		UpdatedBy:   req.Username,
		UpdatedDate: time.Now().UTC(),
	}
	return response.Ok(&res)

}

func (svc *accountSvc) AddMoney(req models.AccountAddMoneyReq) response.Response[any] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[any](valMap)
	}

	dataDB, err := svc.accountBalanceRepo.GetByID(req.AccountID)
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[any]("not found account id")
	}
	if dataDB.UserID != req.UserID {
		return response.Unauthorized[any]("user id unauthorized")
	}

	err = svc.eventProducer.Produce(addMoneyToEvent(req))
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}

	return response.Ok[any](nil)

}

func (svc *accountSvc) Withdrawl(req models.AccountWithdrawlReq) response.Response[any] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[any](valMap)
	}

	dataDB, err := svc.accountBalanceRepo.GetByID(req.AccountID)
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[any]("not found account id")
	}
	if dataDB.UserID != req.UserID {
		return response.Unauthorized[any]("user id unauthorized")
	}
	if dataDB.Amount.Float64 < req.Ammount {
		return response.BadRequest[any]("ืnot enough money")
	}

	err = svc.eventProducer.Produce(withdrawlToEvent(req))
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}

	return response.Ok[any](nil)
}
