package accountdetail_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type accountDetailRepoMock struct {
	mock.Mock
}

func NewAccountDetailRepoMock() *accountDetailRepoMock {
	return &accountDetailRepoMock{}
}

func (m *accountDetailRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *accountDetailRepoMock) Insert(tx *sql.Tx, req orm.AccountDetail) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *accountDetailRepoMock) Update(tx *sql.Tx, req orm.AccountDetail) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *accountDetailRepoMock) GetByID(accountID string) (*orm.AccountDetail, error) {
	args := m.Called(accountID)
	return args.Get(0).(*orm.AccountDetail), args.Error(1)
}
