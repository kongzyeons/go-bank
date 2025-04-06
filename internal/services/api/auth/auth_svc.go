package auth_service

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/jwt"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
	"github.com/redis/go-redis/v9"
)

type AuthSvc interface {
	Register(req models.AuthRegisterReq) response.Response[any]
	Login(req models.AuthLoginReq) response.Response[*models.AuthLoginRes]
	Ping(req models.AuthPingReq) response.Response[*models.AuthPingRes]
	Refresh(req models.AuthRefreshReq) response.Response[*models.AuthRefreshRes]
	Logout(req models.AuthLogoutReq) response.Response[any]
}

type authSvc struct {
	db               *sqlx.DB
	redisClient      *redis.Client
	userRepo         user_repo.UserRepo
	userGreetingRepo usergreeting_repo.UserGreetingRepo
	transectionRepo  transaction_repo.TransactionRepo
}

func NewAuthSvc(
	db *sqlx.DB,
	redisClient *redis.Client,
	userRepo user_repo.UserRepo,
	userGreetingRepo usergreeting_repo.UserGreetingRepo,
	transectionRepo transaction_repo.TransactionRepo,
) AuthSvc {
	return &authSvc{
		db:               db,
		redisClient:      redisClient,
		userRepo:         userRepo,
		userGreetingRepo: userGreetingRepo,
		transectionRepo:  transectionRepo,
	}
}

func (svc *authSvc) Register(req models.AuthRegisterReq) response.Response[any] {
	// validate
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[any](valMap)
	}

	_, err := strconv.Atoi(req.Password)
	if err != nil {
		return response.BadRequest[any]("passeord must number")
	}

	dataDB, err := svc.userRepo.GetUnique(strings.TrimSpace(req.Username))
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}
	if dataDB != nil {
		return response.BadRequest[any]("username already exists")
	}

	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[any](err, err.Error())
	}

	// insert user
	userID, err := svc.userRepo.Insert(
		tx,
		toUserRepoInsert(req),
	)
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}

	// insert user_greeting
	err = svc.userGreetingRepo.Insert(
		tx,
		toUserGreeting(userID),
	)

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[any](err, err.Error())
	}

	return response.Ok[any](nil)
}
func (svc *authSvc) Login(req models.AuthLoginReq) response.Response[*models.AuthLoginRes] {
	// validate
	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[*models.AuthLoginRes](valMap)
	}

	_, err := strconv.Atoi(req.Password)
	if err != nil {
		return response.BadRequest[*models.AuthLoginRes]("passeord must number")
	}

	// find username from repo
	dataDB, err := svc.userRepo.GetUnique(strings.TrimSpace(req.Username))
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.AuthLoginRes]("not found name")
	}

	if dataDB.Password != req.Password {
		return response.BadRequest[*models.AuthLoginRes]("password invalid")
	}

	key := fmt.Sprintf("authSvc::%s", dataDB.UserID)

	// get redis
	refTokenJson, err := svc.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.Nil {
			return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
		}
	}
	if refTokenJson != "" {
		return response.BadRequest[*models.AuthLoginRes]("user already logged in")
	}

	// set jwt
	accToken, err := jwt.GenToken(jwt.GenTokenReq{
		UserID:       dataDB.UserID,
		Username:     req.Username,
		TimeDulation: 15 * time.Minute,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}
	refToken, err := jwt.GenToken(jwt.GenTokenReq{
		UserID:       dataDB.UserID,
		Username:     req.Username,
		TimeDulation: 24 * time.Minute,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}

	// set redis
	valu, err := json.Marshal(orm.Auth{
		AccToken: accToken,
		RefToken: refToken,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}
	err = svc.redisClient.Set(context.Background(), key, string(valu), 24*time.Hour).Err()
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}

	// transection
	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}
	err = svc.transectionRepo.Insert(tx, orm.Transaction{
		UserID:      dataDB.UserID,
		Name:        types.NewNullString("auth:login"),
		CreatedBy:   dataDB.Name,
		CreatedDate: time.Now().UTC(),
	})
	if err != nil {
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}
	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[*models.AuthLoginRes](err, err.Error())
	}

	res := &models.AuthLoginRes{
		AccToken: accToken,
		RefToken: refToken,
	}

	return response.Ok(&res)
}

func (svc *authSvc) Ping(req models.AuthPingReq) response.Response[*models.AuthPingRes] {
	res := &models.AuthPingRes{
		UserID:   req.UserID,
		Username: req.Username,
	}
	return response.Ok(&res)
}

func (svc *authSvc) Refresh(req models.AuthRefreshReq) response.Response[*models.AuthRefreshRes] {
	key := fmt.Sprintf("authSvc::%v", req.UserID)

	authTokenStr, err := svc.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.Nil {
			return response.InternalServerError[*models.AuthRefreshRes](err, err.Error())
		}
	}
	if err == redis.Nil || authTokenStr == "" {
		return response.Unauthorized[*models.AuthRefreshRes]("not found user_id")
	}

	var authToken orm.Auth
	err = json.Unmarshal([]byte(authTokenStr), &authToken)
	if err != nil {
		return response.Unauthorized[*models.AuthRefreshRes](err.Error())
	}

	if authToken.RefToken != req.RefToken {
		return response.BadRequest[*models.AuthRefreshRes]("token invalid")
	}

	// set jwt
	accToken, err := jwt.GenToken(jwt.GenTokenReq{
		UserID:       req.UserID,
		Username:     req.UserID,
		TimeDulation: 15 * time.Minute,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthRefreshRes](err, err.Error())
	}
	refToken, err := jwt.GenToken(jwt.GenTokenReq{
		UserID:       req.UserID,
		Username:     req.UserID,
		TimeDulation: 24 * time.Hour,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthRefreshRes](err, err.Error())
	}

	valu, err := json.Marshal(orm.Auth{
		AccToken: accToken,
		RefToken: refToken,
	})
	if err != nil {
		return response.InternalServerError[*models.AuthRefreshRes](err, err.Error())
	}
	err = svc.redisClient.Set(context.Background(), key, string(valu), 24*time.Hour).Err()
	if err != nil {
		return response.InternalServerError[*models.AuthRefreshRes](err, err.Error())
	}

	res := &models.AuthRefreshRes{
		AccToken: accToken,
		RefToken: refToken,
	}
	return response.Ok(&res)
}
func (svc *authSvc) Logout(req models.AuthLogoutReq) response.Response[any] {
	key := fmt.Sprintf("authSvc::%v", req.UserID)

	authTokenStr, err := svc.redisClient.Get(context.Background(), key).Result()
	if err != nil {
		if err != redis.Nil {
			return response.InternalServerError[any](err, err.Error())
		}
	}
	if err == redis.Nil || authTokenStr == "" {
		return response.Unauthorized[any]("not found user_id")
	}

	err = svc.redisClient.Del(context.Background(), key).Err()
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}

	// transection
	// begin transection
	tx, err := svc.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[any](err, err.Error())
	}
	err = svc.transectionRepo.Insert(tx, orm.Transaction{
		UserID:      req.UserID,
		Name:        types.NewNullString("auth:logout"),
		CreatedBy:   req.Username,
		CreatedDate: time.Now().UTC(),
	})
	if err != nil {
		return response.InternalServerError[any](err, err.Error())
	}
	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return response.InternalServerError[any](err, err.Error())
	}

	return response.Ok[any](nil)
}
