package debitcardstatus_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type debitCardStatusRepoMock struct {
	mock.Mock
}

func NewDebitCardStatusRepoMock() *debitCardStatusRepoMock {
	return &debitCardStatusRepoMock{}
}

func (m *debitCardStatusRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *debitCardStatusRepoMock) Insert(tx *sql.Tx, req orm.DebitCardStatus) error {
	args := m.Called(tx, req)
	return args.Error(0)
}
