package accountflag_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type accountFlagRepoMock struct {
	mock.Mock
}

func NewAccountFlagRepoMock() *accountFlagRepoMock {
	return &accountFlagRepoMock{}
}

func (m *accountFlagRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *accountFlagRepoMock) Insert(tx *sql.Tx, req orm.AccountFlag) error {
	args := m.Called(tx, req)
	return args.Error(0)
}
