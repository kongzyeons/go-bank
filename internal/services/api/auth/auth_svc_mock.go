package auth_service

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

type authSvcMock struct {
	mock.Mock
}

func NewAuthSvcMock() *authSvcMock {
	return &authSvcMock{}
}

func (m *authSvcMock) Register(req models.AuthRegisterReq) response.Response[any] {
	args := m.Called(req)
	return args.Get(0).(response.Response[any])
}

func (m *authSvcMock) Login(req models.AuthLoginReq) response.Response[*models.AuthLoginRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AuthLoginRes])
}

func (m *authSvcMock) Ping(req models.AuthPingReq) response.Response[*models.AuthPingRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AuthPingRes])
}

func (m *authSvcMock) Refresh(req models.AuthRefreshReq) response.Response[*models.AuthRefreshRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AuthRefreshRes])
}

func (m *authSvcMock) Logout(req models.AuthLogoutReq) response.Response[any] {
	args := m.Called(req)
	return args.Get(0).(response.Response[any])
}
