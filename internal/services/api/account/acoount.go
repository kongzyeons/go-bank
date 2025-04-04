package account_svc

import (
	"math"

	"github.com/kongzyeons/go-bank/internal/models"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type AccountSvc interface {
	GetList(req models.AccountGetListReq) response.Response[*models.AccountGetListRes]
}

type accountSvc struct {
	accountRepo account_repo.AccountRepo
}

func NewAccountSvc(accountRepo account_repo.AccountRepo) AccountSvc {
	return &accountSvc{
		accountRepo: accountRepo,
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
