package accountbalance_repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountBalanceRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.AccountBalance) error
	Update(tx *sql.Tx, req orm.AccountBalance) error
	GetByID(accountID string) (res *orm.AccountBalance, err error)
}

type accountBalanceRepo struct {
	db *sqlx.DB
}

func NewaccountBalanceRepo(db *sqlx.DB) AccountBalanceRepo {
	return &accountBalanceRepo{
		db: db,
	}
}

func (repo *accountBalanceRepo) CreateTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS public.account_balances (
		account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
		user_id UUID DEFAULT gen_random_uuid(),
		amount DECIMAL(15,2) NULL,
		created_by VARCHAR(100) NOT NULL,
		created_date TIMESTAMPTZ NOT NULL DEFAULT (NOW() AT TIME ZONE 'UTC'),
		updated_by VARCHAR(100) NULL,
		updated_date TIMESTAMPTZ NULL
	);
	create index account_balances_user_id_idx on public.account_balances using  btree (user_id);
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
func (repo *accountBalanceRepo) Insert(tx *sql.Tx, req orm.AccountBalance) error {
	params := make([]interface{}, 7)
	params[0] = req.AccountID
	params[1] = req.UserID
	params[2] = req.Amount.NullFloat64
	params[3] = req.CreatedBy
	params[4] = req.CreatedDate
	params[5] = req.UpdatedBy.NullString
	params[6] = req.UpdatedDate.NullTime

	tableInsert := `INSERT INTO public.account_balances`
	col := `(
		account_id,
		user_id,
		amount,
		created_by,
		created_date,
		updated_by,
		updated_date
	)`
	values := `VALUES (
		$1, $2, $3, $4, $5, $6, $7
	)`

	query := fmt.Sprintf(`%s %s %s;`, tableInsert, col, values)

	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}

func (repo *accountBalanceRepo) Update(tx *sql.Tx, req orm.AccountBalance) error {
	params := make([]interface{}, 4)
	params[0] = req.Amount.NullFloat64
	params[1] = req.UpdatedBy.NullString
	params[2] = req.UpdatedDate.NullTime
	params[3] = req.AccountID

	tableUpdate := fmt.Sprintf("UPDATE %s SET", "public.account_balances")

	update := `
		amount = $1,
		updated_by = $2,
		updated_date = $3
	`

	where := `WHERE account_id = $4`

	query := fmt.Sprintf("%s %s %s;", tableUpdate, update, where)
	_, err := tx.Exec(query, params...)
	if postgresql.IsSQLReallyError(err) {
		return err
	}

	return nil
}

func (repo *accountBalanceRepo) GetByID(accountID string) (res *orm.AccountBalance, err error) {
	params := make([]interface{}, 1)
	params[0] = accountID

	sl := `SELECT *`
	from := fmt.Sprintf(`FROM %s`, "account_balances")
	condition := `WHERE account_id = $1`
	query := fmt.Sprintf(`%s %s %s`, sl, from, condition)

	var result orm.AccountBalance
	err = repo.db.Get(&result, query, params...)
	if postgresql.IsSQLReallyError(err) {
		return nil, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	return &result, err
}
