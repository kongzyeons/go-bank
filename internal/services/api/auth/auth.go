package auth_service

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
)

type AuthSvc interface {
	Register(req models.RegisterReq) response.Response[any]
	Login(req models.LoginReq) response.Response[*models.LoginRes]
	Ping(req models.PingReq) response.Response[*models.PingRes]
	Logout() response.Response[any]
}

type authSvc struct {
}

func NewAuthSvc() AuthSvc {
	return &authSvc{}
}

func (svc *authSvc) Register(req models.RegisterReq) response.Response[any] {
	return response.Ok[any](nil)
}
func (svc *authSvc) Login(req models.LoginReq) response.Response[*models.LoginRes] {
	res := &models.LoginRes{}
	return response.Ok(&res)
}
func (svc *authSvc) Ping(req models.PingReq) response.Response[*models.PingRes] {
	res := &models.PingRes{}
	return response.Ok(&res)
}
func (svc *authSvc) Logout() response.Response[any] {
	return response.Ok[any](nil)
}
