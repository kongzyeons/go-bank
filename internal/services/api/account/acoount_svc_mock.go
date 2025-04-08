package account_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

type accountSvcMock struct {
	mock.Mock
}

func NewAccountSvcMock() *accountSvcMock {
	return &accountSvcMock{}
}

func (m *accountSvcMock) GetList(req models.AccountGetListReq) response.Response[*models.AccountGetListRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AccountGetListRes])
}

func (m *accountSvcMock) Edit(req models.AccountEditReq) response.Response[*models.AccountEditRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AccountEditRes])
}

func (m *accountSvcMock) GetQrcode(req models.AccountGetQrcodeReq) response.Response[*models.AccountGetQrcodeRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AccountGetQrcodeRes])
}

func (m *accountSvcMock) SetIsmain(req models.AccountSetIsmainReq) response.Response[*models.AccountSetIsmainRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.AccountSetIsmainRes])
}

func (m *accountSvcMock) AddMoney(req models.AccountAddMoneyReq) response.Response[any] {
	args := m.Called(req)
	return args.Get(0).(response.Response[any])
}

func (m *accountSvcMock) Withdrawl(req models.AccountWithdrawlReq) response.Response[any] {
	args := m.Called(req)
	return args.Get(0).(response.Response[any])
}
