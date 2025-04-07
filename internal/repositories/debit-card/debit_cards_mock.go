package debitcard_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type debitCardRepoMock struct {
	mock.Mock
}

func NewDebitCardRepoMock() *debitCardRepoMock {
	return &debitCardRepoMock{}
}

func (m *debitCardRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *debitCardRepoMock) CreateTableView() error {
	args := m.Called()
	return args.Error(0)
}

func (m *debitCardRepoMock) Insert(tx *sql.Tx, req orm.DebitCard) (string, error) {
	args := m.Called(tx, req)
	return args.String(0), args.Error(1)
}

func (m *debitCardRepoMock) GetList(req models.DebitCardGetListReq) ([]orm.DebitCardVW, int64, error) {
	args := m.Called(req)
	return args.Get(0).([]orm.DebitCardVW), args.Get(1).(int64), args.Error(2)
}
