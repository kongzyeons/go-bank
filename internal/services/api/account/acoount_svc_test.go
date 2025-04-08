package account_svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/internal/queues"
	account_repo "github.com/kongzyeons/go-bank/internal/repositories/account"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	accountdetail_repo "github.com/kongzyeons/go-bank/internal/repositories/account-detail"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	"github.com/kongzyeons/go-bank/internal/utils/response"
	"github.com/kongzyeons/go-bank/internal/utils/types"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestNewAccountSvcMock(t *testing.T) {
	repo := NewAccountSvcMock()

	repo.On("GetList", mock.Anything).Return(response.Response[*models.AccountGetListRes]{})
	repo.On("Edit", mock.Anything).Return(response.Response[*models.AccountEditRes]{})
	repo.On("GetQrcode", mock.Anything).Return(response.Response[*models.AccountGetQrcodeRes]{})
	repo.On("SetIsmain", mock.Anything).Return(response.Response[*models.AccountSetIsmainRes]{})
	repo.On("AddMoney", mock.Anything).Return(response.Response[any]{})
	repo.On("Withdrawl", mock.Anything).Return(response.Response[any]{})

	_ = repo.GetList(models.AccountGetListReq{})
	_ = repo.Edit(models.AccountEditReq{})
	_ = repo.GetQrcode(models.AccountGetQrcodeReq{})
	_ = repo.SetIsmain(models.AccountSetIsmainReq{})
	_ = repo.AddMoney(models.AccountAddMoneyReq{})
	_ = repo.Withdrawl(models.AccountWithdrawlReq{})
}

func TestGetList(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.AccountGetListReq
		key        string
		data       string
		value      string
		errGet     error
		errSet     error
		dataDB     []orm.AccountVW
		total      int64
		errGetlist error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountGetListReq{
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
			req: models.AccountGetListReq{
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
			req: models.AccountGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
			value:      "test",
			errGetlist: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
			value: "test",
		},
		{nameTest: "test",
			req: models.AccountGetListReq{
				UserID:  "test",
				Page:    1,
				PerPage: 1,
			},
			dataDB: []orm.AccountVW{{}},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, _, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			// if tc.errBeginTx == nil {
			// 	mockDB.ExpectBegin()
			// }
			// if tc.errCommit == nil {
			// 	mockDB.ExpectCommit()
			// }

			rc, redisMock := redismock.NewClientMock()
			defer rc.Close()
			reqJson, _ := json.Marshal(tc.req)
			key := fmt.Sprintf("accountSvc::%s", string(reqJson))
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

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountRepo.On("GetList", mock.Anything).Return(
				tc.dataDB, tc.total, tc.errGetlist,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.GetList(tc.req)

		})

	}
}

func TestEdit(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.AccountEditReq
		errGetByID error
		dataDB     *orm.AccountDetail
		errBeginTx error
		errUpdate  error
		errInsert  error
		errCommit  error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			dataDB: &orm.AccountDetail{UserID: ""},
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			dataDB:    &orm.AccountDetail{UserID: "test"},
			errUpdate: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			dataDB:    &orm.AccountDetail{UserID: "test"},
			errInsert: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			dataDB:    &orm.AccountDetail{UserID: "test"},
			errCommit: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountEditReq{
				AccountID: "test",
				UserID:    "test",
				Username:  "test",
				Name:      "test",
				Color:     "test",
			},
			dataDB: &orm.AccountDetail{UserID: "test"},
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
			// defer rc.Close()
			// reqJson, _ := json.Marshal(tc.req)
			// key := fmt.Sprintf("accountSvc::%s", string(reqJson))
			// if tc.errGet != nil {
			// 	redisMock.ExpectGet(key).SetErr(tc.errGet)
			// }
			// if tc.value != "" {
			// 	redisMock.ExpectGet(key).SetVal(tc.value)
			// }
			// if tc.errSet != nil {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			// } else {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			// }

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountDetailRepo.On("GetByID", mock.Anything).Return(
				tc.dataDB, tc.errGetByID,
			)

			accountDetailRepo.On("Update", mock.Anything, mock.Anything).Return(
				tc.errUpdate,
			)

			transectionRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errInsert,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.Edit(tc.req)

		})

	}
}

func TestGetQrcode(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.AccountGetQrcodeReq
		dataDB     *orm.Account
		errGetByID error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountGetQrcodeReq{
				AccountID: "test",
				UserID:    "test",
			},
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountGetQrcodeReq{
				AccountID: "test",
				UserID:    "test",
			},
		},
		{nameTest: "test",
			req: models.AccountGetQrcodeReq{
				AccountID: "test",
				UserID:    "test",
			},
			dataDB: &orm.Account{},
		},
		{nameTest: "test",
			req: models.AccountGetQrcodeReq{
				AccountID: "test",
				UserID:    "test",
			},
			dataDB: &orm.Account{
				UserID: "test",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, _, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			// if tc.errBeginTx == nil {
			// 	mockDB.ExpectBegin()
			// }
			// if tc.errCommit == nil {
			// 	mockDB.ExpectCommit()
			// }

			rc, _ := redismock.NewClientMock()
			// defer rc.Close()
			// reqJson, _ := json.Marshal(tc.req)
			// key := fmt.Sprintf("accountSvc::%s", string(reqJson))
			// if tc.errGet != nil {
			// 	redisMock.ExpectGet(key).SetErr(tc.errGet)
			// }
			// if tc.value != "" {
			// 	redisMock.ExpectGet(key).SetVal(tc.value)
			// }
			// if tc.errSet != nil {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			// } else {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			// }

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountRepo.On("GetByID", mock.Anything).Return(
				tc.dataDB, tc.errGetByID,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.GetQrcode(tc.req)

		})

	}
}

