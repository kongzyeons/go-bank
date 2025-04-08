package account_queue

import (
	"encoding/json"
	"errors"
	"reflect"
	"testing"

	"github.com/kongzyeons/go-bank/internal/models/events"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	accountbalance_repo "github.com/kongzyeons/go-bank/internal/repositories/account-balance"
	transaction_repo "github.com/kongzyeons/go-bank/internal/repositories/transaction"
	"github.com/kongzyeons/go-bank/pkg/line"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestHandle(t *testing.T) {
	testCases := []struct {
		nameTest   string
		topic      string
		eventBytes []byte
		errBeginTx error
		errCommit  error
		dataDB     *orm.AccountBalance
		errGetByID error
		errUpdate  error
		errInsert  error
	}{
		{nameTest: "test"},
		{nameTest: "test",
			topic:      reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: []byte{},
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errUpdate: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errInsert: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errCommit: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountAddMoneyEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountAddMoneyEvent{})
				return jsonStr
			}(),
			dataDB: &orm.AccountBalance{},
		},

		{nameTest: "test",
			topic:      reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: []byte{},
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
			errGetByID: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errUpdate: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errInsert: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
			dataDB:    &orm.AccountBalance{},
			errCommit: errors.New(""),
		},
		{nameTest: "test",
			topic: reflect.TypeOf(events.AccountWithldrawEvent{}).Name(),
			eventBytes: func() []byte {
				jsonStr, _ := json.Marshal(events.AccountWithldrawEvent{})
				return jsonStr
			}(),
			dataDB: &orm.AccountBalance{},
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

			accountBalanceRepo := accountbalance_repo.NewAccountBalanceRepoMock()
			transectionRepo := transaction_repo.NewTransactionRepoMock()
			lineMessage := line.NewLineAPIMock()

			lineMessage.On("SendMessage", mock.Anything).Return(
				nil,
			)

			accountBalanceRepo.On("GetByID", mock.Anything).Return(
				tc.dataDB, tc.errGetByID,
			)
			accountBalanceRepo.On("Update", mock.Anything, mock.Anything).Return(
				tc.errUpdate,
			)

			transectionRepo.On("Insert", mock.Anything, mock.Anything).Return(
				tc.errInsert,
			)

			svc := NewAccountEventHandler(
				db, accountBalanceRepo, transectionRepo, lineMessage,
			)

			svc.Handle(tc.topic, tc.eventBytes)

		})
	}
}
