package user_svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/stretchr/testify/mock"
)

func TestNewUserSvcMock(t *testing.T) {
	repo := NewUserSvcMock()

	// Prepare mock response
	mockResponse := response.Response[*models.UserGetGeetingRes]{}
	repo.On("GetGeeting", mock.Anything).Return(mockResponse)

	// Test the mock method
	_ = repo.GetGeeting(models.UserGetGeetingReq{})
}

func TestNewUserSvc(t *testing.T) {
	NewUserSvc(nil, nil)
}

func TestGetGeeting(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.UserGetGeetingReq
		key        string
		value      string
		data       string
		errGet     error
		dataDB     *orm.UserGreeting
		errGetByID error
		errSet     error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.UserGetGeetingReq{
				UserID:   "test",
				Username: "test",
			},
			value: "test",
		},
		{nameTest: "test",
			req: models.UserGetGeetingReq{
				UserID:   "test",
				Username: "test",
			},
			value: func() string {
				jsonStr, _ := json.Marshal(models.UserGetGeetingRes{})
				return string(jsonStr)
			}(),
		},
		{nameTest: "test",
			req: models.UserGetGeetingReq{
				UserID:   "test",
				Username: "test",
			},
			value:      "test",
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			req: models.UserGetGeetingReq{
				UserID:   "test",
				Username: "test",
			},
			value:  "test",
			dataDB: &orm.UserGreeting{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {

			userGreetingRepo := usergreeting_repo.NewUserGreetingRepoMock()
			userGreetingRepo.On("GetByID", mock.Anything).Return(
				tc.dataDB, tc.errGetByID,
			)

			rc, redisMock := redismock.NewClientMock()
			reqJson, _ := json.Marshal(tc.req)
			key := fmt.Sprintf("userSvc::%s", string(reqJson))

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

			svc := NewUserSvc(
				rc, userGreetingRepo,
			)

			svc.GetGeeting(tc.req)

		})
	}
}