func TestSetIsmain(t *testing.T) {
	testCases := []struct {
		nameTest    string
		req         models.AccountSetIsmainReq
		dataDB1     *orm.AccountDetail
		errGetByID1 error
		dataDB2     *orm.AccountDetail
		errGetByID2 error
		errBeginTx  error
		errUpdate   error
		errInsert   error
		errCommit   error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test",
				Username:        "test",
				UserID:          "test",
			},
			errGetByID1: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test",
				Username:        "test",
				UserID:          "test",
			},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				IsManinAccount: true,
			},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test2",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				UserID: "test",
			},
			errGetByID2: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test2",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				UserID: "test",
			},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test2",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				UserID: "test",
			},
			dataDB2: &orm.AccountDetail{},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test2",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				UserID: "test",
			},
			dataDB2: &orm.AccountDetail{
				IsManinAccount: true,
			},
		},
		{nameTest: "test",
			req: models.AccountSetIsmainReq{
				AccountID:       "test",
				AccountIDIsmain: "test2",
				Username:        "test",
				UserID:          "test",
			},
			dataDB1: &orm.AccountDetail{
				UserID: "test",
			},
			dataDB2: &orm.AccountDetail{
				UserID:         "test",
				IsManinAccount: true,
			},
			errUpdate: errors.New(""),
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
			// defer rc.Close()
			// reqJson, _ := json.Marshal(tc.req)
			// key := fmt.Sprintf("accountSvc::%s", string(reqJson))
			// if tc.errGet != nil {
			// 	redisMock.ExpectGet(key).SetErr(tc.errGet)
			// }
			// if tc.value != "" {
			// 	redisMock.ExpectGet(key).SetVal(tc.value)
			// }
			// if tc.errSet != nil {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			// } else {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			// }

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountDetailRepo.On("GetByID", strings.TrimSpace(tc.req.AccountID)).Return(
				tc.dataDB1, tc.errGetByID1,
			)

			accountDetailRepo.On("GetByID", strings.TrimSpace(tc.req.AccountIDIsmain)).Return(
				tc.dataDB2, tc.errGetByID2,
			)

			accountDetailRepo.On("Update", mock.Anything, mock.Anything).Return(
				tc.errUpdate,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.SetIsmain(tc.req)

		})

	}
}

func TestAddMoney(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.AccountAddMoneyReq
		errGetByID error
		dataDB     *orm.AccountBalance
		errProduce error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountAddMoneyReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountAddMoneyReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
		},
		{nameTest: "test",
			req: models.AccountAddMoneyReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{},
		},
		{nameTest: "test",
			req: models.AccountAddMoneyReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{
				UserID: "test",
			},
			errProduce: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountAddMoneyReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{
				UserID: "test",
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, _, _ := postgresql.InitDatabaseMock()
			defer db.Close()

			rc, _ := redismock.NewClientMock()
			// defer rc.Close()
			// reqJson, _ := json.Marshal(tc.req)
			// key := fmt.Sprintf("accountSvc::%s", string(reqJson))
			// if tc.errGet != nil {
			// 	redisMock.ExpectGet(key).SetErr(tc.errGet)
			// }
			// if tc.value != "" {
			// 	redisMock.ExpectGet(key).SetVal(tc.value)
			// }
			// if tc.errSet != nil {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			// } else {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			// }

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountBalanceRepo.On("GetByID", strings.TrimSpace(tc.req.AccountID)).Return(
				tc.dataDB, tc.errGetByID,
			)

			eventProducer.On("Produce", mock.Anything).Return(
				tc.errProduce,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.AddMoney(tc.req)

		})

	}
}

func TestWithdrawl(t *testing.T) {
	testCases := []struct {
		nameTest   string
		req        models.AccountWithdrawlReq
		errGetByID error
		dataDB     *orm.AccountBalance
		errProduce error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
		},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{},
		},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{
				UserID: "test",
			},
		},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{
				UserID: "test",
				Amount: types.NewNullFloat64(2),
			},
			errProduce: errors.New(""),
		},
		{nameTest: "test",
			req: models.AccountWithdrawlReq{
				UserID:    "test",
				Username:  "test",
				AccountID: "test",
				Ammount:   1,
				Currency:  "test",
			},
			dataDB: &orm.AccountBalance{
				UserID: "test",
				Amount: types.NewNullFloat64(2),
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, _, _ := postgresql.InitDatabaseMock()
			defer db.Close()

			rc, _ := redismock.NewClientMock()
			// defer rc.Close()
			// reqJson, _ := json.Marshal(tc.req)
			// key := fmt.Sprintf("accountSvc::%s", string(reqJson))
			// if tc.errGet != nil {
			// 	redisMock.ExpectGet(key).SetErr(tc.errGet)
			// }
			// if tc.value != "" {
			// 	redisMock.ExpectGet(key).SetVal(tc.value)
			// }
			// if tc.errSet != nil {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetErr(tc.errSet)
			// } else {
			// 	redisMock.ExpectSet(key, tc.data, time.Second*10).SetVal("ok")
			// }

			eventProducer := queues.NewEventProducerMock()
			accountRepo := account_repo.NewAccountRepoMock()
			accountDetailRepo := accountdetail_repo.NewAccountDetailRepoMock()
			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()

			accountBalanceRepo.On("GetByID", strings.TrimSpace(tc.req.AccountID)).Return(
				tc.dataDB, tc.errGetByID,
			)

			eventProducer.On("Produce", mock.Anything).Return(
				tc.errProduce,
			)

			svc := NewAccountSvc(
				db, rc, eventProducer,
				accountRepo, accountDetailRepo,
				accountBalanceRepo, transectionRepo,
			)

			svc.Withdrawl(tc.req)

		})

	}
}
