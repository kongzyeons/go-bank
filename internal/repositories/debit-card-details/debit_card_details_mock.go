package debitcarddetails_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type debitCardSDetailRepoMock struct {
	mock.Mock
}

func NewDebitCardSDetailRepoMock() *debitCardSDetailRepoMock {
	return &debitCardSDetailRepoMock{}
}

func (m *debitCardSDetailRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *debitCardSDetailRepoMock) Insert(tx *sql.Tx, req orm.DebitCardDetail) error {
	args := m.Called(tx, req)
	return args.Error(0)
}
