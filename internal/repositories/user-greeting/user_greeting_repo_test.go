package usergreeting_repo

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

func TestNewUserGreetingRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewUserGreetingRepo(db)
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
			repo := NewUserGreetingRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\t\tCREATE TABLE IF NOT EXISTS public.user_greetings (\n\t\t\tuser_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\t\tgreeting VARCHAR(255) NULL,\n\t\t\tdummy_col_2 varchar(255) DEFAULT NULL,\n\t\t\tcreated_by VARCHAR(100) NULL,\n\t\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\t\tupdated_by VARCHAR(100) NULL,\n\t\t\tupdated_date TIMESTAMPTZ NULL\n\t\t);\n\t"

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
	req := orm.UserGreeting{}

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
			repo := NewUserGreetingRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.user_greetings (\n\t\tuser_id,\n\t\tgreeting, \n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6\n\t);"

			params := make([]driver.Value, 6)
			params[0] = req.UserID
			params[1] = req.Greeting.NullString
			params[2] = req.CreatedBy.NullString
			params[3] = req.CreatedDate
			params[4] = req.UpdatedBy.NullString
			params[5] = req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				mockDB.ExpectExec(query).
					WithArgs(params...).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			repo.Insert(tx, req)

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
			repo := NewUserGreetingRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM public.user_greetings WHERE user_id = $1"

			if tc.errGet == nil {
				if tc.errNoRows != nil {
					mockDB.ExpectQuery(query).
						WithArgs(params...).
						WillReturnError(sql.ErrNoRows)
				} else {
					mockDB.ExpectQuery(query).
						WithArgs(params...).
						WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow("test"))
				}
			}

			repo.GetByID("test")
		})
	}
}
