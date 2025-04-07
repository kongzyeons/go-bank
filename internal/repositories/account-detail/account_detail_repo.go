package accountdetail_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountDetailRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.AccountDetail) error
	Update(tx *sql.Tx, req orm.AccountDetail) error
	GetByID(accountID string) (res *orm.AccountDetail, err error)
}

type accountDetailRepo struct {
	db *sqlx.DB
}

func NewAccountDetailRepo(db *sqlx.DB) AccountDetailRepo {
	return &accountDetailRepo{
		db: db,
	}
}

func (repo *accountDetailRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.account_details (
		account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		name varchar(100) NULL,
		color VARCHAR(10) NULL,
		is_main_account bool DEFAULT false NOT NULL,
		progress BIGINT NULL,
		dummy_col_5 varchar(255) DEFAULT NULL,
		created_by VARCHAR(100) NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index account_details_user_id_idx on public.account_details using  btree (user_id);
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
func (repo *accountDetailRepo) Insert(tx *sql.Tx, req orm.AccountDetail) error {
	params := make([]interface{}, 10)
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

	tableInsert := `INSERT INTO public.account_details`
	col := `(
		account_id,
		user_id,
		name,
		color,
		is_main_account,
		progress,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}

func (repo *accountDetailRepo) Update(tx *sql.Tx, req orm.AccountDetail) error {
	params := make([]interface{}, 7)
	params[0] = req.Name.NullString
	params[1] = req.Color.NullString
	params[2] = req.IsManinAccount
	params[3] = req.Progress.NullInt64
	params[4] = req.UpdatedBy.NullString
	params[5] = req.UpdatedDate.NullTime
	params[6] = req.AccountID

	tableUpdate := fmt.Sprintf("UPDATE %s SET", "account_details")

	update := `
		name = $1,
		color = $2,
		is_main_account = $3,
		progress = $4,
		updated_by = $5,
		updated_date = $6
	`

	where := `WHERE account_id = $7`

	query := fmt.Sprintf("%s %s %s;", tableUpdate, update, where)
	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}
	return nil

}
func (repo *accountDetailRepo) GetByID(accountID string) (res *orm.AccountDetail, err error) {
	params := make([]interface{}, 1)
	params[0] = accountID

	sl := `SELECT *`
	from := fmt.Sprintf(`FROM %s`, "account_details")
	condition := `WHERE account_id = $1`
	query := fmt.Sprintf(`%s %s %s`, sl, from, condition)

	var result orm.AccountDetail
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}
