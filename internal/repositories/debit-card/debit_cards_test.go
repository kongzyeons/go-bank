package debitcard_repo

import (
	"context"
	"database/sql/driver"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestNewDebitCardRepoMock(t *testing.T) {
	repo := NewDebitCardRepoMock()

	repo.On("CreateTable").Return(nil)
	repo.On("CreateTableView").Return(nil)
	repo.On("Insert", mock.Anything, mock.Anything).Return("", nil)
	repo.On("GetList", mock.Anything).Return([]orm.DebitCardVW{}, int64(0), nil)

	_ = repo.CreateTable()
	_ = repo.CreateTableView()
	_, _ = repo.Insert(nil, orm.DebitCard{})
	_, _, _ = repo.GetList(models.DebitCardGetListReq{})
}

func TestNewDebitCardRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewDebitCardRepo(db)
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
			repo := NewDebitCardRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.debit_cards (\n\t\tcard_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\tname VARCHAR(100) NULL, \n\t\tdummy_col_7 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index debit_cards_user_id_idx on public.debit_cards using  btree (user_id);\n"

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

func TestCreateTableView(t *testing.T) {
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
			repo := NewDebitCardRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\t\tCREATE OR REPLACE VIEW vw_debit_cards AS\n\t\tSELECT \n\t\t\tdc.card_id,\n\t\t\tdc.user_id,\n\t\t\tdc.\"name\",\n\t\t\tdcs.status,\n\t\t\tdcd.\"number\",\n\t\t\tdcd.issuer,\n\t\t\tdcd2.color,\n\t\t\tdcd2.border_color,\n\t\t\tdc.created_by,\n\t\t\tdc.created_date,\n\t\t\tdc.updated_by,\n\t\t\tdc.updated_date\n\t\tFROM public.debit_cards dc  \n\t\tLEFT JOIN public.debit_card_status dcs\n\t\t\tON dc.card_id = dcs.card_id\n\t\tLEFT JOIN public.debit_card_details dcd \n\t\t\tON dc.card_id = dcd.card_id \n\t\tLEFT JOIN public.debit_card_design dcd2 \n\t\t\tON dc.card_id = dcd2.card_id;\n"

				mockDB.ExpectExec(query).
					WillReturnResult(sqlmock.NewResult(0, 0))
			}

			if tc.errCommit == nil {
				mockDB.ExpectCommit()
			}

			repo.CreateTableView()

		})
	}
}

func TestInsert(t *testing.T) {
	req := orm.DebitCard{}

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
			repo := NewDebitCardRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.debit_cards (\n\t\tuser_id,\n\t\tname,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6\n\t) RETURNING card_id;"

			params := make([]driver.Value, 6)
			params[0] = req.UserID
			params[1] = req.Name.NullString
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

func TestGetList(t *testing.T) {
	req := models.DebitCardGetListReq{
		SearchText: "test",
		UserID:     "test",
		Status:     "test",
		SortBy: struct {
			Field     string       "json:\"field\" example:\"updatedDate\""
			FieldType reflect.Kind "json:\"-\""
			Mode      string       "json:\"mode\" example:\"desc\""
		}{
			Field: "test",
		},
		Page:    1,
		PerPage: 1,
	}
	testCases := []struct {
		nameTest  string
		errGet    error
		total     int64
		errSelect error
	}{
		{nameTest: "test", errGet: errors.New("")},
		{nameTest: "test"},
		{nameTest: "test", total: 1, errSelect: errors.New("")},
		{nameTest: "test", total: 1},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewDebitCardRepo(db)

			var params []driver.Value
			if req.SearchText != "" {
				searchText := strings.TrimSpace(req.SearchText)
				params = append(params, `%\`+searchText+"%")
			}
			if req.Status != "" {
				params = append(params, req.Status)
			}
			if req.UserID != "" {
				params = append(params, req.UserID)
			}

			if tc.errGet == nil {
				query := "SELECT COUNT(*) from vw_debit_cards WHERE true AND (name ILIKE $1) AND (status = $2) AND (user_id = $3);"
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(tc.total))
			}

			if tc.errSelect == nil {
				// limit
				if req.PerPage > 0 {
					params = append(params, req.PerPage)
				}
				// offset
				if req.Page > 0 {
					params = append(params, (req.Page-1)*req.PerPage)
				}

				query := "SELECT * from vw_debit_cards WHERE true AND (name ILIKE $1) AND (status = $2) AND (user_id = $3) ORDER BY test  LIMIT $4 OFFSET $5;"
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow("test"))
			}

			repo.GetList(req)

		})
	}
}
