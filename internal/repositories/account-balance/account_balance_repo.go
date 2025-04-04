package accountbalance_repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kongzyeons/go-bank/internal/models/orm"
	"github.com/kongzyeons/go-bank/pkg/postgresql"
)

type AccountBalanceRepo interface {
	CreateTable() error
	Insert(tx *sql.Tx, req orm.AccountBalance) error
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
