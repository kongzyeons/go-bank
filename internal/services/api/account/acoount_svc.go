package account_svc

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type AccountSvc interface {
	GetList(req models.AccountGetListReq) response.Response[*models.AccountGetListRes]
	Edit(req models.AccountEditReq) response.Response[*models.AccountEditRes]
}

type accountSvc struct {
	db                *sqlx.DB
	accountRepo       account_repo.AccountRepo
	accountDetailRepo accountdetail_repo.AccountDetailRepo
	transectionRepo   transaction_repo.TransactionRepo
}

func NewAccountSvc(
	db *sqlx.DB,
	accountRepo account_repo.AccountRepo,
	accountDetailRepo accountdetail_repo.AccountDetailRepo,
	transectionRepo transaction_repo.TransactionRepo,
) AccountSvc {
	return &accountSvc{
		db:                db,
		accountRepo:       accountRepo,
		accountDetailRepo: accountDetailRepo,
		transectionRepo:   transectionRepo,
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
