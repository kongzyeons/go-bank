package debitcardstatus_repo

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestNewDebitCardStatusRepoMock(t *testing.T) {
	repo := NewDebitCardStatusRepoMock()

	repo.On("CreateTable").Return(nil)
	repo.On("Insert", mock.Anything, mock.Anything).Return(nil)

	_ = repo.CreateTable()
	_ = repo.Insert(nil, orm.DebitCardStatus{})
}

func TestNewDebitCardStatusRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewDebitCardStatusRepo(db)
}

func TestCreateTable(t *testing.T) {
	testCases := []struct {
		nameTest   string
		errBeginTx error
		errExec    error
		errCommit  error
	}{
		{nameTest: "test", errExec: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewDebitCardStatusRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.debit_card_status (\n\t\tcard_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tstatus VARCHAR(20) NULL,\n\t\tdummy_col_8 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index debit_card_status_user_id_idx on public.debit_card_status using  btree (user_id);\n"

				mockDB.ExpectExec(query).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			repo.CreateTable()

		})
	}
}

func TestInsert(t *testing.T) {
	req := orm.DebitCardStatus{}

	testCases := []struct {
		nameTest    string
		errQueryRow error
	}{
		{nameTest: "test", errQueryRow: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewDebitCardStatusRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.debit_card_status (\n\t\tcard_id,\n\t\tuser_id,\n\t\tstatus,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7\n\t);"

			params := make([]driver.Value, 7)
			params[0] = req.CardID
			params[1] = req.UserID
			params[2] = req.Status.NullString
			params[3] = req.CreatedBy.NullString
			params[4] = req.CreatedDate
			params[5] = req.UpdatedBy.NullString
			params[6] = req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, req)

		})
	}
}
