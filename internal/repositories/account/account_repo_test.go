package account_repo

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
	"github.com/stretchr/testify/mock"
)

func TestNewAccountRepoMock(t *testing.T) {
	repo := NewAccountRepoMock()
	repo.On("CreateTable").Return(nil)
	repo.On("CreateTableView").Return(nil)
	repo.On("Insert", mock.Anything, mock.Anything).Return("", nil)
	repo.On("GetList", mock.Anything).Return([]orm.AccountVW{}, int64(0), nil)
	repo.On("GetByID", mock.Anything).Return(&orm.Account{}, nil)

	repo.CreateTable()
	repo.CreateTableView()
	repo.Insert(nil, orm.Account{})
	repo.GetList(models.AccountGetListReq{})
	repo.GetByID("")

}

func TestNewAccountRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewAccountRepo(db)
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
			repo := NewAccountRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\tCREATE TABLE IF NOT EXISTS public.accounts (\n\t\taccount_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\ttype VARCHAR(50) NULL,\n\t\tcurrency VARCHAR(10) NULL,\n\t\taccount_number VARCHAR(20) NULL,\n\t\tissuer VARCHAR(100) NULL,\n\t\tdummy_col_3 varchar(255) DEFAULT NULL,\n\t\tcreated_by VARCHAR(100) NULL,\n\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\tupdated_by VARCHAR(100) NULL,\n\t\tupdated_date TIMESTAMPTZ NULL\n\t);\n\tcreate index accounts_user_id_idx on public.accounts using  btree (user_id);\n"
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
			repo := NewAccountRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\t\tCREATE OR REPLACE VIEW vw_account AS\n\t\tSELECT \n\t\t\ta.account_id,\n\t\t\ta.user_id,\n\t\t\tad.is_main_account,\n\t\t\tad.\"name\",\n\t\t\ta.\"type\",\n\t\t\ta.account_number,\n\t\t\ta.issuer,\n\t\t\tab.amount,\n\t\t\ta.currency,\n\t\t\tad.color,\n\t\t\tad.progress,\n\t\t\ta.created_by,\n\t\t\ta.created_date,\n\t\t\ta.updated_by,\n\t\t\ta.updated_date\n\t\tFROM public.accounts a\n\t\tLEFT JOIN public.account_balances ab \n\t\t\tON a.account_id = ab.account_id \n\t\tLEFT JOIN public.account_details ad \n\t\t\tON a.account_id = ad.account_id;\n"

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
	testCases := []struct {
		nameTest    string
		req         orm.Account
		errQueryRow error
	}{
		{nameTest: "test", errQueryRow: errors.New("")},
		{nameTest: "test"},
	}
	for _, tc := range testCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			db, mockDB, _ := postgresql.InitDatabaseMock()
			defer db.Close()
			repo := NewAccountRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.accounts (\n\t\tuser_id,\n\t\ttype,\n\t\tcurrency,\n\t\taccount_number,\n\t\tissuer,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7, $8, $9\n\t) RETURNING account_id;"

			params := make([]driver.Value, 9)
			params[0] = tc.req.UserID
			params[1] = tc.req.Type.NullString
			params[2] = tc.req.Currency.NullString
			params[3] = tc.req.AccountNumber.NullString
			params[4] = tc.req.Issuer.NullString
			params[5] = tc.req.CreatedBy.NullString
			params[6] = tc.req.CreatedDate
			params[7] = tc.req.UpdatedBy.NullString
			params[8] = tc.req.UpdatedDate.NullTime

			if tc.errQueryRow == nil {
				var defaultUUID uuid.UUID // zero value
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow(defaultUUID))
			}

			repo.Insert(tx, tc.req)

		})
	}
}

func TestGetList(t *testing.T) {
	ismain := true
	req := models.AccountGetListReq{
		SearchText:     "test",
		IsManinAccount: &ismain,
		UserID:         "test",
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
			repo := NewAccountRepo(db)

			var params []driver.Value
			if req.SearchText != "" {
				searchText := strings.TrimSpace(req.SearchText)
				params = append(params, `%\`+searchText+"%")
				params = append(params, `%\`+searchText+"%")
			}
			if req.IsManinAccount != nil {
				params = append(params, req.IsManinAccount)
			}
			if req.UserID != "" {
				params = append(params, req.UserID)
			}

			if tc.errGet == nil {
				query := "SELECT COUNT(*) from vw_account WHERE true AND (name ILIKE $1 OR type ILIKE $2) AND (is_main_account = $3) AND (user_id = $4);"
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

				query := "SELECT * from vw_account WHERE true AND (name ILIKE $1 OR type ILIKE $2) AND (is_main_account = $3) AND (user_id = $4) ORDER BY test  LIMIT $5 OFFSET $6;"
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"account_id"}).AddRow("test"))
			}

			repo.GetList(req)

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
			repo := NewAccountRepo(db)

			params := make([]driver.Value, 1)
			params[0] = "test"

			query := "SELECT * FROM accounts WHERE account_id = $1"

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
