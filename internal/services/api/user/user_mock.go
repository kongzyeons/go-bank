package user_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

type userSvcMock struct {
	mock.Mock
}

func NewUserSvcMock() *userSvcMock {
	return &userSvcMock{}
}

func (m *userSvcMock) GetGeeting(req models.UserGetGeetingReq) response.Response[*models.UserGetGeetingRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.UserGetGeetingRes])
}
