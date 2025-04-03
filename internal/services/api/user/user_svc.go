package user_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type UserSvc interface {
	GetGeeting(req models.UserGetGeetingReq) response.Response[*models.UserGetGeetingRes]
	SetGreeting(req models.UserSetGreetingReq) response.Response[any]
}

type userSvc struct {
	userGreetingRepo usergreeting_repo.UserGreetingRepo
}

func NewUserSvc(userGreetingRepo usergreeting_repo.UserGreetingRepo) UserSvc {
	return &userSvc{
		userGreetingRepo: userGreetingRepo,
	}
}

func (svc *userSvc) GetGeeting(req models.UserGetGeetingReq) response.Response[*models.UserGetGeetingRes] {
	dataDB, err := svc.userGreetingRepo.GetByID(req.UserID)
	if err != nil {
		return response.InternalServerError[*models.UserGetGeetingRes](err, err.Error())
	}
	if dataDB == nil {
		return response.Notfound[*models.UserGetGeetingRes]("not found user id")
	}

	res := &models.UserGetGeetingRes{
		Username: req.Username,
		Greeting: dataDB.Greeting,
	}
	return response.Ok(&res)
}

func (svc *userSvc) SetGreeting(req models.UserSetGreetingReq) response.Response[any] {
	// OPTIONAL :

	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[any](valMap)
	}

	return response.Ok[any](nil)
}
