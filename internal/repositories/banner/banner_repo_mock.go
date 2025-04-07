package banner_repo

import (
	"database/sql"

	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/stretchr/testify/mock"
)

type bannerRepoMock struct {
	mock.Mock
}

func NewBannerRepoMock() *bannerRepoMock {
	return &bannerRepoMock{}
}

func (m *bannerRepoMock) CreateTable() error {
	args := m.Called()
	return args.Error(0)
}

func (m *bannerRepoMock) Insert(tx *sql.Tx, req orm.Banner) error {
	args := m.Called(tx, req)
	return args.Error(0)
}

func (m *bannerRepoMock) GetList(req models.BannerGetListReq) ([]orm.Banner, int64, error) {
	args := m.Called(req)
	return args.Get(0).([]orm.Banner), args.Get(1).(int64), args.Error(2)
}
