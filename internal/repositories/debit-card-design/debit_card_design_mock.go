package debitcarddesign_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type debitCarddesignRepoMock struct {
	mock.Mock
}

func NewDebitCarddesignRepoMock() *debitCarddesignRepoMock {
	return &debitCarddesignRepoMock{}
}

func (m *debitCarddesignRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *debitCarddesignRepoMock) Insert(tx *sql.Tx, req orm.DebitCardDesign) error {
	args := m.Called(tx, req)
	return args.Error(0)
}
