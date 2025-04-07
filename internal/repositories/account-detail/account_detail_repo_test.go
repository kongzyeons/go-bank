package accountdetail_repo

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

func TestNewAccountDetailRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewAccountDetailRepo(db)
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
			repo := NewAccountDetailRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.account_details (\n\t\taccount_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tname varchar(100) NULL,\n\t\tcolor VARCHAR(10) NULL,\n\t\tis_main_account bool DEFAULT false NOT NULL,\n\t\tprogress BIGINT NULL,\n\t\tdummy_col_5 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index account_details_user_id_idx on public.account_details using  btree (user_id);\n\t"

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
	req := orm.AccountDetail{}

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
			repo := NewAccountDetailRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.account_details (\n\t\taccount_id,\n\t\tuser_id,\n\t\tname,\n\t\tcolor,\n\t\tis_main_account,\n\t\tprogress,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7, $8, $9, $10\n\t);"

			params := make([]driver.Value, 10)
			params[0] = req.AccountID
			params[1] = req.UserID
			params[2] = req.Name.NullString
			params[3] = req.Color.NullString
			params[4] = req.IsManinAccount
			params[5] = req.Progress.NullInt64
			params[6] = req.CreatedBy.NullString
			params[7] = req.CreatedDate
			params[8] = req.UpdatedBy.NullString
			params[9] = req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, req)

		})
	}
}

func TestUpdate(t *testing.T) {
	req := orm.AccountDetail{}
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
			repo := NewAccountDetailRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "UPDATE account_details SET \n\t\tname = $1,\n\t\tcolor = $2,\n\t\tis_main_account = $3,\n\t\tprogress = $4,\n\t\tupdated_by = $5,\n\t\tupdated_date = $6\n\t WHERE account_id = $7;"

			params := make([]driver.Value, 7)
			params[0] = req.Name.NullString
			params[1] = req.Color.NullString
			params[2] = req.IsManinAccount
			params[3] = req.Progress.NullInt64
			params[4] = req.UpdatedBy.NullString
			params[5] = req.UpdatedDate.NullTime
			params[6] = req.AccountID

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
			repo := NewAccountDetailRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM account_details WHERE account_id = $1"

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
