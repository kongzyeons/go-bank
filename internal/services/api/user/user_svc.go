package user_svc

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/kongzyeons/go-bank/internal/models"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
	"github.com/redis/go-redis/v9"
)

type UserSvc interface {
	GetGeeting(req models.UserGetGeetingReq) response.Response[*models.UserGetGeetingRes]
}

type userSvc struct {
	redisClient      *redis.Client
	userGreetingRepo usergreeting_repo.UserGreetingRepo
}

func NewUserSvc(
	redisClient *redis.Client,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
) UserSvc {
	return &userSvc{
		redisClient:      redisClient,
		userGreetingRepo: userGreetingRepo,
	}
}

func (svc *userSvc) GetGeeting(req models.UserGetGeetingReq) response.Response[*models.UserGetGeetingRes] {
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.UserGetGeetingRes](valMap)
	}

	// get redis
	reqJson, err := json.Marshal(req)
	if err != nil {
		return response.InternalServerError[*models.UserGetGeetingRes](err, err.Error())
	}
	key := fmt.Sprintf("userSvc::%s", string(reqJson))
	if dataDBJson, err := svc.redisClient.Get(context.Background(), key).Result(); err == nil {
		var res *models.UserGetGeetingRes
		if json.Unmarshal([]byte(dataDBJson), &res) == nil {
			return response.Ok(&res)
		}
	}

	dataDB, err := svc.userGreetingRepo.GetByID(req.UserID)
	if err != nil {
		return response.InternalServerError[*models.UserGetGeetingRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.UserGetGeetingRes]("not found user id")
	}

	res := &models.UserGetGeetingRes{
		Username: req.Username,
		Greeting: dataDB.Greeting.String,
	}

	// Redis SET
	if data, err := json.Marshal(res); err == nil {
		svc.redisClient.Set(context.Background(), key, string(data), time.Minute*5)
	}

	return response.Ok(&res)
}
