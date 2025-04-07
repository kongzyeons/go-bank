package accountbalance_repo

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestNewAccountRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewaccountBalanceRepo(db)
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
			repo := NewaccountBalanceRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.account_balances (\n\t\taccount_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tamount DECIMAL(15,2) NULL,\n\t\tdummy_col_4 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index account_balances_user_id_idx on public.account_balances using  btree (user_id);\n"

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
	testCases := []struct {
		nameTest    string
		req         orm.AccountBalance
		errQueryRow error
	}{
		{nameTest: "test", errQueryRow: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewaccountBalanceRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.account_balances (\n\t\taccount_id,\n\t\tuser_id,\n\t\tamount,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7\n\t);"

			params := make([]driver.Value, 7)
			params[0] = tc.req.AccountID
			params[1] = tc.req.UserID
			params[2] = tc.req.Amount.NullFloat64
			params[3] = tc.req.CreatedBy.NullString
			params[4] = tc.req.CreatedDate
			params[5] = tc.req.UpdatedBy.NullString
			params[6] = tc.req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, tc.req)

		})
	}
}

func TestUpdate(t *testing.T) {
	req := orm.AccountBalance{}
	testCases := []struct {
		nameTest string
		errExec  error
	}{
		{nameTest: "test", errExec: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewaccountBalanceRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "UPDATE public.account_balances SET \n\t\tamount = $1,\n\t\tupdated_by = $2,\n\t\tupdated_date = $3\n\t WHERE account_id = $4;"

			params := make([]driver.Value, 4)
			params[0] = req.Amount.NullFloat64
			params[1] = req.UpdatedBy.NullString
			params[2] = req.UpdatedDate.NullTime
			params[3] = req.AccountID

			if tc.errExec == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Update(tx, req)

		})
	}
}

func TestGetByID(t *testing.T) {
	testCases := []struct {
		nameTest  string
		errGet    error
		errNoRows error
	}{
		{nameTest: "test", errGet: errors.New("")},
		{nameTest: "test", errNoRows: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewaccountBalanceRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM account_balances WHERE account_id = $1"

			if tc.errGet == nil {
				if tc.errNoRows != nil {
					mockDB.ExpectQuery(query).
						WithArgs(params...).
						WillReturnError(sql.ErrNoRows)
				} else {
					mockDB.ExpectQuery(query).
						WithArgs(params...).
						WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow("test"))
				}
			}

			repo.GetByID("test")
		})
	}
}
