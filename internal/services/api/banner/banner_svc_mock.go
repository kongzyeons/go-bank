package banner_svc

import (
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

type bannerSvcMock struct {
	mock.Mock
}

func NewBannerSvcMock() *bannerSvcMock {
	return &bannerSvcMock{}
}

func (m *bannerSvcMock) GetList(req models.BannerGetListReq) response.Response[*models.BannerGetListRes] {
	args := m.Called(req)
	return args.Get(0).(response.Response[*models.BannerGetListRes])
}
