package debitcard_repo

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

type DebitCardRepo interface {
	CreateTable() error
	CreateTableView() error
	Insert(tx *sql.Tx, req orm.DebitCard) (cardID string, err error)
	GetList(req models.DebitCardGetListReq) (res []orm.DebitCardVW, total int64, err error)
}

type debitCardRepo struct {
	db *sqlx.DB
}

func NewDebitCardRepo(db *sqlx.DB) DebitCardRepo {
	return &debitCardRepo{
		db: db,
	}
}

func (repo *debitCardRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.debit_cards (
		card_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		name VARCHAR(100) NULL, 
		dummy_col_7 varchar(255) DEFAULT NULL,
		created_by VARCHAR(100) NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index debit_cards_user_id_idx on public.debit_cards using  btree (user_id);
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

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *debitCardRepo) CreateTableView() error {
	query := `
		CREATE OR REPLACE VIEW vw_debit_cards AS
		SELECT 
			dc.card_id,
			dc.user_id,
			dc."name",
			dcs.status,
			dcd."number",
			dcd.issuer,
			dcd2.color,
			dcd2.border_color,
			dc.created_by,
			dc.created_date,
			dc.updated_by,
			dc.updated_date
		FROM public.debit_cards dc  
		LEFT JOIN public.debit_card_status dcs
			ON dc.card_id = dcs.card_id
		LEFT JOIN public.debit_card_details dcd 
			ON dc.card_id = dcd.card_id 
		LEFT JOIN public.debit_card_design dcd2 
			ON dc.card_id = dcd2.card_id;
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

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (repo *debitCardRepo) Insert(tx *sql.Tx, req orm.DebitCard) (cardID string, err error) {
	params := make([]interface{}, 6)
	params[0] = req.UserID
	params[1] = req.Name.NullString
	params[2] = req.CreatedBy.NullString
	params[3] = req.CreatedDate
	params[4] = req.UpdatedBy.NullString
	params[5] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.debit_cards`
	col := `(
		user_id,
		name,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6
	)`
	returning := `RETURNING card_id`

	query := fmt.Sprintf(`%s %s %s %s;`, tableInsert, col, values, returning)

	err = tx.QueryRow(query, params...).Scan(&cardID)
	if postgresql.IsSQLReallyError(err) {
		return cardID, err
	}
	return cardID, nil
}

func (repo *debitCardRepo) GetList(req models.DebitCardGetListReq) (res []orm.DebitCardVW, total int64, err error) {
	params := []interface{}{}

	sl := `SELECT *`
	from := `from vw_debit_cards`

	// condition
	condition := "WHERE true"
	conditionSearch := ""
	if req.SearchText != "" {
		searchText := strings.TrimSpace(req.SearchText)
		conditionSearch = `AND (name ILIKE ?)`
		params = append(params, `%\`+searchText+"%")
	}
	conditionStatus := ""
	if req.Status != "" {
		conditionStatus = `AND (status = ?)`
		params = append(params, req.Status)
	}
	conditionUserID := ""
	if req.UserID != "" {
		conditionUserID = `AND (user_id = ?)`
		params = append(params, req.UserID)
	}
	condition = fmt.Sprintf(`%s %s %s %s`, condition, conditionSearch, conditionStatus, conditionUserID)

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
		if (req.SortBy.FieldType == reflect.String) && req.SortBy.Field != "user_id" && req.SortBy.Field != "card_id" {
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

	query := fmt.Sprintf(`%s %s %s %s %s %s;`,
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
