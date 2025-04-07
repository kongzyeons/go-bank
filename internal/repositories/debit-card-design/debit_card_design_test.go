package debitcarddesign_repo

import (
	"context"
	"database/sql/driver"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

func TestNewDebitCarddesignRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewDebitCarddesignRepo(db)
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
			repo := NewDebitCarddesignRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.debit_card_design (\n\t\tcard_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tcolor VARCHAR(10) NULL,\n\t\tborder_color VARCHAR(10) NULL,\n\t\tdummy_col_9 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index debit_card_design_user_id_idx on public.debit_card_design using  btree (user_id);\n"

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
	req := orm.DebitCardDesign{}

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
			repo := NewDebitCarddesignRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.debit_card_design (\n\t\tcard_id,\n\t\tuser_id,\n\t\tcolor,\n\t\tborder_color,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7, $8\n\t);"

			params := make([]driver.Value, 8)
			params[0] = req.CardID
			params[1] = req.UserID
			params[2] = req.Color.NullString
			params[3] = req.BorderColor.NullString
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
