package banner_svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	banner_repo "github.com/kongzyeons/go-bank/internal/repositories/banner"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestNewBannerSvcMock(t *testing.T) {
	repo := NewBannerSvcMock()

	// Prepare mock response
	mockResponse := response.Response[*models.BannerGetListRes]{}
	repo.On("GetList", mock.Anything).Return(mockResponse)

	// Test the mock method
	_ = repo.GetList(models.BannerGetListReq{})
}

func TestGetList(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.BannerGetListReq
		key        string
		value      string
		data       string
		errGet     error
		errSet     error
		dataDB     []orm.Banner
		total      int64
		errGetList error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.BannerGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
				SortBy: struct {
					Field     string       "json:\"field\" example:\"updatedDate\""
					FieldType reflect.Kind "json:\"-\""
					Mode      string       "json:\"mode\" example:\"desc\""
				}{
					Field: "test",
					Mode:  "test",
				},
			},
		},
		{nameTest: "test",
			req: models.BannerGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
				SortBy: struct {
					Field     string       "json:\"field\" example:\"updatedDate\""
					FieldType reflect.Kind "json:\"-\""
					Mode      string       "json:\"mode\" example:\"desc\""
				}{
					Field: "test",
					Mode:  "asc",
				},
			},
		},
		{nameTest: "test",
			req: models.BannerGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
			errGetList: errors.New(""),
		},
		{nameTest: "test",
			req: models.BannerGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
		},
		{nameTest: "test",
			req: models.BannerGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
			dataDB: []orm.Banner{{}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			rc, redisMock := redismock.NewClientMock()
			reqJson, _ := json.Marshal(tc.req)
			key := fmt.Sprintf("debitcardSvc::%s", string(reqJson))

			if tc.errGet != nil {
				redisMock.ExpectGet(key).SetErr(tc.errGet)
			}
			if tc.value != "" {
				redisMock.ExpectGet(key).SetVal(tc.value)
			}
			if tc.errSet != nil {
				redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			} else {
				redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			}

			bannerRepo := banner_repo.NewBannerRepoMock()
			bannerRepo.On("GetList", mock.Anything).Return(
				tc.dataDB, tc.total, tc.errGetList,
			)

			svc := NewBannerSvc(
				rc, bannerRepo,
			)
			svc.GetList(tc.req)

		})
	}
}
