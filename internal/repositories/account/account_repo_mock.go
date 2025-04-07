package account_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type accountRepoMock struct {
	mock.Mock
}

func NewAccountRepoMock() *accountRepoMock {
	return &accountRepoMock{}
}

func (m *accountRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}
func (m *accountRepoMock) CreateTableView() error {
	args := m.Called()
	return args.Error(0)
}
func (m *accountRepoMock) Insert(tx *sql.Tx, req orm.Account) (accountID string, err error) {
	args := m.Called(tx, req)
	return args.String(0), args.Error(1)
}
func (m *accountRepoMock) GetList(req models.AccountGetListReq) (res []orm.AccountVW, total int64, err error) {
	args := m.Called(req)
	return args.Get(0).([]orm.AccountVW), args.Get(1).(int64), args.Error(2)
}
func (m *accountRepoMock) GetByID(accountID string) (res *orm.Account, err error) {
	args := m.Called(accountID)
	return args.Get(0).(*orm.Account), args.Error(1)
}
