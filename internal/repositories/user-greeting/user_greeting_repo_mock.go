package usergreeting_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type userGreetingRepoMock struct {
	mock.Mock
}

func NewUserGreetingRepoMock() *userGreetingRepoMock {
	return &userGreetingRepoMock{}
}

func (m *userGreetingRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *userGreetingRepoMock) Insert(tx *sql.Tx, req orm.UserGreeting) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *userGreetingRepoMock) GetByID(id string) (*orm.UserGreeting, error) {
	args := m.Called(id)
	return args.Get(0).(*orm.UserGreeting), args.Error(1)
}
