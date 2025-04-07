package user_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type userRepoMock struct {
	mock.Mock
}

func NewUserRepoMock() *userRepoMock {
	return &userRepoMock{}
}

func (m *userRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *userRepoMock) Insert(tx *sql.Tx, req orm.User) (string, error) {
	args := m.Called(tx, req)
	return args.String(0), args.Error(1)
}

func (m *userRepoMock) InsertMock(tx *sql.Tx, req orm.User) (string, error) {
	args := m.Called(tx, req)
	return args.String(0), args.Error(1)
}

func (m *userRepoMock) GetByID(id string) (*orm.User, error) {
	args := m.Called(id)
	return args.Get(0).(*orm.User), args.Error(1)
}

func (m *userRepoMock) GetUnique(name string) (*orm.User, error) {
	args := m.Called(name)
	return args.Get(0).(*orm.User), args.Error(1)
}
