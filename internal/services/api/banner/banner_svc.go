package banner_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
	"github.com/redis/go-redis/v9"
)

type BannerSvc interface {
	GetList(req models.BannerGetListReq) response.Response[*models.BannerGetListRes]
}

type bannerSvc struct {
	redisClient *redis.Client
	bannerRepo  banner_repo.BannerRepo
}

func NewBannerSvc(
	redisClient *redis.Client,
	bannerRepo banner_repo.BannerRepo,
) BannerSvc {
	return &bannerSvc{
		redisClient: redisClient,
		bannerRepo:  bannerRepo,
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

	// get redis
	reqJson, err := json.Marshal(req)
	if err != nil {
		return response.InternalServerError[*models.BannerGetListRes](err, err.Error())
	}
	key := fmt.Sprintf("bannerSvc::%s", string(reqJson))
	if dataDBJson, err := svc.redisClient.Get(context.Background(), key).Result(); err == nil {
		var res *models.BannerGetListRes
		if json.Unmarshal([]byte(dataDBJson), &res) == nil {
			return response.Ok(&res)
		}
	}

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

	// Redis SET
	if data, err := json.Marshal(res); err == nil {
		svc.redisClient.Set(context.Background(), key, string(data), time.Second*10)
	}

	return response.Ok(&res)
}
