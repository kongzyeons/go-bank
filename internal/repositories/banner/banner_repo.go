package banner_repo

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type BannerRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.Banner) error
	GetList(req models.BannerGetListReq) (res []orm.Banner, total int64, err error)
}

type bannerRepo struct {
	db *sqlx.DB
}

func NewBannerRepo(db *sqlx.DB) BannerRepo {
	return &bannerRepo{db: db}
}

func (repo *bannerRepo) CreateTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS public.banners (
			banner_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			user_id UUID DEFAULT gen_random_uuid(),
			title VARCHAR(255) NULL,
			description VARCHAR(255) NULL,
			image VARCHAR(255) NULL,
			dummy_col_11 varchar(255) DEFAULT NULL,
			created_by VARCHAR(100) NULL,
			created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
			updated_by VARCHAR(100) NULL,
			updated_date TIMESTAMPTZ NULL
		);
		create index banners_user_id_idx on public.banners using  btree (user_id);
`
	tx, err := repo.db.BeginTx(context.Background(), nil)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(query)
	if err != nil {
		return err
	}

	//commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (repo *bannerRepo) Insert(tx *sql.Tx, req orm.Banner) error {
	params := make([]interface{}, 8)
	params[0] = req.UserID
	params[1] = req.Title.NullString
	params[2] = req.Description.NullString
	params[3] = req.Image.NullString
	params[4] = req.CreatedBy.NullString
	params[5] = req.CreatedDate
	params[6] = req.UpdatedBy.NullString
	params[7] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.banners`
	col := `(
		user_id,
		title, 
		description,
		image,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}

func (repo *bannerRepo) GetList(req models.BannerGetListReq) (res []orm.Banner, total int64, err error) {
	params := []interface{}{}

	sl := `SELECT *`
	from := `FROM public.banners`

	// condition
	condition := "WHERE true"
	conditionSearch := ""
	if req.SearchText != "" {
		searchText := strings.TrimSpace(req.SearchText)
		conditionSearch = `AND (title ILIKE ?)`
		params = append(params, `%\`+searchText+"%")
	}
	conditionUserID := ""
	if req.UserID != "" {
		conditionUserID = `AND (user_id = ?)`
		params = append(params, req.UserID)
	}

	condition = fmt.Sprintf(`%s %s %s`, condition, conditionSearch, conditionUserID)

	queryCount := fmt.Sprintf(`SELECT COUNT(*) %s %s;`, from, condition)
	err = repo.db.Get(&total, repo.db.Rebind(queryCount), params...)
	if postgresql.IsSQLReallyError(err) {
		return res, total, err
	}
	if total == 0 {
		return res, total, nil
	}

	// order
	order := ""
	if req.SortBy.Field != "" {
		if (req.SortBy.FieldType == reflect.String) && req.SortBy.Field != "user_id" && req.SortBy.Field != "banner_id" {
			order = fmt.Sprintf(`ORDER BY %s COLLATE "th-TH-x-icu" %s`, req.SortBy.Field, req.SortBy.Mode)
		} else {
			order = fmt.Sprintf(`ORDER BY %s %s`, req.SortBy.Field, req.SortBy.Mode)
		}
	}
	// limit
	limit := ""
	if req.PerPage > 0 {
		limit = `LIMIT ?`
		params = append(params, req.PerPage)
	}
	// offset
	offset := ""
	if req.Page > 0 {
		offset = `OFFSET ?`
		params = append(params, (req.Page-1)*req.PerPage)
	}

	query := fmt.Sprintf(`%s %s %s %s %s %s`,
		sl, from,
		condition,
		order, limit, offset,
	)
	err = repo.db.Select(&res, repo.db.Rebind(query), params...)
	if postgresql.IsSQLReallyError(err) {
		return res, total, err
	}

	return res, total, err
}
