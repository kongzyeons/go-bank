package account_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountRepo interface {
	CreateTable() error
	CreateTableView() error
	Insert(tx *sql.Tx, req orm.Account) (accountID string, err error)
	GetList(req models.AccountGetListReq) (res []orm.AccountVW, total int64, err error)
	GetByID(accountID string) (res *orm.Account, err error)
}

type accountRepo struct {
	db *sqlx.DB
}

func NewAccountRepo(db *sqlx.DB) AccountRepo {
	return &accountRepo{
		db: db,
	}
}

func (repo *accountRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.accounts (
		account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		type VARCHAR(50) NULL,
		currency VARCHAR(10) NULL,
		account_number VARCHAR(20) NULL,
		issuer VARCHAR(100) NULL,
		dummy_col_3 varchar(255) DEFAULT NULL,
		created_by VARCHAR(100) NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index accounts_user_id_idx on public.accounts using  btree (user_id);
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

func (repo *accountRepo) CreateTableView() error {
	query := `
		CREATE OR REPLACE VIEW vw_account AS
		SELECT 
			a.account_id,
			a.user_id,
			ad.is_main_account,
			ad."name",
			a."type",
			a.account_number,
			a.issuer,
			ab.amount,
			a.currency,
			ad.color,
			ad.progress,
			a.created_by,
			a.created_date,
			a.updated_by,
			a.updated_date
		FROM public.accounts a
		LEFT JOIN public.account_balances ab 
			ON a.account_id = ab.account_id 
		LEFT JOIN public.account_details ad 
			ON a.account_id = ad.account_id;
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

func (repo *accountRepo) Insert(tx *sql.Tx, req orm.Account) (accountID string, err error) {
	params := make([]interface{}, 9)
	params[0] = req.UserID
	params[1] = req.Type.NullString
	params[2] = req.Currency.NullString
	params[3] = req.AccountNumber.NullString
	params[4] = req.Issuer.NullString
	params[5] = req.CreatedBy.NullString
	params[6] = req.CreatedDate
	params[7] = req.UpdatedBy.NullString
	params[8] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.accounts`
	col := `(
		user_id,
		type,
		currency,
		account_number,
		issuer,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	)`
	returning := `RETURNING account_id`

	query := fmt.Sprintf(`%s %s %s %s;`, tableInsert, col, values, returning)

	err = tx.QueryRow(query, params...).Scan(&accountID)
	if postgresql.IsSQLReallyError(err) {
		return accountID, err
	}
	return accountID, nil
}

func (repo *accountRepo) GetList(req models.AccountGetListReq) (res []orm.AccountVW, total int64, err error) {
	params := []interface{}{}

	sl := `SELECT *`
	from := `from vw_account`

	// condition
	condition := "WHERE true"
	conditionSearch := ""
	if req.SearchText != "" {
		searchText := strings.TrimSpace(req.SearchText)
		conditionSearch = `AND (name ILIKE ? OR type ILIKE ?)`
		params = append(params, `%\`+searchText+"%")
		params = append(params, `%\`+searchText+"%")
	}
	conditionIsmainAccount := ""
	if req.IsManinAccount != nil {
		conditionIsmainAccount = `AND (is_main_account = ?)`
		params = append(params, req.IsManinAccount)
	}
	conditionUserID := ""
	if req.UserID != "" {
		conditionUserID = `AND (user_id = ?)`
		params = append(params, req.UserID)
	}
	condition = fmt.Sprintf(`%s %s %s %s`, condition, conditionSearch, conditionIsmainAccount, conditionUserID)

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
		if (req.SortBy.FieldType == reflect.String) && req.SortBy.Field != "user_id" && req.SortBy.Field != "account_id" {
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

func (repo *accountRepo) GetByID(accountID string) (res *orm.Account, err error) {
	params := make([]interface{}, 1)
	params[0] = accountID

	sl := `SELECT *`
	from := fmt.Sprintf(`FROM %s`, "accounts")
	condition := `WHERE account_id = $1`
	query := fmt.Sprintf(`%s %s %s`, sl, from, condition)

	var result orm.Account
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}
