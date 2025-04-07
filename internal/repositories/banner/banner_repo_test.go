package banner_repo

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

func TestNewBannerRepoMock(t *testing.T) {
	repo := NewBannerRepoMock()

	repo.On("CreateTable").Return(nil)
	repo.On("Insert", mock.Anything, mock.Anything).Return(nil)
	repo.On("GetList", mock.Anything).Return([]orm.Banner{}, int64(0), nil)

	_ = repo.CreateTable()
	_ = repo.Insert(nil, orm.Banner{})
	_, _, _ = repo.GetList(models.BannerGetListReq{})
}

func TestNewBannerRepo(t *testing.T) {
	db, _, _ := postgresql.InitDatabaseMock()
	defer db.Close()
	NewBannerRepo(db)
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
			repo := NewBannerRepo(db)

			if tc.errBeginTx == nil {
				mockDB.ExpectBegin()
			}
			if tc.errExec == nil {
				query := "\n\t\tCREATE TABLE IF NOT EXISTS public.banners (\n\t\t\tbanner_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),\n\t\t\tuser_id UUID DEFAULT gen_random_uuid(),\n\t\t\ttitle VARCHAR(255) NULL,\n\t\t\tdescription VARCHAR(255) NULL,\n\t\t\timage VARCHAR(255) NULL,\n\t\t\tdummy_col_11 varchar(255) DEFAULT NULL,\n\t\t\tcreated_by VARCHAR(100) NULL,\n\t\t\tcreated_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),\n\t\t\tupdated_by VARCHAR(100) NULL,\n\t\t\tupdated_date TIMESTAMPTZ NULL\n\t\t);\n\t\tcreate index banners_user_id_idx on public.banners using  btree (user_id);\n"

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
	req := orm.Banner{}

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
			repo := NewBannerRepo(db)

			mockDB.ExpectBegin()
			tx, _ := db.BeginTx(context.Background(), nil)

			query := "INSERT INTO public.banners (\n\t\tuser_id,\n\t\ttitle, \n\t\tdescription,\n\t\timage,\n\t\tcreated_by,\n\t\tcreated_date,\n\t\tupdated_by,\n\t\tupdated_date\n\t) VALUES (\n\t\t$1, $2, $3, $4, $5, $6, $7, $8\n\t);"

			params := make([]driver.Value, 8)
			params[0] = req.UserID
			params[1] = req.Title.NullString
			params[2] = req.Description.NullString
			params[3] = req.Image.NullString
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

func TestGetList(t *testing.T) {
	req := models.BannerGetListReq{
		SearchText: "test",
		UserID:     "test",
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
			repo := NewBannerRepo(db)

			var params []driver.Value
			if req.SearchText != "" {
				searchText := strings.TrimSpace(req.SearchText)
				params = append(params, `%\`+searchText+"%")
			}
			if req.UserID != "" {
				params = append(params, req.UserID)
			}

			if tc.errGet == nil {
				query := "SELECT COUNT(*) FROM public.banners WHERE true AND (title ILIKE $1) AND (user_id = $2);"
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

				query := "SELECT * FROM public.banners WHERE true AND (title ILIKE $1) AND (user_id = $2) ORDER BY test  LIMIT $3 OFFSET $4"
				mockDB.ExpectQuery(query).
					WithArgs(params...).
					WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow("test"))
			}

			repo.GetList(req)

		})
	}
}
