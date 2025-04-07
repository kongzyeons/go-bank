package accountbalance_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type accountBalanceRepoMock struct {
	mock.Mock
}

func NewAccountBalanceRepoMock() *accountBalanceRepoMock {
	return &accountBalanceRepoMock{}
}

func (m *accountBalanceRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *accountBalanceRepoMock) Insert(tx *sql.Tx, req orm.AccountBalance) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *accountBalanceRepoMock) Update(tx *sql.Tx, req orm.AccountBalance) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *accountBalanceRepoMock) GetByID(accountID string) (*orm.AccountBalance, error) {
	args := m.Called(accountID)
	return args.Get(0).(*orm.AccountBalance), args.Error(1)
}
