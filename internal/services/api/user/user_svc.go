package user_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/validation"
)

type UserSvc interface {
	SetGreeting(req models.UserSetGreetingReq) response.Response[any]
}

type userSvc struct {
}

func NewUserSvc() UserSvc {
	return &userSvc{}
}

func (svc *userSvc) SetGreeting(req models.UserSetGreetingReq) response.Response[any] {
	// OPTIONAL :

	if valMap := validation.ValidateReq(&req); len(valMap) > 0 {
		return response.ValidationFailed[any](valMap)
	}

	return response.Ok[any](nil)
}
