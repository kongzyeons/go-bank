package transaction_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type transactionRepoMock struct {
	mock.Mock
}

func NewTransactionRepoMock() *transactionRepoMock {
	return &transactionRepoMock{}
}

func (m *transactionRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *transactionRepoMock) Insert(tx *sql.Tx, req orm.Transaction) error {
	args := m.Called(tx, req)
	return args.Error(0)
}
