package auth_service

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	user_repo "github.com/kongzyeons/go-bank/internal/repositories/user"
	usergreeting_repo "github.com/kongzyeons/go-bank/internal/repositories/user-greeting"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthSvcMock(t *testing.T) {
	repo := NewAuthSvcMock()

	// Prepare mock responses
	mockResponseAny := response.Response[any]{}
	mockResponseLogin := response.Response[*models.AuthLoginRes]{}
	mockResponsePing := response.Response[*models.AuthPingRes]{}
	mockResponseRefresh := response.Response[*models.AuthRefreshRes]{}

	repo.On("Register", mock.Anything).Return(mockResponseAny)
	repo.On("Login", mock.Anything).Return(mockResponseLogin)
	repo.On("Ping", mock.Anything).Return(mockResponsePing)
	repo.On("Refresh", mock.Anything).Return(mockResponseRefresh)
	repo.On("Logout", mock.Anything).Return(mockResponseAny)

	// Test the mock methods
	_ = repo.Register(models.AuthRegisterReq{})
	_ = repo.Login(models.AuthLoginReq{})
	_ = repo.Ping(models.AuthPingReq{})
	_ = repo.Refresh(models.AuthRefreshReq{})
	_ = repo.Logout(models.AuthLogoutReq{})
}

func TestRegister(t *testing.T) {
	testCases := []struct {
		nameTest                  string
		req                       models.AuthRegisterReq
		dataDB                    *orm.User
		errGetUnique              error
		errBeginTx                error
		erruserRepoInsert         error
		erruserGreetingRepoInsert error
		errCommit                 error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "test12",
			},
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
			errGetUnique: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
			dataDB: &orm.User{},
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
			erruserRepoInsert: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
			erruserGreetingRepoInsert: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
			errCommit: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthRegisterReq{
				Username: "test",
				Password: "123456",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			rc, _ := redismock.NewClientMock()

			userRepo := user_repo.NewUserRepoMock()
			userGreetingRepo := usergreeting_repo.NewUserGreetingRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			userRepo.On("GetUnique", mock.Anything).Return(
				tc.dataDB, tc.errGetUnique,
			)

			userRepo.On("Insert", mock.Anything, mock.Anything).Return(
				"test", tc.erruserRepoInsert,
			)

			userGreetingRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.erruserGreetingRepoInsert,
			)
			svc := NewAuthSvc(
				db, rc,
				userRepo, userGreetingRepo,
				transectionRepo,
			)

			svc.Register(tc.req)
		})
	}
}

func TestLogin(t *testing.T) {
	testCases := []struct {
		nameTest     string
		req          models.AuthLoginReq
		dataDB       *orm.User
		errGetUnique error
		key          string
		value        string
		errGet       error
		errSet       error
		errBeginTx   error
		errInsert    error
		errCommit    error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "test12",
			},
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
			errGetUnique: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
			dataDB: &orm.User{},
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
			dataDB: &orm.User{
				UserID:   "test",
				Password: types.NewNullString("123456"),
			},
			errGet: errors.New(""),
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
			dataDB: &orm.User{
				UserID:   "test",
				Password: types.NewNullString("123456"),
			},
			value: "test",
		},
		{nameTest: "test",
			req: models.AuthLoginReq{
				Username: "test",
				Password: "123456",
			},
			dataDB: &orm.User{
				UserID:   "test",
				Password: types.NewNullString("123456"),
			},
			errGet: redis.Nil,
			value:  "test",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			rc, redisMock := redismock.NewClientMock()
			key := fmt.Sprintf("authSvc::%s", "test")

			if tc.errGet != nil {
				redisMock.ExpectGet(key).SetErr(tc.errGet)
			}
			if tc.value != "" {
				redisMock.ExpectGet(key).SetVal(tc.value)
			}
			if tc.errSet != nil {
				redisMock.ExpectSet(key, tc.value, time.Hour*24).SetErr(tc.errSet)
			} else {
				redisMock.ExpectSet(key, tc.value, time.Hour*24).SetVal("ok")
			}

			userRepo := user_repo.NewUserRepoMock()
			userGreetingRepo := usergreeting_repo.NewUserGreetingRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			userRepo.On("GetUnique", mock.Anything).Return(
				tc.dataDB, tc.errGetUnique,
			)

			svc := NewAuthSvc(
				db, rc,
				userRepo, userGreetingRepo,
				transectionRepo,
			)

			svc.Login(tc.req)
		})
	}
}
