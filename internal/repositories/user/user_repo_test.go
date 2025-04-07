package user_repo

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestNewAccountRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewUserRepo(db)
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
			repo := NewUserRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\t\tCREATE TABLE IF NOT EXISTS public.users (\n\t\t\tuser_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\t\tname VARCHAR(100) NULL,\n\t\t\tpassword VARCHAR(6) NULL,\n\t\t\tdummy_col_1 varchar(255) DEFAULT NULL,\n\t\t\tcreated_by VARCHAR(100) NULL,\n\t\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\t\tupdated_by VARCHAR(100) NULL,\n\t\t\tupdated_date TIMESTAMPTZ NULL\n\t\t);\n\t\tcreate index users_name_idx on public.users using  btree (name);\n"

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

func TestInsertMock(t *testing.T) {
	req := orm.User{}

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
			repo := NewUserRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.users (\n\t\tuser_id,\n\t\tname,\n\t\tpassword, \n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7\n\t) RETURNING user_id;"

			var defaultUUID uuid.UUID // zero value
			params := make([]driver.Value, 7)
			params[0] = defaultUUID.String()
			params[1] = req.Name
			params[2] = req.Password.NullString
			params[3] = req.CreatedBy.NullString
			params[4] = req.CreatedDate
			params[5] = req.UpdatedBy.NullString
			params[6] = req.UpdatedDate.NullTime
			if tc.errQueryRow == nil {
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(defaultUUID))
			}

			repo.InsertMock(tx, req)

		})
	}
}

func TestInsert(t *testing.T) {
	req := orm.User{}

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
			repo := NewUserRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.users (\n\t\tname,\n\t\tpassword, \n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6\n\t) RETURNING user_id;"

			params := make([]driver.Value, 6)
			params[0] = req.Name
			params[1] = req.Password.NullString
			params[2] = req.CreatedBy.NullString
			params[3] = req.CreatedDate
			params[4] = req.UpdatedBy.NullString
			params[5] = req.UpdatedDate.NullTime
			if tc.errQueryRow == nil {
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow("test"))
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
			repo := NewUserRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM public.users WHERE user_id = $1"

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

func TestGetUnique(t *testing.T) {
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
			repo := NewUserRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM public.users WHERE name = $1"

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

			repo.GetUnique("test")
		})
	}
}
