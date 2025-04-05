package debitcard_svc

import (
	"math"

	"github.com/kongzyeons/go-bank/internal/models"
	debitcard_repo "github.com/kongzyeons/go-bank/internal/repositories/debit-card"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type DebitCardSvc interface {
	GetList(req models.DebitCardGetListReq) response.Response[*models.DebitCardGetListRes]
}

type debitcardSvc struct {
	debitCardRepo debitcard_repo.DebitCardRepo
}

func NewDebitCardSvc(debitCardRepo debitcard_repo.DebitCardRepo) DebitCardSvc {
	return &debitcardSvc{
		debitCardRepo: debitCardRepo,
	}
}

func (svc *debitcardSvc) GetList(req models.DebitCardGetListReq) response.Response[*models.DebitCardGetListRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.DebitCardGetListRes](valMap)
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
		return response.ValidationFailed[*models.DebitCardGetListRes](valMap)
	}
	fieldSort, err := validation.GetField(req.SortBy.Field, "json", models.AccountGetListResult{})
	if err != nil {
		return response.BadRequest[*models.DebitCardGetListRes]("field sort not found")
	}
	if fieldSort.Tag.Get("db") == "" {
		return response.BadRequest[*models.DebitCardGetListRes]("field sort not found")
	}
	req.SortBy.Field = fieldSort.Tag.Get("db")
	req.SortBy.FieldType = fieldSort.Type.Kind()

	dataDB, total, err := svc.debitCardRepo.GetList(req)
	if err != nil {
		return response.InternalServerError[*models.DebitCardGetListRes](err, err.Error())
	}

	results := make([]models.DebitCardGetListResult, len(dataDB))
	if len(dataDB) <= 0 {
		res := &models.DebitCardGetListRes{
			Results:      results,
			TotalPages:   0,
			TotalResults: total,
			Page:         req.Page,
			PerPage:      req.PerPage,
		}
		return response.Ok(&res)
	}
	for i := range dataDB {
		results[i] = toDebitCardGetListResult(dataDB[i])
	}
	totalPages := int64(math.Ceil(float64(total) / float64(req.PerPage)))
	res := &models.DebitCardGetListRes{
		Results:      results,
		TotalPages:   totalPages,
		TotalResults: total,
		Page:         req.Page,
		PerPage:      req.PerPage,
	}
	return response.Ok(&res)
}
