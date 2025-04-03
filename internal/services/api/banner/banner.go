package banner_svc

import (
	"math"

	"github.com/kongzyeons/go-bank/internal/models"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type BannerSvc interface {
	GetList(req models.BannerGetListReq) response.Response[*models.BannerGetListRes]
}

type bannerSvc struct {
	bannerRepo banner_repo.BannerRepo
}

func NewBannerSvc(bannerRepo banner_repo.BannerRepo) BannerSvc {
	return &bannerSvc{
		bannerRepo: bannerRepo,
	}
}

func (svc *bannerSvc) GetList(req models.BannerGetListReq) response.Response[*models.BannerGetListRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.BannerGetListRes](valMap)
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
		return response.ValidationFailed[*models.BannerGetListRes](valMap)
	}
	fieldSort, err := validation.GetField(req.SortBy.Field, "json", models.BannerGetListResult{})
	if err != nil {
		return response.BadRequest[*models.BannerGetListRes]("field sort not found")
	}
	if fieldSort.Tag.Get("db") == "" {
		return response.BadRequest[*models.BannerGetListRes]("field sort not found")
	}
	req.SortBy.Field = fieldSort.Tag.Get("db")
	req.SortBy.FieldType = fieldSort.Type.Kind()

	dataDB, total, err := svc.bannerRepo.GetList(req)
	if err != nil {
		return response.InternalServerError[*models.BannerGetListRes](err, err.Error())
	}

	results := make([]models.BannerGetListResult, len(dataDB))
	if len(dataDB) <= 0 {
		res := &models.BannerGetListRes{
			Results:      results,
			TotalPages:   0,
			TotalResults: total,
			Page:         req.Page,
			PerPage:      req.PerPage,
		}
		return response.Ok(&res)
	}
	for i := range dataDB {
		results[i] = toBannerGetListResult(dataDB[i])
	}
	totalPages := int64(math.Ceil(float64(total) / float64(req.PerPage)))
	res := &models.BannerGetListRes{
		Results:      results,
		TotalPages:   totalPages,
		TotalResults: total,
		Page:         req.Page,
		PerPage:      req.PerPage,
	}
	return response.Ok(&res)
}
