package transaction_repo

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

func TestNewTransactionRepoMock(t *testing.T) {
	repo := NewTransactionRepoMock()

	repo.On("CreateTable").Return(nil)
	repo.On("Insert", mock.Anything, mock.Anything).Return(nil)

	_ = repo.CreateTable()
	_ = repo.Insert(nil, orm.Transaction{})
}

func TestNewTransactionRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewTransactionRepo(db)
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
			repo := NewTransactionRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.transactions (\n\t\ttransaction_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tname VARCHAR(100) NULL,\n\t\timage VARCHAR(255) NULL,\n\t\tisBank bool DEFAULT false NOT NULL,\n\t\tdummy_col_6 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC')\n\t);\n\tcreate index transactions_user_id_idx on public.transactions using  btree (user_id);\n"

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
	req := orm.Transaction{}

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
			repo := NewTransactionRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.transactions (\n\t\tuser_id,\n\t\tname,\n\t\timage,\n\t\tisBank,\n\t\tcreated_by,\n\t\tcreated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6\n\t);"

			params := make([]driver.Value, 6)
			params[0] = req.UserID
			params[1] = req.Name
			params[2] = req.Image.NullString
			params[3] = req.IsBank
			params[4] = req.CreatedBy.NullString
			params[5] = req.CreatedDate

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, req)

		})
	}
}
