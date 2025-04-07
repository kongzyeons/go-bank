package accountflag_repo

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestNewAccountFlagRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewAccountFlagRepo(db)
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
			repo := NewAccountFlagRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.account_flags (\n \t\tflag_id SERIAL PRIMARY KEY, \n\t\taccount_id UUID DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tflag_type VARCHAR(50) NOT NULL,\n\t\tflag_value VARCHAR(30) NOT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index account_flags_account_id_idx on public.account_flags using  btree (account_id);\n\tcreate index account_flags_user_id_idx on public.account_flags using  btree (user_id);\n"

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
	req := orm.AccountFlag{}

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
			repo := NewAccountFlagRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.account_flags (\n\t\taccount_id,\n\t\tuser_id,\n\t\tflag_type,\n\t\tflag_value,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7, $8\n\t);"

			params := make([]driver.Value, 8)
			params[0] = req.AccountID
			params[1] = req.UserID
			params[2] = req.FlagType
			params[3] = req.FlagValue
			params[4] = req.CreatedBy.NullString
			params[5] = req.CreatedDate
			params[6] = req.UpdatedBy.NullString
			params[7] = req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, req)

		})
	}
}
