package debitcard_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

type debitCardSvcMock struct {
	mock.Mock
}

func NewDebitCardSvcMock() *debitCardSvcMock {
	return &debitCardSvcMock{}
}

func (m *debitCardSvcMock) GetList(req models.DebitCardGetListReq) response.Response[*models.DebitCardGetListRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.DebitCardGetListRes])
}
